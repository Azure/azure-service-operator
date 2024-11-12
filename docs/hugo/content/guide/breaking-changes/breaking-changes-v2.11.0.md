---
title: "v2.11.0 Breaking Changes"
linkTitle: "v2.11.0"
weight: -45  # This should be 5 lower than the previous breaking change document
---

## Breaking changes

#### Moved all the "ARM" variants of the CRD types into dedicated subpackages

This is only breaking for consumers of the Go package, not for users of the YAML, and only for those using the ARM types directly.

## Upcoming Breaking changes

### Deprecated managedclusters.containerservice.azure.com API versions

- The v1api20210501 and v1api20231102preview versions will be removed in ASO release 2.12.
- The v1api20230201 version will be removed in ASO release 2.13.

We recommend you move to use a different CRD version to avoid errors.
