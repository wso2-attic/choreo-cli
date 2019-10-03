/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/cmd"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/application"
	cmdCommon "github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/login"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

func main() {
	consoleWriter := os.Stdout

	cliConfig, err := config.InitConfig()
	if err != nil {
		cmdCommon.ExitWithError(consoleWriter, "Error loading configs", err)
	}

	command := initCommands(cliConfig)

	if err := command.Execute(); err != nil {
		cmdCommon.ExitWithError(consoleWriter, "Error executing "+cmdCommon.GetAbsoluteCommandName()+" command", err)
	}
}

func initCommands(cliConfig *config.CliConfig) cobra.Command {
	command := cobra.Command{
		Use:   cmdCommon.GetAbsoluteCommandName() + " COMMAND",
		Short: "Manage integration applications with " + cmdCommon.ProductName + " platform",
	}

	command.AddCommand(cmd.NewVersionCommand(cliConfig))
	command.AddCommand(login.NewLoginCommand(cliConfig))
	command.AddCommand(application.NewApplicationCommand(cliConfig))

	return command
}
