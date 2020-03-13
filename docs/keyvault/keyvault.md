# Key Vault Operator

## Resources Supported

The Azure Key Vault operator suite consists of the following operators:

-  KeyVault - Deploys an Azure Key Vault given the location and resource group
-  KeyVaultKey - Deploys an Azure Key Vault key given the location and resource group

## Deploying Key Vault Resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

### KeyVault

For instance, this is the sample YAML for the Azure SQL server:

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: KeyVault
metadata:
  name: keyvaultsample123
  labels: # Provide tags to add to the Key Vault as labels
    tag1: value1
    tag2: value2
spec:
  resourceGroup: resourcegroup-azure-operators
  location: westus
  enableSoftDelete: false
  networkPolicies: 
    bypass: AzureServices # AzureServices or None
    defaultAction: Allow # Allow or Deny
    ipRules: 
      - 172.16.0.0/24
      - 172.16.1.0/24
    virtualNetworkRules:
      - /subscriptions/<subid>/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1
      - /subscriptions/<subid>/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet2
  accessPolicies:
    - tenantID: <tenantID>
      clientID: <clientID>
      applicationID: <appID>
      permissions:
        keys: # backup create decrypt delete encrypt get import list purge record restore sign unwrapKey update verify wrapKey
          - list
          - get
        secrets: # backup delete get list purge recover restore set
          - list
          - get 
        certificates: # backup create delete deleteissuers get getissuers import list listissuers managecontacts manageissuers purge recover restore setissuers update 
          - list
          - get
        storage: # backup delete deleteas get getas list listas purge recovver regenratekey restore set setas update
          - list
          - get
  ```

The value for kind, `KeyVault` is the Custom Resource Definition (CRD) name.
`keyvaultsample123` is the name of the SQL server resource that will be created.

The values under `spec` provide the values for a location where you want to create the Key Vault to be created, and the Resource Group in which you want to create it under.

Once you've updated the YAML with the settings you need, and you have the operator running, you can create a Key Vault resource using the command:

```bash
kubectl apply -f config/samples/azure_v1alpha1_keyvault.yaml
```

### Access Policies

Key Vault access is governed via a management plane and a data plane.  The management plane is where you manage Key Vault resources, and the data plane is where you work with data stored in a Key Vault.  

In order to access a Key Vault in either plane, all callers must have proper authentication and authorization.  Both planes use Azure Active Directory (AAD) for authentication and authorization.  For authorization, the management plane uses role-based access control (RBAC), and the data plane uses Access Policies.

Access control for the two planes work independently.  Therefore you will notice that Azure Service Operators account for data plane access via Access Policy administration.  Notice the `accessPolicies` yaml entries in the sample above.  These are critical entries for properly securing and authorizing Key Vault data plane access.  Access Control for the management plane is administered using RBAC that is also controlled by AAD.  Therefore, authorization of the management plane is best achieved using Azure Powershell, the Azure CLI, and the Azure Portal.

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

## View and Troubleshoot Key Vault Resources

You can view your created Key Vault resources using the steps [here](viewresources.md).

## Delete a Key Vault Resource

To delete an existing resource from Kubernetes and Azure, use the following command:

```shell
kubectl delete <Kind> <instancename>
```

For instance, deleting the above Key Vault instance would look like this.

```shell
kubectl delete KeyVaultKey vaultsample123
```

The following message should appear:

`keyvault.azure.microsoft.com keyvaultsample123 deleted.`