#!/bin/bash

# This script generates helm manifest and replaces required values in helm chart.

set -e

KUBE_RBAC_PROXY=$1
LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE=$2
PUBLIC_REGISTRY=$3
VERSION=$4
DIR=$5

echo "Generating helm chart manifest"
sed -i "s@\($PUBLIC_REGISTRY\)\(.*\)@\1azureserviceoperator:$VERSION@g" "$DIR"charts/azure-service-operator/values.yaml
rm -rf "$DIR"charts/azure-service-operator/templates/generated
rm -rf "$DIR"charts/azure-service-operator/crds # remove generated files
mkdir "$DIR"charts/azure-service-operator/templates/generated
mkdir "$DIR"charts/azure-service-operator/crds # create dirs for generated files
kustomize build "$DIR"config/default -o "$DIR"charts/azure-service-operator/templates/generated
find "$DIR"charts/azure-service-operator/templates/generated/*_customresourcedefinition_* -exec mv '{}' "$DIR"charts/azure-service-operator/crds \; # move CRD definitions to crd folder
rm "$DIR"charts/azure-service-operator/templates/generated/*_namespace_* # remove namespace as we will let Helm manage it
sed -i "s@$LOCAL_REGISTRY_CONTROLLER_DOCKER_IMAGE@{{.Values.image.repository}}@g" "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # Replace hardcoded ASO image
sed -i '/metrics-addr/i \  \ {{if .Values.metrics.enable}}' "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # Add metrics flow control
sed -i "1,/metrics-addr=.*/s/\(metrics-addr=\)\(.*\)/\1{{.Values.metrics.address | default \"127.0.0.1:8080\" }}/g" "$DIR"charts/azure-service-operator/templates/generated/*_deployment_*
sed -i '/metrics-addr/a \  \ {{ end }}' "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # End metrics flow control
grep -E $KUBE_RBAC_PROXY "$DIR"charts/azure-service-operator/templates/generated/*_deployment_* # Ensure that what we're about to try to replace actually exists (if it doesn't we want to fail)
sed -i "s@$KUBE_RBAC_PROXY.*@{{.Values.image.kubeRBACProxy}}@g" "$DIR"charts/azure-service-operator/templates/generated/*_deployment_*
find "$DIR"charts/azure-service-operator/templates/generated/ -type f -exec sed -i "s@azureserviceoperator-system@{{.Release.Namespace}}@g" {} \;
sed -i "1,/version:.*/s/\(version: \)\(.*\)/\1$VERSION/g" "$DIR"charts/azure-service-operator/Chart.yaml  # find version key and update the value with the current version
helm package "$DIR"charts/azure-service-operator -d "$DIR"charts # package the necessary files into a tar file
helm repo index "$DIR"charts # update index.yaml for Helm Repository
