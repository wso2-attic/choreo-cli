/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package common

import (
	"bytes"
	"errors"
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestGetAbsoluteCommandName(t *testing.T) {
	generatedCommand := GetAbsoluteCommandName("foo")
	expectedCommand := commandRoot + " foo"
	test.AssertString(t, expectedCommand, generatedCommand, "Generated command is incorrect")
}

func TestGetAbsoluteCommandNameLong(t *testing.T) {
	generatedCommand := GetAbsoluteCommandName("foo", "bar", "abc")
	expectedCommand := commandRoot + " foo bar abc"
	test.AssertString(t, expectedCommand, generatedCommand, "Generated command is incorrect")
}

func TestGetLocalBindAddress(t *testing.T) {
	address := GetLocalBindAddress(9999)
	test.AssertString(t, ":9999", address, "Generated address is not correct")
}

func TestPrintErrorMessage(t *testing.T) {
	var b bytes.Buffer
	PrintErrorMessage(&b, "test message")

	expect := "test message\n"
	test.AssertString(t, expect, b.String(), "Incorrect error message format")
}

func TestPrintError(t *testing.T) {
	var b bytes.Buffer
	PrintError(&b, "test message", errors.New("test error"))

	expect := "\ntest message: test error\n"
	test.AssertString(t, expect, b.String(), "Incorrect error message format")
}

func TestPrintInfo(t *testing.T) {
	var b bytes.Buffer
	PrintInfo(&b, "test message")

	expect := "test message\n"
	test.AssertString(t, expect, b.String(), "Incorrect info message format")
}

func TestGetStringOrDefaultReturnValue(t *testing.T)  {
	getValue := func(key string) string { return "test-value"}
	output := GetStringOrDefault(getValue, "key", "default")

	test.AssertString(t, "test-value", output, "GetStringOrDefault should return value when available")
}

func TestGetStringOrDefaultReturnDefault(t *testing.T)  {
	getValue := func(key string) string { return ""}
	output := GetStringOrDefault(getValue, "key", "default")

	test.AssertString(t, "default", output, "GetStringOrDefault should return value when available")
}
