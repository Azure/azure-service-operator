ACR_NAME='azureserviceoperator'
echo "ACR_NAME"
echo $ACR_NAME
echo  $azureserviceoperator_image
echo $azureserviceoperator_image_base
echo $azureserviceoperator_image_latest
echo $azureserviceoperator_image_public
az acr login --name $ACR_NAME
docker pull $ACR_NAME.azurecr.io/$(azureserviceoperator_image)
docker tag $ACR_NAME.azurecr.io/$(azureserviceoperator_image)  $ACR_NAME.azurecr.io/$(azureserviceoperator_image_latest)
docker tag $ACR_NAME.azurecr.io/$(azureserviceoperator_image)  $ACR_NAME.azurecr.io/$(azureserviceoperator_image_public)
docker image ls
docker push $ACR_NAME.azurecr.io/$(azureserviceoperator_image_public)  
docker push $ACR_NAME.azurecr.io/$(azureserviceoperator_image_latest)
