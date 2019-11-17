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
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
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
		out:         ctx.Out(),
		debug:       ctx.DebugOut(),
		skipVerify:  skipVerify,
		backendUrl:  ctx.EnvConfig().GetStringOrDefault(BackendUrl, EnvConfigs[BackendUrl]),
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

func (c *cliClient) getRestResource(resourcePath string, v interface{}) error {
	resp, err := c.makeHttpCall(resourcePath, "GET", nil)
	if err != nil {
		return err
	}

	defer closeResource(c.out, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Error response received. Server error: " + string(body))
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return errors.New("Error decoding the response. Reason: " + err.Error())
	}

	return nil
}

func (c *cliClient) createRestResource(resourcePath string, data interface{}) error {
	jsonStr, err := json.Marshal(data)

	if err != nil {
		common.PrintError(c.debug, "Error converting data into JSON format. Reason: ", err)
		return newInternalError()
	}

	resp, err := c.makeHttpCall(resourcePath, "POST", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	defer closeResource(c.out, resp.Body)()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New(string(body))
	}

	return nil
}

func (c *cliClient) createRestResourceWithResponse(resourcePath string, requestData interface{}, responseData interface{}) error {
	jsonStr, err := json.Marshal(requestData)

	if err != nil {
		common.PrintError(c.debug, "Error converting data into JSON format. Reason: ", err)
		return newInternalError()
	}

	resp, err := c.makeHttpCall(resourcePath, "POST", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	defer closeResource(c.out, resp.Body)()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	err = json.Unmarshal(body, responseData)
	if err != nil {
		return errors.New("Error decoding the response. Reason: " + err.Error())
	}
	return nil
}

func (c *cliClient) makeHttpCall(resourcePath string, method string, dataReader io.Reader) (*http.Response, error) {
	completeUrl := c.backendUrl + resourcePath
	req, err := http.NewRequest(method, completeUrl, dataReader)
	if err != nil {
		common.PrintErrorMessage(c.debug, err.Error())
		return nil, fmt.Errorf("error creating server request")
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.skipVerify},
	}
	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Do(req)
	if err != nil {
		common.PrintErrorMessage(c.debug, err.Error())
		return nil, fmt.Errorf("error communicating with the server")
	}

	return resp, nil
}
