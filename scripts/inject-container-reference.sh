#!/usr/bin/env bash
set -euo pipefail

container_reference="$1"
now="$(date --rfc-3339=seconds)"
cluster_version_file="bundle/manifests/azure-service-operator.clusterserviceversion.yaml"

# Replace the controller:latest reference in the deployment (embedded
# in the CSV) with the SHA one.
sed -i "s!controller:latest!${container_reference}!g" $cluster_version_file

# Insert containerImage and createdAt into metadata.annotations.
yq eval -i ".metadata.annotations.containerImage = \"${container_reference}\" | .metadata.annotations.createdAt = \"${now}\"" $cluster_version_file
