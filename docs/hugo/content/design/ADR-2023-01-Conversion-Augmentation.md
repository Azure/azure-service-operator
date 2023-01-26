---
title: '2022-12: Conversion Augmentation'
---

## Context

The Kubernetes ecosystem has [strong requirements](https://book.kubebuilder.io/multiversion-tutorial/api-changes.html) around backward and forward compatibility of resource versions:

> In Kubernetes, all versions must be safely round-tripable through each other. This means that if we convert from version 1 to version 2, and then back to version 1, we must not lose information. Thus, any change we make to our API must be compatible with whatever we supported in v1, and also need to make sure anything we add in v2 is supported in v1.

To achieve this, we are [code generating conversions](https://azure.github.io/azure-service-operator/design/versioning/) between versions of our resources.

However, we've always known that we'd need the capability to augment these conversions with some hand-written code, if only because we would quickly hit a point of diminishing returns if we try to handle everything with the generator. 

There are three different ways we can implement this:

* Generic interfaces
* Code-gen interfaces
* Configuration

### Generic interface implementation

In our `genruntime` package, we'll define a new generic interface to capture the semantics of conversion augmentation:

```go
type Assignable[T conversion.Convertible] interface {
	AssignPropertiesTo(dst T) error
	AssignPropertiesFrom(src T) error
}
```
Conversions always need to be bi-directional, so we force both `AssignPropertiesTo` and `AssignPropertiesFrom` to be implemented together.

We then update the code generation for the `AssignProperties_To_*` and `AssignProperties_From_*` methods to look for this interface and call it if it's present. The type used for the type parameter will be the "upstream" type in our conversion chain.

For example, for the `v1beta20210101` version of `BatchAccount`, the existing method `AssignProperties_From_BatchAccount` will updated with the following additional code:

``` go
var sourceAsAny any = source
if src, ok := sourceAsAny.(genruntime.Assignable[*v20210101s.BatchAccount]); ok {
    err := src.AssignFrom(source)
    if err != nil {
        return errors.Wrap(
            err, 
            "calling custom AssignFrom() for conversion from v20210101s.BatchAccount")
    }
}
```

**Pro**: Minimizes the amount of code we need to generate because we only need an extra clause in the existing conversion methods.

**Con**: Requires the user to correctly provide the right type parameter when implementing the interface to augment the property conversions. For example, in this above situation, only an implementation of `genruntime.Assignable[*v20210101s.BatchAccount]` will be invoked; if the user instead implemented `genruntime.Assignable[**v20210101.BatchAccount]` the code would compile, but be silently ignored.

**Con**: Go prohibits doing a type cast on a concrete type, so we need to stash `source` in an intermediate variable first.

### Code-gen interface implementation

Instead of using a generic interface, we can define a new interface for each version of each resource. This interface will be implemented by the user and will be called by the code generator.

For example, in the `v1beta20210101` package, we'll generate an interface for the conversion hook:

``` go
type assignableBatchAccount interface {
	AssignPropertiesToBatchAccount(dst *v20210101s.BatchAccount) error
	AssignPropertiesFromBatchAccount(src *v20210101s.BatchAccount) error
}
```

This interface would then be detected and used if present:

``` go
var sourceAsAny any = source
if src, ok := sourceAsAny.(assignableBatchAccount); ok {
    err := src.AssignPropertiesFromBatchAccount(source)
    if err != nil {
        return errors.Wrap(
            err, 
            "calling custom AssignFrom() for conversion from v20210101s.BatchAccount")
    }
}
```

**Pro**: The interface is only accessible (visible) within the package and has the correct parameter types included. It's not possible for the user to *think* they've implemented the hook point and have it silently ignored. Either the compiler will complain, or the code will work.

**Con**: We don't have any other interfaces being defined by the code generator, so this is a new pattern. (We do have interface implementations).

**Con**: Go prohibits doing a type cast on a concrete type, so we need to stash `source` in an intermediate variable first.

### Configuration implementation

Add new configuration where we can specify that a particular resource needs tweaking. For the Batch account, above, this might look like:

``` yaml
objectModelConfiguration:
  batch:
    2021-01-01:
      BatchAccount:
        $customConversion: true
```

When found, we'd directly generate a call to the conversion functions were needed:

``` go
err := source.AssignPropertiesFromBatchAccount(source)
if err != nil {
    return errors.Wrap(
        err, 
        "calling custom AssignFrom() for conversionfrom v20210101s.BatchAccount")
}
```

**Pro**: The configuration is very explicit and easy to understand.

**Con**: We need to add a new configuration option and the code to process it.

**Con**: This is inconsistent with other extension points we have, which don't use configuration and are automatically detected once implemented.

## Decision

We've previously placed a high value on making things as simple as possible for future developers. Internal consistency within the code generator and making it difficult to do things the wrong way are both important, so we'll adopt the **code-gen interface** approach.


## Status

Proposed.

## Consequences


## Experience Report


## References



