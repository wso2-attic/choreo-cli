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
	"testing"

	cl "github.com/wso2/choreo-cli/internal/pkg/client"
	dt "github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/client"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/config"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/runtime"

	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestCreateAndDeployApp(t *testing.T) {
	var b bytes.Buffer
	deployCommand := NewDeployCommand(&runtime.MockCliContext{
		MockOut:        &b,
		MockUserConfig: config.NewMockConfigHolder(map[string]string{cl.AccessToken: "some-token"}),
		MockEnvConfig:  config.NewMockConfigHolder(map[string]string{}),
		MockClient: &client.MockClient{CreateAndDeployApp_: func(repoUrl string) (d dt.DeploymentDetails, e error) {
			return dt.DeploymentDetails{DeploymentUrl: "https://development.choreo.dev/choreoapps/appe1231bbf533d3f2e1287f437ff17d7c8", ApplicationId: "appe1231bbf533d3f2e1287f437ff17d7c8"}, nil
		}},
	})
	deployCommand.SetArgs([]string{"https://github.com/someuser/myapp"})
	_ = deployCommand.Execute()

	expect := "A new application is created for deployment with Id: appe1231bbf533d3f2e1287f437ff17d7c8" +
		"\nOnce deployed, the app can be accessed from" +
		" https://development.choreo.dev/choreoapps/appe1231bbf533d3f2e1287f437ff17d7c8" + "\n"
	test.AssertString(t, expect, b.String(), "Deployment command output is not as expected")
}

func TestCreateAndDeployAppWithName(t *testing.T) {
	var b bytes.Buffer
	deployCommand := NewDeployCommand(&runtime.MockCliContext{
		MockOut:        &b,
		MockUserConfig: config.NewMockConfigHolder(map[string]string{cl.AccessToken: "some-token"}),
		MockEnvConfig:  config.NewMockConfigHolder(map[string]string{}),
		MockClient: &client.MockClient{CreateAndDeployAppWithName_:
		func(appName, repoUrl string) (d dt.DeploymentDetails, e error) {
			return dt.DeploymentDetails{DeploymentUrl: "https://development.choreo.dev/choreoapps/appe1231bbf533d3f2e1287f437ff17d7c8", ApplicationId: "appe1231bbf533d3f2e1287f437ff17d7c8"}, nil
		}},
	})
	deployCommand.SetArgs([]string{"-n", "hello-app","https://github.com/someuser/myapp"})
	_ = deployCommand.Execute()

	expect := "A new application is created for deployment with Id: appe1231bbf533d3f2e1287f437ff17d7c8" +
		"\nOnce deployed, the app can be accessed from" +
		" https://development.choreo.dev/choreoapps/appe1231bbf533d3f2e1287f437ff17d7c8" + "\n"
	test.AssertString(t, expect, b.String(), "Deployment command output is not as expected")
}
