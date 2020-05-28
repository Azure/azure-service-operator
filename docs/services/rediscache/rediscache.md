# Redis Cache Operator

This operator deploys an Azure Cache for Redis into a specified resource group at the specified location.

Learn more about Azure Cache for Redis [here](https://docs.microsoft.com/en-us/azure/azure-cache-for-redis/cache-overview).

Here is a [sample YAML](/config/samples/azure_v1alpha1_rediscache.yaml) to provision an Azure Cache for Redis.

The spec is comprised of the following fields:
- Location
- ResourceGroupName
- Properties
    - SKU
        - Name
        - Family
        - Capacity
    - EnableNonSslPort
    - SubnetID
    - StaticIP
    - Configuration
- SecretName
- KeyVaultToStoreSecrets

### Required Fields

A Redis Cache needs the following fields to deploy, along with a location and resource group.

* `Properties.SKU.Name` Select a SKU, where the options are: _Basic_, _Standard_, and _Premium_.
* `Properties.SKU.Family` Select a SKU Family, where the options are: _C_, _P_. If you selected a _Premium_ SKU, then the corresponding SKU Family is _P_.
* `Properties.SKU.Capacity` Set the desired capacity 
* `EnableNonSslPort` defaults to True

### Optional Fields

* `SecretName` specify the name of the secret. If none is given, it will fall back to the name of the redis cache.
* `KeyVaultToStoreSecrets` specify a Key Vault to store primary and secondary credentials in. If none is given, it will default to storing credentials as a Kube Secret.
* `Properties.SubnetID` specify a subnet ID to place the Redis Cache in
* `Properties.StaticIP` specify a statis IP for the Redis Cache
* `Properties.Configuration` specify configuration values as key value pairs for the Redis Cache

### Secrets

After creating an Azure Cache for Redis instance, the operator stores a JSON formatted secret with the following fields. For more details on where the secrets are stored, look [here](/docs/secrets.md).

* `primaryKey`
* `secondaryKey`

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
