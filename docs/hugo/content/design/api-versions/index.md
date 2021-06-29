---
title: API Versions
---

# API Versions

Specification for how we will ensure the ARM API version we use for interaction with ARM matches the version originally requested by a user when they created the resource in their Kubernetes cluster.

## Why do we need this?

Sometimes, in addition to structural changes, there are behaviour changes between ARM API versions. It's therefore important that we use the requested API version when interacting with ARM to ensure that we get the behaviour requested.

### Example

Revisting the CRM example from the [Versioning](../versioning/) specification, consider what happens if we have two available versions of the resource `Person`, lets call them **v1** and **v2**. In **v2** the new properties `PostalAddress` and `ResidentialAddress` are mandatory, requiring that everyone have a both a mailing address and a home. 

![example](example.png)

If we have a valid **v1** `Person`, trying to submit that through the **v2** ARM API will fail because it's missing these addresses.

## Proposed Solution

**Preserve the API Version**: the original API version used to create the custom resource (CR) will be preserved so that we know which API version to use when interacting with ARM.

**New conversion interfaces**: to avoid having to version-convert an entire resource when we are interested in only its `Spec` or `Status`, we'll define and implement new interfaces to allow direct conversion.

**Reconciler integration**: to cleanly plug into the existing reconciler, we'll introduce a new interface and hide the complexity of version-conversion.

Full details of each of these steps follow, below.

## Preserve the API Version

We need to preserve the original API Version of each resource, and use that to create an appropriate resource for ARM.

When generating storage variants, we'll inject a new `OriginalVersion` property of type **string** into the Spec of each resource, providing a place to capture the API version that was originally used to create the resource.

To populate the `OriginalVersion` property on each storage spec, we'll inject an `OriginalVersion()` method (returning **string**) into the API variant of each spec.

![preservation](preservation.png)

API version shown on the left, corresponding Storage version shown on the right.

For each API spec, generated `AssignPropertiesTo*()` method  will read the value of `OriginalVersion()` and write it to the `OriginalVersion` property on the storage variant. The `AssignPropertiesFrom*()` method will ignore `OriginalVersion`.

For each Storage spec, the generated `AssignPropertiesTo*()` and `AssignPropertiesFrom*()` methods will copy the `OriginalVersion` property between versions, preserving the original value.

## New conversion interfaces

We don't want the overhead of doing a complete conversion of each resource if we're only interested in conversion of the associated `Spec` or `Status`. For example, when submitting a request to ARM we need only do a `Spec` conversion; similiarly, when processing a response from ARM, we need only do a `Status` conversion.

We already have the required infrastructure in the form of the `AssignProperties*()` methods, making it simple to implement these new interfaces.

For conversion of `Spec` types:

``` go
type ConvertibleSpec interface {
    // ConvertFromSpec populates the current spec from the passed one
    ConvertFromSpec(source ConvertibleSpec) error
    
    // ConvertToSpec populates the passed spec from the current one
    ConvertToSpec(destination ConvertibleSpec) error
}
```

As we have done with implementations of the traditional [`Convertible`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion#Convertible) interface, these will be implemented on Specs that are further from our hub, and will chain towards that hub.

For conversion of `Status` types, a similar approach:

``` go
type ConvertibleStatus interface {
    // ConvertFromStatus populates the current status from the passed one
    ConvertFromStatus(source ConvertibleStatus) error
    
    // ConvertToStatus populates the passed status from the current one
    ConvertToStatus(destination ConvertibleStatus) error
}
```

## Reconciler integration

When the reconciler is working with ARM, we'll use the `OriginalVersion` property on the storage spec to pivot from the current version to the original version.

To keep the reconciler a simple as possible, we'll handle this with generated code hidden behind a simple interface:

``` go
type GeneratedResource interface {
    // GetSpec returns the spec of the resource ready for 
    // serialization to ARM
    GetSpec(scheme runtime.Scheme) (genruntime.ARMTransformer, error)

    // NewStatus returns a new, empty, status instance ready for
    // population from ARM
    NewStatus(scheme runtime.Scheme) (genruntime.FromARMConverter, error)

    // SetStatus updates the current resource with the status provided
    // The instance passed here should always be the one returned by
    // NewStatus()
    SetStatus(status genruntime.FromARMConverter) error
}
```

The implementation strategy used for API types and Storage types will differ, as detailed below. Conversion between versions of a resource requires access to a `runtime.Scheme` instance, which is why each method has one as a parameter.

### Implementation for API versions

For all API versions of resources, the generated implementation of this interface will work directly with the `Spec` and `Status` of that API type.

``` go
// GetSpec returns the Spec of this resource
func (p *Person) GetSpec(_ runtime.Scheme) (genruntime.ARMTransformer, error) {
    return &p.Spec, nil
}

// NewStatus returns a new blank status, ready to be populated
func (p *Person) NewStatus(_ runtime.Scheme) (genruntime.FromARMConverter, error) {
    return &Person_Status{}, nil
}

// SetStatus accepts a status and updates this resource
// Only a status of the expected type is handled; anything else will error
func (p *Person) SetStatus(status interface{}) error {
    s, ok := status.(*Person_Status);
    if !ok {
        return errors.Errorf(
            "expected Status of type %T, but received %T", p.Status, status)
    } 

    p.Status = *s
    return nil
}
```

### Implementation for storage versions

For the hub Storage version of resources, the generated implementation will transparently handle conversion between versions as dictated by the `OriginalVersion()` function described earlier.

We pivot to the original version by using the `OriginalGVK()` function to obtain a new, empty resource from the schema.

``` go
// GetSpec returns the spec of the current type using 
// the original API version from when the CR was created
func (p *Person) GetSpec(scheme runtime.Scheme) (genruntime.ARMTransformer, error) {
    gvk := p.OriginalGVK()
    
    // Use gvk to create an empty resource of the expected version 
    originalObj, err := scheme.New(gvk)
    if err != nil {
        return errors.Wrapf(
            err, "unable to create requested version %s of resource", gvk)
    }
    resource, ok := originalObj.(genruntime.GeneratedResource)
    if !ok {
        return errors.Wrapf(
            err, "unable to convert %T to genruntime.GeneratedResource", originalObj)
    }

    // Extract the Spec and populate
    resultSpec := resource.GetSpec()
    resultSpec.ConvertFromSpec(p.Spec)

    return &resultSpec, nil
}
```

The `NewStatus()` method needs to pivot to the original version of the resource so that can return the right kind of status.

``` go
func (p *Person) NewStatus(scheme runtime.Scheme) (genruntime.ARMTransformer, error) {
    gvk := p.OriginalGVK()

    // Use gvk to create an empty resource of the expected version
    original, err := scheme.New(gvk)
    if err != nil {
        return errors.Wrapf(
            err, "unable to create requested version %s of resource", gvk)
    }
    resource, ok := originalObj.(genruntime.GeneratedResource)
    if !ok {
        return errors.Wrapf(
            err, "unable to convert %T to genruntime.GeneratedResource", originalObj)
    }

    // Return an empty Status from the empty resource
    return resource.NewStatus(scheme)
}
```

After the reconciler has populated an empty `Status` instance, `SetStatus()` is called for the update. It will be the appropriate API `Spec` as that is returned by `NewStatus()`, so all we need to do is to populate our Hub status.

And, `SetStatus()` needs to convert the provided status to the hub variant.

``` go
func (p *Person) SetStatus(status interface{}) error {
    if _, ok := status.(*PersonStatus); ok {
        // We already have the required type, can short circuit
        p.Status = resource.Status
        return nil
    }

    // Populate an updated status instance
    source := status.(genruntime.ConvertibleStatus)
    destination := &Person_Status{}
    err := source.ConvertToStatus(destination)
    if err != nil {
        return errors.Wrapf(err, "populating %T from %T", destination, source)
    }

    // Copy the converted status over
    p.Status = resource.Status
    return nil
}
```
