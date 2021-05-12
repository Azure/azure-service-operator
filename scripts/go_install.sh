#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

if [ -z "${1}" ]; then
  echo "must provide module as first parameter"
  exit 1
fi

if [ -z "${GOBIN}" ]; then
  echo "GOBIN is not set. Must set GOBIN to install the bin in a specified directory."
  exit 1
fi

tmp_dir=$(mktemp -d -t goinstall_XXXXXXXXXX)
function clean {
  rm -rf "${tmp_dir}"
}
trap clean EXIT
cd "${tmp_dir}"

# create a new module in the tmp directory
go mod init fake/mod

# install the golang module specified as the first argument
go get "${1}"
