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
	"io/ioutil"
	"net/http"
)

const oauthStateResourcePath = "/internal/source_code/github/oauth/state"

func (c *cliClient) CreateOauthStateString() (string, error) {
	req, err := NewRequest(c.backendUrl, c.accessToken, "GET", oauthStateResourcePath, nil)
	if err != nil {
		return "", err
	}

	httpClient := NewClient(c.skipVerify)
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer closeResource(c.out, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Error response received for state request. Server error: " + string(body))
	}

	var stateObj struct {
		State string `json:"state"`
	}
	err = json.Unmarshal(body, &stateObj)
	if err != nil {
		return "", errors.New("Error decoding the response. Reason: "+ err.Error())
	}

	return stateObj.State, nil
}
