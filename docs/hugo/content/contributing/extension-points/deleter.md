---
title: Deleter
linktitle: Deleter
weight: 30
---

## Description

`Deleter` allows a resource to customize how ASO deletes it from Azure. ASO invokes the extension after Kubernetes sets the resource's deletion timestamp and before issuing the standard ARM DELETE request.

The extension returns an `extensions.DeleteResult` describing what ASO should do next. This allows the extension to complete or block deletion, delegate to the standard ARM deletion path, or start a long-running operation that ASO monitors.

## Interface Definition

See the [Deleter interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/deleter.go) and [DeleteResult definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/delete_result.go) in the source code.

`Delete()` has the following signature:

```go
Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (extensions.DeleteResult, error)
```

## Deletion Results

Return one of the following results to describe the deletion state:

| Result | Meaning |
| --- | --- |
| `extensions.BlockDelete(message)` | Keep the finalizer and retry later. ASO exposes `message` in the resource's Ready condition. |
| `extensions.DeleteCompleted()` | Deletion is complete. ASO can remove the finalizer without issuing another ARM DELETE request. |
| `extensions.MonitorDelete(pollerResponse)` | A long-running deletion operation has started. ASO stores its resume token and requeues until it completes. |

Return `extensions.DeleteResult{}` with an error. ASO ignores the result when the error is non-nil and retries reconciliation according to the error classification.

If you want the deletion to proceed, call the `next()` function, which invokes the standard ARM DELETE operation. The result from `next()` can be returned directly or modified before returning.

## Motivation

The `Deleter` extension exists to handle cases where:

1. **Pre-deletion operations**: Resources that need to perform cleanup before being deleted from Azure (e.g., canceling subscriptions, disabling features)
2. **Multi-step deletion**: Resources requiring multiple API calls in a specific order to delete properly
3. **Dependent resource cleanup**: Resources that need to ensure dependent resources are handled before deletion
4. **Soft-delete handling**: Resources with soft-delete capabilities that may need special deletion modes
5. **Conditional deletion**: Resources that should skip Azure deletion under certain circumstances (e.g., externally managed resources)
6. **Coordinated deletion**: Resources that need to coordinate with other Azure services during deletion

## When to Use

Implement `Deleter` when:

- Pre-deletion operations must be performed, such as canceling or disabling a resource.
- Multiple Azure API calls are needed for complete deletion.
- Deletion order matters across related resources.
- Custom error handling is needed during deletion.
- Soft-delete or purge operations require special logic.
- The resource should be preserved in Azure in some scenarios.

Do **not** use `Deleter` when:

- The standard DELETE operation works correctly.
- You only need to clean up Kubernetes resources; use finalizers instead.
- The logic should apply to all resources; modify the controller instead.
- You're working around an Azure API bug; fix or report the bug instead.

## Example: Subscription Alias Deletion

See the [full implementation in alias_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/subscription/customizations/alias_extensions.go).

**Key aspects of this implementation:**

1. **Type assertions**: For both resource type and hub version
2. **Conditional logic**: Checks if subscription ID is available
3. **Pre-deletion operation**: Cancels subscription before deleting alias
4. **Error handling**: Returns errors that prevent finalizer removal
5. **Chain pattern**: Calls `next()` to perform standard deletion
6. **Logging**: Clear logging of each step for debugging

## Common Patterns

### Pattern 1: Simple Pre-deletion Operation

```go
func (ex *ResourceExtension) Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (extensions.DeleteResult, error) {
    resource := obj.(*myservice.MyResource)

    // Perform cleanup operation
    log.V(Status).Info("Performing pre-deletion cleanup")
    if err := ex.performCleanup(ctx, resource, armClient); err != nil {
        return extensions.DeleteResult{}, eris.Wrap(err, "cleanup failed")
    }

    // Proceed with standard deletion
    return next(ctx, log, resolver, armClient, obj)
}
```

### Pattern 2: Conditional Deletion

```go
func (ex *ResourceExtension) Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (extensions.DeleteResult, error) {
    resource := obj.(*myservice.MyResource)

    // Check if resource should be preserved in Azure
    if ex.shouldPreserve(resource) {
        log.V(Status).Info("Skipping Azure deletion, resource marked for preservation")
        // Report completion without calling next(), allowing ASO to remove the finalizer.
        return extensions.DeleteCompleted(), nil
    }

    // Proceed with normal deletion
    return next(ctx, log, resolver, armClient, obj)
}
```

### Pattern 3: Temporarily Block Deletion

Use `BlockDelete()` when a prerequisite can become ready later. ASO keeps the finalizer, updates the Ready condition with the supplied message, and retries the deletion.

```go
func (ex *ResourceExtension) Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (extensions.DeleteResult, error) {
    resource := obj.(*myservice.MyResource)

    ready, err := ex.dependentsAreDeleted(ctx, resource, armClient)
    if err != nil {
        return extensions.DeleteResult{}, eris.Wrap(err, "checking dependent resources")
    }

    if !ready {
        return extensions.BlockDelete("Waiting for dependent resources to be deleted"), nil
    }

    return next(ctx, log, resolver, armClient, obj)
}
```

### Pattern 4: Soft Delete with a Long-running Purge

This pattern applies when the standard soft-delete operation completes synchronously and the subsequent purge returns an ARM long-running operation. Return the purge poller response through `MonitorDelete()` so ASO persists and resumes the operation.

```go
func (ex *ResourceExtension) Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (extensions.DeleteResult, error) {
    resource := obj.(*myservice.MyResource)

    // For this resource type, standard deletion moves the resource into its
    // soft-deleted state synchronously.
    result, err := next(ctx, log, resolver, armClient, obj)
    if err != nil {
        return result, err
    }

    if resource.Spec.DeleteMode == nil || *resource.Spec.DeleteMode != "Purge" {
        return result, nil
    }

    log.V(Status).Info("Purging soft-deleted resource")
    pollerResponse, err := ex.beginPurge(ctx, resource, armClient)
    if err != nil {
        return extensions.DeleteResult{}, eris.Wrap(err, "starting purge")
    }

    return extensions.MonitorDelete(pollerResponse), nil
}
```

Do not start the purge until the soft-delete operation has completed. If the standard DELETE operation for a resource can return a long-running operation, model the soft-delete and purge as separate idempotent stages.

## Deletion Lifecycle

Understanding the deletion process:

1. **User deletes resource**: `kubectl delete` sets deletion timestamp
2. **Finalizer blocks deletion**: ASO finalizer prevents immediate removal from Kubernetes
3. **Deleter invoked**: Custom `Delete()` method is called
4. **Pre-deletion logic**: Extension performs custom operations
5. **Result returned**: The extension returns a `DeleteResult` or delegates to `next()`.
6. **Deletion handled**: ASO blocks, completes, or monitors deletion based on the result.
7. **Operation monitored**: For `MonitorDelete()`, ASO stores the resume token and requeues until Azure reports completion.
8. **Finalizer removed**: ASO removes the finalizer after deletion completes, allowing Kubernetes to remove the resource.

If any step returns an error or blocks deletion, ASO keeps the finalizer and retries on a later reconciliation.

## Error Handling

Proper error handling in deleters is critical:

```go
// Transient error - will retry
return extensions.DeleteResult{}, eris.Wrap(err, "temporary failure")

// Permanent error with condition
return extensions.DeleteResult{}, conditions.NewReadyConditionImpactingError(
    err,
    conditions.ConditionSeverityError,
    conditions.ReasonFailed)

// Prerequisite not ready - Ready condition is updated and deletion retries later
return extensions.BlockDelete("Waiting for dependent resources to be deleted"), nil

// Custom deletion completed - finalizer can be removed
return extensions.DeleteCompleted(), nil

// Long-running deletion started - ASO stores and resumes the poller
return extensions.MonitorDelete(pollerResponse), nil

// Run the standard ARM deletion path
return next(ctx, log, resolver, armClient, obj)
```

## Testing

When testing `Deleter` extensions:

1. **Test successful deletion**: Verify the happy path works
2. **Test pre-deletion operations**: Ensure cleanup logic executes
3. **Test error scenarios**: Verify error handling prevents finalizer removal
4. **Test idempotency**: Multiple calls should be safe
5. **Test conditional paths**: Cover all branching logic
6. **Test each result**: Verify completed, blocked, and monitored deletion paths.
7. **Test poller handling**: Verify long-running operations return the expected poller response.

## Important Notes

- **Always call `next()` unless**: You have a very specific reason to skip Azure deletion
- **Handle missing IDs gracefully**: Resource might not have been created in Azure yet
- **Return the correct result**: Use `DeleteCompleted()`, `BlockDelete()`, or `MonitorDelete()` rather than constructing a controller result.
- **Let ASO monitor operations**: Return the poller response instead of storing resume tokens or scheduling retries in the extension.
- **Log clearly**: Deletion issues are hard to debug, good logging helps
- **Be idempotent**: Deletion might be called multiple times
- **Don't leak resources**: Ensure Azure resources are eventually deleted

## Related Extension Points

- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Validate before operations
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Validate after operations
- [SuccessfulCreationHandler]({{< relref "successful-creation-handler" >}}): Handle successful creation
