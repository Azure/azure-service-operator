#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

print_usage() {
  echo "Usage: deploy.sh [-r]"
  echo "    -r: If the ASO CI subscription should have role assigned to it. Only should be used in ASO CI contexts."
  echo "Note that this deploys to the currently selected az cli subscription"
}

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source ${SCRIPT_DIR}/vars.sh

ASSIGN_CI_ROLE=0

while getopts 'r' flag; do
  case "${flag}" in
    r) ASSIGN_CI_ROLE=1 ;;
    *) print_usage
       exit 1 ;;
  esac
done

if [ -z "${KIND_OIDC_STORAGE_ACCOUNT_RG}" -o -z "${KIND_OIDC_STORAGE_ACCOUNT}" ]; then
  echo "Missing required environment variables KIND_OIDC_STORAGE_ACCOUNT_RG and KIND_OIDC_STORAGE_ACCOUNT"
fi

az group create --name "${KIND_OIDC_STORAGE_ACCOUNT_RG}" --location westus

az storage account create --resource-group "${KIND_OIDC_STORAGE_ACCOUNT_RG}" --name "${KIND_OIDC_STORAGE_ACCOUNT}" \
  --allow-blob-public-access false \
  --allow-shared-key-access false

# Enable static website serving
az storage blob service-properties update --account-name "${KIND_OIDC_STORAGE_ACCOUNT}" --static-website --auth-mode login
az storage container create --account-name "${KIND_OIDC_STORAGE_ACCOUNT}" --name "${KIND_OIDC_STORAGE_CONTAINER}" --auth-mode login

# Individual identity assignments not managed here but instead managed through the portal

if [[ "$ASSIGN_CI_ROLE" -eq 1 ]]; then
  # Assign permissions to the CI identity
  AZURE_SUBSCRIPTION_ID=$(az account show --output tsv --query id)

  IDENTITY_OBJECT_ID=$(az identity show -g 1es-pool -n aso-ci-identity --query principalId -otsv)
  az role assignment create \
    --assignee-object-id ${IDENTITY_OBJECT_ID} \
    --assignee-principal-type "ServicePrincipal" \
    --role "Storage Blob Data Contributor" \
    --scope /subscriptions/${AZURE_SUBSCRIPTION_ID}/resourceGroups/${KIND_OIDC_STORAGE_ACCOUNT_RG}/providers/Microsoft.Storage/storageAccounts/${KIND_OIDC_STORAGE_ACCOUNT}
fi

echo "Set the following environment variables to use this storage account"
echo "export KIND_OIDC_STORAGE_ACCOUNT_RG=${KIND_OIDC_STORAGE_ACCOUNT_RG}"
echo "export KIND_OIDC_STORAGE_ACCOUNT=${KIND_OIDC_STORAGE_ACCOUNT}"
