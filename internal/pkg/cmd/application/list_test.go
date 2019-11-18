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
	"github.com/wso2/choreo-cli/internal/pkg/test"
	"testing"
)

func TestShowAppsCommandWithApps(t *testing.T) {
	var b bytes.Buffer
	var apps = []runtime.Application{{Id: "appee067864d5714595b9938812a8342b59", Name: "hello World",
		Description: "My first hello world app"}, {Id: "app3a1f1edb0dce4db080f9135250600d7c", Name: "app1"}}
	expect := `  ID                                    APPLICATION NAME   DESCRIPTION               
 ------------------------------------- ------------------ -------------------------- 
  appee067864d5714595b9938812a8342b59   hello World        My first hello world app  
  app3a1f1edb0dce4db080f9135250600d7c   app1                                         
`
	showApps(&b, apps)
	test.AssertString(t, expect, b.String(), "showApps command output is not as expected")
}

func TestShowAppsCommandWithoutApps(t *testing.T) {
	var b bytes.Buffer
	var apps []runtime.Application
	expect := `No applications to show!
`
	showApps(&b, apps)
	test.AssertString(t, expect, b.String(), "showApps command output is not as expected")
}
