/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

type GetConfig func(entry int) string
type SetConfig func(keyEntry int, value string)

func GetUserConfigReader(cliConfig Config, configDefinition map[int]KeyEntry) GetConfig {
	return getConfigReader(cliConfig, configDefinition)
}

func GetEnvironmentConfigReader(cliConfig Config, configDefinition map[int]KeyEntry) GetConfig {
	return getConfigReader(cliConfig.GetEnvironmentConfig(), configDefinition)
}

func getConfigReader(reader Reader, configDefinition map[int]KeyEntry) GetConfig {
	return func(entry int) string {
		return reader.GetStringForKeyEntry(configDefinition[entry])
	}
}

func GetUserConfigWriter(cliConfig Config, configDefinition map[int]KeyEntry) SetConfig {
	return func(entry int, value string) {
		cliConfig.SetStringForKeyEntry(configDefinition[entry], value)
	}
}
