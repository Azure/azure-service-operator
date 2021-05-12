#!/bin/sh

set -eu

GIT_ROOT=$(git rev-parse --show-toplevel)
TOOL_DEST=$GIT_ROOT/hack/tools

if [ ! -f "$TOOL_DEST/task" ]; then # check for local installation
    if [ ! -f "/usr/local/bin/task" ]; then # or devcontainer installation
        $GIT_ROOT/.devcontainer/install-dependencies.sh local # otherwise, install the tools
    fi
fi

export PATH="$TOOL_DEST:$TOOL_DEST/kubebuilder/bin:$PATH"

echo "Entering $SHELL with expanded PATH (use 'exit' to quit):"
echo "Try running 'task -l' to see possible commands."
$SHELL
