---
title: "v2.0.0 Breaking Changes"
linkTitle: "v2.0.0"
weight: 90
---

## Upgrades from releases prior to v2.0.0-beta.5 are disallowed

We changed how we manage CRDs in this release (see [PR#2769](https://github.com/Azure/azure-service-operator/pull/2769)), and as a result if using Helm you _must_ upgrade from v2.0.0-beta.5 to v2.0.0.

You cannot upgrade from v2.0.0-beta.4 or earlier directly to v2.0.0. This is enforced with a Helm upgrade hook.

This restriction is just for upgrades to the v2.0.0 version, although [we always recommend](https://azure.github.io/azure-service-operator/guide/upgrading/#recommended-upgrade-pattern) upgrading one version at a time.

## Alpha CRD versions have been removed

**You cannot successfully upgrade to v2.0.0 until you have followed the migration steps below:**

1. If you have never installed an `alpha` version of ASOv2, or used an `alpha` version of a CRD, no action is required.
   If you're not sure, run the below steps anyway. They will not do anything if alpha was never used.
2. Ensure the cluster is running the `v2.0.0-beta.5` version of ASO.
3. Install the `asoctl` tool by following the [installation instructions](../../../tools/asoctl/#installation)
4. Run `asoctl clean crds`. This will migrate your resources to a newer version if needed and remove the `alpha` CRD 
   versions from the CRD `Status.StoredVersions` collection. For more details on this command see the 
   [asoctl clean crds documentation](../../../tools/asoctl/#clean-crds)

Once you have successfully executed the above steps, you can upgrade to `v2.0.0`. When upgrading to `v2.0.0` the following
error means that you need to run the above steps:

```
"msg"="failed to apply CRDs" "error"="failed to apply CRD storageaccountsqueueservicesqueues.storage.azure.com: CustomResourceDefinition.apiextensions.k8s.io \"storageaccountsqueueservicesqueues.storage.azure.com\" is invalid: status.storedVersions[0]: Invalid value: \"v1alpha1api20210401storage\": must appear in spec.versions" 
```

## ResourceGroup Status.ProvisioningState field is now Status.Properties.ProvisioningState

We believe that this is unlikely to break users as tooling always uses the `Conditions` field rather than `ProvisioningState` to track resource provisioning progress, but calling it out nonetheless for completeness.
