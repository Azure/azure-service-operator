# Reconciler Extensions

## Context

We are discovering inconsistencies in the way different Azure Resource Providers behave.

These inconsistencies are anticipated. Given the number of disparate teams independently implementing their providers and given the way the ARM guidance has changed over time, differences in behaviour are to be expected.

For example, the Azure Redis Resource Provider can return a HTTP conflict (409) with a retry message indicating that it's not a fatal error but a transient one.

Currently we handle this with special case behaviour in `error_classifier.go` (see [PR#2008](https://github.com/Azure/azure-service-operator/pull/2008) for details).

However, directly changing our generic reconciler is a problem. The growing complexity makes this unsustainable, and there is a risk that changes made to support a new resource type will break an existing one.

We need a way to customize its behaviour on a per-resource basis, allowing us to maintain consistent functionality within the operator.

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
    // Classify the provided error.
    // cloudError is the error returned from ARM.
    // details contains the initial analysis by the ASO reconciler, and may be changed if desired.
    // apiVersion is the ARM API version used for the request
    ClassifyError(cloudError *genericarmclient.CloudError, details *CloudErrorDetails, apiVersion string)
}
```

In the `cache/extensions` package, the new extension point will be implemented:

``` go
var _ ErrorClassifierExtension = &RedisExtensions{}

func (e *RedisExtensions) ClassifyError(
    cloudError *genericarmclient.CloudError, 
    details *CloudErrorDetails, 
    apiVersion string) {

    inner := cloudError.InnerError
    if inner == nil {
        return
    }

    code := to.String(inner.Code)
    if code == "Conflict" {
        // For Redis, Conflict can be retried as it can occur because something doesn't yet exist
        details.Classification = CloudErrorRetryable
    }
}
```

This code goes in a different file so that it's not obliterated next time we rerun the generator.

In the generic reconciler, we make use of the extension:

``` go
func (r *AzureDeploymentReconciler) makeReadyConditionFromError(
    cloudError *genericarmclient.CloudError) conditions.Condition {

	var severity conditions.ConditionSeverity
	errorDetails := ClassifyCloudError(cloudError)

    ext := r.FindErrorClassifierExtension()
    ext.ClassifyError(cloudError, &errorDetails, apiVersion /* found where? */)

    // ... existing code elided ...
}
```

The helper method `FindErrorClassifierExtension()` does the lookup for the extension, and returns a null implementation if not found, so the consuming code doesn't need to worry about null checking.

## Status

Proposed.

## Consequences

TBC.

## Experience Report

TBC.

## References


