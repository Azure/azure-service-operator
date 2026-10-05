---
title: "ASO v2.21 Release Notes"
date: 2026-08-31
description: "Release notes for Azure Service Operator v2.21.0"
type: blog
---

We're excited to announce the release of Azure Service Operator v2.21.0! This release completes the rollout of smarter deletion, removes retired Redis Enterprise API versions, expands support for Entra ID and several Azure services, and adds more control over operator security and authentication.

## 🎇 Headline Features

The migration to [smart deletion prechecks](https://github.com/Azure/azure-service-operator/issues/5108) is now complete. ASO checks whether a resource or its parent still exists before issuing a `DELETE` to Azure, avoiding unnecessary calls and reducing error churn when you remove large resource graphs. This release enables the precheck for the remaining API groups, including `compute`, `containerinstance`, `datafactory`, `dbforpostgresql`, `devices`, `insights`, `keyvault`, `machinelearningservices`, `network`, `redhatopenshift`, `resources`, `servicebus`, and `storage`.

This release removes the retired and non-functional Redis Enterprise API versions `v1api20210301` and `v1api20230701` ([#5640](https://github.com/Azure/azure-service-operator/pull/5640)). Before upgrading, check your manifests and migrate any Redis Enterprise resources using these versions to a supported API version.

Eight more resource groups now support [simplified API versioning](https://github.com/Azure/azure-service-operator/issues/4831): `authorization`, `cache`, `cdn`, `cognitiveservices`, `containerinstance`, `dataprotection`, `sql`, and `subscription`. You can use resources in these groups with the new `v` prefix while the corresponding `v1api` versions remain available for backward compatibility.

## 🎉 New and improved resource support

ASO now supports Microsoft Entra ID [`Application`]({{< relref "/reference/entra/v1" >}}#Application) resources, allowing you to manage application registrations alongside your Azure resources.

This release adds support for:

- Database Watcher [`Watcher`]({{< relref "/reference/databasewatcher/v20241001preview" >}}#Watcher), [`Target`]({{< relref "/reference/databasewatcher/v20241001preview" >}}#Target), and [`SharedPrivateLink`]({{< relref "/reference/databasewatcher/v20241001preview" >}}#SharedPrivateLink) resources.
- Container Registry [`RegistryCacheRule`]({{< relref "/reference/containerregistry/v20251101" >}}#RegistryCacheRule) resources.
- API Management [`Diagnostic`]({{< relref "/reference/apimanagement/v20240501" >}}#Diagnostic) resources.
- Kusto [`ClusterPrincipalAssignment`]({{< relref "/reference/kusto/v20240413" >}}#ClusterPrincipalAssignment) resources.
- SQL [`ServersKey`]({{< relref "/reference/sql/v20250101" >}}#ServersKey) and [`ServersEncryptionProtector`]({{< relref "/reference/sql/v20250101" >}}#ServersEncryptionProtector) resources.
- Network [`DdosProtectionPlan`]({{< relref "/reference/network/v20250301" >}}#DdosProtectionPlan) resources.

Also, new API versions add support for Managed Cassandra, Redis Enterprise public network access, Flex Consumption function apps, and the latest Event Grid and SQL capabilities.

## ✨ Other improvements

- ASO can now [acquire credentials from an ASO-created managed identity](https://github.com/Azure/azure-service-operator/pull/5490).
- A new [`--tls-min-version` flag](https://github.com/Azure/azure-service-operator/pull/5620) configures the minimum TLS version used by the webhook and metrics servers.
- CRD comparison is faster, improving operator start-up when many CRDs are installed.
- ASO-managed CRDs can carry additional labels.
- AKS agent pools are requeued while control-plane upgrades are in progress, and Front Door custom-domain route association errors are retried sooner.

## 🐛 Bug fixes

- Fixed escaping of special characters in PostgreSQL and Azure SQL user passwords.
- Added the missing Application Gateway private-link IP configuration model.

## 🔭 Looking ahead

ASO v2.22 will default the webhook and metrics servers to TLS 1.3. You will still be able to enable TLS 1.2 with the `--tls-min-version` flag when required.

## 🙏 Thank You

Thank you to all our contributors for making this release possible, including first-time contributors [andreidorin-oprea](https://github.com/andreidorin-oprea) and [Marek Veber](https://github.com/marek-veber).

See the [**full release notes**](https://github.com/Azure/azure-service-operator/releases/tag/v2.21.0) on GitHub for the complete list of changes.
