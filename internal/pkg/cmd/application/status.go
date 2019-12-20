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
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

func NewStatusCommand(cliContext runtime.CliContext) *cobra.Command {
	const cmdStatus = "status"
	cmd := &cobra.Command{
		Use:   cmdStatus + " APP_ID",
		Short: "Get status of an application",
		Example: fmt.Sprint(common.GetAbsoluteCommandName(cmdApplication, cmdStatus),
			" app2bb24a9488fc46b05212eac1db9dcc81"),
		Args: cobra.ExactArgs(1),
		Run:  runAppStatusCommand(cliContext),
	}
	return cmd
}

func runAppStatusCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		failIfUserNotLoggedIn(cliContext)

		status, err := cliContext.Client().GetApplicationStatus(args[0])
		if err != nil {
			common.ExitWithErrorMessage(cliContext.Out(), "Unable to retrieve state of the application")
		} else {
			common.PrintInfo(cliContext.Out(), status)
		}
	}
}
