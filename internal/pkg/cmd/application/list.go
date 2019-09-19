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
	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"github.com/wso2/choreo/components/cli/internal/pkg/client"
	"github.com/wso2/choreo/components/cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo/components/cli/internal/pkg/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func NewListCommand(cliConfig config.Config) *cobra.Command {

	const cmdLs = "ls"
	cmd := &cobra.Command{
		Use:     cmdLs,
		Short:   "List applications",
		Example: common.GetAbsoluteCommandName(cmdApplication, cmdLs),
		Args:    cobra.NoArgs,
		Run:     runListAppCommand(cliConfig),
	}
	return cmd
}

func runListAppCommand(cliConfig config.Config) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		listApps(cliConfig)
	}
}

func listApps(cliConfig config.Config) {

	req, err := client.NewRequest(cliConfig, "GET", pathApplications, nil)

	if err != nil {
		log.Print("Error creating post request for listing applications: ", err)
		return
	}

	httpClient := client.NewClient(cliConfig)

	resp, err := httpClient.Do(req)
	if err == nil {
		showListAppsResult(resp)
	} else {
		log.Print("Error making post request for listing applications: ", err)
	}
}

func showListAppsResult(resp *http.Response) {

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
		common.PrintInfo("Error listing applications.")
		fmt.Println("Error: ", string(body))
	}
	err = resp.Body.Close()
	if err != nil {
		log.Print(err)
	}
}
