#!/bin/bash

# This script generates helm manifest and replaces required values in helm chart.

set -e

LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=$1
PUBLIC_REGISTRY=$2
VERSION=$3
DIR=$4

ASO_CHART="$DIR"charts/azure-service-operator
GEN_FILES_DIR="$ASO_CHART"/templates/generated
TEMP_DIR="$GEN_FILES_DIR"/temp
IF_CLUSTER="{{- if or (eq .Values.multitenant.enable false) (eq .Values.azureOperatorMode \"webhooks\") }}"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

echo "Updating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$ASO_CHART"/values.yaml

# Chart replacements
sed -i "s/\(version: \)\(.*\)/\1${VERSION//v}/g" "$ASO_CHART"/Chart.yaml  # find version key and update the value with the current version

mkdir "$TEMP_DIR"
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
