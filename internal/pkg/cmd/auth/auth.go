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
	"github.com/spf13/cobra"
	"github.com/wso2/choreo-cli/internal/pkg/config"
)

func NewAuthCommand(cliConfig config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     cmdAuth,
		Short:   "Manage authentication and authorization",
		Args:    cobra.ExactArgs(1),
	}
	cmd.AddCommand(NewConnectCommand(cliConfig))
	return cmd
}
