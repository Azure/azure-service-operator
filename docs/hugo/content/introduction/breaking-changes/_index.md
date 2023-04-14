---
title: "Breaking Changes"
linkTitle: "Breaking Changes"
---
We go to great lengths to avoid breaking changes as much as possible, as we're well aware that they can cause issues for our users. However, they do occasionally occur, so we've committed to providing good documentation each time this occurs.

## Upcoming Breaking Changes

### No earlier than **v2.3.0**

Beta CRD versions (any version with `v1beta` prefix) will be deprecated and removed.
We recommend you start using `v1api` prefixed versions now. 

You can easily swap from a `v1beta` version to a `v1api` version by just replacing `v1beta` with `v1api` in your CRD YAML, the resource shapes are the same.

## Released Breaking Changes

### v2.0.0

Breaking changes are:

* Upgrades from releases prior to v2.0.0-beta.5 are disallowed
* Alpha CRD versions have been removed
* Structure change for ResourceGroup status

For more information see [v2.0.0 Breaking Changes](./breaking-changes-v2.0.0).

### v2.0.0-beta.4

In the `beta.4` release of Azure Service Operator (ASO) we pivoted to using Azure Swagger API Specifications as the sole source of truth for our code generator. This change brought with it a significant improvement in fidelity - the code we generate is now much closer to what the Azure Swagger API Specifications describe. Unfortunately, this change also brings with it a number of breaking changes requiring simple manual modifications to pre-existing resources.

Breaking changes are:

* The discriminator value for polymorphic types has been changed to match the name of the property used to specify that option.
* Enumerated properties that previously had a base type but no enumeration values have been updated to include the enumeration values.
* Objects that were incorrectly generated as nested properties have been directly inlined
* ARM Id fields are now correctly identified as references, allowing linking to a resource in Kubernetes instead of only in Azure.
* Status properties that cannot be set by the end user on a Spec that were included in the Spec in error
* Sub-resources that were incorrectly inlined into the parent resource have been moved to a separate resource.
* Properties that previously included on Spec but actually had no function have been removed.
* Validation rules have been tightened, or added to properties that previously had no validation rules.
* 
For detailed information, including an exhaustive list of all affected resource properties, see [v2.0.0 Breaking Changes](./breaking-changes-v2.0.0-beta.4).

## Policies

We know that breaking changes are a problem for many of our users, so we're making the following commitments.

**No breaking changes to existing resource versions**

We will not introduce a breaking change in an existing version of a resource such as `v1api20230101`.

In the unlikely event of such a change being forced upon us by an upstream team, we will take all reasonable steps to reduce the impact as far as possible. (We consider this to be extremely unlikely because of the stringent API change criteria used by Azure.)

**Minimal breaking changes in new resource versions**

We *may* introduce breaking changes between versions of resources, such as between `v1api20230101` and `v1api20230201`. These two versions correspond to two different Azure API versions (that is, `2023-01-01` and `2023-02-01`), and if the upstream Azure service introduces a breaking change in the newer of these two API versions we may pass that breaking change along. 

### Mitigation techniques

**Automatic conversion**: By manually augmenting the intra-version conversions produced by our code-generator, we may be able to automatically upgrade resources from the older version to the new. Migration would be zero-touch for our users.

We've done this before when a new version of a resource API introduced a breaking change, see [managed_cluster_conversion_overrides.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/containerservice/v1api20210501storage/managed_cluster_conversion_overrides.go).

**Custom Tooling**: For our **v2.0.0** release, we provided a specialized tool (`asoctl`) specifically designed to smooth the way for upgrading users. In addition to the existing two modes, we may add further functions to cater for other scenarios.

For users upgrading from ASO v1, [`asoctl import azure-resource`](https://azure.github.io/azure-service-operator/tools/asoctl/#import-azure-resource) provides a way to scaffold an ASO v2 resource based on an existing Azure resource. 

For clusters that once had an alpha release of ASO v2 installed, [`asoctl clean crds`](https://azure.github.io/azure-service-operator/tools/asoctl/#clean-crds) does the cleanup required to ensure the upgrade from `v2.0.0-beta.5` to `v2.0.0` goes smoothly.

We may, in future, create additional tooling designed to smooth the way for users encountering what would otherwise be breaking changes.

**Documentation**: As a last resort, if we have a breaking change that we can't avoid (or can't transparently mitigate), we'll explicitly document it, both here and in the release notes for that version.
