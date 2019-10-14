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
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

type NoWriter struct {
}

func (c *NoWriter) Write(p []byte) (int, error) {
	//do nothing
	return 0, nil
}

type CliContextData struct {
	userConfig    runtime.UserConfig
	envConfig     runtime.EnvConfig
	verboseWriter io.Writer
}

func (c *CliContextData) Out() io.Writer {
	return os.Stdout
}

func (c *CliContextData) DebugOut() io.Writer {
	return c.verboseWriter
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
	cobra.OnInitialize(cobraOnInit(cliContext, &command))

	if err := command.Execute(); err != nil {
		common.ExitWithError(cliContext.Out(), "Error executing "+common.GetAbsoluteCommandName()+" command", err)
	}
}

func initConfig(cliContext *CliContextData) {
	userConfig, err := config.InitUserConfig()
	if err != nil {
		common.ExitWithError(cliContext.Out(), "Error loading user configs", err)
	}
	cliContext.userConfig = userConfig

	envConfig, err := config.InitEnvConfig()
	if err != nil {
		common.ExitWithError(cliContext.Out(), "Error loading env configs", err)
	}
	cliContext.envConfig = envConfig
}

func initCommands(cliContext *CliContextData) cobra.Command {
	command := cobra.Command{
		Use:   common.GetAbsoluteCommandName() + " COMMAND",
		Short: "Manage integration applications with " + common.ProductName + " platform",
	}
	command.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	command.AddCommand(cmd.NewVersionCommand(cliContext))
	command.AddCommand(auth.NewAuthCommand(cliContext))
	command.AddCommand(application.NewApplicationCommand(cliContext))

	return command
}

func cobraOnInit(cliContext *CliContextData, command *cobra.Command) func() {
	return func() {
		verbose, err := command.PersistentFlags().GetBool("verbose")
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error retrieving verbose flag value", err)
		}
		if verbose {
			cliContext.verboseWriter = os.Stdout
		} else {
			cliContext.verboseWriter = &NoWriter{}
		}
	}
}
