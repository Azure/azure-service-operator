---
title: Clarifying object structure
---
# Clarifying object structure

Today we have resources that look like:
```yaml
apiVersion: microsoft.storage.infra.azure.com/v1alpha1api20190401
kind: StorageAccount
metadata:
  name: samplekubestorage
  namespace: default
spec:
  azureName: mykubestorage
  location: westcentralus
  kind: BlobStorage
  sku:
    name: Standard_LRS
  owner:
    name: k8sinfra-sample-rg
  accessTier: Hot
  tags:
    tag1: tag1
    tag2: tag2
```

## The problem
There's no good way with this object structure to differentiate stuff that is for Azure directly, versus stuff that is
for the operator. `owner` _almost_ falls into this category already, but there are other likely upcoming properties that
definitely fall into this category: 
 - `SecretConfiguration`: details about where/how we should store secrets created by this entity.
 - `Credentials`: per object credentials used to support CAPZ-like scenarios.

The problem also manifests in `Status` where ideally we would distinguish between properties from Azure directly (the
result of a GET on the resource) and properties that we are presenting. For example, we may want to have a status field
for `deploymentId` documenting the deployment ID used to create the resource, or the `error` if there was an error.
If we do that there's no easy way for the customer to understand that field is provided by ASO and not by Storage.

## Proposal
We introduce an additional level of hierarchy specifically to clarify what is coming or going directly from or to Azure.
This is similar to what Crossplane does (see their 
[MySQLServer](https://github.com/crossplane/provider-azure/blob/master/examples/database/mysqlserver.yaml)). Unlike Crossplane though
we will push the operator specific properties down a level and leave the Azure properties at the top level.

On the `spec` we could call this something like `forOperator` or `operator`, and on the `status` `fromOperator` or `operator`.

Our structure would then look like:
```yaml
apiVersion: microsoft.storage.infra.azure.com/v1alpha1api20190401
kind: StorageAccount
metadata:
  name: samplekubestorage
  namespace: default
spec:
  azureName: mykubestorage
  location: westcentralus
  kind: BlobStorage
  sku:
    name: Standard_LRS
  accessTier: Hot
  tags:
  tag1: tag1
  tag2: tag2
  owner:
      name: k8sinfra-sample-rg
  operatorSpec:
    credentials:
      credentials
```

Similarly, a `status` might look like:
```yaml
status:
  id: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/mykubestorage
  location: westcentralus
  kind: BlobStorage
  sku:
    name: Standard_LRS
  accessTier: Hot
  tags:
    tag1: tag1
    tag2: tag2
  operatorStatus:
    deploymentId: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.Deployments/deployments/1234
    goalSeekingState: GoalMet
```

# FAQ

Q: What about the `owner` field?
A: It is such an important field and cuts across both Azure (since it's partially about resource relationships in Azure)
   and Kubernetes (since it's partially about resource relationships in Kubernetes) so it makes sense to leave it at 
   the top level.

Q: What about the `azureName` field?
A: It is for Azure, so stays at the top level.

Q: Do we actually have any properties that need to move TODAY because of this change?
A: No, because `state` and `deploymentId` are currently resource annotations, not fields in the `status`, and on the `spec` the only
   fields which might belong in `operatorSpec` are `owner` and `azureName` which as per the above FAQ are staying at the top level.

# Open questions

1. What should we call the property which we push things down "into"? Here are some options:
  * `operator` (for both)
  * `operatorData` (for both)
  * `forOperator` / `fromOperator`
  * `operatorSpec` / `operatorStatus`
  * ??? / `operatorState`
