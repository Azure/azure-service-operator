# Tips

For your Cassandra resources to deploy correctly, ensure the Azure Cosmos DB Service Principal has the `Microsoft.Network/virtualNetworks/subnets/join/action` to your subnet. Without it, deployment will fail.

See https://learn.microsoft.com/en-us/azure/managed-instance-apache-cassandra/add-service-principal

You can create the relevant RoleAssignment with ASO ... BUT you currently need to look up the ObjectID manually:

``` bash
az ad sp show --id a232010e-820c-4083-83bb-3ace5fc29d0b --query id -o tsv
```

Review `refs/v20220401_roleassignment_cassandra-dc-subnet.yaml` and `refs/v20220401_roleassignment_cassandra-mgmt-subnet.yaml` to see how this ObjectID is used.

