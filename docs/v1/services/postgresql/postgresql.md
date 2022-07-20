# PostgreSQL Operator

## Resources Supported

The PostgreSQL operator suite consists of the following operators.

1. PostgreSQL server - Deploys an `Azure Database for PostgreSQL server` given the Location and Resource group
2. PostgreSQL database - Deploys a database under the given `Azure Database for PostgreSQL server`
3. PostgreSQL firewall rule - Deploys a firewall rule to allow access to the `Azure Database for PostgreSQL server` from the specified IP range

### PostgreSQL server

Here is a [sample YAML](/config/samples/azure_v1alpha1_postgresqlserver.yaml) for the PostgreSQL server.

The value for kind, `PostgreSQLServer` is the Custom Resource Definition (CRD) name.
`postgresqlserver-sample` is the name of the PostgreSQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the PostgreSQL server at and the Resource group in which you want to create it under. It also contains other values that are required to create the server like the `serverVersion`, `sslEnforcement` and the `sku` information.

If `sslEnforcement` is enabled, applications can connect to the PostgreSQL server using SSL. If you would like to connect using the full SSL verification enabled (sslmode=verify-full) that validates the server certificate, you would need the root certificate installed on your client. [This link](https://docs.microsoft.com/en-us/azure/postgresql/concepts-ssl-connection-security) documents the root certificate to use.

**Note** The root certificate documented in the above link (`https://www.digicert.com/CACerts/BaltimoreCyberTrustRoot.crt.pem`) only applies to instances on Azure Public Cloud. If you are deploying Azure Database for PostgreSQL instances on Azure China Cloud, you should use this one - `https://dl.cacerts.digicert.com/DigiCertGlobalRootCA.crt.pem`

Along with creating the PostgreSQL server, this operator also generates the admin username and password for the PostgreSQL server and stores it in a kube secret or keyvault (based on what is specified) with the same name as the PostgreSQL server.

This secret contains the following fields.

- `fullyqualifiedservername` : Fully qualified name of the PostgreSQL server such as postgresqlservername.postgres.database.azure.com
- `postgresqlservername` : PostgreSQL server name
- `username` : Server admin
- `password` : Password for the server admin
- `fullyqualifiedusername` : Fully qualified user name that is required by some apps such as <username>@<postgresqlserver>

For more information on where and how secrets are stored, look [here](/docs/v1/howto/secrets.md)

##### Read Replicas in Azure Database for PostgreSQL

The PostgreSQL server operator can also be used to create Read Replicas given the `sourceserverid`.

The replica inherits the admin account from the master server. All user accounts on the master server are replicated to the read replicas. 

For more information on read replicas, refer [here](https://docs.microsoft.com/en-us/azure/postgresql/concepts-read-replicas)

### PostgreSQL Database

Here is a [sample YAML](/config/samples/azure_v1alpha1_postgresqldatabase.yaml) for PostgreSQL database

Update the `resourcegroup` to where you want to provision the PostgreSQL database. `server` is the name of the PostgreSQL server where you want to create the database in.

### PostgreSQL firewall rule

The PostgreSQL firewall rule operator allows you to add a firewall rule to the PostgreSQL server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_postgresqlfirewallrule.yaml) for PostgreSQL firewall rule

The `server` indicates the PostgreSQL server on which you want to configure the new PostgreSQL firewall rule on and `resourceGroup` is the resource group of the PostgreSQL server. The `startIpAddress` and `endIpAddress` indicate the IP range of sources to allow access to the SQL server.

*Note*: When the `startIpAddress` and `endIpAddress` are 0.0.0.0, it denotes a special case that adds a firewall rule to allow all Azure services to access the SQL server.

## Deploy, view and delete resources

You can follow the steps [here](/docs/v1/howto/resourceprovision.md) to deploy, view and delete resources.
