/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package login

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	cmdContext "github.com/wso2/choreo-cli/internal/pkg/cmd/context"
	"github.com/wso2/choreo-cli/internal/pkg/config"
	"golang.org/x/oauth2"
)

func NewLoginCommand(cliContext cmdContext.CliContext, cliConfig config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "login",
		Short:   "Login to " + common.ProductName,
		Example: common.GetAbsoluteCommandName("login"),
		Args:    cobra.NoArgs,
		Run:     createLoginFunction(cliContext, cliConfig),
	}
}

func createLoginFunction(cliContext cmdContext.CliContext, cliConfig config.Config) func(cmd *cobra.Command, args []string) {
	getEnvConfig := createEnvConfigReader(cliConfig)
	setUserConfig := config.CreateUserConfigWriter(cliConfig)
	consoleWriter := cliContext.Out()

	return func(cmd *cobra.Command, args []string) {
		codeServicePort := common.GetFirstOpenPort(callBackDefaultPort)
		oauth2Conf := createOauth2Conf(callbackUrlContext, codeServicePort, getEnvConfig)
		authCodeChannel, server := startAuthCodeReceivingService(codeServicePort, oauth2Conf, setUserConfig, consoleWriter)
		openBrowserForAuthentication(consoleWriter, oauth2Conf)
		<-authCodeChannel
		stopAuthCodeServer(server)

		common.PrintInfo("Successfully logged in to " + common.ProductName + ".")
	}
}

func getAccessToken(authCode string, conf *oauth2.Config) (string, error) {
	token, err := conf.Exchange(context.Background(), authCode)

	if err == nil {
		return token.AccessToken, nil
	} else {
		return "", err
	}
}

func stopAuthCodeServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down the authcode receiving server: %s", err)
	}
}

func startAuthCodeReceivingService(port int, oauth2Conf *oauth2.Config, setUserConfig config.SetConfig, consoleWriter io.Writer) (<-chan bool, *http.Server) {
	oauthDone := make(chan bool)

	mux := http.NewServeMux()
	server := &http.Server{Addr: common.GetLocalBindAddress(port), Handler: mux}
	mux.HandleFunc(callbackUrlContext, func(httpWriter http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			sendErrorToBrowser(consoleWriter, httpWriter)
			common.PrintError(consoleWriter, "Login to "+common.ProductName+" failed due to an error parsing the received query parameters", err)
			oauthDone <- false
			return
		}

		code := request.Form.Get("code")

		if code == "" {
			sendErrorToBrowser(consoleWriter, httpWriter)
			common.PrintErrorMessage(consoleWriter, "Login to Choreo failed due to receiving a blank auth code from the IDP")
			oauthDone <- false
			return
		} else {
			if err := exchangeAuthCodeForToken(code, oauth2Conf, httpWriter, consoleWriter, setUserConfig); err != nil {
				sendErrorToBrowser(consoleWriter, httpWriter)
				common.PrintError(consoleWriter, "Could not exchange auth code for an access token", err)
				oauthDone <- false
				return
			}
		}

		oauthDone <- true
	})

	go listenForAuthCode(server, consoleWriter)

	return oauthDone, server
}

func sendErrorToBrowser(consoleWriter io.Writer, httpWriter http.ResponseWriter) {
	message := "Login to Choreo failed due to an internal error. Please try again."
	sendBrowserResponse(consoleWriter, httpWriter, http.StatusInternalServerError, message)
}

func exchangeAuthCodeForToken(code string, oauth2Conf *oauth2.Config, writer http.ResponseWriter, consoleWriter io.Writer, setUserConfig config.SetConfig) error {
	token, err := getAccessToken(code, oauth2Conf)
	if err != nil {
		return err
	}
	setUserConfig(client.AccessToken, token)
	sendBrowserResponse(consoleWriter, writer, http.StatusOK, "Login to Choreo is successful. Please return to the CLI.")
	return nil
}

func sendBrowserResponse(consoleWriter io.Writer, httpWriter http.ResponseWriter, status int, message string) {
	httpWriter.WriteHeader(status)
	httpWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	content := ` <!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>CLI Login</title>
  </head>
  <body>
    <h2>%s</h2>
  </body>
</html> `

	if _, err := fmt.Fprintf(httpWriter, content, message); err != nil {
		common.PrintError(consoleWriter, "Error while sending response to auth code redirect", err)
	}
}

func openBrowserForAuthentication(consoleWriter io.Writer, conf *oauth2.Config) {
	hubAuthUrl := conf.AuthCodeURL("state")
	if err := common.OpenBrowser(hubAuthUrl); err != nil {
		common.ExitWithError(consoleWriter, "Couldn't open browser for "+common.ProductName+" login", err)
	}
}

func createOauth2Conf(context string, port int, getEnvConfig config.GetConfig) *oauth2.Config {
	callBackUrlTemplate := "http://localhost:%d" + context
	redirectUrl := fmt.Sprintf(callBackUrlTemplate, port)

	conf := &oauth2.Config{
		ClientID:    getEnvConfig(clientId),
		RedirectURL: redirectUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:   getEnvConfig(authUrl),
			TokenURL:  getEnvConfig(tokenUrl),
			AuthStyle: 1,
		},
	}

	return conf
}

func listenForAuthCode(server *http.Server, consoleWriter io.Writer) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		common.ExitWithError(consoleWriter,"Error while initializing auth code accepting service", err)
	}
}
