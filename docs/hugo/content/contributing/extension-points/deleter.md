---
title: Deleter
linktitle: Deleter
weight: 30
---

## Description

`Deleter` allows resources to customize how the reconciler deletes them from Azure. This extension is invoked when a resource has a deletion timestamp in Kubernetes (indicating the user wants to delete it) and gives the resource control over the deletion process.

The interface is called after Kubernetes marks the resource for deletion but before the standard ARM DELETE operation. This allows resources to perform cleanup, handle special deletion scenarios, or coordinate multiple deletion operations.

## Interface Definition

```go
type Deleter interface {
    Delete(
        ctx context.Context,
        log logr.Logger,
        resolver *resolver.Resolver,
        armClient *genericarmclient.GenericClient,
        obj genruntime.ARMMetaObject,
        next DeleteFunc,
    ) (ctrl.Result, error)
}

type DeleteFunc = func(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
) (ctrl.Result, error)
```

**Parameters:**
- `ctx`: The current operation context
- `log`: Logger for the current operation
- `resolver`: Helper for resolving resource references
- `armClient`: Client for making ARM API calls
- `obj`: The Kubernetes resource being deleted
- `next`: The default deletion implementation to call

**Returns:**
- `ctrl.Result`: Reconciliation result (e.g., requeue timing)
- `error`: Error if deletion fails (will prevent finalizer removal)

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

- ✅ Pre-deletion operations must be performed (e.g., canceling, disabling)
- ✅ Multiple Azure API calls are needed for complete deletion
- ✅ Deletion order matters across related resources
- ✅ Custom error handling is needed during deletion
- ✅ Soft-delete or purge operations require special logic
- ✅ The resource should be preserved in Azure in some scenarios

Do **not** use `Deleter` when:

- ❌ The standard DELETE operation works correctly
- ❌ You only need to clean up Kubernetes resources (use finalizers)
- ❌ The logic should apply to all resources (modify the controller)
- ❌ You're working around an Azure API bug (fix/report the bug)

## Example: Subscription Alias Deletion

The Subscription Alias resource uses `Deleter` to cancel the subscription before deleting the alias:

```go
var _ extensions.Deleter = &AliasExtension{}

func (extension *AliasExtension) Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (ctrl.Result, error) {
    // Type assert to the specific resource type
    typedObj, ok := obj.(*storage.Alias)
    if !ok {
        return ctrl.Result{}, eris.Errorf(
            "cannot run on unknown resource type %T, expected *subscription.Alias",
            obj)
    }

    // Type assert hub version to catch breaking changes
    var _ conversion.Hub = typedObj

    // Get the subscription ID from the alias status
    subscriptionID, ok := getSubscriptionID(typedObj)
    if !ok {
        // SubscriptionID isn't populated, skip cancellation and proceed with deletion
        log.V(Status).Info("No subscription ID found, proceeding with alias deletion")
        return next(ctx, log, resolver, armClient, obj)
    }

    // Create Azure SDK client for subscription operations
    subscriptionClient, err := armsubscription.NewClient(
        armClient.Creds(),
        armClient.ClientOptions())
    if err != nil {
        return ctrl.Result{}, eris.Wrapf(err, "failed to create subscription client")
    }

    // Cancel the subscription before deleting the alias
    log.V(Status).Info("Canceling subscription", "subscriptionId", subscriptionID)
    _, err = subscriptionClient.Cancel(ctx, subscriptionID, nil)
    if err != nil {
        return ctrl.Result{}, eris.Wrapf(err, "failed to cancel subscription %q", subscriptionID)
    }

    log.V(Status).Info("Subscription canceled, proceeding with alias deletion")

    // Now perform the standard deletion of the alias
    return next(ctx, log, resolver, armClient, obj)
}
```

**Key aspects of this example:**

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
) (ctrl.Result, error) {
    resource := obj.(*myservice.MyResource)

    // Perform cleanup operation
    log.V(Status).Info("Performing pre-deletion cleanup")
    if err := ex.performCleanup(ctx, resource, armClient); err != nil {
        return ctrl.Result{}, eris.Wrap(err, "cleanup failed")
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
) (ctrl.Result, error) {
    resource := obj.(*myservice.MyResource)

    // Check if resource should be preserved in Azure
    if ex.shouldPreserve(resource) {
        log.V(Status).Info("Skipping Azure deletion, resource marked for preservation")
        // Return success without calling next() - finalizer will be removed
        return ctrl.Result{}, nil
    }

    // Proceed with normal deletion
    return next(ctx, log, resolver, armClient, obj)
}
```

### Pattern 3: Multi-step Deletion with Retry

```go
func (ex *ResourceExtension) Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (ctrl.Result, error) {
    resource := obj.(*myservice.MyResource)

    // Step 1: Disable the resource
    if !ex.isDisabled(resource) {
        log.V(Status).Info("Disabling resource before deletion")
        if err := ex.disableResource(ctx, resource, armClient); err != nil {
            return ctrl.Result{}, eris.Wrap(err, "failed to disable resource")
        }
        // Requeue to wait for disable to complete
        return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
    }

    // Step 2: Wait for dependent resources to be cleaned up
    if ex.hasDependents(ctx, resource) {
        log.V(Status).Info("Waiting for dependent resources to be deleted")
        return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
    }

    // Step 3: Proceed with deletion
    log.V(Status).Info("All prerequisites met, proceeding with deletion")
    return next(ctx, log, resolver, armClient, obj)
}
```

### Pattern 4: Soft Delete with Purge Option

```go
func (ex *ResourceExtension) Delete(
    ctx context.Context,
    log logr.Logger,
    resolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    obj genruntime.ARMMetaObject,
    next extensions.DeleteFunc,
) (ctrl.Result, error) {
    resource := obj.(*myservice.MyResource)

    // Perform standard deletion (moves to soft-deleted state)
    result, err := next(ctx, log, resolver, armClient, obj)
    if err != nil {
        return result, err
    }

    // If purge is requested, purge the soft-deleted resource
    if resource.Spec.DeleteMode != nil && *resource.Spec.DeleteMode == "Purge" {
        log.V(Status).Info("Purging soft-deleted resource")
        if err := ex.purgeResource(ctx, resource, armClient); err != nil {
            return ctrl.Result{}, eris.Wrap(err, "failed to purge resource")
        }
    }

    return ctrl.Result{}, nil
}
```

## Deletion Lifecycle

Understanding the deletion process:

1. **User deletes resource**: `kubectl delete` sets deletion timestamp
2. **Finalizer blocks deletion**: ASO finalizer prevents immediate removal from Kubernetes
3. **Deleter invoked**: Custom `Delete()` method is called
4. **Pre-deletion logic**: Extension performs custom operations
5. **Standard deletion**: `next()` sends DELETE to ARM
6. **ARM deletion completes**: Azure resource is removed
7. **Finalizer removed**: Kubernetes removes the resource

If any step fails, the process pauses and will retry on the next reconciliation.

## Error Handling

Proper error handling in deleters is critical:

```go
// Transient error - will retry
return ctrl.Result{}, eris.Wrap(err, "temporary failure")

// Permanent error with condition
return ctrl.Result{}, conditions.NewReadyConditionImpactingError(
    err,
    conditions.ConditionSeverityError,
    conditions.ReasonFailed)

// Requeue for later retry
return ctrl.Result{RequeueAfter: 1 * time.Minute}, nil

// Success
return ctrl.Result{}, nil
```

## Testing

When testing `Deleter` extensions:

1. **Test successful deletion**: Verify the happy path works
2. **Test pre-deletion operations**: Ensure cleanup logic executes
3. **Test error scenarios**: Verify error handling prevents finalizer removal
4. **Test idempotency**: Multiple calls should be safe
5. **Test conditional paths**: Cover all branching logic
6. **Test requeue behavior**: Verify multi-step deletions requeue correctly

Example test structure:

```go
func TestMyResourceExtension_Delete(t *testing.T) {
    t.Run("successful deletion with cleanup", func(t *testing.T) {
        // Test full deletion flow
    })

    t.Run("cleanup fails blocks deletion", func(t *testing.T) {
        // Test error handling
    })

    t.Run("conditional preservation", func(t *testing.T) {
        // Test skip deletion logic
    })

    t.Run("multi-step deletion", func(t *testing.T) {
        // Test requeue behavior
    })
}
```

## Important Notes

- **Always call `next()` unless**: You have a very specific reason to skip Azure deletion
- **Handle missing IDs gracefully**: Resource might not have been created in Azure yet
- **Return appropriate Results**: Use `RequeueAfter` for async operations
- **Log clearly**: Deletion issues are hard to debug, good logging helps
- **Be idempotent**: Deletion might be called multiple times
- **Don't leak resources**: Ensure Azure resources are eventually deleted

## Related Extension Points

- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Validate before operations
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Validate after operations
- [SuccessfulCreationHandler]({{< relref "successful-creation-handler" >}}): Handle successful creation
