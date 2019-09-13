/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein is strictly forbidden, unless permitted by WSO2 in accordance with
 * the WSO2 Commercial License available at http://wso2.com/licenses. For specific
 * language governing the permissions and limitations under this license,
 * please see the license as well as any agreement you’ve entered into with
 * WSO2 governing the purchase of this software and any associated services.
 */

package login

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo/components/cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo/components/cli/internal/pkg/config"
	"golang.org/x/oauth2"
)

func NewLoginCommand(cliConfig config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "login",
		Short:   "Login to " + common.ProductName,
		Example: common.GetAbsoluteCommandName("login"),
		Args:    cobra.NoArgs,
		Run:     createLoginFunction(cliConfig),
	}
}

func createLoginFunction(cliConfig config.Config) func(cmd *cobra.Command, args []string) {
	getEnvConfig := config.GetEnvironmentConfigReader(cliConfig, envConfigs)
	setUserConfig := config.GetUserConfigWriter(cliConfig,userConfigs)

	return func(cmd *cobra.Command, args []string) {
		authCode, oauth2Conf := getAuthCode(getEnvConfig)

		token, err := getAccessToken(authCode, oauth2Conf)
		if err != nil {
			common.ExitWithError("Could not get an access token", err)
		}

		setUserConfig(accessToken, token)
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

func getAuthCode(getEnvConfig config.GetConfig) (string, *oauth2.Config) {
	codeServicePort := common.GetFirstOpenPort(callBackDefaultPort)

	authCodeChannel := startAuthCodeReceivingService(codeServicePort, callbackUrlContext)

	oauth2Conf := openBrowserForAuthentication(callbackUrlContext, codeServicePort, getEnvConfig)

	authCode := <-authCodeChannel

	return authCode, oauth2Conf
}

func startAuthCodeReceivingService(port int, urlContext string) chan string {
	authCodeChannel := make(chan string)
	go listenForAuthCode(common.GetLocalBindAddress(port), urlContext, authCodeChannel)
	return authCodeChannel
}

func openBrowserForAuthentication(context string, port int, getEnvConfig config.GetConfig) *oauth2.Config {
	callBackUrlTemplate := "http://localhost:%d" + context
	redirectUrl := fmt.Sprintf(callBackUrlTemplate, port)
	conf := createOauth2Conf(redirectUrl, getEnvConfig)
	hubAuthUrl := getHubUrl(conf)
	if err := common.OpenBrowser(hubAuthUrl); err != nil {
		common.ExitWithError("Couldn't open browser for " + common.ProductName + "  login", err)
	}

	return conf
}

func getHubUrl(conf*oauth2.Config) string {
	return conf.AuthCodeURL("state")
}

func createOauth2Conf(redirectUrl string, getEnvConfig config.GetConfig) *oauth2.Config {
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

func listenForAuthCode(addrString string, callbackUrlContext string, authCodeChannel chan<- string) {
	mux := http.NewServeMux()
	server := http.Server{Addr: addrString, Handler: mux}
	mux.HandleFunc(callbackUrlContext, func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			common.ExitWithError("Error parsing received query parameters", err)
		}

		code := request.Form.Get("code")

		if code == "" {
			writer.WriteHeader(http.StatusBadRequest)
			common.ExitWithErrorMessage("Blank auth code received from IDP")
		}

		writer.WriteHeader(http.StatusOK)
		authCodeChannel <- code
	})

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		common.ExitWithError("Error while initializing auth code accepting service", err)
	}
}
