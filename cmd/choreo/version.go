/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein is strictly forbidden, unless permitted by WSO2 in accordance with
 * the WSO2 Commercial License available at http://wso2.com/licenses. For specific
 * language governing the permissions and limitations under this license,
 * please see the license as well as any agreement you’ve entered into with
 * WSO2 governing the purchase of this software and any associated services.
 */

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wso2/choreo/components/cli/internal/pkg/build"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Get " + productName + " client version information",
		Example: getAbsoluteCommandName("version"),
		Args:    cobra.NoArgs,
		Run:     runVersion,
	}
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf(" Version:\t\t%s\n", build.GetBuildVersion())
	fmt.Printf(" Git commit:\t\t%s\n", build.GetBuildGitRevision())
	fmt.Printf(" Built:\t\t\t%s\n", build.GetBuildTime())
	fmt.Printf(" OS/Arch:\t\t%s\n", build.GetBuildPlatform())
}
