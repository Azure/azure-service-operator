---
title: "v2.22.0 Breaking Changes"
linkTitle: "v2.22.0"
weight: -85  # This should be 5 lower than the previous breaking change document
---

## Breaking changes

### TLS 1.3 by default

The webhook and secure metrics servers now require TLS 1.3 by default. Clients that only support TLS 1.2 can no longer connect unless you explicitly enable TLS 1.2.

Upgrade clients to support TLS 1.3 where possible. If you still need TLS 1.2, set the controller environment variable `TLS_MIN_VERSION=VersionTLS12`. For Helm installations, set the following value:

```yaml
tlsMinVersion: VersionTLS12
```

See [controller TLS settings]({{< relref "aso-controller-settings-options#tls_min_version" >}}) and [#5748](https://github.com/Azure/azure-service-operator/pull/5748) for details.

### Removal of Azure Database for MariaDB resources

ASO no longer supports resources in the `dbformariadb.azure.com` group following the retirement of Azure Database for MariaDB. The resource APIs and controller registrations have been removed.

Before upgrading, replace MariaDB resources in your deployment manifests and migrate any dependent workloads to a supported database service. ASO does not automatically migrate MariaDB databases to another service.

See [#5690](https://github.com/Azure/azure-service-operator/pull/5690) for details.

## Future breaking changes

### Global credential defaults under consideration

We are considering a configuration option to disable global credentials by default in a future release. Global credentials are already optional, and we recommend using namespace-scoped or per-resource credentials instead. Users would still be able to explicitly enable global credentials.

This is a proposal, not a change to global credential behavior in v2.22.0. No target release has been confirmed.

Please share your use cases and concerns, particularly if disabling global credentials by default would cause problems, in [#5719](https://github.com/Azure/azure-service-operator/issues/5719). See [credential scopes]({{< relref "authentication#credential-scope" >}}) for alternatives.
