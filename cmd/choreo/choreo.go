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
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/cmd"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/application"
	cmdCommon "github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/login"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

type CliContextData struct {
	config config.Config
}

func (c *CliContextData) Out() io.Writer {
	return os.Stdout
}

func (c *CliContextData) Config() config.Config {
	return c.config
}

func main() {
	cliContext := &CliContextData{}

	initConfig(cliContext)
	command := initCommands(cliContext)

	if err := command.Execute(); err != nil {
		cmdCommon.ExitWithError(cliContext.Out(), "Error executing "+cmdCommon.GetAbsoluteCommandName()+" command", err)
	}
}

func initConfig(cliContext *CliContextData) {
	cliConfig, err := config.InitConfig()
	if err != nil {
		cmdCommon.ExitWithError(cliContext.Out(), "Error loading configs", err)
	}
	cliContext.config = cliConfig
}

func initCommands(cliContext runtime.CliContext) cobra.Command {
	command := cobra.Command{
		Use:   cmdCommon.GetAbsoluteCommandName() + " COMMAND",
		Short: "Manage integration applications with " + cmdCommon.ProductName + " platform",
	}

	command.AddCommand(cmd.NewVersionCommand(cliContext))
	command.AddCommand(login.NewLoginCommand(cliContext))
	command.AddCommand(application.NewApplicationCommand(cliContext))

	return command
}
