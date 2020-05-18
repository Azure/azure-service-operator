# Information about the Resource post deployment

Many of the Azure resources have access information like connection strings, access keys, admin user password etc. that is required by applications consuming the resource. This information is stored as secrets after resource creation.

The operator provides two options to store these secrets:

1. **Kubernetes secrets**: This is the default option. If the secretname is not specified in the Spec, a secret with the same name as  ObjectMeta.name is created.

2. **Azure Keyvault secrets**: You can specify the name of the Azure Keyvault to use to store secrets through the environment variable `AZURE_OPERATOR_KEYVAULT`.

If the secretname is not specified in the Spec, the secret in Keyvault is normally created with the name `<kubernetes-namespace-of=resource>-<ObjectMeta.name>`. The namespace is preprended to the name to avoid name collisions across Kubernetes namespaces.


```
export AZURE_OPERATOR_KEYVAULT=OperatorSecretKeyVault
```

Some things to note about this Keyvault:
(i) The Keyvault should have an access policy added for the identity under which the Operator runs as. This access policy should include  at least `get`, `set`, `list` and `delete` Secret permissions.

(ii) You can use a Keyvault with "Soft delete" enabled. However, you cannot use a Keyvault with "Purge Protection" enabled, as this prevents the secrets from being deleted and causes issues if a resource with the same name is re-deployed.


## Per resource Keyvault

In addition to being able to specify an Azure Keyvault to store secrets, you also have the option to specify a different Keyvault per resource.

Some situations may require that you use a different Keyvault to store the admin password for the Azure SQL server from the Keyvault used to store the connection string for eventhubs. You can specify the Keyvault name in the Spec field  `keyVaultToStoreSecrets`. When this is specified, the secrets from provisioning of that resource will be stored in this Keyvault instead of the global one configured for the operator.