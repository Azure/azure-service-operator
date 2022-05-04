---
title: '2022-01: Reconciler Extensions'
---
## Context

We are discovering inconsistencies in the way different Azure Resource Providers behave.

These inconsistencies are anticipated. Given the number of disparate teams independently implementing their providers and given the way the ARM guidance has changed over time, differences in behaviour are to be expected.

For example, the Azure Redis Resource Provider can return a HTTP conflict (409) with a retry message indicating that it's not a fatal error but a transient one.

Currently we handle this with special case behaviour in `error_classifier.go` (see [PR#2008](https://github.com/Azure/azure-service-operator/pull/2008) for details).

However, directly changing our generic reconciler is a problem. The growing complexity and increasing number of resources makes this unsustainable, resulting in a risk that changes made to support a new resource type will break an existing one.

We need a way to precisely customize behaviour on a per-resource basis, allowing us to maintain consistent functionality within the operator.

When a new version of a resource is introduced, we don't want to lose any existing customizations. This rules out hosting the extensions directly on the current hub type.

It's quite likely that some extension points will need to import our generated resource types to allow per-version manipulation. This rules out referencing extensions from the generated resource types.

We don't want extension points to duplicate functionality already present in the generic reconciler as this runs the risk of inconsistencies (especially if we change the generic reconciler itself), and because it introduces a significant overhead to implementation. Instead, extension points will provide an opportunity to modify the current behaviour. An extension point that does nothing should have the same effect as one that hasn't been implemented.

## Decision

To provide a stable implementation location for extensions, we'll introduce a dedicated `customization` package alongside the existing versioned packages for each supported service. This package will provide customization for all supported resource versions. For example, alongside the existing **batch** packages `v1alpha1api20210101`, and `v1alpha1api20210101storage`, we will generate `customization`.

Into the `customization` package we will codegen a skeleton framework. For each resource there will be a separate customization host type, suffixed with `Extensions`. For example, for the **compute** resources `Disk`, `VirtualMachine`, and `VirtualMachineScaleSet` we will create the extension types `DiskExtensions`, `VirtualMachineExtensions`, and `VirtualMachineScaleSetExtensions`.

We'll enhance the existing registration file (`controller_resources_gen.go`) to also provide registration of extensions. They will be indexed by the GVK of each resource, allowing easy lookup by the generic reconciler.

For each supported extension point in our reconciler, we'll define a separate interface containing a single method with exactly the parameters required for that extension point. These interfaces will be found in the `genruntime/extensions` package.

Implementers will assert type compatibility with the desired extension interface to ensure at compile time they have the correct method signature. This will also ensure we get compilation errors if we unexpectedly need to modify the signature of an extension point in the future.

The generic reconciler will identify available extensions by testing for the presence of the extension interface.

Extension method signature will follow a common pattern:

* The parameter list will start with any input parameters required.
  These parameters will be custom for each extension point.
* A `next` parameter, a **func** that the extension should call to invoke the default behaviour.  
  This allows the extension to act both *before* and *after* the default behaviour, or even to *skip* the default behaviour if necessary.
* An `apiVersion` parameter that receives the actual ARM API version being used.  
  This enables extensions to do different things based on the ARM API version, such as correcting the behaviour of one version.
* A `log` parameter allowing additional information to be logged.  
  Some entry/exit information will be automatically logged, so extensions will only need to log any additions
* Optionally, an extension point may return a value.  
  If multiple return values are needed, these will be wrapped in a custom struct, with a public factory method.
* There will always be an `error` return value, and it will always be the last one.

### Example

To allow customization of error handling for Azure Redis, we'll introduce an extension point for classification of errors.

We create the required extension interface in the package `genruntime/extensions`:

``` go
// error_classifier_extension.go

package extensions

type ErrorClassifier interface {
    // Classify the provided error, returning details about the error including whether it is fatal 
    // or can be retried.
    // cloudError is the error returned from ARM.
    // next is the function to call for standard classification behaviour.
    // apiVersion is the ARM API version used for the request.
    // log is a logger than can be used for telemetry.
    ClassifyError(
        cloudError *genericarmclient.CloudError, 
        next func(*genericarmclient.CloudError) (*CloudErrorDetails, error),
        apiVersion string, 
        log logr.Logger) (*CloudErrorDetails, error)
}
```

In the `cache/customization` package, a host type for the extension point will be generated:

``` go
// redis_extensions.go

package customization

type RedisExtensions struct{}
```

This type doesn't contain any members as we require extension points to be pure functions; if they contain mutable state, we run the risk of introducing concurrency and sequence errors that would be hard to diagnose and fix.

Manual implementation of the extension point will be in a different file so that it's not obliterated next time we re-run the generator:

``` go
// redis_extensions.go

package customization

var _ extensions.ErrorClassifier = &RedisExtensions{}

func (e *RedisExtensions) ClassifyError(
    cloudError *genericarmclient.CloudError, 
    next func(*genericarmclient.CloudError) *CloudErrorDetails,
    apiVersion string, 
    log logr.Logger) (*CloudErrorDetails, error)

    err, result := next(cloudError)
    if err != nil {
        return nil, err
    }

    if inner := cloudError.InnerError; inner != nil {
        code := to.String(inner.Code)
        if code == "Conflict" {
            // For Azure Redis, Conflict errors may be retried, as they are 
            // returned when a required dependency is still being created
            result.Classification = CloudErrorRetryable
        }
    }

    return result, nil
}
```

Passing in the default behaviour as `next` allows extensions to augment and/or supplant the behaviour with a great deal of flexibility.

In the generic reconciler, we make use of the extension:

``` go
// azure_generic_arm_reconciler.go

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

The helper method `FindErrorClassifierExtension()` does the lookup for the extension, and always returns a default implementation that can be invoked. The consuming code therefore doesn't need to worry about null checking, only error handling. To ensure consistent telemetry no matter who implements an extension, this default implementation will include logging.

``` go
// error_classifier_extension.go

package extensions

//TODO: This needs a better name
type ErrorClassifierDefault struct {
    extension ErrorClassifierExtension
}

var _ ErrorClassifier = &ErrorClassifierDefault{}

// NewErrorClassifierDefault creates a new ErrorClassifier that will invoke the passed host if it implements 
// the ErrorClassifier interface.
func NewErrorClassifier(host interface{}) *LoggingErrorClassifier {
    result := &ErrorClassifier{}

    // Hook up our host if appropriate
    if extension, ok := host.(ErrorClassifierExtension); ok {
        result.extension = extension
    }

    return result
}

func (classifier *ErrorClassifier) ClassifyError(
    cloudError *genericarmclient.CloudError, 
    next func(*genericarmclient.CloudError) (*CloudErrorDetails, error),
    apiVersion string, 
    log logr.Logger) (*CloudErrorDetails, error) {

    log.V(Verbose).Info("Classifying error", "error", cloudError)

    var details *CloudErrorDetails
    var err error
    if classifier.extension != null {
      log.V(Verbose).Info("Invoking extension")
      details, err = classifier.extension.ClassifyError(cloudError, next, apiVersion, log)
    } else {
      details, err = next(cloudError, next, apiVersion, log)
    }

    if err != nil {
      log.Error("Failure classifying error", "error", err)
    } else {
      log.V(Verbose).Info("Success classifying error", "result", details)
    }

    return details, err
}
```

## Status

Proposed.

## Consequences

TBC.

## Experience Report

TBC.

## References
