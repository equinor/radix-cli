# Radix CLI

Radix CLI is the command-line interface for the Radix platform, enabling users to automate application management.

For general usage documentation, see the [Radix CLI docs](https://radix.equinor.com/docs/topic-radix-cli/).

## Installation

### Linux & Mac

1. Choose a [release version](https://github.com/equinor/radix-cli/releases).
2. Download and extract the appropriate `.tar.gz` file for your platform into your `bin` folder.

| OS    | AMD64                                       | ARM64                                      |
|-------|---------------------------------------------|--------------------------------------------|
| Mac   | `radix-cli_<version>_Darwin_x86_64.tar.gz`  | `radix-cli_<version>_Darwin_arm64.tar.gz`  |
| Linux | `radix-cli_<version>_Linux_x86_64.tar.gz`   | `radix-cli_<version>_Linux_arm64.tar.gz`   |

**Example installation (Mac, AMD64):**

```bash
LATEST_VERSION=$(
    curl --silent "https://api.github.com/repos/equinor/radix-cli/releases/latest" |
        grep '"tag_name":' |
        sed -E 's/.*"v([^"]+)".*/\1/'
)

rx_tar="radix-cli_${LATEST_VERSION}_Darwin_x86_64.tar.gz"
curl -OL "https://github.com/equinor/radix-cli/releases/download/v${LATEST_VERSION}/${rx_tar}"
tar -xf "${rx_tar}"
sudo mv rx /usr/local/bin/rx
rm "${rx_tar}"
```

**Example installation (Linux, AMD64):**

```bash
LATEST_VERSION=$(
    curl --silent "https://api.github.com/repos/equinor/radix-cli/releases/latest" |
        grep '"tag_name":' |
        sed -E 's/.*"v([^"]+)".*/\1/'
)

rx_tar="radix-cli_${LATEST_VERSION}_Linux_x86_64.tar.gz"
curl -OL "https://github.com/equinor/radix-cli/releases/download/v${LATEST_VERSION}/${rx_tar}"
tar -xf "${rx_tar}"
sudo mv rx /usr/local/bin/rx
rm "${rx_tar}"
```

### Windows

1. Download the appropriate binary from [the latest release](https://github.com/equinor/radix-cli/releases/latest):
    - **AMD64:** `radix-cli_<version>_Windows_x86_64.tar.gz`
    - **ARM64:** `radix-cli_<version>_Windows_arm64.tar.gz`
2. Extract the archive using either the built-in tar command or a tool such as *WinZip*, *WinRar*, or *7zip*:

```batch
tar -xf radix-cli_1.26.0_Windows_x86_64.tar.gz
```

3. Ensure the directory containing `rx.exe` is included in your global `PATH` environment variable.

### Install with Go

```sh
go install github.com/equinor/radix-cli/cli/rx@latest
```

## Usage

### Authentication

The `login` command authenticates your CLI session with Radix. Choose the method that best fits your environment:

- **Azure Interactive Login (default):**  
  Opens a browser for interactive login.  
  ```
  rx login
  ```

- **Device Code Authentication:**  
  Uses Azure device code flow.  
  ```
  rx login --device-code
  ```

- **GitHub Workload Identity:**  
  For running in GitHub workflows.  
  ```
  rx login --github-credentials --azure-client-id <client-id>
  ```

- **Azure Client Secret:**  
  Authenticate using client ID and secret.  
  ```
  rx login --azure-client-id <client-id> --azure-client-secret <client-secret>
  ```

- **Federated Credentials:**  
  Authenticate using a federated token and client ID.  
  ```
  rx login --azure-client-id <client-id> --federated-token-file <path-to-token-file>
  ```

> **Note:** Authentication flags are mutually exclusive. Provide the appropriate combination for your scenario.

#### Logging Out

End your session with:
```
rx logout
```

### Example Commands

- **List applications:**
  ```
  rx list applications
  ```

- **Get pipeline job logs:**
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
export APP_SERVICE_ACCOUNT_TOKEN=<service-account-token> # Acquire a token with: az account get-access-token --resource 6dae42f8-4368-4678-94ff-3960e28e3630 --query accessToken -o tsv
docker run -it -e APP_SERVICE_ACCOUNT_TOKEN=$APP_SERVICE_ACCOUNT_TOKEN ghcr.io/equinor/radix/rx:latest get cluster-config --token-environment
```

### Running in GitHub Workflows

Use the [GitHub Action for Radix CLI](https://github.com/equinor/radix-github-actions) to run `rx` within GitHub workflows.

## Troubleshooting

**Problem:** Failed to acquire a token from Azure AD  
**Solution:** Remove your `$HOME/.radix` folder

```bash
rm -rf $HOME/.radix/
```

## Development

Radix CLI uses the [cobra framework](https://github.com/spf13/cobra) for command handling.

### Merging Changes

- All changes must be merged into the `master` branch via **pull requests** using **squash commits**.
- Squash commit messages must follow the [Conventional Commits](https://www.conventionalcommits.org/en/about/) specification.

### Generating Client Stubs

Client code is generated from the Radix API server's swagger contract definition using [go-swagger](https://github.com/go-swagger/go-swagger/blob/master/docs/install.md):

```sh
make generate-client
```

## Release Process

- Merging into `master` triggers the **Prepare release pull request** workflow.
    - This workflow analyzes commit messages to determine if a version bump (major, minor, patch) is needed.
    - It creates a release pull request for the new stable version (e.g., `1.2.3`).
- Merging the release pull request triggers the **Create releases and tags** workflow:
    - Reads the version from `version.txt`, creates a GitHub release, and tags it.
- The new tag triggers the **CD** workflow:
    - Builds and pushes container image tags (version and `latest`) to `ghcr.io`
    - Builds and uploads Radix CLI binaries to the GitHub release.

## Contributing

Interested in contributing? See the [contributing guide](./CONTRIBUTING.md).

## Security

See [security information](./security.md) for details.
