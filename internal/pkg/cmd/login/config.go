/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package login

import "github.com/wso2/choreo-cli/internal/pkg/config"

const (
	clientId    = "login.oauth2.clientId"
	authUrl     = "login.oauth2.authUrl"
	tokenUrl    = "login.oauth2.tokenUrl"
)

func createEnvConfigReader(cliConfig config.Config) func(string) string {
	return config.CreateEnvironmentConfigReader(cliConfig, map[string]string{
		clientId: "choreocliapplication",
		authUrl:  "https://id.development.choreo.dev/oauth2/authorize",
		tokenUrl: "https://id.development.choreo.dev/oauth2/token",
	})
}