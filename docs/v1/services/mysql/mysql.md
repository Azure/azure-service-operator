# MySQL Operator

## Resources Supported

The MySQL operator suite consists of the following operators.

1. MySQL server - Deploys an `Azure Database for MySQL server` given the Location, Resource group and other properties. This operator also helps creating read replicas for MySQL server.
2. MySQL database - Deploys a database under the given `Azure Database for MySQL server`
3. MySQL firewall rule - Deploys a firewall rule to allow access to the `Azure Database for MySQL server` from the specified IP range
4. MySQL virtual network rule - Deploys a virtual network rule place the `Azure Database for MySQL server` into an Azure Virtual Network Subnet.
4. MySQL administrator - Sets the AAD Administrator of the `Azure Database for MySQL server` to the specified AAD identity.
5. MySQL user - Deploys a user into the `Azure Database for MySQL server`.
6. MySQL AAD User - Deploys an AAD user into the `Azure Database for MySQL server`.

> **Deprecation notice**: Azure Database for MySQL - Single Server is on the retirement path and is [scheduled for retirement by September 16, 2024](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server).  
> Existing instances can be migrated to Azure Database for MySQL - Flexible Server using the Azure Database migration Service.  
> Azure Database for MySQL - Flexible Server is [fully supported in ASO v2](https://azure.github.io/azure-service-operator/reference/#dbformysql).

### MySQL server

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlserver.yaml) for the MySQL server.

The value for kind, `MySQLServer` is the Custom Resource Definition (CRD) name.
`mysqlserver-sample` is the name of the MySQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the server at and the Resource group in which you want to create it under. It also contains other values that are required to create the server like the `serverVersion`, `sslEnforcement` and the `sku` information.

The `adminSecret` is optional and if provided must point to a Kubernetes secret containing a `username` and `password` field. If not specified, the operator will generate an administrator account `username` and `password`.
Along with creating the MySQL server, this operator also generates the admin `username` and `password` for the MySQL server and stores it in a kube secret or keyvault (based on what is specified). The generated secret is named according to [secrets naming](/docs/secrets.md).

This secret contains the following fields.
| Secret field               | Content                                                                                       |
| -------------------------- | --------------------------------------------------------------------------------------------- |
| `fullyQualifiedServerName` | Fully qualified name of the MySQL server. Example: `mysqlserver.mysql.database.azure.com`.    |
| `mySqlServerName`          | MySQL server name.                                                                            |
| `username`                 | Server admin account name.                                                                    |
| `password`                 | Server admin account password.                                                                |
| `fullyQualifiedUsername`   | Fully qualified user name that is required by some apps. Example: `<username>@<mysqlserver>`. |

For more information on where and how secrets are stored, look [here](/docs/secrets.md)

#### Read Replicas in Azure Database for MySQL

The MySQL server operator can also be used to create Read Replicas given the `sourceserverid` and the `location`.

The replica inherits all other properties including the admin username and password from the source server.

The operator reads the admin username and password for the source server from its secret (if available) and creates a secret with the same fields as described above for the replica.

For more information on read replicas, refer [here](https://docs.microsoft.com/en-us/azure/mysql/concepts-read-replicas)

### MySQL Database

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqldatabase.yaml) for MySQL database

Update the `resourcegroup` to where you want to provision the MySQL database. `server` is the name of the MySQL server where you want to create the database in.

### MySQL firewall rule

The MySQL firewall rule operator allows you to add a firewall rule to the MySQL server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlfirewallrule.yaml) for MySQL firewall rule

The `server` indicates the MySQL server on which you want to configure the new MySQL firewall rule on and `resourceGroup` is the resource group of the MySQL server. The `startIpAddress` and `endIpAddress` indicate the IP range of sources to allow access to the server.

*Note*: When the `startIpAddress` and `endIpAddress` are 0.0.0.0, it denotes a special case that adds a firewall rule to allow all Azure services to access the server.

### MySQL virtual network rule

The MySQL virtual network rule operator allows you to add virtual network rules to the MySQL server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlvnetrule.yaml) for MySQL virtual network rule. 

The `server` indicates the MySQL server on which you want to configure the new MySQL virtual network rule on and `resourceGroup` is the resource group of the MySQL server. Provide the virtual network name and subnet name in the variables `vNetName` and `subnetName`, and `vNetResourceGroup` is the resource group the virtual network is located in. The `ignoreMissingServiceEndpoint` indicates whether or not to create virtual network rule before the virtual network has vnet service endpoint enabled.

*Note*: When using MySQL Virtual Network Rules, the `Basic` SKU is not a valid op

### MySQL administrator

The MySQL administrator operator allows you to add an [AAD administrator](https://docs.microsoft.com/azure/mysql/concepts-azure-ad-authentication) to the MySQL server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlserveradministrator.yaml).

### MySQL user

The MySQL user operator allows you to add a new user to an existing MySQL database. 

Here is a [sample YAML](/config/samples/azure_v1alpha2_mysqluser.yaml) for MySQL user. 

The `resourceGroup` is the resource group of the MySQL server and MySQL database, provide the MySQL server name in `server` and MySQL database name in `dbName`. 

The operator supports granting the user global privileges using the `roles` field, which can contain zero or more privileges from the following list:

* `CREATE TABLESPACE`
* `CREATE USER`
* `FILE`
* `PROCESS`
* `RELOAD`
* `REPLICATION CLIENT`
* `REPLICATION SLAVE`
* `SHOW DATABASES`
* `SHUTDOWN`
* `SUPER`

Users can also be granted privileges on all objects in a specific database using the `databaseRoles` field - this is a map of database name to a list of privileges that should be granted in that database.
The following privileges can be set:

* `SELECT`
* `INSERT`
* `UPDATE`
* `DELETE`
* `CREATE`
* `DROP`
* `REFERENCES`
* `INDEX`
* `ALTER`
* `CREATE TEMPORARY TABLES`
* `LOCK TABLES`
* `EXECUTE`
* `CREATE VIEW`
* `SHOW VIEW`
* `CREATE ROUTINE`
* `ALTER ROUTINE`
* `EVENT`
* `TRIGGER`

If the special value `ALL` is used, all of the above privileges will be granted.
`ALL` can't be used in the `roles` field because the MySQL administrator user the operator uses doesn't have sufficient privileges to grant all global privileges.

The username is defined by `username`. The MySQL server admin secret is stored in the secret with name `adminSecret` in the  keyvault named `adminSecretKeyVault`. 

### MySQL AAD user
The MySQL AAD user operator allows you to add a new AAD user to an existing MySQL database.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlaaduser.yaml).

This controller is only avilable when using [Managed Identity authentication](https://github.com/Azure/azure-service-operator/blob/main/docs/v1/howto/managedidentity.md) with ASO.
Attempting to use it without Managed Identity will result in an authentication error.

The AAD identity the operator is running as must have permissions to create users in the MySQLServer. 
This is most commonly granted by making the operator managed identity the MySQL Administrator using the [MySQL administrator](#mysql-administrator) operator described above.

MySQL AAD users can be granted server- and database-level privileges using the `roles` and `databaseRoles` fields in the same way as [MySQL users](#mysql-user).

## Deploy, view and delete resources

You can follow the steps [here](/docs/v1/howto/resourceprovision.md) to deploy, view and delete resources.
