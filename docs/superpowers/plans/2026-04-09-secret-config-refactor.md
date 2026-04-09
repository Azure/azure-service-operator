# Migrate `$isSecret` to `$secret` with `Secrecy` Type — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace boolean `$isSecret: true/false` with string-valued `$secret: always/never` throughout the code generator, keeping backward compat for migration.

**Architecture:** New `Secrecy` string type in `astmodel` package. Config layer parses both `$secret` (new) and `$isSecret` (legacy, temporary). AST model uses `Secrecy` type instead of `bool`. All error messages reference only `$secret`.

**Tech Stack:** Go, YAML configuration

---

### Task 1: Create the `Secrecy` type in `astmodel`

**Files:**
- Create: `v2/tools/generator/internal/astmodel/secrecy.go`

- [ ] **Step 1: Create `secrecy.go`**

Create `v2/tools/generator/internal/astmodel/secrecy.go`:

```go
/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// Secrecy classifies whether a property contains a secret value.
type Secrecy string

const (
	SecrecyAlways = Secrecy("always") // Property always contains a secret
	SecrecyNever  = Secrecy("never")  // Property never contains a secret
)
```

- [ ] **Step 2: Verify it compiles**

Run: `cd v2/tools/generator && go build ./internal/astmodel/...`
Expected: Success, no errors.

- [ ] **Step 3: Commit**

```bash
git add v2/tools/generator/internal/astmodel/secrecy.go
git commit -m "Add Secrecy type to astmodel package"
```

---

### Task 2: Update `PropertyDefinition` in `astmodel`

**Files:**
- Modify: `v2/tools/generator/internal/astmodel/property_definition.go`

- [ ] **Step 1: Change the field type and add new methods**

In `v2/tools/generator/internal/astmodel/property_definition.go`:

Replace the field declaration (line ~53):
```go
	isSecret bool
```
with:
```go
	secrecy Secrecy
```

Replace `WithIsSecret` method (lines ~149-157):
```go
// WithIsSecret returns a new PropertyDefinition with IsSecret set to the specified value
func (property *PropertyDefinition) WithIsSecret(secret bool) *PropertyDefinition {
	if secret == property.isSecret {
		return property
	}

	result := property.copy()
	result.isSecret = secret
	return result
}
```
with:
```go
// WithSecrecy returns a new PropertyDefinition with the Secrecy classification set to the specified value
func (property *PropertyDefinition) WithSecrecy(secrecy Secrecy) *PropertyDefinition {
	if secrecy == property.secrecy {
		return property
	}

	result := property.copy()
	result.secrecy = secrecy
	return result
}
```

Update `IsSecret` method (lines ~389-391):
```go
// IsSecret returns true iff the property is a secret.
func (property *PropertyDefinition) IsSecret() bool {
	return property.isSecret
}
```
with:
```go
// IsSecret returns true iff the property is a secret.
func (property *PropertyDefinition) IsSecret() bool {
	return property.secrecy == SecrecyAlways
}
```

Update `Equals` method — change the comparison (line ~493):
```go
		property.isSecret == o.isSecret &&
```
to:
```go
		property.secrecy == o.secrecy &&
```

- [ ] **Step 2: Verify it compiles**

Run: `cd v2/tools/generator && go build ./...`
Expected: Compilation errors in files that call `WithIsSecret` — these will be fixed in subsequent tasks. Verify only `property_definition.go` itself compiles cleanly:
Run: `cd v2/tools/generator && go vet ./internal/astmodel/...`

- [ ] **Step 3: Fix `jsonast.go` call site**

In `v2/tools/generator/internal/jsonast/jsonast.go` (line ~543), replace:
```go
			property = property.WithIsSecret(true)
```
with:
```go
			property = property.WithSecrecy(astmodel.SecrecyAlways)
```

- [ ] **Step 4: Fix `add_secrets.go` call sites for `WithIsSecret`**

In `v2/tools/generator/internal/codegen/pipeline/add_secrets.go`:

Replace (line ~84):
```go
				propWithSecret := prop.WithIsSecret(isSecret)
				it = it.WithProperty(propWithSecret)
```
with (temporary — will be cleaned up in Task 4 when config types change):
```go
				secrecyValue := astmodel.SecrecyNever
				if isSecret {
					secrecyValue = astmodel.SecrecyAlways
				}
				it = it.WithProperty(prop.WithSecrecy(secrecyValue))
```

Replace (line ~281):
```go
				it = it.WithProperty(prop.WithIsSecret(false))
```
with:
```go
				it = it.WithProperty(prop.WithSecrecy(astmodel.SecrecyNever))
```

- [ ] **Step 5: Verify full build**

Run: `cd v2/tools/generator && go build ./...`
Expected: Success, no errors.

- [ ] **Step 6: Run generator unit tests**

Run: `cd v2/tools/generator && go test ./... -tags=noexit -count=1 -timeout 5m`
Expected: All tests pass.

- [ ] **Step 7: Commit**

```bash
git add -A
git commit -m "Replace isSecret bool with Secrecy type in PropertyDefinition"
```

---

### Task 3: Update `PropertyConfiguration` in config layer

**Files:**
- Modify: `v2/tools/generator/internal/config/property_configuration.go`

- [ ] **Step 1: Add `$secret` tag, keep `$isSecret` for backward compat**

In `v2/tools/generator/internal/config/property_configuration.go`, update the tag constants block (lines ~68-76). Replace:
```go
	isSecretTag                       = "$isSecret"                       // Bool specifying whether a property contains a secret
```
with:
```go
	isSecretTag                       = "$isSecret"                       // Deprecated: use secretTag instead. Bool specifying whether a property contains a secret. Kept for backward compatibility during migration.
	secretTag                         = "$secret"                         // String specifying the secrecy classification of a property (always or never)
```

- [ ] **Step 2: Change field type from `configurable[bool]` to `configurable[Secrecy]`**

Add import for astmodel at the top of the file:
```go
import (
	"strings"

	"github.com/rotisserie/eris"
	"gopkg.in/yaml.v3"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)
```

Replace the field (line ~31):
```go
	IsSecret                       configurable[bool]                // Specify whether this property is a secret
```
with:
```go
	Secrecy                        configurable[astmodel.Secrecy]    // Specify the secrecy classification of this property
```

- [ ] **Step 3: Update `NewPropertyConfiguration` initialization**

Replace (line ~87):
```go
		IsSecret:                       makeConfigurable[bool](isSecretTag, scope),
```
with:
```go
		Secrecy:                        makeConfigurable[astmodel.Secrecy](secretTag, scope),
```

- [ ] **Step 4: Update YAML parsing — replace old `$isSecret` handler, add both `$secret` and `$isSecret`**

Replace the existing `$isSecret` parsing block (lines ~127-137):
```go
		// $isSecret: <bool>
		if strings.EqualFold(lastID, isSecretTag) && c.Kind == yaml.ScalarNode {
			var isSecret bool
			err := c.Decode(&isSecret)
			if err != nil {
				return eris.Wrapf(err, "decoding %s", isSecretTag)
			}

			pc.IsSecret.Set(isSecret)
			continue
		}
```
with:
```go
		// $secret: <string> (preferred)
		if strings.EqualFold(lastID, secretTag) && c.Kind == yaml.ScalarNode {
			switch strings.ToLower(c.Value) {
			case string(astmodel.SecrecyAlways):
				pc.Secrecy.Set(astmodel.SecrecyAlways)
			case string(astmodel.SecrecyNever):
				pc.Secrecy.Set(astmodel.SecrecyNever)
			default:
				return eris.Errorf("unknown %s value: %s.", secretTag, c.Value)
			}

			continue
		}

		// $isSecret: <bool> (legacy, for backward compatibility during migration)
		if strings.EqualFold(lastID, isSecretTag) && c.Kind == yaml.ScalarNode {
			var isSecret bool
			err := c.Decode(&isSecret)
			if err != nil {
				return eris.Wrapf(err, "decoding %s", isSecretTag)
			}

			if isSecret {
				pc.Secrecy.Set(astmodel.SecrecyAlways)
			} else {
				pc.Secrecy.Set(astmodel.SecrecyNever)
			}

			continue
		}
```

- [ ] **Step 5: Verify it compiles**

Run: `cd v2/tools/generator && go build ./internal/config/...`
Expected: Compilation errors in `object_model_configuration.go` (refers to `IsSecret`) — fixed in next task. The config package itself should compile.

- [ ] **Step 6: Commit**

```bash
git add -A
git commit -m "Update PropertyConfiguration to use Secrecy type with $secret tag"
```

---

### Task 4: Update `ObjectModelConfiguration` and fix `add_secrets.go`

**Files:**
- Modify: `v2/tools/generator/internal/config/object_model_configuration.go`
- Modify: `v2/tools/generator/internal/codegen/pipeline/add_secrets.go`

- [ ] **Step 1: Update `ObjectModelConfiguration` field and initialization**

In `v2/tools/generator/internal/config/object_model_configuration.go`:

Replace the field (line ~54):
```go
	IsSecret                       propertyAccess[bool]
```
with:
```go
	Secrecy                        propertyAccess[astmodel.Secrecy]
```

Replace the initialization (lines ~116-117):
```go
	result.IsSecret = makePropertyAccess[bool](
		result, func(c *PropertyConfiguration) *configurable[bool] { return &c.IsSecret })
```
with:
```go
	result.Secrecy = makePropertyAccess[astmodel.Secrecy](
		result, func(c *PropertyConfiguration) *configurable[astmodel.Secrecy] { return &c.Secrecy })
```

- [ ] **Step 2: Fully update `add_secrets.go`**

In `v2/tools/generator/internal/codegen/pipeline/add_secrets.go`:

Replace the config lookup and usage block (the section that does `Lookup` and the temporary bool-to-Secrecy bridge):
```go
			isSecret, isSecretConfigured := config.ObjectModelConfiguration.IsSecret.Lookup(ctx, prop.PropertyName())
			if ctx.IsStatus() && !isSecretConfigured {
				isSecret, isSecretConfigured = config.ObjectModelConfiguration.IsSecret.Lookup(strippedTypeName, prop.PropertyName())
			}

			// If it's not a secret, but it looks like a secret, and we don't have any configuration to tell us for
			// sure, request configuration so we know for sure.
			if !prop.IsSecret() && maybeSecret && !isSecretConfigured {
				// Property might be a secret, but isn't already configured as one,
				// and we don't have config to tell us for sure
				return nil, eris.Errorf(
					"property %s might be a secret and must be configured with $isSecret",
					prop.PropertyName())
			}

			if isSecretConfigured {
				secrecyValue := astmodel.SecrecyNever
				if isSecret {
					secrecyValue = astmodel.SecrecyAlways
				}
				it = it.WithProperty(prop.WithSecrecy(secrecyValue))
			}
```
with:
```go
			secrecy, secrecyConfigured := config.ObjectModelConfiguration.Secrecy.Lookup(ctx, prop.PropertyName())
			if ctx.IsStatus() && !secrecyConfigured {
				secrecy, secrecyConfigured = config.ObjectModelConfiguration.Secrecy.Lookup(strippedTypeName, prop.PropertyName())
			}

			// If it's not a secret, but it looks like a secret, and we don't have any configuration to tell us for
			// sure, request configuration so we know for sure.
			if !prop.IsSecret() && maybeSecret && !secrecyConfigured {
				// Property might be a secret, but isn't already configured as one,
				// and we don't have config to tell us for sure
				return nil, eris.Errorf(
					"property %s might be a secret and must be configured with $secret",
					prop.PropertyName())
			}

			if secrecyConfigured {
				it = it.WithProperty(prop.WithSecrecy(secrecy))
			}
```

Replace the verify-consumed block (lines ~118-123):
```go
	err := config.ObjectModelConfiguration.IsSecret.VerifyConsumed()
	if err != nil {
		return nil, eris.Wrap(
			err,
			"Found unused $isSecret configurations; these need to be fixed or removed.")
	}
```
with:
```go
	err := config.ObjectModelConfiguration.Secrecy.VerifyConsumed()
	if err != nil {
		return nil, eris.Wrap(
			err,
			"Found unused $secret configurations; these need to be fixed or removed.")
	}
```

- [ ] **Step 3: Verify full build**

Run: `cd v2/tools/generator && go build ./...`
Expected: Success, no errors.

- [ ] **Step 4: Run generator unit tests**

Run: `cd v2/tools/generator && go test ./... -tags=noexit -count=1 -timeout 5m`
Expected: Some test failures in tests still referencing `IsSecret` — fixed in next task.

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "Update ObjectModelConfiguration and add_secrets pipeline to use Secrecy"
```

---

### Task 5: Update all tests

**Files:**
- Modify: `v2/tools/generator/internal/config/property_configuration_test.go`
- Modify: `v2/tools/generator/internal/config/type_configuration_test.go`
- Modify: `v2/tools/generator/internal/codegen/pipeline/create_arm_types_test.go`

- [ ] **Step 1: Update `property_configuration_test.go`**

In `v2/tools/generator/internal/config/property_configuration_test.go`:

Add import for astmodel (if not already present):
```go
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
```

Replace the "WhenSpecified" test (lines ~95-105):
```go
func TestPropertyConfiguration_IsSecret_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.IsSecret.Set(true)

	isSecret, ok := property.IsSecret.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(isSecret).To(BeTrue())
}
```
with:
```go
func TestPropertyConfiguration_Secrecy_WhenSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")
	property.Secrecy.Set(astmodel.SecrecyAlways)

	secrecy, ok := property.Secrecy.Lookup()
	g.Expect(ok).To(BeTrue())
	g.Expect(secrecy).To(Equal(astmodel.SecrecyAlways))
}
```

Replace the "WhenNotSpecified" test (lines ~107-117):
```go
func TestPropertyConfiguration_IsSecret_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	isSecret, ok := property.IsSecret.Lookup()

	g.Expect(isSecret).To(BeFalse())
	g.Expect(ok).To(BeFalse())
}
```
with:
```go
func TestPropertyConfiguration_Secrecy_WhenNotSpecified_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	property := NewPropertyConfiguration("Property")

	secrecy, ok := property.Secrecy.Lookup()

	g.Expect(secrecy).To(Equal(astmodel.Secrecy("")))
	g.Expect(ok).To(BeFalse())
}
```

- [ ] **Step 2: Update `type_configuration_test.go`**

In `v2/tools/generator/internal/config/type_configuration_test.go`:

Replace YAML snippets in both duplicate-key tests (lines ~197 and ~215). Replace all occurrences of:
```yaml
  $isSecret: true
```
with:
```yaml
  $secret: always
```

There are two test functions that use this YAML:
- `TestTypeConfiguration_UnmarshalYAML_WhenDuplicateProperties_ReturnsError`
- `TestTypeConfiguration_UnmarshalYAML_WhenDuplicatePropertiesCaseInsensitive_ReturnsError`

- [ ] **Step 3: Update `create_arm_types_test.go`**

In `v2/tools/generator/internal/codegen/pipeline/create_arm_types_test.go`:

Replace all three occurrences of (lines ~301, ~310, ~319):
```go
				pc.IsSecret.Set(true)
```
with:
```go
				pc.Secrecy.Set(astmodel.SecrecyAlways)
```

Ensure the test file imports `astmodel`:
```go
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
```

- [ ] **Step 4: Run all generator unit tests**

Run: `cd v2/tools/generator && go test ./... -tags=noexit -count=1 -timeout 5m`
Expected: All tests pass.

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "Update tests to use Secrecy type"
```

---

### Task 6: Migrate `azure-arm.yaml`

**Files:**
- Modify: `v2/azure-arm.yaml`

- [ ] **Step 1: Update the doc comment block**

In `v2/azure-arm.yaml`, replace the doc comment (lines ~1001-1007):
```yaml
# $isSecret: <bool>
#     Specifies that the property references a secret
#     Secrets are string values read from a secure storage by ASO when needed.
#     Set to `true` to flag this property as a secret. This is an override for when the Swagger
#     is incomplete. If you specify $isSecret: true, you should also open a PR to update the 
#     upstream Swagger repo spec with the x-ms-secret annotation.
#     Here's a reference PR: https://github.com/Azure/azure-rest-api-specs/pull/19399
```
with:
```yaml
# $secret: always | never
#     Specifies the secrecy classification of a property.
#     Secrets are string values read from a secure storage by ASO when needed.
#     Set to `always` to flag this property as a secret. This is an override for when the Swagger
#     is incomplete. If you specify $secret: always, you should also open a PR to update the
#     upstream Swagger repo spec with the x-ms-secret annotation.
#     Here's a reference PR: https://github.com/Azure/azure-rest-api-specs/pull/19399
```

- [ ] **Step 2: Replace all `$isSecret: true` with `$secret: always`**

Run:
```bash
cd /workspaces/azure-service-operator-2
sed -i 's/\$isSecret: true/$secret: always/g' v2/azure-arm.yaml
```

This should replace 79 occurrences (the doc comment was already updated in step 1, which changed the line that contained `$isSecret: true` in the comment text).

- [ ] **Step 3: Replace all `$isSecret: false` with `$secret: never`**

Run:
```bash
sed -i 's/\$isSecret: false/$secret: never/g' v2/azure-arm.yaml
```

This should replace 17 occurrences.

- [ ] **Step 4: Verify no remaining `$isSecret` in the YAML (except none should remain)**

Run:
```bash
grep -n '$isSecret' v2/azure-arm.yaml
```

Expected: No output (zero matches).

- [ ] **Step 5: Commit**

```bash
git add v2/azure-arm.yaml
git commit -m "Migrate azure-arm.yaml from \$isSecret to \$secret"
```

---

### Task 7: Full verification

**Files:** None (verification only)

- [ ] **Step 1: Format code**

Run: `task format-code`
Expected: Success. Fix any formatting issues if they arise.

- [ ] **Step 2: Run generator quick checks**

Run: `task generator:quick-checks`
Expected: Builds generator, runs unit tests, generates code, lints — all pass.

- [ ] **Step 3: Run controller quick checks**

Run: `task controller:quick-checks`
Expected: Builds controller, runs unit tests, lints — all pass.

- [ ] **Step 4: Verify no generated code changes**

Run:
```bash
git diff --stat v2/api/
```
Expected: No output — zero changes under `v2/api/`.

- [ ] **Step 5: Verify no uncommitted changes**

Run: `git status`
Expected: Clean working tree (or only the plan/spec docs).

- [ ] **Step 6: Final commit (if formatting produced changes)**

```bash
git add -A
git commit -m "Format code after secret config refactor"
```
