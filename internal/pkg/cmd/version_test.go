/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestPrintVersionInfo(t *testing.T) {
	var b bytes.Buffer
	printVersionInfo(&b)

	expect := ` Version:		unknown
 Git commit:		unknown
 Built:			unknown
 OS/Arch:		unknown
`
	test.AssertString(t, expect, b.String(), "Version command output is not as expected")
}

func TestNewVersionCommand(t *testing.T) {
	var b bytes.Buffer
	versionCommand := NewVersionCommand(&mockContext{out: &b})
	versionCommand.Run(nil,nil)

	expect := ` Version:		unknown
 Git commit:		unknown
 Built:			unknown
 OS/Arch:		unknown
`
	test.AssertString(t, expect, b.String(), "Version command output is not as expected")
}

type mockContext struct {
	out io.Writer
}

func (c *mockContext) Out() io.Writer {
	return c.out
}
