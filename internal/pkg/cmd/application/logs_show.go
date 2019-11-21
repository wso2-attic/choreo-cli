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

const numOfLinesFlagName = "number-of-lines"

func NewShowLogsCommand(cliContext runtime.CliContext) *cobra.Command {
	const cmdShow = "show"
	cmd := &cobra.Command{
		Use:   cmdShow + " APP_ID",
		Short: "Show logs of a deployed application",
		Example: fmt.Sprint(common.GetAbsoluteCommandName(cmdApplication, cmdLogs, cmdShow),
			" app1234567890abcd"),
		Args: cobra.ExactArgs(1),
		Run:  runShowLogsCommand(cliContext),
	}
	cmd.Flags().UintP(numOfLinesFlagName, "n", 0, "Specify number of log lines which should be fetched")
	return cmd
}

func runShowLogsCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if !client.IsUserLoggedIn(cliContext) {
			common.ExitWithErrorMessage(cliContext.Out(), "Please login first")
		}

		linesCount, err := cmd.Flags().GetUint(numOfLinesFlagName)
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error reading description flag", err)
		}

		logs, err := cliContext.Client().FetchLogs(args[0], linesCount)
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error occurred while fetching logs of the application", err)
		} else {
			common.PrintInfo(cliContext.Out(), logs)
		}
	}
}
