# radix-cli

NOTE: This project is currently work in progress

The command line interface for Radix, which is to enable users of Radix platform in automation around their application on the platform. This document is for developers of the Radix CLI, or anyone interested in poking around.

## How to run

### Run using docker image

```
alias rx="docker run -it -v <your home dir>/.radix:/home/radix-cli/.radix radixdev.azurecr.io/rx:latest"
```

Typically your home dir will be `/Users/<username>/` on a Mac, or `<root>\Users\<username>` on a Window machine

### Modes of running

There are generally two modes of running the CLI:

1. Interactively

CLI will use users privileges to access the Radix API server. Context information is stored in the \$HOME/.radix folder. First time you run it (i.e. `rx list applications`) a prompt is provided for you to authenticated with Azure using a device code flow. A message like this appears in your terminal:

`To sign in, use a web browser to open the page https://microsoft.com/devicelogin and enter the code ABCDEFGHI to authenticate.`

2. Using a machine user

CLI can also use a machine user for authenticating with the API server. This will be done through a bearer token of a service account connected to your application. The service account token will be provided to you under configuration page of your application. There are two ways to feed this token to the CLI. Either using standard input. The should be done like this:

`echo <your service account token> | rx --token-stdin list applications`

Alternatively, you can use an environment variable for the token:

```
export APP_SERVICE_ACCOUNT_TOKEN=<your service account token>
rx --token-environment list applications
```

### Available commands

```
  build       Will trigger build of a Radix application
  get-context Gets current context (used in interactive mode)
  help        Help about any command
  list        Lists Radix resources
  set-context Sets the context to be either production, playground or development (used in interactive mode)
```

### Global command arguments

These are global arguments for all commands. Default will use context=production, unless otherwise stated. --cluster and --environment are meant for Radix platform developers only, to test against a custom cluster and api environment

```
General flags:
  -k, --cluster string       Set cluster to override context
  -c, --context string       Use context production|playground|development regardless of current context
  -e, --environment string   The API environment to run with (default "prod")
      --from-config          Read and use radix config from file as context
  -h, --help                 help for rx
      --token-environment    Take the token from environment variable APP_SERVICE_ACCOUNT_TOKEN
      --token-stdin          Take the token from stdin
```

## Generate client stubs

Client code is generated from swagger contract definition of the latest contract of the Radix API server. We use go-swagger (https://github.com/go-swagger/go-swagger/blob/master/docs/install.md). The generated code should not be checked in, but will be generated on build of the CLI. When go-swagger is installed you can generate code using this command:

```
make generate client
```

## Building and releasing

We are making releases available as github releases using go-releaser (https://goreleaser.com/). The release process is controlled by the .goreleaser.yml file. To make a release:

```
make release VERSION=0.0.1 RELASE_NOTE="First release"
```

## Development

We are using the cobra framework for handling commands (https://github.com/spf13/cobra). Add a command by:

```
cobra add <commandName>
```
