# Azure SQL Operator

## Resources Supported

The Azure SQL operator suite consists of the following operators.

1. Azure SQL server - Deploys an Azure SQL server given the location and Resource group
2. Azure SQL database - Deploys an SQL database given the SQL server
3. Azure SQL firewall rule - Deploys a firewall rule to allow access to the SQL server from specific IPs
4. Azure SQL VirtualNetwork rule - Deploys a VirtualNetwork rule to allow access to the SQL server from specific VNets
4. Azure SQL Action - Allows you to roll the password for the specified SQL server
5. Azure SQL failover group - Deploys a failover group on a specified Azure SQL server given the secondary server and the databases to failover
6. Azure SQL User - Creates an user on the specified Azure SQL database and stores the username/password as secrets
7. Azure SQL Managed User - Binds a specified Managed Identity as a user on the SQL database

### Azure SQL server

Here is a [sample YAML](/config/samples/azure_v1alpha1_azuresqlserver.yaml) for the Azure SQL server.

The value for kind, `AzureSqlServer` is the Custom Resource Definition (CRD) name and `sqlserver-sample` in this case is the name of the SQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the Azure SQL server at and the Resource group in which you want to create it under.

Deploying a SQL Server instance requires that you deploy a ResourceGroup, an AzureSqlDatabase, and AzureSqlServer CRDs.

The project maintains a [set of samples](https://github.com/Azure-Samples/azure-service-operator-samples) of how to utilize the Azure Service Operator.

As part of this, there is a sample voting app that utilizes a SQL server. Clone this repo, adjust the names in the resource group and SQL Server CRDs such that they are unique and then deploy them with:

```bash
$ kubectl apply -f azure-votes-sql/manifests/azure_v1_resourcegroup.yaml
$ kubectl apply -f azure-votes-sql/manifests/azure_v1_sqldatabase.yaml
$ kubectl apply -f azure-votes-sql/manifests/azure_v1_sqlserver.yaml
```

Along with creating the SQL server, this operator also generates the admin username and password for the SQL server and stores it in a kube secret with the same name as the SQL server.

This secret contains the following fields.

- `fullyqualifiedservername` : Fully qualified name of the SQL server such as sqlservername.database.windows.net
- `sqlservername` : SQL server name
- `username` : Server admin
- `password` : Password for the server admin
- `fullyqualifiedusername` : Fully qualified user name that is required by some apps such as <username>@<sqlserver>

You can retrieve this secret using the following command for the sample YAML

```bash
kubectl get secret sqlserver-sample -o yaml
```

This would show you the details of the secret. `username` and `password` in the `data` section are the base64 encoded admin credentials to the SQL server.

```bash
apiVersion: v1alpha1
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
    kind: AzureSqlServer
    name: sqlserver-sample-888
    uid: 08fdbf42-ead8-11e9-91e0-025000000001
  resourceVersion: "131163"
  selfLink: /api/v1/namespaces/default/secrets/sqlserver-sample-888
  uid: 0aeb2429-ead8-11e9-91e0-025000000001
type: Opaque
```

### SQL Database

Here is a [sample YAML](/config/samples/azure_v1alpha1_azuresqldatabase.yaml) for SQL database

Update the `location` and the `resourcegroup` to where you want to provisiong the SQL database. `server` is the name of the Azure SQL server where you want to create the database in.
The `edition` represents the SQL database edition you want to use when creating the resource and should be one of the values above.

### SQL firewall rule

The SQL firewall operator allows you to add a SQL firewall rule to the SQL server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_azuresqlfirewallrule.yaml) for SQL firewall rule

The `server` indicates the SQL server on which you want to configure the new SQL firewall rule on and `resourcegroup` is the resource group of the SQL server. The `startipaddress` and `endipaddress` indicate the IP range of sources to allow access to the SQL server.

When the `startipadress` and `endipaddress` are 0.0.0.0, it is a special case that adds a firewall rule to allow all Azure services to access the SQL server.

### SQL VirtualNetwork rule

The SQL VNet rule operator allows you to add a SQL VNet (VirtualNetwork()) rule to the SQL server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_azuresqlvnetrule.yaml) for SQL VNet rule

The `server` indicates the SQL server on which you want to configure the new SQL VNet rule on and `resourceGroup` is the resource group of the SQL server.
`vNetResourceGroup`, `vNetName` and `subnetName` identify the Subnet on the particular VirtualNetwork that should be added to the rule.
`ignoreMissingServiceEndpoint` tells the operator to go ahead and add the rule with the specified VirtualNetwork/Subnet even if that VirtualNetwork does not have the Microsoft.Sql Service Endpoint enabled

### SQL Action

The SQL Action operator is used to trigger an action on the SQL server. Right now, the supported actions are `rolladmincreds`, which rolls the password for the SQL server to a new one, and `rollusercreds`, which rolls the password for the SQL Database User.

The `name` is a name for the action that we want to trigger. The type of action is determined by the value of `actionname` in the spec which is `rolladmincreds` or `rollusercreds` if you want to roll the password (Note: This action name should be exact for the password to be rolled). The `resourcegroup` and `servername` identify the SQL server on which the action should be triggered on.

Here is a [sample YAML](/config/samples/azure_v1alpha1_azuresqlaction.yaml) for rolling the admin or user password of the SQL server

Once you apply this, the kube or Key Vault secret with the same name as the SQL server is updated with the rolled password.

#### rolladmincreds

This action will roll the password for the SQL Server. 

Optional parameters:
  * `serverSecretKeyVault` - specify a Key Vault to store the admin secret credentials in. If there is no Key Vault specified, it will default to the global secret client.
  * `serverAdminSecretName` - specify a secret name for the admin secret. If there is no secret name specified, it will fallback to the default name.

#### rollusercreds

This action will roll the password for the SQL Database User. 

Required parameters:
  * `dbName` - name of the database
  * `dbUser` - name of the user

Optional parameters:
  * `userSecretKeyVault` - specify a Key Vault to store the user secret credentials in. If there is no Key Vault specified, it will default to the global secret client. 
  * `serverSecretKeyVault` - specify a Key Vault to store the admin secret credentials in. If there is no Key Vault specified, it will default to the global secret client.
  * `serverAdminSecretName` - specify a secret name for the admin secret. If there is no secret name specified, it will fallback to the default name.

### SQL failover group

The SQL failover group operator is used to create a failover group across two Azure SQL servers (one primary, one secondary). The servers should already be provisioned and deployed different regions. The specified databases will be replicated from the primary server and created on the secondary.

Here is a [sample YAML](/config/samples/azure_v1alpha1_azuresqlfailovergroup.yaml) for creating a failover group

The `name` is a name for the failover group that we want to create. `server` is the primary SQL server on which the failover group is created, `location` and `resourcegroup` are the location and the resource group of the primary SQL server. `failoverpolicy` can be "automatic" or "manual". `failovergraceperiod` is the time in minutes. `secondaryserver` is the secondary SQL server to failover to and `secondaryserverresourcegroup` is the resource group that the server is in. `databaselist` is the list of databased on the primary SQL server that should replicate to the secondary SQL server, when there is a failover action.

Once you apply this, a secret with the same name as the SQL failovergroup is also stored. This secret contains the fields for primary/secondary failovergroup listener endpoints (`readWriteListenerEndpoint` and `readOnlyListenerEndpoint`) and the primary/secondary SQL server names (`azureSqlPrimaryServerName` and `azureSqlSecondaryServerName`)

### SQL database user

The SQL user operator is used to create a user on the specified Azure SQL database. This user is more restrictive than the admin user created on the SQL server and is so recommended to use. The operator creates the user on the database by auto generating a strong password, and also stores the username and password as a secret (name can be specified in the YAML), so applications can use them.

User credentials are persisted in the secret store that has been configured for the operator runtime. The default secret contains the following fields:

- `azureSqlDatabaseName`
- `azureSqlServerName`
- `azureSqlServerNamespace`
- `fullyQualifiedServerName`
- `username`
- `password`

When Key Vault is configured, each value is Base64 encoded and the set of secrets is persisted as a single JSON document in the Key Vault.
The default secret name prefix in Key Vault is `azuresqluser-<serverName>-<azureSqlDatabaseName>`. Users can set the `keyVaultSecretPrefix` parameter to override this value.

Additionally, some client libraries support connecting directly to Key Vault to retrieve secrets. Users can set the `keyVaultSecretFormats` parameter so that explicit connection strings for their desired formats are added to the Key Vault. Each secret will be named after the secret prefix followed by the format name, for example: `azuresqluser-<serverName>-<azureSqlDatabaseName>-adonet`.
Here is a [sample YAML](/config/samples/azure_v1alpha1_azuresqluser.yaml) for creating a database user

The `name` is used to generate the username on the database. The exact name is not used but rather a UUID is appended to this to make it unique. `server` and `dbname` qualify the database on which you want to create the user on. `adminsecret` is the name of the secret where the username and password will be stored. `roles` specify the security roles that this user should have on the specified database.

### SQL Managed user

The SQL Managed user operator is used to enable the specified managed identity as a database user on the specified database. It also grants the specified database roles to the user.

**Pre-requisite: This operator requires that the operator is running as a Managed Identity user and that this Managed Identity user is configured as the AAD admin on the SQL server**

The operator also stores the following secrets (if using Keyvault) in addition to enabling Database access for the managed identity.

- Secret named `<KeyVaultSecretPrefix>-clientid` with value `<ManagedIdentityClientId>`
- Secret named `<KeyVaultSecretPrefix>-server` with value `Spec.Server`
- Secret named `<KeyVaultSecretPrefix>-dbName` with value `Spec.DbName`

*Note* If `KeyVaultSecretPrefix` is not specified `instance.Namespace` is used appended with `instance.Name`

If you are using Kube for storing secrets, there will be one secret with name `<instance.Name>` created with the above name-value in the secret map

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.

## Demo

Watch this demo <https://bit.ly/2lUIX6Y> to observe how you would you use the Azure SQL Operator from a real application.

In this demo, we use YAML to deploy the application and the SQL server. Once the SQL server is provisioned, the connection details are stored as secrets which the application can use to access the SQL server.
