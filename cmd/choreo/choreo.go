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
	"github.com/wso2/choreo-cli/internal/pkg/cmd/auth"
	cmdCommon "github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

type CliContextData struct {
	userConfig runtime.UserConfig
	envConfig  runtime.EnvConfig
}

func (c *CliContextData) Out() io.Writer {
	return os.Stdout
}

func (c *CliContextData) UserConfig() runtime.UserConfig {
	return c.userConfig
}

func (c *CliContextData) EnvConfig() runtime.EnvConfig {
	return c.envConfig
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
	userConfig, err := config.InitUserConfig()
	if err != nil {
		cmdCommon.ExitWithError(cliContext.Out(), "Error loading user configs", err)
	}
	cliContext.userConfig = userConfig

	envConfig, err := config.InitEnvConfig()
	if err != nil {
		cmdCommon.ExitWithError(cliContext.Out(), "Error loading env configs", err)
	}
	cliContext.envConfig = envConfig
}

func initCommands(cliContext runtime.CliContext) cobra.Command {
	command := cobra.Command{
		Use:   cmdCommon.GetAbsoluteCommandName() + " COMMAND",
		Short: "Manage integration applications with " + cmdCommon.ProductName + " platform",
	}

	command.AddCommand(cmd.NewVersionCommand(cliContext))
	command.AddCommand(auth.NewAuthCommand(cliContext))
	command.AddCommand(application.NewApplicationCommand(cliContext))

	return command
}
