# Project: RAPID Testing

Overall goal is to migrate from the existing use of gopter tests to instead use the rapid library. Because this is a big change, we're going to tackle it step by step. This file is our overall plan, including reference information, and should be updated as a part of each step.

The rapid library is available from https://github.com/flyingmutant/rapid

The gopter library we're replacing is github.com/leanovate/gopter

## Phase I: Research gopter

Read up on the gopter library and how it's used in the ASO code generator.

The pipeline stage is in `json_serialization_test_cases.go`.

A sample of the generated gopter code:
`v2/api/batch/v20240701/batch_account_types_gen_test.go`

The code generator is found at `v2/tools/generator`. It's complex, so don't try to read all of it in one go.

Add notes to this section detailing what you find.

### How gopter is used in the code generator

#### Test Case Types

There are three test case types that use gopter, all in `v2/tools/generator/internal/testcases/`:

1. **`JSONSerializationTestCase`** (`json_serialization_test_case.go`) — Tests that a type can be round-tripped through JSON serialization. For each type, it generates:
   - A **test runner function** (e.g. `Test_BatchAccount_Spec_WhenSerializedToJson_DeserializesAsEqual`) that configures gopter parameters and runs the property test.
   - A **test method function** (e.g. `RunJSONSerializationTestForBatchAccount_Spec`) that serializes to JSON and back, comparing results. Returns `string` (empty on success, diff on failure).
   - A **generator global variable** (e.g. `batchAccount_SpecGenerator`) — a lazily-initialized `gopter.Gen`.
   - A **generator factory method** (e.g. `BatchAccount_SpecGenerator()`) that builds the generator using `gen.Struct(reflect.TypeOf(...), generators)`.
   - **Independent generators factory** (e.g. `AddIndependentPropertyGeneratorsForBatchAccount_Spec()`) for primitive-typed properties.
   - **Related generators factory** (e.g. `AddRelatedPropertyGeneratorsForBatchAccount_Spec()`) for complex-typed properties that reference other generated types.

2. **`PropertyAssignmentTestCase`** (`property_assignment_test_case.go`) — Tests round-trip property assignment (conversion to next storage version and back). Uses gopter with `prop.ForAll` and **reuses the generator defined by `JSONSerializationTestCase`** — it calls the same `XGenerator()` function.

3. **`ResourceConversionTestCase`** (`resource_conversion_test_case.go`) — Tests round-trip conversion from a resource version to the hub version and back. Also uses gopter with `prop.ForAll` and **reuses the same `XGenerator()` function**.

**Critical dependency**: Both `PropertyAssignmentTestCase` and `ResourceConversionTestCase` depend on the generators created by `JSONSerializationTestCase`. Their pipeline stages (`InjectPropertyAssignmentTests`, `InjectResourceConversionTestCases`) both declare `InjectJSONSerializationTestsID` as a prerequisite. The generated test code calls the `XGenerator()` function by name — it doesn't directly generate its own generators.

#### How Generators are Built

The `JSONSerializationTestCase` builds gopter generators in two phases to handle cyclic type references:

1. **Phase 1**: Create generators for properties with primitive/simple types only (`createIndependentGenerator`). These map to gopter's built-in generators:
   - `gen.AlphaString()` for strings
   - `gen.UInt32()`, `gen.Int()`, `gen.Float64()`, `gen.Bool()` for primitives
   - `gen.SliceOf()` for arrays
   - `gen.MapOf()` for maps
   - `gen.PtrOf()` for optional types
   - Enum types get `gen.OneConstOf()` with their possible values
   - Validated types get the underlying generator plus a `.Map()` cast

2. **Phase 2**: Assign the partial generator (with only independent fields) to the global cache variable (breaking cycles), then add generators for complex properties (`createRelatedGenerator`) which call the generator factory methods of other types.

The final generator is built using `gen.Struct(reflect.TypeOf(Subject{}), generators)` which constructs a struct with the given field generators.

**Large object handling**: Types with >50 properties skip independent generators entirely (gopter limitation due to Go runtime struct reflection limits — see https://github.com/golang/go/issues/54669). Only related generators are used.

#### OneOf Types

`JSONSerializationTestCase` has special handling for OneOf types (discriminated unions) via `createGeneratorMethodForOneOf()`. Instead of `gen.Struct()`, it:
1. Creates related generators (one per OneOf option)
2. Iterates over the generator map, creating a `gen.Struct()` for each individual property
3. Combines them with `gen.OneGenOf(gens...)` to randomly select one option per test run

#### How Tests are Registered

Test cases are attached to types via the `astmodel.TestCase` interface. The pipeline stage `InjectJSONSerializationTests` (in `v2/tools/generator/internal/codegen/pipeline/json_serialization_test_cases.go`) walks all type definitions, creating a `JSONSerializationTestCase` for each property container and attaching it to the type definition via `TestCaseInjector.Inject()`.

When code is exported, `FileDefinition.AsAst()` calls `AsFuncs()` on each test case to generate the Go AST declarations that get written to `*_types_gen_test.go` files.

The factory (`objectSerializationTestCaseFactory`) also adjusts `MinSuccessfulTests`:
- Resources: 20 (spec and status are tested independently)
- Spec types: 80 (100 minus the 20 done by resource tests)
- Status types: 80 (same logic)
- Other types: 100 (default)

#### Package References

Defined in `v2/tools/generator/internal/astmodel/std_references.go`:
- `GopterReference` — `github.com/leanovate/gopter`
- `GopterGenReference` — `github.com/leanovate/gopter/gen`
- `GopterPropReference` — `github.com/leanovate/gopter/prop`
- `CmpReference` — `github.com/google/go-cmp/cmp`
- `CmpOptsReference` — `github.com/google/go-cmp/cmp/cmpopts`
- `DiffReference` — `github.com/kylelemons/godebug/diff`
- `PrettyReference` — `github.com/kr/pretty`

#### Pipeline Registration

In `v2/tools/generator/internal/codegen/code_generator.go`, the `createAllPipelineStages()` function registers the three test injection stages consecutively:
```go
pipeline.InjectJSONSerializationTests(idFactory).UsedFor(pipeline.ARMTarget),
pipeline.InjectPropertyAssignmentTests(idFactory).UsedFor(pipeline.ARMTarget),
pipeline.InjectResourceConversionTestCases(idFactory).UsedFor(pipeline.ARMTarget),
```

The JSON serialization stage uses `RequiresPostrequisiteStages("simplifyDefinitions")`, meaning other stages that add types declare themselves as needing to run before this stage.

#### Shared Helper

`v2/tools/generator/internal/testcases/shared.go` contains `idOfGeneratorMethod()` which generates consistent generator method names (e.g. `BatchAccount_SpecGenerator`). This is used by all three test case types.

## Phase II: Skeleton

Add a new code generator pipeline stage that executes immediately after the existing stage.

Initially this will do nothing due to a deny list that blocks it for all groups - use the similar deny list in `v2/internal/reconcilers/arm/azure_generic_arm_reconciler_instance.go` as a model.

Additionally, modify the gopter generator in `json_serialization_test_cases` to use the same deny list, but inverted so that we get either gopter or rapid tests for each group.

### Deny List Model

The model to follow is `skipDeletionPrecheck` in `v2/internal/reconcilers/arm/azure_generic_arm_reconciler_instance.go` (line ~920). It's a `sets.NewString(...)` containing group names. The code generator already uses the `set` package from `github.com/Azure/azure-service-operator/v2/internal/set` — use `set.Make[string](...)` and `set.Contains()` for this.

### Detailed Steps

#### Step 1: Add the rapid module dependency

```bash
cd v2
go get pgregory.net/rapid
```

#### Step 2: Add rapid package reference

In `v2/tools/generator/internal/astmodel/std_references.go`, add alongside the gopter references:

```go
RapidReference = MakeExternalPackageReference("pgregory.net/rapid")
```

#### Step 3: Create the migration list

Create `v2/tools/generator/internal/testcases/rapid_migration.go` containing:

- A `set.Set[string]` of group names that should **continue using gopter**. This set is populated with all current groups — any group NOT in this set will use rapid. This means new groups automatically get rapid tests, and we migrate existing groups by removing them from this set.

```go
import "github.com/Azure/azure-service-operator/v2/internal/set"

// gopterGroups is the set of groups that still use gopter-based property tests.
// To migrate a group to rapid, remove it from this set.
var gopterGroups = set.Make(
    "alertsmanagement",
    "apimanagement",
    "app",
    "appconfiguration",
    "authorization",
    "batch",
    "cache",
    "cdn",
    "cognitiveservices",
    "compute",
    "containerinstance",
    "containerregistry",
    "containerservice",
    "datafactory",
    "dataprotection",
    "dbformariadb",
    "dbformysql",
    "dbforpostgresql",
    "devices",
    "documentdb",
    "entra",
    "eventgrid",
    "eventhub",
    "insights",
    "keyvault",
    "kubernetesconfiguration",
    "kusto",
    "machinelearningservices",
    "managedidentity",
    "monitor",
    "network",
    "network.frontdoor",
    "notificationhubs",
    "operationalinsights",
    "quota",
    "redhatopenshift",
    "resources",
    "search",
    "servicebus",
    "signalrservice",
    "sql",
    "storage",
    "subscription",
    "synapse",
    "web",
)
```

- Two exported functions:
  - `UseRapidForGroup(group string) bool` — returns `!gopterGroups.Contains(group)` (groups NOT in the set get rapid)
  - `UseGopterForGroup(group string) bool` — returns `gopterGroups.Contains(group)` (groups IN the set keep gopter)

The group name is the value returned by `InternalPackageReference.Group()` — e.g. `"batch"`, `"compute"`, etc.

This design means:
- All existing groups start on gopter (no generated code changes in Phase II)
- Any new group automatically gets rapid tests
- Migration happens by removing a group from `gopterGroups`

#### Step 4: Modify existing gopter injection to respect the allow-list

In `v2/tools/generator/internal/codegen/pipeline/json_serialization_test_cases.go`, modify `InjectJSONSerializationTests` to skip types whose group is in the rapid migration list.

Extract the group from the type definition's name:
```go
if ref, ok := def.Name().PackageReference().(astmodel.InternalPackageReference); ok {
    if testcases.UseRapidForGroup(ref.Group()) {
        continue // Skip — rapid stage will handle this group
    }
}
```

Place this check inside the loop before `factory.NeedsTest(def)`.

#### Step 5: Create the new pipeline stage

Create `v2/tools/generator/internal/codegen/pipeline/inject_rapid_serialization_tests.go`:

```go
const InjectRapidSerializationTestsStageID = "injectRapidSerializationTests"

func InjectRapidSerializationTests(idFactory astmodel.IdentifierFactory) *Stage {
    stage := NewStage(
        InjectRapidSerializationTestsStageID,
        "Add rapid-based test cases to verify JSON serialization",
        func(ctx context.Context, state *State) (*State, error) {
            // Phase II: no-op for existing groups (all are in gopterGroups)
            // New groups not in gopterGroups would get rapid tests
            return state, nil
        })

    stage.RequiresPostrequisiteStages("simplifyDefinitions")

    return stage
}
```

#### Step 6: Register the new stage

In `v2/tools/generator/internal/codegen/code_generator.go`, add the new stage immediately after `InjectJSONSerializationTests`:

```go
pipeline.InjectJSONSerializationTests(idFactory).UsedFor(pipeline.ARMTarget),
pipeline.InjectRapidSerializationTests(idFactory).UsedFor(pipeline.ARMTarget),
pipeline.InjectPropertyAssignmentTests(idFactory).UsedFor(pipeline.ARMTarget),
pipeline.InjectResourceConversionTestCases(idFactory).UsedFor(pipeline.ARMTarget),
```

Also update `InjectPropertyAssignmentTests` and `InjectResourceConversionTestCases` to add `InjectRapidSerializationTestsStageID` as a prerequisite (since they reuse generators).

#### Step 7: Update golden test expectations

Run `./hack/tools/task generator:unit-tests` (with `-update` flag if needed). The following golden files will need updating to include the new stage:
- `v2/tools/generator/internal/codegen/testdata/TestGolden_NewARMCodeGeneratorFromConfigCreatesRightPipeline.golden`
- `v2/tools/generator/internal/codegen/testdata/TestGolden_NewTestCodeGeneratorCreatesRightPipeline.golden`

### Validation

- `./hack/tools/task quick-checks` passes
- `./hack/tools/task generator:unit-tests` passes (with updated golden files)
- No generated code changes for existing groups (all are in `gopterGroups`)
- Golden file tests use synthetic group names not in `gopterGroups`, so they will automatically exercise the rapid code path. This means: golden file tests that generate test cases will produce rapid-style code. The golden files for those tests will need updating.

## Phase III: Implementation

Expand the skeleton from Phase II so that it fully implements the required tests using rapid. Enable just the "batch" group for this new testing style.

### Key Design Decision: Generator Interface Compatibility

The `PropertyAssignmentTestCase` and `ResourceConversionTestCase` call the generator by name (e.g. `BatchAccount_SpecGenerator()`). Currently this function returns `gopter.Gen`. When we switch a group to rapid, we need to ensure these other test cases still work.

**Approach**: When a group uses rapid, ALL three test case types for that group must use rapid generators. The `XGenerator()` function signature will change from returning `gopter.Gen` to returning `*rapid.Generator[T]`. This means we also need rapid-flavored versions of the property assignment and resource conversion test runners.

This is necessary because:
- The generator return type changes (`gopter.Gen` → `*rapid.Generator[T]`)
- The test runner pattern changes (`gopter.NewProperties()` + `prop.ForAll()` → `rapid.Check()`)
- The test method signature changes (gopter tests take a value and return `string`; rapid tests use `*rapid.T` and call `t.Fatal`)

### Pre-existing Phase II Infrastructure

The following files and changes already exist from Phase II — Phase III builds on them:

**Files already created:**
- `v2/tools/generator/internal/testcases/rapid_migration.go` — The `gopterGroups` set and `UseRapidForGroup()`/`UseGopterForGroup()` functions
- `v2/tools/generator/internal/codegen/pipeline/inject_rapid_serialization_tests.go` — No-op pipeline stage skeleton with `InjectRapidSerializationTestsStageID`

**Files already modified:**
- `v2/tools/generator/internal/astmodel/std_references.go` — `RapidReference` added
- `v2/tools/generator/internal/codegen/pipeline/json_serialization_test_cases.go` — Skips types in rapid groups
- `v2/tools/generator/internal/codegen/pipeline/property_assignment_test_cases.go` — Skips types in rapid groups; `InjectRapidSerializationTestsStageID` added as prerequisite
- `v2/tools/generator/internal/codegen/pipeline/resource_conversion_test_cases.go` — Skips types in rapid groups; `InjectRapidSerializationTestsStageID` added as prerequisite
- `v2/tools/generator/internal/codegen/code_generator.go` — `InjectRapidSerializationTests` stage registered after `InjectJSONSerializationTests`
- `v2/go.mod` and `v2/tools/generator/go.mod` — `pgregory.net/rapid` dependency added
- Golden test files updated for new pipeline stage

### Detailed Steps

#### Step 1: Create `RapidJSONSerializationTestCase`

Create `v2/tools/generator/internal/testcases/rapid_json_serialization_test_case.go` implementing the `astmodel.TestCase` interface.

**Generated test structure** (target output for each type):

```go
func Test_BatchAccount_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
    t.Parallel()
    rapid.Check(t, RunJSONSerializationTestForBatchAccount_Spec)
}

func RunJSONSerializationTestForBatchAccount_Spec(t *rapid.T) {
    subject := BatchAccount_SpecGenerator().Draw(t, "subject")

    bin, err := json.Marshal(subject)
    if err != nil {
        t.Fatal(err)
    }

    var actual BatchAccount_Spec
    err = json.Unmarshal(bin, &actual)
    if err != nil {
        t.Fatal(err)
    }

    match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
    if !match {
        actualFmt := pretty.Sprint(actual)
        subjectFmt := pretty.Sprint(subject)
        result := diff.Diff(subjectFmt, actualFmt)
        t.Errorf(result)
    }
}

var batchAccount_SpecGenerator *rapid.Generator[BatchAccount_Spec]

func BatchAccount_SpecGenerator() *rapid.Generator[BatchAccount_Spec] {
    if batchAccount_SpecGenerator != nil {
        return batchAccount_SpecGenerator
    }

    batchAccount_SpecGenerator = rapid.Custom(func(t *rapid.T) BatchAccount_Spec {
        var result BatchAccount_Spec
        result.AccountName = rapid.Ptr(rapid.String(), true).Draw(t, "AccountName")
        result.Location = rapid.Ptr(rapid.String(), true).Draw(t, "Location")
        result.Identity = IdentityGenerator().Draw(t, "Identity")
        // ... all properties, independent and related, in a single closure
        return result
    })

    return batchAccount_SpecGenerator
}
```

**Important: rapid doesn't have `gen.Struct()`** — we generate struct construction field-by-field inside `rapid.Custom()`. Each property is drawn from its respective rapid generator.

**No `AddIndependentPropertyGeneratorsFor` / `AddRelatedPropertyGeneratorsFor` helpers**: Unlike gopter which populates a `map[string]gopter.Gen` and passes it to `gen.Struct()`, rapid constructs each field inline within the `rapid.Custom()` closure. This eliminates the separate helper methods entirely.

**No two-phase cycle breaking needed**: Gopter's `gen.Struct(reflect.TypeOf(T{}), generators)` requires the generator map to be fully populated at construction time — if `AddRelatedPropertyGeneratorsForX()` calls `YGenerator()` which calls `XGenerator()` recursively, we need a partial-generator already cached. With rapid, `rapid.Custom(fn)` does NOT execute `fn` at construction time — `fn` runs only at draw time. So we can:
1. Create a single `rapid.Custom()` with ALL fields (independent + related) inside the closure
2. Assign it to the global cache variable immediately
3. If another generator's draw triggers a recursive call back to this generator, it finds the cached value

Draw-time cycles (A draws B draws A) terminate naturally because optional fields use `rapid.Ptr(g, true)` which sometimes generates `nil`.

**No >50 property limitation**: Unlike gopter (which uses `gen.Struct()` with reflection and hits a Go runtime limit at ~50 fields — see https://github.com/golang/go/issues/54669), `rapid.Custom()` constructs the struct field-by-field with no reflection. Generate generators for ALL properties regardless of count.

#### Step 2: Implement rapid generator mapping for each property type

Create helper methods on `RapidJSONSerializationTestCase`:

**Independent generators** (`createIndependentGenerator`):
| ASO Type | Gopter | Rapid |
|----------|--------|-------|
| `string` | `gen.AlphaString()` | `rapid.String()` (full UTF-8 — broader coverage than gopter's alpha-only) |
| `uint32` | `gen.UInt32()` | `rapid.Uint32()` |
| `int` | `gen.Int()` | `rapid.Int()` |
| `float64` | `gen.Float64()` | `rapid.Float64()` |
| `bool` | `gen.Bool()` | `rapid.Bool()` |
| `*T` (optional) | `gen.PtrOf(g)` | `rapid.Ptr(g, true)` (`allowNil=true` for optional fields) |
| `[]T` (array) | `gen.SliceOf(g)` | `rapid.SliceOf(g)` |
| `map[K]V` | `gen.MapOf(k, v)` | `rapid.MapOf(k, v)` |
| Enum | `gen.OneConstOf(v1, v2, ...)` | `rapid.SampledFrom([]T{v1, v2, ...})` |
| Validated | generator + `g.Map(castFn)` (method) | `rapid.Map(generator, castFn)` (standalone function, NOT a method) |

**Related generators** (`createRelatedGenerator`):
| ASO Type | Rapid |
|----------|-------|
| `InternalTypeName` (same pkg) | `OtherTypeGenerator()` (same naming via `idOfGeneratorMethod`) |
| `InternalTypeName` (other pkg) | `pkg.OtherTypeGenerator()` |
| `*T` (optional, non-OneOf) | `rapid.Ptr(g, true)` |
| `*T` (optional, OneOf context) | `rapid.Map(g, func(it T) *T { return &it })` — forces non-nil pointer for OneOf members |
| `[]T` (array) | `rapid.SliceOf(g)` |
| `map[K]V` | `rapid.MapOf(keyGen, valueGen)` (key from independent, value from related) |
| `ValidatedType` | Recurse to underlying element type |

**Key API difference**: In gopter, `.Map()` is a method on `gopter.Gen` (e.g. `g.Map(castFn)`). In rapid, `rapid.Map()` is a standalone generic function: `rapid.Map[U, V](g *Generator[U], fn func(U) V) *Generator[V]`. The code generator must emit `rapid.Map(g, fn)` not `g.Map(fn)`.

**Property filtering (`removeByPackage`)**: The same property filtering from the gopter implementation must be applied. Properties with types from these packages are excluded from generation:
- `GenRuntimeReference`, `GenRuntimeConfigMapsReference`, `GenRuntimeSecretsReference`, `GenRuntimeCoreReference`
- `APIMachineryRuntimeReference`, `APIMachinerySchemaReference`
- `APIExtensionsReference`, `APIExtensionsJSONReference`
- `GenRuntimeConditionsReference`

#### Step 3: Generate generic type expressions in AST

The rapid API uses Go generics extensively. The generated code must include generic type expressions like `*rapid.Generator[BatchAccount_Spec]`.

To construct `*rapid.Generator[T]` in the `dave/dst` AST, use `dst.IndexExpr` (used for single-type-parameter generics in Go 1.18+):

```go
// *rapid.Generator[BatchAccount_Spec]
generatorType := &dst.StarExpr{
    X: &dst.IndexExpr{
        X: &dst.SelectorExpr{
            X:   dst.NewIdent(rapidPackage),   // "rapid"
            Sel: dst.NewIdent("Generator"),
        },
        Index: subjectTypeExpr,  // e.g. dst.NewIdent("BatchAccount_Spec")
    },
}
```

This pattern already exists in the codebase — see `kubernetes_admissions_validator.go` (line ~409) which uses `dst.IndexExpr` for generic interface instantiation.

The variable declaration becomes:
```go
// var batchAccount_SpecGenerator *rapid.Generator[BatchAccount_Spec]
astbuilder.VariableDeclaration(globalID, generatorType, comment)
```

The function return type also uses this pattern for the generator factory method.

#### Step 4: Handle OneOf types

For OneOf types, use `rapid.OneOf()` instead of `gen.OneGenOf()`:
- Iterate over properties (each is a OneOf option)
- For each option, create a `rapid.Custom()` that sets just that one property on an otherwise-zero-valued struct
- For optional (pointer) OneOf properties, use `rapid.Map(g, func(it T) *T { return &it })` to force a non-nil pointer (same semantic as the gopter code's `.Map()` wrapper)
- Combine all per-option generators with `rapid.OneOf(g1, g2, ...)` — all generators must be `*rapid.Generator[ParentStruct]`

**Note**: `rapid.OneOf` requires all argument generators to have the same type `*Generator[V]`. Since each per-option generator produces the parent struct (with one field set), they all share the same type.

#### Step 5: Create rapid-flavored PropertyAssignment and ResourceConversion test cases

Since these test cases reuse the generator and the generator return type changes, we need rapid versions.

Create `v2/tools/generator/internal/testcases/rapid_property_assignment_test_case.go`:
```go
func Test_X_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
    t.Parallel()
    rapid.Check(t, func(t *rapid.T) {
        subject := XGenerator().Draw(t, "subject")
        copied := subject.DeepCopy()    // preserve original
        var other OtherVersion
        err := copied.AssignProperties_To_OtherVersion(&other)
        if err != nil { t.Fatalf("AssignTo: %v", err) }
        var actual CurrentVersion
        err = actual.AssignProperties_From_OtherVersion(&other)
        if err != nil { t.Fatalf("AssignFrom: %v", err) }
        match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
        if !match { ... t.Errorf(result) }
    })
}
```

Create `v2/tools/generator/internal/testcases/rapid_resource_conversion_test_case.go`:
```go
func Test_X_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
    t.Parallel()
    rapid.Check(t, func(t *rapid.T) {
        subject := XGenerator().Draw(t, "subject")
        copied := subject.DeepCopy()
        var hub HubVersion
        err := copied.ConvertTo(&hub)
        if err != nil { t.Fatalf("ConvertTo: %v", err) }
        var actual CurrentVersion
        err = actual.ConvertFrom(&hub)
        if err != nil { t.Fatalf("ConvertFrom: %v", err) }
        match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
        if !match { ... t.Errorf(result) }
    })
}
```

These mirror the existing test case types but:
- Use `rapid.Check()` instead of `gopter.NewProperties()` + `prop.ForAll()`
- Use `.Draw(t, "label")` to get values from generators
- Use `t.Fatal()` / `t.Fatalf()` / `t.Errorf()` instead of returning `string`

Both implement the `astmodel.TestCase` interface: `Name()`, `References()`, `RequiredImports()`, `AsFuncs(subject TypeName, ctx *CodeGenerationContext) ([]dst.Decl, error)`, `Equals()`.

Their `RequiredImports()` must include: `testing`, `os`, `astmodel.RapidReference`, `astmodel.CmpReference`, `astmodel.CmpOptsReference`, `astmodel.DiffReference`, `astmodel.PrettyReference`, plus the parameter type or hub type package reference. They do NOT include gopter references.

#### Step 6: Add rapid injection to property assignment and resource conversion pipeline stages

Phase II already added the gopter-skipping logic (`UseRapidForGroup` → `continue`) to both `InjectPropertyAssignmentTests` and `InjectResourceConversionTestCases`. Phase III adds the complementary rapid injection.

After the existing loop that handles gopter test cases, add a second loop (or modify the existing one) to inject rapid-flavored test cases for groups where `UseRapidForGroup()` returns `true`:

```go
// Existing loop already skips rapid groups for gopter injection
// Add rapid injection for rapid groups:
for _, d := range state.Definitions() {
    if ref, ok := d.Name().PackageReference().(astmodel.InternalPackageReference); ok {
        if !testcases.UseRapidForGroup(ref.Group()) {
            continue // Skip — gopter loop already handled this group
        }
    }
    if rapidFactory.NeedsTest(d) {
        updated, err := rapidFactory.AddTestTo(d)
        // ...
    }
}
```

#### Step 7: Implement the rapid JSON serialization pipeline stage action

In `inject_rapid_serialization_tests.go`, flesh out the stage action. The factory should mirror `objectSerializationTestCaseFactory` from `json_serialization_test_cases.go`:

```go
func(ctx context.Context, state *State) (*State, error) {
    factory := makeRapidSerializationTestCaseFactory(idFactory)
    modifiedDefinitions := make(astmodel.TypeDefinitionSet)
    var errs []error
    for _, def := range state.Definitions() {
        ref, ok := def.Name().PackageReference().(astmodel.InternalPackageReference)
        if !ok || !testcases.UseRapidForGroup(ref.Group()) {
            continue
        }
        if factory.NeedsTest(def) {
            updated, err := factory.AddTestTo(def)
            if err != nil {
                errs = append(errs, err)
            } else {
                modifiedDefinitions[updated.Name()] = updated
            }
        }
    }
    if len(errs) > 0 {
        return nil, kerrors.NewAggregate(errs)
    }
    return state.WithOverlaidDefinitions(modifiedDefinitions), nil
}
```

The factory's `NeedsTest()` must replicate the same checks as `objectSerializationTestCaseFactory.NeedsTest()`:
- Type must be a property container (`astmodel.AsPropertyContainer`)
- Must not be a webhook type (`astmodel.IsWebhookPackageReference`)
- Must not be in the suppressions list (ARM OneOf types that don't round-trip)
- Empty property containers still get test cases (generators are needed by other types)

#### Step 8: Enable "batch" group

Remove `"batch"` from the `gopterGroups` set in `rapid_migration.go`. This causes all batch types to get rapid-based tests instead of gopter-based tests.

**Note**: Golden file tests use a synthetic group name (`"person"`) that is NOT in `gopterGroups`, so they already exercise the rapid code path. Validate with golden tests BEFORE enabling a real group.

#### Step 9: Regenerate and validate

```bash
# First validate with golden tests (synthetic "person" group uses rapid path)
./hack/tools/task generator:unit-tests

# Update golden files if needed
./hack/tools/task generator:update-golden-tests

# Then enable batch and regenerate
./hack/tools/task controller:generate-types

# Build to verify compilation
./hack/tools/task controller:build

# Run the generated tests for batch specifically
cd v2 && go test ./api/batch/... -v -count=1

# Run generator unit tests
./hack/tools/task generator:unit-tests

# Full test suite
./hack/tools/task controller:test
```

### File Summary

**New files to create (Phase III):**
- `v2/tools/generator/internal/testcases/rapid_json_serialization_test_case.go` — The rapid JSON serialization test case
- `v2/tools/generator/internal/testcases/rapid_property_assignment_test_case.go` — Rapid-flavored property assignment test case
- `v2/tools/generator/internal/testcases/rapid_resource_conversion_test_case.go` — Rapid-flavored resource conversion test case

**Files to modify (Phase III):**
- `v2/tools/generator/internal/testcases/rapid_migration.go` — Remove `"batch"` from `gopterGroups`
- `v2/tools/generator/internal/codegen/pipeline/inject_rapid_serialization_tests.go` — Flesh out the no-op stage with the factory implementation
- `v2/tools/generator/internal/codegen/pipeline/property_assignment_test_cases.go` — Add rapid test case injection loop (gopter skipping already done in Phase II)
- `v2/tools/generator/internal/codegen/pipeline/resource_conversion_test_cases.go` — Add rapid test case injection loop (gopter skipping already done in Phase II)
- Golden test files — Update for new generated test case output

## Resolved Questions

These questions have been asked and answered — kept for reference.

1. **Rapid import path**: Use `pgregory.net/rapid` as the Go import path.

2. **Test case scope**: All three test case types (JSON serialization, property assignment, resource conversion) will be migrated. The goal is to eventually remove all gopter tests.

3. **String generation**: Use `rapid.String()` (full UTF-8). We expect serialization and conversion code to handle UTF-8 properly.

4. **Large objects (>50 properties)**: Generate independent generators for all properties — no skip. The gopter limitation (reflection on large structs) doesn't apply to rapid.

5. **Test iteration count**: Use rapid's default of 100 for all types. No graduated counts.

6. **Golden file tests**: The inverted list design (gopter set lists known groups; anything else gets rapid) means golden file tests with synthetic group names automatically exercise the rapid code path.

7. **`rapid.Ptr` not `rapid.Pointer`**: The correct function name is `rapid.Ptr[E any](elem *Generator[E], allowNil bool) *Generator[*E]`. There is no `rapid.Pointer` function.

8. **`rapid.Map` is a standalone function**: Unlike gopter where `.Map()` is a method on `gopter.Gen`, rapid's `Map` is a standalone generic function: `rapid.Map[U, V](g *Generator[U], fn func(U) V) *Generator[V]`. The code generator must emit `rapid.Map(g, fn)` not `g.Map(fn)`.

9. **No two-phase cycle breaking**: Unlike gopter's `gen.Struct()` which requires generators at construction time (necessitating a two-phase pattern), `rapid.Custom()` defers execution to draw time. A single `rapid.Custom()` with all fields suffices; no partial generator assignment is needed. See Phase III Step 1 for details.

10. **`rapid.Deferred` exists but doesn't solve cycles**: `rapid.Deferred[V](fn func() *Generator[V]) *Generator[V]` defers generator construction, but if the deferred function triggers recursive access to the same generator, it will still recurse infinitely. The single-phase `rapid.Custom()` + global cache approach is the correct cycle-breaking strategy.

## Appendix: astbuilder Vocabulary

The `v2/tools/generator/internal/astbuilder` package provides helpers for constructing Go AST nodes (using the `dave/dst` library). These are the primary tools for generating test code.

### Function Declarations
- `NewTestFuncDetails(testingPackage, testName, body...)` — Creates a `func Test_...(t *testing.T)` function
- `FuncDetails.DefineFunc()` — Emits a complete function declaration
- `FuncDetails.AddStatements(stmts...)` — Appends statements to body
- `FuncDetails.AddParameter(id, type)` — Adds a named parameter
- `FuncDetails.AddReturns(types...)` — Adds return types
- `FuncDetails.AddComments(comment...)` — Adds doc comments

### Assignments
- `SimpleAssignment(lhs, rhs)` — `lhs = rhs`
- `ShortDeclaration(id, rhs)` — `id := rhs`
- `QualifiedAssignment(lhs, sel, tok, rhs)` — `lhs.sel = rhs`
- `SimpleAssignmentWithErr(lhs, tok, rhs)` — `lhs, err := rhs`

### Function Calls
- `CallFunc(name, args...)` — `name(args...)`
- `CallQualifiedFunc(pkg, name, args...)` — `pkg.name(args...)`
- `CallExpr(expr, name, args...)` — `expr.name(args...)`
- `CallFuncAsStmt(...)` / `CallQualifiedFuncAsStmt(...)` / `CallExprAsStmt(...)` — Call as statement

### Declarations & Variables
- `NewVariable(name, structName)` — `var name structName`
- `VariableDeclaration(ident, type, comment)` — Global `var ident type` with comment
- `LocalVariableDeclaration(ident, type, comment)` — Local `var ident type`

### Control Flow
- `SimpleIf(cond, stmts...)` — `if cond { ... }`
- `ReturnIfNotNil(check, returns...)` — `if check != nil { return ... }`
- `ReturnIfNil(check, returns...)` — `if check == nil { return ... }`
- `Returns(exprs...)` — `return expr, ...`
- `IterateOverSlice(item, list, stmts...)` — `for _, item := range list { ... }`

### Literals
- `StringLiteral(content)` / `StringLiteralf(format, args...)` — Quoted string literal
- `IntLiteral(value)` — Integer literal

### Expressions
- `QualifiedTypeName(pkg, name)` — `pkg.name`
- `Selector(expr, names...)` — `expr.name1.name2...`
- `AddrOf(expr)` — `&expr`
- `Dereference(expr)` / `PointerTo(expr)` — `*expr`
- `NotExpr(expr)` — `!expr`
- `AreEqual(lhs, rhs)` / `AreNotEqual(lhs, rhs)` — `==` / `!=`

### Collections
- `MakeMap(key, value)` — `make(map[key]value)`
- `InsertMap(map, key, rhs)` — `map[key] = rhs`
- `SliceLiteral(type, items...)` — `[]type{items...}`
- `AppendItemToSlice(lhs, rhs)` — `lhs = append(lhs, rhs)`

### Composite Literals
- `NewCompositeLiteralBuilder(type)` — Builds `Type{ Field: value, ... }`
- `.AddField(name, value)` — Adds `name: value`
- `.Build()` — Emits the literal

### Statements
- `Statements(stmts...)` — Flatten mixed `dst.Stmt`/`[]dst.Stmt` into a cloned slice
- `StatementBlock(stmts...)` — Wrap in `{ ... }`

### Comments
- `AddComment(list, comment)` — Adds `// comment`
- `AddWrappedComment(list, comment)` — Adds word-wrapped comment

