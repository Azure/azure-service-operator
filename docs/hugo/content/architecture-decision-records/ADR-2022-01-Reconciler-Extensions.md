# Reconciler Extensions

## Context

We are discovering inconsistencies in the way different Azure Resource Providers behave.

These inconsistencies are anticipated. Given the number of disparate teams independently implementing their providers and given the way the ARM guidance has changed over time, differences in behaviour are to be expected.

For example, the Azure Redis Resource Provider can return a HTTP conflict (409) with a retry message indicating that it's not a fatal error but a transient one.

Currently we handle this with special case behaviour in `error_classifier.go` (see [PR#2008](https://github.com/Azure/azure-service-operator/pull/2008) for details).

However, directly changing our generic reconciler is a problem. The growing complexity and increasing number of resources makes this unsustainable, resulting in a risk that changes made to support a new resource type will break an existing one.

We need a way to precisely customize behaviour on a per-resource basis, allowing us to maintain consistent functionality within the operator.

When a new version of a resource is introduced, we don't want to lose any existing customizations. This rules out hosting the extensions directly on the current hub type.

It's quite likely that some extension points will need to import our generated resource types to allow per-version manipulation. This rules out referencing extensions from the generated resource types.

We don't want extension points to reimplement functionality already present in the generic reconciler as this runs the risk of inconsistencies (especially if we change the generic reconciler itself), and introduces a significant overhead to implementation. Instead, the extension points will provide an opportunity to modify the information already held. An extension point that does nothing should have the same effect as one that hasn't been implemented.

## Decision

To provide a stable implementation location for extensions, we'll introduce a dedicated package alongside the existing versioned packages for each supported service. This package will provide extension support for all supported resource versions. For example, alongside the existing **batch** packages `v1alpha1api20210101`, and `v1alpha1api20210101storage`, we will generate `extensions`.

Into the `extensions` package we will codegen a skeleton framework. For each resource there will be a separate extension type, suffixed with `Extensions`. For example, for the **compute** resources `Disk`, `VirtualMachine`, and `VirtualMachineScaleSet` we will create the extension types `DiskExtensions`, `VirtualMachineExtensions`, and `VirtualMachineScaleSetExtensions`.

We'll enhance the existing registration file (`controller_resources_gen.go`) to also provide registration of extensions. They will be indexed by the GVK of each resource, allowing easy lookup by the generic reconciler.

For each supported extension, we'll define an interface. Implementers will assert type compatibility with this extension interface to ensure they have the correct method signature. The generic reconciler will identify available extensions by testing for the presence of the interface.

### Example

To allow customization of error handling for Azure Redis, we'll introduce an extension point, which would look something like this:

We create the required extension interface in `genruntime`:

``` go
type ErrorClassifierExtension interface {
    // Classify the provided error, returning details about the classification.
    // cloudError is the error returned from ARM.
    // apiVersion is the ARM API version used for the request
    // next is the function to call for standard classification behaviour
    ClassifyError(
        cloudError *genericarmclient.CloudError, 
        apiVersion string, 
        next func(*genericarmclient.CloudError) (*CloudErrorDetails, error)) (*CloudErrorDetails, error)
}
```

In the `cache/extensions` package, a host type for the extension point will be generated in `redis_extensions_gen.go`:

```
type RedisExtensions struct{}
```

**TODO**: Should this have any members?

Manual implementation of the extension point will be in `redist_extensions.go` (a different file so that it's not obliterated next time we re-run the generator):

``` go
var _ ErrorClassifierExtension = &RedisExtensions{}

func (e *RedisExtensions) ClassifyError(
    cloudError *genericarmclient.CloudError, 
    next func(*genericarmclient.CloudError) *CloudErrorDetails) (*CloudErrorDetails, error)

    err, result := next(cloudError)
    if err != nil {
        return nil, err
    }

    if inner := cloudError.InnerError; inner != nil {
        code := to.String(inner.Code)
        if code == "Conflict" {
            // For Redis, Conflict can be retried as it can occur because something doesn't yet exist
            result.Classification = CloudErrorRetryable
        }
    }

    return result, nil
}
```

Passing in the default behavour as `next` allows extensions to augmenet and/or supplant the behaviour with a great deal of flexibility.

In the generic reconciler, we make use of the extension:

``` go
func (r *AzureDeploymentReconciler) makeReadyConditionFromError(
    cloudError *genericarmclient.CloudError) conditions.Condition {
    // ... existing code elided ...

    ext := r.FindErrorClassifierExtension()
    errorDetails, err := ext.ClassifyError(cloudError, apiVersion, r.ClassifyCloudError)
    if err != nil {
        // handle error
    }

    // ... existing code elided ...
}
```

The helper method `FindErrorClassifierExtension()` does the lookup for the extension, and returns a null implementation if not found, so the consuming code doesn't need to worry about null checking.

**TODO**: Should we actually return a wrapper around the extension that provides logging and other useful features?

## Status

Proposed.

## Consequences

TBC.

## Experience Report

TBC.

## References


