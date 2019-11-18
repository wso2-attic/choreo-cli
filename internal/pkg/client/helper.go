/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package client

import (
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

func IsUserLoggedIn(cliContext runtime.UserConfigHolder) bool {

	token := cliContext.UserConfig().GetString(AccessToken)
	return token != ""
}
