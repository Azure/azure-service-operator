# Note: This is not currently used as the Helm chart at the time of release doesn't have the correct container image version.

export HELM_EXPERIMENTAL_OCI=1
ACR_NAME='asorelease'

echo "ACR_NAME", $ACR_NAME
echo "AZURESERVICEOPERATOR_IMAGE: $AZURESERVICEOPERATOR_IMAGE"
echo "AZURESERVICEOPERATOR_IMAGE_BASE $AZURESERVICEOPERATOR_IMAGE_BASE"
echo "AZURESERVICEOPERATOR_IMAGE_LATEST: $AZURESERVICEOPERATOR_IMAGE_LATEST"
echo "AZURESERVICEOPERATOR_IMAGE_PUBLIC: $AZURESERVICEOPERATOR_IMAGE_PUBLIC"
echo "AZURESERVICEOPERATOR_IMAGE_VERSION: $AZURESERVICEOPERATOR_IMAGE_VERSION"

echo $AZURE_CLIENT_SECRET | helm registry login $ACR_NAME.azurecr.io --username $AZURE_CLIENT_ID --password-stdin

helm chart pull $ACR_NAME.azurecr.io/candidate/k8s/asohelmchart:$AZURESERVICEOPERATOR_IMAGE_VERSION
helm chart export $ACR_NAME.azurecr.io/candidate/k8s/asohelmchart:$AZURESERVICEOPERATOR_IMAGE_VERSION --destination .
cd azure-service-operator
helm chart save . $ACR_NAME.azurecr.io/public/k8s/asohelmchart:$AZURESERVICEOPERATOR_IMAGE_VERSION
helm chart save . $ACR_NAME.azurecr.io/public/k8s/asohelmchart:latest
helm chart push $ACR_NAME.azurecr.io/public/k8s/asohelmchart:$AZURESERVICEOPERATOR_IMAGE_VERSION
helm chart push $ACR_NAME.azurecr.io/public/k8s/asohelmchart:latest
