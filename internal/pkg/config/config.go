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

func InitUserConfig() (*ViperConfigHolder, error) {
	userConfig := viper.New()
	if err := loadConfigFile(userConfig, userConfigFileName); err != nil {
		return nil, err
	}
	return &ViperConfigHolder{viperInstance: userConfig}, nil
}

func InitEnvConfig() (*ViperConfigHolder, error) {
	envConfig := viper.New()
	if getEnvAsBool(enableEnvConfigPropertyName, false) {
		if err := loadConfigFile(envConfig, environmentConfigFileName); err != nil {
			return nil, err
		}
	}
	return &ViperConfigHolder{viperInstance: envConfig}, nil
}

func loadConfigFile(v *viper.Viper, configFileName string) error {
	configDirectory, err := getConfigDirectory()
	absoluteConfigFileDirectory := filepath.Join(configDirectory, configFileName)
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

func getConfigDirectory() (string, error) {
	homeDirectoryLocation, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	absoluteConfigDirectory := filepath.Join(homeDirectoryLocation, configFileDir)
	return absoluteConfigDirectory, nil
}
