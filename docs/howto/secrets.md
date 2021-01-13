# Information about the Resource post deployment

Many of the Azure resources have access information like connection strings, access keys, admin user password etc. that is required by applications consuming the resource.
This information is stored by the operator as secrets after resource creation.

The operator provides two options to store these secrets:

1. **Kubernetes secrets**: This is the default option. If the secretname is not specified in the Spec, a secret with the same name as  ObjectMeta.name is created.

2. **Azure Key Vault secrets**: You can specify the default Azure Key Vault to store secrets in through via the `AZURE_OPERATOR_KEYVAULT` field in the `azureoperatorsettings` secret.
This can be set by exporting it as an environment variable and then including it in the secret creation.
    ```
    export AZURE_OPERATOR_KEYVAULT=mykeyvault
    kubectl create secret generic azureoperatorsettings \
            ... (other settings) \
            --from-literal=AZURE_OPERATOR_KEYVAULT=${AZURE_OPERATOR_KEYVAULT}
    ```
    Or if you're using Helm:
    ```
    helm upgrade --install aso aso/azure-service-operator \
            --create-namespace \
            --namespace=azureoperator-system \
            ... \
            --set azureOperatorKeyvault=$AZURE_OPERATOR_KEYVAULT
    ```

If the secretname is not specified in the Spec, the secret in Key Vault is created with the name `<kubernetes-namespace-of-resource>-<ObjectMeta.name>`.
If secretname _is_ specified in the Spec, the actual secret name searched for is `<kubernetes-namespace-of-resource>-<spec.secretName>`.


Some things to note about the Key Vault you use with the operator:
1. The KeyVault should have an access policy added for the identity under which the Operator runs as.
   This access policy should include  at least `get`, `set`, `list` and `delete` Secret permissions.
2. You can use a Key Vault with "Soft delete" enabled. However, you cannot use a Key Vault with "Purge Protection" enabled, as this prevents the
   secrets from being deleted and causes issues if a resource with the same name is re-deployed.

## Per resource Key Vault

In addition to being able to specify an Azure Key Vault to store secrets, you also have the option to specify a different Key Vault per resource.

Some situations may require that you use a different Key Vault to store the admin password for the Azure SQL server from the Key Vault used to store the connection string for eventhubs.
You can specify the Key Vault name in the Spec field  `keyVaultToStoreSecrets`. When this is specified, the secrets produced when provisioning the resource in question will be stored
in the specified Key Vault instead of the global one configured for the operator.

## Format of secrets in Key Vault
In some scenarios it may be helpful to understand the format of the secrets stored in Key Vault. One such scenario is if you created a parent resource with a tool other than the operator,
but you now would like to manage child resources with the operator. A common example of this is a single persistent SQL server whose administrator account will be used to provision other databases
and users.

The secrets are all stored in Key Vault as a `secret` entity whose value is JSON serialized key-value pairs where the values have been base64 encoded.
For example, an Azure SQL Server created in namespace `default` with name `my-sql-server` will have a secret named `default-my-sql-server` that looks like:
```
{
  "username": "aGFzMTFzMnU=",
  "password": "XTdpMmQqNsd7YlpFdEApMw==",
  "fullyqualifiedusername": "aGFzMTUzMnVAc3Fsc2VydmVyLXNhbXBsZS04ODg=",
  "sqlservername": "c3Fsc2VyfmVyLXNhbXBsZS04ODg=",
  "fullyqualifiedservername": "c3Fsc2VydmVyLXNhbXBsZS04ODguZGF0YWJhc2Uud2luZG93cy5uZXQ="
}
```

For more details about what fields are in each secret, see the documentation for the resource in question, for example: [azuresql](../services/azuresql.md) and [mysql](../services/mysql.md)
