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

The model to follow is `skipDeletionPrecheck` in `v2/internal/reconcilers/arm/azure_generic_arm_reconciler_instance.go` (line ~920). It's a `sets.NewString(...)` containing group names. However, since the code generator is in a different module (`v2/tools/generator`), we should use a plain `map[string]bool` or a set type available within the generator module.

UPDATE: the code generator already uses a proper "set" abstraction, we should use that instead of monkeying around with a `map[string]bool`.

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

#### Step 3: Create the migration allow-list

Create `v2/tools/generator/internal/testcases/rapid_migration.go` containing:

- A `map[string]bool` (or `set.Set[string]`) of group names that should use rapid instead of gopter. Initially, this set is **empty** (all groups use gopter).

UPDATE: We don't want any new groups to use gopter, we want them to use rapid. Populate the set will all current groups and only generate gopter tests for groups IN the set; generate rapid tests for any group not in the set.

- Two exported functions:
  - `UseRapidForGroup(group string) bool` — returns true if the group should generate rapid tests
  - `UseGopterForGroup(group string) bool` — returns `!UseRapidForGroup(group)` (the inverse)

The group name is the value returned by `InternalPackageReference.Group()` — e.g. `"batch"`, `"compute"`, etc.

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
            // Phase II: no-op — the allow-list is empty so no groups match
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
- No generated code changes (since the rapid list is empty, all groups still use gopter)

## Phase III: Implementation

Expand the skeleton from Phase II so that it fully implements the required tests using rapid. Enable just the "batch" group for this new testing style.

### Key Design Decision: Generator Interface Compatibility

The `PropertyAssignmentTestCase` and `ResourceConversionTestCase` call the generator by name (e.g. `BatchAccount_SpecGenerator()`). Currently this function returns `gopter.Gen`. When we switch a group to rapid, we need to ensure these other test cases still work.

**Approach**: When a group uses rapid, ALL three test case types for that group must use rapid generators. The `XGenerator()` function signature will change from returning `gopter.Gen` to returning `*rapid.Generator[T]`. This means we also need rapid-flavored versions of the property assignment and resource conversion test runners.

This is necessary because:
- The generator return type changes (`gopter.Gen` → `*rapid.Generator[T]`)
- The test runner pattern changes (`gopter.NewProperties()` + `prop.ForAll()` → `rapid.Check()`)
- The test method signature changes (gopter tests take a value and return `string`; rapid tests use `*rapid.T` and call `t.Fatal`)

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

    generators := make(map[string]gopter.Gen)  // See note below about generator map type
    AddIndependentPropertyGeneratorsForBatchAccount_Spec(generators)

    // Phase 1: assign partial generator to break cycles
    batchAccount_SpecGenerator = rapid.Custom(func(t *rapid.T) BatchAccount_Spec {
        // Generate struct using independent generators only
        ...
    })

    AddRelatedPropertyGeneratorsForBatchAccount_Spec(generators)

    // Phase 2: replace with full generator
    batchAccount_SpecGenerator = rapid.Custom(func(t *rapid.T) BatchAccount_Spec {
        // Generate struct using all generators
        ...
    })

    return batchAccount_SpecGenerator
}
```

**Important: rapid doesn't have `gen.Struct()`** — we need to generate struct construction field-by-field inside `rapid.Custom()`. Each property is drawn from its respective rapid generator.

#### Step 2: Implement rapid generator mapping for each property type

Create helper methods on `RapidJSONSerializationTestCase`:

**Independent generators** (`createIndependentGenerator`):
| ASO Type | Gopter | Rapid |
|----------|--------|-------|
| `string` | `gen.AlphaString()` | `rapid.StringMatching("[a-zA-Z]*")` (see Q5) |
| `uint32` | `gen.UInt32()` | `rapid.Uint32()` |
| `int` | `gen.Int()` | `rapid.Int()` |
| `float64` | `gen.Float64()` | `rapid.Float64()` |
| `bool` | `gen.Bool()` | `rapid.Bool()` |
| `*T` (optional) | `gen.PtrOf(g)` | `rapid.Pointer(g, true)` |
| `[]T` (array) | `gen.SliceOf(g)` | `rapid.SliceOf(g)` |
| `map[K]V` | `gen.MapOf(k, v)` | `rapid.MapOf(k, v)` |
| Enum | `gen.OneConstOf(v1, v2, ...)` | `rapid.SampledFrom([]T{v1, v2, ...})` |
| Validated | generator + `.Map()` cast | generator + `rapid.Map()` or `rapid.Custom()` cast |

**Related generators** (`createRelatedGenerator`):
- For types in same package: call `OtherTypeGenerator()` (same naming convention via `idOfGeneratorMethod`)
- For types in other packages: call `pkg.OtherTypeGenerator()`
- The return type changes from `gopter.Gen` to `*rapid.Generator[T]`, but the function call pattern is the same.

#### Step 3: Handle the two-phase generator pattern for cycle breaking

The cycle-breaking pattern works with rapid using `rapid.Custom()`:

1. Create a `rapid.Custom(func(t *rapid.T) T { ... })` with only independent fields set
2. Assign it to the global `*rapid.Generator[T]` variable
3. Create a new `rapid.Custom(func(t *rapid.T) T { ... })` with all fields (including related types that may call back into this generator)
4. Assign the full generator to the global variable

Since `rapid.Custom()` captures its closure at creation time, and the global variable is read inside other generators' closures at draw-time (not creation-time), this correctly breaks cycles.

#### Step 4: Handle OneOf types

For OneOf types, use `rapid.OneOf()` instead of `gen.OneGenOf()`:
- Iterate over properties (each is a OneOf option)
- Create a `rapid.Custom()` for each that sets just that one property
- Combine with `rapid.OneOf(g1, g2, ...)`

#### Step 5: Create rapid-flavored PropertyAssignment and ResourceConversion test cases

Since these test cases reuse the generator and the generator return type changes, we need rapid versions:

Create `v2/tools/generator/internal/testcases/rapid_property_assignment_test_case.go`:
```go
func Test_X_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
    t.Parallel()
    rapid.Check(t, func(t *rapid.T) {
        subject := XGenerator().Draw(t, "subject")
        // ... assignment round-trip test ...
    })
}
```

Create `v2/tools/generator/internal/testcases/rapid_resource_conversion_test_case.go`:
```go
func Test_X_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
    t.Parallel()
    rapid.Check(t, func(t *rapid.T) {
        subject := XGenerator().Draw(t, "subject")
        // ... conversion round-trip test ...
    })
}
```

These mirror the existing test case types but:
- Use `rapid.Check()` instead of `gopter.NewProperties()` + `prop.ForAll()`
- Use `.Draw(t, "label")` to get values from generators
- Use `t.Fatal()` / `t.Errorf()` instead of returning `string`

#### Step 6: Modify property assignment and resource conversion pipeline stages

Update `InjectPropertyAssignmentTests` and `InjectResourceConversionTestCases` in their respective pipeline stage files to also check `UseRapidForGroup()`. When a group uses rapid, inject the rapid-flavored test case instead of the gopter-flavored one.

#### Step 7: Implement the rapid JSON serialization pipeline stage action

In `inject_rapid_serialization_tests.go`, flesh out the stage action:

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

#### Step 8: Enable "batch" group

Update the migration allow-list in `rapid_migration.go`:

```go
var rapidMigrationGroups = map[string]bool{
    "batch": true,
}
```

UPDATE: To match earlier changes, this step will involve removing "batch" from the existing list.

#### Step 9: Regenerate and validate

```bash
# Regenerate all types (will produce rapid tests for batch, gopter for everything else)
./hack/tools/task controller:generate-types

# Build to verify compilation
./hack/tools/task controller:build

# Run the generated tests for batch specifically
cd v2 && go test ./api/batch/... -v -count=1

# Run generator unit tests (will need golden file updates)
./hack/tools/task generator:unit-tests

# Full test suite
./hack/tools/task controller:test
```

### File Summary

**New files to create:**
- `v2/tools/generator/internal/testcases/rapid_migration.go` — The group allow/deny list
- `v2/tools/generator/internal/testcases/rapid_json_serialization_test_case.go` — The rapid JSON serialization test case
- `v2/tools/generator/internal/testcases/rapid_property_assignment_test_case.go` — Rapid-flavored property assignment test case
- `v2/tools/generator/internal/testcases/rapid_resource_conversion_test_case.go` — Rapid-flavored resource conversion test case
- `v2/tools/generator/internal/codegen/pipeline/inject_rapid_serialization_tests.go` — The pipeline stage

**Files to modify:**
- `v2/tools/generator/internal/astmodel/std_references.go` — Add `RapidReference`
- `v2/tools/generator/internal/codegen/pipeline/json_serialization_test_cases.go` — Skip rapid groups
- `v2/tools/generator/internal/codegen/pipeline/property_assignment_test_cases.go` — Use rapid version for rapid groups
- `v2/tools/generator/internal/codegen/pipeline/resource_conversion_test_cases.go` — Use rapid version for rapid groups
- `v2/tools/generator/internal/codegen/code_generator.go` — Register new pipeline stage
- `v2/go.mod` / `v2/go.sum` — Add rapid dependency
- Golden test files — Update pipeline stage lists

## Open Questions

1. **Rapid import path**: The Go module path for rapid is `pgregory.net/rapid` (this is what `go get` uses). The GitHub URL is `github.com/flyingmutant/rapid`. We should use `pgregory.net/rapid` as the import path. Please confirm.

A: The correct import path is "pgregory.net/rapid" as suggested.

2. **ResourceConversionTestCase and PropertyAssignmentTestCase scope**: These two test case types also use gopter and **share the generators** created by `JSONSerializationTestCase`. Phase III above includes creating rapid-flavored versions of both. Is this the intended scope, or should we limit Phase III to just the JSON serialization test and leave property assignment / resource conversion tests using gopter? (Note: limiting to just JSON serialization is NOT possible without significant refactoring — the generator return type changes, which breaks the call from these other tests.)

A: The goal is to eventually remove all gopter tests, so conversion of all test styles is appropriate.

3. **String generation strategy**: Gopter uses `gen.AlphaString()` (only `[a-zA-Z]` characters). Rapid's `rapid.String()` generates arbitrary Unicode. Should we constrain rapid to match gopter's behavior (using `rapid.StringMatching("[a-zA-Z]*")`) or embrace broader Unicode fuzzing? Broader fuzzing may surface JSON encoding issues with special characters but could also produce false positives if the types don't handle Unicode properly.

A: We expect all the serialization and conversion code to handle UTF-8 encoded text properly, so let's test all of that.

4. **Large object handling (>50 properties)**: Gopter skips independent generators for types with >50 properties due to a Go runtime limitation with `reflect.TypeOf` on large structs. Since rapid uses `rapid.Custom()` with field-by-field construction (no reflection), this limitation likely doesn't apply. Should we generate independent generators for large objects in the rapid version? This would improve test coverage.

A: Improve the test coverage by including large objects in the tests.

5. **Test iteration count**: Gopter is configured with `MinSuccessfulTests` (80-100 depending on type). Rapid defaults to 100 iterations. Should we keep the default 100 for all types, or replicate the graduated counts (20 for resources, 80 for spec/status)?

A: Keep the default of 100 for now. We expect rapid to be much quicker, so the additional tests won't be a performance issue.

6. **Golden file test coverage**: The existing golden file tests use synthetic test schemas (not real Azure resource groups like "batch"). They will exercise the rapid code path only if we either (a) add a synthetic test that uses a group name in the allow-list, or (b) use a different mechanism to force rapid generation in tests. Which approach is preferred? Option (a) is cleaner.

A: I've requested that a deny list is used - any group not known to require gopter tests should receive rapid tests; as long as the golden file tests use a unique group name, they should automatically get tests written with rapid.


QUESTION: A followup question from me - have you reviewed the helper functions in the astbuilder package - these form a vocabularly for your use when creating the abstract-syntax-tree for the generated Go code.

