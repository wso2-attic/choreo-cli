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
	CreateAndDeployApp_     func(deploymentRequest runtime.DeploymentInput) (runtime.DeploymentOut, error)
	FetchLogs_              func(appId string, linesCount uint) (string, error)
	CreateOauthStateString_ func() (string, error)
	DeleteApp_              func(appId string) error
	GetStatus_              func(appId string) (string, error)
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

func (c *MockClient) CreateAndDeployApp(deploymentRequest runtime.DeploymentInput) (runtime.DeploymentOut, error) {
	if c.CreateAndDeployApp_ != nil {
		return c.CreateAndDeployApp_(deploymentRequest)
	}
	return runtime.DeploymentOut{
		DeploymentUrl: "",
		ApplicationId: "",
	}, nil
}

func (c *MockClient) FetchLogs(appId string, linesCount uint) (string, error) {
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

func (c *MockClient) DeleteApp(appId string) error {
	if c.DeleteApp_ != nil {
		return c.DeleteApp_(appId)
	}
	return nil
}

func (c *MockClient) GetApplicationStatus(appId string) (string, error) {
	if c.GetStatus_ != nil {
		return c.GetStatus_(appId)
	}
	return "", nil
}
