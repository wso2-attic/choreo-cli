/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package github

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/config"
	"golang.org/x/oauth2"
)

func StartAuthFlow(cliConfig config.Config) bool {

	getEnvConfig := createEnvironmentConfigReader(cliConfig)

	localServerPort := common.GetFirstOpenPort(localServerBasePort)
	localServerUrl := "http://localhost:" + strconv.Itoa(localServerPort) + localServerPath

	conf := &oauth2.Config{
		ClientID: getEnvConfig(clientID),
		Scopes:   []string{"repo"},
		Endpoint: oauth2.Endpoint{
			AuthURL: githubAuthUrl,
		},
		RedirectURL: localServerUrl,
	}

	state := common.GetRandomString(10)
	authUrl := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)

	doneChannel := make(chan bool)
	mux := http.NewServeMux()
	server := &http.Server{Addr: common.GetLocalBindAddress(localServerPort), Handler: mux}
	mux.HandleFunc(localServerPath, callbackHandler(getEnvConfig, doneChannel, state))

	go func() {
		// returns ErrServerClosed on graceful close
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("Error while ListenAndServe: ", err)
		}
	}()

	fmt.Println("You will now be taken to your browser for authentication")
	err := common.OpenBrowser(authUrl)
	if err != nil {
		log.Fatal("Error opening browser: ", err)
	}

	isAuthorized := <-doneChannel
	time.Sleep(2 * time.Second)
	shutdownServer(server)

	return isAuthorized
}

func callbackHandler(getEnvConfig func(key string) string, doneChannel chan bool, state string) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		var isAuthorized = false
		var code, receivedState = "", ""
		queryParts, err := url.ParseQuery(request.URL.RawQuery)
		if err == nil {
			code = queryParts["code"][0]
			receivedState = queryParts["state"][0]
		}

		var browserMsg string
		if code == "" {
			log.Print("Error obtaining auth code")
			browserMsg = createBrowserContent(false)
		} else if receivedState != state {
			log.Print("Error: Sent and received states are different")
			browserMsg = createBrowserContent(false)
		} else {
			statusCode := doPostBackend(getEnvConfig, code, state)
			if statusCode == http.StatusOK {
				isAuthorized = true
			}
			browserMsg = createBrowserContent(isAuthorized)
		}

		responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err = fmt.Fprintf(responseWriter, browserMsg)
		if err != nil {
			log.Print("Error displaying browser content: ", err)
		}
		flusher, ok := responseWriter.(http.Flusher)
		if !ok {
			log.Print("Error in casting the flusher", err)
		}
		flusher.Flush()

		doneChannel <- isAuthorized
	}
}

func shutdownServer(server *http.Server) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Print("Error shutting down: ", err)
	}
}

func createBrowserContent(isAuthorized bool) string {

	var msg string
	if isAuthorized == true {
		msg = "<h2>Authorization successful !</h2>"
	} else {
		msg = "<h2>Authorization failed !</h2>"
	}
	msg = msg + "<h2>Please return to the CLI</h2>"

	return msg
}

func doPostBackend(getEnvConfig func(key string) string, code string, state string) int {

	postUrl := getEnvConfig(backendUrl) + backendGithubAuthPath
	jsonStr := []byte(`{"code":"` + code + `", "state":"` + state + `"}`)
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Print("Error creating post request to submit auth code: ", err)
		return -1
	}
	req.Header.Set("Content-Type", "application/json")

	skipVerify, _ := strconv.ParseBool(getEnvConfig(skipVerify))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("==> ", string(body))

		err = resp.Body.Close()
		if err != nil {
			log.Print(err)
		}
		return resp.StatusCode
	} else {
		log.Print("Error making post request to submit auth code: ", err)
		return -1
	}
}
