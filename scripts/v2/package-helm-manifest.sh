#!/bin/bash

# This script generates helm manifest and replaces required values in helm chart.

set -e

LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=$1
PUBLIC_REGISTRY=$2
VERSION=$3
DIR=$4

ASO_CHART="$DIR"charts/azure-service-operator
GEN_FILES_DIR="$ASO_CHART"/templates/generated
IF_CLUSTER="{{- if or (eq .Values.multitenant.enable false) (eq .Values.azureOperatorMode \"webhooks\") }}"
IF_TENANT="{{- if or (eq .Values.multitenant.enable false) (eq .Values.azureOperatorMode \"watchers\") }}"
IF_CRDS="{{- if and (eq .Values.installCRDs true) (or (eq .Values.multitenant.enable false) (eq .Values.azureOperatorMode \"webhooks\")) }}"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Matches and adds helm flow control to a file
flow_control(){
  MATCHER_IF=$1
  MATCHER_END=$2
  IF_CLAUSE=$3
  TARGET=$4

  sed -i "/$MATCHER_IF/i \  \ $IF_CLAUSE" "$TARGET"
  sed -i "/$MATCHER_END/a \  \ {{- end }}" "$TARGET"
}

echo "Updating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$ASO_CHART"/values.yaml

# Chart replacements
sed -i "s/\(version: \)\(.*\)/\1${VERSION//v}/g" "$ASO_CHART"/Chart.yaml  # find version key and update the value with the current version

# Helm chart packaging, indexing and updating dependencies
echo "Packaging helm charts"
helm package "$ASO_CHART" -d "$DIR"charts # package the ASOv2 helm files into a tar file
helm repo index "$DIR"charts # update index.yaml for Helm Repository
