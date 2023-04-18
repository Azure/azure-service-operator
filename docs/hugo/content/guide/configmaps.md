---
title: Sharing data through ConfigMaps
---

Resources supported by Azure Service Operator may support dynamic input in the form of reading a `ConfigMap`.
They may also support storing certain output into a `ConfigMap` (such as ManagedIdentity `principalId` and `clientId`). 

## How to provide ConfigMap data to ASO

Resources may have fields in their `spec` that allow a reference to a Kubernetes `ConfigMap`.
In general, these fields are usually optional, and you can choose to specify the raw value _or_ source it from a `ConfigMap` for something more dynamic.

For example, in order to create a `RoleAssignment`, you must specify the `principalId` of the identity it applies to. 
You can hardcode that `principalId` or you can refer to it dynamically as a `ConfigMap` key.

**Example (from [the RoleAssignment sample](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/authorization/v1beta/v1beta20200801preview_roleassignment.yaml)):**
```yaml
apiVersion: authorization.azure.com/v1beta20200801preview
kind: RoleAssignment
metadata:
  name: fee8b6b1-fe6e-481d-8330-0e950e4e6b86 # Should be UUID
  namespace: default
spec:
  location: westcentralus
  # This resource can be owner by any resource. In this example we've chosen a resource group for simplicity
  owner:
    name: aso-sample-rg
    group: resources.azure.com
    kind: ResourceGroup
  # This is the Principal ID of the AAD identity to which the role will be assigned
  principalIdFromConfig:
    name: identity-settings
    key: principalId
  roleDefinitionReference:
    # This ARM ID represents "Contributor" - you can read about other built in roles here: https://docs.microsoft.com/en-us/azure/role-based-access-control/built-in-roles
    armId: /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c
```

This reads the `principalId` value from a `ConfigMap`. Alternatively it could have been specified explicitly via the 
`spec.principalId` field instead of using `spec.principalIdFromConfig`.

Note that not every field for every resource supports importing data from `ConfigMap`. To check if a resource supports it,
consult its [documentation](../../reference/authorization/v1beta20200801preview#authorization.azure.com/v1beta20200801preview.RoleAssignment).

## How to export ConfigMap data from ASO

Some resources support saving data into a `ConfigMap`. The individual properties can be exported to a `ConfigMap` of your choosing by
configuring the `.spec.operatorSpec.configMaps` field. The data will be written to the destination(s) you specify once the resource has 
successfully been provisioned in Azure.
The resource will not move to [Condition](../conditions) `Ready=True` 
until the data has been written.

**Example (from [the UserAssignedIdentity sample](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/managedidentity/v1beta/v1beta20181130_userassignedidentity.yaml)):**
```yaml
apiVersion: managedidentity.azure.com/v1beta20181130
kind: UserAssignedIdentity
metadata:
  name: sampleuserassignedidentity
  namespace: default
spec:
  location: westcentralus
  owner:
    name: aso-sample-rg
  operatorSpec:
    configMaps:
      # Export the principalId and clientId to a ConfigMap for use by our application and/or
      # other ASO resources such as RoleAssignments
      principalId:
        name: identity-settings
        key: principalId
      clientId:
        name: identity-settings
        key: clientId
```

