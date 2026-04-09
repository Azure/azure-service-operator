# Migrate `$isSecret` to `$secret` with `Secrecy` type

**Date:** 2026-04-09
**Status:** Approved

## Goal

Replace the boolean `$isSecret: true/false` configuration property with a string-valued `$secret: always/never` property throughout the code generator. This is preparation for introducing additional values (e.g. `optionally`) in a later change.

The generated output under `v2/api/` must remain identical after migration.

## Constraints

- Backward compatible (temporary): the YAML parser accepts both `$isSecret: true/false` (legacy) and `$secret: always/never` (new). Legacy values are converted to the new `Secrecy` type during deserialization (`true` → `always`, `false` → `never`). This eases migration; `$isSecret` support will be removed once migration is complete.
- Error messages only reference `$secret` (the new tag), not `$isSecret`.
- No changes to generated output — `controller:generate-types` produces identical code
- Follow existing patterns for string-valued config properties (e.g. `$conversionStrategy`)

## New Go Type

Define `Secrecy` in the `astmodel` package (avoids circular import since `config` imports `astmodel`):

```go
// Secrecy classifies whether a property contains a secret value
type Secrecy string

const (
    SecrecyAlways = Secrecy("always") // Property always contains a secret
    SecrecyNever  = Secrecy("never")  // Property never contains a secret
)
```

Place in a new file `v2/tools/generator/internal/astmodel/secrecy.go`.

## Files Changed

### 1. New file: `astmodel/secrecy.go`

- Type `Secrecy string`
- Constants `SecrecyAlways`, `SecrecyNever`

### 2. `config/property_configuration.go`

- Add new tag constant: `secretTag = "$secret"`; keep `isSecretTag = "$isSecret"` for backward compat
- Change field: `IsSecret configurable[bool]` → `Secrecy configurable[astmodel.Secrecy]`
- Update `NewPropertyConfiguration`: `makeConfigurable[astmodel.Secrecy](secretTag, scope)`
- Update YAML parsing to handle both tags:
  - `$secret`: Switch on `strings.ToLower(c.Value)` for `"always"` / `"never"`, error on unknown (same pattern as `$conversionStrategy`)
  - `$isSecret` (legacy): Decode as bool, then convert `true` → `SecrecyAlways`, `false` → `SecrecyNever`

### 3. `config/object_model_configuration.go`

- Change field: `IsSecret propertyAccess[bool]` → `Secrecy propertyAccess[astmodel.Secrecy]`
- Update initialization: `makePropertyAccess[astmodel.Secrecy](result, func(c *PropertyConfiguration) *configurable[astmodel.Secrecy] { return &c.Secrecy })`

### 4. `astmodel/property_definition.go`

- Change field: `isSecret bool` → `secrecy Secrecy`
- Add new method: `WithSecrecy(Secrecy) *PropertyDefinition`
- Add new method: `Secrecy() Secrecy`
- Update `IsSecret() bool` to return `property.secrecy == SecrecyAlways` (keeps convenience for downstream)
- Remove `WithIsSecret(bool)` (no longer needed)
- Update equality: `property.isSecret == o.isSecret` → `property.secrecy == o.secrecy`

### 5. `codegen/pipeline/add_secrets.go`

- Change `config.ObjectModelConfiguration.IsSecret.Lookup(...)` → `config.ObjectModelConfiguration.Secrecy.Lookup(...)`
- Variable `isSecret` (bool) → `secrecy` (Secrecy); check `secrecy == astmodel.SecrecyAlways` where needed
- `prop.WithIsSecret(isSecret)` → `prop.WithSecrecy(secrecy)`
- `prop.WithIsSecret(false)` → `prop.WithSecrecy(astmodel.SecrecyNever)` (in `removeSecretProperties`)
- Error message: `"must be configured with $isSecret"` → `"must be configured with $secret"`
- Verify error: `"Found unused $isSecret configurations"` → `"Found unused $secret configurations"`
- `VerifyConsumed()` call: `config.ObjectModelConfiguration.IsSecret` → `config.ObjectModelConfiguration.Secrecy`
- All error messages reference only `$secret`, not `$isSecret`
- Secret detector struct's `isSecret bool` field: unchanged (internal heuristic, not config)

### 6. `codegen/pipeline/export_controller_type_registrations.go`

- `prop.IsSecret()` — no change needed (convenience method still returns bool)

### 7. `armconversion/convert_to_arm_function_builder.go`

- No changes needed (operates on transformed types, not the isSecret flag)

### 8. `jsonast/jsonast.go`

- `property.WithIsSecret(true)` → `property.WithSecrecy(astmodel.SecrecyAlways)`

### 9. YAML configuration files

**`v2/azure-arm.yaml`** (~97 changes):
- All `$isSecret: true` → `$secret: always`
- All `$isSecret: false` → `$secret: never`
- Doc comment block updated to document `$secret` syntax (no mention of legacy `$isSecret`)

**`hack/crossplane/azure-crossplane.yaml`**: No changes (0 occurrences)
**`v2/azure-stack.yaml`**: No changes (0 occurrences)

### 10. Tests

**`config/property_configuration_test.go`**:
- `property.IsSecret.Set(true)` → `property.Secrecy.Set(astmodel.SecrecyAlways)`
- `property.IsSecret.Lookup()` → `property.Secrecy.Lookup()`
- Expected values change from bool to Secrecy constants

**`config/type_configuration_test.go`**:
- YAML snippets: `$isSecret: true` → `$secret: always`

**`codegen/pipeline/create_arm_types_test.go`**:
- `pc.IsSecret.Set(true)` → `pc.Secrecy.Set(astmodel.SecrecyAlways)`

## Verification

After all changes:
1. `task format-code` — passes
2. `task generator:quick-checks` — builds generator, runs unit tests, generates code
3. `task controller:quick-checks` — builds controller, runs unit tests
4. No diff under `v2/api/` — generated code is identical
