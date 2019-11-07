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
)

func TestCreateOauth2Conf(t *testing.T) {
	getEnvConfig := func(key string) string {
		switch key {
		case clientId:
			return "testClientId"
		case authUrl:
			return "http://localhost/auth"
		case tokenUrl:
			return "http://localhost/token"
		default:
			return "ERROR_THIS_SHOULD_NOT_BE_RETURNED_AT_ALL"
		}
	}
	oauth2Conf := initOauthClient("/oauth-context", 8765, getEnvConfig)

	test.AssertString(t,"testClientId", oauth2Conf.oauthConf.ClientID, "Client ID is not correct in oauth conf")
	test.AssertString(t,"http://localhost:8765/oauth-context", oauth2Conf.oauthConf.RedirectURL, "RedirectURL is not correct in oauth conf")
	test.AssertString(t,"http://localhost/auth", oauth2Conf.oauthConf.Endpoint.AuthURL,
		"AuthURL is not correct in oauth conf")
	test.AssertString(t,"http://localhost/token", oauth2Conf.oauthConf.Endpoint.TokenURL,
		"TokenURL is not correct in oauth conf")
}
