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
	cl "github.com/wso2/choreo-cli/internal/pkg/client"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/client"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/config"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/runtime"
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestNewDeployCommand(t *testing.T) {
	var b bytes.Buffer
	deployCommand := NewDeployCommand(&runtime.MockCliContext{
		MockOut:        &b,
		MockUserConfig: config.NewMockConfigHolder(map[string]string{cl.AccessToken: "some-token"}),
		MockEnvConfig:  config.NewMockConfigHolder(map[string]string{}),
		MockClient: &client.MockClient{DeployApp_: func(repoUrl string) (s string, e error) {
			return "https://development.choreo.dev/choreapps/123456/myapp", nil
		}},
	})
	deployCommand.Run(nil, []string{"https://github.com/someuser/myapp"})

	expect := "Deployment request submitted! Once deployed, the app can be accessed from: https://development.choreo.dev/choreapps/123456/myapp" + "\n"
	test.AssertString(t, expect, b.String(), "Deployment command output is not as expected")
}
