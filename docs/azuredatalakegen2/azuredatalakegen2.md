# Azure Data Lake Gen2 Operator 
## Resources Supported
The Azure Data Lake Gen2 operator can be used to provision the following resources
1. Azure Data Lake Gen2 enabled storage account - Deploys a data lake gen2 enabled storage account using the storage operator
2. Azure Data Lake Gen2 Filesystem - Deploys a filesystem inside of a data lake gen2 enabled storage account

## Deploying Data Lake Resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

**Note**  Don't forget to set the Service Principal ID, Service Principal secret, Tenant ID and Subscription ID as environment variables

### Azure Data Lake Gen 2 Enabled Storage Account
This is the sample YAML for the Data Lake Gen2 enabled storage account

```yaml
    apiVersion: azure.microsoft.com/v1alpha1
    kind: Storage
    metadata:
        name: adlsaccountsample
    spec:
        location: westus
        resourceGroup: resourcegroup-azure-operators
        sku:
            name: Standard_LRS
        kind: StorageV2
        accessTier: Hot
        supportsHttpsTrafficOnly: true
        dataLakeEnabled: true
```
Important Values:
- Kind: the Custom Resource Definition (CRD) name
- Metadata.Name: the name of the data lake enabled storage account that will be created
- Spec.Location: Azure region where you want to create the data lake enabled storage account
- Spec.ResourceGroup: Name of the resource group under which you want to create the data lake enabled storage account
- Spec.Kind: The kind of storage account that you want to create. This value must be `StorageV2` in order for it to be a data lake.
- Spec.DataLakeWEnabled: If set to `true`, the storage account will have hierarchical namespace enabled and will, therefore, be a data lake enabled storage account

To create an instance of the Data Lake Enabled Storage Account:

```shell
kubectl apply -f config/samples/azure_v1_azuredatalakegen2storage.yaml
```
### Azure Data Lake Gen 2 FileSystem
This is the sample YAML for the Data Lake Gen2 FileSystem

```yaml
    apiVersion: azure.microsoft.com/v1alpha1
    kind: AzureDataLakeGen2FileSystem
    metadata:
        name: adls-filesystem-sample
    spec:
        storageAccountName: adlsaccountsample
        resourceGroup: resourcegroup-azure-operators
```
Important Values:
- Kind: the Custom Resource Definition (CRD) name
- Metadata.Name: the name of the data lake enabled storage account that will be created
- Spec.Location: Name of the data lake enabled storage account in which you would like to create the filesystem
- Spec.ResourceGroup: Name of the resource group under which the storage account lives

To create an instance of the Azure Data Lake Gen2 FileSystem:

```shell
kubectl apply -f config/samples/azure_v1_azuredatalakegen2filesystem.yaml
```

## View and Troubleshoot ADLS Gen2 Resources

To view your created data lake resources run the following command:

```shell
kubectl get <CRD>
```

For instance, you can get the Azure Data Lake Gen2 FileSystem provisioned using the command

```shell
kubectl get AzureDataLakeGen2FileSystem
```
You should see the AzureDataLakeGen2FileSystem instances as below 

```shell
NAME                     AGE
adls-filesystem-sample   5s
```

If you want to see more details about a particular resource instance such as the `Status` or `Events`, you can use the below command

```shell
kubectl describe <Kind> <instance name>
```

For instance, the below command is used to get more details about the `adls-filesystem-sample` instance

```shell
kubectl describe AzureDataLakeGen2FileSystem adls-filesystem-sample
```

The expected output should look like this:

```shell
Name:         adls-filesystem-sample
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"azure.microsoft.com/v1alpha1","kind":"AzureDataLakeGen2FileSystem","metadata":{"annotations":{},"name":"adls-filesystem",
                "namespace":"default"}...
API Version:  azure.microsoft.com/v1alpha1
Kind:         AzureDataLakeGen2FileSystem
Metadata:
  Creation Timestamp:  2019-10-24T23:59:22Z
  Finalizers:
    filesystem.finalizers.azure.com
  Generation:        2
  Resource Version:  1708
  Self Link:         /apis/azure.microsoft.com/v1alpha1/namespaces/default/azuredatalakegen2filesystems/adls-filesystem
  UID:               3d75a97b-1969-4d00-bf75-85516c27f43c
Output:
Spec:
  Resource Group:        resourcegroup-azure-operators
  Storage Account Name:  adlsaccountsample
Events:
  Type     Reason   Age   From                                    Message
  ----     ------   ----  ----                                    -------
  Normal   Updated  41s   AzureDataLakeGen2FileSystem-controller  finalizer filesystem.finalizers.azure.com added
  Warning  Failed   40s   AzureDataLakeGen2FileSystem-controller  Couldn't create resource in azure
  Normal   Updated  8s    AzureDataLakeGen2FileSystem-controller  adls-filesystem-sample provisioned
```
The `Events` have a chronological record of what occurred through the process of provisioning the resource.

## Delete an ADLSGen2 Resource

To delete an existing resource from Kubernetes and Azure, use the following command.

```shell
kubectl delete <Kind> <instancename>
```

For instance, deleting the above FileSystem instance would look like this.

```shell
kubectl delete AzureDataLakeGen2FileSystem adls-filesystem-sample
```

The output should look like this:

```shell
azuredatalakegen2filesystem.azure.microsoft.com "adls-filesystem-sample" deleted
```