---
title: "ASO v2.17 Release Notes"
date: 2025-12-08
description: "Release notes for Azure Service Operator v2.17.0"
type: blog
---

We're pleased to announce the release of Azure Service Operator v2.17.0! This release brings new PostgreSQL and Log Analytics resources, improved secret exports across several services, and the beginning of our transition to simplified API versioning.

## ⚠️ Breaking changes

This release includes breaking changes. Please review the [breaking changes documentation](https://azure.github.io/azure-service-operator/guide/breaking-changes/) before upgrading.

### Removed Fleet API version

We've removed the `containerservice` v20230315preview versions of Fleet resources, as the API has been deprecated by Azure. If you allow the operator to manage its own CRDs via `--crd-pattern`, no action is needed—the operator will take care of removing these versions automatically. If you manage the CRD versions yourself, you'll need to run [asoctl clean crds](https://azure.github.io/azure-service-operator/tools/asoctl/#clean-crds) before upgrading.

### Upcoming breaking changes

We're giving you advance notice of upcoming deprecations:

- In [ASO v2.18](https://github.com/Azure/azure-service-operator/milestone/30), we will remove `containerservice` ManagedCluster and AgentPool API versions `v1api20230201` and `v1api20231001`.
- In [ASO v2.19](https://github.com/Azure/azure-service-operator/milestone/31), we will remove `containerservice` ManagedCluster and AgentPool API version `v1api20240402preview`.

## 🎉 New and improved resource support

### New resources and API versions

We've added support for several new resources and updated API versions:

- [PostgreSQL Flexible Server Administrator](https://github.com/Azure/azure-service-operator/pull/5041) resource for managing database administrators.
- New [dbforpostgresql API version v1api20250801](https://github.com/Azure/azure-service-operator/pull/5018) with the latest PostgreSQL features.
   - _Special thanks to [tjololo](https://github.com/tjololo) for his contribution!_
- New [operationalinsights API version v1api20250701](https://github.com/Azure/azure-service-operator/pull/5026) for Log Analytics workspaces.

### Enhanced secret exports

We've expanded the ability to export secrets across several services:

- [RedisEnterprise keys](https://github.com/Azure/azure-service-operator/pull/5010) can now be exported to Kubernetes secrets.
- [Shared keys for operationalinsights workspaces](https://github.com/Azure/azure-service-operator/pull/5011) are now available for export.
- [Service Bus TopicAuthorizationRule](https://github.com/Azure/azure-service-operator/pull/5039) now supports secret exports.

### Improved resource references

- PrivateDns [A and AAAA record IP addresses](https://github.com/Azure/azure-service-operator/pull/5027) can now be read from ConfigMaps.

## 🔧 Infrastructure and technical improvements

### Simplified versioning

Resources introduced in this release are using simplified versioning, as we're deprecating the `v1api` prefix (see [#4831](https://github.com/Azure/azure-service-operator/issues/4831)). We've laid the [foundation for migration](https://github.com/Azure/azure-service-operator/pull/5031) to the new versioning style, with [batch resources](https://github.com/Azure/azure-service-operator/pull/5032) being the first to migrate.

### Smarter resource handling

- ASO will [no longer attempt to delete](https://github.com/Azure/azure-service-operator/pull/4987) Azure sub-resources that become a permanent part of their parent resource (such as `FlexibleServersConfiguration` for a PostgreSQL server). Instead of filling the ASO log with deletion errors, you'll now see a helpful warning.
- The operator now applies a [default reconcile policy](https://github.com/Azure/azure-service-operator/pull/5044) when one isn't explicitly configured.
- CRD versions to deprecate are now [auto-detected](https://github.com/Azure/azure-service-operator/pull/5050), streamlining our release process.

### Documentation improvements

- Updated documentation for [ASO v2.16](https://github.com/Azure/azure-service-operator/pull/4993) including [breaking changes](https://github.com/Azure/azure-service-operator/pull/4996).
- Added [badges to our README](https://github.com/Azure/azure-service-operator/pull/5003) for better visibility into project status.
- Documented our [approach for rare property conversions](https://github.com/Azure/azure-service-operator/pull/5005).
- Set up [blog pages](https://github.com/Azure/azure-service-operator/pull/5008) for the ASO website.
   - _Special thanks to [jakjang](https://github.com/jakjang) for his contribution!_

## 🙏 Thank You

Thank you to all our contributors for making this release possible! We're especially grateful to our new contributors [jakjang](https://github.com/jakjang) and [tjololo](https://github.com/tjololo) for their first contributions to ASO. Your engagement helps make ASO better for everyone!

**Full Changelog**: [v2.16.0...v2.17.0](https://github.com/Azure/azure-service-operator/compare/v2.16.0...v2.17.0)


