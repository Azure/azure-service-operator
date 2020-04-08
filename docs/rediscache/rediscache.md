# Redis Cache Operator

This operator will deploy an Azure Cache for Redis into a resource group and location.

Learn more about Azure Cache for Redis [here](https://docs.microsoft.com/en-us/azure/azure-cache-for-redis/cache-overview).

A Redis Cache is comprised of the following fields:
- Location
- ResourceGroupName
- Properties
    - SKU
        - RedisCacheSku
        - RedisCacheSkuFamily
        - Capacity
    - EnableNonSslPort
    - SubnetID
    - StaticIP
    - Configuration
- SecretName
- KeyVaultToStoreSecrets

### Required Fields

A Redis Cache needs the following fields to deploy, along with a location and resource group.

* `Properties.SKU.RedisCacheSku` Select a SKU, where the options are: _Basic_, _Standard_, and _Premium_.
* `Properties.SKU.RedisCacheSkuFamily` Select a SKU Family, where the options are: _C_, _P_. If you selected a _Premium_ SKU, then the corresponding SKU Family is _P_.
* `Properties.SKU.Capacity` Set the desired capacity 
* `EnableNonSslPort` defaults to True

### Optional Fields

* `SecretName` specify the name of the secret. If none is given, it will fall back to a default name.
* `KeyVaultToStoreSecrets` specify a Key Vault to store primary and secondary credentials in. If none is given, it will default to storing credentials as a Kube Secret.
* `Properties.SubnetID` specify a subnet ID to place the Redis Cache in
* `Properties.StaticIP` specify a statis IP for the Redis Cache
* `Properties.Configuration` provide a configuration to the Redis Cache

## Deploy

Follow the steps [here](/docs/development.md) or [here](/docs/deploy.md) to either run the operator locally or in a real Kubernetes cluster.

You can find a sample Redis Cache YAML [here](/config/samples/azure_v1alpha1_rediscache.yaml).