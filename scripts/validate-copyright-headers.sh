#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

echo "==> Checking copyright headers <=="

regx="(Copyright.*Microsoft.*Licensed.*MIT)|(generated)"
files=$(find . -type f -iname '*.go' ! -path './vendor/*' ! -path './hack/tools/*' ! -path './test/e2e/vendor/*')
licRes=$(for file in $files; do
           head -4 "$file" | tr -d '\n' | grep -Eiq "$regx" || echo "$file";
         done)

if [ -n "$licRes" ]; then
        echo "Copyright header check failed:";
        echo "${licRes}";
        exit 1;
fi