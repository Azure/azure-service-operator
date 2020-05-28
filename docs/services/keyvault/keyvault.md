# Key Vault Operator

## Resources Supported

The Azure Key Vault operator suite consists of the following operators:

- KeyVault - Deploys an Azure Key Vault given the location and resource group
- KeyVaultKey - Deploys an Azure Key Vault key given the location and resource group

### KeyVault

You can find a sample YAML for KeyVault [here](/config/samples/azure_v1alpha1_keyvault_simple.yaml)

The value for Kind, `KeyVault` is the Custom Resource Definition (CRD) name.
`Name` is the name of the KeyVault resource that will be created.

The values under `spec` provide the values for a location where you want to create the Key Vault in, and the Resource Group in which you want to create it under.

You can also configure `Access Policies` and `Network policies` for the KeyVault as part of the Spec. A sample YAML with these settings can be found [here](/config/samples/azure_v1alpha1_keyvault.yaml)

#### Access Policies

Key Vault access is governed via a management plane and a data plane.  The management plane is where you manage Key Vault resources, and the data plane is where you work with data stored in a Key Vault.  

In order to access a Key Vault in either plane, all callers must have proper authentication and authorization.  Both planes use Azure Active Directory (AAD) for authentication and authorization.  For authorization, the management plane uses role-based access control (RBAC), and the data plane uses Access Policies.

Access control for the two planes work independently.  Therefore you will notice that Azure Service Operators account for data plane access via Access Policy administration.  Notice the `accessPolicies` yaml entries in the sample above.  These are critical entries for properly securing and authorizing Key Vault data plane access.  Access Control for the management plane is administered using RBAC that is also controlled by AAD.  Therefore, authorization of the management plane is best achieved using Azure Powershell, the Azure CLI, and the Azure Portal.

#### Network Policies

KeyVault access can be restricted to only certain virtual networks or IP ranges. You can choose to allow access from all Azure services too. Network policies in the Keyvault spec allow you to configure these.

### KeyVaultKey Operator

The KeyVaultKey operator serves as an operator that allows for declarative management of Key Vault keys - one of the three resources available for storage and management in Key Vault; keys, secrets, and certificates.  Keys can be leveraged for various use cases.

For example, one of the most common use cases for Key Vault keys is [SQL Server Always Encrypted](https://docs.microsoft.com/sql/relational-databases/security/encryption/always-encrypted-database-engine) data encryption.  The KeyVaultKey operator will allow for generation of an key that can be used to leverage data encryption via managed keys in Key Vault:

```sql
ALTER COLUMN ENCRYPTION KEY key_name
    [ ADD | DROP ] VALUE
    (
        COLUMN_MASTER_KEY = column_master_key_name
        [ , ALGORITHM = 'algorithm_name' , ENCRYPTED_VALUE =  varbinary_literal ]
    ) [;]
```

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
