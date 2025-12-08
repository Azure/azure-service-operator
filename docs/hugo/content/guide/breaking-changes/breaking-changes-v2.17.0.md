---
title: "v2.17.0 Breaking Changes"
linkTitle: "v2.17.0"
weight: -50  # This should be 5 lower than the previous breaking change document
---
## Breaking changes

* Removed containerservice v20230315preview versions of Fleet resources as the API has been deprecated by Azure. If you allow the operator to manage its own CRDs via `--crd-pattern`, no action is needed as the operator will take care of removing these versions. If you manage the CRD versions yourself, you'll need to run [asoctl clean crds](https://azure.github.io/azure-service-operator/tools/asoctl/#clean-crds) before upgrading.

## Upcoming breaking changes

* Will remove containerservice ManagedCluster and AgentPool api version v1api20230201 and v1api20231001 in v2.18 of ASO.
* Will remove containerservice ManagedCluster and AgentPool api version v1api20240402preview in v2.19 of ASO.
