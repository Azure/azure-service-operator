---
title: "v2.19.0 Breaking Changes"
linkTitle: "v2.19.0"
weight: -60  # This should be 5 lower than the previous breaking change document
---
## Breaking changes

* Removed containerservice v1api20240402preview versions of ManagedCluster, ManagedClustersAgentPool, and TrustedAccessRoleBinding resources as the API has been deprecated by Azure. If you allow the operator to manage its own CRDs via `--crd-pattern`, no action is needed as the operator will take care of removing these versions. If you manage the CRD versions yourself, you'll need to run [asoctl clean crds](https://azure.github.io/azure-service-operator/tools/asoctl/#clean-crds) before upgrading.

## Upcoming breaking changes
