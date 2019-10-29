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
	"bytes"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/config"
	"io"
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestNewDeployCommandWithoutLogin(t *testing.T) {
	var b bytes.Buffer
	deployCommand := NewDeployCommand(&mockContext{
		out:        &b,
		userConfig: config.NewMockConfigHolder(map[string]string{}),
		envConfig:  config.NewMockConfigHolder(map[string]string{}),
	})
	deployCommand.Run(nil, nil)

	expect := `Please login first
`
	test.AssertString(t, expect, b.String(), "Deployment command output is not as expected")
}

type mockContext struct {
	out        io.Writer
	debugOut   io.Writer
	userConfig runtime.UserConfig
	envConfig  runtime.EnvConfig
	apiClient  runtime.Client
}

func (c *mockContext) Out() io.Writer {
	return c.out
}

func (c *mockContext) DebugOut() io.Writer {
	return c.debugOut
}

func (c *mockContext) UserConfig() runtime.UserConfig {
	return c.userConfig
}

func (c *mockContext) EnvConfig() runtime.EnvConfig {
	return c.envConfig
}

func (c *mockContext) Client() runtime.Client {
	return c.apiClient
}
