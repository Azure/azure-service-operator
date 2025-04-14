---
title: Database Users
---

ASOv1 supports 3 types of database user:
- Azure SQL
- MySQL
- PostgreSQL

Unlike other resources, these users do not exist as first class entities in the Azure API. This means
that `asoctl` will not import them.

To migrate database users from ASOv1 to ASOv2, you must manually transform the ASOv1 DatabaseUser shape
into the ASOv2 shape.

## Azure SQL

ASOv1 `AzureSQLUser` [example](https://github.com/Azure/azure-service-operator/blob/asov1/config/samples/azure_v1alpha1_azuresqluser.yaml).

ASOv2 `sql.azure.com/User` [example](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/sql/v1api/v1api/v1_user.yaml).

## MySQL

ASOv1 `MySQLUser` [example](https://github.com/Azure/azure-service-operator/blob/asov1/config/samples/azure_v1alpha2_mysqluser.yaml).

ASOv1 `MySQLAADUser` [example](https://github.com/Azure/azure-service-operator/blob/asov1/config/samples/azure_v1alpha2_mysqlaaduser.yaml).

ASOv2 `dbformysql.azure.com/User` [example](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/dbformysql/v1api/v1_user.yaml).

ASOv2 `dbformysql.azure.com/User` in AAD configuration [example](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/dbformysql/v1api/v1_user_aad.yaml).
- Note that MySQL Flexible server has certain requirements for how it is set up to work with AAD [here](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/dbformysql/v1api/v1_user_aad.yaml#L3).
  So not only do you need to create the user in ASOv2 you must also ensure that the FlexibleServer has the right 
  [FlexibleServerAdministrator](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/dbformysql/v1api20230630/v1api20230630_flexibleserversadministrator.yaml) configured, with the right permissions. 

## PostgreSQL

ASOv1 `PostgreSQLUser` [example](https://github.com/Azure/azure-service-operator/blob/asov1/config/samples/azure_v1alpha1_postgresqluser.yaml).

ASOv2 `dbforpostgresql.azure.com/User` [example](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/dbforpostgresql/v1api/v1_user.yaml).