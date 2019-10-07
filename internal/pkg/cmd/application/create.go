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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

const descriptionFlagName = "description"

func NewCreateCommand(cliContext runtime.CliContext) *cobra.Command {
	const cmdCreate = "create"
	cmd := &cobra.Command{
		Use:   cmdCreate + " APP_NAME",
		Short: "Create an application",
		Example: fmt.Sprint(common.GetAbsoluteCommandName(cmdApplication, cmdCreate),
			" app1 -d \"My first app\""),
		Args: cobra.ExactArgs(1),
		Run:  runCreateAppCommand(cliContext),
	}
	cmd.Flags().StringP(descriptionFlagName, "d", "", "Specify description for the application")
	return cmd
}

func runCreateAppCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		description, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error reading description flag", err)
		}

		createApp(cliContext, &Application{Name: args[0], Description: description})
	}
}

func createApp(cliContext runtime.CliContext, application *Application) {
	err := createNewApp(cliContext, application)
	if err != nil {
		common.ExitWithError(cliContext.Out(), "Error occurred while creating the application. Reason: ", err)
	}
}

func createNewApp(cliContext runtime.CliContext, application *Application) error {
	jsonStr, err := json.Marshal(application)
	if err != nil {
		common.ExitWithError(cliContext.Out(), "Error converting application data into JSON format. Reason: ", err)
	}

	req, err := client.NewRequest(cliContext, "POST", pathApplications, bytes.NewBuffer(jsonStr))
	if err != nil {
		common.ExitWithError(cliContext.Out(), "Error creating post request for application creation. Reason: ", err)
	}

	httpClient := client.NewClient(cliContext)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		defer closeResource(cliContext.Out(), resp.Body)()
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}

func closeResource(consoleWriter io.Writer, res io.Closer) func() {
	return func() {
		if err := res.Close(); err != nil {
			common.PrintError(consoleWriter, "Error closing resource. Reason: ", err)
		}
	}
}
