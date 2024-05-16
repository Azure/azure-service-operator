#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

# Deal with required parameters
if [[ -n "$AZURE_MI_CLIENT_ID" ]]; then
  cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: aso-controller-settings
  namespace: tenant1-system
stringData:
  AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
  AZURE_TENANT_ID: "$AZURE_TENANT_ID"
  AZURE_CLIENT_ID: "$AZURE_MI_CLIENT_ID"
  USE_WORKLOAD_IDENTITY_AUTH: "true"
  AZURE_TARGET_NAMESPACES: "t1-items,t1-more"
EOF
else # TODO: We possibly should just remove the below path?
  cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: aso-controller-settings
  namespace: tenant1-system
stringData:
  AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
  AZURE_TENANT_ID: "$AZURE_TENANT_ID"
  AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
  AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
  AZURE_TARGET_NAMESPACES: "t1-items,t1-more"
EOF
fi
