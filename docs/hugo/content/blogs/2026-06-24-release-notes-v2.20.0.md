---
title: "ASO v2.20 Release Notes"
date: 2026-06-24
description: "Release notes for Azure Service Operator v2.20.0"
type: blog
---

We're excited to announce the release of Azure Service Operator v2.20.0! This release extends simplified versioning to four more resource groups, rolls smart deletion prechecks out across more than twenty groups, adds a way to resolve Azure resource names from ConfigMaps at reconciliation time, and introduces support for Microsoft Foundry projects.

## 🎇 Headline Features

Four more resource groups now support [simplified API versioning](https://github.com/Azure/azure-service-operator/issues/4831): `compute`, `app`, `eventgrid`, and `datafactory` ([#5348](https://github.com/Azure/azure-service-operator/pull/5348), [#5342](https://github.com/Azure/azure-service-operator/pull/5342), [#5344](https://github.com/Azure/azure-service-operator/pull/5344), [#5343](https://github.com/Azure/azure-service-operator/pull/5343)). You can now use resources in these groups with the new `v` prefix (e.g. `v20250701` instead of `v1api20250701`), while the legacy `v1api` versions remain available for backward compatibility. We strongly recommend using the new version format wherever it's available, as it sorts correctly for `kubectl get` (see [#4147](https://github.com/Azure/azure-service-operator/issues/4147)) and the `v1api` versions will eventually be removed after a deprecation period.

[Smart deletion prechecks](https://github.com/Azure/azure-service-operator/pull/5394) are now enabled for a large set of groups: `alertsmanagement`, `apimanagement`, `app`, `appconfiguration`, `cache`, `cdn`, `cognitiveservices`, `containerregistry`, `dataprotection`, `dbformariadb`, `eventgrid`, `eventhub`, `kusto`, `managedidentity`, `monitor`, `notificationhubs`, `operationalinsights`, `quota`, `search`, `signalrservice`, `sql`, `synapse`, and `web`. Before issuing a `DELETE` to Azure, ASO checks whether the resource or its parent has already been removed, avoiding unnecessary calls and reducing churn during cleanup of large resource graphs.

ASO can now [resolve Azure resource names from a ConfigMap at reconciliation time](https://github.com/Azure/azure-service-operator/pull/5407) via the new `AzureNameFromConfig` field. This allows the Azure name of a resource to be supplied dynamically (for example, by another controller or pipeline) rather than baked into the resource YAML.

Operator-managed Secrets and ConfigMaps can now carry [custom annotations and labels](https://github.com/Azure/azure-service-operator/pull/5413), making it easier to integrate ASO-exported credentials with tools that rely on metadata (such as service meshes, secret synchronisation tooling, or GitOps controllers).

## 🎉 New and improved resource support

A new [Project resource](https://github.com/Azure/azure-service-operator/pull/5475) has been added to the `cognitiveservices` group, providing support for Microsoft Foundry (formerly Azure AI Studio) projects under existing `AIServices` accounts.

The `apimanagement` group gains a new [ProductGroup resource](https://github.com/Azure/azure-service-operator/pull/5367) for associating groups with API Management products.

## ✨ Other improvements

- Enum validation on `insights` dataflow and datasource streams has been [relaxed](https://github.com/Azure/azure-service-operator/pull/5420), allowing new stream types to be used without waiting for an ASO release.

## 🐛 Bug fixes

- Fixed a bug where a generated [password or password snippet for Sql or PostgreSQL could contain invalid characters](https://github.com/Azure/azure-service-operator/pull/5401).
- Fixed a bug where the [wrong certificate was used for the metrics endpoint](https://github.com/Azure/azure-service-operator/pull/5366).
- Fixed a bug where the [subscription check for skipped resources was incorrect](https://github.com/Azure/azure-service-operator/pull/5360).
- Fixed a bug where the [Helm PodDisruptionBudget template file name was incorrect](https://github.com/Azure/azure-service-operator/pull/5337).

### asoctl

- Fixed a bug where [import failed when looking for Quota resources](https://github.com/Azure/azure-service-operator/pull/5432).
- Fixed a bug where [`ManagedIdentityCredential` was attempted when IMDS was not available](https://github.com/Azure/azure-service-operator/pull/5436).

## 📚 Documentation

- Added new [versioning documentation](https://github.com/Azure/azure-service-operator/pull/5361) explaining the migration from `v1api` to the new version format.
- Added a [warning about subscription ID configuration](https://github.com/Azure/azure-service-operator/pull/5362).
- [Improved the security documentation](https://github.com/Azure/azure-service-operator/pull/5421).
- Added a new [ADR documenting decisions on templating](https://github.com/Azure/azure-service-operator/pull/5410).
- [Updated cache samples](https://github.com/Azure/azure-service-operator/pull/5332) and added an advanced Entra example.

## 🙏 Thank You

Thank you to all our contributors for making this release possible!

See the [**Full Release notes**](https://github.com/Azure/azure-service-operator/releases/tag/v2.20.0) on GitHub.
