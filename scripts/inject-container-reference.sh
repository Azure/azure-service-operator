#!/usr/bin/env bash
set -euo pipefail

sha_reference="$1"
now="$(date --rfc-3339=seconds)"
csvfile="bundle/manifests/azure-service-operator.clusterserviceversion.yaml"

# Replace the controller:latest reference in the deployment (embedded
# in the CSV) with the SHA one.
sed -i "s!controller:latest!${sha_reference}!g" $csvfile

# Insert containerImage and createdAt into metadata.annotations.
yq eval -i ".metadata.annotations.containerImage = \"${sha_reference}\" | .metadata.annotations.createdAt = \"${now}\"" $csvfile
