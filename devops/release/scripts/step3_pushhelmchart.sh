# Note: This is not currently used as the Helm chart at the time of release doesn't have the correct container image version.

export HELM_EXPERIMENTAL_OCI=1
ACR_NAME='asorelease'

echo "ACR_NAME", $ACR_NAME
echo "azureserviceoperator_image: $azureserviceoperator_image"
echo "azureserviceoperator_image_base $azureserviceoperator_image_base"
echo "azureserviceoperator_image_latest: $azureserviceoperator_image_latest"
echo "azureserviceoperator_image_public: $azureserviceoperator_image_public"
echo "azureserviceoperator_image_version: $azureserviceoperator_image_version"

echo $AZURE_CLIENT_SECRET | helm registry login $ACR_NAME.azurecr.io --username $AZURE_CLIENT_ID --password-stdin

helm chart pull $ACR_NAME.azurecr.io/candidate/k8s/asohelmchart:$azureserviceoperator_image_version
helm chart export $ACR_NAME.azurecr.io/candidate/k8s/asohelmchart:$azureserviceoperator_image_version --destination .
cd azure-service-operator
helm chart save . $ACR_NAME.azurecr.io/public/k8s/asohelmchart:$azureserviceoperator_image_version
helm chart save . $ACR_NAME.azurecr.io/public/k8s/asohelmchart:latest
helm chart push $ACR_NAME.azurecr.io/public/k8s/asohelmchart:$azureserviceoperator_image_version
helm chart push $ACR_NAME.azurecr.io/public/k8s/asohelmchart:latest
