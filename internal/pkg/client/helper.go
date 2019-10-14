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
	"io"
	"net/http"
)

// set InsecureSkipVerify option if required
func NewClient(skipVerify bool) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}
	return &http.Client{Transport: tr}
}

// creates the http request for the given path with Authorization header set
func NewRequest(backendUrl string, accessToken string, method, path string, body io.Reader) (*http.Request, error) {
	completeUrl := backendUrl + path
	req, err := http.NewRequest(method, completeUrl, body)

	if err == nil {
		req.Header.Set("Authorization", "Bearer "+accessToken)
		req.Header.Set("Content-Type", "application/json")
	}

	return req, err
}
