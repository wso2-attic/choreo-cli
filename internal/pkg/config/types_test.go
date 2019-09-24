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
	"testing"

	"github.com/spf13/viper"
)

func TestViperConfigHolderReadValue(t *testing.T) {
	v := viper.New()
	v.Set("foo", "fooValue")
	configHolder := &ViperConfigHolder{viperInstance: v}

	got := configHolder.GetString("foo")
	if got != "fooValue" {
		t.Errorf("ViperConfigHolder did not read the value correctly; %s; want %s", "fooValue", got)
	}
}

func TestViperConfigHolderReadValueNotDefault(t *testing.T) {
	v := viper.New()
	v.Set("foo", "fooValue")
	configHolder := &ViperConfigHolder{viperInstance: v}

	got := configHolder.GetStringOrDefault("foo", "fooDefault")
	if got != "fooValue" {
		t.Errorf("ViperConfigHolder did not read the value correctly; %s; want %s", "fooValue", got)
	}
}

func TestViperConfigHolderReadDefault(t *testing.T) {
	v := viper.New()
	configHolder := &ViperConfigHolder{viperInstance: v}

	got := configHolder.GetStringOrDefault("foo", "fooDefault")
	if got != "fooDefault" {
		t.Errorf("ViperConfigHolder did not read the default correctly; %s; want %s", "fooDefault", got)
	}
}

func TestViperConfigHolderWrite(t *testing.T) {
	v := viper.New()
	configHolder := &ViperConfigHolder{viperInstance: v}

	configHolder.SetString("foo", "fooValue")
	got := configHolder.GetString("foo")
	if got != "fooValue" {
		t.Errorf("ViperConfigHolder did not write the value correctly; %s; want %s", "fooValue", got)
	}
}