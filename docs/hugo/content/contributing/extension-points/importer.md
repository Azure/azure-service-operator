---
title: Importer
linktitle: Importer
weight: 50
---

## Description

`Importer` allows resources to customize their behavior during the import process when using `asoctl import`. This extension is invoked when importing existing Azure resources into Kubernetes as ASO resources, giving resources the ability to skip import, modify the imported resource, or perform additional validation.

The interface is called during the import workflow, after retrieving the resource from Azure but before writing it to Kubernetes. This allows resources to filter out unwanted resources or adjust the imported representation.

## Interface Definition

See the [Importer interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/importer.go) in the source code.

## Motivation

The `Importer` extension exists to handle cases where:

1. **System-managed resources**: Resources that are created/managed by Azure and shouldn't be imported
2. **Default values**: Resources that only have default settings and don't need to be managed
3. **Read-only resources**: Resources that can't be modified after creation
4. **Filtering criteria**: Resources that don't meet certain criteria for management
5. **Import validation**: Resources that need validation before allowing import
6. **Resource transformation**: Adjusting the imported resource to fit Kubernetes conventions

Many Azure services automatically create child resources or configurations. These often shouldn't be managed by the operator as they're managed by Azure.

## When to Use

Implement `Importer` when:

- ✅ System-managed resources should be excluded from import
- ✅ Default/empty configurations don't need management
- ✅ Read-only resources can't be managed via ASO
- ✅ Import should be conditional based on resource properties
- ✅ Imported resources need validation or transformation
- ✅ Certain resource states should prevent import

Do **not** use `Importer` when:

- ❌ All resources of the type should be imported
- ❌ The filtering logic applies to all resources (add to asoctl)
- ❌ You want to modify how resources are retrieved from Azure (that's not import)
- ❌ You're trying to fix generator issues (fix the generator)

## Example: MySQL Configuration Import Filtering

See the [full implementation in flexible_servers_configuration_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/dbformysql/customizations/flexible_servers_configuration_extensions.go).

**Key aspects of this implementation:**

1. **Chain pattern**: Calls `next()` first to get the imported resource
2. **Type checking**: Safely handles being called on wrong resource type
3. **Multiple filters**: Several conditions that trigger skip
4. **Clear reasons**: Each skip includes explanation for user feedback
5. **Default behavior**: Returns original result when no filters apply

## Common Patterns

### Pattern 1: Skip System-Managed Resources

```go
func (ex *ResourceExtension) Import(
    ctx context.Context,
    rsrc genruntime.ImportableResource,
    owner *genruntime.ResourceReference,
    next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
    result, err := next(ctx, rsrc, owner)
    if err != nil {
        return extensions.ImportResult{}, err
    }

    resource := rsrc.(*myservice.MyResource)

    // Skip resources managed by the system
    if resource.Spec.ManagedBy != nil && *resource.Spec.ManagedBy == "System" {
        return extensions.ImportSkipped("system-managed resources shouldn't be imported"), nil
    }

    return result, nil
}
```

### Pattern 2: Validate Before Import

```go
func (ex *ResourceExtension) Import(
    ctx context.Context,
    rsrc genruntime.ImportableResource,
    owner *genruntime.ResourceReference,
    next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
    resource := rsrc.(*myservice.MyResource)

    // Validate resource before importing
    if err := ex.validateForImport(resource); err != nil {
        return extensions.ImportResult{}, eris.Wrapf(err,
            "resource %s failed import validation", resource.Name)
    }

    // Proceed with import
    return next(ctx, rsrc, owner)
}
```

### Pattern 3: Skip Based on Property Values

```go
func (ex *ResourceExtension) Import(
    ctx context.Context,
    rsrc genruntime.ImportableResource,
    owner *genruntime.ResourceReference,
    next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
    result, err := next(ctx, rsrc, owner)
    if err != nil {
        return extensions.ImportResult{}, err
    }

    resource := rsrc.(*myservice.MyResource)

    // Skip resources in certain states
    if resource.Status.ProvisioningState != nil {
        state := *resource.Status.ProvisioningState
        if state == "Deleting" || state == "Failed" {
            return extensions.ImportSkipped(
                fmt.Sprintf("resource in %s state shouldn't be imported", state)), nil
        }
    }

    return result, nil
}
```

### Pattern 4: Transform Before Import

```go
func (ex *ResourceExtension) Import(
    ctx context.Context,
    rsrc genruntime.ImportableResource,
    owner *genruntime.ResourceReference,
    next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
    resource := rsrc.(*myservice.MyResource)

    // Modify the resource before importing
    // For example, clear fields that shouldn't be managed
    if resource.Spec.AutoGeneratedField != nil {
        resource.Spec.AutoGeneratedField = nil
    }

    // Proceed with import of modified resource
    return next(ctx, rsrc, owner)
}
```

### Pattern 5: Conditional Import Based on Tags

```go
func (ex *ResourceExtension) Import(
    ctx context.Context,
    rsrc genruntime.ImportableResource,
    owner *genruntime.ResourceReference,
    next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
    result, err := next(ctx, rsrc, owner)
    if err != nil {
        return extensions.ImportResult{}, err
    }

    resource := rsrc.(*myservice.MyResource)

    // Skip resources with specific tags
    if resource.Spec.Tags != nil {
        if managed, ok := resource.Spec.Tags["managedBy"]; ok && managed == "external" {
            return extensions.ImportSkipped("resource tagged as externally managed"), nil
        }
    }

    return result, nil
}
```

## Import Workflow

Understanding the import process:

1. **Resource discovery**: `asoctl import` queries Azure for resources
2. **Resource retrieval**: Full resource details fetched from Azure
3. **Conversion**: Azure representation converted to ASO format
4. **Importer invoked**: Custom `Import()` method called
5. **Skip or continue**: Based on result, resource is written to YAML or skipped
6. **User notification**: Skipped resources reported to user with reasons

## Skip vs. Error

It's important to distinguish between skipping and errors:

- **Skip (`ImportSkipped`)**: Resource should not be imported, but import continues
  - Used for: System resources, defaults, read-only items
  - Effect: Resource excluded from output, reason logged
  - User sees: "Skipped N resources" with reasons

- **Error (`return error`)**: Import process failed
  - Used for: Validation failures, unexpected states, bugs
  - Effect: Import may abort or continue based on error handling
  - User sees: Error message, stack trace

## Testing

When testing `Importer` extensions:

1. **Test successful import**: Verify normal resources import correctly
2. **Test skip conditions**: Cover all skip scenarios with reasons
3. **Test error cases**: Verify proper error handling
4. **Test transformation**: If modifying resources, verify changes
5. **Test type safety**: Ensure handling of unexpected types

## User Experience

When resources are skipped during import, users see output like:

```bash
$ asoctl import azure-resource myserver --output yaml
Imported 12 resources
Skipped 5 resources:
  - myserver-config-1: system-defaults don't need to be imported
  - myserver-config-2: readonly configuration can't be set
  - myserver-config-3: default value is the same as the current value
  ...
```

Clear, informative skip reasons help users understand why certain resources weren't imported.

## Important Notes

- **Call `next()` first**: Allows default import logic to run
- **Provide clear reasons**: Skip reasons are shown to users
- **Be consistent**: Similar resources should skip for similar reasons
- **Don't skip too broadly**: Only skip resources that truly shouldn't be managed
- **Document skip logic**: Comment why specific conditions trigger skips
- **Test with real imports**: Verify skips work in practice with `asoctl import`

## Related Extension Points

- [PreReconciliationChecker]({{< relref "pre-reconciliation-checker" >}}): Validate before reconciliation
- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Validate after reconciliation
- [ARMResourceModifier]({{< relref "arm-resource-modifier" >}}): Modify before sending to Azure

## Related Tools

- **asoctl import**: The CLI tool that uses the Importer extension
- **Resource specifications**: Define what makes a resource importable
- **Generator**: Creates the base ImportableResource implementation
