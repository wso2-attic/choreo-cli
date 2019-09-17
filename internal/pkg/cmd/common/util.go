/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
 * You may not alter or remove any copyright or other notice from copies of this content.
 */

package common

import (
	"fmt"
	"os"
)

func GetAbsoluteCommandName(command string) string {
	return commandRoot + " " + command
}

func PrintErrorMessage(message string) {
	fmt.Printf("\n\n%s\n\n", message)
}

func PrintError(message string, err error) {
	fmt.Printf("\n\n%s: %v\n\n", message, err)
}

func PrintInfo(message string) {
	fmt.Println(message)
}

func ExitWithError(message string, err error) {
	PrintError(message, err)
	os.Exit(1)
}

func ExitWithErrorMessage(message string) {
	PrintErrorMessage(message)
	os.Exit(1)
}