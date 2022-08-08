#!/bin/bash

# This script generates helm manifest and replaces required values in helm chart.

set -e

KUBE_RBAC_PROXY="gcr.io/kubebuilder/kube-rbac-proxy"
LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=$1
PUBLIC_REGISTRY=$2
VERSION=$3
CHART_DIR=$4

# Manifest purge and generation
echo "Generating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$CHART_DIR"/values.yaml
rm -rf "$CHART_DIR"/templates/generated # remove generated files
rm -rf ./v2/charts/azure-service-operator-crds/templates
mkdir -p ./v2/charts/azure-service-operator-crds/templates/crds # create dir for generated files
mkdir "$CHART_DIR"/templates/generated
kustomize build ./v2/config/default -o "$CHART_DIR"/templates/generated
rm "$CHART_DIR"/templates/generated/*_namespace_* # remove namespace as we will let Helm manage it

# Sed Replacements
grep -E $KUBE_RBAC_PROXY "$CHART_DIR"/templates/generated/*_deployment_* > /dev/null # Ensure that what we're about to try to replace actually exists (if it doesn't we want to fail)
sed -i "s@$KUBE_RBAC_PROXY.*@{{.Values.image.kubeRBACProxy}}@g" "$CHART_DIR"/templates/generated/*_deployment_*
sed -i "s@$LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE@{{.Values.image.repository}}@g" "$CHART_DIR"/templates/generated/*_deployment_* # Replace hardcoded ASO image
sed -i '/default-logs-container: manager/a \  \ {{- if .Values.podAnnotations }}\n \  \ {{ toYaml .Values.podAnnotations | indent 6 }}\n \  \ {{- end }}' "$CHART_DIR"/templates/generated/*_deployment_* # Add pod annotations
sed -i "s/\(version: \)\(.*\)/\1$VERSION/g" "$CHART_DIR"/Chart.yaml  # find version key and update the value with the current version for both main and subchart

# Metrics Configuration
sed -i '/metrics-addr/i \  \ {{if .Values.metrics.enable}}' "$CHART_DIR"/templates/generated/*_deployment_* # Add metrics flow control
sed -i "1,/metrics-addr=.*/s/\(metrics-addr=\)\(.*\)/\1{{ tpl .Values.metrics.address . }}/g" "$CHART_DIR"/templates/generated/*_deployment_*
sed -i '/metrics-addr/a \  \ {{ end }}' "$CHART_DIR"/templates/generated/*_deployment_* # End metrics flow control
sed -i 's/containerPort: 8080/containerPort: {{ .Values.metrics.port | default 8080 }}/g' "$CHART_DIR"/templates/generated/*_deployment_*
sed -i '1 i {{- if .Values.metrics.enable -}}' "$CHART_DIR"/templates/generated/*controller-manager-metrics-service*
sed -i 's/port: 8080/port: {{ .Values.metrics.port | default 8080 }}/g' "$CHART_DIR"/templates/generated/*controller-manager-metrics-service*
sed -i -e '$a{{- end }}' "$CHART_DIR"/templates/generated/*controller-manager-metrics-service*

# Azure-Service-Operator-crds actions
# We had to split charts here here as with a single chart, we were running into the max size issue with helm
# See https://github.com/helm/helm/issues/9788
find "$CHART_DIR"/templates/generated/*_customresourcedefinition_* -exec mv '{}' ./v2/charts/azure-service-operator-crds/templates/crds \; # move CRD definitions to crds chart folder
sed -i "s/\(version: \)\(.*\)/\1$VERSION/g" ./v2/charts/azure-service-operator-crds/Chart.yaml  # find version key and update the value with the current version for crds chart

echo "Successfully generated helm manifest for $CHART_DIR"