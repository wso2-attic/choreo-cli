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
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

func CreateClient(ctx runtime.CliContext) *cliClient {
	skipVerify, _ := strconv.ParseBool(ctx.EnvConfig().GetStringOrDefault(SkipVerify, EnvConfigs[SkipVerify]))
	return &cliClient{
		out: ctx.Out(),
		debug: ctx.DebugOut(),
		skipVerify: skipVerify,
		backendUrl: ctx.EnvConfig().GetStringOrDefault(BackendUrl, EnvConfigs[BackendUrl]),
		accessToken: ctx.UserConfig().GetStringOrDefault(AccessToken, UserConfigs[AccessToken]),
	}
}

type cliClient struct {
	out         io.Writer
	debug       io.Writer
	skipVerify  bool
	accessToken string
	backendUrl  string
}

func closeResource(consoleWriter io.Writer, res io.Closer) func() {
	return func() {
		if err := res.Close(); err != nil {
			common.PrintError(consoleWriter, "Error closing resource. Reason: ", err)
		}
	}
}

func newInternalError() error {
	return errors.New("internal error occurred")
}

func (c *cliClient) getHttpResource(resourcePath string, v interface{}) error {
	req, err := NewRequest(c.backendUrl, c.accessToken, "GET", resourcePath, nil)

	if err != nil {
		return err
	}

	httpClient := NewClient(c.skipVerify)
	resp, err := httpClient.Do(req)
	if err != nil {
		common.PrintErrorMessage(c.debug, err.Error())
		return errors.New("error communicating with the server")
	}

	defer closeResource(c.out, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New("Error response received. Server error: " + string(body))
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return errors.New("Error decoding the response. Reason: "+ err.Error())
	}

	return nil
}
