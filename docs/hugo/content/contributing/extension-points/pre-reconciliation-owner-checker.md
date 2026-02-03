---
title: PreReconciliationOwnerChecker
linktitle: PreReconciliationOwnerChecker
weight: 76
---

## Description

`PreReconciliationOwnerChecker` allows resources to perform _owner_ validation or checks before ARM reconciliation begins. This extension is invoked _before_ `PreReconciliationChecker` and before sending any requests to Azure, giving resources the ability to block reconciliation based solely on the owner's state.

If this extension blocks the reconcile, ASO won't even attempt to GET the resource itself. This is critical for resources where the owner's state can completely block access to the child resource.

## Interface Definition

See the [PreReconciliationOwnerChecker interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/prereconciliation_owner_checker.go) in the source code.

**Key characteristics:**

- The `owner` parameter is **guaranteed to be non-nil** - the extension is not called if there is no owner
- The owner's status is **freshly updated from Azure** before the check is called (a GET is issued on the owner), ensuring you always have up-to-date owner state to make decisions

## Motivation

The `PreReconciliationOwnerChecker` extension exists to handle a specific class of Azure resources where:

1. **Owner state blocks all access**: The parent resource's state can prevent any operations on child resources, including GET
2. **Cannot determine child state**: You cannot query the child resource to check its state when the owner is in certain states
3. **Avoid wasted API calls**: Attempting to GET or PUT a child when the owner blocks access wastes API quota and generates errors
4. **Owner-dependent access**: Some Azure services completely shut down child resources when the parent is in maintenance, updating, or powered-off states

The most notable example is **Azure Data Explorer (Kusto)**, where you cannot even GET a database when the cluster is powered off or updating. Without this extension, the controller would repeatedly attempt to access the database, failing each time, consuming request quota, and filling ASO logs with errors.

## When to Use

Implement `PreReconciliationOwnerChecker` when:

- ✅ The owner's state can block child resource operations, causing them to fail
- ✅ You want to avoid wasting API quota on operations that will definitely fail

Do **not** use `PreReconciliationOwnerChecker` when:

- ❌ You can safely GET the resource regardless of owner state (use `PreReconciliationChecker` instead)
- ❌ You need access to the resource's current state to make the decision
- ❌ The owner's state doesn't affect access to the child
- ❌ The check can be done after retrieving the resource

## PreReconciliationOwnerChecker vs PreReconciliationChecker

These two extension points are similar, as each allows an advance check to block reconciliation attempts that are doomed to fail. However, they differ in **when** they are invoked and **what information** is available to make the decision. `PreReconciliationOwnerChecker` runs earlier and only has access to the owner, while `PreReconciliationChecker` runs later and has access to the resource itself.

Assuming both extensions are implemented, the reconciliation flow is:

1. GET updated owner status from Azure
2. Invoke `PreReconciliationOwnerChecker` to check whether reconcile should proceed based on owner state
3. If allowed, GET the resource from Azure
4. Invoke `PreReconciliationChecker` to check whether reconcile should proceed based on resource state
5. If allowed, continue with PUT/PATCH/DELETE as needed

## Example: Kusto Database Owner State Check

**Problem:** Kusto databases cannot be accessed _at all_ when the cluster is stopped or updating.

**Solution:** Check cluster state before attempting any database operations.

```go
func (ext *DatabaseExtension) PreReconcileOwnerCheck(
 ctx context.Context,
 owner genruntime.MetaObject,
 resourceResolver *resolver.Resolver,
 armClient_ *genericarmclient.GenericClient,
 log logr.Logger,
 next extensions.PreReconcileOwnerCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
 // Check to see if the owning cluster is in a state that will block us from reconciling
 if cluster, ok := owner.(*kusto.Cluster); ok {
  // If our owning *cluster* is in a state that will reject any PUT, then we should skip
  // reconciliation of the database as there's no point in even trying.
  // One way this can happen is when we reconcile the cluster, putting it into an `Updating`
  // state for a period. While it's in that state, we can't even try to reconcile the database as
  // the operation will fail with a `Conflict` error.
  // Checking the state of our owning cluster allows us to "play nice with others" and not use up
  // request quota attempting to make changes when we already know those attempts will fail.
  state := cluster.Status.ProvisioningState
  if state != nil && clusterProvisioningStateBlocksReconciliation(state) {
   return extensions.BlockReconcile(
     fmt.Sprintf("Owning cluster is in provisioning state %q", *state)),
    nil
  }
 }

 return next(ctx, owner, resourceResolver, armClient_, log)
}

```

See the [full implementation in database_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/kusto/customizations/database_extensions.go).

**Key aspects of this implementation:**

1. **Owner type assertion**: Checks if owner is a Kusto Cluster
2. **State checking**: Examines cluster's provisioning state (always up-to-date from Azure)
3. **Clear blocking messages**: Provides specific reason for blocking
4. **Helper function**: Encapsulates state-checking logic
5. **Calls next**: Proceeds when cluster is in acceptable state

## Check Results

The extension returns one of two results, or an error:

### Proceed

```go
return extensions.ProceedWithReconcile(), nil
```

- Owner state permits reconciliation
- Reconciliation will continue (GET will be attempted on the resource)
- Normal reconciliation flow proceeds

### Block

```go
return extensions.BlockReconcile("parent is updating"), nil
```

- Owner state blocks reconciliation
- No GET or other operations attempted on the child resource
- Resource requeued to try again later
- Condition set with the blocking reason

> **Note:** The owner itself has already been fetched from Azure before this check runs, ensuring the state you're checking is current.

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
2. **Owner resolution**: Owner resource identified from Kubernetes
3. **Owner refresh**: GET issued to Azure to refresh owner status **← Fresh state guaranteed**
4. **PreReconciliationOwnerChecker invoked**: Extension checks owner state **← YOU ARE HERE**
5. **Decision point**:
   - If **blocked**: Stop here, requeue, set condition
   - If **proceed**: Continue to next step
6. **GET resource**: Fetch resource from Azure (only if not blocked)
7. **PreReconciliationChecker**: Check resource state (if implemented)
8. **ARM operations**: PUT/PATCH/DELETE as needed

## Testing

When testing `PreReconciliationOwnerChecker` extensions:

1. **Test with wrong owner type**: Verify graceful handling
2. **Test blocking states**: Cover all states that should block
3. **Test proceed states**: Verify reconciliation proceeds when appropriate
4. **Test error handling**: Verify proper error returns

> **Note:** You don't need to test with nil owner - the extension is never called when the owner is nil.

## Important Notes

- **No resource access**: You do **not** receive the resource being reconciled
- **Owner only**: Make decisions based solely on owner state
- **Earlier in flow**: Runs before GET, unlike PreReconciliationChecker
- **Owner guaranteed non-nil**: The extension is not called when there is no owner
- **Fresh owner state**: Owner is fetched from Azure before check runs
- **Type assertions**: Verify owner is expected type before accessing fields
- **Call next()**: Call the next checker in the chain when checks pass
- **Clear reasons**: Provide helpful blocking messages for users
- **Logging**: Log decisions to help debugging

## Related Extension Points

- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Check resource state
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Validate after reconciliation

## Best Practices

1. **Type assert safely**: Verify owner type before accessing fields
2. **Document blocking states**: Comment which states block and why
3. **Use helper functions**: Encapsulate state-checking logic
4. **Log decisions**: Help debugging by logging why checks block
5. **Clear messages**: Users need to understand why reconciliation is blocked
6. **Call next()**: Enable check chaining
7. **Test thoroughly**: Cover all owner states and edge cases
