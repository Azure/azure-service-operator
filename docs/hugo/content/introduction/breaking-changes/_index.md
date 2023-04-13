---
title: "Breaking Changes"
linkTitle: "Breaking Changes"
---
We go to great lengths to avoid breaking changes as much as possible, as we're well aware that they can cause issues for our users. However, they do occasionally occur, so we've committed to providing good documentation each time this occurs.

## Released Breaking Changes

### v2.0.0

Breaking changes are:

* Upgrades from releases prior to v2.0.0-beta.5 are disallowed
* Alpha CRD versions have been removed

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

## Upcoming Breaking Changes

Beta CRD versions (any version with `v1beta` prefix) will be deprecated no sooner than **v2.3.0**. 
We recommend you start using `v1api` prefixed versions now. 

You can easily swap from a `v1beta` version to a `v1api` version by just replacing `v1beta` with `v1api` in your CRD YAML, the resource shapes are the same.

## Policies

Prior to the v2.0.0 GA release, our policy was *avoid breaking changes*. 
We're now upgrading that to *no avoidable breaking changes*.

If we do encounter potential breaking changes, we have a number of potential mitigations.

### Conversion augmentation

If the automatically generated intra-version conversions are not sufficient, we can manually augment the conversion by implementation an appropriate interface, providing zero-touch mitigation. 

This might occur if a new version of a resource API introduced a breaking change. For an example of this, see [managed_cluster_conversion_overrides.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/containerservice/v1api20210501storage/managed_cluster_conversion_overrides.go).

The advantage of this approach is that migration is zero-touch for our users.

### Custom Tooling

For our **v2.0.0** release, we provided a specialized tool (`asoctl clean crds`) to allow users to upgrade their clusters. For this release, the tooling allows users to upgrade their clusters smoothly even though we are removing the previously deprecated `v1alpha1` resource versions.

We may, in future, create additional tooling designed to smooth the way for users encountering what would otherwise be breaking changes.

### Documentation

As a last resort, if we have a breaking change that we can't avoid, we'll explicitly document it, both here and in the release notes for that version.
