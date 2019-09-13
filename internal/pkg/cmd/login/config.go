/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package login

import "github.com/wso2/choreo/components/cli/internal/pkg/config"

const (
	clientId = iota
	authUrl
	tokenUrl
	accessToken
)

var envConfigs = map[int]config.KeyEntry{
	clientId: {
		Key:          "login.oauth2.clientId",
		DefaultValue: "uEJMEFl4OFHbm54id3xdZiCHPS0a",
	},
	authUrl: {
		Key:          "login.oauth2.authUrl",
		DefaultValue: "https://id.development.choreo.dev/oauth2/authorize",
	},
	tokenUrl: {
		Key:          "login.oauth2.tokenUrl",
		DefaultValue: "https://id.development.choreo.dev/oauth2/token",
	},
}

var userConfigs = map[int]config.KeyEntry{
	accessToken: {
		Key:          "login.oauth2.accessToken",
		DefaultValue: "",
	},
}
