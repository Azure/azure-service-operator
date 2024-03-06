#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

print_usage() {
  echo "Usage: wait-for-operator-ready.sh [-c]"
  echo "    -c: Do NOT wait for CRDs to reach established state"
  echo "    -n: Specify the namespace for the operator (default is 'azureserviceoperator-system')"
}

OPERATOR_NAMESPACE="azureserviceoperator-system"  # Default namespace
CHECK_ESTABLISHED=1
while getopts 'c:n:' flag; do
  case "${flag}" in
    c) CHECK_ESTABLISHED=0 ;;
    n) OPERATOR_NAMESPACE="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done

function all_crds_have_cabundle() {
  for crd in $(kubectl get crd -l "app.kubernetes.io/name == azure-service-operator" -o name); do
    cabundle=$(kubectl get "$crd" -o jsonpath='{.spec.conversion.webhook.clientConfig.caBundle}')
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

function wait_for_crds_established() {
    until kubectl wait --for=condition=established --timeout=1m crd -l 'app.kubernetes.io/name == azure-service-operator'; do
      sleep 5
    done
}

if [[ "$CHECK_ESTABLISHED" -eq 1 ]]; then
  echo "Waiting for CRDs established..."
  # This has to be a timeout wrapping kubectl wait as we're racing with CRDs being added, and kubectl wait will fail if nothing matches the -l filter
  export -f wait_for_crds_established
  timeout 2m bash -c wait_for_crds_established
fi

echo "Waiting for pod ready..."
kubectl wait --for=condition=ready --timeout=3m pod -n "$OPERATOR_NAMESPACE" -l control-plane=controller-manager

echo "Waiting for CRD cabundle..."
export -f all_crds_have_cabundle
export -f wait_for_crds_cabundle
timeout 2m bash -c wait_for_crds_cabundle

echo "The operator is ready"
exit 0
