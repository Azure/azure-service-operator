#!/bin/bash

set -e

print_usage() {
  echo "Usage: kustomize-build.sh -k crd|operator -v <version> [-o <DESTINATION>]"
  echo "    -o: Output to a specific folder. If this is not set results are streamed to stdout, otherwise results are saved in a file per resource."
  echo "    -k (required): Specifies 'crd' or 'operator', indicating if you want the operator YAML or the CRD yaml"
  echo "    -v (required): The operator version"
}

while getopts 'o:k:v:' flag; do
  case "${flag}" in
    o) DESTINATION="${OPTARG}" ;;
    k) KIND="${OPTARG}" ;;
    v) VERSION="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done

if [[ -z "$KIND" ]] || [[ -z "$VERSION" ]]; then
  print_usage
  exit 1
fi

if [[ "$KIND" != "crd" ]] && [[ "$KIND" != "operator" ]]; then
  print_usage
  exit 1
fi

ENVSUBST_VARS='${VERSION}'
export VERSION=${VERSION}

if [[ -n "$DESTINATION" ]]; then
  kustomize build config/default -o "${DESTINATION}"

  if [[ "$KIND" == "operator" ]]; then
    find "${DESTINATION}"/*_customresourcedefinition_* -delete
  else
    find "${DESTINATION}"/* -not -name "*_customresourcedefinition_*" -delete
  fi

  # envsubst on all the files
  for file in $(find "$DESTINATION" -type f); do
    tmp=$(mktemp)
    cp --attributes-only --preserve $file $tmp
    cat $file | envsubst $ENVSUBST_VARS > $tmp
    mv $tmp $file
  done

elif [[ "$KIND" == "crd" ]]; then
  kustomize build config/default | yq e '. | select(.kind == "CustomResourceDefinition")' - | envsubst $ENVSUBST_VARS
elif [[ "$KIND" == "operator" ]]; then
  kustomize build config/default | yq e '. | select(.kind != "CustomResourceDefinition")' - | envsubst $ENVSUBST_VARS
fi
