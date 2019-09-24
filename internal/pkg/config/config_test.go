/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

import "testing"


func TestCreateUserConfigReaderValue(t *testing.T) {
	mockConfig := &MockConfig{
		userConfigHolder: &MockConfigHolder{entries: map[string]string{"foo": "fooValue"}},
	}

	reader := CreateUserConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	if got != "fooValue" {
		t.Errorf("CreateUserConfigReader did not create the reader correctly; %s; want %s", "fooValue", got)
	}
}

func TestCreateUserConfigReaderDefault(t *testing.T) {
	mockConfig := &MockConfig{
		userConfigHolder: &MockConfigHolder{entries: map[string]string{}},
	}

	reader := CreateUserConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	if got != "fooDefault" {
		t.Errorf("CreateUserConfigReader did not create the reader correctly; %s; want %s", "fooDefault", got)
	}
}

func TestCreateEnvConfigReaderValue(t *testing.T) {
	mockConfig := &MockConfig{
		envConfigHolder: &MockConfigHolder{entries: map[string]string{"foo": "fooValue"}},
	}

	reader := CreateEnvironmentConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	if got != "fooValue" {
		t.Errorf("CreateEnvConfigReader did not create the reader correctly; %s; want %s", "fooValue", got)
	}
}

func TestCreateEnvConfigReaderDefault(t *testing.T) {
	mockConfig := &MockConfig{
		envConfigHolder: &MockConfigHolder{entries: map[string]string{}},
	}

	reader := CreateEnvironmentConfigReader(mockConfig, map[string]string{"foo":"fooDefault"})

	got := reader("foo")
	if got != "fooDefault" {
		t.Errorf("CreateEnvConfigReader did not create the reader correctly; %s; want %s", "fooDefault", got)
	}
}

func TestCreateUserConfigWrite(t *testing.T) {
	mockConfig := &MockConfig{
		userConfigHolder: &MockConfigHolder{entries: map[string]string{}},
	}

	writer := CreateUserConfigWriter(mockConfig)
	writer("foo", "fooValue")

	got := mockConfig.GetString("foo")
	if got != "fooValue" {
		t.Errorf("CreateUserConfigWriter did not create the reader correctly; %s; want %s", "fooValue", got)
	}
}