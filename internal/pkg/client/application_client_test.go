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
	"fmt"
	"testing"

	"github.com/wso2/choreo-cli/internal/pkg/cmd/runtime"
	"github.com/wso2/choreo-cli/internal/pkg/test"
)

func TestListApps(t *testing.T) {
	var b bytes.Buffer
	client := &cliClient{
		out:   &b,
		debug: &b,
		httpClient: &mockHttpClient{
			getRestResourceImpl: func(resourcePath string, v interface{}) error {
				apps := v.(*[]runtime.Application)
				*apps = append(*apps, runtime.Application{Name: "app1"})
				*apps = append(*apps, runtime.Application{Name: "app2", Description: "text"})
				return nil
			},
		},
	}

	apps, _ := client.ListApps()

	test.AssertString(t, "app1", apps[0].Name, "Incorrect app name returned")
	test.AssertString(t, "app2", apps[1].Name, "Incorrect app name returned")
	test.AssertString(t, "text", apps[1].Description, "Incorrect app description returned")
}

func TestListAppsError(t *testing.T) {
	var b bytes.Buffer
	client := &cliClient{
		out:   &b,
		debug: &b,
		httpClient: &mockHttpClient{
			getRestResourceImpl: func(resourcePath string, v interface{}) error {
				return fmt.Errorf("mock HTTP error")
			},
		},
	}

	_, err := client.ListApps()

	test.AssertNonNil(t, err, "An error should be returned")
}

func TestCreateApp(t *testing.T) {
	var actual *runtime.ApplicationRequest

	var b bytes.Buffer
	client := &cliClient{
		out:   &b,
		debug: &b,
		httpClient: &mockHttpClient{
			createRestResourceImpl: func(resourcePath string, data interface{}) error {
				actual = data.(*runtime.ApplicationRequest)
				return nil
			},
		},
	}

	_ = client.CreateNewApp("app1", "")

	test.AssertString(t, actual.Name, "app1", "App data not created correctly")
}

func TestCreateAppError(t *testing.T) {
	var b bytes.Buffer
	client := &cliClient{
		out:   &b,
		debug: &b,
		httpClient: &mockHttpClient{
			createRestResourceImpl: func(resourcePath string, data interface{}) error {
				return fmt.Errorf("mock HTTP error")
			},
		},
	}

	err := client.CreateNewApp("app1", "")

	test.AssertNonNil(t, err, "An error should be returned")
}

func TestDeployAppError(t *testing.T) {
	var b bytes.Buffer
	client := &cliClient{
		out:   &b,
		debug: &b,
		httpClient: &mockHttpClient{
			createRestResourceWithResponseImpl: func(resourcePath string, requestData interface{},
				responseData interface{}) error {
				return fmt.Errorf("mock HTTP error")
			},
		},
	}

	_, err := client.CreateAndDeployApp("http://github.com/test/test")

	test.AssertNonNil(t, err, "An error should be returned")
}

func TestDeployApp(t *testing.T) {
	var b bytes.Buffer
	client := &cliClient{
		out:   &b,
		debug: &b,
		httpClient: &mockHttpClient{
			createRestResourceWithResponseImpl: func(resourcePath string, requestData interface{},
				responseData interface{}) error {
				apps := responseData.(*runtime.DeploymentDetails)
				apps.DeploymentUrl = "http://example.com/apps/url"
				apps.ApplicationId = "appd83d56f4c6ff40428e8dd057c5b94bd5"
				return nil
			},
		},
	}

	deploymentDetails, _ := client.CreateAndDeployApp("http://github.com/test/test")

	test.AssertString(t, "http://example.com/apps/url", deploymentDetails.DeploymentUrl,
		"The app URL should be returned")
	test.AssertString(t, "appd83d56f4c6ff40428e8dd057c5b94bd5", deploymentDetails.ApplicationId,
		"The app id should be returned")
}
