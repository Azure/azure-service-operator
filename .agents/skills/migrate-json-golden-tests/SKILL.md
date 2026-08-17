---
name: migrate-json-golden-tests
description: "**WORKFLOW SKILL** — Migrate JSON-based golden file test scenarios (TestGolden) to programmatic Go unit tests in the pipeline package. USE FOR: converting testdata/<GroupName>/*.json scenarios into Go tests that construct types with astmodel and run specific pipeline stages. DO NOT USE FOR: adding new test scenarios from scratch, fixing existing Go unit tests, or code review."
disable-model-invocation: true
---

# Skill: Migrate JSON Golden File Tests to Pipeline Unit Tests

This skill guides the migration of JSON-based golden file test scenarios (run by `TestGolden` in `golden_files_test.go`) into programmatic Go unit tests in the `pipeline` package. The new tests construct types using `astmodel` APIs and run specific pipeline stages, providing focused, maintainable unit tests.

## Background

### Two Test Systems

1. **JSON golden file tests** (`v2/tools/generator/internal/codegen/golden_files_test.go`):
   - JSON files in `testdata/<GroupName>/` directories define schemas
   - `config.yaml` specifies options (`hasArmResources`, `pipelines`)
   - Runs through the FULL code generator pipeline (80+ stages)
   - Tests are discovered automatically by walking the testdata directory
   - Output includes webhooks, storage types, conversion functions, etc.

2. **Pipeline unit tests** (e.g., `v2/tools/generator/internal/codegen/pipeline/create_arm_types_test.go`):
   - Types are constructed programmatically using `astmodel` APIs and `test` helpers
   - Runs only the specific pipeline stages being tested
   - Uses `test.AssertPackagesGenerateExpectedCode()` for golden file comparison
   - Golden files stored in `pipeline/testdata/<TestFuncName>/` directories

The unit tests are preferred because they are faster, more focused, and don't depend on the JSON schema scanner.

## Prerequisites

- You know the **group name** (the testdata subdirectory name, e.g., `ArmResource`) to migrate.
- The testdata directory exists at `v2/tools/generator/internal/codegen/testdata/<GroupName>/`.

## Step-by-Step Procedure

### Step 1: Examine the Existing JSON Scenarios

1. List all `.json` files in `v2/tools/generator/internal/codegen/testdata/<GroupName>/`.
2. Read `config.yaml` in the same directory to understand test configuration:
   - `hasArmResources: true/false` — determines which pipeline stages run
   - `pipelines:` — which pipelines to test (azure, crossplane)
3. Read each JSON file to understand what it tests:
   - What resource definitions exist (`resourceDefinitions` section)
   - What types are defined (`definitions` section)
   - Property types (string, integer, boolean, array, map, enum, $ref, oneOf)
   - Required vs optional properties (`required` arrays)
   - Special features (resource references, nested resources, discriminated unions)

### Step 2: Determine Where to Place New Tests

Tests always go in a `*_test.go` file in the `pipeline` package (`v2/tools/generator/internal/codegen/pipeline/`), typically alongside the pipeline stage being tested.

To determine the specific file:
- Look at the `config.yaml` to understand what features the scenarios exercise (e.g., `hasArmResources: true` suggests ARM type creation is involved).
- Examine what the JSON scenarios test and find the corresponding pipeline stage.
- Place tests in the test file for that stage:
  - ARM type creation and conversion → `create_arm_types_test.go`
  - Enum handling → the enum-related test file
  - Type alias removal → the type alias test file
  - Other stage-specific scenarios → the corresponding `<stage_name>_test.go` file

### Step 3: Check for Existing Coverage

Before creating new tests, check the target test file for tests that already cover the same scenarios:

1. Read the target test file completely.
2. For each JSON scenario, determine if an existing test constructs equivalent types and runs the same pipeline stages.
3. Note which scenarios already have coverage and which need new tests.
4. Do NOT recreate tests that already exist — just note the match.

### Step 4: Create the New Tests

For each JSON scenario that needs a new test, create a Go test function.

#### 4a: Map JSON Schema to astmodel Types

| JSON Schema Pattern | astmodel Equivalent |
|---|---|
| `"type": "string"` | `astmodel.StringType` |
| `"type": "integer"` | `astmodel.IntType` |
| `"type": "boolean"` | `astmodel.BoolType` |
| `"type": "object"` with no properties | `astmodel.NewMapType(astmodel.StringType, astmodel.AnyType)` |
| `{}` (empty schema) | `astmodel.AnyType` |
| `"type": "array", "items": X` | `astmodel.NewArrayType(X)` |
| `"additionalProperties": X` | `astmodel.NewMapType(astmodel.StringType, X)` |
| `"$ref": "#/definitions/Foo"` | Reference the TypeName of the Foo definition |
| `"enum": [values]` | `astmodel.NewEnumType(baseType, astmodel.MakeEnumValue(id, `\`"value"\``)...)` — **value must be backtick-quoted** |
| `"oneOf": [refs]` | Object with optional properties for each variant + `astmodel.OneOfFlag.ApplyTo()` |

#### 4b: Handle Required/Optional Properties

The JSON schema scanner wraps ALL property types in optional and then annotates required ones:

```go
// Required property:
prop := astmodel.NewPropertyDefinition("Name", "name", SomeType).MakeTypeOptional().MakeRequired()

// Optional property:
prop := astmodel.NewPropertyDefinition("Name", "name", SomeType).MakeTypeOptional()
// or equivalently:
prop := astmodel.NewPropertyDefinition("Name", "name", astmodel.NewOptionalType(SomeType))
```

**IMPORTANT**: `MakeRequired()` panics if the property type is not already optional. Always call `MakeTypeOptional()` first.

#### 4c: Create Resource Structures

For ARM resources (the common case):

```go
spec := test.CreateSpec(test.Pkg2020, "ResourceName", properties...)
status := test.CreateStatus(test.Pkg2020, "ResourceName")
resource := test.CreateARMResource(test.Pkg2020, "ResourceName", spec, status, test.Pkg2020APIVersion)

defs := make(astmodel.TypeDefinitionSet)
defs.AddAll(resource, status, spec, /* other type defs... */ test.Pkg2020APIVersion)
```

For resources with a nested `properties` object (common ARM pattern):

```go
propsObj := test.CreateObjectDefinition(pkg, "ResourceNameProperties", prop1, prop2, ...)
propsProp := astmodel.NewPropertyDefinition("Properties", "properties", propsObj.Name()).MakeTypeOptional()
spec := test.CreateSpec(pkg, "ResourceName", test.NameProperty, propsProp)
```

For resource ownership (parent-child relationships):

```go
resourceBRT, _ := astmodel.AsResourceType(resourceB.Type())
resourceB = resourceB.WithType(resourceBRT.WithOwner(resourceA.Name()))
```

#### 4d: Choose Pipeline Stages

Determine which pipeline stages to run by examining what the existing tests in the target file use. Look at the other tests already in the file to understand the standard pattern for that stage, then replicate it for your new tests.

For example, `create_arm_types_test.go` uses this standard set:

```go
state, err := RunTestPipeline(
    NewState(defs),
    CreateARMTypes(cfg, idFactory, logr.Discard()),
    ApplyARMConversionInterface(idFactory, cfg),
    SimplifyDefinitions(),
    StripUnreferencedTypeDefinitions(),
)
```

If the scenario exercises additional features, add the relevant stages. Look at existing tests in the same file for examples of how to include stages for specific features:

| Scenario | Additional Stages |
|---|---|
| **Resource references** | Configure OMC with `ReferenceType.Set(config.ReferenceTypeARM)` per property, then add `ApplyCrossResourceReferencesFromConfig(configuration, logr.Discard())` and `TransformCrossResourceReferences(configuration, idFactory)` before the core stages. |
| **Config maps** | Add `AddConfigMaps(configuration)` before the core stages |
| **Secrets** | Add `AddSecrets(configuration)` before the core stages |
| **Flattening** | Add `FlattenProperties(logr.Discard())` after the core stages |
| **JSON/Any type fields** | Add `ReplaceAnyTypeWithJSON()` before the core stages |
| **OneOf** | Standard stages work; the type must have `OneOfFlag` applied via `ApplyObjectTransformation` |

For resource references, use standard pipeline stages with OMC configuration — do NOT write custom test helper functions:

```go
omc := config.NewObjectModelConfiguration()
g.Expect(
    omc.ModifyProperty(
        specProperties.Name(),
        someProperty.PropertyName(),
        func(pc *config.PropertyConfiguration) error {
            pc.ReferenceType.Set(config.ReferenceTypeARM)
            return nil
        },
    ),
).To(Succeed())

configuration := config.NewConfiguration()
configuration.ObjectModelConfiguration = omc

state, err := RunTestPipeline(
    NewState(defs),
    ApplyCrossResourceReferencesFromConfig(configuration, logr.Discard()),
    TransformCrossResourceReferences(configuration, idFactory),
    CreateARMTypes(omc, idFactory, logr.Discard()),
    ApplyARMConversionInterface(idFactory, omc),
    SimplifyDefinitions(),
    StripUnreferencedTypeDefinitions(),
)
```

#### 4e: Assert and Generate Golden Files

```go
g.Expect(err).ToNot(HaveOccurred())
test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
```

#### 4f: Test Naming and Comments

- Name tests descriptively: `TestCreateARMTypes_SimpleResourceMapProperties`
- Add a comment describing what the test covers:
  ```go
  // TestCreateARMTypes_SimpleResourceMapProperties tests that an ARM resource with various map property
  // types (maps of objects, maps of maps, maps of arrays, maps of enums, maps of strings) generates
  // correct ARM types and conversions.
  ```
- Do NOT reference the old JSON file name in comments — the files will be deleted.

#### 4g: Generate Golden Files

Run the new tests with `-update` to create golden files:

```bash
cd v2/tools/generator
go test ./internal/codegen/pipeline/ -run "TestName1|TestName2|..." -update -v
```

Then verify they pass without `-update`:

```bash
go test ./internal/codegen/pipeline/ -run "TestName1|TestName2|..." -v
```

Also run all existing tests to confirm no regressions:

```bash
go test ./internal/codegen/pipeline/ -run "TestCreate" -v
```

### Step 5: Compare Old and New Golden Files

Generate a comparison report between the new pipeline test golden files and the old JSON test golden files. The expected differences fall into two categories:

**Expected/benign differences** (present in all scenarios due to running fewer pipeline stages):
- **Package name**: New uses `person` (from `test.Pkg2020`), old uses `test` (from JSON schema URL)
- **API version string**: New uses `"v2020"`, old uses the version from JSON schema (e.g., `"2020-01-01"`)
- **APIVersion/Type enum fields on ARM spec**: Absent in new (not created by unit pipeline), present in old
- **Status type**: New has `Status string` field, old may differ
- **`// Generated from:` comments**: Absent in new, present in old
- **Enum value formatting**: New uses `Color_blue = Color(blue)`, old uses `Color_Blue = Color("blue")`
- **Additional interfaces/scaffolding**: Old may include extra interfaces (Conditioner, Convertible, Exporter, etc.) that are generated by pipeline stages not included in the unit test

**Potentially meaningful differences to investigate**:
- If types or fields that are central to what the test scenario is testing are missing or have a different structure, a pipeline stage may need to be added to the unit test. Check what stage produces the expected output and add it.
- If types appear with the wrong representation (e.g., `interface{}` instead of a typed wrapper), look for a transformation stage that converts between representations and add it to the pipeline.

Run the comparison:

```bash
for scenario in "NewTestName:old_json_name"; do
  new_name="${scenario%%:*}"
  old_name="${scenario##*:}"
  diff "pipeline/testdata/${new_name}/person-v20200101-arm.golden" \
       "testdata/<GroupName>/${old_name}_azure_arm.golden"
done
```

Present the report to the user and get confirmation before proceeding.

### Step 6: Delete the Old Tests

Once the user confirms the comparison looks good:

```bash
rm -rf v2/tools/generator/internal/codegen/testdata/<GroupName>/
```

Then verify:

1. All remaining golden file tests pass:
   ```bash
   go test ./internal/codegen/ -run "TestGolden" -v -count=1
   ```

2. All new pipeline tests pass:
   ```bash
   go test ./internal/codegen/pipeline/ -v -count=1
   ```

## Available Test Helpers

### Package: `test` (`v2/tools/generator/internal/test/`)
- `test.Pkg2020`, `test.Pkg2021`, `test.Pkg2022` — package references
- `test.Pkg2020APIVersion` — API version enum definition
- `test.NameProperty`, `test.FullNameProperty`, `test.FamilyNameProperty`, `test.KnownAsProperty`, `test.RestrictedNameProperty` — reusable properties
- `test.CreateSpec()`, `test.CreateStatus()`, `test.CreateARMResource()`, `test.CreateResource()` — resource builders
- `test.CreateObjectDefinition()` — creates an ObjectType with properties
- `test.AssertPackagesGenerateExpectedCode()` — golden file assertion
- `test.CreateFolderForTest()` — option for subtests needing unique golden file folders

### Package: `astmodel` (`v2/tools/generator/internal/astmodel/`)
- `astmodel.NewPropertyDefinition(name, jsonName, type)` — with `.MakeTypeOptional()`, `.MakeRequired()`, `.WithDescription()`
- `astmodel.NewEnumType(baseType, values...)`, `astmodel.MakeEnumValue(id, value)` — enums. **The `value` parameter is a literal Go expression that appears in generated code; for string enums it must include quotes:** `` astmodel.MakeEnumValue("blue", `"blue"`) ``
- `astmodel.NewArrayType(element)` — arrays
- `astmodel.NewMapType(key, value)` — maps
- `astmodel.NewOptionalType(element)` — optional wrapper
- `astmodel.OneOfFlag.ApplyTo(objectType)` — oneOf flag
- `astmodel.MakeTypeDefinition(name, type)` — type definition
- `astmodel.MakeInternalTypeName(pkg, name)` — type name
- `astmodel.AsResourceType(type)` — cast to resource type (for ownership)
- `resourceType.WithOwner(ownerName)` — set resource ownership

### Pipeline Stages (`v2/tools/generator/internal/codegen/pipeline/`)
- `RunTestPipeline(state, stages...)` — run pipeline stages in sequence
- `NewState(defs)` — create initial state
- `CreateARMTypes()`, `ApplyARMConversionInterface()` — core ARM stages
- `SimplifyDefinitions()`, `StripUnreferencedTypeDefinitions()` — cleanup stages
- `FlattenProperties()` — flattening stage
- `ReplaceAnyTypeWithJSON()` — converts `interface{}` to `v1.JSON`
- `ApplyCrossResourceReferencesFromConfig()`, `TransformCrossResourceReferences()` — resource reference stages
- `AddConfigMaps()`, `AddSecrets()` — configmap/secret stages

## Common Pitfalls

1. **`MakeRequired()` panics** if the property type is not already optional — always call `MakeTypeOptional()` first.
2. **Do NOT reference old JSON file names** in test comments — the JSON files will be deleted.
3. **Do NOT write custom test helper functions** for cross-resource references — use the standard `ApplyCrossResourceReferencesFromConfig` and `TransformCrossResourceReferences` pipeline stages with OMC configuration.
4. **Resource ownership** is set via `ResourceType.WithOwner()` which takes `InternalTypeName`, not a string pointer.
5. **Golden file directories** are named after the test function — test naming matters for file organization.
6. **The old and new tests exercise different pipeline depths** — the new tests are more focused unit tests; differences in output scaffolding are expected and benign.
7. **`MakeEnumValue` value parameter must be quoted for string enums** — The second argument to `MakeEnumValue` is a literal Go expression emitted in generated code. For string enums, wrap the value in backtick quotes: `` astmodel.MakeEnumValue("blue", `"blue"`) ``. Without quotes, the generated code will reference undefined identifiers (e.g., `Color(blue)` instead of `Color("blue")`).
