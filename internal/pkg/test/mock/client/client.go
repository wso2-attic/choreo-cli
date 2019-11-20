/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package client

import "github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"

type MockClient struct {
	CreateNewApp_           func(name string, desc string) error
	ListApps_               func() ([]runtime.Application, error)
	DeployApp_              func(repoUrl string) (string, error)
	FetchLogs_              func(appId string, linesCount string) (string, error)
	CreateOauthStateString_ func() (string, error)
}

func (c *MockClient) CreateNewApp(name string, desc string) error {
	if c.CreateNewApp_ != nil {
		return c.CreateNewApp_(name, desc)
	}
	return nil
}

func (c *MockClient) ListApps() ([]runtime.Application, error) {
	if c.ListApps_ != nil {
		return c.ListApps_()
	}
	return nil, nil
}

func (c *MockClient) DeployApp(repoUrl string) (string, error) {
	if c.DeployApp_ != nil {
		return c.DeployApp_(repoUrl)
	}
	return "", nil
}

func (c *MockClient) FetchLogs(appId string, linesCount string) (string, error) {
	if c.FetchLogs_ != nil {
		return c.FetchLogs_(appId, linesCount)
	}
	return "", nil
}

func (c *MockClient) CreateOauthStateString() (string, error) {
	if c.CreateOauthStateString_ != nil {
		return c.CreateOauthStateString_()
	}
	return "", nil
}
