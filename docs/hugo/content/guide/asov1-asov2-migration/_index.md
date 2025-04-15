---
title: ASOv1 to ASOv2 migration guide
weight: 5
layout: single
cascade:
  - type: docs
  - render: always
---
ASOv1 has been [deleted](https://github.com/Azure/azure-service-operator/blob/main/README.md#aso-v1). 
We strongly recommend migrating to ASOv2 to continue getting new features and support.

## Assumptions

- Migrating a single cluster.
  - If you have ASOv1 installed on multiple clusters, apply these steps to each cluster.
- ASOv1 is installed in its default namespace `azureoperator-system`.
  - If it is installed into a different namespace, replace all occurrences of `azureoperator-system` in this document with 
    the namespace it's installed into. 
- ASOv2 is not yet installed in the cluster.
- ASOv1 is managing a single subscription, so ASOv2 will also be managing that same single subscription via a single 
  global credential.
  - See [ASOv2 credential scope](../authentication/#credential-scope) for other credential scopes that allow for 
    managing multiple subscriptions and/or using multiple identities.

## Performing the migration

### Check the version of cert-manager

Ensure the version of cert-manager in your cluster supports `v1` resources.

```
kubectl api-resources --api-group cert-manager.io`
```
should show `v1` resources.

{{% alert title="Warning" color="warning" %}}
Ensure that you've verified each ASOv1 resource you delete has the `skipreconcile=true` annotation before you delete it
in Kubernetes or else the deletion will propagate to Azure and delete the underlying Azure resource as well, which you do not want.
{{% /alert %}}

{{% alert title="Note" %}}
We strongly recommend ensuring that you're running the latest version of cert-manager
(1.14.4 at the time this article was written).
{{% /alert %}}

{{% alert title="Note" %}}
We also recommend ensuring that ASOv1 is configured to use `cert-manager.io/v1` resources.
You can run `helm upgrade` and pass `--set certManagerResourcesAPIVersion=cert-manager.io/v1` to ensure ASOv1 is using
the v1 versions of the cert-manager resources.
{{% /alert %}}

### Install ASOv2

Follow the [standard instructions](../../#installation). We recommend you use the same credentials as ASOv1 is currently using.

{{% alert title="Note" %}}
Make sure to follow the [guidelines](../crd-management/) for setting the `--crdPattern`. Configure ASOv2
to only manage the CRDs which you need.

For example: if you are using storage accounts, resource groups, redis cache's, cosmosdbs, eventhubs, and SQL Azure,
then you would configure
`--set crdPattern='resources.azure.com/*;cache.azure.com/*;documentdb.azure.com/*;eventhub.azure.com/*;sql.azure.com/*;storage.azure.com/*'`
{{% /alert %}}

### Stop ASOv1 reconciliation

Mark each ASOv1 resource with the `skipreconcile=true` annotation.
This annotation ensures that ASOv1 will no longer update or delete the resource in question in Azure.

To annotate a specific resource:
```
kubectl annotate -n <namespace> storageaccount <name> skipreconcile=true
```

To annotate all ASOv1 resources in a given namespace:
```
kubectl annotate $(kubectl api-resources -o name | grep azure.microsoft.com | paste -sd "," - | xargs kubectl get -A -o name) -n <namespace> skipreconcile=true`
```

Once you have annotated the resources, double-check that they all have the `skipreconcile` annotation. 

{{% alert title="Warning" color="warning" %}}
If you delete an ASOv1 resource that does not have this annotation it _**will**_ delete the underlying
Azure resource which may cause downtime or an outage.
{{% /alert %}}

### Use `asoctl` to import the resources from Azure into ASOv2

The commandline tool `asoctl` is a utility for working with ASO. Here we are going to use it to import your existing 
Azure resources into ASOv2.

Download the latest version of [asoctl](../../tools/asoctl/).

Export the resources you want to transfer using the `asoctl import azure-resource` command. `import azure-resource`
imports the referenced resource and all its child resources by default. This means if you want to import the resources 
from a specific resource group being managed by ASOv1 you can just refer to the resource group.

When running this command it's recommended you use at least asoctl v2.7.0 and pass the 
`-namespace` (or `-n`) and `-annotation` (or `-a`) arguments to import the resources to a particular namespace with 
the `serviceoperator.azure.com/reconcile-policy: skip` annotation.
```
asoctl import azure-resource /subscriptions/<subid>/resourceGroups/<rg> -o resources.yaml -namespace <namespace> -annotation serviceoperator.azure.com/reconcile-policy=skip
```

If you have a large number of resources you may want to also use the `--output-folder` option to render a single 
resource per file, rather than one file with all resources.

This should produce a file similar to the one shown here (this is a simplified file showing just a resource group and
a single resource for the sake of brevity. It's likely that your file is much larger):

```yaml
---
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-test
  namespace: ns1
  annotations:
    serviceoperator.azure.com/reconcile-policy: skip
spec:
  azureName: aso-test
  location: westus
---
apiVersion: storage.azure.com/v1api20230101
kind: StorageAccount
metadata:
  name: aso-test-storage1
  namespace: ns1
  annotations:
    serviceoperator.azure.com/reconcile-policy: skip
spec:
  accessTier: Hot
  allowBlobPublicAccess: false
  allowCrossTenantReplication: false
  azureName: asoteststorage1
  encryption:
    keySource: Microsoft.Storage
    services:
      blob:
        enabled: true
        keyType: Account
      file:
        enabled: true
        keyType: Account
  kind: StorageV2
  location: westus
  minimumTlsVersion: TLS1_0
  networkAcls:
    bypass: AzureServices
    defaultAction: Allow
  owner:
    name: aso-test
  sku:
    name: Standard_RAGRS
    tier: Standard
  supportsHttpsTrafficOnly: true
```

### Examine `resources.yaml` to ensure it has the resources you expect

Once you have `resources.yaml` locally, examine it to ensure that it has the resources you expect.

{{% alert title="Note" %}}
`asoctl` is importing the resources from Azure. This means it cannot preserve any Kubernetes labels/annotations
you have on your ASOv1 resources. If your resources need certain labels or annotations, add those to `resources.yaml`
manually or use the `--annotation` and `--label` arguments to `asoctl`.
{{% /alert %}}

{{% alert title="Note" %}}
By default `asoctl` imports everything into the `default` namespace. Before applying any YAML
created by `asoctl`, ensure that the namespaces for the resources are correct. You can do this by manually
modifying the YAML, using a tool such as Kustomize, or using the `--namespace` argument to `asoctl`.
{{% /alert %}} 

### Configure `resources.yaml` to export the secrets you need from Azure

These are secrets such as storage account keys which ASOv1 has automatically exported into a Kubernetes secret.

ASOv1 exports secrets from Azure according to the 
[Secret naming](https://github.com/Azure/azure-service-operator/blob/asov1/docs/v1/howto/secrets.md#secret-naming)
rules. ASOv1 always exports these secrets, there is no user configuration required on either the name of the secret or its values.
In contrast, ASOv2 requires the user to opt-in to secret export, via the `spec.operatorSpec.secrets` property.

{{% alert title="Note" %}}
ASOv2 also splits secrets into two categories: Secrets provided by you, and secrets generated
by Azure. Exporting secrets from Azure is done via the `spec.operatorSpec.secrets`, while supplying secrets to Azure
is done by passing the reference to a secret key/value, such as via the
[administratorLoginPassword field of Azure SQL](../../reference/sql/v1api20211101#sql.azure.com/v1api20211101.Server_Spec).
This means that there may be cases where a single secret in ASOv1 becomes 2 secrets in ASOv2, one for inputs and one for outputs.
{{% /alert %}}

You may need to configure `resouces.yaml` to export corresponding secrets. To determine if, for a given namespace, 
there are secrets being written by ASOv1 and consumed by your applications, check the following two things:

1. Are there secrets in the namespace owned by ASO? 
   - Use `kubectl get secrets -n ns1 -o json | jq -r '.items[] | select(.metadata.ownerReferences[]? | 
     select(.apiVersion | contains("azure.microsoft.com")))'`
     to find ASOv1 owned secrets. 
2. Do you have pods or services consuming those secrets? If they aren't being consumed anywhere, you don't 
   need them and they can be ignored/discarded.

Here's an example of secrets created by ASOv1 for various resources:
```sh
kubectl get secrets -n ns1
NAME                                         TYPE     DATA   AGE
azuresqlserver-azuresql-migration-sample-1   Opaque   5      106m
cosmosdb-cosmosdb-sample-1                   Opaque   10     6h42m
eventhub-eventhub-sample-1                   Opaque   7      101m
rediscache-rediscache-sample-1               Opaque   2      80m
storageaccount-cutoverteststorage1           Opaque   5      105m
```

Secrets owned by ASOv1 will have an `ownerReferences` section like this:
```
  ownerReferences:
  - apiVersion: azure.microsoft.com/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: Eventhub
    name: eventhub-sample-1
    uid: 0a034f50-3060-4114-98e6-b5085326da0d
```

Once you've identified the set of secrets which are exported by ASOv1 and which are being consumed by your applications,
configure ASOv2 to export similar secrets by using the `spec.operatorSpec.secrets` field. See [examples](#examples) for 
examples of various resource types.


{{% alert title="Note" %}}
It is strongly recommended that you export the ASOv2 secrets to a _different_ secret name than the
ASOv1 secret. ASOv2 will not allow you to overwrite an existing secret with `spec.operatorSpec.secrets`.
You'll get an error that looks like
"cannot overwrite Secret ns1/storageaccount-cutoverteststorage1 which is not owned by StorageAccount.storage.azure.com
ns1/aso-migration-test-cutoverteststorage1".
Instead of overwriting the same secret, create a different secret with ASOv2 (recommend including a suffix to
identify it such as -asov2). Then, when you're ready, swap your deployment to use the new ASOv2 secrets.
{{% /alert %}}

### Configure `resources.yaml` to source secret values from Kubernetes

Some Azure services require user-supplied secrets, for example: Azure SQL Server requires an administrator password.
ASOv1 generated these passwords for you, but in ASOv2 you control these passwords and their lifecycle (including rollover).

`asoctl` does not automatically include secret password fields in its output. Instead, you must supply a Kubernetes
secret containing the password which ASOv2 will supply to Azure.

The following ASOv1 resources automatically generated usernames and passwords which you will need to take ownership of
when migrating to ASOv2:
- Azure SQL Server
- Azure SQL User
- My SQL Server
- My SQL User
- PostgreSQL SQL Server
- PostgreSQL SQL User
- VM
- VMSS

The process for each of these resources is the same:
1. Find the ASOv1 managed secret which contains the current password.
2. Create a new secret containing the same password.
3. Update the ASOv2 `resources.yaml` to supply the new password in the correct location.

See [examples](#examples) for examples of how to do this with various resource types.

### Configure `resources.yaml` to have annotation [serviceoperator.azure.com/reconcile-policy: skip](../annotations/#serviceoperatorazurecomreconcile-policy)

This can be done automatically by `asoctl` during resource import with the 
`-annotation serviceoperator.azure.com/reconcile-policy=skip` argument or manually by modifying `resources.yaml` 
afterward.

{{% alert title="Warning" color="warning" %}}
Be _very_ careful issuing deletions of ASOv2 resources once you've imported them if they do not have
the `serviceoperator.azure.com/reconcile-policy=skip` annotation. Just like ASOv1,
ASOv2 _**will**_ delete the underlying Azure resource by default!
It is _strongly_ recommended that you use asoctl's `-annotation serviceoperator.azure.com/reconcile-policy=skip` flag
and only remove that annotation a few resources at a time to ensure things are working as you expect.
{{% /alert %}}

### Apply `resources.yaml` to your cluster

This will create the ASOv2 resources. At this point you should be in a state where:
- ASOv1 is not managing the resources because they have all been annotated with `skipreconcile=true`.
- ASOv2 is not managing the resources because they have all been annotated with `serviceoperator.azure.com/reconcile-policy=skip`.

Check that the ASOv2 resources are in the correct namespaces and that you're seeing the secrets and configmaps you expect 
(these will be exported by ASOv2 even when annotated with `serviceoperator.azure.com/reconcile-policy=skip`).

Remove the `serviceoperator.azure.com/reconcile-policy=skip` from the ASOv2 resources (one at a time if desired) and swap existing 
deployments to source data from the ASOv2 secrets/configmaps.

## Rolling back

If at any point something isn't working, you can roll back by marking the ASOv2 resource with the 
`serviceoperator.azure.com/reconcile-policy=skip` annotation and removing the `skipreconcile=true` annotation
from the ASOv1 resource. Don't forget to swap your deployments to use the secrets from ASOv1 if they had been updated 
to rely on the ASOv2 secrets.

## Cleaning up

Once you've migrated the ASOv1 resources to ASOv2 and been running successfully for a while, you can delete the ASOv1 resources
in Kubernetes with `kubectl delete`.

{{% alert title="Warning" color="warning" %}}
Make sure that each ASOv1 resource you delete has the `skipreconcile=true` annotation before you delete it
in Kubernetes or else the deletion will propagate to Azure and delete the underlying Azure resource as well, which you do not want.
{{% /alert %}}

## Examples

- [Storage](./storage/)
- [Storage Container](./storagecontainer/)
- [EventHub](./eventhub/)
- [Azure SQL](./azuresql/)
- [Redis](./redis/)
- [MySQL](./mysql/)
- [PostgreSQL](./postgresql/)
- [Database users](./databaseusers/)

## Gotchas

- `resourcegroup` is the Kubernetes "short name" for both ASOv1 and ASOv2 resource groups. This means that running
  `kubectl get resourcegroups` is ambiguous. Kubernetes will use ASOv1's `resourcegroup` asuming it was installed first.
   If ASOv1 and ASOv2 are installed into the same cluster, we recommend using the fully specified name for querying 
   resourcegroups from ASOv2:  `kubectl get resourcegroup.resources.azure.com`.
- ASOv2 does not support keyVault secrets.
- ASOv2 does not support MySQL single server or PostgreSQL single server, as these are being deprecated. 
  See [MySQL](./mysql/) and [PostgreSQL](./postgresql/) for more details.

## Helpful commands

View all ASOv1 resources in the cluster:
```sh
kubectl api-resources -o name | grep azure.microsoft.com | paste -sd "," - | xargs kubectl get -A
```

View all ASOv2 resources in the cluster:
```sh
kubectl api-resources -o name | grep azure.com | paste -sd "," - | xargs kubectl get -A
```

Annotate a single ASOv1 resource to skip all reconciliation:
```sh
kubectl annotate -n ns1 storageaccount <name> skipreconcile=true
```

Annotate existing ASOv1 resources in namespace `ns1` to skip all reconciliation:
```sh
kubectl annotate $(kubectl api-resources -o name | grep azure.microsoft.com | paste -sd "," - | xargs kubectl get -A -o name) -n ns1 skipreconcile=true
```

Find secrets owned/created by ASOv1:
```sh
kubectl get secrets -n ns1 -o json | jq -r '.items[] | select(.metadata.ownerReferences[]? | select(.apiVersion | contains("azure.microsoft.com")))'
```

Annotate existing ASOv2 resources in namespace `ns1` to skip all reconciliation:
```sh
kubectl annotate $(kubectl api-resources -o name | grep azure.com | paste -sd "," - | xargs kubectl get -A -o name) -n ns1 serviceoperator.azure.com/reconcile-policy=skip
```

Clear the `serviceoperator.azure.com/reconcile-policy=skip` annotation:
```sh
kubectl annotate $(kubectl api-resources -o name | grep azure.com | paste -sd "," - | xargs kubectl get -A -o name) -n ns1 serviceoperator.azure.com/reconcile-policy-
```
