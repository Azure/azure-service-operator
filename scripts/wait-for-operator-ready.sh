#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

function all_crds_have_cabundle() {
  for crd in $(kubectl api-resources -o name | grep '\.azure\.com'); do
    cabundle=$(kubectl get crd "$crd" -o jsonpath='{.spec.conversion.webhook.clientConfig.caBundle}')
    if [ -z "$cabundle" ]; then
      echo "$crd has no CA bundle"
      return 1
    fi
  done
  return 0
}

function wait_for_crds_cabundle() {
  until all_crds_have_cabundle; do
    sleep 5
  done
}

kubectl wait --for=condition=ready --timeout=2m installedresourcedefinitions -n azureserviceoperator-system aso-installed-resources
kubectl wait --for=condition=established --timeout=1m crd -l 'serviceoperator.azure.com/version'
kubectl wait --for=condition=ready --timeout=2m pod -n azureserviceoperator-system -l control-plane=controller-manager

# Wait for all CRDs to have CA bundles set, or 30s.
SECONDS=0

export -f all_crds_have_cabundle
export -f wait_for_crds_cabundle
timeout 1m bash -c wait_for_crds_cabundle

echo "All CRDs have CA bundles"
exit 0
