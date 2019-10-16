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
	"io/ioutil"
	"net/http"

	"github.com/wso2/choreo-cli/internal/pkg/cmd/common"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

const pathApplications = "/applications"

func (c *cliClient) ListApps() ([]runtime.Application, error) {
	var apps []runtime.Application
	req, err := NewRequest(c.backendUrl, c.accessToken, "GET", pathApplications, nil)
	if err != nil {
		return nil, err
	}

	httpClient := NewClient(c.skipVerify)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer closeResource(c.out, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	err = json.Unmarshal(body, &apps)
	if err != nil {
		return nil, errors.New("Error converting json into applications. Reason: "+ err.Error())
	}

	return apps, nil
}

func (c *cliClient) CreateNewApp(name string, desc string) error {
	application := &runtime.Application{
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


