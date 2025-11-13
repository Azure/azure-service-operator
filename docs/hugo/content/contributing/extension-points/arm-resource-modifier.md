---
title: ARMResourceModifier
linktitle: ARMResourceModifier
weight: 10
---

## Description

`ARMResourceModifier` allows resources to modify the payload that will be sent to Azure Resource Manager (ARM) immediately before it is transmitted. This extension point is invoked after the standard resource conversion and validation, but before the HTTP request is made to ARM.

The interface is invoked during PUT and PATCH operations to ARM, giving the resource an opportunity to make last-minute adjustments to the ARM payload based on runtime conditions, Azure state, or complex business logic that cannot be expressed in the generated code.

## Interface Definition

```go
type ARMResourceModifier interface {
    ModifyARMResource(
        ctx context.Context,
        armClient *genericarmclient.GenericClient,
        armObj genruntime.ARMResource,
        obj genruntime.ARMMetaObject,
        kubeClient kubeclient.Client,
        resolver *resolver.Resolver,
        log logr.Logger,
    ) (genruntime.ARMResource, error)
}
```

**Parameters:**
- `ctx`: The current operation context
- `armClient`: Client for making additional ARM API calls if needed
- `armObj`: The ARM resource representation about to be sent to Azure
- `obj`: The Kubernetes resource being reconciled
- `kubeClient`: Client for accessing the Kubernetes cluster
- `resolver`: Helper for resolving resource references
- `log`: Logger for the current operation

**Returns:**
- Modified `genruntime.ARMResource` that will be sent to ARM
- Error if modification fails (will block the ARM request)

## Motivation

The `ARMResourceModifier` extension exists to handle cases where:

1. **Azure resource state affects the payload**: Some Azure resources require different payloads based on their current state in Azure (e.g., creation vs. update)

2. **Complex conditional logic**: Business logic that depends on multiple factors and cannot be expressed declaratively in the resource schema

3. **Azure-specific quirks**: Handling special cases or undocumented Azure behavior that varies by resource type

4. **Dynamic payload construction**: Building parts of the payload at runtime based on information retrieved from Azure or Kubernetes

5. **Soft-delete scenarios**: Resources with soft-delete capabilities (like Key Vault) may need special handling to recover or purge existing resources

## When to Use

Implement `ARMResourceModifier` when:

- ✅ The resource has different creation modes or requires conditional field population
- ✅ You need to query Azure or Kubernetes state to determine the correct payload
- ✅ The resource has soft-delete and requires recovery or purge logic
- ✅ Child resources need to be included in parent payloads (e.g., VNET subnets)
- ✅ Field values must be computed at reconciliation time based on external state

Do **not** use `ARMResourceModifier` when:

- ❌ The logic could be handled by field defaults or validation
- ❌ The change should apply to all resources (modify the generator instead)
- ❌ Simple field transformations (use conversion functions)
- ❌ The issue is a generator bug (fix the generator)

## Example: Key Vault CreateMode Handling

The Key Vault resource uses `ARMResourceModifier` to handle different creation modes based on whether a soft-deleted vault exists:

```go
var _ extensions.ARMResourceModifier = &VaultExtension{}

func (ex *VaultExtension) ModifyARMResource(
    ctx context.Context,
    armClient *genericarmclient.GenericClient,
    armObj genruntime.ARMResource,
    obj genruntime.ARMMetaObject,
    kubeClient kubeclient.Client,
    resolver *resolver.Resolver,
    log logr.Logger,
) (genruntime.ARMResource, error) {
    // Type assert to the specific resource type
    kv, ok := obj.(*keyvault.Vault)
    if !ok {
        return nil, eris.Errorf(
            "Cannot run VaultExtension.ModifyARMResource() with unexpected resource type %T",
            obj)
    }

    // Type assert hub version to catch breaking changes
    var _ conversion.Hub = kv

    // Exit early if no special handling needed
    if kv.Spec.Properties == nil || kv.Spec.Properties.CreateMode == nil {
        return armObj, nil
    }

    // Get resource context
    id, err := ex.getOwner(ctx, kv, resolver)
    if err != nil {
        return nil, eris.Wrap(err, "failed to get and parse resource ID from KeyVault owner")
    }

    // Create Azure SDK client to check for soft-deleted vaults
    vc, err := armkeyvault.NewVaultsClient(id.SubscriptionID, armClient.Creds(), armClient.ClientOptions())
    if err != nil {
        return nil, eris.Wrap(err, "failed to create new VaultsClient")
    }

    // Determine the correct create mode based on Azure state
    createMode := *kv.Spec.Properties.CreateMode
    if createMode == CreateMode_CreateOrRecover {
        // Check if soft-deleted vault exists and adjust createMode accordingly
        createMode, err = ex.handleCreateOrRecover(ctx, kv, vc, id, log)
        if err != nil {
            return nil, eris.Wrapf(err, "error checking for existence of soft-deleted KeyVault")
        }
    }

    if createMode == CreateMode_PurgeThenCreate {
        // Purge the soft-deleted vault before creating
        err = ex.handlePurgeThenCreate(ctx, kv, vc, log)
        if err != nil {
            return nil, eris.Wrapf(err, "error purging soft-deleted KeyVault")
        }
        createMode = CreateMode_Default
    }

    // Modify the ARM payload with the determined createMode
    spec := armObj.Spec()
    err = reflecthelpers.SetProperty(spec, "Properties.CreateMode", &createMode)
    if err != nil {
        return nil, eris.Wrapf(err, "error setting CreateMode to %s", createMode)
    }

    return armObj, nil
}
```

**Key aspects of this example:**

1. **Type assertions**: Both for the resource type and hub version
2. **Early exit**: Returns original payload if no modification needed
3. **External queries**: Checks Azure for soft-deleted vault state
4. **Conditional logic**: Different behavior based on createMode and vault state
5. **Payload modification**: Uses reflection to safely modify the ARM payload
6. **Error handling**: Returns detailed errors that prevent ARM submission

## Common Patterns

### Pattern 1: Querying Current Azure State

```go
func (ex *ResourceExtension) ModifyARMResource(...) (genruntime.ARMResource, error) {
    // Get the resource ID
    resourceID, hasResourceID := genruntime.GetResourceID(obj)
    if !hasResourceID {
        // Not yet claimed, return unmodified
        return armObj, nil
    }

    // Query current state from Azure
    raw := make(map[string]any)
    _, err = armClient.GetByID(ctx, resourceID, apiVersion, &raw)
    if err != nil {
        // Handle NotFound appropriately
        var responseError *azcore.ResponseError
        if eris.As(err, &responseError) && responseError.StatusCode == http.StatusNotFound {
            return armObj, nil
        }
        return nil, err
    }

    // Use Azure state to modify payload
    // ...
    
    return armObj, nil
}
```

### Pattern 2: Including Child Resources

```go
func (ex *ResourceExtension) ModifyARMResource(...) (genruntime.ARMResource, error) {
    // Get existing child resources from Azure
    children, err := getRawChildCollection(raw, "subnets")
    if err != nil {
        return nil, eris.Wrapf(err, "failed to get child resources")
    }

    // Merge with desired state in payload
    err = setChildCollection(armObj.Spec(), children, "Subnets")
    if err != nil {
        return nil, eris.Wrapf(err, "failed to set child resources")
    }

    return armObj, nil
}
```

### Pattern 3: Conditional Field Population

```go
func (ex *ResourceExtension) ModifyARMResource(...) (genruntime.ARMResource, error) {
    // Determine value based on runtime conditions
    value, err := ex.computeValue(ctx, obj, armClient)
    if err != nil {
        return nil, err
    }

    // Set the computed value on the ARM resource
    spec := armObj.Spec()
    err = reflecthelpers.SetProperty(spec, "Properties.FieldName", value)
    if err != nil {
        return nil, err
    }

    return armObj, nil
}
```

## Testing

When testing `ARMResourceModifier` extensions:

1. **Test all code paths**: Cover each conditional branch
2. **Mock ARM responses**: Use envtest with recorded responses
3. **Verify payload changes**: Assert that the returned armObj has expected modifications
4. **Test error handling**: Ensure errors are properly classified and reported
5. **Test edge cases**: No ResourceID, resource not found, nil fields, etc.

## Related Extension Points

- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): For validation before ARM operations
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): For validation after ARM operations
- [ErrorClassifier]({{< relref "error-classifier" >}}): For handling errors from modified requests
