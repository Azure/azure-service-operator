#!/bin/bash
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

# Generates raw changelog between the last release tag and a target commit using git log.
# Usage: generate-raw-changelog.sh [--api-only] [TARGET_COMMIT]
#   --api-only:     list only the v2/api/ groups that changed (not the full changelog)
#   TARGET_COMMIT:  optional commit/ref to generate changelog up to (default: HEAD)

set -euo pipefail

API_ONLY=false
if [[ "${1:-}" == "--api-only" ]]; then
    API_ONLY=true
    shift
fi

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

TARGET="${1:-HEAD}"

# Determine the last release tag reachable from TARGET (v2.x.y format, sorted by version)
LAST_TAG=$(git tag --sort=-version:refname --merged "$TARGET" | grep -E '^v2\.[0-9]+\.[0-9]+$' | head -1 || true)
if [[ -z "$LAST_TAG" ]]; then
    echo "Error: no v2.x.y release tag found merged into $TARGET" >&2
    exit 1
fi
echo "Last release: $LAST_TAG" >&2
echo "Generating changelog for $LAST_TAG..$TARGET" >&2

if [[ "$API_ONLY" == true ]]; then
    # List only the v2/api/ groups that changed
    git diff --name-only "$LAST_TAG".."$TARGET" -- 'v2/api/' | awk -F'/' '/^v2\/api\//{print $3}' | sort -u
else
    # List commits (PRs) since the last release
    git log "$LAST_TAG".."$TARGET" --oneline
fi
