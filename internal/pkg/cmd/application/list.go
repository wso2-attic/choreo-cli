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
	"io"

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

		apps, err := cliContext.Client().ListApps()
		if err == nil {
			showApps(cliContext.Out(), apps)
		} else {
			common.ExitWithError(cliContext.Out(), "Error while retrieving application details", err)
		}
	}
}

func showApps(consoleOut io.Writer, apps []runtime.Application) {
	if len(apps) == 0 {
		common.PrintInfo(consoleOut, "No applications to show!")
	}

	printer := tableprinter.New(consoleOut)
	printer.Print(apps)
}
