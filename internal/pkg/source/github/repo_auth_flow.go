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
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

func StartAuthFlow(cliConfig config.Config) bool {

	getEnvConfig := config.CreateEnvironmentConfigReader(cliConfig, client.EnvConfigs)

	localServerPort := common.GetFirstOpenPort(localServerBasePort)
	localServerUrl := "http://localhost:" + strconv.Itoa(localServerPort) + localServerPath
	authRequestUrl := getEnvConfig(client.BackendUrl) + backendOauthRequestPath + "?user_redirect=" + url.QueryEscape(localServerUrl)

	doneChannel := make(chan bool)
	mux := http.NewServeMux()
	server := &http.Server{Addr: common.GetLocalBindAddress(localServerPort), Handler: mux}
	mux.HandleFunc(localServerPath, callbackHandler(getEnvConfig, doneChannel))

	go func() {
		// returns ErrServerClosed on graceful close
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			common.ExitWithError("Error while ListenAndServe: ", err)
		}
	}()

	common.PrintInfo("You will now be taken to your browser for authentication")
	err := common.OpenBrowser(authRequestUrl)
	if err != nil {
		common.ExitWithError("Error opening browser: ", err)
	}

	isAuthorized := <-doneChannel
	time.Sleep(2 * time.Second)
	shutdownServer(server)

	if isAuthorized {
		common.PrintInfo("Authorization successful")
	} else {
		common.PrintErrorMessage("Authorization failed")
	}
	return isAuthorized
}

func callbackHandler(getEnvConfig func(key string) string, doneChannel chan bool) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		var isAuthorized = false
		queryParts, err := url.ParseQuery(request.URL.RawQuery)
		if err == nil && queryParts["is_authorized"][0] == "true" {
			isAuthorized = true
		}

		title := "Github Authorization"
		message := "Please return to the CLI."
		if isAuthorized {
			common.SendBrowserResponse(responseWriter, http.StatusOK, title, "Authorization successful !", message)
		} else {
			common.SendBrowserResponse(responseWriter, http.StatusInternalServerError, title, "Authorization failed !", message)
		}

		doneChannel <- isAuthorized
	}
}

func shutdownServer(server *http.Server) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		common.PrintError("Error shutting down the local server. Reason: ", err)
	}
}
