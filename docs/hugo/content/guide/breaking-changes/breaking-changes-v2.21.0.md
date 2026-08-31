---
title: "v2.21.0 Breaking Changes"
linkTitle: "v2.21.0"
weight: -80  # This should be 5 lower than the previous breaking change document
---

## Breaking changes

* `containerservice`: `ManagedCluster.spec.aadProfile.serverAppSecret` is now correctly identified as a secret.
* `synapse`: `Workspace.spec.encryption.kekIdentity.useSystemAssignedIdentity` is now correctly typed as a Boolean instead of `v1.JSON`. Valid configurations are unaffected, but invalid configurations may require correction.
* `cache`: Removed `RedisEnterprise` and `RedisEnterpriseDatabase` API versions `v1api20210301` and `v1api20230701`. Users should migrate to `v20250401` before upgrading. (#5640)

## Future breaking changes

These breaking changes are planned for future releases:

### TLS 1.3 (v2.22)

In v2.22 we'll default to using TLS 1.3, but will provide a command line flag to enable TLS 1.2 if needed.
