#!/bin/sh

set -eu

GIT_ROOT=$(git rev-parse --show-toplevel)
TOOL_DEST=$GIT_ROOT/hack/tools

# This will be fast if everything is already installed
$GIT_ROOT/.devcontainer/install-dependencies.sh --skip-installed

# Setup envtest binaries
# NB: if you change this, .devcontainer/Dockerfile also likely needs updating
source <(setup-envtest use -i -p env 1.23.5) # this sets KUBEBUILDER_ASSETS
export PATH="$KUBEBUILDER_ASSETS:$TOOL_DEST:$PATH"

echo "Entering $SHELL with expanded PATH (use 'exit' to quit)."
echo "Try running 'task -l' to see possible commands."
$SHELL
