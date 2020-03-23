#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

# Turn colors in this script off by setting the NO_COLOR variable in your
# environment to any value:
#
# $ NO_COLOR=1 test.sh
NO_COLOR=${NO_COLOR:-""}
if [[ -z "${NO_COLOR}" ]]; then
  header=$'\e[1;33m'
  reset=$'\e[0m'
else
  header=''
  reset=''
fi

function header_text {
  echo "$header$*$reset"
}

header_text "==> Checking copyright headers <=="

files=$(find . -type f -iname '*.go' ! -path './vendor/*' ! -path './hack/tools/*' ! -path './test/e2e/vendor/*')
licRes=$(for file in $files; do
           awk 'NR<=4' "$file" | grep -Eq "(Copyright|generated|GENERATED)" || echo "$file";
         done)

if [ -n "$licRes" ]; then
        echo "Copyright header check failed:";
        echo "${licRes}";
        exit 1;
fi