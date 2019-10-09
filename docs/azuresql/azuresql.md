# Azure SQL Operator

## Resources Supported

The Azure SQL operator can be used to provision the following resources.

1. Azure SQL server - Deploys an Azure SQL server given the location and Resource group
2. Azure SQL database - Deploys an SQL database given the SQL server
3. Azure SQL firewall rule - Deploys a firewall rule to allow access to the SQL server from specific IPs
4. Azure SQL Action - Allows you to roll the password for the specified SQL server

## Deploying SQL Resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

**Note**  Don't forget to set the Service Principal ID, Service Principal secret, Tenant ID and Subscription ID as environment variables

### Azure SQL server

For instance, this is the sample YAML for the Azure SQL server.

  ```yaml
    apiVersion: azure.microsoft.com/v1
    kind: SqlServer
    metadata:
    name: sqlserver-sample
    spec:
     location: westus
     resourcegroup: resourceGroup1
  ```

The value for kind, `SqlServer` is the Custom Resource Definition (CRD) name.
`sqlserver-sample` is the name of the SQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the SQL server at and the Resource group in which you want to create it under. The `allowazureserviceaccess' boolean allows you to specify if you want Azure services to have access to your SQL server.

Once you've updated the YAML with the settings you need, and you have the operator running, you can create a Custom SQL server resource using the command.

```bash
kubectl apply -f config/samples/azure_v1_sqlserver.yaml
```

Along with creating the SQL server, this operator also generates the admin username and password for the SQL server and stores it in a kube secret with the same name as the SQL server.

You can retrieve this secret using the following command for the sample YAML

```bash
kubectl get secret sqlserver-sample -o yaml
```

This would show you the details of the secret. `username` and `password` in the `data` section are the base64 encoded admin credentials to the SQL server.

```bash
apiVersion: v1
data:
  fullyqualifiedservername: c3Fsc2VydmVyLXNhbXBsZS04ODguZGF0YWJhc2Uud2luZG93cy5uZXQ=
  fullyqualifiedusername: aGFzMTUzMnVAc3Fsc2VydmVyLXNhbXBsZS04ODg=
  password: XTdpMmQqNsd7YlpFdEApMw==
  sqlservername: c3Fsc2VyfmVyLXNhbXBsZS04ODg=
  username: aGFzMTFzMnU=
kind: Secret
metadata:
  creationTimestamp: "2019-10-09T21:02:02Z"
  name: sqlserver-sample-888
  namespace: default
  ownerReferences:
  - apiVersion: azure.microsoft.com/v1
    blockOwnerDeletion: true
    controller: true
    kind: SqlServer
    name: sqlserver-sample-888
    uid: 08fdbf42-ead8-11e9-91e0-025000000001
  resourceVersion: "131163"
  selfLink: /api/v1/namespaces/default/secrets/sqlserver-sample-888
  uid: 0aeb2429-ead8-11e9-91e0-025000000001
type: Opaque
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

The SQL firewall operator allows you to add a SQL firewall rule to the SQL server.

Below is the sample YAML for SQL firewall rule

```yaml
apiVersion: azure.microsoft.com/v1
kind: SqlFirewallRule
metadata:
  name: sqlf-allowazuresvcaccess
spec:
  resourcegroup: ResourceGroup1
  server:  sqlserver-sample
  
  # this IP range enables Azure Service access
  startipaddress: 0.0.0.0
  endipaddress: 0.0.0.0
```

The `server` indicates the SQL server on which you want to configure the new SQL firewall rule on and `resourcegroup` is the resource group of the SQL server. The `startipaddress` and `endipaddress` indicate the IP range of sources to allow access to the SQL server.

When the `startipadress` and `endipaddress` are 0.0.0.0, it is a special case that adds a firewall rule to allow all Azure services to access the SQL server.

### SQL Action

The SQL Action operator is used to trigger an action on the SQL server. Right now, the only action supported is `rollcreds` which rolls the password for the SQL server to a new one.

Below is a sample YAML for rolling the password

```yaml
apiVersion: azure.microsoft.com/v1
kind: SqlAction
metadata:
  name: Sql-rollcreds-action
spec:
  resourcegroup: ResourceGroup1
  actionname: rollcreds
  servername: sqlserver-sample
```

The `name` is a name for the action that we want to trigger. The type of action is determined by the value of `actionname` in the spec which in this case is `rollcreds`. The `resourcegroup` and `servername` identify the SQL server on which the action should be triggered on.

Once you apply this, the kube secret with the name as the SQL server is updated with the rolled password.

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

## Demo

Watch this demo <https://bit.ly/2lUIX6Y> to observe how you would you use the Azure SQL Operator from a real application.

In this demo, we use YAML to deploy the application and the SQL server. Once the SQL server is provisioned, the connection details are stored as secrets which the application can use to access the SQL server.
