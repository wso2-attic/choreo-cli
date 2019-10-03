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
	"os"
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestGetEnvAsBool(t *testing.T) {
	envConfigName := "choreo.config.test/TestGetEnvAsBool"
	if err := os.Setenv(envConfigName, "true"); err != nil {
		t.Errorf("Could not set environment variables in the OS; Error: %s", err)
		return
	}

	defer func() {
		if err := os.Unsetenv(envConfigName); err != nil {
			t.Logf("Could not unset the environment variable; Error: %s", err)
		}
	}()

	got := getEnvAsBool(envConfigName, false)
	test.AssertBool(t, true, got, "Reading a boolean from an environment variable failed")
}

func TestGetEnvAsBoolDefault(t *testing.T) {
	envConfigName := "choreo.config.test/TestGetEnvAsBoolDefault"

	got := getEnvAsBool(envConfigName, true)
	test.AssertBool(t, true, got, "Returning default when environment variable does not exist failed")
}
