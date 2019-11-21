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
	"fmt"
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

const pathApplications = "/applications"
const pathApplicationDeployment = pathApplications + "/deployments"
const pathApplicationLogs = pathApplications + "/logs"

func (c *cliClient) ListApps() ([]runtime.Application, error) {
	var apps []runtime.Application

	err := c.httpClient.getRestResource(pathApplications, &apps)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (c *cliClient) CreateNewApp(name string, desc string) error {
	application := &runtime.ApplicationRequest{
		Name:        name,
		Description: desc,
	}

	if err := c.httpClient.createRestResource(pathApplications, application); err != nil {
		return err
	}

	return nil
}

func (c *cliClient) DeployApp(repoUrl string) (runtime.DeploymentDetails, error) {
	var deploymentRequest = struct {
		RepoUrl string `json:"repo_url"`
	}{
		RepoUrl: repoUrl,
	}

	var deploymentDetails runtime.DeploymentDetails

	err := c.httpClient.createRestResourceWithResponse(pathApplicationDeployment, &deploymentRequest, &deploymentDetails)

	return deploymentDetails, err
}

func (c *cliClient) FetchLogs(appId string, linesCount uint) (string, error) {

	pathWithQueryParam := pathApplicationLogs + "/" + appId
	if linesCount > 0 {
		pathWithQueryParam += "?lines_count=" + fmt.Sprint(linesCount)
	}

	var logsDetails struct {
		Logs string `json:"logs"`
	}
	err := c.httpClient.getRestResource(pathWithQueryParam, &logsDetails)
	if err != nil {
		return "", err
	}

	return logsDetails.Logs, nil
}
