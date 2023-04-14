---
title: "v2.0.0 Breaking Changes"
linkTitle: "v2.0.0"
weight: 90
---


## Upgrades from releases prior to v2.0.0-beta.5 are disallowed

We changed how we manage CRDs in this release (see [PR#2769](https://github.com/Azure/azure-service-operator/pull/2769)), and as a result if using Helm you _must_ upgrade from v2.0.0-beta.5 to v2.0.0.

You cannot upgrade from v2.0.0-beta.4 or earlier directly to v2.0.0. This is enforced with a Helm upgrade hook.

This restriction is just for upgrades to the v2.0.0 version, although [we always recommend](https://azure.github.io/azure-service-operator/introduction/upgrading/#recommended-upgrade-pattern) upgrading one version at a time.

## Alpha CRD versions have been removed

**You cannot successfully upgrade to v2.0.0 until you have followed our [migration guide](TODO)**.

Fresh installations of v2.0.0 are unaffected.

## ResourceGroup Status.ProvisioningState field is now Status.Properties.ProvisioningState

We believe that this is unlikely to break users as tooling always uses the `Conditions` field rather than `ProvisioningState` to track resource provisioning progress, but calling it out nonetheless for completeness.

