# Choreo CLI

## Building from source

1. Go to the `$GOPATH/src/github.com/wso2` directory. You might have to create the missing directories.
    ```
    $ cd $GOPATH/src/github.com/wso2
    ```

2. Clone choreo repository in `$GOPATH/src/github.com/wso2`.

    ```
    $ git clone git@github.com:wso2/choreo.git
    ```

3. Go inside the repository and execute the make goal `build-cli-all` to generate cli binaries for all platforms. The generated distributions can be found in `$GOPATH/src/github.com/wso2/choreo/builder/target`.
    ```
    $ cd choreo
    $ make build-cli-all
    $ ls builder/target
    choreo-cli-0.0.1-SNAPSHOT-linux-x64.tar.gz   choreo-cli-0.0.1-SNAPSHOT-windows-x64.tar.gz   choreo-cli-0.0.1-SNAPSHOT-macosx-x64.tar.gz
    ```