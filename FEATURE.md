# Entra Application Support

Implementation plan for Entra Application as a handwritten resource, following the pattern established by SecurityGroup.

Ref: #5217

## Overview

Add a new `Application` CRD under the `entra.azure.com` API group (v1), allowing users to manage Entra ID Application registrations through Kubernetes. The implementation mirrors the existing `SecurityGroup` resource: handwritten types, a dedicated reconciler using the Microsoft Graph SDK (`models.Applicationable`), webhooks for defaulting and validation, and ConfigMap/Secret export of key identifiers (AppId, ObjectId).

## Files to Create

### 1. API Types — `v2/api/entra/v1/application_types.go`

Define the following types, following the SecurityGroup pattern:

- **`Application`** — top-level resource struct
  - Embeds `metav1.TypeMeta`, `metav1.ObjectMeta`
  - Has `Spec ApplicationSpec` and `Status ApplicationStatus`
  - Implements `conditions.Conditioner` (GetConditions/SetConditions)
  - Implements `conversion.Hub` (marker)
  - Kubebuilder markers: `+kubebuilder:object:root=true`, `+kubebuilder:resource:categories={azure,entra}`, `+kubebuilder:subresource:status`, `+kubebuilder:storageversion`, printer columns for Ready/Severity/Reason/Message
  - RBAC markers for `entra.azure.com` group, resource `applications`

- **`ApplicationList`** — standard list type

- **`ApplicationSpec`** — user-configurable fields
  - `DisplayName *string` (required) — display name for the app
  - `Description *string` — free text description
  - `SignInAudience *SignInAudience` — enum: `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount`, `PersonalMicrosoftAccount`
  - `IdentifierUris []string` — application ID URIs
  - `Web *WebApplication` — redirect URIs and implicit grant settings for web apps
  - `Spa *SpaApplication` — redirect URIs for single-page apps
  - `PublicClient *PublicClientApplication` — redirect URIs for desktop/mobile apps
  - `Tags []string` — custom identification strings
  - `IsFallbackPublicClient *bool` — fallback application type
  - `GroupMembershipClaims *string` — groups claim configuration
  - `OperatorSpec *ApplicationOperatorSpec` — operator-specific configuration
  - `OriginalVersion() string` method

  Sub-types as needed:
  - `WebApplication` — `RedirectUris []string`, `ImplicitGrantSettings *ImplicitGrantSettings`
  - `ImplicitGrantSettings` — `EnableIdTokenIssuance *bool`, `EnableAccessTokenIssuance *bool`
  - `SpaApplication` — `RedirectUris []string`
  - `PublicClientApplication` — `RedirectUris []string`
  - `SignInAudience` — kubebuilder enum type

  Conversion methods:
  - `AssignToApplication(model models.Applicationable)` — maps spec fields to Microsoft Graph Application model
  - These should use the existing `models.Applicationable` interface from `msgraph-sdk-go`

- **`ApplicationStatus`** — observed state
  - `EntraID *string` — the object ID of the application
  - `AppId *string` — the application (client) ID assigned by Entra
  - `DisplayName *string`
  - `Conditions []conditions.Condition`
  - `AssignFromApplication(model models.Applicationable)` method

- **`ApplicationOperatorSpec`**
  - `CreationMode *CreationMode` — reuse existing `CreationMode` from `creation_mode.go`
  - `ConfigMaps *ApplicationOperatorConfigMaps`
  - `CreationAllowed() bool` and `AdoptionAllowed() bool` methods (same pattern as SecurityGroup)

- **`ApplicationOperatorConfigMaps`**
  - `EntraID *genruntime.ConfigMapDestination` — object ID
  - `AppId *genruntime.ConfigMapDestination` — the application (client) ID

Register types in `init()` via `SchemeBuilder.Register`.

### 2. Webhook — `v2/api/entra/v1/webhook/application_webhook_types.go`

- **`Application_Webhook`** struct implementing `webhook.CustomDefaulter` and `webhook.CustomValidator`
- Mutating webhook path: `/mutate-entra-azure-com-v1-application`
- Validating webhook path: `/validate-entra-azure-com-v1-application`
- `Default()` — ensures OperatorSpec exists, defaults CreationMode to `AdoptOrCreate`
- `ValidateCreate/ValidateUpdate/ValidateDelete` — skeleton implementations with the same delegation pattern as SecurityGroup

### 3. Reconciler — `v2/internal/reconcilers/entra/entra_application_reconciler.go`

- **`EntraApplicationReconciler`** struct with same fields as SecurityGroup reconciler:
  - `reconcilers.ReconcilerCommon`
  - `ResourceResolver`, `CredentialProvider`, `Config`, `EntraClientFactory`
- Implements `genruntime.Reconciler`
- Constructor: `NewEntraApplicationReconciler(...)`

Key methods (mirroring SecurityGroup reconciler):

| Method | Behavior |
|--------|----------|
| `CreateOrUpdate` | Check annotation for existing EntraID → update; else try adopt by DisplayName → update; else create new |
| `Delete` | Get EntraID from annotation → call `client.Applications().ByApplicationId(id).Delete()` → handle 404 gracefully |
| `UpdateStatus` | Load application by ID → `status.AssignFromApplication(app)` → save ConfigMaps |
| `Claim` | No-op (same as SecurityGroup) |

Helper methods:
- `update(ctx, id, app, log)` — PATCH via `client.Applications().ByApplicationId(id).Patch(ctx, model, nil)`
- `tryAdopt(ctx, app, log)` — search by displayName using OData filter: `displayName eq '{name}'` via `client.Applications().Get(ctx, options)`
- `create(ctx, app, log)` — POST via `client.Applications().Post(ctx, model, nil)`
- `loadApplicationByID(ctx, id, client)` — GET via `client.Applications().ByApplicationId(id).Get(ctx, nil)`
- `loadApplicationsByDisplayName(ctx, displayName, client)` — GET with OData filter
- `saveAssociatedKubernetesResources(ctx, app, log)` — export EntraID and AppId to ConfigMaps
- `isNotFound(err)` — can reuse the existing `isNotFound` from SecurityGroup reconciler (or extract to shared helper)
- `asApplication(obj)` — type assertion helper

### 4. Shared Helpers (Refactoring)

The SecurityGroup reconciler has a `TODO` comment about extracting common code. Consider extracting these shared utilities into the `entra` package:

- `isNotFound(err) bool` — already present on SecurityGroup reconciler, identical logic needed for Application
- `saveAssociatedKubernetesResources` could potentially be generalized, but initial implementation can duplicate

This refactoring is optional for the initial implementation; it can be done as a follow-up.

## Files to Modify

### 5. Controller Registration — `v2/internal/controllers/controller_resources.go`

In `GetKnownStorageTypes()`, add a `registration.StorageType` entry:
```go
knownStorageTypes = append(
    knownStorageTypes,
    &registration.StorageType{
        Obj:  &entrav1.Application{},
        Name: "entra_application",
        Reconciler: entrareconciler.NewEntraApplicationReconciler(
            clients.KubeClient,
            clients.EntraConnectionFactory,
            resourceResolver,
            positiveConditions,
            options.Config),
        Predicate: makeStandardPredicate(),
        Indexes:   []registration.Index{},
        Watches:   []registration.Watch{},
    })
```

In `GetKnownTypes()`, add a `registration.KnownType` entry for webhook registration:
```go
knownTypes = append(
    knownTypes,
    &registration.KnownType{
        Obj:       &entrav1.Application{},
        Defaulter: &entrav1webhook.Application_Webhook{},
        Validator: &entrav1webhook.Application_Webhook{},
    })
```

### 6. Handcrafted Type Configuration — `v2/azure-arm.yaml`

Add under the existing `entra` section alongside SecurityGroup:
```yaml
entra:
  v1: # Handcrafted
    Application:
      $supportedFrom: <next-release-version>
    SecurityGroup:
      $supportedFrom: v2.14.0
```


## Files to Generate

### 8. CRD — `v2/config/crd/generated/bases/entra.azure.com_applications.yaml`

Generated by running `controller-gen` (via `./hack/tools/task controller:generate-types`). The kubebuilder markers on the Application types will produce the CRD automatically.

### 9. Webhook Configuration

Generated by `controller-gen` from the webhook markers:
- `v2/config/webhook/generated/bases/entra-mwh.yaml` — add mutating webhook entry for Application
- `v2/config/webhook/generated/bases/entra-vwh.yaml` — add validating webhook entry for Application

### 10. DeepCopy — `v2/api/entra/v1/zz_generated.deepcopy.go`

Generated by `controller-gen`. The `+kubebuilder:object:root=true` markers will trigger generation of `DeepCopyObject()` and `DeepCopyInto()` for `Application`, `ApplicationList`, and all sub-types.

## Additional Files

### 11. Sample — `v2/samples/entra/v1/v1_application.yaml`

```yaml
apiVersion: entra.azure.com/v1
kind: Application
metadata:
  name: aso-sample-application
  namespace: default
spec:
  displayName: ASO Sample Application
  signInAudience: AzureADMyOrg
  operatorSpec:
    configmaps:
      appId:
        key: appId
        name: application-configmap
      entraID:
        key: entraID
        name: application-configmap
```

### 12. Integration Test — `v2/internal/controllers/entra_application_v1_test.go`

Test covering the CRUD lifecycle:
1. **Create** — create Application resource with ConfigMap export, wait for Ready
2. **Read** — verify ConfigMap contains correct AppId and EntraID
3. **Update** — patch DisplayName, verify status reflects update
4. **Delete** — delete resource, verify cleanup

## Microsoft Graph API Endpoints

The reconciler will use these Microsoft Graph v1.0 endpoints via the `msgraph-sdk-go` client:

| Operation | SDK Call |
|-----------|---------|
| Create | `client.Applications().Post(ctx, model, nil)` |
| Get by ID | `client.Applications().ByApplicationId(id).Get(ctx, nil)` |
| Update | `client.Applications().ByApplicationId(id).Patch(ctx, model, nil)` |
| Delete | `client.Applications().ByApplicationId(id).Delete(ctx, nil)` |
| List/Filter | `client.Applications().Get(ctx, &options)` with OData `$filter=displayName eq '...'` |

## Implementation Order

1. API types (`application_types.go`) + run `controller-gen` to generate deepcopy and CRD
2. Webhook (`application_webhook_types.go`)
3. Reconciler (`entra_application_reconciler.go`)
4. Controller registration (`controller_resources.go`)
5. Configuration (`azure-arm.yaml`)
6. Sample YAML
7. Integration test
8. Documentation
9. Run full validation: `./hack/tools/task quick-checks && ./hack/tools/task controller:build && ./hack/tools/task controller:test`

## Key Design Decisions

- **Reuse `CreationMode`**: The existing `creation_mode.go` in `v2/api/entra/v1/` is resource-agnostic and can be reused directly by Application.
- **Reuse `EntraClientCache`/`Connection`**: The existing Entra client infrastructure works for any Microsoft Graph resource — no changes needed.
- **Adopt-by-DisplayName**: Same pattern as SecurityGroup — search by `displayName` filter, error if multiple matches found.
- **Minimal spec fields**: Start with the most commonly needed fields (DisplayName, SignInAudience, IdentifierUris, Web/SPA/PublicClient redirect URIs). Additional fields (AppRoles, RequiredResourceAccess, OptionalClaims, etc.) can be added incrementally in follow-up PRs.
- **ConfigMap export**: Export both `EntraID` (object ID) and `AppId` (client ID), since both are commonly needed by downstream consumers.
