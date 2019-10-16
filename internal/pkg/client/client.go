/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package client

import (
	"io"
	"strconv"

	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

func CreateClient(ctx runtime.CliContext) *cliClient {
	skipVerify, _ := strconv.ParseBool(ctx.EnvConfig().GetStringOrDefault(SkipVerify, EnvConfigs[SkipVerify]))
	return &cliClient{
		out: ctx.Out(),
		skipVerify: skipVerify,
		backendUrl: ctx.EnvConfig().GetStringOrDefault(BackendUrl, EnvConfigs[BackendUrl]),
		accessToken: ctx.UserConfig().GetStringOrDefault(AccessToken, UserConfigs[AccessToken]),
	}
}

type cliClient struct {
	out         io.Writer
	skipVerify  bool
	accessToken string
	backendUrl  string
}

func closeResource(consoleWriter io.Writer, res io.Closer) func() {
	return func() {
		if err := res.Close(); err != nil {
			common.PrintError(consoleWriter, "Error closing resource. Reason: ", err)
		}
	}
}
