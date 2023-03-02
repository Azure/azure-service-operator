#!/bin/bash

set -e

print_usage() {
  echo "Usage: kustomize-build.sh -k crd|operator [-o <DESTINATION>]"
  echo "    -o: Output to a specific folder. If this is not set results are streamed to stdout, otherwise results are saved in a file per resource."
  echo "    -k: Specifies 'crd' or 'operator', indicating if you want the operator YAML or the CRD yaml"
  echo "    -k must be set, -o is optional"
}

while getopts 'o:k:' flag; do
  case "${flag}" in
    o) DESTINATION="${OPTARG}" ;;
    k) KIND="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done

if [[ -z "$KIND" ]]; then
  print_usage
  exit 1
fi

if [[ "$KIND" != "crd" ]] && [[ "$KIND" != "operator" ]]; then
  print_usage
  exit 1
fi

if [[ -n "$DESTINATION" ]]; then
  kustomize build config/default -o "${DESTINATION}"

  if [[ "$KIND" == "operator" ]]; then
    find "${DESTINATION}"/*_customresourcedefinition_* -not -name "*installedresourcedefinitions.serviceoperator.azure.com*" -delete
  else
    find "${DESTINATION}"/* -not -name "*_customresourcedefinition_*" -delete
    find "${DESTINATION}"/* -name "*installedresourcedefinitions.serviceoperator.azure.com*" -delete
  fi

#  rm -rf "${DESTINATION}/operator"
#  rm -rf "${DESTINATION}/crds"
#  mkdir -p "${DESTINATION}/operator"
#  mkdir -p "${DESTINATION}/crds"
#
#  kustomize build config/default -o "${DESTINATION}/operator"
#
#  find "${DESTINATION}/operator"/*_customresourcedefinition_* -not -name "*installedresourcedefinitions.serviceoperator.azure.com*" -exec mv '{}' "${DESTINATION}/crds/" \;
elif [[ "$KIND" == "crd" ]]; then
  kustomize build config/default | yq e '. | select(.kind == "CustomResourceDefinition" and .metadata.name != "installedresourcedefinitions.serviceoperator.azure.com")' -
elif [[ "$KIND" == "operator" ]]; then
  kustomize build config/default | yq e '. | select(.kind != "CustomResourceDefinition" or .metadata.name == "installedresourcedefinitions.serviceoperator.azure.com")' -
fi
