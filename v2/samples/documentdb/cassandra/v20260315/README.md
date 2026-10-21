# Tips

For your Cassandra resources to deploy correctly, the Azure Cosmos DB service
principal needs `Microsoft.Network/virtualNetworks/subnets/join/action` on both
subnets.

See https://learn.microsoft.com/en-us/azure/managed-instance-apache-cassandra/add-service-principal

Creating the Cassandra cluster registers the service principal in the tenant.
The `ServicePrincipal` sample adopts it and exports its tenant-specific object
ID to the `cassandra-service-principal` ConfigMap. Both `RoleAssignment`
resources read that value from the ConfigMap and grant Network Contributor on
their respective subnets.
