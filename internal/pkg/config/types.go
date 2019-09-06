/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

import "github.com/spf13/viper"

type Config interface {
	GetString(key string) string
}

type ViperConfig struct {
	viperInstance *viper.Viper
}

func (cliConfig *ViperConfig) GetString(key string) string {
	return cliConfig.viperInstance.GetString(key)
}
