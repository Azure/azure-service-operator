---
title: 2023-07 Package References
---

## Context

We at one point discussed the idea of pushing the spec and status types down into dedicated subpackages, given the resources themselves greater promenance. More recently, we've discussed pushing the ARM types (as an implementation detail) into a dedicated subpackage, in order to provide more structure to the generated code, and to reduce namespace pollution for users who choose to directly consume our object structure.

Doing this would require introduction of a new implementation of the `PackageReference` interface, let's call it `SubPackageReference` for this discussion.

As this interface is implemented, we'll need to revisit all the code that consumes `PackageReference` to make sure that it works as expected with the new subclass.

Ideally, everything would just work. However, our existing object model in this area is a little complex:

{{< figure src="current.png" >}}

<!-- yuml.me class diagram

[<<interface>>;PackageReference]

[PackageReference]<>--[LocalPackageReference]
[PackageReference]<>--[StoragePackageReference]
[PackageReference]<>--[ExternalPackageReference]

[TypeName;Name string]--packageReference >[PackageReference]

[TypeDefinition]--name >[TypeName]
[TypeDefinition]--theType >[Type]

[Type]<>--[TypeName]

[<<interface>>;LocalLikePackageReference]<>--[LocalPackageReference]
[LocalLikePackageReference]<>--[StoragePackageReference]

[StoragePackageReference]-inner >[PackageReference]

-->

The presence of the extra interface `LocalLikePackageReference`, of predicates like `IsExternalPackageReference()`, and of method implementations on `ExternalPackageReference` that panic, point to ways this object model doesn't properly capture our needs.

In particular, this model allows a number of illegal (or impossible) states to be captured - such as a TypeDefinition for a type in an external package, or a StoragePackageReference that's a variant of an ExternalPackage. 

To protect against inadvertent propagation of these illegal states, we have a number of checks throughout the codebase that check for these illegal states and panic if they're ever encountered. This is not ideal, as it's easy to miss a check, and it's not immediately obvious why the check is needed in the first place.

### Option 1

We could just add `SubPackageReference` as a new implementation of `PackageReference`, and leave the existing model as-is. This would be the simplest option, but it would leave the existing model in place, and would leave the existing codebase with the same issues it has today.

* PRO: Simplest option.
* PRO: Least amount of work.
* CON: Leaves existing model in place with existing issues.
* CON: Still requires us to check all uses of `PackageReference`.

### Option 2

Add `SubPackageReference` as a new implementation of `PackageReference`, and pare down the methods on `PackageReference` to just those that make sense across all implementations (e.g. removing `GroupVersion()` and `TryGroupVersion()`). Enhance and make more extensive use of `LocalLikePackageReference`.

<!-- yuml.me class diagram

[<<interface>>;PackageReference]

[PackageReference]<>--[LocalPackageReference]
[PackageReference]<>--[StoragePackageReference]
[PackageReference]<>--[ExternalPackageReference]

[TypeName;Name string]--packageReference >[PackageReference]

[TypeDefinition]--name >[TypeName]
[TypeDefinition]--theType >[Type]

[Type]<>--[TypeName]

[<<interface>>;LocalLikePackageReference]<>--[LocalPackageReference]
[LocalLikePackageReference]<>--[StoragePackageReference]

[StoragePackageReference]-inner >[PackageReference]

[PackageReference]<>--[SubPackageReference]
[SubPackageReference]-parent >[PackageReference]

-->

* PRO: Somewhat tidies up the existing model.
* CON: Still allows illegal states to be represented.
* CON: Requires a type downcast in every case where we need to now use `LocalLikePackageReference` instead of `PackageReference`.

### Option 3

As a part of adding `SubPackageReference`, we might slightly rework the domain model to make illegal states impossible to represent.

Not only will this be easier to understand, but it should streamline some parts of our existing codebase by removing the need to check for illegal states.

{{< figure src="option-3.png" >}}

<!-- yuml.me class diagram

[<<interface>>;PackageReference]
[PackageReference]<>--[LocalPackageReference]
[PackageReference]<>--[StoragePackageReference]
[PackageReference]<>--[SubPackageReference]

[TypeName;Name string]--packageReference >[PackageReference]

[TypeDefinition]--name >[TypeName]
[TypeDefinition]--theType >[Type]

[Type]<>--[TypeName]
[Type]<>--[ExternalTypeName]

[StoragePackageReference]-inner >[PackageReference]
[SubPackageReference]-parent >[PackageReference]

[ExternalTypeName;Name string]--packageReference >[ExternalPackageReference]

-->

We only ever need an `ExternalPackageReference` when naming a type that we import from elsewhere, so we make that a separate case, with a dedicated `ExternalTypeName` that can only reference such an external type.

Breaking the implementation link between `ExternalPackageReference` and `PackageReference` means it's no longer possible for a `TypeDefinition` to be destined for an external package; nor is it possible for a StoragePackageReference to be based on the same. 

The new implementation `SubPackageReference` declares a `parent` package for nesting.

* PRO: Makes illegal states impossible to represent.
* PRO: Removes need for type downcasts.
* PRO: Removes need for checks for illegal states.
* PRO: Makes the model easier to understand.
* CON: Requires slightly more rework of existing code.

### Option 4

As for Option 4, but we also rename types to better reflect the new semantics. Given the introduction of `ExternalPackageReference`, maybe we use `InternalPackageReference`:

{{< figure src="option-4.png" >}}

<!-- yuml.me class diagram

[<<interface>>;InternalPackageReference]
[InternalPackageReference]<>--[ResourcePackageReference]
[InternalPackageReference]<>--[StoragePackageReference]
[InternalPackageReference]<>--[SubPackageReference]

[InternalTypeName;Name string]--packageReference >[InternalPackageReference]

[TypeDefinition]--name >[InternalTypeName]
[TypeDefinition]--theType >[Type]

[Type]<>--[InternalTypeName]
[Type]<>--[ExternalTypeName]

[StoragePackageReference]-inner >[InternalPackageReference]
[SubPackageReference]-parent >[InternalPackageReference]

[ExternalTypeName;Name string]--packageReference >[ExternalPackageReference]

-->


## Status

Option #4 proposed.

## Consequences

TBC.

## Experience Report

TBC

## References

* [#1588](https://github.com/Azure/azure-service-operator/issues/1588) - Improve file layout for generated types

* [#3116](https://github.com/Azure/azure-service-operator/issues/3116) - Refactor ARM types to have only a single copy shared for both Spec/StatusTBC

