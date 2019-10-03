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
	"testing"

	"github.com/spf13/viper"
	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestViperConfigHolderReadValue(t *testing.T) {
	v := viper.New()
	v.Set("foo", "fooValue")
	configHolder := createMockConfigHolder(v)

	got := configHolder.GetString("foo")
	test.AssertString(t, "fooValue", got, "ViperConfigHolder did not read the value correctly")
}

func TestViperConfigHolderReadValueNotDefault(t *testing.T) {
	v := viper.New()
	v.Set("foo", "fooValue")
	configHolder := createMockConfigHolder(v)

	got := configHolder.GetStringOrDefault("foo", "fooDefault")
	test.AssertString(t, "fooValue", got, "ViperConfigHolder did not return the value correctly")
}

func TestViperConfigHolderReadDefault(t *testing.T) {
	v := viper.New()
	configHolder := createMockConfigHolder(v)

	got := configHolder.GetStringOrDefault("foo", "fooDefault")
	test.AssertString(t, "fooDefault", got, "ViperConfigHolder did not read the default")
}

func TestViperConfigHolderWrite(t *testing.T) {
	v := viper.New()
	configHolder := createMockConfigHolder(v)

	configHolder.SetString("foo", "fooValue")
	got := configHolder.GetString("foo")
	test.AssertString(t, "fooValue", got, "ViperConfigHolder did not write the value correctly")
}

func createMockConfigHolder(v *viper.Viper) *ViperConfigHolder {
	return &ViperConfigHolder{viperInstance: v, writer: os.Stdout}
}
