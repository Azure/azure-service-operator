---
title: ARMResourceModifier
linktitle: ARMResourceModifier
weight: 10
---

## Description

`ARMResourceModifier` allows resources to modify the payload that will be sent to Azure Resource Manager (ARM) immediately before it is transmitted. This extension point is invoked after the standard resource conversion and validation, but before the HTTP request is made to ARM.

The interface is invoked during PUT and PATCH operations to ARM, giving the resource an opportunity to make last-minute adjustments to the ARM payload based on runtime conditions, Azure state, or complex business logic that cannot be expressed in the generated code.

## Interface Definition

See the [ARMResourceModifier interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/arm_resource_modifier.go) in the source code.

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

The Key Vault resource uses `ARMResourceModifier` to handle different creation modes based on whether a soft-deleted vault exists.

See the [full implementation in vault_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/keyvault/customizations/vault_extensions.go).

**Key aspects of this implementation:**

1. **Type assertions**: Both for the resource type and hub version
2. **Early exit**: Returns original payload if no modification needed
3. **External queries**: Checks Azure for soft-deleted vault state
4. **Conditional logic**: Different behavior based on createMode and vault state
5. **Payload modification**: Uses reflection to safely modify the ARM payload
6. **Error handling**: Returns detailed errors that prevent ARM submission

## Common Patterns

### Pattern 1: Querying Current Azure State

See implementation in [virtual_network_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/network/customizations/virtual_network_extensions.go).

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

See implementation in [virtual_network_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/network/customizations/virtual_network_extensions.go).

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
