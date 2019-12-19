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

func NewDeployCommand(cliContext runtime.CliContext) *cobra.Command {
	const cmdDeploy = "deploy"
	cmd := &cobra.Command{
		Use:   cmdDeploy + " GITHUB_REPO_URL",
		Short: "Deploy an application",
		Example: fmt.Sprint(common.GetAbsoluteCommandName(cmdApplication, cmdDeploy),
			" https://github.com/wso2/choreo-ballerina-hello"),
		Args: checkArgCount(1, 2),
		Run:  runDeployAppCommand(cliContext),
	}
	cmd.Flags().StringP("name", "n", "", "the name to be used for the created application")
	return cmd
}

func checkArgCount(min int, max int) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) > max {
			return fmt.Errorf("accepts at most %d arg(s), received %d", max, len(args))
		} else if len(args) < min {
			return fmt.Errorf("requires at least %d arg(s), only received %d", min, len(args))
		}
		return nil
	}
}

func runDeployAppCommand(cliContext runtime.CliContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		failIfUserNotLoggedIn(cliContext)

		appName, err := cmd.Flags().GetString("name")
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error while reading the application name flag value", err)
		}
		var deploymentRequest runtime.DeploymentInput
		deploymentRequest.Name = appName
		var msg string
		if len(args) == 2 {
			deploymentRequest.ApplicationId = args[0]
			deploymentRequest.Url = args[1]
			msg = "The application with id %s has been updated\nOnce deployed, the app can be accessed from %s"
		} else {
			deploymentRequest.Url = args[0]
			msg = "A new application is created for deployment with Id: %s" +
				"\nOnce deployed, the app can be accessed from %s"
		}
		deploymentDetails, err := cliContext.Client().CreateAndDeployApp(deploymentRequest)
		if err != nil {
			common.ExitWithError(cliContext.Out(), "Error occurred while deploying the application", err)
		} else {
			common.PrintInfo(cliContext.Out(), fmt.Sprintf(msg, deploymentDetails.ApplicationId,
				deploymentDetails.DeploymentUrl))
		}
	}
}
