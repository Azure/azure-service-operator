---
description: "**WORKFLOW SKILL** — Migrate an Azure resource group from legacy to hybrid versioning mode in ASO. USE FOR: moving a group (e.g., appconfiguration, cache, network) from VersionMigrationModeLegacy to VersionMigrationModeHybrid in the code generator, updating samples, renaming CRUD tests, and re-recording sample tests. DO NOT USE FOR: adding new resources (use new-resource.instructions.md), general debugging, or code review."
---

# Skill: Migrate Resource Group from Legacy to Hybrid Versioning

This skill guides the migration of an Azure resource group from `VersionMigrationModeLegacy` to `VersionMigrationModeHybrid` in Azure Service Operator. Hybrid mode generates new-style version packages (`v20XXXXXX`) for all resources while retaining backward-compatible `v1api20XXXXXX` packages for resources introduced in v2.16.0 and earlier.

## Prerequisites

- You know the **group name** (lowercase, e.g., `appconfiguration`, `cache`, `web`) to migrate.
- You can run `task` commands from the repository root.
- For re-recording sample tests, you need Azure credentials configured (`AZURE_TENANT_ID`, `AZURE_SUBSCRIPTION_ID`, etc.).

## Step-by-Step Procedure

### Step 1: Update legacy.go

**File:** `v2/tools/generator/internal/astmodel/legacy.go`

Find the target group's entry in the `versionMigrationModes` map and change it from `VersionMigrationModeLegacy` to `VersionMigrationModeHybrid`.

**Example:**

```go
// Before
"appconfiguration": VersionMigrationModeLegacy,

// After
"appconfiguration": VersionMigrationModeHybrid,
```

> **Important:** The map is sorted alphabetically. Do not reorder entries.

### Step 2: Run the Code Generator

Run the generator to produce the new `v20XXXXXX` packages alongside the existing `v1api20XXXXXX` packages:

```bash
task controller:generate-types
```

This creates new version directories under `v2/api/<group>/` (e.g., `v2/api/appconfiguration/v20220501/` alongside the existing `v2/api/appconfiguration/v1api20220501/`).

Verify the new directories were created:

```bash
ls v2/api/<group>/
```

You should see both `v1api*` and `v*` directories for each version.

### Step 3: Copy and Update Samples

For each existing `v1api<version>/` directory under `v2/samples/<group>/`:

1. **Create a new directory** named `v<version>/` (dropping the `v1api` prefix, so `v1api20220301` becomes `v20220301`).
2. **Copy all YAML files** from the `v1api<version>/` directory to the new `v<version>/` directory.
3. **Rename each file** from `v1api<version>_<resource>.yaml` to `v<version>_<resource>.yaml`.
4. **Update the `apiVersion` field** in each copied YAML file from `<group>.azure.com/v1api<version>` to `<group>.azure.com/v<version>`.

**Example for `web` group, version `20220301`:**

```bash
# Create new directory
mkdir -p v2/samples/web/v20220301

# Copy and rename files
cp v2/samples/web/v1api20220301/v1api20220301_serverfarm.yaml v2/samples/web/v20220301/v20220301_serverfarm.yaml
cp v2/samples/web/v1api20220301/v1api20220301_site.yaml v2/samples/web/v20220301/v20220301_site.yaml
cp v2/samples/web/v1api20220301/v1api20220301_sitessourcecontrol.yaml v2/samples/web/v20220301/v20220301_sitessourcecontrol.yaml
```

Then edit each new file to update `apiVersion`:

```yaml
# Before
apiVersion: web.azure.com/v1api20220301

# After
apiVersion: web.azure.com/v20220301
```

All other fields remain **identical** — only the `apiVersion` line changes.

> **Tip:** If a group has multiple version directories (e.g., `v1api20220301/`, `v1api20230101/`), repeat this for each one.

### Step 4: Rename CRUD Tests

Not all CRUD test files use the `v1api` prefix in their names or function signatures. Many tests already use the `v20*` naming convention (e.g., `Test_Web_Site_v20220301_CRUD`). **If a test file and its function names do NOT contain `v1api`, no renaming is needed** — skip to Step 5 for that test.

Only perform the steps below for tests that **do** have `v1api` in the test function name, file name, or import paths:

#### 4a. Rename the test file

If the test file name contains `v1api`, rename it to drop the prefix:

```bash
# Example
mv v2/internal/controllers/appconfiguration_keyvalue_v1api20240601_crud_test.go \
   v2/internal/controllers/appconfiguration_keyvalue_v20240601_crud_test.go
```

> **Note:** Some test files may already use the `v20*` naming in the filename even though they have `v1api` references in the code. In that case, no file rename is needed — just update the contents.

#### 4b. Update import paths

Change Go import paths from the `v1api<version>` package to the `v<version>` package:

```go
// Before
appconfig "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1api20220501"

// After
appconfig "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v20220501"
```

Do this for **all imports** in the file that reference the migrated group's legacy version packages.

#### 4c. Update test function names (if needed)

Rename all test function names and helper function names to drop the `v1api` prefix from the version:

```go
// Before
func Test_AppConfiguration_KeyValue_v1api20240601_CRUD(t *testing.T) {
// After
func Test_AppConfiguration_KeyValue_v20240601_CRUD(t *testing.T) {
```

Also update any helper functions or subtest references within the file:

```go
// Before
AppConfiguration_KeyValue_v1api20240601_CRUD(tc, cs)
// After
AppConfiguration_KeyValue_v20240601_CRUD(tc, cs)
```

> **Note:** Some tests may already use the `v20*` naming in the test name even though they have `v1api` references in the code. In that case, no test function rename is needed. You can also skip Step 5 as well.

### Step 5: Re-record CRUD Tests

> **STOP:** You **MUST** have the following environment variables set to record the tests: `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and `ENTRA_APP_ID`. If any are missing, do not proceed because the tests will fail. Tell the user what is missing and stop the process.

If the test function was renamed in Step 4, the old recording files no longer match. Delete the old recording files and re-record the tests by running them against Azure:

```bash
# Delete old recording(s)
rm -f "v2/internal/controllers/recordings/Test_AppConfiguration_KeyValue_v1api20240601_CRUD.yaml"
# If subtests exist, delete the directory too
rm -rf "v2/internal/controllers/recordings/Test_AppConfiguration_KeyValue_v1api20240601_CRUD/"

# Re-record with the new test name
TEST_FILTER=Test_AppConfiguration_KeyValue_v20240601_CRUD task controller:test-integration-envtest
```

The general pattern is:

```bash
TEST_FILTER=<NewTestFunctionName> task controller:test-integration-envtest
```

> **Important:** If a test function name did NOT have `v1api` in it (some groups already used `v20*` naming), no re-recording is needed.

### Step 6: Run the Samples Tests

Run the samples tests in record mode to create (or re-record) the sample test recordings for the new `v<version>` sample directories:

```bash
TEST_FILTER='Test_Samples_CreationAndDeletion/Test_<GroupNameWithLeadingCapital>*' task controller:test-samples
```

**Examples:**

```bash
# For appconfiguration
TEST_FILTER='Test_Samples_CreationAndDeletion/Test_Appconfiguration*' task controller:test-samples

# For web
TEST_FILTER='Test_Samples_CreationAndDeletion/Test_Web*' task controller:test-samples

# For cache
TEST_FILTER='Test_Samples_CreationAndDeletion/Test_Cache*' task controller:test-samples
```

The test name format is `Test_<GroupNameWithLeadingCapital>_<version>_CreationAndDeletion`. The group name is title-cased using `cases.Title(language.English)`, which capitalizes only the first letter (e.g., `appconfiguration` → `Appconfiguration`). For groups with dots in the name (e.g., `network.frontdoor`), each dot-separated segment is title-cased independently (e.g., `Test_NetworkFrontdoor_v20220501_CreationAndDeletion`).

After recording, tidy the recordings:

```bash
task controller:tidy-samples-recordings
```

### Step 7: Verify

Run the full verification sequence as specified in the project's copilot instructions:

```bash
task format-code
task generator:quick-checks
task controller:quick-checks
```

Fix any issues and re-run from the beginning if anything fails.

## Summary of Changes

| Area | What Changes |
|------|-------------|
| `legacy.go` | Group entry: `VersionMigrationModeLegacy` → `VersionMigrationModeHybrid` |
| `v2/api/<group>/` | New `v<version>/` directories generated alongside existing `v1api<version>/` |
| `v2/samples/<group>/` | New `v<version>/` directories with copied+updated sample YAML files |
| CRUD test files | File names, import paths, function names updated to drop `v1api` prefix |
| Recording files | Renamed to match updated test function names |
| Sample test recordings | New recordings created for `v<version>` sample directories |

## Common Mistakes

- **Don't forget subtest helper functions.** If a test file has helper functions like `AppConfiguration_KeyValue_v1api20240601_CRUD(tc, cs)`, those need renaming too.
- **Don't modify generated files by hand.** `*_gen.go` files are regenerated by `task controller:generate-types`. Only edit `legacy.go` and non-generated files.
- **Check all version directories.** If a group has multiple API versions (e.g., `v1api20220501` and `v1api20240601`), create hybrid samples and update tests for all of them.
- **Recording files must match test names exactly.** The test framework looks up recordings by `t.Name()`. A mismatch will cause tests to re-record or fail.
- **Title-case the group name correctly in `TEST_FILTER`.** The test name uses `cases.Title(language.English)`, which only capitalizes the first letter: `appconfiguration` → `Appconfiguration`, `containerservice` → `Containerservice`.
