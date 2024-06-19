#!/bin/bash

# package-helm-manifest.sh script is used to copy the generated files by kustomize and package the helm chart.
# The generated files are updated when a new resource is added. To eliminate the manual process of updating generated files, we use this script for automation.
# Generated files include below files:
# - admissionregistration.k8s.io_v1_mutatingwebhookconfiguration_azureserviceoperator-mutating-webhook-configuration.yaml
# - admissionregistration.k8s.io_v1_validatingwebhookconfiguration_azureserviceoperator-validating-webhook-configuration.yaml
# - rbac.authorization.k8s.io_v1_clusterrole_azureserviceoperator-manager-role.yaml

# Above files are always updated when a new resource is added

set -e

print_usage() {
  echo "Usage: package-helm-manifest.sh -d <DIRECTORY> -v <VERSION> -r <PUBLIC_REGISTRY> -i <LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE>"
}

LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=
PUBLIC_REGISTRY=
VERSION=
DIR=

while getopts 'i:r:v:d:' flag; do
  case "${flag}" in
    i) LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE="${OPTARG}" ;;
    r) PUBLIC_REGISTRY="${OPTARG}" ;;
    v) VERSION="${OPTARG}" ;;
    d) DIR="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done

if [[ -z "$DIR" ]] || [[ -z "$VERSION" ]] || [[ -z "$PUBLIC_REGISTRY" ]] || [[ -z "$LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE" ]]; then
  print_usage
  exit 1
fi

ASO_CHART="$DIR"charts/azure-service-operator
GEN_FILES_DIR="$ASO_CHART"/templates/generated
TEMP_DIR="$GEN_FILES_DIR"/temp
IF_CLUSTER="{{- if or (eq .Values.multitenant.enable false) (eq .Values.azureOperatorMode \"webhooks\") }}"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

echo "Updating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$ASO_CHART"/values.yaml

# Chart replacements
sed -i "s/\(version: \)\(.*\)/\1${VERSION//v}/g" "$ASO_CHART"/Chart.yaml  # find version key and update the value with the current version

mkdir -p "$TEMP_DIR"
${SCRIPT_DIR}/kustomize-build.sh -v "$VERSION" -k operator -o "$TEMP_DIR"

echo "Making sed replacements and copying generated yamls"
for file in $(find "$TEMP_DIR" -type f)
do
  # Append cluster or tenant guards to each file
  if [[ $file == *"mutating-webhook-configuration"* ]] ||
     [[ $file == *"validating-webhook-configuration"* ]] ||
     [[ $file == *"azureserviceoperator-manager-role.yaml" ]] ; then
        sed -i "1 s/^/$IF_CLUSTER\n/;$ a {{- end }}" "$file"
        sed -i 's/azureserviceoperator-system/{{ .Release.Namespace }}/g' "$file"
        mv -f "$file" "$GEN_FILES_DIR"
  fi
done

rm -rf "$TEMP_DIR"

# Helm chart packaging, indexing and updating dependencies
echo "Packaging helm charts"
helm package "$ASO_CHART" -d "$DIR"charts # package the ASOv2 helm files into a tar file
helm repo index "$DIR"charts # update index.yaml for Helm Repository
