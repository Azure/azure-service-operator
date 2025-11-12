---
title: "v2.16.0 Breaking Changes"
linkTitle: "v2.16.0"
weight: -55  # This should be 5 lower than the previous breaking change document
---
## Breaking changes

* Removed containerservice v1api20210501 v1api20231102preview api verisons. This was required to fit the new versions in the CRD. See [number of versions supported in a CRD](https://azure.github.io/azure-service-operator/guide/crd-management/#number-of-versions-supported-in-a-crd). If you allow the operator to manage its own CRDs via `--crd-pattern`, no action is needed the operator will take care of removing these versions. If you manage the CRD versions yourself, you'll need to run [asoctl clean crds](https://azure.github.io/azure-service-operator/tools/asoctl/#clean-crds) before upgrading.

## Upcoming breaking changes

* Will remove containerservice fleet API version v1api20230315preview in the next release of ASO (v2.17)
