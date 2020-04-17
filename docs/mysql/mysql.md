# MySQL Operator

The MySQL Suite is made up of the following operators:
* MySQL Server
* MySQL Database
* MySQL Firewall Rule
* MySQL Virtual Network Rule

Learn more about Azure Database for MySQL [here](https://docs.microsoft.com/en-us/azure/mysql/).

## MySQL Server

Create a MySQL Server in a specified resource group and location.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlserver.yaml) to provision a MySQL Server.

The spec is comprised of the following fields:

* `sku`
    * `name` The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
    * `tier` The tier of the particular SKU. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'
    * `family` The family of hardware.
    * `size` The size code
    * `capacity` The scale up/out capacity, representing server's compute units.
* `serverVersion` The version of a server. Possible values include: "8.0", "5.7", "5.6"
* `sslEnforcement` Enable ssl enforcement or not when connect to server

## MySQL Database

Create a MySQL Database in a specified resource group and location for an existing MySQL Server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqldatabase.yaml) to provision a MySQL Server.

## MySQL Firewall Rule

Create a MySQL Firewall Rule in a specified resource group and location for an existing MySQL Server.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlfirewallrule.yaml) to provision a Firewall Rule for your MySQL Server.

The spec is comprised of the following fields:

* `server` the name of the existing MySQL Server
* `startIpAddress` starting IP Address Range 
* `endIpAddress` ending IP Address Range

When the `startIpAddress` and `endIpAddress` are '0.0.0.0', it is a special case that adds a firewall rule to allow all Azure services to access the MySQL server.

## MySQL Virtual Network Rule

Create a MySQL Database in a specified resource group and location. Requires a MySQL server to already be provisioned.

Here is a [sample YAML](/config/samples/azure_v1alpha1_mysqlvnetrule.yaml) to provision a Virtual Network Rule for your MySQL Server.

The spec is comprised of the following fields:

* `server` the name of the existing MySQL Server
* `vNetResourceGroup` resource group of the Virtual Network
* `vNetName` name of the virtual network
* `subnetName` name of the subnet within the Virtual Network
* `ignoreMissingServiceEndpoint` Create firewall rule before the virtual network has vnet service endpoint enabled.

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
