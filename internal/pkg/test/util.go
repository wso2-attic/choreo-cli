/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package test

import "testing"

func AssertString(t *testing.T, expected string, actual string, message string) {
	if expected != actual {
		t.Errorf("%s; expected {%s} actual {%s}", message, expected, actual)
	}
}

func AssertBool(t *testing.T, expected bool, actual bool, message string) {
	if expected != actual {
		t.Errorf("%s; expected {%t} actual {%t}", message, expected, actual)
	}
}

func AssertNonNil(t *testing.T, actual interface{}, message string) {
	if actual == nil {
		t.Errorf("%s; actual value is nil", message)
	}
}
