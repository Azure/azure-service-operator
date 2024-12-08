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

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source ${SCRIPT_DIR}/workloadidentitystorage/vars.sh

# Generate names and save in files
mkdir -p "${DIR}/azure"

# Note that this resource group isn't used by this script anymore but is still used by make-mi-fic.py
# to produce a new MI + FIC.
RESOURCE_GROUP="${PREFIX}-rg-wi$(openssl rand -hex 6)"

# If somehow the files already exist then the resource also already exists and we shouldn't do anything
if [ -f "$DIR/azure/oidcid.txt" ]; then
  # Nothing to do, no existing rg
  echo "[INF] Using existing OIDC key $(cat ${DIR}/azure/oidcid.txt)"
  exit 0
fi

if [ -f "$DIR/azure/rg.txt" ]; then
   # Nothing to do, no existing rg
  echo "[INF] Using existing RG $(cat ${DIR}/azure/rg.txt)"
  exit 0
fi

# This is just an identifier to differentiate different blobs from one another so that we can support
# multiple OIDC issuers in a single storage account. It's not actually involved in the authentication flow anywhere
# it's purely ASO bookkeeping.
OIDC_IDENTIFIER="asowi$(openssl rand -hex 6)"
echo ${OIDC_IDENTIFIER} > "${DIR}/azure/oidcid.txt"

# Save the RG we're using too
echo ${RESOURCE_GROUP} > "${DIR}/azure/rg.txt"

# Generate the OIDC keys
openssl genrsa -out "$DIR/sa.key" 2048
openssl rsa -in "$DIR/sa.key" -pubout -out "$DIR/sa.pub"

if [ -z "${KIND_OIDC_STORAGE_ACCOUNT_SUBSCRIPTION-}" ]; then
  echo "[DBG] KIND_OIDC_STORAGE_ACCOUNT_SUBSCRIPTION is not set, fetching the current subscription"
  KIND_OIDC_STORAGE_ACCOUNT_SUBSCRIPTION=$(az account show --output tsv --query id)
fi

echo "[DBG] KIND_OIDC_STORAGE_ACCOUNT_SUBSCRIPTION: ${KIND_OIDC_STORAGE_ACCOUNT_SUBSCRIPTION}"

# There's already a trailing / so we don't need to add one between the web endpoint and the OIDC Identifier
echo "[DBG] Retrieving OIDC issuer URL"
ISSUER_URL="$(az storage account show --subscription ${KIND_OIDC_STORAGE_ACCOUNT_SUBSCRIPTION} --name "${KIND_OIDC_STORAGE_ACCOUNT}" -o json | jq -r .primaryEndpoints.web)${OIDC_IDENTIFIER}"
echo "${ISSUER_URL}" > "${DIR}/azure/saissuer.txt"
echo "[DBG] OIDC issuer URL: ${ISSUER_URL}"

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

CREATION_TIME="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"

echo "[INF] Creating resource group ${RESOURCE_GROUP}"
az group create -l westus -n "${RESOURCE_GROUP}" --tags "CreatedAt=${CREATION_TIME}"

echo "[INF] Uploading OIDC configuration to storage account"
echo "[DBG] KIND_OIDC_STORAGE_ACCOUNT: ${KIND_OIDC_STORAGE_ACCOUNT}"
echo "[DBG] KIND_OIDC_STORAGE_CONTAINER: ${KIND_OIDC_STORAGE_CONTAINER}"
echo "[DBG] OIDC_IDENTIFIER: ${OIDC_IDENTIFIER}"
echo "[DBG] CREATION_TIME: ${CREATION_TIME}"

az storage blob upload \
  --account-name "${KIND_OIDC_STORAGE_ACCOUNT}" \
  --container-name "${KIND_OIDC_STORAGE_CONTAINER}" \
  --file "${DIR}/openid-configuration.json" \
  --tags "CreatedAt=${CREATION_TIME}" \
  --name "${OIDC_IDENTIFIER}/.well-known/openid-configuration" \
  --auth-mode login

azwi jwks --public-keys "${DIR}/sa.pub" --output-file "${DIR}/jwks.json"

az storage blob upload \
  --account-name "${KIND_OIDC_STORAGE_ACCOUNT}" \
  --container-name "${KIND_OIDC_STORAGE_CONTAINER}" \
  --file "${DIR}/jwks.json" \
  --tags "CreatedAt=${CREATION_TIME}" \
  --name "${OIDC_IDENTIFIER}/openid/v1/jwks" \
  --auth-mode login
