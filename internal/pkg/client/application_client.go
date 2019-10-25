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
	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
)

const pathApplications = "/applications"

func (c *cliClient) ListApps() ([]runtime.Application, error) {
	var apps []runtime.Application

	err := c.getRestResource(pathApplications, &apps)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (c *cliClient) CreateNewApp(name string, desc string) error {
	application := &runtime.Application{
		Name:        name,
		Description: desc,
	}

	if err := c.createRestResource(pathApplications, application); err != nil {
		return err
	}

	return nil
}


