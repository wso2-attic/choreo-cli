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
	"fmt"

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
		Run:  createAppCreateCommand(cliContext),
	}
	cmd.Flags().StringP(descriptionFlagName, "d", "", "Specify description for the application")
	return cmd
}

func createAppCreateCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if !client.IsUserLoggedIn(cliContext) {
			common.ExitWithErrorMessage(cliContext.Out(), "Please login first")
		}

		description, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error reading description flag", err)
		}

		err = cliContext.Client().CreateNewApp(args[0], description)
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error occurred while creating the application", err)
		}
	}
}
