# radix-cli

NOTE: This project is currently work in progress

The command line interface for Radix, which is to enable users of Radix platform in automation around their application on the platform. This document is for developers of the Radix CLI, or anyone interested in poking around.

## How to run

### Run using docker image

```
alias rx="docker run -it -v <your home dir>/.radix:/home/radix-cli/.radix radixdev.azurecr.io/rx:latest"
```

Typically your home dir will be `/Users/<username>/` on a Mac, or `<root>\Users\<username>` on a Window machine

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
