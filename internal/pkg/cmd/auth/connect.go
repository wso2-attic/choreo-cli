/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/source/github"
	"strings"
)

func NewConnectCommand(cliContext runtime.CliContext) *cobra.Command {

	const cmdConnect = "connect"
	cmd := &cobra.Command{
		Use:     cmdConnect,
		Short:   "Connect to a source code provider",
		Example: fmt.Sprint(common.GetAbsoluteCommandName(cmdAuth, cmdConnect), " github"),
		Args:    cobra.ExactArgs(1),
		Run:     runConnectCommand(cliContext),
	}
	return cmd
}

func runConnectCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {

	consoleWriter := cliContext.Out()

	return func(cmd *cobra.Command, args []string) {

		if !client.IsUserLoggedIn(cliContext) {
			common.ExitWithErrorMessage(consoleWriter, "Please login first")
		}

		if strings.ToLower(args[0]) == sourceProviderGithub {
			if github.PerformGithubAuthorization(cliContext) {
				common.PrintInfo(consoleWriter, "GitHub authorization successful")
			} else {
				common.PrintErrorMessage(consoleWriter, "GitHub authorization failed")
			}
		} else {
			common.PrintErrorMessage(consoleWriter, "Unsupported source provider specified: "+args[0]+
				". At the moment we only support GitHub.")
		}
	}
}
