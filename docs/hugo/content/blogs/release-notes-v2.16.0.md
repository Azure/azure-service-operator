---
title: "ASO v2.16 Release Notes"
date: 2025-10-29
description: "Release notes for Azure Service Operator v2.16.0"
---

We're excited to announce the release of Azure Service Operator v2.16.0! This release comes with a host of improvements and fixes, along with support being added for some oft-requested resources. 

## ⚠️ Breaking changes

This release includes breaking changes. Please review the [breaking changes documentation](https://azure.github.io/azure-service-operator/guide/breaking-changes/) before upgrading.

## 🎉 New and improved resource support

### New resource support

We've added support for three new resources: 

- [Azure Compute Capacity Reservation Groups](https://github.com/Azure/azure-service-operator/pull/4980) allowing capacity reservation manifests to be declaratively included alongside VM manifests 
   - _Special thanks to [bingikarthik](https://github.com/bingikarthik) for his contribution!_
- [Managing Azure quotas](https://github.com/Azure/azure-service-operator/pull/4979) for Kubernetes native management of Azure quotas
   - _Special thanks to [bingikarthik](https://github.com/bingikarthik) for his contribution!_
- [NetworkWatchers/flowLogs](https://github.com/Azure/azure-service-operator/issues/4614) for management of Network Security Group Flow Logs.

### Enhanced resource references

We've enhanced existing resource references with additional functionality:

- You can now specify Role Assignments using [well-known role defintion names](https://github.com/Azure/azure-service-operator/pull/4923) instead of GUIDs.
- UserAssignedIdentity now supports [config map](https://github.com/Azure/azure-service-operator/pull/4940).
- A system-managed identity can now be [specified in the app resources' identity references](https://github.com/Azure/azure-service-operator/pull/4924).
- Selected ResourceReferences now support [WellKnown names](https://github.com/Azure/azure-service-operator/pull/4922).
- Fixed 2 broken [ResourceReferenceProperties](https://github.com/Azure/azure-service-operator/pull/4925), `UserAssignedIdentityReference` and `KeyVaultArmReference`, which were previously seen as just strings but are now recognized as proper references.

## 🐛 Bug fixes

- Fixed a [high-priority regression](https://github.com/Azure/azure-service-operator/pull/4966) from ASO v2.15.0 where resources failed with the `"/v1, Kind=Secret is not cached"` error during resource claims.
- Kusto database reconciliation should [no longer be attempted](https://github.com/Azure/azure-service-operator/pull/4976) while a cluster is stopped. 
- `NetcfgSubnetRangeOutsideVnet` error is [now retryable](https://github.com/Azure/azure-service-operator/pull/4931) for VirtualNetworksSubnet to avoid the resources becoming stuck in certain scenarios.
- Added a missing parameter for [default node pool](https://github.com/Azure/azure-service-operator/issues/4942) in Node Auto-Provisioning.
- Fixed [select annotation changed predicate](https://github.com/Azure/azure-service-operator/pull/4967) from mistakenly classifying creates as containing annotation changes.
- Fixed `asoctl` import failures for [DataCollectionRule in PostgreSQL tables](https://github.com/Azure/azure-service-operator/issues/4919).

## 🔧 Infrastructure and technical improvements

### Code quality and tooling
- Updated Azure REST API specs submodule and generated code. Completed in [#4941](https://github.com/Azure/azure-service-operator/pull/4941).
- Stop generating explict local variables for loop aliasing effects. Completed in [#4949](https://github.com/Azure/azure-service-operator/pull/4949).
- Improved diagnostics from TypeTransformers. Completed in [#4937](https://github.com/Azure/azure-service-operator/pull/4937).
- Improved documentation comments in `arm` packages. Completed in [#4914](https://github.com/Azure/azure-service-operator/pull/4914).
- Fixed linter issues from newly published linter. Completed in [#4916](https://github.com/Azure/azure-service-operator/pull/4916).
- Deprecated `$export` in favor of `$exportAs` in generator configuration. Completed in [#4879](https://github.com/Azure/azure-service-operator/pull/4879).
- Added "UnsupportedResourceType" to asoctl exclusions for extension resources. Completed in [#4934](https://github.com/Azure/azure-service-operator/pull/4934).

## 🙏 Thank You

Thank you to the community for the continued engagement in ASO! From opening issues or requests, to submitting PRs of your own, the community helps ASO keep going! 
