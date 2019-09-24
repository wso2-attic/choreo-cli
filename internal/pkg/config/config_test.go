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
)

func TestCreateUserConfigReaderValue(t *testing.T) {
	mockConfig := &CliConfig{
		userConfigHolder: &MockConfigHolder{entries: map[string]string{"foo": "fooValue"}},
	}

	reader := CreateUserConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	assertString(t, "fooValue", got, "CreateUserConfigReader did not create the reader correctly")
}

func TestCreateUserConfigReaderDefault(t *testing.T) {
	mockConfig := &CliConfig{
		userConfigHolder: &MockConfigHolder{entries: map[string]string{}},
	}

	reader := CreateUserConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	assertString(t, "fooDefault", got, "CreateUserConfigReader did not create the reader correctly")
}

func TestCreateEnvConfigReaderValue(t *testing.T) {
	mockConfig := &CliConfig{
		envConfigHolder: &MockConfigHolder{entries: map[string]string{"foo": "fooValue"}},
	}

	reader := CreateEnvironmentConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	assertString(t, "fooValue", got, "CreateEnvConfigReader did not create the reader correctly")
}

func TestCreateEnvConfigReaderDefault(t *testing.T) {
	mockConfig := &CliConfig{
		envConfigHolder: &MockConfigHolder{entries: map[string]string{}},
	}

	reader := CreateEnvironmentConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	assertString(t, "fooDefault", got, "CreateEnvConfigReader did not create the reader correctly")
}

func TestCreateUserConfigWrite(t *testing.T) {
	mockConfig := &CliConfig{
		userConfigHolder: &MockConfigHolder{entries: map[string]string{}},
	}

	writer := CreateUserConfigWriter(mockConfig)
	writer("foo", "fooValue")

	got := mockConfig.GetString("foo")
	assertString(t, "fooValue", got, "CreateUserConfigWriter did not create the reader correctly")
}
