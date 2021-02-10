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
AZURESERVICEOPERATOR_IMAGE=$(head -n 1 _Azure.azure-service-operator/drop/azure-service-operator.txt)
AZURESERVICEOPERATOR_IMAGE_PUBLIC=$(echo ${AZURESERVICEOPERATOR_IMAGE//candidate/public})
AZURESERVICEOPERATOR_IMAGE_BASE=$(echo $AZURESERVICEOPERATOR_IMAGE | cut -d':' -f 1)
AZURESERVICEOPERATOR_IMAGE_VERSION=$(echo $AZURESERVICEOPERATOR_IMAGE | cut -d':' -f 2)
AZURESERVICEOPERATOR_IMAGE_LATEST=$(echo ${AZURESERVICEOPERATOR_IMAGE_BASE//candidate/public})':latest'

echo "image: $AZURESERVICEOPERATOR_IMAGE"
echo "image_base: $AZURESERVICEOPERATOR_IMAGE_BASE"
echo "image_latest: $AZURESERVICEOPERATOR_IMAGE_LATEST"
echo "image_public: $AZURESERVICEOPERATOR_IMAGE_PUBLIC"

echo "##vso[task.setvariable variable=AZURESERVICEOPERATOR_IMAGE_LATEST]$AZURESERVICEOPERATOR_IMAGE_LATEST"
echo "##vso[task.setvariable variable=AZURESERVICEOPERATOR_IMAGE]$AZURESERVICEOPERATOR_IMAGE"
echo "##vso[task.setvariable variable=AZURESERVICEOPERATOR_IMAGE_PUBLIC]$AZURESERVICEOPERATOR_IMAGE_PUBLIC"
echo "##vso[task.setvariable variable=AZURESERVICEOPERATOR_IMAGE_BASE]$AZURESERVICEOPERATOR_IMAGE_BASE"
echo "##vso[task.setvariable variable=AZURESERVICEOPERATOR_IMAGE_VERSION]$AZURESERVICEOPERATOR_IMAGE_VERSION"

mkdir release/config -p
cp -r _Azure.azure-service-operator/drop/setup.yaml ./release/config
# Not currently including Helm charts as part of release so skip this
# cp -r _Azure.azure-service-operator/drop/*.tgz ./release/config
IMG=${AZURESERVICEOPERATOR_IMAGE_PUBLIC/"public/"/"mcr.microsoft.com/"}
echo "updating the manager image "
echo $IMG
sed -i'' -e 's@image: candidate/k8s@image: mcr.microsoft.com/k8s@' ./release/config/setup.yaml
echo  ${AZURESERVICEOPERATOR_IMAGE_PUBLIC/"public/"/"mcr.microsoft.com/"} >> ./release/notes.txt
cat ./release/config/setup.yaml
echo 'ls ./release'
ls ./release
zip -r release.zip ./release/*
ls
