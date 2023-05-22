#!/usr/bin/env bash
set -euo pipefail

container_reference="$1"
cluster_version_file=$2
now="$(date -u)"

# Replace the controller:latest reference in the deployment (embedded
# in the CSV) with the SHA one.
sed -i "s!controller:latest!${container_reference}!g" $cluster_version_file

# Insert containerImage and createdAt into metadata.annotations.
yq eval -i ".metadata.annotations.containerImage = \"${container_reference}\" | .metadata.annotations.createdAt = \"${now}\"" $cluster_version_file

# Remove cert volumes and volume mounts from the CSV deployment - the
# ones here are for cert-manager, OLM will set its own to get the
# webhook certificates installed in the pod.
yq eval -i 'del(.spec.install.spec.deployments[0].spec.template.spec.containers[] | select(.name == "manager").volumeMounts)' $cluster_version_file
yq eval -i 'del(.spec.install.spec.deployments[0].spec.template.spec.volumes)' $cluster_version_file
