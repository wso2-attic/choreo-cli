/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package client

type mockHttpClient struct {
	getRestResourceImpl                func(resourcePath string, v interface{}) error
	createRestResourceImpl             func(resourcePath string, data interface{}) error
	createRestResourceWithResponseImpl func(resourcePath string, requestData interface{},
		responseData interface{}) error
}

func (c *mockHttpClient) getRestResource(resourcePath string, v interface{}) error {
	if c.getRestResourceImpl != nil {
		return c.getRestResourceImpl(resourcePath, v)
	}
	return nil
}

func (c *mockHttpClient) createRestResource(resourcePath string, data interface{}) error {
	if c.createRestResourceImpl != nil {
		return c.createRestResourceImpl(resourcePath, data)
	}
	return nil
}

func (c *mockHttpClient) createRestResourceWithResponse(resourcePath string,
	requestData interface{}, responseData interface{}) error {
	if c.createRestResourceWithResponseImpl != nil {
		return c.createRestResourceWithResponseImpl(resourcePath, requestData, responseData)
	}
	return nil
}
