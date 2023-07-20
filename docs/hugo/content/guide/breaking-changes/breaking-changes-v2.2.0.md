---
title: "v2.2.0 Breaking Changes"
linkTitle: "v2.2.0"
weight: 80
---

## ManagedClusters ManagedClusterServicePrincipalProfile.Secret field is now marked as a SecretReference
  We always try to avoid breaking changes, but in this case, allowing raw passwords in the spec is a security problem and as such we've
  decided to make a break to correct this issue.

**Action required:** If the `ContainerService/ManagedClusters` resource is used in your cluster and the `ManagedClusterServicePrincipalProfile.Secret` property is set, do the following before upgrading ASO:

1. Annotate the resource with `serviceoperator.azure.com/reconcile-policy: skip` to prevent ASO from trying to reconcile the resource while you are upgrading.
2. Download the current YAML for the resource using `kubectl` if you don't have it elsewhere.
3. Create a kubernetes secret containing the value for `ManagedClusterServicePrincipalProfile.Secret`.
4. Edit downloaded YAML in step 2, and add a secret key and name reference. Example [here](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/compute/v1api/v1api20201201_virtualmachine.yaml#L18).
5. Delete the resource from your cluster using `kubectl delete`. Your Azure resource will be left untouched because of the `reconcile-policy` annotation you added above.
6. [Upgrade ASO]( {{< relref "upgrading" >}} ) in your cluster.
7. Apply the updated YAML to your cluster using `kubectl apply`. If any errors occur, address them.
8. If the `reconcile-policy` annotation is still present, remove it from the resource.

## Removed un-used Status properties

The below fields are never returned from the service and end up being an empty string always. The changes here do not affect the users, hence **no action is required**. 
* MachineLearningServices:
  * UserAccountCredentials_STATUS.AdminUserPassword
  * UserAccountCredentials_STATUS.AdminUserSshPublicKey
  * VirtualMachineSshCredentials_STATUS.Password

* Synapse:
    * Workspace_STATUS.SqlAdministratorLoginPassword
