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
	"github.com/spf13/cobra"
	"github.com/wso2/choreo/components/cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo/components/cli/internal/pkg/config"
)

type Application struct {
	Name        string `json:"name" header:"Application Name"`
	Description string `json:"description" header:"Description"`
}

type Applications []Application

func NewApplicationCommand(cliConfig config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "application",
		Short:   "Commands related to an environment",
		Example: common.GetAbsoluteCommandName("environment"),
		Args:    cobra.ExactArgs(1),
	}
	cmd.AddCommand(NewCreateCommand(cliConfig))
	cmd.AddCommand(NewListCommand(cliConfig))
	return cmd
}
