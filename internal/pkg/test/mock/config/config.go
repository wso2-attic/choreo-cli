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
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

type MockConfigHolder struct {
	entries map[string]string
}

func (configHolder *MockConfigHolder) GetString(key string) string {
	return configHolder.entries[key]
}

func (configHolder *MockConfigHolder) GetStringOrDefault(key string, defaultValue string) string {
	return common.GetStringOrDefault(configHolder.GetString, key, defaultValue)
}

func (configHolder *MockConfigHolder) SetString(key string, value string) {
	configHolder.entries[key] = value
}

func NewMockConfigHolder(entries map[string]string) *MockConfigHolder {
	return &MockConfigHolder{entries: entries}
}

type MockEnvConfigHolder struct {
	configHolder *MockConfigHolder
}

func (c *MockEnvConfigHolder) EnvConfig() runtime.EnvConfig {
	return c.configHolder
}

func NewMockEnvConfigHolder(entries map[string]string) *MockEnvConfigHolder {
	return &MockEnvConfigHolder{configHolder: NewMockConfigHolder(entries)}
}