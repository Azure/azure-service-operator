---
title: PostReconciliationChecker
linktitle: PostReconciliationChecker
weight: 70
---

## Description

`PostReconciliationChecker` allows resources to perform validation or checks after ARM reconciliation completes successfully. This extension is invoked after Azure operations succeed but before the Ready condition is marked successful, giving resources the ability to defer the Ready status until additional conditions are met.

The interface is called at the end of the reconciliation process, after the resource has been successfully created or updated in Azure. It provides a final opportunity to verify that the resource is truly ready for use.

## Interface Definition

See the [PostReconciliationChecker interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/postreconciliation_checker.go) in the source code.

## Motivation

The `PostReconciliationChecker` extension exists to handle cases where:

1. **Async operations**: Azure resource is created but still initializing or provisioning
2. **Manual approval required**: Resources that require external approval before being considered ready
3. **Dependent state**: Resource readiness depends on the state of other Azure resources
4. **Complex readiness criteria**: Determining if a resource is "ready" requires more than just ARM success
5. **Validation**: Additional checks needed to ensure resource is in expected state
6. **Gradual rollout**: Resource created but waiting for deployment to complete

The default behavior marks resources as Ready immediately after successful ARM operations. Some resources need to wait for additional conditions before being truly ready.

## When to Use

Implement `PostReconciliationChecker` when:

- ✅ Resource continues initializing after ARM operations complete
- ✅ Manual approval or external processes must complete first
- ✅ Dependent resources need to reach certain states
- ✅ Complex validation is needed to determine readiness
- ✅ Resource provisioning happens asynchronously in Azure
- ✅ Users should wait before using the resource

Do **not** use `PostReconciliationChecker` when:

- ❌ ARM success means the resource is ready
- ❌ The check should happen before reconciliation (use PreReconciliationChecker)
- ❌ You're validating the spec (do that in webhooks)
- ❌ You're modifying the resource (use other extensions)

## Example: Private Endpoint Approval Check

See the [full implementation in private_endpoints_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/network/customizations/private_endpoints_extensions.go).

**Key aspects of this implementation:**

1. **Type assertions**: For both resource type and hub version
2. **Status inspection**: Examines resource status to determine readiness
3. **Clear failure messages**: Provides actionable information to users
4. **No ARM calls**: Uses existing status data
5. **Conditional result**: Returns success or failure based on state
6. **No error return**: Check itself succeeded, but resource not ready yet

## Common Patterns

### Pattern 1: Check Resource State

```go
func (ex *ResourceExtension) PostReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Check if resource is in expected state
    if resource.Status.ProvisioningState == nil {
        return extensions.PostReconcileCheckResultFailure(
            "provisioning state not yet available"), nil
    }

    state := *resource.Status.ProvisioningState
    if state != "Succeeded" {
        return extensions.PostReconcileCheckResultFailure(
            fmt.Sprintf("resource still provisioning: %s", state)), nil
    }

    return extensions.PostReconcileCheckResultSuccess(), nil
}
```

### Pattern 2: Query Azure for Additional State

```go
func (ex *ResourceExtension) PostReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Get resource ID for additional queries
    resourceID, hasID := genruntime.GetResourceID(resource)
    if !hasID {
        return extensions.PostReconcileCheckResultFailure(
            "resource ID not available"), nil
    }

    // Query Azure for deployment status
    deploymentReady, err := ex.checkDeploymentStatus(ctx, resourceID, armClient)
    if err != nil {
        // Check failed (not the same as check returning failure)
        return extensions.PostReconcileCheckResult{}, err
    }

    if !deploymentReady {
        return extensions.PostReconcileCheckResultFailure(
            "deployment still in progress"), nil
    }

    return extensions.PostReconcileCheckResultSuccess(), nil
}
```

### Pattern 3: Check Dependent Resources

```go
func (ex *ResourceExtension) PostReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Check if dependent resources are ready
    if resource.Spec.DependencyReference != nil {
        ready, err := ex.isDependencyReady(ctx, resource.Spec.DependencyReference, resourceResolver)
        if err != nil {
            return extensions.PostReconcileCheckResult{}, err
        }

        if !ready {
            return extensions.PostReconcileCheckResultFailure(
                "waiting for dependent resource to be ready"), nil
        }
    }

    return extensions.PostReconcileCheckResultSuccess(), nil
}
```

### Pattern 4: Timeout After Extended Wait

```go
func (ex *ResourceExtension) PostReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Check how long we've been waiting
    if resource.Status.CreationTime != nil {
        waitTime := time.Since(resource.Status.CreationTime.Time)
        if waitTime > 30*time.Minute {
            // Give up waiting, but don't fail permanently
            log.V(Warning).Info(
                "Resource taking longer than expected to be ready",
                "waitTime", waitTime)
            // Could return success here to stop blocking, or keep waiting
        }
    }

    // Perform the actual readiness check
    if !ex.isReady(resource) {
        return extensions.PostReconcileCheckResultFailure(
            "resource initialization in progress"), nil
    }

    return extensions.PostReconcileCheckResultSuccess(), nil
}
```

### Pattern 5: Call Next for Chain

```go
func (ex *ResourceExtension) PostReconcileCheck(
    ctx context.Context,
    obj genruntime.MetaObject,
    owner genruntime.MetaObject,
    resourceResolver *resolver.Resolver,
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
    next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
    resource := obj.(*myservice.MyResource)

    // Perform custom check first
    if !ex.customReadinessCheck(resource) {
        return extensions.PostReconcileCheckResultFailure(
            "custom readiness check failed"), nil
    }

    // Call next checker in the chain (if any)
    // The default implementation always returns success
    return next(ctx, obj, owner, resourceResolver, armClient, log)
}
```

## Check Results

The extension returns one of two results:

### Success

```go
return extensions.PostReconcileCheckResultSuccess(), nil
```

- Resource is ready
- Ready condition will be marked True
- Reconciliation completes successfully

### Failure

```go
return extensions.PostReconcileCheckResultFailure("reason for not ready"), nil
```

- Resource is not yet ready
- Warning condition set with the provided reason
- Reconciliation will retry later
- Resource requeued for another check

### Error

```go
return extensions.PostReconcileCheckResult{}, fmt.Errorf("check failed: %w", err)
```

- The check itself failed (couldn't determine readiness)
- Error condition set on resource
- Reconciliation will retry later

## Failure vs. Error

It's critical to distinguish between check failure and check error:

- **Failure**: "The resource is not ready yet" (expected, will retry)
  - Example: "waiting for approval", "deployment in progress"
  - Returns `PostReconcileCheckResultFailure(reason), nil`

- **Error**: "I couldn't determine if the resource is ready" (unexpected)
  - Example: "failed to query Azure", "invalid state"
  - Returns `PostReconcileCheckResult{}, error`

## Reconciliation Impact

When a post-reconciliation check fails:

1. **Condition set**: Warning condition added to resource status
2. **Reconciliation requeued**: Controller will try again later
3. **No Ready condition**: Resource not marked as Ready
4. **User visibility**: Users see the warning condition with reason

This continues until the check succeeds or the resource is deleted.

## Testing

When testing `PostReconciliationChecker` extensions:

1. **Test success case**: Verify check passes when resource is ready
2. **Test failure cases**: Cover all scenarios that should defer readiness
3. **Test error handling**: Verify proper error returns
4. **Test with real status**: Use realistic status values
5. **Test retry behavior**: Verify requeue happens correctly

## Important Notes

- **Call `next()` if appropriate**: Allows for check chaining (rarely needed)
- **Don't modify the resource**: This is for validation only
- **Be patient**: Checks may run many times before succeeding
- **Provide clear reasons**: Failure messages shown to users
- **Log appropriately**: Help debugging without noise
- **Handle nil values**: Status fields may not be populated yet
- **Consider performance**: This runs on every reconciliation

## Common Readiness Scenarios

Here are typical reasons to use post-reconciliation checks:

1. **Async provisioning**: Azure resource created but still initializing
2. **Manual approval**: External human approval required
3. **Deployment rollout**: Waiting for deployment to all regions/instances
4. **Certificate generation**: Waiting for certs to be issued
5. **DNS propagation**: Waiting for DNS changes to propagate
6. **Dependent services**: Waiting for related services to initialize

## Related Extension Points

- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Check before reconciliation
- [KubernetesSecretExporter]({{< relref "kubernetes-secret-exporter" >}}): Runs at similar time
- [SuccessfulCreationHandler]({{< relref "successful-creation-handler" >}}): Runs after initial creation
- [ARMResourceModifier]({{< relref "arm-resource-modifier" >}}): Modify before reconciliation

## Best Practices

1. **Keep checks fast**: Runs frequently, avoid expensive operations
2. **Be idempotent**: Check may run multiple times
3. **Use status fields**: Prefer examining status over ARM calls
4. **Provide context**: Clear failure messages help users
5. **Consider timeouts**: Don't wait forever for readiness
6. **Log decisions**: Help debugging by logging why checks fail/succeed
