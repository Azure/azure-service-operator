#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail


print_usage() {
  echo "Usage: make-aks.sh -l <LOCATION> [-g <RESOURCE_GROUP> -a <ACR_NAME> -c <CLUSTER_NAME>]"
}

# TODO maybe just make all of these arguments required?
RESOURCE_GROUP=
ACR_NAME=
CLUSTER_NAME=
LOCATION=
while getopts 'g:l:a:c:' flag; do
  case "${flag}" in
    g) RESOURCE_GROUP="${OPTARG}" ;;
    l) LOCATION="${OPTARG}" ;;
    a) ACR_NAME="${OPTARG}" ;;
    c) CLUSTER_NAME="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done

# Deal with required parameters
if [[ -z "$LOCATION" ]]; then
  print_usage
  exit 1
fi

if [[ -z "${RESOURCE_GROUP}" ]]; then
  RESOURCE_GROUP="$(hostname)-aso-rg"
fi

if [[ -z "${ACR_NAME}" ]]; then
  ACR_NAME="$(hostname)-aso-acr"
fi

if [[ -z "${ACR_NAME}" ]]; then
  CLUSTER_NAME="$(hostname)-aso-aks"
fi

echo "Creating RG: ${RESOURCE_GROUP}..."
az group create --name ${RESOURCE_GROUP} --location ${LOCATION} -o none

echo "Creating ACR: ${ACR_NAME}..."
az acr create --resource-group ${RESOURCE_GROUP} --name ${ACR_NAME} --sku Basic --admin-enabled -o none

echo "Creating AKS cluster: ${CLUSTER_NAME}..."
az aks create --resource-group ${RESOURCE_GROUP} --name ${CLUSTER_NAME} --attach-acr ${ACR_NAME} \
  --enable-managed-identity --node-count 3 --generate-ssh-keys --network-dataplane cilium \
  --network-plugin azure --network-plugin-mode overlay --tier standard -o none
az aks get-credentials --resource-group ${RESOURCE_GROUP} --name ${CLUSTER_NAME} --overwrite-existing

