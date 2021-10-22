# Redis Cache Operator

## Resources Supported

The RedisCache operator suite consists of the following operators.

1. Redis Cache - Deploys an Azure Cache for redis into a specified resource group at the specified location
2. Redis Cache Firewall Rule - Deploys a firewall rule to allow access to the RedisCache from the specified IP range

### RedisCache
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

#### Required Fields

A Redis Cache needs the following fields to deploy, along with a location and resource group.

* `Properties.SKU.Name` Select a SKU, where the options are: _Basic_, _Standard_, and _Premium_.
* `Properties.SKU.Family` Select a SKU Family, where the options are: _C_, _P_. If you selected a _Premium_ SKU, then the corresponding SKU Family is _P_.
* `Properties.SKU.Capacity` Set the desired capacity 
* `EnableNonSslPort` defaults to True

#### Optional Fields

* `SecretName` specify the name of the secret. If none is given, it will fall back to the name of the redis cache.
* `KeyVaultToStoreSecrets` specify a Key Vault to store primary and secondary credentials in. If none is given, it will default to storing credentials as a Kube Secret.
* `Properties.SubnetID` specify a subnet ID to place the Redis Cache in
* `Properties.StaticIP` specify a statis IP for the Redis Cache
* `Properties.Configuration` specify configuration values as key value pairs for the Redis Cache

#### Secrets

After creating an Azure Cache for Redis instance, the operator stores a JSON formatted secret with the following fields. For more details on where the secrets are stored, look [here](/docs/secrets.md).

* `primaryKey`
* `secondaryKey`

## RedisCache firewall rule

The RedisCache firewall rule allows you to add a firewall rule to RedisCache.

Here is a [sample YAML](https://github.com/Azure/azure-service-operator/blob/main/config/samples/azure_v1alpha1_rediscachefirewallrule.yaml) for RedisCache firewall rule

The `redisCache` indicates the RedisCache on which you want to configure the new RedisCache firewall rule on and `resourceGroup` is the resource group of the RedisCache. The `startIP` and `endIP`  under Properties indicates the IP range of sources to allow access to the RedisCache.

_Note:_ When the `startIP` and `endIP` are 0.0.0.0, it denotes a special case that adds a firewall rule to allow all Azure services to access the RedisCache.

## RedisCache action

The RedisCache action allows you to regenerate keys and reboot the RedisCache cluster.

Here is a [sample YAML](/config/samples/azure_v1alpha1_rediscacheaction.yaml) for RedisCache action.

The `cacheName` indicates the RedisCache on which you want to perform the action and `resourceGroup` is the resource group of the RedisCache. The `actionName` corresponds to one of the supported actions listed below. 

### RedisCache action - Roll Keys
The `secretName` field is used to update the RedisCache secret. The `keyVaultToStoreSecrets` field is used to specify a KeyVault instance where the RedisCache secret exists. The following "roll" actions are supported:
- `rollprimarykey` - regenerates primary key and updates the secret
- `rollsecondarykey` - regenerates secondary key and updates the secret
- `rollallkeys` - regenerates primary and secondary keys and updates the secret

### RedisCache action - Reboot
The `shardID` field is used to specify a specific RedisCache shard to reboot. The following "reboot" actions are supported:
- `rebootprimarynode` - reboots all primary nodes in the RedisCache cluster
- `rebootsecondarynode` - reboots all secondary nodes in the RedisCache cluster
- `rebootallnodes` - reboots all nodes (primary & secondary) in the RedisCache cluster

## Deploy, view and delete resources

You can follow the steps [here](/docs/v1/howto/resourceprovision.md) to deploy, view and delete resources.
