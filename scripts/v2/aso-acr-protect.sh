#!/bin/bash

# -e immediate exit on error
# -u treat unset variables as an error
set -eu

IMAGES=$(az acr repository show-tags --name asorelease --repository public/k8s/azureserviceoperator --output tsv)

for i in ${IMAGES}
do
    az acr repository update --name asorelease --image public/k8s/azureserviceoperator:${i} --write-enabled false --delete-enabled false
done
