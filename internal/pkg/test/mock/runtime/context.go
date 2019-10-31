/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package runtime

import (
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"io"
)

type MockCliContext struct {
	MockOut        io.Writer
	MockDebugOut   io.Writer
	MockUserConfig runtime.UserConfig
	MockEnvConfig  runtime.EnvConfig
	MockClient     runtime.Client
}

func (c *MockCliContext) Out() io.Writer {
	return c.MockOut
}

func (c *MockCliContext) DebugOut() io.Writer {
	return c.MockDebugOut
}

func (c *MockCliContext) UserConfig() runtime.UserConfig {
	return c.MockUserConfig
}

func (c *MockCliContext) EnvConfig() runtime.EnvConfig {
	return c.MockEnvConfig
}

func (c *MockCliContext) Client() runtime.Client {
	return c.MockClient
}
