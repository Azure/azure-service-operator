#!/bin/bash

# This script generates helm manifest and replaces required values in helm chart.

set -e

KUBE_RBAC_PROXY=$1
LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=$2
PUBLIC_REGISTRY=$3
VERSION=$4
DIR=$5

# Manifest purge and generation
echo "Generating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$DIR"charts/azure-service-operator/values.yaml
rm -rf "$DIR"charts/azure-service-operator/templates/generated # remove generated files
rm -rf "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/templates
mkdir -p "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/templates/crds # create dirs for generated files
mkdir "$DIR"charts/azure-service-operator/templates/generated
kustomize build "$DIR"config/default -o "$DIR"charts/azure-service-operator/templates/generated
rm "$DIR"charts/azure-service-operator/templates/generated/*_namespace_* # remove namespace as we will let Helm manage it

# Sed Replacements
grep -E $KUBE_RBAC_PROXY "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* > /dev/null # Ensure that what we're about to try to replace actually exists (if it doesn't we want to fail)
sed -i "s@$KUBE_RBAC_PROXY.*@{{.Values.image.kubeRBACProxy}}@g" "$DIR"charts/azure-service-operator/templates/generated/*_deployment_*
sed -i "s@$LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE@{{.Values.image.repository}}@g" "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # Replace hardcoded ASO image
sed -i '/default-logs-container: manager/a \  \ {{ toYaml .Values.podAnnotations | indent 6 }}' "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # Add pod annotations
sed -i "s/\(version: \)\(.*\)/\1$VERSION/g" "$DIR"charts/azure-service-operator/Chart.yaml  # find version key and update the value with the current version for both main and subchart

# Metrics Configuration
sed -i '/metrics-addr/i \  \ {{if .Values.metrics.enable}}' "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # Add metrics flow control
sed -i "1,/metrics-addr=.*/s/\(metrics-addr=\)\(.*\)/\1{{ tpl .Values.metrics.address . }}/g" "$DIR"charts/azure-service-operator/templates/generated/*_deployment_*
sed -i '/metrics-addr/a \  \ {{ end }}' "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # End metrics flow control
sed -i 's/containerPort: 8080/containerPort: {{ .Values.metrics.port | default 8080 }}/g' "$DIR"charts/azure-service-operator/templates/generated/*_deployment_*
sed -i '1 i {{- if .Values.metrics.enable -}}' "$DIR"charts/azure-service-operator/templates/generated/*controller-manager-metrics-service*
sed -i 's/port: 8080/port: {{ .Values.metrics.port | default 8080 }}/g' "$DIR"charts/azure-service-operator/templates/generated/*controller-manager-metrics-service*
sed -i -e '$a{{- end }}' "$DIR"charts/azure-service-operator/templates/generated/*controller-manager-metrics-service*

# Azure-Service-Operator-crds actions
# We had to split charts here here as with a single chart, we were running into the max size issue with helm
# See https://github.com/helm/helm/issues/9788
find "$DIR"charts/azure-service-operator/templates/generated/*_customresourcedefinition_* -exec mv '{}' "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/templates/crds \; # move CRD definitions to crds chart folder
sed -i "1,/version:.*/s/\(version: \)\(.*\)/\1$VERSION/g" "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/Chart.yaml  # find version key and update the value with the current version for crds chart

# Helm chart packaging, indexing and updating dependencies
echo "Packaging helm charts"
helm package "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds -d "$DIR"charts # package the CRD helm files into a tar file
helm template "$DIR"charts/azure-service-operator --dependency-update > /dev/null # Update the crds dependency
helm package "$DIR"charts/azure-service-operator -d "$DIR"charts # package the ASOv2 helm files into a tar file
helm repo index "$DIR"charts # update index.yaml for Helm Repository
