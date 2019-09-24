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

type MockConfigHolder struct {
	entries map[string]string
}

func (configHolder *MockConfigHolder) GetString(key string) string {
	return configHolder.entries[key]
}

func (configHolder *MockConfigHolder) GetStringOrDefault(key string, defaultValue string) string {
	return getStringOrDefault(configHolder.GetString, key, defaultValue)
}

func (configHolder *MockConfigHolder) SetString(key string, value string) {
	configHolder.entries[key] = value
}

func assertString(t *testing.T, expected string, actual string, message string) {
	if expected != actual {
		t.Errorf("%s; expected: %s; actual %s", message, expected, actual)
	}
}

func assertBool(t *testing.T, expected bool, actual bool, message string) {
	if expected != actual {
		t.Errorf("%s; expected: %t; actual %t", message, expected, actual)
	}
}
