#!/usr/bin/env bash

set -eu

GIT_ROOT=$(git rev-parse --show-toplevel)
TOOL_DEST=$GIT_ROOT/hack/tools

# This will be fast if everything is already installed
$GIT_ROOT/.devcontainer/install-dependencies.sh --skip-installed

# Setup envtest binaries and define KUBEBUILDER_ASSETS
# NB: if you change this, .devcontainer/Dockerfile also likely needs updating
if ! ENVTEST=$("$TOOL_DEST/setup-envtest" use --print env 1.27.1) ; then
    echo "Failed to setup envtest"
    exit 1
fi

source <(echo $ENVTEST)

export PATH="$KUBEBUILDER_ASSETS:$TOOL_DEST:$PATH"

echo "Entering $SHELL with expanded PATH (use 'exit' to quit)."
echo "Try running 'task -l' to see possible commands."
$SHELL
