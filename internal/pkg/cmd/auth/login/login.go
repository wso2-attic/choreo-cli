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
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/config"
	"golang.org/x/oauth2"
)

func NewLoginCommand(cliContext runtime.CliContext) *cobra.Command {
	return &cobra.Command{
		Use:     "login",
		Short:   "Login to " + common.ProductName,
		Example: common.GetAbsoluteCommandName("login"),
		Args:    cobra.NoArgs,
		Run:     createLoginFunction(cliContext),
	}
}

func createLoginFunction(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	getEnvConfig := createEnvConfigReader(cliContext)
	setUserConfig := config.CreateConfigWriter(cliContext.UserConfig())
	consoleWriter := cliContext.Out()

	return func(cmd *cobra.Command, args []string) {
		codeServicePort := common.GetFirstOpenPort(callBackDefaultPort)
		oauthClient := initOauthClient(callbackUrlContext, codeServicePort, getEnvConfig)
		authCodeChannel, server := startAuthCodeReceivingService(codeServicePort, oauthClient, setUserConfig, consoleWriter)
		openBrowserForAuthentication(consoleWriter, oauthClient)
		<-authCodeChannel
		stopAuthCodeServer(server)

		common.PrintInfo(consoleWriter, "Successfully logged in to "+common.ProductName+".")
	}
}

func getAccessToken(authCode string, conf *oauthClient) (string, error) {
	token, err := conf.exchange(authCode)

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

func startAuthCodeReceivingService(port int, oauth2Conf *oauthClient, setUserConfig config.SetConfig, consoleWriter io.Writer) (<-chan bool, *http.Server) {
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
	common.SendBrowserResponse(consoleWriter, httpWriter, http.StatusInternalServerError, "CLI Login",
		"Login to Choreo failed due to an internal error.", "Please try again.")
}

func exchangeAuthCodeForToken(code string, oauth2Conf *oauthClient, httpWriter http.ResponseWriter, consoleWriter io.Writer, setUserConfig config.SetConfig) error {
	token, err := getAccessToken(code, oauth2Conf)
	if err != nil {
		return err
	}
	setUserConfig(client.AccessToken, token)
	common.SendBrowserResponse(consoleWriter, httpWriter, http.StatusOK, "CLI Login",
		"Login to Choreo is successful.", "Please return to the CLI.")
	return nil
}

func openBrowserForAuthentication(consoleWriter io.Writer, client *oauthClient) {
	hubAuthUrl := client.authCodeURL("state")
	if err := common.OpenBrowser(hubAuthUrl); err != nil {
		common.ExitWithError(consoleWriter, "Couldn't open browser for "+common.ProductName+" login", err)
	}
}

func initOauthClient(context string, port int, getEnvConfig config.GetConfig) *oauthClient {
	callBackUrlTemplate := "http://localhost:%d" + context
	redirectUrl := fmt.Sprintf(callBackUrlTemplate, port)

	skipVerify, _ := strconv.ParseBool(common.GetStringOrDefault(getEnvConfig, client.SkipVerify, client.EnvConfigs[client.SkipVerify]))
	return createOathClient(getEnvConfig(clientId), redirectUrl, getEnvConfig(authUrl), getEnvConfig(tokenUrl), skipVerify)
}

func listenForAuthCode(server *http.Server, consoleWriter io.Writer) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		common.ExitWithError(consoleWriter, "Error while initializing auth code accepting service", err)
	}
}

type oauthClient struct {
	oauthConf *oauth2.Config
	ctx context.Context
}

func (c *oauthClient) authCodeURL(state string) string {
	return c.oauthConf.AuthCodeURL(state)
}

func (c *oauthClient) exchange(code string) (*oauth2.Token, error) {
	return c.oauthConf.Exchange(c.ctx, code)
}

func createOathClient(clientId, redirectUrl, authUrl, tokenUrl string, insecure bool) *oauthClient {
	oauthConf := &oauth2.Config{
		ClientID:     clientId,
		Endpoint:     oauth2.Endpoint{
			AuthURL:   authUrl,
			TokenURL:  tokenUrl,
			AuthStyle: 1,
		},
		RedirectURL:  redirectUrl,
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	httpClient := &http.Client{Transport: tr}
	ctx := context.WithValue(context.TODO(), oauth2.HTTPClient, httpClient)

	return &oauthClient{
		oauthConf: oauthConf,
		ctx:       ctx,
	}
}
