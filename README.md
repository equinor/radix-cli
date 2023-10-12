# radix-cli

The command line interface for Radix, which is to enable users of Radix platform in automation around their application on the platform. This document is for developers of the Radix CLI, or anyone interested in poking around.

Radix CLI in the [Radix documentation](https://radix.equinor.com/docs/topic-radix-cli/)

## Installation

### Linux or Mac

#### Binaries

Pick the appropriate binaries for your machine

`radix-cli_<version>_Darwin_i386.tar.gz` `radix-cli_<version>_Darwin_x86_64.tar.gz`

`radix-cli_<version>_Linux_i386.tar.gz` `radix-cli_<version>_Linux_x86_64.tar.gz`

`radix-cli_<version>_Linux_armv6.tar.gz` `radix-cli_<version>_Linux_arm64.tar.gz`

Pick a [version](https://github.com/equinor/radix-cli/releases) of the cli you want to install, or the latest version, then download and extract the tar file into the `bin` folder like the following example (replacing the version and architecture with the one you picked).

```bash
LATEST_VERSION=$(
    curl --silent "https://api.github.com/repos/equinor/radix-cli/releases/latest" |
        grep '"tag_name":' |
        sed -E 's/.*"v([^"]+)".*/\1/'
)

rx_tar=radix-cli_${LATEST_VERSION}_Darwin_x86_64.tar.gz
sudo curl -OL "https://github.com/equinor/radix-cli/releases/download/v${LATEST_VERSION}/${rx_tar}"
tar -xf ${rx_tar}

mv rx /usr/local/bin/rx
rm ${rx_tar}
```

#### Or run using Docker image

Authenticate with github via docker using a token with _read:packages_ access. Make sure you also enable single sign-on for Equinor after [generating your token](https://github.com/settings/tokens). Replace `<github username>` and `<access token>`.

```bash
docker login -u <github username> -p <access token> ghcr.io

alias rx="docker run -it -v ${HOME}/.radix:/home/radix-cli/.radix ghcr.io/equinor/radix/rx:latest"
```

(Typically your `HOME` variable will be `/Users/<username>` on a Mac and `/home/<username>` on Linux)

### Windows

#### Binaries

Visit https://github.com/equinor/radix-cli/releases/latest and download the appropriate binaries for your machine.

`radix-cli_<version>_Windows_i386.tar.gz` (32 bit)
`radix-cli_<version>_Windows_x86_64.tar.gz` (64 bit)

Either run the tar command to extract the contents (replacing the filename with the one you downloaded)

```batch
tar -xf radix-cli_0.0.16_Windows_x86_64.tar.gz
```

or use a third-party tool like _WinZip_, _WinRar_ or _7zip_ to extract it.

Make sure the directory path you put the executable into is in the global `PATH` environment variable to use the `rx` command anywhere.

#### Or run using Docker image

See docker for linux/mac above for authentication guide.

If your terminal has a profile or auto-run script, you can add the following to it:

```batch
DOSKEY rx=docker run -it -v %HOME%:/home/radix-cli ghcr.io/equinor/radix/rx:latest $*
```

If not, you must add a new script file called `rx.bat` in a directory, present in `PATH`, with the following content

```batch
docker run -it -v %HOME%:/home/radix-cli ghcr.io/equinor/radix/rx:latest $*
```

(Typically your `HOME` variable will be `C:\Users\<username>`)

## Modes of running

There are generally two modes of running the CLI. Both cases may use configuration in your `<home>/.radix` folder:

### Interactively

CLI will use users privileges to access the Radix API server. Context information is stored in the `<home>/.radix` folder. First time you run it (i.e. `rx get logs environment -a <your application> -e <your environment>`) a prompt is provided for you to authenticated with Azure using a device code flow. A message like this appears in your terminal:

`To sign in, use a web browser to open the page https://microsoft.com/devicelogin and enter the code ABCDEFGHI to authenticate.`

### Using a machine user

CLI can also use a machine user for authenticating with the API server. This will be done through a bearer token of a service account connected to your application. The service account token will be provided to you under configuration page of your application. For more information on this see [this](https://www.radix.equinor.com/guides/deploy-only/#machine-user-token) guide. There are two ways to feed this token to the CLI. Either using standard input. This should be done like this:

`echo <your service account token> | rx --token-stdin list applications`

Alternatively, you can use an environment variable for the token:

```
export APP_SERVICE_ACCOUNT_TOKEN=<your service account token>
rx --token-environment get application
```

Note that using your own token obtained through `az account get-access-token` may not work, because the size of the token may be too big.

### Using docker image

* Login to the packages: `docker login ghcr.io/equinor`
* Set the machine-user token to the environment variable: `export APP_SERVICE_ACCOUNT_TOKEN=<your service account token>`
* Run the command within the container (example to watch pipeline job logs with a command `rx get logs job -a your-application-name -c playground -j your-job-name`): 
```shell
docker run -it -e APP_SERVICE_ACCOUNT_TOKEN=$APP_SERVICE_ACCOUNT_TOKEN  ghcr.io/equinor/radix-cli/rx:latest --token-environment get logs job -a your-application-name -c playground -j your-job-name
```

## Problems encountered

Problem: Failed to acquire a token from Azure AD
Solution: Remove your `<home>/.radix` folder

```
rm -rf $HOME/.radix/
```

## Development

We are using the [cobra framework](https://github.com/spf13/cobra) for handling commands. Add a command by:

```
cobra add <commandName>
```

### Generate client stubs

Client code is generated from swagger contract definition of the latest contract of the Radix API server. We use [go-swagger](https://github.com/go-swagger/go-swagger/blob/master/docs/install.md). Install it by:
```
go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
```
The generated code should not be checked in, but will be generated on build of the CLI. When go-swagger is installed you can generate code using this command:
```
make generate-client
```

NOTE: If there is a change to the API, you make need to point to the API environment which holds the correct swagger definition.

### Building and releasing

We are making releases available as GitHub releases using [go-releaser](https://goreleaser.com/). The release process is controlled by the `.goreleaser.yml` file. 

To make a release:
1. Set the version number in the constant `version` in the file `cmd/root.go`. The version will be shown with the command `rx --version`
2. Create and push the new version as a tag: `git tag v0.0.1` and `git push origin v0.0.1`
3. If something goes wrong:
   - open the GitHub repository and delete [created tag](https://github.com/equinor/radix-cli/tags/) (with release)
   - delete it locally ` git tag -d v0.0.1`
   - reset changes `git reset --hard`
   - tag the commit againg and push: `git tag v0.0.1` and `git push origin v0.0.1`

To generate a local version for debugging purposes, it can be built using:

```
CGO_ENABLED=0 GOOS=darwin go build -ldflags "-s -w" -a -installsuffix cgo -o ./rx
```

### Security

There is an app registration associated with the Radix CLI, `Omnia Radix CLI`, with API permissions to `Omnia Radix Web Console - Platform Clusters` to allow for the device code flow when running in interactive mode
