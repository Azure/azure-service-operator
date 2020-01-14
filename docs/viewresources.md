# View and Troubleshoot Custom Resources

To view your created custom resource, run the following command:

```shell
kubectl get <CRD>
```

where CRD is the Custom Resource Definition name or `Kind` for the resource.

For instance, you can get the Azure SQL servers provisioned using the command

```shell
kubectl get AzureSqlServer
```

You should see the AzureSqlServer instances as below

```shell
NAME                  AGE
sqlserver-sample      1h
```

If you want to see more details about a particular resource instance such as the `Status` or `Events`, you can use the below command

```shell
kubectl describe <Kind> <instance name>
```

For instance, the below command is used to get more details about the `sqlserver-sample` instance

```shell
kubectl describe AzureSqlServer sqlserver-sample
```

```shell
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
