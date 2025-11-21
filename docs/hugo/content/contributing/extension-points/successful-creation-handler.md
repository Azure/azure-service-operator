---
title: SuccessfulCreationHandler
linktitle: SuccessfulCreationHandler
weight: 80
---

## Description

`SuccessfulCreationHandler` allows resources to perform custom logic immediately after they are successfully created in Azure for the first time. This extension is invoked once after the initial ARM PUT operation succeeds, giving resources the opportunity to capture information, set fields, or perform initialization that depends on the Azure-assigned resource ID.

The interface is called after the resource exists in Azure and has been assigned an ID, but before subsequent reconciliations. It runs exactly once per resource, during its initial creation.

## Interface Definition

See the [SuccessfulCreationHandler interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/successful_resource_modifier.go) in the source code.


## Motivation

The `SuccessfulCreationHandler` extension exists to handle cases where:

1. **Derived IDs**: Some resources need to compute or override their resource ID based on Azure's response

2. **Child resource IDs**: Parent resources may need to set special IDs for child resources to reference

3. **Post-creation initialization**: One-time setup that can only happen after the resource exists in Azure

4. **Status field initialization**: Setting status fields that depend on the Azure resource ID

5. **Special ARM ID handling**: Resources with non-standard ARM ID structures that need custom handling

Most resources receive their ARM ID through the standard process. Some resources have special requirements that need custom handling when the resource is first created.

## When to Use

Implement `SuccessfulCreationHandler` when:

- ✅ Resource ID needs custom computation after creation
- ✅ Child resources need special ID references set on the parent
- ✅ One-time initialization required after Azure creation
- ✅ Resource has non-standard ARM ID structure
- ✅ Status fields need to be set based on the created resource

Do **not** use `SuccessfulCreationHandler` when:

- ❌ The standard resource ID handling works correctly
- ❌ The logic should run on every reconciliation (use other extensions)
- ❌ You're modifying the spec (that should be done elsewhere)
- ❌ The initialization doesn't depend on the resource existing in Azure

## Example: Subscription Alias ID Override

See the [full implementation in alias_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/subscription/customizations/alias_extensions.go).

**Key aspects of this implementation:**

1. **Type assertions**: For both resource type and hub version
2. **Status access**: Retrieves information from the populated status
3. **ID override**: Sets a custom child resource ID reference
4. **Validation**: Checks that required status fields are present
5. **Error handling**: Returns error if required data missing


## Common Patterns

### Pattern 1: Set Computed Status Fields

```go
func (ex *ResourceExtension) Success(obj genruntime.ARMMetaObject) error {
    resource := obj.(*myservice.MyResource)

    // Type assert hub version
    var _ conversion.Hub = resource

    // Get the Azure resource ID
    resourceID, hasID := genruntime.GetResourceID(resource)
    if !hasID {
        return fmt.Errorf("resource ID not set after creation")
    }

    // Compute derived values from the resource ID
    // For example, extract subscription ID, resource group, etc.
    parsed, err := arm.ParseResourceID(resourceID)
    if err != nil {
        return fmt.Errorf("failed to parse resource ID: %w", err)
    }

    // Set status fields that depend on the ID
    resource.Status.SubscriptionID = &parsed.SubscriptionID
    resource.Status.ResourceGroup = &parsed.ResourceGroupName

    return nil
}
```

### Pattern 2: Initialize One-Time Configuration

```go
func (ex *ResourceExtension) Success(obj genruntime.ARMMetaObject) error {
    resource := obj.(*myservice.MyResource)
    var _ conversion.Hub = resource

    // Perform one-time initialization that depends on the resource existing
    if err := ex.initializePostCreation(resource); err != nil {
        return fmt.Errorf("post-creation initialization failed: %w", err)
    }

    return nil
}
```

### Pattern 3: Set Child Resource Reference

```go
func (ex *ResourceExtension) Success(obj genruntime.ARMMetaObject) error {
    resource := obj.(*myservice.MyResource)
    var _ conversion.Hub = resource

    // Extract a specific value from status to use as child resource ID
    if resource.Status.SpecialID == nil {
        return fmt.Errorf("special ID not populated in status")
    }

    // Set this as the reference ID for child resources
    specialID := *resource.Status.SpecialID
    genruntime.SetChildResourceIDOverride(
        resource,
        fmt.Sprintf("/subscriptions/.../providers/.../resources/%s", specialID))

    return nil
}
```

### Pattern 4: Validate Creation Result

```go
func (ex *ResourceExtension) Success(obj genruntime.ARMMetaObject) error {
    resource := obj.(*myservice.MyResource)
    var _ conversion.Hub = resource

    // Verify that the created resource has expected properties
    if resource.Status.Endpoint == nil {
        return fmt.Errorf("endpoint not populated after creation")
    }

    if resource.Status.State == nil || *resource.Status.State != "Active" {
        return fmt.Errorf("resource not in expected state after creation")
    }

    return nil
}
```

### Pattern 5: Record Creation Metadata

```go
func (ex *ResourceExtension) Success(obj genruntime.ARMMetaObject) error {
    resource := obj.(*myservice.MyResource)
    var _ conversion.Hub = resource

    // Record when the resource was created
    now := metav1.Now()
    resource.Status.CreatedAt = &now

    // Record the original configuration
    resource.Status.InitialSKU = resource.Spec.SKU

    return nil
}
```

## Invocation Timing

Understanding when this extension runs:

1. **Resource created in Kubernetes**: User applies YAML
2. **Resource claimed**: Gets Azure resource ID assigned
3. **Initial ARM PUT**: First creation request sent to Azure
4. **ARM returns success**: Resource now exists in Azure
5. **Status updated**: Resource status populated from Azure response
6. **SuccessfulCreationHandler called**: Extension runs **once**
7. **Ready condition set**: Resource marked as Ready (if no errors)

Subsequent reconciliations do **not** trigger this extension again.

## Success vs. Update

It's important to note:
- **Success handler**: Runs once after initial creation
- **Updates**: Do not trigger the success handler
- **Recreation**: If a resource is deleted and recreated, the handler runs again

## Error Handling

Errors from the success handler are significant:

```go
// Success handling failed
return fmt.Errorf("failed to initialize: %w", err)
```

If the success handler returns an error:
1. The error is recorded in conditions
2. The Ready condition is not set
3. Reconciliation will retry
4. Handler may be called multiple times until it succeeds

Make sure your success handler is **idempotent** - safe to call multiple times.

## Testing

When testing `SuccessfulCreationHandler` extensions:

1. **Test successful handling**: Verify handler succeeds with good data
2. **Test with missing data**: Verify error handling for missing status fields
3. **Test ID override**: Verify child resource IDs set correctly
4. **Test idempotency**: Verify multiple calls are safe
5. **Test status modifications**: Verify status fields set correctly


## Common Scenarios

Here are typical reasons to use success handlers:

1. **ID Override**: Setting custom child resource IDs (subscriptions, managed resources)
2. **Derived Status**: Computing status fields from the resource ID or response
3. **One-Time Setup**: Configuration that only makes sense after creation
4. **State Initialization**: Setting initial state based on Azure response
5. **Reference Setup**: Establishing references to related resources

## Important Notes

- **Runs once per resource**: After initial creation, not on updates
- **Status is populated**: Resource status contains Azure response data
- **Resource has ID**: Azure resource ID is available via `GetResourceID`
- **Be idempotent**: Handler may be called multiple times if it errors
- **Don't modify spec**: This is for status and metadata only
- **Return errors carefully**: Errors prevent Ready condition
- **Type assert hub**: Catch breaking changes with hub version assertion

## Idempotency Requirement

Since the handler might be called multiple times (if errors occur), ensure your handler is idempotent:

```go
// Good - idempotent
func (ex *ResourceExtension) Success(obj genruntime.ARMMetaObject) error {
    resource := obj.(*myservice.MyResource)
    
    // Set a value (safe to do multiple times)
    computed := computeValue(resource)
    resource.Status.ComputedField = &computed
    
    return nil
}

// Bad - not idempotent
func (ex *ResourceExtension) Success(obj genruntime.ARMMetaObject) error {
    resource := obj.(*myservice.MyResource)
    
    // Append to a list (would grow on each call)
    resource.Status.History = append(resource.Status.History, "Created")
    
    return nil
}
```

## Alternative Approaches

Before implementing a success handler, consider:

1. **Status fields**: Can the generator populate this from Azure?
2. **Conversion functions**: Can this be handled in conversion?
3. **Controller logic**: Should this be in the generic controller?
4. **ARMResourceModifier**: Would modifying the request work better?

Use `SuccessfulCreationHandler` when the logic truly requires:
- The resource to exist in Azure first
- One-time execution after creation
- Custom ID handling that can't be generated

## Related Extension Points

- [ARMResourceModifier]({{< relref "arm-resource-modifier" >}}): Modifies the creation request
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Runs after every reconciliation
- [KubernetesSecretExporter]({{< relref "kubernetes-secret-exporter" >}}): Also runs after successful creation

## Best Practices

1. **Keep it simple**: Only do what's necessary
2. **Validate status**: Check that required status fields are populated
3. **Be idempotent**: Handle multiple calls gracefully
4. **Clear errors**: Return descriptive error messages
5. **Don't call ARM**: This is for processing what Azure returned, not making new requests
6. **Type assert hub**: Catch breaking changes early
7. **Document why**: Comment why the success handler is needed
