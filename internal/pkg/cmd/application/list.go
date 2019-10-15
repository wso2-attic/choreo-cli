/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package application

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

func NewListCommand(cliContext runtime.CliContext) *cobra.Command {

	const cmdList = "list"
	cmd := &cobra.Command{
		Use:     cmdList,
		Short:   "List applications",
		Example: common.GetAbsoluteCommandName(cmdApplication, cmdList),
		Args:    cobra.NoArgs,
		Run:     runListAppCommand(cliContext),
	}
	return cmd
}

func runListAppCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		if !client.IsUserLoggedIn(cliContext) {
			common.ExitWithErrorMessage(cliContext.Out(), "Please login first")
		}

		listApps(cliContext)
	}
}

func listApps(cliContext runtime.CliContext) {

	backendUrl := cliContext.EnvConfig().GetStringOrDefault(client.BackendUrl, client.EnvConfigs[client.BackendUrl])
	accessToken := cliContext.UserConfig().GetStringOrDefault(client.AccessToken, client.UserConfigs[client.AccessToken])
	req, err := client.NewRequest(backendUrl, accessToken, "GET", pathApplications, nil)

	if err != nil {
		log.Print("Error creating post request for listing applications: ", err)
		return
	}

	skipVerify, _ := strconv.ParseBool(cliContext.
		EnvConfig().GetStringOrDefault(client.SkipVerify, client.EnvConfigs[client.SkipVerify]))
	httpClient := client.NewClient(skipVerify)

	resp, err := httpClient.Do(req)
	if err == nil {
		showListAppsResult(cliContext.Out(), resp)
	} else {
		log.Print("Error making post request for listing applications: ", err)
	}
}

func showListAppsResult(consoleWriter io.Writer, resp *http.Response) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading json body: ", err)
		return
	}
	if resp.StatusCode == http.StatusOK {
		var apps Applications
		err := json.Unmarshal(body, &apps)
		if err != nil {
			log.Print("Error converting json into applications: ", err)
		}
		printer := tableprinter.New(os.Stdout)
		printer.Print(apps)
	} else {
		common.PrintInfo(consoleWriter, "Error listing applications.")
		fmt.Println("Error: ", string(body))
	}
	err = resp.Body.Close()
	if err != nil {
		log.Print(err)
	}
}
