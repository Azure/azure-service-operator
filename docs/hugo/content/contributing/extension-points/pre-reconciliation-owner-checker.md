---
title: PreReconciliationOwnerChecker
linktitle: PreReconciliationOwnerChecker
weight: 76
---

## Description

`PreReconciliationOwnerChecker` is a specialized variant of `PreReconciliationChecker` that validates only the owner's state before reconciliation, without accessing the resource itself. This extension is invoked before any ARM operations, including GET requests on the resource being reconciled.

The key difference from `PreReconciliationChecker` is that this extension **avoids** performing GET operations on the resource itself, only checking the parent/owner resource. This is critical for resources where the owner's state can completely block access to the child resource.

## Interface Definition

See the [PreReconciliationOwnerChecker interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/prereconciliation_owner_checker.go) in the source code.

## Motivation

The `PreReconciliationOwnerChecker` extension exists to handle a specific class of Azure resources where:

1. **Owner state blocks all access**: The parent resource's state can prevent any operations on child resources, including GET
2. **Cannot determine child state**: You cannot query the child resource to check its state when the owner is in certain states
3. **Avoid wasted API calls**: Attempting to GET or PUT a child when the owner blocks access wastes API quota and generates errors
4. **Owner-dependent access**: Some Azure services completely lock down child resources when the parent is in maintenance, updating, or powered-off states

The most notable example is **Azure Data Explorer (Kusto)**, where you cannot even GET a database when the cluster is powered off or updating. Without this extension, the controller would repeatedly attempt to access the database, failing each time and consuming request quota.

## When to Use

Implement `PreReconciliationOwnerChecker` when:

- ✅ The owner's state can block **all** access to the child resource, including GET operations
- ✅ You need to avoid GET requests on the resource itself
- ✅ The parent resource has states that completely prevent child access (powered off, updating, locked)
- ✅ Attempting operations when the owner is in certain states always fails
- ✅ You want to avoid wasting API quota on operations that will definitely fail

Do **not** use `PreReconciliationOwnerChecker` when:

- ❌ You can safely GET the resource regardless of owner state (use `PreReconciliationChecker` instead)
- ❌ You need access to the resource's current state to make the decision
- ❌ The owner's state doesn't affect access to the child
- ❌ The check can be done after retrieving the resource

## PreReconciliationOwnerChecker vs PreReconciliationChecker

Understanding the difference is critical:

| Aspect              | PreReconciliationOwnerChecker       | PreReconciliationChecker          |
| ------------------- | ----------------------------------- | --------------------------------- |
| **When invoked**    | Before any ARM calls, including GET | After GET, before PUT/PATCH       |
| **Resource access** | No - only owner checked             | Yes - resource status available   |
| **Use case**        | Owner blocks all access             | Validation based on current state |
| **GET avoidance**   | Yes - no GET performed              | No - GET already performed        |
| **Parameters**      | Only `owner`                        | Both `obj` and `owner`            |

**Rule of thumb:** If you can GET the resource to check its state, use `PreReconciliationChecker`. If even GET fails when the owner is in certain states, use `PreReconciliationOwnerChecker`.

## Example: Kusto Database Owner State Check

See the [full implementation in database_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/kusto/customizations/database_extensions.go).

**Key aspects of this implementation:**

1. **Owner type assertion**: Checks if owner is a Kusto Cluster
2. **Nil handling**: Gracefully handles nil owner (ARM ID references)
3. **State checking**: Examines cluster's provisioning state
4. **Clear blocking messages**: Provides specific reason for blocking
5. **Helper function**: Encapsulates state-checking logic
6. **Calls next**: Proceeds when cluster is in acceptable state

## Common Patterns

### Pattern 1: Block on Owner State

```go
func (ex *ResourceExtension) PreReconcileOwnerCheck(
    ctx context.Context,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileOwnerCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    // Handle nil owner
    if owner == nil {
        // No owner to check, proceed
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    // Type assert to specific owner type
    parent, ok := owner.(*parentservice.ParentResource)
    if !ok {
        // Not the expected owner type, proceed
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    // Check owner state
    if parent.Status.State != nil && *parent.Status.State == "PoweredOff" {
        return extensions.BlockReconcile("parent is powered off"), nil
    }

    if parent.Status.State != nil && *parent.Status.State == "Updating" {
        return extensions.BlockReconcile("parent is updating"), nil
    }

    // Owner in acceptable state
    return next(ctx, owner, resourceResolver, armClient, log)
}
```

### Pattern 2: Multiple Blocking States

```go
func (ex *ResourceExtension) PreReconcileOwnerCheck(
    ctx context.Context,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileOwnerCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    if owner == nil {
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    parent, ok := owner.(*parentservice.ParentResource)
    if !ok {
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    // Check multiple conditions that block access
    if parent.Status.State != nil {
        state := *parent.Status.State
        
        blockedStates := map[string]string{
            "Stopped":   "parent is stopped and children cannot be accessed",
            "Stopping":  "parent is stopping, wait for it to complete",
            "Deleting":  "parent is being deleted",
            "Updating":  "parent is updating, operations will conflict",
            "Migrating": "parent is migrating, children unavailable",
        }
        
        if reason, isBlocked := blockedStates[state]; isBlocked {
            return extensions.BlockReconcile(reason), nil
        }
    }

    return next(ctx, owner, resourceResolver, armClient, log)
}
```

### Pattern 3: Conditional Blocking Based on Multiple Factors

```go
func (ex *ResourceExtension) PreReconcileOwnerCheck(
    ctx context.Context,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileOwnerCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    if owner == nil {
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    parent, ok := owner.(*parentservice.ParentResource)
    if !ok {
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    // Check both state and SKU
    // Some SKUs allow operations even during certain states
    if parent.Status.State != nil && *parent.Status.State == "Updating" {
        // Premium SKU allows limited operations during updates
        if parent.Spec.SKU != nil && *parent.Spec.SKU.Name != "Premium" {
            return extensions.BlockReconcile(
                "parent is updating and SKU does not support concurrent operations"), nil
        }
    }

    // Check for maintenance mode
    if parent.Status.InMaintenanceMode != nil && *parent.Status.InMaintenanceMode {
        return extensions.BlockReconcile("parent is in maintenance mode"), nil
    }

    return next(ctx, owner, resourceResolver, armClient, log)
}
```

### Pattern 4: Logging for Debugging

```go
func (ex *ResourceExtension) PreReconcileOwnerCheck(
    ctx context.Context,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PreReconcileOwnerCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
    if owner == nil {
        log.V(Debug).Info("No owner to check, proceeding with reconciliation")
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    parent, ok := owner.(*parentservice.ParentResource)
    if !ok {
        log.V(Debug).Info("Owner is not expected type, proceeding")
        return next(ctx, owner, resourceResolver, armClient, log)
    }

    state := parent.Status.State
    log.V(Debug).Info("Checking parent state", 
        "parentName", parent.Name,
        "state", state)

    if state != nil && *state == "Updating" {
        log.V(Info).Info("Blocking reconciliation due to parent state",
            "parentName", parent.Name,
            "state", *state)
        return extensions.BlockReconcile(
            fmt.Sprintf("parent %s is updating", parent.Name)), nil
    }

    log.V(Debug).Info("Parent state allows reconciliation", "state", state)
    return next(ctx, owner, resourceResolver, armClient, log)
}
```

## Check Results

The extension returns one of two results:

### Proceed

```go
return extensions.ProceedWithReconcile(), nil
```

- Owner state permits reconciliation
- Reconciliation will continue (GET will be attempted)
- Normal reconciliation flow proceeds

### Block

```go
return extensions.BlockReconcile("parent is updating"), nil
```

- Owner state blocks reconciliation
- No GET or other operations attempted on the resource
- Resource requeued to try again later
- Condition set with the blocking reason

### Error

```go
return extensions.PreReconcileCheckResult{}, fmt.Errorf("check failed: %w", err)
```

- The check itself failed
- Error condition set on resource
- Reconciliation blocked until error resolved

## Reconciliation Flow

Understanding where this fits in the reconciliation process:

1. **Resource needs reconciliation**: Controller picks up resource
2. **Owner resolution**: Owner resource identified and fetched
3. **PreReconciliationOwnerChecker invoked**: Extension checks owner state **← YOU ARE HERE**
4. **Decision point**:
   - If **blocked**: Stop here, requeue, set condition
   - If **proceed**: Continue to next step
5. **GET resource**: Fetch resource from Azure (only if not blocked)
6. **PreReconciliationChecker**: Check resource state (if implemented)
7. **ARM operations**: PUT/PATCH/DELETE as needed

## Testing

When testing `PreReconciliationOwnerChecker` extensions:

1. **Test with nil owner**: Verify handling when owner is nil
2. **Test with wrong owner type**: Verify graceful handling
3. **Test blocking states**: Cover all states that should block
4. **Test proceed states**: Verify reconciliation proceeds when appropriate
5. **Test error handling**: Verify proper error returns

## Performance Considerations

This extension is more efficient than `PreReconciliationChecker` when:

- Owner checks are sufficient to make the decision
- Avoiding GET saves API quota
- Owner state changes less frequently than resource state
- Multiple child resources share the same owner (check once, benefit many times)

However, it's less flexible because:

- You cannot examine the resource's current state
- You must make decisions based only on owner information
- You cannot access resource-specific fields or status

## Important Notes

- **No resource access**: You do **not** receive the resource being reconciled
- **Owner only**: Make decisions based solely on owner state
- **Earlier in flow**: Runs before GET, unlike PreReconciliationChecker
- **Nil owner**: Always handle the nil owner case gracefully
- **Type assertions**: Verify owner is expected type before accessing fields
- **Call next()**: Call the next checker in the chain when checks pass
- **Clear reasons**: Provide helpful blocking messages for users
- **Logging**: Log decisions to help debugging

## Real-World Scenarios

### Scenario 1: Kusto (Azure Data Explorer)

**Problem:** Kusto databases cannot be accessed at all when the cluster is stopped or updating.

**Solution:** Check cluster state before attempting any database operations.

```go
// Block on: Creating, Updating, Deleting, Stopped, Stopping
// Allow on: Running, Succeeded
```

### Scenario 2: Virtual Machine Scale Set Instances

**Problem:** Cannot manage individual VMSS instances when the scale set is updating or deleting.

**Solution:** Check VMSS state before attempting instance operations.

```go
// Block on: Updating, Deallocating, Deleting
// Allow on: Running, Succeeded
```

### Scenario 3: Container Service Agent Pools

**Problem:** Agent pools cannot be modified when the cluster is upgrading.

**Solution:** Check AKS cluster upgrade state before pool operations.

```go
// Block on: Upgrading, Updating
// Allow on: Succeeded, Running
```

## Related Extension Points

- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Check resource state (use this unless you need owner-only checks)
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Validate after reconciliation
- [ARMResourceModifier]({{< relref "arm-resource-modifier" >}}): Modify payloads (runs after checks pass)

## Best Practices

1. **Prefer PreReconciliationChecker**: Use this only when necessary
2. **Handle nil owner**: Always check for nil before dereferencing
3. **Type assert safely**: Verify owner type before accessing fields
4. **Document blocking states**: Comment which states block and why
5. **Use helper functions**: Encapsulate state-checking logic
6. **Log decisions**: Help debugging by logging why checks block
7. **Clear messages**: Users need to understand why reconciliation is blocked
8. **Call next()**: Enable check chaining
9. **Test thoroughly**: Cover all owner states and edge cases
