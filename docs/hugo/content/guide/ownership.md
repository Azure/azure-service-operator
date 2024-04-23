---
title: Ownership
linktitle: Ownership
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

## What is an owner?

Most Azure Service Operator resources have an `owner` field, like so:
```
  owner:
    name: aso-sample-rg
```

The Go type for this is either 
[KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)
or 
[ArbitraryOwnerReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ArbitraryOwnerReference), depending
on if the resource supports only a single kind of owner (`StorageAccount`s are always owned by `ResourceGroup`s),
or many kinds of owner (`RoleAssignment` can be owned by any type of resource).

You can tell what type of resource the owner should point to by looking at the reference documentation. For example:

[StorageAccount owner](https://azure.github.io/azure-service-operator/reference/storage/v1api20230101/#storage.azure.com/v1api20230101.StorageAccount)
> Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource

[BlobService owner](https://azure.github.io/azure-service-operator/reference/storage/v1api20230101/#storage.azure.com/v1api20230101.StorageAccounts_BlobService_Spec)
> Owner is expected to be a reference to a storage.azure.com/StorageAccount resource

[RoleAssignment owner](https://azure.github.io/azure-service-operator/reference/authorization/v1api20220401/#authorization.azure.com/v1api20220401.RoleAssignment)
> This resource is an extension resource, which means that any other Azure resource can be its owner.

## What is ownership for?

Ownership tells ASO two things:

### What "path" is this resource at? 

For example: which `ResourceGroup` should this `StorageAccount` be deployed into?
The owner represents the bolded part of this storage account ARM ID: 
> **/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}**/providers/Microsoft.Storage/storageAccounts/mystorageaccount

and

> **/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}**/agentPools/myagentpool

for an AKS `AgentPool`.

### Which Kubernetes resource owns the resource?

This is for [garbage collection](https://kubernetes.io/docs/concepts/architecture/garbage-collection).

{{% alert title="Note" %}}
Unlike BICEP/ARM templates, which use two fields `parent` and `scope`, ASO uses the same `owner` property regardless of 
if the resource is a standard resource such as a `StorageAccount` with only one type of owner (a `ResourceGroup`), 
_or_ if the resource is an extension resource such as `RoleAssignment`. 

This is because the actual semantics of `parent` and `scope` are identical, each defines the URI path of the
resource and what Kubernetes resource owns it for garbage collection. 
{{% /alert %}}

## Using ARM ID for ownership

The `owner` field supports pointing to an ARM ID rather than the Kubernetes resource. This mechanism of ownership
exists to make using ASO alongside other resource management tools easier.

```
owner:
  armId: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg
```

This option should only be used if ASO is _not_ managing the owning resource. If the resource is being managed by ASO then
always prefer the `owner.name` option.

Disadvantages of using `owner.armId`:

- No automatic Kubernetes garbage collection if parent resource is deleted. Instead, resources whose owner.armId points 
  to a non-existent Azure resource will report an error in their ready condition.
- Makes it more difficult to deploy the same ASO resources into multiple namespaces and have them end up in multiple 
  subscriptions, as now the subscriptionId is in some `owner.armId` fields and will need to be templated using 
  something like Kustomize or Helm.

## Other reading

- [Type references and ownership design document]({{< relref "type-references-and-ownership" >}}).
