/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package login

import (
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/test"
	"github.com/wso2/choreo-cli/internal/pkg/test/mock/config"
)

func TestCreateEnvConfigReaderDefault(t *testing.T) {
	configHolder := config.NewMockEnvConfigHolder(map[string]string{})
	getEnvConfig := createEnvConfigReader(configHolder)
	clientIdValue := getEnvConfig(clientId)

	test.AssertString(t, "choreocliapplication", clientIdValue, "Default value not returned for env config")
}

func TestCreateEnvConfigReaderValue(t *testing.T) {
	configHolder := config.NewMockEnvConfigHolder(map[string]string{clientId: "testclientid"})
	getEnvConfig := createEnvConfigReader(configHolder)
	clientIdValue := getEnvConfig(clientId)

	test.AssertString(t, "testclientid", clientIdValue, "Correct value not returned for env config")
}
