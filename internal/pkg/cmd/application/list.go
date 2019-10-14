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
	"io/ioutil"
	"net/http"
	"os"

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

	req, err := client.NewRequest(cliContext, "GET", pathApplications, nil)

	if err != nil {
		common.PrintError(cliContext.DebugOut(), "Error creating post request for listing applications: ", err)
		common.ExitWithErrorMessage(cliContext.Out(), "Internal error occurred")
	}

	httpClient := client.NewClient(cliContext)

	resp, err := httpClient.Do(req)
	if err == nil {
		showListAppsResult(cliContext, resp)
	} else {
		common.PrintError(cliContext.DebugOut(), "Error making post request for listing applications: ", err)
		common.ExitWithErrorMessage(cliContext.Out(), "Error communicating with server")
	}
}

func showListAppsResult(console runtime.ConsoleWriterHolder, resp *http.Response) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.PrintError(console.DebugOut(), "Error reading json body: ", err)
		common.ExitWithErrorMessage(console.Out(), "Internal error occurred")
	}
	if resp.StatusCode == http.StatusOK {
		var apps Applications
		err := json.Unmarshal(body, &apps)
		if err != nil {
			common.PrintError(console.DebugOut(), "Error converting json into applications: ", err)
			common.ExitWithErrorMessage(console.Out(), "Internal error occurred")
		}
		printer := tableprinter.New(os.Stdout)
		printer.Print(apps)
	} else {
		common.PrintErrorMessage(console.DebugOut(), "Received error message: "+string(body))
		common.PrintErrorMessage(console.Out(), "Error listing applications")
	}
	err = resp.Body.Close()
	if err != nil {
		common.PrintError(console.DebugOut(), "Error closing response body: ", err)
	}
}
