# Choreo CLI (*chor*) reference

Choreo CLI (code name `$chor`) reference is based on [Choreo CLI specification](spec.md)

## Chor CLI Commands
- [auth](#chor-auth) - authentication and authorization actions in Choreo
- [version](#chor-version) - get Choreo client version information
- [application](#chor-application) - manage applications

### chor auth

`auth` command is used to manage authentication and authorization in Choreo platform. 
Available sub commands include [login](#chor-auth-login).

#### Synopsis

Manage authentication and authorization.

#### Options

```
  -h, --help   help for application
```

[Back to Command List](#chor-cli-commands)

### chor auth login

`auth login` command can be used login to Choreo. This is required to
perform all the tasks that interacts with Choreo.

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
  -h, --help   help for version
```

[Back to Command List](#chor-cli-commands)

### chor version

`version` command can be used to retrieve version information 
related to the Choreo client. In addition to the cli version name, 
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
Available sub commands include [create](#chor-application-create) and [list](#chor-application-list).

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
  -h, --help   help for list
```

[Back to Command List](#chor-cli-commands)
