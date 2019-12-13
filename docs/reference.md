# Choreo CLI (*chor*) reference

Choreo CLI (code name `$chor`) reference is based on [Choreo CLI specification](../spec)

## Chor CLI Commands
- [auth](#chor-auth) - authentication and authorization actions in Choreo
    - [login](#chor-auth-login) - login to Choreo
    - [connect](#chor-auth-connect) - connect a source code provider
- [version](#chor-version) - get Choreo client version information
- [application](#chor-application) - manage applications
    - [create](#chor-application-create) - create an application
    - [list](#chor-application-list) - list applications
    - [deploy](#chor-application-deploy) - deploy a Ballerina application
    - [logs](#chor-application-logs) - manage application logs
        - [show](#chor-application-logs-show) - show application logs
    - [delete](#chor-application-delete) - delete an application

### chor auth

`auth` command is used to manage authentication and authorization in Choreo platform.

#### Synopsis

Manage authentication and authorization.

#### Options

```
  -h, --help   help for auth
```

[Back to Command List](#chor-cli-commands)

### chor auth login

`auth login` command can be used to login to Choreo. This is required to
perform all the tasks that interact with Choreo.

#### Synopsis

Login to Choreo.

```
auth login
```

#### Examples

```
$ chor auth login
```

#### Options

```
  -h, --help   help for login
```

[Back to Command List](#chor-cli-commands)

### chor auth connect

`auth connect` command is used to connect a source code provider to Choreo. 
At the moment only GitHub is supported.

#### Synopsis

Connect a source code provider.

```
auth connect SOURCE_PROVIDER
```

#### Examples

```
$ chor auth connect github
```

#### Options

```
  -h, --help   help for connect
```

[Back to Command List](#chor-cli-commands)

### chor version

`version` command can be used to retrieve version information 
related to the Choreo client. In addition to the CLI version name, 
Git commit hash, built date and target platform details are also printed.

#### Synopsis

Get Choreo client version information.

```
version
```

#### Examples

```
$ chor version
 Version:		0.0.1
 Git commit:		b086b964ae81e8277842fad625784672bb44a3a7
 Built:			2019-08-15T11:06:22+0530
 OS/Arch:		linux/amd64
```

#### Options

```
  -h, --help   help for version
```

[Back to Command List](#chor-cli-commands)

### chor application

`application` command is used to manage applications created with the Choreo platform.

#### Synopsis

Manage applications.

### Aliases

app

#### Options

```
  -h, --help   help for application
```

[Back to Command List](#chor-cli-commands)

### chor application create

`application create` command is used to create a new application with the Choreo platform.

#### Synopsis

Create an application.

```
application create APP_NAME [options]
```

#### Examples

```
$ chor application create app1 -d "My first app"
```

#### Options

```
  -d, --description string   Specify description for the application
  -h, --help                 help for create
```

[Back to Command List](#chor-cli-commands)

### chor application list

`application list` command is used to list applications created with the Choreo platform.

#### Synopsis

List applications.

```
application list
```

#### Examples

```
$ chor application list
```

#### Options

```
  -h, --help    help for list
```

[Back to Command List](#chor-cli-commands)

### chor application deploy

`application deploy` command is used to deploy a Ballerina application to the Choreo platform.

#### Synopsis

Deploy a Ballerina application.

```
application deploy GITHUB_URL
```

#### Examples

```
$ chor application deploy https://github.com/someuser/choreo-sample
$ chor application deploy -n my-app https://github.com/someuser/choreo-sample
```

#### Options

```
  -h, --help          help for deploy
  -n, --name string   the name to be used for the created application
```

[Back to Command List](#chor-cli-commands)

### chor application logs

`application logs` command is used to manage logs of a deployed application.

#### Synopsis

Manage application logs.

#### Options

```
  -h, --help    help for list
```

[Back to Command List](#chor-cli-commands)

### chor application logs show

`application logs show APP_ID` command is used to show logs of a deployed application.

#### Synopsis

Show logs of a deployed application. Maximum number of log lines shown is 500.
If the number of log lines is not specified 50 lines is shown by default.

```
application logs show APP_ID
```

#### Examples

```
$ chor application logs show app1234567890abcd
```

#### Options

```
  -n, --number-of-lines     specify number of log lines which should be fetched
  -h, --help                help for list
```

[Back to Command List](#chor-cli-commands)

### chor application delete

`application delete` command is used to delete an application created in the Choreo platform. 
It will also remove the app deployment if the application is already deployed.

#### Synopsis

Delete an application.

```
application delete APP_ID
```

#### Examples

```
$ chor application delete a123456788901
```

#### Options

```
  -h, --help          help for deploy
```

[Back to Command List](#chor-cli-commands)

## Global flags 
```
  -v, --verbose     verbose output
```

[Back to Command List](#chor-cli-commands)
