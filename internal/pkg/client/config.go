/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package client

const (
	BackendUrl  = "choreo.backend.url"
	SkipVerify  = "security.certificate.skipVerify"
	AccessToken = "login.oauth2.accessToken"
)

var EnvConfigs = map[string]string{
	BackendUrl: "https://api.development.choreo.dev",
	SkipVerify: "false",
}

var UserConfigs = map[string]string{
	AccessToken: "login.oauth2.accessToken",
}
