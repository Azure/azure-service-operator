# PostgreSQL Operator

## Resources Supported

The Postgre SQL operator suite consists of the following operators.

1. Postgre SQL server - Deploys an `Azure Database for PostgreSQL server` given the Location and Resource group
2. Postgre SQL database - Deploys a database under the given `Azure Database for PostgreSQL server`
3. Postgre SQL firewall rule - Deploys a firewall rule to allow access to the `Azure Database for PostgreSQL server` from the specified IP range

## Deploying PostgreSQL Resources

You can follow the steps [here](/docs/development.md) or [here](/docs/deploy.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

**Note**  Don't forget to set the Service Principal ID, Service Principal secret, Tenant ID and Subscription ID as environment variables. If using Managed Identity, then set the appropriate environment variables too as per the docs above.

### PostgreSQL server

For instance, this is the sample YAML for the PostgreSQL server.

  ```yaml
    apiVersion: azure.microsoft.com/v1alpha1
    kind: PostgreSQLServer
    metadata:
    name: postgresqlserver-sample
    spec:
    location: westus2
    resourceGroup: resourcegroup-azure-operators
    serverVersion: "10"
    sslEnforcement: Enabled
    sku:
        name: B_Gen5_2
        tier: Basic
        family: Gen5
        size: "51200"
        capacity: 2
  ```

The value for kind, `PostgreSQLServer` is the Custom Resource Definition (CRD) name.
`postgresqlserver-sample` is the name of the PostgreSQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the PostgreSQL server at and the Resource group in which you want to create it under. It also contains other values that are required to create the server like the `serverVersion`, `sslEnforcement` and the `sku` information.

Once you've updated the YAML with the settings you need, and you have the operator running, you can create a Custom PostgreSQL server resource using the command.

```shell
kubectl apply -f config/samples/azure_v1alpha1_postgresqlserver.yaml
```

Along with creating the PostgreSQL server, this operator also generates the admin username and password for the PostgreSQL server and stores it in a kube secret or keyvault (based on what is specified) with the same name as the PostgreSQL server.

This secret contains the following fields.

- `fullyqualifiedservername` : Fully qualified name of the PostgreSQL server such as postgresqlservername.postgres.database.azure.com
- `postgresqlservername` : PostgreSQL server name
- `username` : Server admin
- `password` : Password for the server admin
- `fullyqualifiedusername` : Fully qualified user name that is required by some apps such as <username>@<postgresqlserver>

You can retrieve this secret using the following command for the sample YAML

```shell
kubectl get secret postgresqlserver-sample -o yaml
```

This would show you the details of the secret. `username` and `password` in the `data` section are the base64 encoded admin credentials to the PostgreSQL server.

```shell
apiVersion: v1alpha1
data:
  fullyqualifiedservername: c3Fsc2VydmVyLXNhbXBsZS04ODguZGF0YWJhc2Uud2luZG93cy5uZXQ=
  fullyqualifiedusername: aGFzMTUzMnVAc3Fsc2VydmVyLXNhbXBsZS04ODg=
  password: XTdpMmQqNsd7YlpFdEApMw==
  postgresqlservername: c3Fsc2VyfmVyLXNhbXBsZS04ODg=
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
    kind: PostgreSqlServer
    name: sqlserver-sample-888
    uid: 08fdbf42-ead8-11e9-91e0-025000000001
  resourceVersion: "131163"
  selfLink: /api/v1/namespaces/default/secrets/postgresqlserver-sample
  uid: 0aeb2429-ead8-11e9-91e0-025000000001
type: Opaque
```

### PostgreSQL Database

Below is the sample YAML for PostgreSQL database

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: PostgreSQLDatabase
metadata:
  name: postgresqldatabase-sample
spec:
  resourceGroup: resourcegroup-azure-operators
  server: postgresqlserver-sample
```

Update the `resourcegroup` to where you want to provision the PostgreSQL database. `server` is the name of the PostgreSQL server where you want to create the database in.

### PostgreSQL firewall

The PostgreSQL firewall operator allows you to add a firewall rule to the PostgreSQL server.

Below is the sample YAML for PostgreSQL firewall rule

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: PostgreSQLFirewallRule
metadata:
  name: postgresqlfirewallrule-sample
spec:
  resourceGroup: resourcegroup-azure-operators
  server: postgresqlserver-sample
  # this IP range enables Azure Service access
  startIpAddress: 0.0.0.0
  endIpAddress: 0.0.0.0
```

The `server` indicates the PostgreSQL server on which you want to configure the new PostgreSQL firewall rule on and `resourceGroup` is the resource group of the PostgreSQL server. The `startIpAddress` and `endIpAddress` indicate the IP range of sources to allow access to the SQL server.

*Note*: When the `startIpAddress` and `endIpAddress` are 0.0.0.0, it is a special case that adds a firewall rule to allow all Azure services to access the SQL server.

## View and Troubleshoot PostgreSQL Resources

To view your created PostgreSQL resources, such as PostgreSQLServer, run the following command:

```shell
kubectl get <CRD>
```

where CRD is the Custom Resource Definition name or `Kind` for the resource.

For instance, you can get the PostgreSQL servers provisioned using the command

```shell
kubectl get PostgreSqlServer
```

You should see the PostgreSqlServer instances as below

```shell
NAME                         AGE
postgresqlserver-sample      1h
```

If you want to see more details about a particular resource instance such as the `Status` or `Events`, you can use the below command

```shell
kubectl describe <Kind> <instance name>
```

For instance, the below command is used to get more details about the `postgresqlserver-sample` instance

```shell
kubectl describe PostgreSqlServer postgresqlserver-sample
```

```shell
Name:         postgresqlserver-sample
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"azure.microsoft.com/v1alpha1","kind":"PostgreSQLServer","metadata":{"annotations":{},"name":"postgresqlserver-sample","na...
API Version:  azure.microsoft.com/v1alpha1
Kind:         PostgreSQLServer
Metadata:
  Creation Timestamp:  2020-01-06T22:48:36Z
  Finalizers:
    azure.microsoft.com/finalizer
  Generation:        1
  Resource Version:  205041
  Self Link:         /apis/azure.microsoft.com/v1alpha1/namespaces/default/postgresqlservers/postgresqlserver-sample
  UID:               acdb1a76-30d6-11ea-9c12-0e1c98a89af8
Spec:
  Location:        westus2
  Resource Group:  resourcegroup-operators
  Server Version:  10
  Sku:
    Capacity:       2
    Family:         Gen5
    Name:           B_Gen5_2
    Size:           51200
    Tier:           Basic
  Ssl Enforcement:  Enabled
Status:
  Message:      Success
  Provisioned:  true
  State:        Ready
Events:
  Type     Reason           Age    From                         Message
  ----     ------           ----   ----                         -------
  Normal   Added            8m24s  PostgreSQLServer-controller  Object finalizer is added
  Normal   Added            8m24s  PostgreSQLServer-controller  Object finalizer is added
  Normal   Reconciled  4m57s (x2 over 4m57s)  PostgreSQLServer-controller  Successfully reconciled
  Normal   Reconciled  4m55s (x2 over 4m57s)  PostgreSQLServer-controller  Successfully reconciled
```

The `Status` section gives you the current state of the resource that it's Ready and is provisioned.

The `Events` have a chronological record of what occurred through the process of provisioning the resource.

## Delete a PostgreSQL Resource

To delete an existing resource from Kubernetes and Azure, use the following command.

```shell
kubectl delete <Kind> <instancename>
```

For instance, deleting the above PostgreSqlServer instance would look like this.

```shell
kubectl delete PostgreSqlServer postgresqlserver-sample
```

The following message should appear:

`postgresqlserver.azure.microsoft.com postgresqlserver-sample deleted.`
