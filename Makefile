# --------------------------------------------------------------------
# Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
#
#  This software is the property of WSO2 Inc. and its suppliers, if any.
#  Dissemination of any information or reproduction of any material contained
#  herein is strictly forbidden, unless permitted by WSO2 in accordance with
#  the WSO2 Commercial License available at http://wso2.com/licenses. For specific
#  language governing the permissions and limitations under this license,
#  please see the license as well as any agreement you’ve entered into with
#  WSO2 governing the purchase of this software and any associated services.
# -----------------------------------------------------------------------

#### NOTE: Use tab characters instead of spaces for indentation ###

include builder/cli-constants.txt

PROJECT_ROOT := $(realpath $(dir $(abspath $(firstword $(MAKEFILE_LIST)))))
GO_BUILD_DIRECTORY := $(PROJECT_ROOT)/builder/target
GIT_REVISION := $(shell git rev-parse --verify HEAD)

INSTALLER_VERSION ?= $(GIT_REVISION)

# Go build time flags
GO_LDFLAGS := -X $(PROJECT_MODULE)/internal/pkg/build.buildVersion=$(CHOREO_CLI_VERSION)
GO_LDFLAGS += -X $(PROJECT_MODULE)/internal/pkg/build.buildGitRevision=$(GIT_REVISION)
GO_LDFLAGS += -X $(PROJECT_MODULE)/internal/pkg/build.buildTime=$(shell date +%Y-%m-%dT%H:%M:%S%z)
GO_LDFLAGS += -X $(PROJECT_MODULE)/internal/pkg/build.buildPlatform=local

.PHONY: test
test:
	go test $(shell go list ./...)

.PHONY: build-cli
build-cli: clean-cli
	go build -o ${GO_BUILD_DIRECTORY}/chor -ldflags "$(GO_LDFLAGS)" -x $(PROJECT_MODULE)/$(CHORE_CLI_SRC_ROOT)

.PHONY: install-cli
install-cli: clean-cli
	go build -o ${GO_BUILD_DIRECTORY}/chor -ldflags "$(GO_LDFLAGS)" -x $(PROJECT_MODULE)/$(CHORE_CLI_SRC_ROOT); \
	sudo cp ${GO_BUILD_DIRECTORY}/chor /usr/local/bin; \

.PHONY: build-cli-all
build-cli-all: clean-cli
	builder/cli-all-release.sh

.PHONY: clean-cli
clean-cli:
	rm -rf ${GO_BUILD_DIRECTORY}/*
