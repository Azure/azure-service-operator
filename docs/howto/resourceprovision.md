# Deploy, View and Delete Resources

## Deploy Custom Resource

Once the Azure Service Operator is deployed, you can use the YAML files in the `config/samples` folder to create the resources using the following command.

```
kubectl apply -f <YAML file>
```

## Tags in Azure

When deploying resources using the YAML files, you can specify as labels any tags you want to add to the resource in Azure.
If the labels contain characters that are not allowed as tags in Azure (<,>,\,/,%,?), those characters will be replaced by a period(.) and added as Azure tags.

## View and Troubleshoot Custom Resource

To view your created custom resource, run the following command:

```
kubectl get <CRD>
```

where CRD is the Custom Resource Definition name or `Kind` for the resource.

For instance, you can get the Azure SQL servers provisioned using the command

```
kubectl get AzureSqlServer
```

You should see the AzureSqlServer instances as below

```
NAME                  AGE
sqlserver-sample      1h
```

If you want to see more details about a particular resource instance such as the `Status` or `Events`, you can use the below command

```
kubectl describe <Kind> <instance name>
```

For instance, the below command is used to get more details about the `sqlserver-sample` instance

```
kubectl describe AzureSqlServer sqlserver-sample
```

```
Name:         sqlserver-sample234
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"azure.microsoft.com/v1alpha1","kind":"AzureSqlServer","metadata":{"annotations":{},"name":"sqlserver-sample234","namespace":"default"}...
API Version:  azure.microsoft.com/v1alpha1
Kind:         SqlServer
Metadata:
  Creation Timestamp:  2019-09-26T21:30:56Z
  Finalizers:
    azuresqlserver.finalizers.azure.com
  Generation:        1
  Resource Version:  20001
  Self Link:         /apis/azure.microsoft.com/v1/namespaces/default/azuresqlservers/sqlserver-sample234
  UID:               ed1c5d1d-e0a4-11e9-9ee8-52a5c765e9d7
Spec:
  Location:                 westus
  Resourcegroup:            resourceGroup1
Status:
  Message: Success
  Provisioned:  true
  State:        Ready
Events:
  Type    Reason       Age                   From                  Message
  ----    ------       ----                  ----                  -------
  Normal  Updated      2m21s                 SqlServer-controller  finalizer azuresqlserver.finalizers.azure.com added
  Normal  Submitting   2m21s                 SqlServer-controller  starting resource reconciliation
  Normal  Checking     108s (x3 over 2m18s)  SqlServer-controller  instance in NotReady state
  Normal  Checking     76s (x2 over 78s)     SqlServer-controller  instance in Ready state
  Normal  Provisioned  75s (x2 over 76s)     SqlServer-controller  azuresqlserver sqlserver-sample234 provisioned
```

The `Status` section gives you the current state of the resource, it's `State` and if it is `Provisioned`. It also provides a more detailed `Message`

The `Events` have a chronological record of what occurred through the process of provisioning the resource.

## Delete Resource

To delete an existing resource from Kubernetes and Azure, use the following command.

```
kubectl delete <Kind> <instancename>
```

For instance, deleting a AzureSqlServer instance would look like this.

```
kubectl delete azuresqlserver sqlserver-sample
```

The following message should appear:

`azuresqlserver.azure.microsoft.com sqlserver-sample deleted.`

## Delete Kubernetes instances without deleting Azure resources

In some cases, like when you setup a new Kubernetes cluster with the same CRDs and want to take down the older cluster, you might need to delete the Kubernetes instances without impacting the Azure resources (as these are still tracked by CRDs in a different cluster)

In this case, you can use the annotation `skipreconcile` set to `true'. Follow the steps below to accomplish this:

1. Patch the CRD with the annotation above. You can use `kubectl apply -f <config.yaml>`. A sample YAML is below.

```
apiVersion: azure.microsoft.com/v1alpha1
kind: KeyVault
metadata:
  name: keyvaultsample123
  annotations:
    skipreconcile: "true"
spec:
  resourceGroup: resourcegroup-azure-operators
  location: westus
  enableSoftDelete: true
```

2. Delete the CRD now using `kubectl delete`. Since the object is updated with the annotation, the kubernetes instance is deleted without impacting the Azure resource.