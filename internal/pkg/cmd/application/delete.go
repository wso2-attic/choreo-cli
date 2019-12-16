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

func NewDeleteCommand(cliContext runtime.CliContext) *cobra.Command {
	const cmdDelete = "delete"
	cmd := &cobra.Command{
		Use:     cmdDelete + " APP_ID",
		Short:   "Delete an application",
		Example: fmt.Sprint(common.GetAbsoluteCommandName(cmdApplication, cmdDelete), " a12345678901"),
		Args:    cobra.ExactArgs(1),
		Run:     createAppDeleteCommand(cliContext),
	}
	return cmd
}

func createAppDeleteCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if !client.IsUserLoggedIn(cliContext) {
			common.ExitWithErrorMessage(cliContext.Out(), "Please login first")
		}

		err := cliContext.Client().DeleteApp(args[0])
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error occurred while deleting the application", err)
		} else {
			common.PrintInfo(cliContext.Out(), args[0]+" is successfully deleted")
		}
	}
}
