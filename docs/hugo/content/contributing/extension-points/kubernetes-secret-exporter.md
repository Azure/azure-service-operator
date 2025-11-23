---
title: KubernetesSecretExporter
linktitle: KubernetesSecretExporter
weight: 60
---

## Description

`KubernetesSecretExporter` allows resources to export secrets from Azure into Kubernetes Secret objects. This extension is invoked after a resource has been successfully created or updated in Azure, giving the resource the ability to retrieve sensitive data (like connection strings, keys, passwords) and make them available in Kubernetes.

The interface is called during the reconciliation process, after ARM operations succeed but before the Ready condition is marked successful. This ensures secrets are available before dependent resources can use them.

## Interface Definition

See the [KubernetesSecretExporter interface definition](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/kubernetes_secret_exporter.go) in the source code.

## Motivation

The `KubernetesSecretExporter` extension exists to handle cases where:

1. **Azure-generated secrets**: Resources that generate secrets in Azure (keys, connection strings, passwords) that need to be accessible in Kubernetes
2. **Status-based secrets**: Secret values that are part of the resource's status and should be exported to Secrets
3. **API-retrieved secrets**: Secrets that require separate API calls to retrieve (not in the resource response)
4. **Derived secrets**: Secret values that need to be computed or combined from Azure data
5. **User-controlled export**: Users can specify which secrets they want exported via `operatorSpec.secrets`

Many Azure resources generate credentials, keys, or connection strings that applications need. Rather than requiring users to manually retrieve these from Azure, ASO can export them directly to Kubernetes Secrets.

## When to Use

Implement `KubernetesSecretExporter` when:

- ✅ The resource generates secrets in Azure (keys, passwords, tokens)
- ✅ Secret values are available in the resource status
- ✅ Additional ARM API calls are needed to retrieve secrets
- ✅ Users need programmatic access to resource credentials
- ✅ Secrets need to be made available to other Kubernetes resources

Do **not** use `KubernetesSecretExporter` when:

- ❌ The resource doesn't have any secrets to export
- ❌ Secrets are already available through standard Kubernetes mechanisms
- ❌ You're exposing non-sensitive configuration (use ConfigMaps instead)
- ❌ The generator can handle it automatically (enhance the generator)

## Example: User Assigned Identity Secret Export

See the [full implementation in user_assigned_identity_extention_authorization.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/managedidentity/customizations/user_assigned_identity_extention_authorization.go).

**Key aspects of this implementation:**

1. **Early exit**: Returns nil if no secrets were requested
2. **Type safety**: Type assertions for resource and hub version
3. **Collector pattern**: Uses secrets.Collector to build Secret objects
4. **Status values**: Retrieves values from resource status
5. **User control**: Only exports secrets the user specified in operatorSpec
6. **Namespace scoping**: Secrets created in the same namespace as the resource

## Common Patterns

### Pattern 1: Export Status-Based Secrets

```go
func (ex *ResourceExtension) ExportKubernetesSecrets(
    ctx context.Context,
    obj genruntime.MetaObject,
    additionalSecrets set.Set[string],
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
    resource := obj.(*myservice.MyResource)

    // Check if secrets were requested
    if resource.Spec.OperatorSpec == nil || 
       resource.Spec.OperatorSpec.Secrets == nil {
        return nil, nil
    }

    // Create collector
    collector := secrets.NewCollector(resource.Namespace)

    // Add secrets from status
    secrets := resource.Spec.OperatorSpec.Secrets
    collector.AddValue(secrets.ConnectionString, resource.Status.ConnectionString)
    collector.AddValue(secrets.PrimaryKey, resource.Status.PrimaryKey)
    collector.AddValue(secrets.SecondaryKey, resource.Status.SecondaryKey)

    // Build Secret objects
    result, err := collector.Values()
    if err != nil {
        return nil, err
    }

    return &genruntime.KubernetesSecretExportResult{
        Objs: secrets.SliceToClientObjectSlice(result),
    }, nil
}
```

### Pattern 2: Retrieve Secrets via ARM API

Some secrets require additional API calls:

```go
func (ex *ResourceExtension) ExportKubernetesSecrets(
    ctx context.Context,
    obj genruntime.MetaObject,
    additionalSecrets set.Set[string],
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
    resource := obj.(*myservice.MyResource)

    if resource.Spec.OperatorSpec == nil || 
       resource.Spec.OperatorSpec.Secrets == nil {
        return nil, nil
    }

    // Get the resource ID for API calls
    resourceID, hasID := genruntime.GetResourceID(resource)
    if !hasID {
        return nil, fmt.Errorf("resource doesn't have an ID yet")
    }

    // Make ARM API call to retrieve keys
    var keysResponse MyResourceKeysResponse
    apiVersion := "2023-01-01"
    keysURL := fmt.Sprintf("%s/listKeys", resourceID)
    
    _, err := armClient.PostByIDWithResponse(
        ctx,
        keysURL,
        apiVersion,
        nil, // request body
        &keysResponse)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve keys: %w", err)
    }

    // Collect secrets
    collector := secrets.NewCollector(resource.Namespace)
    secrets := resource.Spec.OperatorSpec.Secrets
    collector.AddValue(secrets.Key1, keysResponse.Key1)
    collector.AddValue(secrets.Key2, keysResponse.Key2)

    result, err := collector.Values()
    if err != nil {
        return nil, err
    }

    return &genruntime.KubernetesSecretExportResult{
        Objs: secrets.SliceToClientObjectSlice(result),
    }, nil
}
```

### Pattern 3: Combine Multiple Sources

```go
func (ex *ResourceExtension) ExportKubernetesSecrets(
    ctx context.Context,
    obj genruntime.MetaObject,
    additionalSecrets set.Set[string],
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
    resource := obj.(*myservice.MyResource)

    if resource.Spec.OperatorSpec == nil || 
       resource.Spec.OperatorSpec.Secrets == nil {
        return nil, nil
    }

    collector := secrets.NewCollector(resource.Namespace)
    secrets := resource.Spec.OperatorSpec.Secrets

    // From status
    collector.AddValue(secrets.Endpoint, resource.Status.Endpoint)

    // From ARM API
    keys, err := ex.retrieveKeys(ctx, resource, armClient)
    if err != nil {
        return nil, err
    }
    collector.AddValue(secrets.PrimaryKey, keys.Primary)

    // Computed value
    connectionString := fmt.Sprintf(
        "Endpoint=%s;SharedAccessKey=%s",
        resource.Status.Endpoint,
        keys.Primary)
    collector.AddValue(secrets.ConnectionString, connectionString)

    result, err := collector.Values()
    if err != nil {
        return nil, err
    }

    return &genruntime.KubernetesSecretExportResult{
        Objs: secrets.SliceToClientObjectSlice(result),
    }, nil
}
```

### Pattern 4: Using RawSecrets for Secret Expressions

For advanced scenarios with secret expressions:

```go
func (ex *ResourceExtension) ExportKubernetesSecrets(
    ctx context.Context,
    obj genruntime.MetaObject,
    additionalSecrets set.Set[string],
    armClient *genericarmclient.GenericClient,
    log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
    resource := obj.(*myservice.MyResource)

    // Regular secret export
    collector := secrets.NewCollector(resource.Namespace)
    if resource.Spec.OperatorSpec != nil && 
       resource.Spec.OperatorSpec.Secrets != nil {
        collector.AddValue(
            resource.Spec.OperatorSpec.Secrets.PrimaryKey,
            resource.Status.PrimaryKey)
    }

    secretObjs, err := collector.Values()
    if err != nil {
        return nil, err
    }

    // Raw secrets for expressions (if requested)
    rawSecrets := make(map[string]string)
    if additionalSecrets.Contains("adminCredentials") {
        creds, err := ex.getAdminCredentials(ctx, resource, armClient)
        if err != nil {
            return nil, err
        }
        rawSecrets["adminCredentials"] = creds
    }

    return &genruntime.KubernetesSecretExportResult{
        Objs:       secrets.SliceToClientObjectSlice(secretObjs),
        RawSecrets: rawSecrets,
    }, nil
}
```

## User Specification of Secrets

Users control which secrets to export through the resource spec:

```yaml
apiVersion: managedidentity.azure.com/v1api20230131
kind: UserAssignedIdentity
metadata:
  name: my-identity
  namespace: default
spec:
  location: westus2
  owner:
    name: my-rg
  operatorSpec:
    secrets:
      clientId:
        name: identity-secret
        key: clientId
      principalId:
        name: identity-secret
        key: principalId
      tenantId:
        name: identity-secret
        key: tenantId
```

This creates a Secret named `identity-secret` with three keys.

## Secret Lifecycle

Understanding the secret export process:

1. **Resource reconciled**: ARM operations complete successfully
2. **Extension invoked**: `ExportKubernetesSecrets()` called
3. **Secrets retrieved**: Extension gathers secret values
4. **Secrets created**: Controller creates/updates Secret objects in Kubernetes
5. **Ready condition**: Resource marked as Ready
6. **Updates**: Secrets updated on each reconciliation if values change

## Error Handling

Proper error handling ensures users know when secret export fails:

```go
// Retrieval failed
return nil, fmt.Errorf("failed to retrieve keys from Azure: %w", err)

// Invalid configuration
return nil, fmt.Errorf("operatorSpec.secrets.connectionString requires a destination")

// Partial failure - log warning and continue
log.V(Warning).Info("Failed to retrieve optional secret", "error", err)
// Continue with other secrets
```

## Testing

When testing `KubernetesSecretExporter` extensions:

1. **Test no secrets requested**: Verify nil return when nothing specified
2. **Test secret creation**: Verify Secret objects are correctly created
3. **Test multiple secrets**: Verify all requested secrets are exported
4. **Test API failures**: Verify error handling for ARM call failures
5. **Test secret updates**: Verify secrets update when values change
6. **Test secret references**: Verify correct names and keys

## Security Considerations

When implementing secret export:

1. **Only export when requested**: Don't create secrets unless user specified them
2. **Use appropriate permissions**: Ensure proper RBAC for secret creation
3. **Validate destinations**: Ensure secret references are valid
4. **Handle nil values**: Don't export nil/empty secrets
5. **Log carefully**: Don't log secret values
6. **Follow naming conventions**: Use predictable, documented secret formats

## Important Notes

- **Return nil for no secrets**: If no secrets requested, return `nil, nil`
- **Secrets updated on reconciliation**: Values refresh with each reconciliation
- **Namespace scoping**: Secrets created in same namespace as resource
- **Secret format**: Follow Kubernetes Secret conventions (base64 encoding handled automatically)
- **Error impact**: Export failures prevent Ready condition
- **User control**: Only export what users explicitly request

## Related Extension Points

- [PostReconciliationChecker]({{< relref "post-reconciliation-checker" >}}): Runs at similar time
- [SuccessfulCreationHandler]({{< relref "successful-creation-handler" >}}): Also runs after creation
- [ARMResourceModifier]({{< relref "arm-resource-modifier" >}}): Runs before ARM operations

## Related Interfaces

- **KubernetesConfigExporter**: Similar interface for exporting ConfigMaps instead of Secrets
- **Secrets package**: Helper utilities for building Secret objects
- **SecretDestination**: Spec type for specifying where secrets should be exported
