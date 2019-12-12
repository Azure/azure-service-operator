echo 'ls -d **/*'
ls -d **/*
if [[ ! -d _Azure.azure-service-operator/drop ]] ; then
    echo 'ERROR: Folder "_Azure.azure-service-operator/drop" is not there, aborting.'
    exit -1
fi
if [[ ! -f _Azure.azure-service-operator/drop/azure-service-operator.txt ]] ; then
    echo 'ERROR: File "_Azure.azure-service-operator/drop/azure-service-operator.txt" is not there, aborting.'
    exit -1
fi
azureserviceoperator_image=$(head -n 1 _Azure.azure-service-operator/drop/azure-service-operator.txt)
azureserviceoperator_image_public=$(echo ${azureserviceoperator_image//candidate/public})
azureserviceoperator_image_base=$(echo $azureserviceoperator_image | cut -d':' -f 1)
azureserviceoperator_image_version=$(echo $azureserviceoperator_image | cut -d':' -f 2)
azureserviceoperator_image_latest=$(echo ${azureserviceoperator_image_base//candidate/public})':latest'

echo  $azureserviceoperator_image
echo $azureserviceoperator_image_base
echo $azureserviceoperator_image_latest
echo $azureserviceoperator_image_public

echo "##vso[task.setvariable variable=azureserviceoperator_image_latest]$azureserviceoperator_image_latest"
echo "##vso[task.setvariable variable=azureserviceoperator_image]$azureserviceoperator_image"
echo "##vso[task.setvariable variable=azureserviceoperator_image_public]$azureserviceoperator_image_public"
echo "##vso[task.setvariable variable=azureserviceoperator_image_base]$azureserviceoperator_image_base"
echo "##vso[task.setvariable variable=azureserviceoperator_image_version]$azureserviceoperator_image_version"

mkdir release/config -p
cp -r _Azure.azure-service-operator/drop/setup.yaml ./release/config
IMG=${azureserviceoperator_image_public/"public/"/"mcr.microsoft.com/"}
echo "updating the manager image "
echo $IMG
sed -i'' -e 's@image: docker.io/controllertest:.*@image: '${IMG}'@' ./release/config/setup.yaml
sed -i'' -e 's@image: IMAGE_URL@image: '${IMG}'@' ./release/config/setup.yaml
echo  ${azureserviceoperator_image_public/"public/"/"mcr.microsoft.com/"} >> ./release/notes.txt
cat ./release/config/setup.yaml
echo 'ls ./release'
ls ./release
zip -r release.zip ./release/*
ls
