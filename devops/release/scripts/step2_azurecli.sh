ACR_NAME='azureserviceoperator'

echo "ACR_NAME", $ACR_NAME
echo  "azureserviceoperator_image", $(azureserviceoperator_image)
echo "azureserviceoperator_image_base", $(azureserviceoperator_image_base)
echo "azureserviceoperator_image_latest", $(azureserviceoperator_image_latest)
echo "azureserviceoperator_image_public", $(azureserviceoperator_image_public)
echo "azureserviceoperator_image_version", $(azureserviceoperator_image_version)


az acr login --name $ACR_NAME
docker pull $ACR_NAME.azurecr.io/$(azureserviceoperator_image)
docker tag $ACR_NAME.azurecr.io/$(azureserviceoperator_image)  $ACR_NAME.azurecr.io/$(azureserviceoperator_image_latest)
docker tag $ACR_NAME.azurecr.io/$(azureserviceoperator_image)  $ACR_NAME.azurecr.io/$(azureserviceoperator_image_public)
docker image ls
docker push $ACR_NAME.azurecr.io/$(azureserviceoperator_image_public)  
docker push $ACR_NAME.azurecr.io/$(azureserviceoperator_image_latest)
