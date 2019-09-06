/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein is strictly forbidden, unless permitted by WSO2 in accordance with
 * the WSO2 Commercial License available at http://wso2.com/licenses. For specific
 * language governing the permissions and limitations under this license,
 * please see the license as well as any agreement you’ve entered into with
 * WSO2 governing the purchase of this software and any associated services.
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
