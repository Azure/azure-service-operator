---
title: PreReconciliationChecker
linktitle: PreReconciliationChecker
weight: 75
---

## Description

`PreReconciliationChecker` allows resources to perform validation or checks before ARM reconciliation begins. This extension is invoked before sending any requests to Azure, giving resources the ability to block reconciliation until certain prerequisites are met.

The interface is called early in the reconciliation process, after basic claim/validation but before any ARM GET/PUT/PATCH operations. It provides an opportunity to ensure that conditions necessary for successful reconciliation are satisfied.

## Interface Definition

```go
type PreReconciliationChecker interface {
    PreReconcileCheck(
        ctx context.Context,
        obj MetaObject,
        owner MetaObject,
        resourceResolver *resolver.Resolver,
        armClient *genericarmclient.GenericClient,
        log logr.Logger,
        next PreReconcileCheckFunc,
    ) (PreReconcileCheckResult, error)
}

type PreReconcileCheckFunc func(
    ctx context.Context,
    obj MetaObject,
    owner MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
) (PreReconcileCheckResult, error)

// Helper functions for creating results
func ProceedWithReconcile() PreReconcileCheckResult
func BlockReconcile(reason string) PreReconcileCheckResult
```

**Parameters:**
- `ctx`: The current operation context
- `obj`: The resource about to be reconciled (with fresh status)
- `owner`: The parent resource (can be nil for root resources)
- `resourceResolver`: Helper for resolving resource references
- `armClient`: Client for making ARM API calls if needed
- `log`: Logger for the current operation
- `next`: The default check implementation (usually proceeds)

**Returns:**
- `PreReconcileCheckResult`: Indicates proceed or block with optional reason
- `error`: Error if the check itself fails

## Motivation

The `PreReconciliationChecker` extension exists to handle cases where:

1. **Prerequisite validation**: Ensuring required conditions are met before attempting ARM operations

2. **Owner readiness**: Verifying that parent resources are in a state suitable for child creation

3. **Quota checking**: Validating that operation won't exceed limits

4. **Preventing futile operations**: Blocking reconciliation attempts that cannot possibly succeed

5. **External dependencies**: Waiting for external systems or resources to be ready

6. **Resource state validation**: Ensuring the resource is in an appropriate state for reconciliation

The default behavior attempts reconciliation immediately. Some resources need to verify prerequisites first to avoid unnecessary ARM calls, rate limiting, or error conditions.

## When to Use

Implement `PreReconciliationChecker` when:

- ✅ Parent/owner resources must reach certain states first
- ✅ External dependencies must be satisfied before reconciling
- ✅ Spec validation requires complex logic beyond webhooks
- ✅ Reconciliation would fail due to known prerequisites not being met
- ✅ Rate limiting or quota concerns require gating
- ✅ Resource configuration requires specific ordering

Do **not** use `PreReconciliationChecker` when:

- ❌ Simple validation can be done in admission webhooks
- ❌ The check should happen after reconciliation (use PostReconciliationChecker)
- ❌ You're trying to modify the resource (use other extensions)
- ❌ The default behavior works correctly

## Example: Waiting for Owner Ready State

A common pattern is waiting for the parent resource to be ready:

```go
var _ extensions.PreReconciliationChecker = &MyResourceExtension{}

func (ex *MyResourceExtension) PreReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    // Type assert to specific resource type
    resource, ok := obj.(*myservice.MyResource)
    if !ok {
        return extensions.PreReconcileCheckResult{},
            eris.Errorf("cannot run on unknown resource type %T", obj)
    }

    // Type assert hub version to catch breaking changes
    var _ conversion.Hub = resource

    // Check if owner is provided and ready
    if owner != nil {
        ready := conditions.IsReady(owner)
        if !ready {
            ownerName := owner.GetName()
            return extensions.BlockReconcile(
                fmt.Sprintf("waiting for owner %s to be ready", ownerName)), nil
        }
    }

    // Owner is ready (or not required), proceed with reconciliation
    return extensions.ProceedWithReconcile(), nil
}
```

**Key aspects of this example:**

1. **Type assertions**: For both resource type and hub version
2. **Owner check**: Validates owner state before proceeding
3. **Clear blocking messages**: Provides reason for blocking
4. **No error**: Check succeeded, but reconciliation should wait
5. **Proceeds when ready**: Returns proceed result when checks pass

## Common Patterns

### Pattern 1: Check Owner Ready Condition

```go
func (ex *ResourceExtension) PreReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Ensure owner is ready before creating child
    if owner == nil {
        return extensions.BlockReconcile("owner not found"), nil
    }

    if !conditions.IsReady(owner) {
        return extensions.BlockReconcile(
            fmt.Sprintf("owner %s/%s not ready",
                owner.GetNamespace(), owner.GetName())), nil
    }

    // Owner ready, call next checker in chain
    return next(ctx, obj, owner, resourceResolver, armClient, log)
}
```

### Pattern 2: Validate Required References

```go
func (ex *ResourceExtension) PreReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Check if required references are ready
    if resource.Spec.NetworkReference != nil {
        resolved, err := resourceResolver.ResolveResourceReference(
            ctx, resource.Spec.NetworkReference)
        if err != nil {
            return extensions.PreReconcileCheckResult{}, err
        }

        if !resolved.Found {
            return extensions.BlockReconcile(
                "required network reference not found"), nil
        }

        if !conditions.IsReady(resolved.Resource) {
            return extensions.BlockReconcile(
                "required network not ready"), nil
        }
    }

    return extensions.ProceedWithReconcile(), nil
}
```

### Pattern 3: Validate Configuration Prerequisites

```go
func (ex *ResourceExtension) PreReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Complex validation that can't be done in webhook
    if resource.Spec.AdvancedConfig != nil {
        if err := ex.validateAdvancedConfig(resource.Spec.AdvancedConfig); err != nil {
            // Configuration is invalid, block permanently
            return extensions.PreReconcileCheckResult{}, conditions.NewReadyConditionImpactingError(
                err,
                conditions.ConditionSeverityError,
                conditions.ReasonFailed)
        }
    }

    return extensions.ProceedWithReconcile(), nil
}
```

### Pattern 4: Check Azure Resource State

```go
func (ex *ResourceExtension) PreReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Get resource ID if it exists
    resourceID, hasID := genruntime.GetResourceID(resource)
    if !hasID {
        // Not claimed yet, proceed
        return extensions.ProceedWithReconcile(), nil
    }

    // Query current state from Azure
    var azureState MyResourceState
    apiVersion := "2023-01-01"
    _, err := armClient.GetByID(ctx, resourceID, apiVersion, &azureState)
    if err != nil {
        // Handle error appropriately
        return extensions.PreReconcileCheckResult{}, err
    }

    // Check if resource is in a state that allows updates
    if azureState.Status == "Locked" || azureState.Status == "Deleting" {
        return extensions.BlockReconcile(
            fmt.Sprintf("resource in %s state, cannot reconcile", azureState.Status)), nil
    }

    return extensions.ProceedWithReconcile(), nil
}
```

### Pattern 5: Rate Limiting or Throttling

```go
func (ex *ResourceExtension) PreReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Check if we're being rate limited
    if ex.shouldThrottle(resource) {
        return extensions.BlockReconcile(
            "rate limit reached, waiting before retry"), nil
    }

    // Check quota
    quotaOK, err := ex.checkQuota(ctx, resource, armClient)
    if err != nil {
        return extensions.PreReconcileCheckResult{}, err
    }

    if !quotaOK {
        return extensions.BlockReconcile("quota exceeded"), nil
    }

    return extensions.ProceedWithReconcile(), nil
}
```

## Check Results

The extension returns one of two results:

### Proceed
```go
return extensions.ProceedWithReconcile(), nil
```
- Prerequisites are satisfied
- Reconciliation will continue normally
- ARM operations will be attempted

### Block
```go
return extensions.BlockReconcile("reason for blocking"), nil
```
- Prerequisites not met
- Reconciliation skipped for now
- Resource will be requeued to try again later
- Condition set with the blocking reason

### Error
```go
return extensions.PreReconcileCheckResult{}, fmt.Errorf("check failed: %w", err)
```
- The check itself failed
- Error condition set on resource
- Reconciliation blocked until error resolved

## Block vs. Error

Understanding the difference is important:

- **Block**: "Not ready to reconcile yet, try again later" (transient)
  - Example: "waiting for owner", "resource locked"
  - Returns `BlockReconcile(reason), nil`
  - Resource requeued automatically

- **Error**: "Something went wrong during the check" (needs attention)
  - Example: "invalid configuration", "failed to query Azure"
  - Returns `PreReconcileCheckResult{}, error`
  - May be permanent issue requiring user action

## Reconciliation Impact

When a pre-reconciliation check blocks:

1. **No ARM operations**: No requests sent to Azure
2. **Condition set**: Condition added explaining why blocked
3. **Reconciliation requeued**: Will try again automatically
4. **Resource state preserved**: No changes made to resource
5. **Efficient waiting**: Avoids unnecessary ARM calls

This continues until the check passes or the resource is deleted.

## Testing

When testing `PreReconciliationChecker` extensions:

1. **Test proceed case**: Verify check passes when prerequisites met
2. **Test block cases**: Cover all blocking scenarios
3. **Test error handling**: Verify proper error handling
4. **Test with nil owner**: Handle cases with no owner
5. **Test check chain**: Verify calling next() works correctly

Example test structure:

```go
func TestMyResourceExtension_PreReconcileCheck(t *testing.T) {
    t.Run("proceeds when owner ready", func(t *testing.T) {
        // Test success case
    })

    t.Run("blocks when owner not ready", func(t *testing.T) {
        // Test blocking when owner not ready
    })

    t.Run("blocks when owner missing", func(t *testing.T) {
        // Test blocking when owner is nil
    })

    t.Run("proceeds when no owner required", func(t *testing.T) {
        // Test for resources without owners
    })
}
```

## Performance Considerations

Pre-reconciliation checks run on **every** reconciliation attempt, so:

- **Keep them fast**: Avoid expensive operations when possible
- **Cache results**: If appropriate, cache validation results
- **Minimize ARM calls**: Use status/cache over live queries
- **Fail fast**: Return quickly when blocking
- **Be selective**: Only check what's necessary

## Common Use Cases

1. **Owner/Parent Readiness**: Most common - wait for parent to be ready
2. **Reference Resolution**: Ensure referenced resources exist and are ready
3. **Ordering Dependencies**: Ensure resources created in correct order
4. **State Validation**: Verify resource in appropriate state for updates
5. **Quota/Limit Checks**: Prevent operations that would exceed limits
6. **External System Dependencies**: Wait for external systems to be ready

## Important Notes

- **Call `next()` when checks pass**: Allows for check chaining
- **Don't modify the resource**: This is for validation only
- **Provide clear reasons**: Blocking messages shown to users
- **Be idempotent**: Checks may run many times
- **Handle nil owner**: Owner can be nil for root resources
- **Use conditions package**: For setting appropriate conditions
- **Log decisions**: Help debugging by explaining why checks block

## Relationship to Other Validation

Pre-reconciliation checks complement other validation mechanisms:

- **Admission Webhooks**: Validate at create/update time (prefer for simple validation)
- **PreReconciliationChecker**: Validate at reconciliation time (for complex/runtime checks)
- **PostReconciliationChecker**: Validate after reconciliation (for readiness checks)

Use the right tool for the right job:
- Simple field validation → Admission webhook
- Runtime dependency checks → PreReconciliationChecker
- Async readiness validation → PostReconciliationChecker

## Related Extension Points

- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Check after reconciliation
- [ARMResourceModifier]({{< relref "arm-resource-modifier" >}}): Runs after pre-reconciliation check
- [ErrorClassifier]({{< relref "error-classifier" >}}): Handles errors that could have been prevented

## Best Practices

1. **Validate early**: Catch issues before ARM operations
2. **Clear messaging**: Users need to understand why reconciliation is blocked
3. **Appropriate blocking**: Only block when reconciliation would fail
4. **Consider performance**: These run frequently
5. **Use conditions**: Set appropriate conditions when blocking
6. **Handle edge cases**: Nil owners, missing references, etc.
7. **Test thoroughly**: Cover all blocking scenarios
