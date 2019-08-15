echo 'ls'
ls _Azure.azure-service-operator/drop
ls _Azure.azure-service-operator/drop/*.yaml
azureserviceoperator_image=$(head -n 1 _Azure.azure-service-operator/drop/azure-service-operator.txt)
azureserviceoperator_image_public=$(echo ${azureserviceoperator_image//candidate/public})
azureserviceoperator_image_base=$(echo $azureserviceoperator_image | cut -d':' -f 1)
azureserviceoperator_image_latest=$(echo ${azureserviceoperator_image_base//candidate/public})':latest'
echo  $azureserviceoperator_image
echo $azureserviceoperator_image_base
echo $azureserviceoperator_image_latest
echo $azureserviceoperator_image_public
echo "##vso[task.setvariable variable=azureserviceoperator_image_latest]$azureserviceoperator_image_latest"
echo "##vso[task.setvariable variable=azureserviceoperator_image]$azureserviceoperator_image"
echo "##vso[task.setvariable variable=azureserviceoperator_image_public]$azureserviceoperator_image_public"
mkdir release/config -p
cp -r _Azure.azure-service-operator/drop/setup.yaml ./release/config

IMG=${azureserviceoperator_image_public/"public/"/"mcr.microsoft.com/"}
echo "updating the manager image "
echo $IMG
sed -i -e 's@docker.io/controllertest:1@'${IMG}'@g' ./release/config/setup.yaml
echo  ${azureserviceoperator_image_public/"public/"/"mcr.microsoft.com/"} >> ./release/notes.txt
cat ./release/config/setup.yaml
echo 'ls ./release'
ls ./release
zip -r release.zip ./release/*
ls
