/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

import "github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"

type GetConfig func(key string) string
type SetConfig func(key string, value string)

func CreateConfigReader(reader runtime.Reader, configDefaults map[string]string) GetConfig {
	return func(key string) string {
		return reader.GetStringOrDefault(key, configDefaults[key])
	}
}

func CreateConfigWriter(cliConfig runtime.Writer) SetConfig {
	return func(key string, value string) {
		cliConfig.SetString(key, value)
	}
}
