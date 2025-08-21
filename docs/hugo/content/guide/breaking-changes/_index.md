---
title: "Breaking Changes"
linkTitle: "Breaking Changes"
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---
We go to great lengths to avoid breaking changes as much as possible, as we're well aware that they can cause issues for our users. However, they do occasionally occur, so we've committed to providing good documentation each time this occurs.

## Upcoming Breaking Changes

### Resource Versioning

As first reported in [#4147](https://github.com/Azure/azure-service-operator/issues/4147) and explored in the ADR [Resources and Version Priority]({{< relref "adr-2025-05-version-priority" >}}), the way we're defining our resource versions in ASO is not compliant with Kubernetes rules.

This failure results in Kubernetes' [version priority](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/#version-priority) rules consistently selecting the _oldest_ version of any ASO resource.

This has resulted in confusion, and more importantly, to _data loss for users as they work with their resources_. (See [ASO #4723](https://github.com/Azure/azure-service-operator/issues/4723), [CAPZ #5168](https://github.com/kubernetes-sigs/cluster-api-provider-azure/issues/5168), and [CAPZ #5649](https://github.com/kubernetes-sigs/cluster-api-provider-azure/issues/5649) for examples.)

In upcoming releases of ASO, we're going to gradually migrate resources to a new compliant approach using `v1api` as a consistent prefix to just `v`.

This will, eventually, be a breaking change but we're going to try and make this as painless as possible by introducing the new versioning scheme alongside the old one (we'll support both `v1api20250701` and `v20250701` of the resource simultaneously) wherever possible.

## Policies

We know that breaking changes are a problem for many of our users, so we're making the following commitments to keep such things to a minimum.

**No breaking changes to existing resource versions**

We will not introduce a breaking change in an existing version of a resource such as `v1api20230101`. 

In the case of such a change being forced upon us, we will take all reasonable steps to reduce the impact as far as possible.

The most likely cause of this will be a security issue that requires action - such as if we discover a password field that hasn't been correctly flagged as a secret. 
There's also a possibilty of a change ocurring upstream, but we consider this to be extremely unlikely because of the stringent API change criteria used by Azure.

**Minimal breaking changes in new resource versions**

We *may* introduce breaking changes between versions of resources, such as between `v1api20230101` and `v1api20230201`. These two versions correspond to two different Azure API versions (that is, `2023-01-01` and `2023-02-01`), and if the upstream Azure service introduces a breaking change in the newer of these two API versions we may pass that breaking change along. 

### Mitigation techniques

**Automatic conversion**: By manually augmenting the intra-version conversions produced by our code-generator, we may be able to automatically upgrade resources from the older version to the new. Migration would be zero-touch for our users.

We've done this before when a new version of a resource API introduced a breaking change, see [managed_cluster_conversion_overrides.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/containerservice/v1api20210501/storage/managed_cluster_conversion_overrides.go).

**Custom Tooling**: For our `v2.0.0` release, we provided a specialized tool (`asoctl`) specifically designed to smooth the way for upgrading users. In addition to the existing two modes, we may add further functions to cater for other scenarios.

For users upgrading from ASO v1, [`asoctl import azure-resource`]( {{< relref "asoctl#import-azure-resource" >}} ) provides a way to scaffold an ASO v2 resource based on an existing Azure resource. 

For clusters that once had an alpha release of ASO v2 installed, [`asoctl clean crds`]( {{< relref "asoctl#clean-crds" >}} ) does the cleanup required to ensure the upgrade from `v2.0.0-beta.5` to `v2.0.0` goes smoothly.

We may, in future, create additional tooling designed to smooth the way for users encountering what would otherwise be breaking changes.

**Documentation**: As a last resort, if we have a breaking change that we can't avoid (or can't transparently mitigate), we'll explicitly document it, both here and in the release notes for that version.
