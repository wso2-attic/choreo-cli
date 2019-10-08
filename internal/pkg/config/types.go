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
	"io"
	"os"

	"github.com/spf13/viper"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
)

type ViperConfigHolder struct {
	viperInstance *viper.Viper
	writer        io.Writer
}

func (v *ViperConfigHolder) GetString(key string) string {
	value := v.viperInstance.GetString(key)
	return value
}

func (v *ViperConfigHolder) GetStringOrDefault(key string, defaultValue string) string {
	return common.GetStringOrDefault(v.GetString, key, defaultValue)
}

func (v *ViperConfigHolder) SetString(key string, value string) {
	v.viperInstance.Set(key, value)

	err := v.viperInstance.WriteConfig()
	if err == nil {
		return
	}

	if _, ok := err.(*os.PathError); ok {
		createDirectoryAndWrite(v.viperInstance, key, v.writer)
	} else {
		common.PrintError(v.writer, "Could not write to config. Key: "+key, err)
	}
}

func createDirectoryAndWrite(viperInstance *viper.Viper, key string, writer io.Writer) {
	if err := makeConfigDirectory(); err != nil {
		common.PrintError(writer, "Could not make config directory to config. Key: "+key, err)
	}

	if err := viperInstance.WriteConfig(); err != nil {
		common.PrintError(writer, "Could not write to config. Key: "+key, err)
	}
}

func makeConfigDirectory() error {
	absoluteConfigDirectory, err := getConfigDirectory()
	if err != nil {
		return err
	}

	if err = os.Mkdir(absoluteConfigDirectory, os.ModePerm); err != nil {
		return err
	}

	return nil
}
