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
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/client"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/config"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/runtime"

	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestDeleteApp(t *testing.T) {
	var b bytes.Buffer
	deleteCommand := NewDeleteCommand(&runtime.MockCliContext{
		MockOut:        &b,
		MockUserConfig: config.NewMockConfigHolder(map[string]string{cl.AccessToken: "some-token"}),
		MockEnvConfig:  config.NewMockConfigHolder(map[string]string{}),
		MockClient: &client.MockClient{DeleteApp_: func(appId string) error {
			return nil
		}},
	})
	deleteCommand.SetArgs([]string{"a12345678901"})
	_ = deleteCommand.Execute()

	expected := "a12345678901 is successfully deleted" + "\n"
	test.AssertString(t, expected, b.String(), "App delete command output is not as expected")
}
