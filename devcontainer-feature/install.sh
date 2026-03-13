#!/usr/bin/env bash
set -euo pipefail

VERSION="${VERSION:-latest}"
REPO="equinor/radix-cli"

# Resolve "latest" to an actual version tag
resolve_version() {
    if [ "$VERSION" = "latest" ]; then
        local api_url="https://api.github.com/repos/${REPO}/releases/latest"
        VERSION=$(curl -fsSL "$api_url" | grep '"tag_name"' | sed -E 's/.*"tag_name":\s*"v?([^"]+)".*/\1/')
        if [ -z "$VERSION" ]; then
            echo "ERROR: Failed to resolve latest version from GitHub." >&2
            exit 1
        fi
    fi
    # Strip leading 'v' if present
    VERSION="${VERSION#v}"
}

# Map kernel name to GoReleaser OS name (title case)
get_os() {
    local kernel
    kernel="$(uname -s)"
    case "$kernel" in
        Linux)  echo "Linux" ;;
        Darwin) echo "Darwin" ;;
        MINGW*|MSYS*|CYGWIN*) echo "Windows" ;;
        *)
            echo "ERROR: Unsupported operating system '${kernel}'." >&2
            exit 1
            ;;
    esac
}

# Map machine architecture to GoReleaser arch name
get_arch() {
    local machine
    machine="$(uname -m)"
    case "$machine" in
        x86_64|amd64)   echo "x86_64" ;;
        aarch64|arm64)   echo "arm64" ;;
        *)
            echo "ERROR: Unsupported architecture '${machine}'." >&2
            exit 1
            ;;
    esac
}

install_rx() {
    resolve_version

    local os arch filename url
    os="$(get_os)"
    arch="$(get_arch)"
    filename="radix-cli_${VERSION}_${os}_${arch}.tar.gz"
    url="https://github.com/${REPO}/releases/download/v${VERSION}/${filename}"

    echo "Installing Radix CLI v${VERSION} (${os}/${arch})..."

    local checksum_url="https://github.com/${REPO}/releases/download/v${VERSION}/checksums.txt"
    tmpdir="$(mktemp -d)"
    trap 'rm -rf "$tmpdir"' EXIT

    # Download archive and checksums
    curl -fsSL -o "${tmpdir}/${filename}" "$url"
    curl -fsSL -o "${tmpdir}/checksums.txt" "$checksum_url"

    # Verify checksum
    local expected_checksum actual_checksum
    expected_checksum=$(grep "${filename}" "${tmpdir}/checksums.txt" | awk '{print $1}')
    if [ -z "$expected_checksum" ]; then
        echo "WARNING: Could not find checksum for ${filename}, skipping verification." >&2
    else
        actual_checksum=$(sha256sum "${tmpdir}/${filename}" | awk '{print $1}')
        if [ "$expected_checksum" != "$actual_checksum" ]; then
            echo "ERROR: Checksum mismatch for ${filename}." >&2
            echo "  Expected: ${expected_checksum}" >&2
            echo "  Actual:   ${actual_checksum}" >&2
            exit 1
        fi
        echo "Checksum verified."
    fi

    # Extract and install
    tar -xzf "${tmpdir}/${filename}" -C "$tmpdir"
    install -m 755 "${tmpdir}/rx" /usr/local/bin/rx

    # Install bash completion
    local completion_dir="/etc/bash_completion.d"
    mkdir -p "$completion_dir"
    /usr/local/bin/rx completion bash > "${completion_dir}/rx"

    echo "Radix CLI v${VERSION} installed to /usr/local/bin/rx"
}

install_rx
