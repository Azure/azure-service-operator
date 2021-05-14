#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

bins="./hack/tools/bin"
k="$bins/kubectl"

$k create namespace cert-manager || true
$k apply -f https://github.com/jetstack/cert-manager/releases/download/v0.13.0/cert-manager.yaml

$k create namespace k8s-infra-system || true
cat <<EOF | $k apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: manager-bootstrap-credentials
  namespace: k8s-infra-system
type: Opaque
data:
  subscription-id: $(echo "${AZURE_SUBSCRIPTION_ID}" | tr -d '\n' | base64 | tr -d '\n')
  tenant-id: $(echo "${AZURE_TENANT_ID}" | tr -d '\n' | base64 | tr -d '\n')
  client-id: $(echo "${AZURE_CLIENT_ID}" | tr -d '\n' | base64 | tr -d '\n')
  client-secret: $(echo "${AZURE_CLIENT_SECRET}" | tr -d '\n' | base64 | tr -d '\n')
EOF

sleep 10 # wait for the resources to be created, so we can wait for them

$k wait --for=condition=Ready pod --namespace cert-manager --all --timeout 60s