# Choreo CLI

CLI to manage integration applications with Choreo platform

## Building from source

1. Download and install go 1.13 or later version from https://golang.org/dl/

2. Clone choreo repository (somewhere _outside_ $GOPATH).

    ```
    $ git clone git@github.com:wso2/choreo-cli.git
    ```

3. Go inside the repository and execute the make goal `build-cli-all` to generate cli binaries for all platforms. The generated distributions can be found in `PROJECT/builder/target`.
    ```
    $ cd choreo-cli
    $ make build-cli-all
    $ ls builder/target
    choreo-cli-0.0.1-SNAPSHOT-linux-x64.tar.gz   choreo-cli-0.0.1-SNAPSHOT-windows-x64.tar.gz   choreo-cli-0.0.1-SNAPSHOT-macosx-x64.tar.gz
    ```
    
## Updating dependencies

Update dependency versions if required.

```
go get -u github.com/wso2/choreo-cli/cmd/choreo
```
