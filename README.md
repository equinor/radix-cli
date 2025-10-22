# radix-cli

The command line interface for Radix, which is to enable users of Radix platform in automation around their application on the platform. This document is for developers of the Radix CLI, or anyone interested in poking around.

Radix CLI in the [Radix documentation](https://radix.equinor.com/docs/topic-radix-cli/)

## Installation

### Linux or Mac

Pick a [version](https://github.com/equinor/radix-cli/releases) of the cli you want to install, and download and extract the tar.gz file into the `bin` folder like the following example (replacing the platform and architecture with the one you picked).

|                   | AMD64                                       | ARM64                                      |
| ----------------- | ------------------------------------------- | ------------------------------------------ |
| Mac               | `radix-cli_<version>_Darwin_x86_64.tar.gz`  | `radix-cli_<version>_Darwin_arm64.tar.gz`  |
| Linux             | `radix-cli_<version>_Linux_x86_64.tar.gz`   | `radix-cli_<version>_Linux_arm64.tar.gz`   |

<br/>


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

### Run as Docker image

```bash
docker run -e APP_SERVICE_ACCOUNT_TOKEN=$(az account get-access-token --resource 6dae42f8-4368-4678-94ff-3960e28e3630 | jq .accessToken -r) ghcr.io/equinor/radix/rx:latest get application --token-environment
```

### Install with Go
```sh
go install github.com/equinor/radix-cli/cli/rx@latest
```

### Windows

Visit https://github.com/equinor/radix-cli/releases/latest and download the appropriate binaries for your machine.

- **AMD64**: `radix-cli_<version>_Windows_x86_64.tar.gz`
- **ARM64**: `radix-cli_<version>_Windows_arm64.tar.gz`

Either run the tar command to extract the contents (replacing the filename with the one you downloaded)

```batch
tar -xf radix-cli_1.26.0_Windows_x86_64.tar.gz
```

or use a third-party tool like _WinZip_, _WinRar_ or _7zip_ to extract it.

Make sure the directory path you put the executable into is in the global `PATH` environment variable to use the `rx` command anywhere.

## Modes of running

There are generally two modes of running the CLI. Both cases may use configuration in your `<home>/.radix` folder:

### Interactively

CLI will use users privileges to access the Radix API server. Context information is stored in the `<home>/.radix` folder. First time you run it (i.e. `rx get logs environment -a <your application> -e <your environment>`) a prompt is provided for you to authenticated with Azure using a device code flow. A message like this appears in your terminal:

`To sign in, use a web browser to open the page https://microsoft.com/devicelogin and enter the code ABCDEFGHI to authenticate.`

### Using docker image

* Login to the packages: `docker login ghcr.io/equinor`
* Set the machine-user token to the environment variable: `export APP_SERVICE_ACCOUNT_TOKEN=<your service account token>`
* Run the command within the container (example to watch pipeline job logs with a command `rx get logs job -a your-application-name -c playground -j your-job-name`): 
```shell
docker run -it -e APP_SERVICE_ACCOUNT_TOKEN=$APP_SERVICE_ACCOUNT_TOKEN  ghcr.io/equinor/radix/rx:latest --token-environment get logs job -a your-application-name -c playground -j your-job-name
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
go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0
```
The generated code should not be checked in, but will be generated on build of the CLI. When go-swagger is installed you can generate code using this command:
```
make generate-client
```

NOTE: If there is a change to the API, you make need to point to the API environment which holds the correct swagger definition.

## Release Process

Merging a pull request into `main` triggers the **Prepare release pull request** workflow.  
This workflow analyzes the commit messages to determine whether the version number should be bumped — and if so, whether it's a major, minor, or patch change.  

It then creates a pull request for releasing a new stable version (e.g. `1.2.3`):
Merging this request triggers the **Create releases and tags** workflow, which reads the version stored in `version.txt`, creates a GitHub release, and tags it accordingly.

The new tag triggers the **CD** workflow, which:

- builds and pushes new container image tags (current version and `latest`) to `ghcr.io`
- builds and uploads Radix CLI binaries to the GitHub release.

## Contributing

Want to [contribute](./CONTRIBUTING.md)?

## Security

There is an app registration associated with the Radix CLI, `Omnia Radix CLI`, with API permissions to `Omnia Radix Web Console - Platform Clusters` to allow for the device code flow when running in interactive mode

Read this [Security information](./security.md)
