#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

# Enable tracing in this script off by setting the TRACE variable in your
# environment to any value:
#
# $ TRACE=1 test.sh
TRACE=${TRACE:-""}
if [[ -n "${TRACE}" ]]; then
  set -x
fi

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

json=$(az ad sp create-for-rbac -o json)
client_id=$(echo "$json" | jq -r '.appId')
client_secret=$(echo "$json" | jq -r '.password')
tenant_id=$(echo "$json" | jq -r '.tenant')
subscription_id=$(az account show -o json | jq -r ".id")

echo -e "AZURE_TENANT_ID=$tenant_id\nAZURE_CLIENT_ID=$client_id\nAZURE_CLIENT_SECRET=$client_secret\nAZURE_SUBSCRIPTION_ID=$subscription_id" > "$REPO_ROOT"/.env