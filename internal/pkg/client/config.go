/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package client

import "github.com/wso2/choreo/components/cli/internal/pkg/config"

const (
	BackendUrl = iota
	SkipVerify
	AccessToken
)

var EnvConfigs = map[int]config.KeyEntry{
	BackendUrl: {
		Key:          "choreo.backend.url",
		DefaultValue: "https://api.choreo.dev:8081",
	},
	SkipVerify: {
		Key:          "security.certificate.skipVerify",
		DefaultValue: "false",
	},
}

var UserConfigs = map[int]config.KeyEntry{
	AccessToken: {
		Key:          "login.oauth2.accessToken",
		DefaultValue: "",
	},
}
