---
title: "ASO v2.19 Release Notes"
date: 2026-04-24
description: "Release notes for Azure Service Operator v2.19.0"
type: blog
---

We're excited to announce the release of Azure Service Operator v2.19.0! This release expands simplified versioning to six more resource groups, adds new API Management and Cassandra resources, improves reconciliation performance, and tightens security by converting secret-bearing fields to proper secret references.

## 🎇 Headline Features

Six more resource groups now support [simplified API versioning](https://github.com/Azure/azure-service-operator/issues/4831): `alertsmanagement`, `appconfiguration`, `apimanagement`, `dbformysql`, `synapse`, and `web`. You can use all resources in these groups with the new `v` prefix (e.g. `v20250701` instead of `v1api20250701`), while the legacy `v1api` versions remain available for backward compatibility. We recommend you use this new format for groups that support it. We will continue rolling out this version style to all groups over the next release or two.

The default [maximum reconciliation parallelism has been increased to 4](https://github.com/Azure/azure-service-operator/pull/4822), improving throughput for clusters managing many Azure resources simultaneously.

ASO now [splits webhooks across multiple configurations](https://github.com/Azure/azure-service-operator/pull/5218) to avoid hitting Kubernetes CRD size limits, and includes [reduced CRD management memory load on startup](https://github.com/Azure/azure-service-operator/pull/5304). Together, these changes make ASO more reliable in large-scale deployments.

## ⚠️ Breaking changes

This release includes breaking changes. Please review [breaking changes in v2.19](https://azure.github.io/azure-service-operator/guide/breaking-changes/breaking-changes-v2.19.0/) before upgrading.

- The `containerservice` version `v1api20240402preview` has been removed, as [previously announced](https://github.com/Azure/azure-service-operator/releases/tag/v2.18.0). If you allow the operator to manage its own CRDs, no action is needed.

- The `network.azure.com/VirtualNetworkGateway` field `radiusServerSecret` is now a [secret reference](https://github.com/Azure/azure-service-operator/pull/5295) instead of a plain string. Update your resources to use a Kubernetes Secret for this value.

## 🎉 New and improved resource support

A broad set of new [API Management resources](https://github.com/Azure/azure-service-operator/pull/5301) are now available: **Gateway**, **Certificate**, **Logger**, **User**, **Group**, **ApiPolicy**, **ApiDiagnostic**, and gateway sub-resources.

New [Cassandra resources](https://github.com/Azure/azure-service-operator/pull/5175) are now supported under `documentdb`: **CassandraCluster** and **CassandraDataCenter**. Note that you must run `az ad sp show --id a232010e-820c-4083-83bb-3ace5fc29d0b --query id -o tsv` to obtain the required ObjectId for your subscription. See [the sample](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/documentdb/cassandra/v20251015) for details.

New [CommunicationService resource](https://github.com/Azure/azure-service-operator/pull/5263) added under the `communication` group.

New [dbformysql API versions](https://github.com/Azure/azure-service-operator/pull/5173) `v20241230` and `v20250601preview` are available for MySQL Flexible Server.

A new [RedisEnterpriseDatabaseAccessPolicyAssignment](https://github.com/Azure/azure-service-operator/pull/5290) resource is available under the `cache` group.

The `network` group has been updated to [API version v20250301](https://github.com/Azure/azure-service-operator/pull/5295).

## ✨ Other improvements

- `insights.ActionGroup` now supports [populating `serviceUri` from a secret](https://github.com/Azure/azure-service-operator/pull/5308) for `automationRunbookReceiver` and `webhookReceiver`, allowing safer use of URIs with embedded API keys or passwords. Matching status fields have been removed to prevent accidental exposure of sensitive information.

- Private Link Service Connections on PrivateEndpoints now support [alias-based connections](https://github.com/Azure/azure-service-operator/pull/5288).

## 🐛 Bug fixes

- Fixed a bug where [reconciliation could block](https://github.com/Azure/azure-service-operator/pull/5224) under certain conditions.

## 🙏 Thank You

Thank you to all our contributors for making this release possible!

See the [**Full Release notes**](https://github.com/Azure/azure-service-operator/releases/tag/v2.19.0) on GitHub.
