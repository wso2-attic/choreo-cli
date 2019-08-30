# Choreo Command Line Interface Specification v0.1.0
## Commands and operator style

When someone is dealing with high level objects, Choreo CLI uses *noun-verb* paradigm and when some one is dealing with Choreo platform itself (like to login etc), the CLI will use *verb-noun* paradigm.

## High level objects (nouns)

`Application` - Choreo app is a placeholder. At minimum this represent some code that can be deployed as one unit (service.bal/service.java/service.js). In a real world situation, an app is a composite. (e.g. App = front end component + backend component + ..) In such a situation, a single app has multiple deployable units. The application is associated with verbs like: `create, describe, …`

`Plan` - Plan is an environment topology. Dev-Ops engineers can create plans that can build environments. These are terraform templates. There will be a bunch of default plans in Choreo for developers to choose, or they can ask org’s Dev-Ops engineers to create custom plans. Associated verbs are like: `create, describe, dryrun, …` 
 
`Environment` - An environment can be created to deploy an app in Choreo. Environments can be created anywhere, by default it will be in Choreo, but developers can bring their own environments in AWS, GCP etc. (multi-cloud aspect) by selecting a plan. Associated verbs are like: `create, link, start, stop, restart, destroy, …` 

`Logs` - Top level object to deal with anytype of a log of an app or part of an app. Associated verbs are like: `tail, trace, …`

## Object (nouns) relationship

<img src="assets/images/rel.png" width="400" />


## High level actions (verbs without nouns)

`Login` - is used to log a user in to choreo through CLI

`List` - List objects in Choreo. i.e: `applications, environments, plans, ...`

`Version` - Choreo backend API version & CLI version

## Developer creating an app and running it locally

![local-run](assets/images/local-run.png) 

```
$git clone http://github.com/<user_name>/<app_name>.git
$cd <app_name>
$chor login
```

>Login to id.choreo.dev - *this is a one time thing*

>Also think about “$chor login -i” to provide un/pw through cli itself with a multi-factor auth possibly (backlog item)

```
$chor application create <app_name>
```

>Create the app in Choreo

>This always associate the origin repo with all it’s branches

>This will initialize the app with Choreo specific environment descriptor (for instrumentation etc.)

```
$git add <new files>
```

>$chor create have added new files to the app dir\

```
$git commit -m”new files and changes”
$git push origin master
```
>Assuming the dev is working on the master

```
$chor local run <app_name>
```

>This is to run the application locally.

>This will also connect the app with the Choreo cloud using a secure tunnel and will let the app to  be visualized as a sequence diagram

## Dev-Ops engineer creating a topology plan for aws|gcp|azure

![dev-ops](assets/images/devops.png) 

```
$chor plan create <plan_name> -type aws|gcp|azure
```

>choreo topology plans are created by dev-ops

>Choreo plans are topology plans. These are defined with terraform to setup an environment in IaaSes (aws/gcp/azure/K8s)

```
$chor plan dryrun <plan_name>
```

>Once you create a the devops should be able to do a dryrun of this

## Developer creating an app and deploying it remotely (in Choreo or in aws/gcp/azure..)

![remote-run](assets/images/remote-dep.png) 

>… dev-ops flow above 

```
$git clone http://github.com/<user_name>/<app_name>.git`
$cd <app_name>`
$chor login`
```

>Login to id.choreo.dev. *this is a one time thing*
>*also think about “$chor login -i” to provide un/pw through cli itself with a multi-factor possibly (backlog item)*

```
$chor application create <app_name>
```

>Create the app in Choreo

>Init the app with choreo specific environment descriptor (for instrumentation etc.)

```
$git add <new files>
```

>$chor create have added new files to the app dir

```
$git commit -m”new files and changes”
$git push origin master
```

>Assuming the dev is working on the master

```
$chor environment create <env_name> -plan <plan_name>
```

>Create an environment remotely in choreo cloud

>When creating an environment devs can link a topology plan (which is pre-configured), based on the plan then the environment will be setup in the chosen IaaS

```
$chor environment link <env_name>
```

>This links working source branch of the current app

```
$chor environment link <env_name> -b <upstream/featurebranch1>
```

>This links specific source branch of the current app

```
$chor environment start|stop|restart <env_name>
```

>This starts the app in the environment

```
$chor environment destroy <env_name>
```

>This removes the environment

```
$chor logs tail <env_name>
```

>Similarly we need to get tracing through cli



