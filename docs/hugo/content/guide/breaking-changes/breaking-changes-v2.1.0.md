---
title: "v2.1.0 Breaking Changes"
linkTitle: "v2.1.0"
weight: 80
---

## The operator no longer installs CRDs by default

**Action required:** When installing ASO for the first time, you must now specify `crdPattern` (for Helm) or `--crd-patterns` 
(in operator pod cmdline for raw YAML) to select the subset of CRDs you would like to install.

When upgrading ASO, existing CRDs will be automatically updated to the new version but new CRDs added in that release 
will not automatically be installed. 
This means that when upgrading the operator, if you don't want to use any CRDs newly added in that release you don't 
need to do anything.

**Action required:** When upgrading ASO, if you want to install new CRDs (for example CRDs just added in the version of 
ASO you are upgrading to) you must specify `crdPattern` (Helm) or `--crd-patterns` (YAML) to install the CRDs. 
For example: if you do want to use a newly added CRD (such as `network.azure.com/bastionHosts` mentioned
below), you would need to specify `crdPatterns=network.azure.com/*` when performing the upgrade.

See [CRD management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for more details 
about this change and why it was made.

## `serviceoperator.azure.com/credential-from` no longer supports cross namespace secret references

This was never documented as supported but worked unintentionally. The feature now works as it was always documented: 
allowing references to secrets only if the secret is in the same namespace as the resource itself.

This was a security issue which we had to close.

See [#2919](https://github.com/Azure/azure-service-operator/pull/2919)

## ManagedClusters ManagedClusterServicePrincipalProfile.Secret field is now marked as a SecretReference

We have marked `ManagedClusterServicePrincipalProfile.Secret` property to be a secret reference. If `ContainerService/ManagedClusters` resource is installed in your cluster, 
below steps should be performed before upgrading. If you have `ManagedClusterServicePrincipalProfile.Secret` property set on your `ManagedCluster` resource, follow the steps below:

1. Annotate the resource with `serviceoperator.azure.com/reconcile-policy: skip` to prevent ASO from trying to reconcile the resource while you are upgrading.
2. Download the current YAML for the resource using `kubectl` if you don't have it elsewhere.
4. Create a kubernetes secret containing the value for `ManagedClusterServicePrincipalProfile.Secret`.
5. Edit downloaded YAML in step 3, and add a secret key and name reference. Example [here](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/compute/v1api/v1api20201201_virtualmachine.yaml#L18).
6. Delete the resource from your cluster using `kubectl delete`. Your Azure resource will be left untouched because of the `reconcile-policy` annotation you added above.
7. [Upgrade ASO](../../upgrading) in your cluster.
8. Apply the updated YAML to your cluster using `kubectl apply`. If any errors occur, address them.
9. If the `reconcile-policy` annotation is still present, remove it from the resource.
