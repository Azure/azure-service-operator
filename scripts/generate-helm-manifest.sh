#!/bin/bash

# This script generates helm manifest and replaces required values in helm chart.

set -e

KUBE_RBAC_PROXY=$1
LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=$2
PUBLIC_REGISTRY=$3
VERSION=$4
DIR=$5

ASO_CHART="$DIR"charts/azure-service-operator
GEN_FILES_DIR="$ASO_CHART"/templates/generated
IF_CLUSTER="{{- if or (eq .Values.multitenant.enable false) (eq .Values.azureOperatorMode \"webhooks\") }}"
IF_TENANT="{{- if or (eq .Values.multitenant.enable false) (eq .Values.azureOperatorMode \"watchers\") }}"
IF_CRDS="{{- if eq .Values.installCRDs true }}"

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

# Manifest purge and generation
echo "Generating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$ASO_CHART"/values.yaml
rm -rf "$GEN_FILES_DIR" # remove generated files
mkdir "$GEN_FILES_DIR"

${SCRIPT_DIR}/kustomize-build.sh -v "$VERSION" -k operator -o "$GEN_FILES_DIR"

rm "$GEN_FILES_DIR"/*_namespace_* # remove namespace as we will let Helm manage it

# Chart replacements
sed -i "s/\(version: \)\(.*\)/\1$VERSION/g" "$ASO_CHART"/Chart.yaml  # find version key and update the value with the current version for both main and subchart

# Deployment replacements
grep -E $KUBE_RBAC_PROXY "$GEN_FILES_DIR"/*_deployment_* > /dev/null # Ensure that what we're about to try to replace actually exists (if it doesn't we want to fail)
sed -i "s@$KUBE_RBAC_PROXY.*@{{.Values.image.kubeRBACProxy}}@g" "$GEN_FILES_DIR"/*_deployment_*
sed -i "s@$LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE@{{.Values.image.repository}}@g" "$GEN_FILES_DIR"/*_deployment_* # Replace hardcoded ASO image
# Perl multiline replacements - using this because it's tricky to do these sorts of multiline replacements with sed
perl -0777 -i -pe 's/(template:\n.*metadata:\n.*annotations:\n(\s*))/$1\{\{- if .Values.podAnnotations \}\}\n$2\{\{ toYaml .Values.podAnnotations \}\}\n$2\{\{- end \}\}\n$2/igs' "$GEN_FILES_DIR"/*_deployment_* # Add pod annotations

# Add annotations to aso-installed-resources resource to ensure it deploys at the right time
yq -i e '.metadata.annotations.["helm.sh/hook"] = "post-install,post-upgrade"' "$GEN_FILES_DIR"/*_installedresourcedefinitions_aso-installed-resources.yaml

# Metrics Configuration
flow_control "metrics-addr" "metrics-addr" "{{- if .Values.metrics.enable}}" "$GEN_FILES_DIR"/*_deployment_*
sed -i "1,/metrics-addr=.*/s/\(metrics-addr=\)\(.*\)/\1{{ tpl .Values.metrics.address . }}/g" "$GEN_FILES_DIR"/*_deployment_*
sed -i 's/containerPort: 8080/containerPort: {{ .Values.metrics.port | default 8080 }}/g' "$GEN_FILES_DIR"/*_deployment_*
sed -i '1 i {{- if .Values.metrics.enable -}}' "$GEN_FILES_DIR"/*controller-manager-metrics-service*
sed -i 's/port: 8080/port: {{ .Values.metrics.port | default 8080 }}/g' "$GEN_FILES_DIR"/*controller-manager-metrics-service*
sed -i -e '$a{{- end }}' "$GEN_FILES_DIR"/*controller-manager-metrics-service*
find "$GEN_FILES_DIR" -type f -exec sed -i 's/azureserviceoperator-system/{{ .Release.Namespace }}/g' {} \;

# Perform file level changes for cluster and tenant
for file in $(find "$GEN_FILES_DIR" -type f)
do
  # Append cluster or tenant guards to each file
  if [[ $file == *"clusterrolebinding_azureserviceoperator-manager"* ]]; then
    sed -i "1 s/^/$IF_TENANT\n/;$ a {{- end }}" "$file"
    flow_control "name: azureserviceoperator-manager-rolebinding" "name: azureserviceoperator-manager-rolebinding" "{{- if not .Values.multitenant.enable }}" "$file"
    sed -i "/name: azureserviceoperator-manager-rolebinding/a \  \ {{ else }}\n \ name: azureserviceoperator-manager-rolebinding-{{ .Release.Namespace }}" "$file"
  elif [[ $file != *"leader-election"* ]] && [[ $file != *"_deployment_"* ]]; then
    sed -i "1 s/^/$IF_CLUSTER\n/;$ a {{- end }}" "$file"
  fi

  # Apply CRD guards
  if [[ $file == *"v1api_installedresourcedefinitions"* ]] || [[ $file == *"crd-manager-role"* ]]; then
    sed -i "1 s/^/$IF_CRDS\n/;$ a {{- end }}" "$file"
  fi
done

flow_control "aadpodidbinding" "aadpodidbinding" "$IF_TENANT" "$GEN_FILES_DIR"/*_deployment_*

flow_control "--enable-leader-election" "--enable-leader-election" "$IF_TENANT" "$GEN_FILES_DIR"/*_deployment_*

# TODO: This bit is tricky to exclude kube-rbac-proxy and webhook stuff.
flow_control "serving-certs" "name: https" "$IF_CLUSTER" "$GEN_FILES_DIR"/*_deployment_*
flow_control "- name: cert" "secretName" "$IF_CLUSTER" "$GEN_FILES_DIR"/*_deployment_*

# Helm chart packaging, indexing and updating dependencies
echo "Packaging helm charts"
helm package "$ASO_CHART" -d "$DIR"charts # package the ASOv2 helm files into a tar file
helm repo index "$DIR"charts # update index.yaml for Helm Repository
