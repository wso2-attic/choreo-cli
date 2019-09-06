/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package config

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func InitConfig() (*ViperConfig, error) {
	config := new(ViperConfig)
	
	if err := initializeViper(config); err != nil {
		return nil, err
	}

	return config, nil
}

func initializeViper(config *ViperConfig) error {
	v := viper.New()
	if err := loadConfigFile(v); err != nil {
		return err
	}
	config.viperInstance = v
	return nil
}

func loadConfigFile(v *viper.Viper) error {
	homeDirectoryLocation, err := homedir.Dir()
	absoluteConfigFileDirectory := filepath.Join(homeDirectoryLocation, configFileDir, userConfigFileName)
	if err != nil {
		return err
	}
	v.SetConfigFile(absoluteConfigFileDirectory)
	err = v.ReadInConfig()
	if err != nil {
		// Ignore error if the file is not found
		if _, ok := err.(*os.PathError); !ok {
			return err
		}
	}
	return nil
}

func GetConfigReader(cliConfig Config, configDefinition map[int]KeyEntry) func(entry int) string {
	return func(entry int) string {
		return cliConfig.GetStringForKeyEntry(configDefinition[entry])
	}
}
