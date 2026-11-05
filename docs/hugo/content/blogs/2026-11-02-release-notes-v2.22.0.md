---
title: "ASO v2.22 Release Notes"
date: 2026-11-02
description: "Release notes for Azure Service Operator v2.22.0"
type: blog
---

We're excited to announce the release of Azure Service Operator v2.22.0! This release adds stricter Workload Identity isolation and support for Entra service principals and several new Azure resources. It also defaults the webhook and secure metrics servers to TLS 1.3 and removes support for the retired Azure Database for MariaDB service.

## 🎇 Headline Features

### More control over Workload Identity

ASO now supports an opt-in **strict Workload Identity mode**. Namespace-scoped and per-resource credentials can use separate Kubernetes ServiceAccounts rather than sharing the ASO controller's ServiceAccount subject. This lets you configure separate Federated Identity Credentials for those credentials, improving isolation in multitenant deployments.

Enable this mode with `AZURE_WORKLOAD_IDENTITY_AUTH_MODE=strict`, or set the Helm value `workloadIdentityAuthMode: strict`. The default remains `relaxed`, and the global credential's behavior is unchanged. Strict mode requires additional ServiceAccount and federation setup; see the [credential configuration guide]({{< relref "/guide/authentication/credential-format" >}}) before enabling it.

### Entra service principal creation and adoption

The new Entra [`ServicePrincipal`]({{< relref "/reference/entra/v1" >}}#ServicePrincipal) resource lets you resolve a tenant-local object ID from a known application/client ID, or create a service principal when needed. You can export the resolved ID to a ConfigMap for other resources to consume, removing manual lookup steps from scenarios such as configuring Azure Managed Instance for Apache Cassandra.

Use `AdoptOnly` when an existing principal must not be modified or deleted. The default `AdoptOrCreate` mode manages updates and deletion even for adopted principals. `AdoptOnly` is also available for Entra Application and SecurityGroup resources. See the [Entra service principal guide]({{< relref "/guide/entra-service-principals" >}}) for lifecycle details.

## ⚠️ Breaking changes

The webhook and secure metrics servers now require **TLS 1.3 by default**. If your clients still need TLS 1.2, set `TLS_MIN_VERSION=VersionTLS12`, or use the Helm value `tlsMinVersion: VersionTLS12`.

ASO no longer supports resources in the **`dbformariadb.azure.com` group** following Azure Database for MariaDB retirement. Before upgrading, replace MariaDB resources in your deployment manifests and migrate dependent workloads to a supported database service.

See the [v2.22.0 breaking changes guide]({{< relref "/guide/breaking-changes/breaking-changes-v2.22.0" >}}) for upgrade guidance.

## 🎉 New and improved resource support

Alongside Entra service principals, this release adds support for:

- Service Linker [`Linker`]({{< relref "/reference/servicelinker/v20240401" >}}#Linker) resources.
- Portal [`Dashboard`]({{< relref "/reference/portal/v20250401preview" >}}#Dashboard) resources.
- API Management [`Backend`]({{< relref "/reference/apimanagement/v20250301preview" >}}#Backend) resources with TLS certificate thumbprints.
- SQL [`ManagedInstance`]({{< relref "/reference/sql/v20250101" >}}#ManagedInstance), [`InstanceFailoverGroup`]({{< relref "/reference/sql/v20250101" >}}#InstanceFailoverGroup), [`ManagedInstancesDatabase`]({{< relref "/reference/sql/v20250101" >}}#ManagedInstancesDatabase), and [`ManagedInstancesDatabasesTransparentDataEncryption`]({{< relref "/reference/sql/v20250101" >}}#ManagedInstancesDatabasesTransparentDataEncryption) resources.
- Key Vault [`VaultKey`]({{< relref "/reference/keyvault/v20230701" >}}#VaultKey) resources.
- Service Bus [`QueueAuthorizationRule`]({{< relref "/reference/servicebus/v20240101" >}}#QueueAuthorizationRule) resources, including exporting keys and connection strings to Secrets.
- Event Grid [`Namespace`]({{< relref "/reference/eventgrid/v20250215" >}}#Namespace) and [`NamespaceTopic`]({{< relref "/reference/eventgrid/v20250215" >}}#NamespaceTopic) resources.
- Azure Red Hat OpenShift [`HcpOpenShiftCluster`]({{< relref "/reference/redhatopenshift/v20260901preview" >}}#HcpOpenShiftCluster), [`HcpOpenShiftClustersExternalAuth`]({{< relref "/reference/redhatopenshift/v20260901preview" >}}#HcpOpenShiftClustersExternalAuth), and [`HcpOpenShiftClustersNodePool`]({{< relref "/reference/redhatopenshift/v20260901preview" >}}#HcpOpenShiftClustersNodePool) resources.
- Container Service [`IdentityBinding`]({{< relref "/reference/containerservice/v20260501" >}}#IdentityBinding) resources.

New API versions also expand support for AKS (`v20260501`), Container Registry Registry and RegistryReplication (`v20251101`), Storage's Smart access tier (`v20260601`), and Web Flex Consumption configurations (`v20260715`).

## ✨ Other improvements

- `containerregistry`, `dbforpostgresql`, `devices`, `eventhub`, and `servicebus` now support [simplified API versioning](https://github.com/Azure/azure-service-operator/issues/4831), while legacy `v1api` versions remain available for compatibility.
- ARM requests, polling, final-result requests, and redirects are now restricted to the configured Resource Manager HTTPS origin.
- The release includes protection against compromised eMudhra CA keys (CVE-2026-111943).
- Controllers running as local processes can opt into Azure CLI authentication after `az login` with `--enable-azure-cli-authentication`. This is not intended for the in-cluster operator.
- Leader election and Kubernetes probe timings are configurable, helping operators accommodate different startup and availability requirements.
- Helm users can configure `defaultReconcilePolicy` and the PodDisruptionBudget's `unhealthyPodEvictionPolicy`.
- Entra SecurityGroup owners and members can be managed declaratively.
- Database watcher SharedPrivateLink resources can automatically approve eligible connections to ASO-managed resources referenced through Kubernetes, and wait for approval before reporting readiness. Connections to bare ARM IDs retain their existing behavior.

## 🐛 Bug fixes

- SQL Entra-only authentication now retries while its administrator is still being configured.
- Fixed delete retries that could take more than an hour.
- PostgreSQL replicas are correctly sent as updates when they already exist.
- Database connection string values can no longer be interpreted as additional connection options.
- Workload Identity authentication now honors a configured `AZURE_FEDERATED_TOKEN_FILE` path.
- Fixed invalid truncated owner labels and conversion of nested ARM status conditions.
- Key Vault data-plane requests now include ASO's User-Agent.

## 🔭 Looking ahead

We are considering a configuration option to **disable global credentials by default** in a future release. Users would still be able to explicitly enable them. This is a proposal, not a change to global credential behavior in v2.22.0, and no target release has been confirmed.

Global credentials are already optional, and we recommend using namespace-scoped or per-resource credentials instead. Please share your use cases and concerns, especially if disabling global credentials by default would cause problems, in [the global credential discussion](https://github.com/Azure/azure-service-operator/issues/5719).

## 🙏 Thank You

Thank you to all our contributors and everyone who reported issues, tested changes, and shared feedback to make this release possible!

See the [**full release notes**](https://github.com/Azure/azure-service-operator/releases/tag/v2.22.0) on GitHub for the complete list of changes.
