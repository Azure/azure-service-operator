#!/bin/bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail


print_usage() {
  echo "Usage: delete-kind-wi-storage.sh -d <DIRECTORY>"
}

DIR=

while getopts 'd:' flag; do
  case "${flag}" in
    d) DIR="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done


if [[ -z "$DIR" ]]; then
  print_usage
  exit 1
fi

EXISTS=$(kind get clusters -q | grep "^asov2-wi$" || true)

if [ -f "$DIR/azure/rg.txt" ]; then
  RESOURCE_GROUP=$(cat $DIR/azure/rg.txt)
else
  # Nothing to do, no existing rg
  exit 0
fi

if [[ -z "$EXISTS" ]]; then
  # Nothing to do, no match
  exit 0
fi

if [ -f "$DIR/azure/fic.txt" ]; then
  # Need to manually delete the identity FIC
  FIC=$(cat $DIR/azure/fic.txt)
  if [[ ! -z "${AZURE_IDENTITY_NAME:-}" ]]; then
    az identity federated-credential delete --resource-group "${AZURE_IDENTITY_RG}" --identity-name ${AZURE_IDENTITY_NAME} --name ${FIC} -y
  fi
fi

if [ $(az group exists --name ${RESOURCE_GROUP}) = true ]; then
  echo "Deleting resourceGroup: ${RESOURCE_GROUP}"
  az group delete --name ${RESOURCE_GROUP} -y
  echo "Done deleting resourceGroup: ${RESOURCE_GROUP}"
fi

rm -rf "${DIR}/azure"
