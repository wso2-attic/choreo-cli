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

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

func newLoginCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "login",
		Short:   "Login to " + productName,
		Example: getAbsoluteCommandName("login"),
		Args:    cobra.NoArgs,
		Run:     runLogin,
	}
}

func runLogin(cmd *cobra.Command, args []string) {
	authCode := getAuthCode()
	accessToken := getAccessToken(authCode)
	persistAccessToken(accessToken)
}

func persistAccessToken(accessToken string) {
	panic("persistAccessToken Method not implemented")
}

func getAccessToken(authCode string) string {
	log.Println("Received: " + authCode)
	panic("getAccessToken Method not implemented")
}

func getAuthCode() string {
	log.Printf("Requesting credentials through browser")
	const callBackDefaultPort = 8888
	const callbackUrlContext = "/auth"
	const callBackUrl = "http://localhost:%d" + callbackUrlContext
	const idpUrlPrefix = "https://id.choreo.io/sdk/fidp-select?redirectUrl="

	codeServicePort := getFirstOpenPort(callBackDefaultPort)
	redirectUrl := url.QueryEscape(fmt.Sprintf(callBackUrl, codeServicePort))
	hubAuthUrl := idpUrlPrefix + redirectUrl
	log.Println("Callback URL: " + hubAuthUrl)

	authCodeChannel := make(chan string)
	go listenForAuthCode(getBindAddress(codeServicePort), callbackUrlContext, authCodeChannel)
	authCode := <-authCodeChannel

	return authCode
}

func getBindAddress(port int) string {
	return ":" + strconv.Itoa(port)
}

func getFirstOpenPort(startingPort int) int {
	port := startingPort
	for connection, err := net.Dial("tcp", getBindAddress(port)); err == nil; port++ {
		_ = connection.Close()
	}
	return port
}

func listenForAuthCode(addrString string, callbackUrlContext string, authCodeChannel chan<- string) {
	mux := http.NewServeMux()
	server := http.Server{Addr: addrString, Handler: mux}
	mux.HandleFunc(callbackUrlContext, func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			exitWithError("Error parsing received query parameters", err)
		}

		code := request.Form.Get("code")

		if code == "" {
			writer.WriteHeader(http.StatusBadRequest)
			exitWithErrorMessage("Blank auth code received from IDP")
		}

		writer.WriteHeader(http.StatusOK)
		authCodeChannel <- code
	})

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		exitWithError("Error while initializing auth code accepting service", err)
	}
}
