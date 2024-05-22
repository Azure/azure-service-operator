#!/bin/bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail


print_usage() {
  echo "Usage: create-kind-wi-storage.sh -d <DIRECTORY> -p <PREFIX>"
}

DIR=
PREFIX=

while getopts 'd:p:' flag; do
  case "${flag}" in
    d) DIR="${OPTARG}" ;;
    p) PREFIX="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done

if [[ -z "$DIR" ]] || [[ -z "$PREFIX" ]]; then
  print_usage
  exit 1
fi

# Generate names and save in files
mkdir -p "${DIR}/azure"
RESOURCE_GROUP="${PREFIX}-rg-wi$(openssl rand -hex 6)"
AZURE_STORAGE_ACCOUNT="asowi$(openssl rand -hex 6)"
# This $web container is a special container that serves static web content without requiring public access enablement.
# See https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blob-static-website
AZURE_STORAGE_CONTAINER="\$web"

# If somehow the files already exist then the resource also already exists and we shouldn't do anything
if [ -f "$DIR/azure/rg.txt" ]; then
  # Nothing to do, no existing rg
  echo "Using existing RG $(cat ${DIR}/azure/rg.txt)"
  exit 0
fi

echo ${RESOURCE_GROUP} > "${DIR}/azure/rg.txt"

# Generate the OIDC keys
openssl genrsa -out "$DIR/sa.key" 2048
openssl rsa -in "$DIR/sa.key" -pubout -out "$DIR/sa.pub"

az group create -l westus -n "${RESOURCE_GROUP}" --tags "CreatedAt=$(date --utc +"%Y-%m-%dT%H:%M:%SZ")"

az storage account create --resource-group "${RESOURCE_GROUP}" --name "${AZURE_STORAGE_ACCOUNT}"
# Enable static website serving
az storage blob service-properties update --account-name "${AZURE_STORAGE_ACCOUNT}" --static-website
az storage container create --account-name "${AZURE_STORAGE_ACCOUNT}" --name "${AZURE_STORAGE_CONTAINER}"

ISSUER_URL=$(az storage account show --name "${AZURE_STORAGE_ACCOUNT}" -o json | jq -r .primaryEndpoints.web)
echo "${ISSUER_URL}" > "${DIR}/azure/saissuer.txt"

cat <<EOF > "${DIR}/openid-configuration.json"
{
  "issuer": "${ISSUER_URL}",
  "jwks_uri": "${ISSUER_URL}/openid/v1/jwks",
  "response_types_supported": [
    "id_token"
  ],
  "subject_types_supported": [
    "public"
  ],
  "id_token_signing_alg_values_supported": [
    "RS256"
  ]
}
EOF

az storage blob upload \
  --account-name "${AZURE_STORAGE_ACCOUNT}" \
  --container-name "${AZURE_STORAGE_CONTAINER}" \
  --file "${DIR}/openid-configuration.json" \
  --name .well-known/openid-configuration

azwi jwks --public-keys "${DIR}/sa.pub" --output-file "${DIR}/jwks.json"

az storage blob upload \
  --account-name "${AZURE_STORAGE_ACCOUNT}" \
  --container-name "${AZURE_STORAGE_CONTAINER}" \
  --file "${DIR}/jwks.json" \
  --name openid/v1/jwks
