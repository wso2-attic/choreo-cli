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

type Config interface {
	Reader
	Writer
	GetEnvironmentConfig() Reader
}

type Reader interface {
	GetString(key string) string
	GetStringOrDefault(key string, defaultValue string) string
}

type CliConfig struct {
	userConfigHolder configHolder
	envConfigHolder  configHolder
}

func (c *CliConfig) GetEnvironmentConfig() Reader {
	return c.envConfigHolder
}

func (c *CliConfig) GetString(key string) string {
	return c.userConfigHolder.GetString(key)
}

func (c *CliConfig) GetStringOrDefault(key string, defaultValue string) string {
	return c.userConfigHolder.GetStringOrDefault(key, defaultValue)
}

type configHolder interface {
	Reader
	Writer
}

type ViperConfigHolder struct {
	viperInstance *viper.Viper
	writer        io.Writer
}

func (v *ViperConfigHolder) GetString(key string) string {
	value := v.viperInstance.GetString(key)
	return value
}

func (v *ViperConfigHolder) GetStringOrDefault(key string, defaultValue string) string {
	return getStringOrDefault(v.GetString, key, defaultValue)
}

func getStringOrDefault(predicate func(key string) string, key string, defaultValue string) string {
	if value := predicate(key); value != "" {
		return value
	} else {
		return defaultValue
	}
}

type Writer interface {
	SetString(key string, value string)
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

func (c *CliConfig) SetString(key string, value string) {
	c.userConfigHolder.SetString(key, value)
}
