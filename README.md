# This repository is no longer maintained.

Issue reports and pull requests will not be attended.

# Choreo CLI

[![Build Status](https://github.com/wso2/choreo-cli/workflows/Build/badge.svg)](https://github.com/wso2/choreo-cli/actions?workflow=Build)

CLI to manage integration applications with Choreo platform

## Table of content
- [Getting started](docs/reference.md)
- [Building from source](#building-from-source)

## Getting started with Choreo

### Mac users - Install with Homebrew
```
brew tap wso2/choreo && brew install chor
```
### Linux/Windows users - Download binary and install
1. Download Choreo CLI from the [releases page](https://github.com/wso2/choreo-cli/releases)
2. Extract the tar.gz file
    ```
    tar -xf choreo-cli-{version}-{os}-x64.tar.gz
    ```
3. Navigate to the `bin` directory and run `$chor` command
    ```
    cd choreo-{version}/bin
    ```
    ```
    ./chor
    ```
Visit [Choreo CLI reference](docs/reference.md) for more operations to interact with Choreo 

## Building from source

1. Download and install go 1.13 or later version from https://golang.org/dl/

2. Clone choreo repository (somewhere _outside_ $GOPATH).

    ```
    $ git clone git@github.com:wso2/choreo-cli.git
    ```

3. Go inside the repository and execute the make goal `build-cli-all` to generate CLI binaries for all platforms. The generated distributions can be found in `PROJECT/builder/target`.
    ```
    $ cd choreo-cli
    $ make build-cli-all
    $ ls builder/target
    choreo-cli-0.0.3-SNAPSHOT-linux-x64.tar.gz   choreo-cli-0.0.3-SNAPSHOT-windows-x64.tar.gz   choreo-cli-0.0.3-SNAPSHOT-macosx-x64.tar.gz
    ```
    
### Updating dependencies

Update dependency versions if required.

```
go get -u github.com/wso2/choreo-cli/cmd/choreo
```

## License

Choreo CLI is licensed under the [WSO2 Commercial License](http://wso2.com/licenses).

## Copyright

(c) 2019, [WSO2 Inc.](http://www.wso2.org) All Rights Reserved.
