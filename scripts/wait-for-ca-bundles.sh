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

# Wait for all CRDs to have CA bundles set, or 30s.
SECONDS=0

until all_crds_have_cabundle; do
    if (( SECONDS > 30 )); then
        echo "Timed out waiting for all CRDs to have CA bundles"
    fi

    echo "Waiting 5s"
    sleep 5
done
