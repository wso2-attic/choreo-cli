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

	getEnvConfig := config.CreateConfigReader(cliContext.EnvConfig(), client.EnvConfigs)

	state, err := obtainState(cliContext)
	if state == "" {
		if err != nil {
			common.PrintErrorMessage(cliContext.DebugOut(), err.Error())
		}
		common.ExitWithErrorMessage(cliContext.Out(), "Error while initiating authorization flow")
	}

	localServerPort := common.GetFirstOpenPort(localServerBasePort)
	localServerUrl := "http://localhost:" + strconv.Itoa(localServerPort) + localServerPath
	authRequestUrl := getEnvConfig(client.BackendUrl) + backendOauthRequestPath +
		"?user_redirect=" + url.QueryEscape(localServerUrl) +
		"&state=" + state

	doneChannel := make(chan bool)
	mux := http.NewServeMux()
	server := &http.Server{Addr: common.GetLocalBindAddress(localServerPort), Handler: mux}
	mux.HandleFunc(localServerPath, callbackHandler(cliContext, getEnvConfig, doneChannel))

	go func() {
		// returns ErrServerClosed on graceful close
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			common.ExitWithError(cliContext.Out(), "Error while ListenAndServe: ", err)
		}
	}()

	common.PrintInfo(cliContext.Out(), "You will now be taken to your browser for authentication")
	err = common.OpenBrowser(authRequestUrl)
	if err != nil {
		common.PrintErrorMessage(cliContext.DebugOut(), err.Error())
		common.ExitWithErrorMessage(cliContext.Out(), "Error occurred while opening browser")
	}

	isAuthorized := <-doneChannel
	shutdownServer(cliContext, server)

	return isAuthorized
}

func callbackHandler(console runtime.ConsoleWriterHolder, getEnvConfig func(key string) string, doneChannel chan bool) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		var isAuthorized = false
		queryParts, err := url.ParseQuery(request.URL.RawQuery)
		if err == nil && queryParts["is_authorized"][0] == "true" {
			isAuthorized = true
		}

		title := "GitHub Authorization"
		message := "Please return to the CLI."
		if isAuthorized {
			common.SendBrowserResponse(console.Out(), responseWriter, http.StatusOK, title, "Authorization successful !", message)
		} else {
			common.SendBrowserResponse(console.Out(), responseWriter, http.StatusInternalServerError, title, "Authorization failed !", message)
		}

		doneChannel <- isAuthorized
	}
}

func shutdownServer(console runtime.ConsoleWriterHolder, server *http.Server) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		common.PrintErrorMessage(console.DebugOut(), err.Error())
		common.PrintErrorMessage(console.Out(), "Error shutting down the local server")
	}
}

func obtainState(cliContext runtime.CliContext) (string, error) {

	req, err := client.NewRequest(cliContext, "GET", backendOauthStatePath, nil)
	if err != nil {
		return "", err
	}
	httpClient := client.NewClient(cliContext)
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
