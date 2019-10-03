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
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func GetAbsoluteCommandName(commandComponents ...string) string {
	commandName := commandRoot
	for _, component := range commandComponents {
		commandName = commandName + " " + component
	}
	return commandName
}

func PrintErrorMessage(writer io.Writer, message string) {
	fPrintln(writer, message)
}

func PrintError(message string, err error) {
	fmt.Printf("\n%s: %v\n", message, err)
}

func PrintInfo(message string) {
	fmt.Println(message)
}

func ExitWithError(message string, err error) {
	PrintError(message, err)
	fmt.Println()
	os.Exit(1)
}

func ExitWithErrorMessage(writer io.Writer, message string) {
	PrintErrorMessage(writer, message)
	fPrintln(writer)
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

// OpenBrowser opens up the provided URL in a browser
func OpenBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "openbsd":
		fallthrough
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		r := strings.NewReplacer("&", "^&")
		cmd = exec.Command("cmd", "/c", "start", r.Replace(url))
	}
	if cmd != nil {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			return fmt.Errorf("Failed to open browser: " + err.Error())
		}
		err = cmd.Wait()
		if err != nil {
			return fmt.Errorf("Failed to wait for open browser command to finish: " + err.Error())
		}
		return nil
	} else {
		return errors.New("unsupported platform")
	}
}

func fPrintln(writer io.Writer, message ...interface{}) {
	_, _ = fmt.Fprintln(writer, message...)
}
