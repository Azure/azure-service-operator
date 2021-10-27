#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: aso-controller-settings
  namespace: azureserviceoperator-system
stringData:
  AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
  AZURE_TENANT_ID: "$AZURE_TENANT_ID"
  AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
  AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
EOF
