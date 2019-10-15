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
	"context"
	"encoding/json"
	"fmt"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

func PerformGithubAuthorization(cliContext runtime.CliContext) bool {

	consoleWriter := cliContext.Out()
	getEnvConfig := config.CreateConfigReader(cliContext.EnvConfig(), client.EnvConfigs)

	state, err := obtainState(cliContext)
	if state == "" {
		if err != nil {
			common.PrintErrorMessage(consoleWriter, err.Error())
		}
		common.ExitWithErrorMessage(consoleWriter, "Error while initiating authorization flow")
	}

	localServerPort := common.GetFirstOpenPort(localServerBasePort)
	localServerUrl := "http://localhost:" + strconv.Itoa(localServerPort) + localServerPath
	authRequestUrl := getEnvConfig(client.BackendUrl) + backendOauthRequestPath +
		"?user_redirect=" + url.QueryEscape(localServerUrl) +
		"&state=" + state

	doneChannel := make(chan bool)
	mux := http.NewServeMux()
	server := &http.Server{Addr: common.GetLocalBindAddress(localServerPort), Handler: mux}
	mux.HandleFunc(localServerPath, callbackHandler(consoleWriter, getEnvConfig, doneChannel))

	go func() {
		// returns ErrServerClosed on graceful close
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			common.ExitWithError(consoleWriter, "Error while ListenAndServe: ", err)
		}
	}()

	common.PrintInfo(consoleWriter, "You will now be taken to your browser for authentication")
	err = common.OpenBrowser(authRequestUrl)
	if err != nil {
		common.ExitWithError(consoleWriter, "Error opening browser: ", err)
	}

	isAuthorized := <-doneChannel
	shutdownServer(consoleWriter, server)

	return isAuthorized
}

func callbackHandler(consoleWriter io.Writer, getEnvConfig func(key string) string, doneChannel chan bool) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		var isAuthorized = false
		queryParts, err := url.ParseQuery(request.URL.RawQuery)
		if err == nil && queryParts["is_authorized"][0] == "true" {
			isAuthorized = true
		}

		title := "GitHub Authorization"
		message := "Please return to the CLI."
		if isAuthorized {
			common.SendBrowserResponse(consoleWriter, responseWriter, http.StatusOK, title, "Authorization successful !", message)
		} else {
			common.SendBrowserResponse(consoleWriter, responseWriter, http.StatusInternalServerError, title, "Authorization failed !", message)
		}

		doneChannel <- isAuthorized
	}
}

func shutdownServer(consoleWriter io.Writer, server *http.Server) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		common.PrintError(consoleWriter, "Error shutting down the local server. Reason: ", err)
	}
}

func obtainState(cliContext runtime.CliContext) (string, error) {
	backendUrl := cliContext.EnvConfig().GetStringOrDefault(client.BackendUrl, client.EnvConfigs[client.BackendUrl])
	accessToken := cliContext.UserConfig().GetStringOrDefault(client.AccessToken, client.UserConfigs[client.AccessToken])
	req, err := client.NewRequest(backendUrl, accessToken, "GET", backendOauthStatePath, nil)
	if err != nil {
		return "", err
	}

	skipVerify, _ := strconv.ParseBool(cliContext.
		EnvConfig().GetStringOrDefault(client.SkipVerify, client.EnvConfigs[client.SkipVerify]))
	httpClient := client.NewClient(skipVerify)
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var stateObj struct {
		State string `json:"state"`
	}
	if resp.StatusCode == http.StatusOK {
		err := json.Unmarshal(body, &stateObj)
		if err != nil {
			return "", err
		}
	} else {
		err = fmt.Errorf("error response received for state request: %s", string(body))
		return "", err
	}
	err = resp.Body.Close()
	if err != nil {
		return "", err
	}
	return stateObj.State, nil
}
