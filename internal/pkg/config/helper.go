/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

type GetConfig func(key string) string
type SetConfig func(key string, value string)

func CreateUserConfigReader(cliConfig Config, configDefaults map[string]string) GetConfig {
	return createConfigReader(cliConfig, configDefaults)
}

func CreateEnvironmentConfigReader(cliConfig Config, configDefaults map[string]string) GetConfig {
	return createConfigReader(cliConfig.GetEnvironmentConfig(), configDefaults)
}

func CreateUserConfigWriter(cliConfig Config) SetConfig {
	return func(key string, value string) {
		cliConfig.SetString(key, value)
	}
}

func createConfigReader(reader Reader, configDefaults map[string]string) GetConfig {
	return func(key string) string {
		return reader.GetStringOrDefault(key, configDefaults[key])
	}
}
