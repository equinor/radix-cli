# radix-cli

The command line interface for Radix, which is to enable users of Radix platform in automation around their application on the platform. This document is for developers of the Radix CLI, or anyone interested in poking around.

Radix CLI in the [Radix documentation](https://radix.equinor.com/docs/topic-radix-cli/)

## Installation

### Linux or Mac

Pick a [version](https://github.com/equinor/radix-cli/releases) of the CLI you want to install, and download and extract the tar.gz file into the `bin` folder as shown in the following example.

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

### Install with Go

```sh
go install github.com/equinor/radix-cli/cli/rx@latest
```

## Usage Guide

### Logging In

The `login` command authenticates your CLI session with Radix. There are several authentication options to suit different environments and use cases:

#### Authentication Methods

- **Azure Interactive Login** (default):  
  Authenticate interactively via your browser. If no flags are specified, this method is used.
  ```
  rx login
  ```

- **Device Code Authentication**:  
  Authenticate using Azure Device Code.
  ```
  rx login --device-code
  ```

- **GitHub Workload Identity**:  
  Authenticate using GitHub credentials (Workload Identity). This authentication method is only valid when running `rx` in a GitHub workflow. See [GitHub Action for Radix CLI](https://github.com/equinor/radix-github-actions) for more information.
  ```
  rx login --github-credentials --azure-client-id <client-id>
  ```

- **Azure Client Secret**:  
  Authenticate using Azure Client ID and Client Secret.
  ```
  rx login --azure-client-id <client-id> --azure-client-secret <client-secret>
  ```

- **Federated Credentials**:  
  Authenticate using a federated token file and Azure Client ID.
  ```
  rx login --azure-client-id <client-id> --federated-token-file <path-to-token-file>
  ```

> **Note:** The flags above are mutually exclusive. You must provide the appropriate combinations for your authentication scenario.

### Logging Out

To end your session:
```
rx logout
```

### Example Commands

- **List applications:**
  ```
  rx list applications
  ```

- **Get logs for a pipeline job:**
  ```
  rx get logs pipeline-job --application <app-name> --job <job-name>
  ```

- **Validate a Radix config file:**
  ```
  rx validate radix-config --config-file ./radixconfig.yaml
  ```

### Running via Docker

You can run `radix-cli` in a container:

```bash
export APP_SERVICE_ACCOUNT_TOKEN=<service-account-token> # You can acquire a token with: az account get-access-token --resource 6dae42f8-4368-4678-94ff-3960e28e3630 --query accessToken -o tsv
docker run -it -e APP_SERVICE_ACCOUNT_TOKEN=$APP_SERVICE_ACCOUNT_TOKEN ghcr.io/equinor/radix/rx:latest get cluster-config --token-environment
```

### Running in GitHub workflows

To run `rx` in GitHub workflows, we recommend using the [GitHub Action for Radix CLI](https://github.com/equinor/radix-github-actions).

### Problems encountered

Problem: Failed to acquire a token from Azure AD
Solution: Remove your `<home>/.radix` folder

```
rm -rf $HOME/.radix/
```

## Development

We use the [cobra framework](https://github.com/spf13/cobra) for handling commands.

### ✅ Merging Changes

All changes must be merged into the `master` branch using **pull requests** with **squash commits**.

The squash commit message must follow the [Conventional Commits](https://www.conventionalcommits.org/en/about/) specification.

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
