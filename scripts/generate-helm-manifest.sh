#!/bin/bash

# This script generates helm manifest and replaces required values in helm chart.

set -e

KUBE_RBAC_PROXY=$1
LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=$2
PUBLIC_REGISTRY=$3
VERSION=$4
DIR=$5

GEN_FILES_DIR="$DIR"charts/azure-service-operator/templates/generated
IF_CLUSTER="{{- if or (eq .Values.multiTenant.enable false) (eq .Values.multiTenant.isCluster true) }}"
IF_TENANT="{{- if or (eq .Values.multiTenant.enable false) (eq .Values.multiTenant.isCluster false) }}"

# Matches and adds helm flow control to a file
flow_control(){
  MATCHER_IF=$1
  MATCHER_END=$2
  IF_CLAUSE=$3
  TARGET=$4

sed -i "/$MATCHER_IF/i \  \ $IF_CLAUSE" "$TARGET"
sed -i "/$MATCHER_END/a \  \ {{- end }}" "$TARGET"
}


# Manifest purge and generation
echo "Generating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$DIR"charts/azure-service-operator/values.yaml
rm -rf "$GEN_FILES_DIR" # remove generated files
rm -rf "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/templates
mkdir -p "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/templates/crds # create dirs for generated files
mkdir "$GEN_FILES_DIR"
kustomize build "$DIR"config/default -o "$GEN_FILES_DIR"
rm "$GEN_FILES_DIR"/*_namespace_* # remove namespace as we will let Helm manage it

# Sed Replacements
grep -E $KUBE_RBAC_PROXY "$GEN_FILES_DIR"/*_deployment_* > /dev/null # Ensure that what we're about to try to replace actually exists (if it doesn't we want to fail)
sed -i "s@$KUBE_RBAC_PROXY.*@{{.Values.image.kubeRBACProxy}}@g" "$GEN_FILES_DIR"/*_deployment_*
sed -i "s@$LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE@{{.Values.image.repository}}@g" "$GEN_FILES_DIR"/*_deployment_* # Replace hardcoded ASO image
sed -i '/default-logs-container: manager/a \  \ {{- if .Values.podAnnotations }}\n \  \ {{ toYaml .Values.podAnnotations | indent 6 }}\n \  \ {{- end }}' "$GEN_FILES_DIR"/*_deployment_* # Add pod annotations
sed -i "s/\(version: \)\(.*\)/\1$VERSION/g" "$DIR"charts/azure-service-operator/Chart.yaml  # find version key and update the value with the current version for both main and subchart

# Metrics Configuration
flow_control "metrics-addr" "metrics-addr" "{{- if .Values.metrics.enable}}" "$GEN_FILES_DIR"/*_deployment_*
sed -i "1,/metrics-addr=.*/s/\(metrics-addr=\)\(.*\)/\1{{ tpl .Values.metrics.address . }}/g" "$GEN_FILES_DIR"/*_deployment_*
sed -i '/metrics-addr/a \  \ {{ end }}' "$GEN_FILES_DIR"/*_deployment_* # End metrics flow control
sed -i 's/containerPort: 8080/containerPort: {{ .Values.metrics.port | default 8080 }}/g' "$GEN_FILES_DIR"/*_deployment_*
sed -i '1 i {{- if .Values.metrics.enable -}}' "$GEN_FILES_DIR"/*controller-manager-metrics-service*
sed -i 's/port: 8080/port: {{ .Values.metrics.port | default 8080 }}/g' "$GEN_FILES_DIR"/*controller-manager-metrics-service*
sed -i -e '$a{{- end }}' "$GEN_FILES_DIR"/*controller-manager-metrics-service*

# Perform file level changes for cluster and tenant
for file in $(find "$GEN_FILES_DIR" -type f)
do
    if [[ $file == *"clusterrolebinding_azureserviceoperator-manager"* ]]; then
      sed -i "1 s/^/$IF_TENANT\n/;$ a {{- end }}" "$file"
      sed -i 's/namespace: azureserviceoperator-system/namespace: {{.Release.Namespace}}/g' "$file" # TODO: flow control here
      sed -i 's/azureserviceoperator-manager-rolebinding/azureserviceoperator-manager-rolebinding-{{ .Values.tenantName }}/g' "$file" # TODO: flow control here
    elif [[ $file == *"leader-election"* ]] || [[ $file == *"deployment"* ]]; then
      flow_control "namespace: azureserviceoperator-system" "namespace: azureserviceoperator-system" "$IF_CLUSTER" "$file" # TODO: flow control on L60, 61 like 63 and 64
      sed -i "/namespace: azureserviceoperator-system/a \  \ {{ else }}\n \ namespace: {{.Release.Namespace}}" "$file"
    else
      sed -i "1 s/^/$IF_CLUSTER\n/;$ a {{- end }}" "$file"
    fi
done

flow_control "aadpodidbinding" "aadpodidbinding" "$IF_TENANT" "$GEN_FILES_DIR"/*_deployment_*

flow_control "--enable-leader-election" "--enable-leader-election" "$IF_TENANT" "$GEN_FILES_DIR"/*_deployment_*

flow_control "volumeMounts:" "secretName:" "$IF_CLUSTER" "$GEN_FILES_DIR"/*_deployment_*


# Azure-Service-Operator-crds actions
# We had to split charts here here as with a single chart, we were running into the max size issue with helm
# See https://github.com/helm/helm/issues/9788
find "$GEN_FILES_DIR"/*_customresourcedefinition_* -exec mv '{}' "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/templates/crds \; # move CRD definitions to crds chart folder
sed -i "1,/version:.*/s/\(version: \)\(.*\)/\1$VERSION/g" "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds/Chart.yaml  # find version key and update the value with the current version for crds chart

# Helm chart packaging, indexing and updating dependencies
echo "Packaging helm charts"
helm package "$DIR"charts/azure-service-operator/charts/azure-service-operator-crds -d "$DIR"charts # package the CRD helm files into a tar file
helm template "$DIR"charts/azure-service-operator --dependency-update > /dev/null # Update the crds dependency
helm package "$DIR"charts/azure-service-operator -d "$DIR"charts # package the ASOv2 helm files into a tar file
helm repo index "$DIR"charts # update index.yaml for Helm Repository
