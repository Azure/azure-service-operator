#!/bin/sh

set -eu

GIT_ROOT=$(git rev-parse --show-toplevel)
TOOL_DEST=$GIT_ROOT/hack/tools

# This will be fast if everything is already installed
$GIT_ROOT/.devcontainer/install-dependencies.sh

# Put our tools onto the path
export PATH="$TOOL_DEST:$TOOL_DEST/kubebuilder/bin:$PATH"

# For local dev, make sure we use the local version over a global install
export KUBEBUILDER_ASSETS=$TOOL_DEST/kubebuilder/bin

echo "Entering $SHELL with expanded PATH (use 'exit' to quit)."
echo "Try running 'task -l' to see possible commands."
$SHELL
