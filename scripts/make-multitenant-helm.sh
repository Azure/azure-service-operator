#!/bin/bash

# This script is used to make cluster and tenant level sed replacements to azure-service-operator-multitenant

set -e

CHART_DIR=$1
GEN_FILES_DIR="$CHART_DIR"/templates/generated
IF_CLUSTER="{{- if .Values.isCluster}}"
IF_TENANT="{{- if not .Values.isCluster}}"

# Matches and adds helm flow control to a file
flow_control(){
  MATCHER_IF=$1
  MATCHER_END=$2
  IF_CLAUSE=$3
  TARGET=$4

sed -i "/$MATCHER_IF/i \  \ $IF_CLAUSE" "$TARGET"
sed -i "/$MATCHER_END/a \  \ {{- end }}" "$TARGET"
}

# Perform file level changes for cluster and tenant
for file in $(find "$GEN_FILES_DIR" -type f)
do
    if [[ $file == *"clusterrolebinding_azureserviceoperator-manager"* ]]; then
      sed -i '1 s/^/{{- if not .Values.isCluster }}\n/;$ a {{- end }}' "$file"
      sed -i 's/namespace: azureserviceoperator-system/namespace: {{.Release.Namespace}}/g' "$file"
      sed -i 's/azureserviceoperator-manager-rolebinding/azureserviceoperator-manager-rolebinding-{{ .Values.tenantName }}/g' "$file"
    elif [[ $file == *"leader-election"* ]] || [[ $file == *"deployment"* ]]; then
      flow_control "namespace: azureserviceoperator-system" "namespace: azureserviceoperator-system" "$IF_CLUSTER" "$file"
      sed -i "/namespace: azureserviceoperator-system/a \  \ {{ else }}\n \ namespace: {{.Release.Namespace}}" "$file"
    else
      sed -i '1 s/^/{{- if .Values.isCluster }}\n/;$ a {{- end }}' "$file"
    fi
done

flow_control "aadpodidbinding" "aadpodidbinding" "$IF_TENANT" "$GEN_FILES_DIR"/*_deployment_*

flow_control "--enable-leader-election" "--enable-leader-election" "$IF_TENANT" "$GEN_FILES_DIR"/*_deployment_*

flow_control "volumeMounts:" "secretName:" "$IF_CLUSTER" "$GEN_FILES_DIR"/*_deployment_*
