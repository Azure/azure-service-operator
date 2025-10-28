---
title: Claimer
linktitle: Claimer
weight: 20
---

## Description

`Claimer` allows resources to customize how the reconciler claims ownership of a resource. The claiming process establishes the link between a Kubernetes resource and its corresponding Azure resource by setting the Azure Resource Manager (ARM) ID and other status fields.

This extension is invoked early in the reconciliation process, before any ARM operations are performed. It gives the resource control over how and when it establishes its identity in Azure.

## Interface Definition

```go
type Claimer interface {
    Claim(
        ctx context.Context,
        log logr.Logger,
        obj genruntime.ARMOwnedMetaObject,
        next ClaimFunc,
    ) error
}

type ClaimFunc = func(
    ctx context.Context,
    log logr.Logger,
    obj genruntime.ARMOwnedMetaObject,
) error
```

**Parameters:**
- `ctx`: The current operation context
- `log`: Logger for the current operation
- `obj`: The Kubernetes resource being claimed
- `next`: The default claim implementation to call

**Returns:**
- Error if claiming fails (will prevent reconciliation)

## Motivation

The `Claimer` extension exists to handle cases where:

1. **Custom ID construction**: Some resources need special logic to construct their Azure Resource ID that differs from the standard pattern

2. **Conditional claiming**: Resources that may need to defer claiming until certain conditions are met

3. **Alternative identity sources**: Resources that get their ARM ID from somewhere other than the standard owner hierarchy

4. **Wrapped claim logic**: Resources that need to perform additional operations before or after the standard claim

The standard claiming process constructs the ARM ID from the resource's name, owner hierarchy, and Azure conventions. Some resources need to deviate from this pattern.

## When to Use

Implement `Claimer` when:

- ✅ The resource's ARM ID doesn't follow standard construction patterns
- ✅ Claiming depends on external state or conditions
- ✅ Additional operations must be performed during claiming
- ✅ The resource needs custom owner resolution logic
- ✅ Special validation is required before establishing resource identity

Do **not** use `Claimer` when:

- ❌ The standard claiming logic works correctly
- ❌ The issue is with owner resolution (that's a different concern)
- ❌ You only need to validate after claiming (use PreReconciliationChecker)
- ❌ The resource just needs a custom name (use AzureName field)

## Example: Standard Claim with Logging

Most resources don't need custom claiming. Here's a minimal example that wraps the default behavior with additional logging:

```go
var _ extensions.Claimer = &MyResourceExtension{}

func (ex *MyResourceExtension) Claim(
    ctx context.Context,
    log logr.Logger,
    obj genruntime.ARMOwnedMetaObject,
    next extensions.ClaimFunc,
) error {
    // Type assert to your specific resource type
    resource, ok := obj.(*myservice.MyResource)
    if !ok {
        return eris.Errorf(
            "cannot run on unknown resource type %T",
            obj)
    }

    // Type assert hub version to catch breaking changes
    var _ conversion.Hub = resource

    // Perform pre-claim operations
    log.V(Status).Info("Performing custom claim logic")

    // Call the default claim implementation
    err := next(ctx, log, obj)
    if err != nil {
        return eris.Wrap(err, "failed to claim resource")
    }

    // Perform post-claim operations
    log.V(Status).Info("Claim successful", "armId", genruntime.GetResourceID(obj))

    return nil
}
```

**Key aspects:**
1. **Chain pattern**: The `next` parameter allows calling the default implementation
2. **Wrapping logic**: Additional operations can be performed before and after claiming
3. **Type safety**: Type assertions ensure correct resource type
4. **Error handling**: Errors from the default claim are properly wrapped

## Common Patterns

### Pattern 1: Conditional Claiming

```go
func (ex *ResourceExtension) Claim(
    ctx context.Context,
    log logr.Logger,
    obj genruntime.ARMOwnedMetaObject,
    next extensions.ClaimFunc,
) error {
    resource := obj.(*myservice.MyResource)

    // Check if claiming should be deferred
    if resource.Spec.SomeRequiredField == nil {
        log.V(Status).Info("Deferring claim until required field is set")
        // Return nil to allow reconciliation to continue
        // The resource will be claimed on the next reconciliation
        return nil
    }

    // Proceed with normal claiming
    return next(ctx, log, obj)
}
```

### Pattern 2: Custom Validation Before Claiming

```go
func (ex *ResourceExtension) Claim(
    ctx context.Context,
    log logr.Logger,
    obj genruntime.ARMOwnedMetaObject,
    next extensions.ClaimFunc,
) error {
    resource := obj.(*myservice.MyResource)

    // Validate resource state before claiming
    if err := ex.validateBeforeClaim(resource); err != nil {
        return conditions.NewReadyConditionImpactingError(
            err,
            conditions.ConditionSeverityError,
            conditions.ReasonFailed)
    }

    // Claim the resource
    err := next(ctx, log, obj)
    if err != nil {
        return err
    }

    // Perform post-claim initialization
    return ex.initializeAfterClaim(resource)
}
```

### Pattern 3: Alternative Claiming (Advanced)

In rare cases, you might completely replace the default claiming logic:

```go
func (ex *ResourceExtension) Claim(
    ctx context.Context,
    log logr.Logger,
    obj genruntime.ARMOwnedMetaObject,
    next extensions.ClaimFunc,
) error {
    resource := obj.(*myservice.MyResource)

    // Construct ARM ID using custom logic
    armId, err := ex.constructCustomArmId(ctx, resource)
    if err != nil {
        return err
    }

    // Set the ARM ID directly on the resource
    genruntime.SetResourceID(resource, armId)

    log.V(Status).Info("Claimed resource with custom ARM ID", "armId", armId)

    // Note: Not calling next() - we're completely replacing the claim logic
    return nil
}
```

**⚠️ Warning**: Completely replacing the claim logic is rarely necessary and should be done with caution. Always prefer calling `next()` and wrapping the default behavior.

## Claim Process Overview

Understanding the standard claiming process helps you know what `next()` does:

1. **Owner resolution**: Determines the parent resource
2. **ARM ID construction**: Builds the full ARM resource ID
   - Format: `/subscriptions/{sub}/resourceGroups/{rg}/providers/{provider}/{type}/{name}`
3. **Status update**: Sets `status.id` and `status.name` on the resource
4. **Annotation**: Adds the `serviceoperator.azure.com/resource-id` annotation

When you call `next()`, all of these steps happen automatically.

## Testing

When testing `Claimer` extensions:

1. **Test with default claim**: Verify that calling `next()` works correctly
2. **Test custom logic**: Ensure your pre/post claim operations execute
3. **Test error cases**: Verify proper error handling and propagation
4. **Test conditional paths**: Cover all branches in conditional claiming
5. **Verify resource state**: Check that ARM ID and status are set correctly

Example test structure:

```go
func TestMyResourceExtension_Claim(t *testing.T) {
    t.Run("successful claim with default logic", func(t *testing.T) {
        // Test calling next() succeeds
    })

    t.Run("custom validation fails", func(t *testing.T) {
        // Test error cases before claiming
    })

    t.Run("deferred claiming", func(t *testing.T) {
        // Test conditional claim deferral
    })
}
```

## Notes

- The `Claimer` extension is invoked once per resource when it doesn't have an ARM ID
- After successful claiming, the resource will have its `status.id` populated
- Claiming happens before any ARM API calls are made
- If claiming fails, the reconciliation is blocked until the issue is resolved
- Most resources do **not** need custom claiming - the default logic handles standard cases

## Related Extension Points

- [Deleter]({{< relref "deleter" >}}): Customize deletion behavior (the opposite of claiming)
- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Validate before reconciliation
- [SuccessfulCreationHandler]({{< relref "successful-creation-handler" >}}): Handle successful creation
