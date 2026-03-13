#!/usr/bin/env bash
set -euo pipefail

#
# Publish devcontainer feature as an OCI artifact to ghcr.io using ORAS CLI
#
# Usage: ./publish.sh [namespace]
#   e.g. ./publish.sh richard87/radix-cli
#
# Requires: gh (authenticated), oras, jq, tar
#   Install oras: brew install oras
#

NAMESPACE="${1:-richard87/radix-cli}"
REGISTRY="ghcr.io"
FEATURE_DIR="$(cd "$(dirname "$0")" && pwd)"
FEATURE_JSON="$FEATURE_DIR/devcontainer-feature.json"

FEATURE_ID=$(jq -r '.id' "$FEATURE_JSON")
VERSION=$(jq -r '.version' "$FEATURE_JSON")
REPO="${REGISTRY}/${NAMESPACE}/${FEATURE_ID}"

MAJOR="${VERSION%%.*}"
MINOR="${VERSION%.*}"

echo "Publishing ${REPO}:${VERSION}"

# --- Auth: login to ghcr.io via oras using gh token ---
gh auth token | oras login "${REGISTRY}" -u richard87 --password-stdin

# --- Package feature as tgz ---
TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

TGZ_FILE="$TMPDIR/devcontainer-feature-${FEATURE_ID}.tgz"
tar -czf "$TGZ_FILE" -C "$FEATURE_DIR" devcontainer-feature.json install.sh

# --- Create empty config blob ---
CONFIG_FILE="$TMPDIR/config.json"
echo -n '{}' > "$CONFIG_FILE"

# --- Build annotations ---
METADATA=$(jq -c '.' "$FEATURE_JSON")
ANNOTATIONS="dev.containers.metadata=${METADATA},com.github.package.type=devcontainer_feature"

TGZ_BASENAME="devcontainer-feature-${FEATURE_ID}.tgz"

# --- Push with oras (use relative paths) ---
pushd "$TMPDIR" > /dev/null
oras push \
  --config "config.json:application/vnd.devcontainers" \
  --annotation "$ANNOTATIONS" \
  "${REPO}:${VERSION}" \
  "${TGZ_BASENAME}:application/vnd.devcontainers.layer.v1+tar"
popd > /dev/null

echo "Pushed ${REPO}:${VERSION}"

# --- Tag additional versions ---
for TAG in "$MINOR" "$MAJOR" "latest"; do
  oras tag "${REPO}:${VERSION}" "$TAG"
  echo "Tagged ${REPO}:${TAG}"
done

echo "Done! Published ${REPO}:${VERSION}"
