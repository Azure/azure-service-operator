#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

function create_role_assignment() {
  az role assignment create --assignee "${USER_ASSIGNED_OBJECT_ID}" \
      --role "Owner" \
      --subscription "${AZURE_SUBSCRIPTION_ID}"
}

function retry_create_role_assignment() {
  until create_role_assignment; do
    sleep 5
  done

}

print_usage() {
  echo "Usage: make-mi-fic.sh -g <RESOURCE_GROUP> -i <ISSUER> -d <DIR>"
  echo "    To bring your own identity, set the AZURE_IDENTITY_NAME and AZURE_IDENTITY_RG env variables"
}

RESOURCE_GROUP=
ISSUER=
DIR=

while getopts 'g:i:d:' flag; do
  case "${flag}" in
    g) RESOURCE_GROUP="${OPTARG}" ;;
    i) ISSUER="${OPTARG}" ;;
    d) DIR="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done


if [[ -z "$RESOURCE_GROUP" ]] || [[ -z "$ISSUER" ]] || [[ -z "$DIR" ]]; then
  print_usage
  exit 1
fi

IDENTITIES=$(az identity list --resource-group ${RESOURCE_GROUP} -o table)

if [[ ! -z "$IDENTITIES" ]]; then
  echo "An identity already exists, not creating another one"
  exit 0
fi

EXISTING_IDENTITY=false
if [[ ! -z "${AZURE_IDENTITY_NAME:-}" ]]; then
  IDENTITY_NAME="$AZURE_IDENTITY_NAME"
  RESOURCE_GROUP="$AZURE_IDENTITY_RG"
  EXISTING_IDENTITY=true
else
  IDENTITY_NAME="mi$(openssl rand -hex 6)"
fi

SUBJECT="system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default"

if [ "$EXISTING_IDENTITY" = false ]; then
  az identity create --name ${IDENTITY_NAME} --resource-group ${RESOURCE_GROUP}
fi

az identity federated-credential create \
  --identity-name ${IDENTITY_NAME} \
  --name fic \
  --resource-group ${RESOURCE_GROUP} \
  --issuer ${ISSUER} \
  --subject ${SUBJECT} \
  --audiences "api://AzureADTokenExchange"

export USER_ASSIGNED_CLIENT_ID="$(az identity show --resource-group "${RESOURCE_GROUP}" --name "${IDENTITY_NAME}" --query 'clientId' -otsv)"
export USER_ASSIGNED_OBJECT_ID="$(az identity show --resource-group "${RESOURCE_GROUP}" --name "${IDENTITY_NAME}" --query 'principalId' -otsv)"

# Assumption is if the user brought their own identity that it already has the permisisons it needs
if [ "$EXISTING_IDENTITY" = false ]; then
  export -f create_role_assignment
  export -f retry_create_role_assignment
  timeout 1m bash -c retry_create_role_assignment
fi

echo ${USER_ASSIGNED_CLIENT_ID} > "${DIR}/azure/miclientid.txt"
if [ "$EXISTING_IDENTITY" = true ]; then
  echo "fic" > "${DIR}/azure/fic.txt"
fi
