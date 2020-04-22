# MySQL Operator

## Resources Supported

The MySQL operator suite consists of the following operators.

1. MySQL server - Deploys an `Azure Database for MySQL server` given the Location, Resource group and other properties. This operator also helps creating read replicas for MySQL server.
2. MySQL database - Deploys a database under the given `Azure Database for MySQL server`
3. MySQL firewall rule - Deploys a firewall rule to allow access to the `Azure Database for MySQL server` from the specified IP range

### MySQL server

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlserver.yaml) for the MySQL server.

The value for kind, `MySQLServer` is the Custom Resource Definition (CRD) name.
`mysqlserver-sample` is the name of the MySQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the server at and the Resource group in which you want to create it under. It also contains other values that are required to create the server like the `serverVersion`, `sslEnforcement` and the `sku` information.

Along with creating the MySQL server, this operator also generates the admin username and password for the MySQL server and stores it in a kube secret or keyvault (based on what is specified) with the same name as the MySQL server.

This secret contains the following fields.

- `fullyqualifiedservername` : Fully qualified name of the MySQL server such as mysqlserver.mysql.database.azure.com
- `mysqlservername` : MySQL server name
- `username` : Server admin
- `password` : Password for the server admin
- `fullyqualifiedusername` : Fully qualified user name that is required by some apps such as <username>@<mysqlserver>

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

The `server` indicates the MySQL server on which you want to configure the new MySQL virtual network rule on and `resourceGroup` is the resource group of the MySQL server. Provide the virtual network name and subnet name in the variables `vNetName` and `subnetName`, and `vNetResourceGroup` is the resource group the virtual network is located in. The `ignoreMissingServiceEndpoint` indicates whether or not to create virtual network rule before the virtual network has vnet service endpoint enabled.

*Note*: When using MySQL Virtual Network Rules, the `Basic` SKU is not a valid op

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
