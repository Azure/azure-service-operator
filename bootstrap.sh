#!/bin/bash
SP_NAME=az-service-operator
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml
kubectl -n cert-manager wait --for=condition=ready pod -l app=cert-manager
POD="$(kubectl get pods --all-namespaces|grep cert-manager-webhook|awk '{print $2}'|head --)"
while [[ $(kubectl -n cert-manager get pods $POD -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do echo "waiting for pod" && sleep 1; done

DIR="azure-service-operator"
if [ -d "$DIR" ]; then
  # Take action if $DIR exists. #
  echo "Removing directory ${DIR}... and redownloading new code"
  rm -rf ${DIR}
fi

helm repo add azureserviceoperator https://raw.githubusercontent.com/Azure/azure-service-operator/master/charts
helm repo update
IS_SP_THERE=$(az ad sp list --display-name $SP_NAME -o table )
echo "Checking for pre created service principal"
if [ -z "$IS_SP_THERE" ]
then
      echo "The service principal does not exsist"
else
      echo "Deleting service principal, we will recreate it later in the script"
      az ad sp delete --id http://$SP_NAME
fi

AZURE_CLIENT_SECRET=$(az ad sp create-for-rbac --name $SP_NAME --role Contributor -o tsv | awk '{print $4}')
AZURE_CLIENT_ID=$(az ad sp show --id http://$SP_NAME --query appId -o tsv)
AZURE_TENANT_ID=$(az ad sp show --id http://$SP_NAME --query appOwnerTenantId -o tsv)

export AZURE_CLIENT_ID
export AZURE_TENANT_ID
export AZURE_CLIENT_SECRET

helm upgrade --install az_svc_operator https://github.com/Azure/azure-service-operator/raw/master/charts/azure-service-operator-0.1.0.tgz \
        --create-namespace \
        --namespace=azureoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set azureClientSecret=$AZURE_CLIENT_SECRET \
        --set image.repository="mcr.microsoft.com/k8s/azureserviceoperator:latest"
