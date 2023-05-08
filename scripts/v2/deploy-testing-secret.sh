#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

MODE_SP=false
MODE_WORKLOAD_IDENTITY=false

if [ $1 = "sp" ]; then
  MODE_SP=true
fi

if [ $1 = "workloadidentity" ]; then
  MODE_WORKLOAD_IDENTITY=true
fi

if [ ! $MODE_SP ] && [ ! $MODE_WORKLOAD_IDENTITY ]; then
  echo "MODE parameter '$1' invalid. Allowed values: 'sp', or 'workloadidentity'"
  exit 1
fi

if [ $MODE_SP = true ]; then
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
fi

if [ $MODE_WORKLOAD_IDENTITY = true ]; then
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: aso-controller-settings
  namespace: azureserviceoperator-system
stringData:
  AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
  AZURE_TENANT_ID: "$AZURE_TENANT_ID"
  AZURE_CLIENT_ID: "$AZURE_MI_CLIENT_ID"
  USE_WORKLOAD_IDENTITY_AUTH: "true"
EOF
fi
