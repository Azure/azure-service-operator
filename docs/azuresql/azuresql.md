# Azure SQL Operator

## Resources Supported
The Azure SQL operator can be used to provision the following resources.

1. Azure SQL server - Deploys an Azure SQL server given the location and Resource group
2. SQL database - Deploys an SQL database given the SQL server
3. SQL firewall rule - Deploys a firewall rule to allow access to the SQL server from specific IPs
4. Action (Rolling user credentials for the SQL server) - Allows you to roll the password for the specified SQL server

## Deploying SQL Resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

**Note**  Don't forget to set the Service Principal ID, Service Principal secret, Tenant ID and Subscription ID as environment variables

### Azure SQL server

For instance, this is the sample YAML for the Azure SQL server.
Repl
  ```yaml
    apiVersion: azure.microsoft.com/v1
    kind: SqlServer
    metadata:
    name: sqlserver-sample
    spec:
     location: westus
     resourcegroup: resourceGroup1
     allowazureserviceaccess: true
  ```

The value for kind, `SqlServer` is the Custom Resource Definition (CRD) name.
`sqlserver-sample` is the name of the SQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the SQL server at and the Resource group in which you want to create it under. The `allowazureserviceaccess' boolean allows you to specify if you want Azure services to have access to your SQL server.

Once you've updated the YAML with the settings you need, and you have the operator running, you can create a Custom SQL server resource using the command.

```bash
kubectl apply -f config/samples/azure_v1_sqlserver.yaml
```

### SQL Database

Below is the sample YAML for SQL database

```yaml
apiVersion: azure.microsoft.com/v1
kind: SqlDatabase
metadata:
  name:  sqldatabase-sample
spec:
  location: westus
  resourcegroup: ResourceGroup1
  # Database Editions
  # Basic=0; Business=1; BusinessCritical=2; DataWarehouse=3; Free=4;
  # GeneralPurpose=5; Hyperscale=6; Premium=7; PremiumRS=8; Standard=9;
  # Stretch=10; System=11; System2=12; Web=13
  edition: 0
  server:  sqlserver-sample
```
Update the `location` and the `resourcegroup` to where you want to provisiong the SQL database. `server` is the name of the Azure SQL server where you want to create the database in.
The `edition` represents the SQL database edition you want to use when creating the resource and should be one of the values above.

### SQL firewall

In progress. Will be updated soon.

### SQL Action

In progress. Will be updated soon.

## View and Troubleshoot SQL Resources

To view your created SQL resources, such as sqlserver, run the following command:

```shell
kubectl get <CRD>
```

where CRD is the Custom Resource Definition name or `Kind` for the resource.

For instance, you can get the Azure SQL servers provisioned using the command

```shell
kubectl get SqlServer
```

You should see the SqlServer instances as below

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
kubectl describe SqlServer sqlserver-sample
```

```shell
Name:         sqlserver-sample234
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"azure.microsoft.com/v1","kind":"SqlServer","metadata":{"annotations":{},"name":"sqlserver-sample234","namespace":"default"}...
API Version:  azure.microsoft.com/v1
Kind:         SqlServer
Metadata:
  Creation Timestamp:  2019-09-26T21:30:56Z
  Finalizers:
    sqlserver.finalizers.azure.com
  Generation:        1
  Resource Version:  20001
  Self Link:         /apis/azure.microsoft.com/v1/namespaces/default/sqlservers/sqlserver-sample234
  UID:               ed1c5d1d-e0a4-11e9-9ee8-52a5c765e9d7
Spec:
  Allowazureserviceaccess:  true
  Location:                 westus
  Resourcegroup:            Janani-testRG
Status:
  Provisioned:  true
  State:        Ready
Events:
  Type    Reason       Age                   From                  Message
  ----    ------       ----                  ----                  -------
  Normal  Updated      2m21s                 SqlServer-controller  finalizer sqlserver.finalizers.azure.com added
  Normal  Submitting   2m21s                 SqlServer-controller  starting resource reconciliation
  Normal  Checking     108s (x3 over 2m18s)  SqlServer-controller  instance in NotReady state
  Normal  Checking     76s (x2 over 78s)     SqlServer-controller  instance in Ready state
  Normal  Provisioned  75s (x2 over 76s)     SqlServer-controller  sqlserver sqlserver-sample234 provisioned
```

The `Status` section gives you the current state of the resource that it's Ready and is provisioned.

The `Events` have a chronological record of what occurred through the process of provisioning the resource.

## Delete a SQL Resource

To delete an existing resource from Kubernetes and Azure, use the following command.

```shell
kubectl delete <Kind> <instancename>
```

For instance, deleting the above SqlServer instance would look like this.

```shell
kubectl delete SqlServer sqlserver-sample
```

The following message should appear:

`sqlserver.azure.microsoft.com sqlserver-sample deleted.`

## Demo App

Watch this demo <https://bit.ly/2lUIX6Y> to observe how you would you use the Azure SQL Operator from a real application.
