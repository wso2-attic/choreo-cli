/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

import (
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/test"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/config"
)

func TestCreateUserConfigReaderValue(t *testing.T) {
	mockConfig := config.NewMockConfigHolder(map[string]string{"foo": "fooValue"})

	reader := CreateConfigReader(mockConfig, map[string]string{"foo": "fooDefault"})

	got := reader("foo")
	test.AssertString(t, "fooValue", got, "CreateConfigReader did not create the reader correctly")
}

func TestCreateUserConfigReaderDefault(t *testing.T) {
	mockConfig := config.NewMockConfigHolder(map[string]string{})

	reader := CreateConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	test.AssertString(t, "fooDefault", got, "CreateConfigReader did not create the reader correctly")
}

func TestCreateUserConfigWrite(t *testing.T) {
	mockConfig := config.NewMockConfigHolder(map[string]string{})

	writer := CreateConfigWriter(mockConfig)
	writer("foo", "fooValue")

	got := mockConfig.GetString("foo")
	test.AssertString(t, "fooValue", got, "CreateConfigWriter did not create the writer correctly")
}
