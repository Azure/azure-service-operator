#!/bin/bash
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

# Creates a GitHub release for Azure Service Operator.
# Usage: create-release.sh <VERSION> <RELEASE_NOTES_FILE>
#   VERSION:            the release tag (e.g. v2.20.0)
#   RELEASE_NOTES_FILE: path to a file containing the release notes markdown

set -euo pipefail

if [[ $# -lt 2 ]]; then
    echo "Usage: $0 <VERSION> <RELEASE_NOTES_FILE>" >&2
    echo "  VERSION:            release tag, e.g. v2.20.0" >&2
    echo "  RELEASE_NOTES_FILE: path to a markdown file with the release notes" >&2
    exit 1
fi

VERSION="$1"
RELEASE_NOTES_FILE="$2"

# Validate version format
if [[ ! "$VERSION" =~ ^v2\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: VERSION must match pattern v2.x.y (e.g. v2.20.0), got: $VERSION" >&2
    exit 1
fi

# Validate release notes file exists
if [[ ! -f "$RELEASE_NOTES_FILE" ]]; then
    echo "Error: Release notes file not found: $RELEASE_NOTES_FILE" >&2
    exit 1
fi

# Ensure gh CLI is available
if ! command -v gh &>/dev/null; then
    echo "Error: gh CLI is required but not found. Install from https://cli.github.com/" >&2
    exit 1
fi

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

echo "Creating release $VERSION targeting main..." >&2
echo "Release notes from: $RELEASE_NOTES_FILE" >&2

# Create the release as a draft so it can be reviewed before publishing
gh release create "$VERSION" \
    --repo Azure/azure-service-operator \
    --target main \
    --title "$VERSION" \
    --notes-file "$RELEASE_NOTES_FILE" \
    --draft

echo "" >&2
echo "Draft release $VERSION created successfully." >&2
echo "Review and publish at: https://github.com/Azure/azure-service-operator/releases" >&2
