---
title: "ASO v2.18 Release Notes"
date: 2026-02-18
description: "Release notes for Azure Service Operator v2.18.0"
type: blog
---

We're excited to announce the release of Azure Service Operator v2.18.0! This release adds new AKS preview API support, expands simplified versioning to storage resources, strengthens reconciliation checks, fixes Helm configurability issues, and updates a broad set of dependencies and tooling.

## 🎇 Headline Features

The `storage` group now supports [simplified API versioning](https://github.com/Azure/azure-service-operator/pull/5116). You can use all storage resources with the new `v` prefix, while the legacy `v1api` versions remain available for backward compatibility. This continues our transition toward simplified versioning across all resource groups (see [#4831](https://github.com/Azure/azure-service-operator/issues/4831) for more detail).

ASO now issues a fresh `GET` on the owner ARM resource before calling `PreReconcileOwnerCheck`, ensuring the check always operates on the latest owner status ([#5140](https://github.com/Azure/azure-service-operator/pull/5140)). This should result in faster goal state convergence for complex resource graphs.

Support for the `allowMultiEnvManagement` configuration flag allows per-namespace/per-resource Azure cloud environment settings, allowing users to manage resources across different Azure clouds/environments.

- Special thanks to [subhamrajvanshi](https://github.com/shubhamrajvanshi) for this contribution.

ASO no longer gives up [retrying on common Azure errors](https://github.com/Azure/azure-service-operator/pull/5092), improving reliability for long-running operations in ASO. ASO will currently still give up retrying on some well-known non-retryable errors.

Before issuing a DELETE request to Azure, [ASO may now check](https://github.com/Azure/azure-service-operator/pull/5040) if the resource or the resource's parent has already been removed. We will progressively enable this for all resources over the next few releases.

## ⚠️ Breaking changes

This release includes breaking changes. Please review [breaking changes in v2.18](https://azure.github.io/azure-service-operator/guide/breaking-changes/breaking-changes-v2.18.0/) before upgrading.

- The `containerservice` ManagedCluster and AgentPool API versions `v1api20230201` and `v1api20231001` have been removed, as [previously announced](https://github.com/Azure/azure-service-operator/releases/tag/v2.17.0).

- The API version `2024-01-01-preview` of `insights` has been removed as it was deprecated and removed upstream by Azure. Users of this API version should move to a supported version before upgrading ASO.

In [ASO v2.19](https://github.com/Azure/azure-service-operator/milestone/38), we will remove `containerservice` ManagedCluster and AgentPool API version `v1api20240402preview`.

## 🎉 New and improved resource support

You can now use the latest AKS preview features through [containerservice v20251002preview](https://github.com/Azure/azure-service-operator/pull/5170).

Several `insights` resources now use their latest ARM API versions ([#5136](https://github.com/Azure/azure-service-operator/pull/5136)):

- **DataCollectionEndpoint**, **DataCollectionRule**, and **DataCollectionRuleAssociation** now target API version `2024-03-11`
- **ScheduledQueryRule** now targets `2025-01-01-preview`

New [AppConfiguration child resources](https://github.com/Azure/azure-service-operator/pull/4948) using the latest 2024-06-01 API version have been added to address the requests for KeyValue management capabilities. These resources include `KeyValue`, `Replica`, and `Snapshot`.

The `Version` value for `dbformysql.FlexibleServer` [has been changed](https://github.com/Azure/azure-service-operator/pull/5195) to a simple string to allow new server versions to be used without the need for an API upgrade.

## 🐛 Bug fixes

- [The Helm chart metrics port is now fully configurable](https://github.com/Azure/azure-service-operator/pull/5156). Previously, the `metrics.port` value in `values.yaml` was ignored in favor of a hardcoded `8443`.
  - _Special thanks to [bingikarthik](https://github.com/bingikarthik) for the contribution!_
  -
- 4 known cases where the ARM ID properties were released [without being properly tagged-as/converted-to](https://github.com/Azure/azure-service-operator/pull/4925) `KnownResourceReference` structs were fixed. To [avoid potential breakage](https://github.com/Azure/azure-service-operator/pull/4925), users can add the correct `ResourceReference` property alongside the legacy property.

## 🙏 Thank You

Thank you to all our contributors for making this release possible! A special thanks goes to [bingikarthik](https://github.com/bingikarthik) and [subhamrajvanshi](https://github.com/shubhamrajvanshi) for their contributions to this release.

See the [**Full Release notes**](https://github.com/Azure/azure-service-operator/releases/tag/v2.18.0) on GitHub.
