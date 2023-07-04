---
title: Package References
---

## Context

We at one point discussed the idea of pushing the spec and status types down into dedicated subpackages, given the resources themselves greater promenance. More recently, we've discussed pushing the ARM types (as an implementation detail) into a dedicated subpackage, in order to provide more structure to the generated code, and to reduce namespace pollution for users who choose to consume our object structure.

Doing this would require introduction of a new implementation of the `PackageReference` interface, let's call it `SubPackageReference` for this discussion.

As this interface is implemented, we'll need to revisit all the code that consumes package references to make sure that it works as expected with the new subclass.

Ideally, everything would just work. However, our existing object model in this area is a little complex:

{{ $current := .Resources.GetMatch "current.png" }}
<img src="{{ $current.RelPermalink }}" width="{{ $current.Width }}" height="{{ $current.Height }}">

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

In particular, this model allows a number of illegal (or impossible) states to be captured - such as a TypeDefinition for a type in an external package, or a StoragePackageReference that's a variant of an ExternalPackage. We have code throughout ASO to check for and panic if these are ever used.

## Decision

As a part of adding `SubPackageReference`, we should slightly rework the domain model to make illegal states impossible to represent.

Not only will this be easier to understand, but it should streamline some parts of our existing codebase by removing the need to check for illegal states.

{{ $proposed := .Resources.GetMatch "proposed.png" }}
<img src="{{ $proposed.RelPermalink }}" width="{{ $proposed.Width }}" height="{{ $proposed.Height }}">

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

## Status

TBC

## Consequences

TBC

## Experience Report

TBC

## References

TBC

