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

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source ${SCRIPT_DIR}/workloadidentitystorage/vars.sh

if [ -f "$DIR/azure/fic.txt" ]; then
  # Need to manually delete the identity FIC
  FIC=$(cat $DIR/azure/fic.txt)
  if [[ ! -z "${AZURE_IDENTITY_NAME:-}" ]]; then
    az identity federated-credential delete --resource-group "${AZURE_IDENTITY_RG}" --identity-name ${AZURE_IDENTITY_NAME} --name ${FIC} -y
  fi
fi

if [ -f "$DIR/azure/roleassignmentid.txt" ]; then
  # Need to delete the role assignment as well so we don't leak them
  ROLE_ASSIGNMENT_ID=$(cat $DIR/azure/roleassignmentid.txt)
  echo "Deleting role assignment: ${ROLE_ASSIGNMENT_ID}"
  az role assignment delete --ids "${ROLE_ASSIGNMENT_ID}"
fi

if [ -f "$DIR/azure/oidcid.txt" ]; then
  OIDC_IDENTIFIER=$(cat $DIR/azure/oidcid.txt)
  echo "Deleting blobs associated with OIDC ID: ${OIDC_IDENTIFIER}"
  az storage blob delete-batch \
    --account-name "${KIND_OIDC_STORAGE_ACCOUNT}" \
    --source "${KIND_OIDC_STORAGE_CONTAINER}" \
    --pattern "${OIDC_IDENTIFIER}/*" \
    --auth-mode login
  echo "Done deleting blobs associated with OIDC ID: ${OIDC_IDENTIFIER}"
fi

if [ -f "$DIR/azure/rg.txt" ]; then
  RESOURCE_GROUP=$(cat $DIR/azure/rg.txt)

  if [ $(az group exists --name ${RESOURCE_GROUP}) = true ]; then
    echo "Deleting resourceGroup: ${RESOURCE_GROUP}"
    az group delete --name ${RESOURCE_GROUP} -y
    echo "Done deleting resourceGroup: ${RESOURCE_GROUP}"
  fi
fi

echo "Deleting directory ${DIR}"
rm -rf "${DIR}"
