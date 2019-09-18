/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package client

import (
	"crypto/tls"
	"github.com/wso2/choreo/components/cli/internal/pkg/config"
	"io"
	"net/http"
	"strconv"
)

// set InsecureSkipVerify option if required
func NewClient(cliConfig config.Config) *http.Client {

	getEnvConfig := config.GetEnvironmentConfigReader(cliConfig, EnvConfigs)
	skipVerify, _ := strconv.ParseBool(getEnvConfig(SkipVerify))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}
	return &http.Client{Transport: tr}
}

// creates the http request for the given path with Authorization header set
func NewRequest(cliConfig config.Config, method, path string, body io.Reader) (*http.Request, error) {

	getEnvConfig := config.GetEnvironmentConfigReader(cliConfig, EnvConfigs)
	getUserConfig := config.GetUserConfigReader(cliConfig, UserConfigs)

	completeUrl := getEnvConfig(BackendUrl) + path
	req, err := http.NewRequest(method, completeUrl, body)

	if err == nil {
		req.Header.Set("Authorization", "Bearer "+getUserConfig(AccessToken))
		req.Header.Set("Content-Type", "application/json")
	}

	return req, err
}
