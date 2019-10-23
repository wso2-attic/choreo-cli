#!/usr/bin/env bash

# -----------------------------------------------------------------------
# Copyright (c) 2019, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
#
# This software is the property of WSO2 Inc. and its suppliers, if any.
# Dissemination of any information or reproduction of any material contained
# herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
# You may not alter or remove any copyright or other notice from copies of this content.
# -----------------------------------------------------------------------

source builder/cli-constants.txt

function generatePlatformArchive() {
    pushd "$TEMP_BUILD_DIRECTORY" || exit

    PLATFORM_ARCHIVE_NAME=choreo-cli-$CHOREO_CLI_VERSION-${OS_PLATFORM}-${PLATFORM_ARCHITECTURE}.tar.gz

    tar czf "$PLATFORM_ARCHIVE_NAME" "$TEMP_PLATFORM_BUILD_DIRECTORY_NAME" > /dev/null 2>&1
    rm -rf "$TEMP_PLATFORM_BUILD_DIRECTORY"

    popd || exit
}

echo "Building Choreo CLI $CHOREO_CLI_VERSION"

TARGET_PLATFORMS="linux/amd64/linux/x64 darwin/amd64/macosx/x64 windows/amd64/windows/x64"

GIT_REVISION=$(git rev-parse --verify HEAD)
BUILD_SCRIPT_ROOT=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
PROJECT_ROOT=$BUILD_SCRIPT_ROOT/..
TEMP_BUILD_DIRECTORY=$BUILD_SCRIPT_ROOT/target
TEMP_PLATFORM_BUILD_DIRECTORY_NAME=choreo-$CHOREO_CLI_VERSION
TEMP_PLATFORM_BUILD_DIRECTORY=$TEMP_BUILD_DIRECTORY/$TEMP_PLATFORM_BUILD_DIRECTORY_NAME

rm -rf "${TEMP_BUILD_DIRECTORY:?}"/*

# Go build time flags
GO_LDFLAGS="-X $PROJECT_MODULE/internal/pkg/build.buildVersion=$CHOREO_CLI_VERSION"
GO_LDFLAGS+=" -X $PROJECT_MODULE/internal/pkg/build.buildGitRevision=$GIT_REVISION"
GO_LDFLAGS+=" -X $PROJECT_MODULE/internal/pkg/build.buildTime=$(date +%Y-%m-%dT%H:%M:%S%z)"

for platform in ${TARGET_PLATFORMS}
do
    echo "Building CLI binary for $platform"

    PLATFORM_DATA_ARRAY=(${platform//\// })
    GOOS_DATA=${PLATFORM_DATA_ARRAY[0]}
    GOARCH_DATA=${PLATFORM_DATA_ARRAY[1]}
    OS_PLATFORM=${PLATFORM_DATA_ARRAY[2]}
    PLATFORM_ARCHITECTURE=${PLATFORM_DATA_ARRAY[3]}

    AGGREGATED_GO_LDFLAGS="$GO_LDFLAGS -X $PROJECT_MODULE/internal/pkg/build.buildPlatform=$GOOS_DATA/$GOARCH_DATA"

    GOOS=$GOOS_DATA GOARCH=$GOARCH_DATA go build -o "$TEMP_PLATFORM_BUILD_DIRECTORY"/bin/chor \
                            -ldflags "$AGGREGATED_GO_LDFLAGS" -x "$PROJECT_MODULE/$CHORE_CLI_SRC_ROOT"

    cp "$PROJECT_ROOT"/LICENSE "$TEMP_PLATFORM_BUILD_DIRECTORY"

    generatePlatformArchive
done

echo "Choreo CLI build completed"
ls -lh "$TEMP_BUILD_DIRECTORY"
