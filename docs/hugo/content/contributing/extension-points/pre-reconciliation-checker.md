---
title: PreReconciliationChecker
linktitle: PreReconciliationChecker
weight: 75
---

## Description

`PreReconciliationChecker` allows resources to perform validation or checks before ARM reconciliation begins. This extension is invoked before sending any requests to Azure, giving resources the ability to block reconciliation until certain prerequisites are met.

The interface is called early in the reconciliation process, after basic claim/validation but before any ARM GET/PUT/PATCH operations. It provides an opportunity to ensure that conditions necessary for successful reconciliation are satisfied.

## Interface Definition

See the [PreReconciliationChecker interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/prereconciliation_checker.go) in the source code.

## Motivation

The `PreReconciliationChecker` extension exists to handle cases where:

1. **Prerequisite validation**: Ensuring required conditions are met before attempting ARM operations
2. **Preventing futile operations**: Blocking reconciliation attempts that cannot possibly succeed
3. **External dependencies**: Waiting for external systems or resources to be ready

The default behavior of ASO is to attempt reconciliation immediately. Some resources need to verify prerequisites first to avoid unnecessary ARM calls, rate limiting, or error conditions.

## When to Use

Implement `PreReconciliationChecker` when:

- ✅ Resources must reach certain states first
- ✅ External dependencies must be satisfied before reconciling
- ✅ Reconciliation would fail due to known prerequisites not being met

Do **not** use `PreReconciliationChecker` when:

- ❌ The check should happen after reconciliation (use PostReconciliationChecker)
- ❌ You're trying to modify the resource (use other extensions)
- ❌ The default behavior works correctly

## Example: Validate referenced resources are ready

In this example, we block reconciliation if a referenced network resource is not found or not ready, preventing futile ARM calls.

```go
func (ex *ResourceExtension) PreReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
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

## Check Result

The extension returns one of two results, or an error.

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

## When to use Block vs. Error

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
4. **Test check chain**: Verify calling next() works correctly

## Performance Considerations

Pre-reconciliation checks run on **every** reconciliation attempt, so:

- **Keep them fast**: Avoid expensive operations when possible
- **Cache results**: If appropriate, cache validation results
- **Minimize ARM calls**: Use status/cache over live queries (status is refreshed for you)
- **Fail fast**: Return quickly when blocking
- **Be selective**: Only check what's necessary

## Common Use Cases

1. **Reference Resolution**: Ensure referenced resources exist and are ready
2. **Ordering Dependencies**: Ensure resources created in correct order
3. **External System Dependencies**: Wait for external systems to be ready

> **Note:** For owner/parent readiness checks, use [PreReconciliationOwnerChecker]({{< relref "pre-reconciliation-owner-checker" >}}) instead.

## Important Notes

- **Call `next()` when checks pass**: Allows for check chaining
- **Don't modify the resource**: This extension is for validation only
- **Provide clear reasons**: Blocking messages shown to users
- **Be idempotent**: Checks may run many times
- **Use conditions package**: For setting appropriate conditions
- **Log decisions**: Help debugging by explaining why checks block
- **Use PreReconciliationOwnerChecker for owner checks**: Owner validation should use a separate (dedicated) extension

## Related Extension Points

- [PreReconciliationOwnerChecker]({{< relref "pre-reconciliation-owner-checker" >}}): Specialized owner checks
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Check after reconciliation

## Best Practices

1. **Validate early**: Catch issues before ARM operations
2. **Clear messaging**: Users need to understand why reconciliation is blocked
3. **Appropriate blocking**: Only block when reconciliation would fail
4. **Consider performance**: These run frequently
5. **Use conditions**: Set appropriate conditions when blocking
6. **Handle edge cases**: Missing references, invalid states, etc.
7. **Test thoroughly**: Cover all blocking scenarios
