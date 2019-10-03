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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/context"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

func NewCreateCommand(cliContext context.CliContext) *cobra.Command {

	const cmdCreate = "create"
	cmd := &cobra.Command{
		Use:   cmdCreate + " APP_NAME",
		Short: "Create an application",
		Example: fmt.Sprint(common.GetAbsoluteCommandName(cmdApplication, cmdCreate),
			" app1 -d \"My first app\""),
		Args: cobra.ExactArgs(1),
		Run:  runCreateAppCommand(cliContext),
	}
	cmd.Flags().StringP("description", "d", "", "Specify description for the application")
	return cmd
}

func runCreateAppCommand(cliContext context.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		description, _ := cmd.Flags().GetString("description")

		app := Application{args[0], description}
		createApp(cliContext.Config(), app)
	}
}

func createApp(cliConfig config.Config, application Application) {

	jsonStr, err := json.Marshal(application)
	if err != nil {
		log.Print("Error converting application into json: ", err)
		return
	}
	req, err := client.NewRequest(cliConfig, "POST", pathApplications, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Print("Error creating post request for application creation: ", err)
		return
	}

	httpClient := client.NewClient(cliConfig)

	resp, err := httpClient.Do(req)
	if err == nil {
		showCreateAppResult(resp)
	} else {
		log.Print("Error making post request for application creation: ", err)
	}
}

func showCreateAppResult(resp *http.Response) {

	if resp.StatusCode == http.StatusCreated {
		common.PrintInfo("Application created successfully.")
	} else {
		common.PrintInfo("Error creating application.")
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print("Error reading json body: ", err)
		}
		fmt.Println("Error: ", string(body))
	}
	err := resp.Body.Close()
	if err != nil {
		log.Print(err)
	}
}
