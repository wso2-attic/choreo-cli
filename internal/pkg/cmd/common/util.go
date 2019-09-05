/*
 * Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein is strictly forbidden, unless permitted by WSO2 in accordance with
 * the WSO2 Commercial License available at http://wso2.com/licenses. For specific
 * language governing the permissions and limitations under this license,
 * please see the license as well as any agreement you’ve entered into with
 * WSO2 governing the purchase of this software and any associated services.
 */

package common

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func GetAbsoluteCommandName(command string) string {
	return commandRoot + " " + command
}

func ExitWithError(message string, err error) {
	fmt.Printf("\n\n%s: %v\n\n", message, err)
	os.Exit(1)
}

func ExitWithErrorMessage(message string) {
	fmt.Printf("\n\n%s\n\n", message)
	os.Exit(1)
}

func GetLocalBindAddress(port int) string {
	return ":" + strconv.Itoa(port)
}

func GetFirstOpenPort(startingPort int) int {
	port := startingPort
	for connection, err := net.Dial("tcp", GetLocalBindAddress(port)); err == nil; port++ {
		_ = connection.Close()
	}
	return port
}