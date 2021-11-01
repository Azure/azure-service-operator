#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

MODE_SP=false
MODE_AAD_POD_IDENTITY=false

if [ $1 = "sp" ]; then
  MODE_SP=true
fi

if [ $1 = "aadpodidentity" ]; then
  MODE_AAD_POD_IDENTITY=true
fi

if [ ! $MODE_SP ] && [ ! $MODE_AAD_POD_IDENTITY ]; then
  echo "MODE parameter '$1' invalid. Allowed values: 'sp', or 'aadpodidentity'"
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

if [ $MODE_AAD_POD_IDENTITY = true ]; then
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
---
apiVersion: v1
kind: Secret
metadata:
  name: aad-sp-secret
  namespace: azureserviceoperator-system
type: Opaque
stringData:
  clientSecret: "$AZURE_CLIENT_SECRET"
---
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentity
metadata:
  name: sp-identity
  namespace: azureserviceoperator-system
spec:
  type: 1
  tenantID: "$AZURE_TENANT_ID"
  clientID: "$AZURE_CLIENT_ID"
  clientPassword: {"name":"aad-sp-secret","namespace":"azureserviceoperator-system"}
---
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentityBinding
metadata:
  name: sp-binding
  namespace: azureserviceoperator-system
spec:
  azureIdentity: sp-identity
  selector: aso-manager-binding
EOF
fi

