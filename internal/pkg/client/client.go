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
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

const pathApplications = "/applications"

func CreateClient(ctx runtime.CliContext) *cliClient {
	skipVerify, _ := strconv.ParseBool(ctx.EnvConfig().GetStringOrDefault(SkipVerify, EnvConfigs[SkipVerify]))
	return &cliClient{
		out: ctx.Out(),
		skipVerify: skipVerify,
		backendUrl: ctx.EnvConfig().GetStringOrDefault(BackendUrl, EnvConfigs[BackendUrl]),
		accessToken: ctx.UserConfig().GetStringOrDefault(AccessToken, UserConfigs[AccessToken]),
	}
}

type cliClient struct {
	out         io.Writer
	skipVerify  bool
	accessToken string
	backendUrl  string
}

type Application struct {
	Name        string `json:"name" header:"Application Name"`
	Description string `json:"description" header:"Description"`
}

func (c *cliClient) CreateNewApp(name string, desc string) error {
	application := &Application{
		Name:        name,
		Description: desc,
	}
	jsonStr, err := json.Marshal(application)
	if err != nil {
		common.ExitWithError(c.out, "Error converting application data into JSON format. Reason: ", err)
	}

	req, err := NewRequest(c.backendUrl, c.accessToken, "POST", pathApplications, bytes.NewBuffer(jsonStr))
	if err != nil {
		common.ExitWithError(c.out, "Error creating post request for application creation. Reason: ", err)
	}

	httpClient := NewClient(c.skipVerify)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		defer closeResource(c.out, resp.Body)()
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}

func closeResource(consoleWriter io.Writer, res io.Closer) func() {
	return func() {
		if err := res.Close(); err != nil {
			common.PrintError(consoleWriter, "Error closing resource. Reason: ", err)
		}
	}
}
