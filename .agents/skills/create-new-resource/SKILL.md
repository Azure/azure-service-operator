---
name: create-new-resource
description: Use when adding a new Azure resource (or a new API version of an existing resource) to Azure Service Operator, or when a GitHub issue requests that ASO support a new ARM resource type.
---

# Add a New Code-Generated Resource to ASO

Adding a new resource is a **generator-driven** workflow: 95% of the Go code is produced by the ASO code generator from Azure OpenAPI specs. Your job is to configure the generator, iterate on its errors, hand-write a CRUD test, and record a sample. Most of the work happens in `v2/azure-arm.yaml` and its per-group imports under `v2/azure-arm/`, not in Go.

The workflow is not linear — writing the test frequently sends you back to update generator configuration. Expect to loop.

For the canonical narrative reference (with screenshots and background), see [`docs/hugo/content/contributing/add-a-new-code-generated-resource/`](../../../docs/hugo/content/contributing/add-a-new-code-generated-resource/_index.md).

## Prerequisites

Verify all of these **before** starting. Missing any one of them will stall the workflow partway through.

1. **Resource identified.** From the issue (or user), extract the Azure type and API version, e.g. `Microsoft.Synapse/workspaces` at `2021-06-01`. If not given, look up the resource in an ARM/Bicep template or in [`azure-rest-api-specs`](https://github.com/Azure/azure-rest-api-specs/tree/main/specification). Prefer the latest **stable** (non-preview) `api-version` unless the issue requires a preview feature.
2. **GVK derived.** Convert Azure type + version to Kubernetes GVK:
   - **group** — resource provider, `Microsoft.` prefix stripped, lowercased. `Microsoft.Synapse` → `synapse`.
   - **version** — `v` prefix + `api-version` with non-alphanumeric characters removed. `2023-07-01` → `v20230701`; `2021-04-01-preview` → `v20210401preview`. (The older `v1api` prefix only applies to resources introduced before ASO v2.16 — the migration modes tracked in [`v2/tools/generator/internal/astmodel/legacy.go`](../../../v2/tools/generator/internal/astmodel/legacy.go) exist purely to keep those grandfathered. Any resource or version added today uses plain `v`, regardless of that file.)
   - **kind** — Azure type minus resource provider, singularised. `Microsoft.Synapse/workspaces` → `Workspace`.
3. **Repository ready for generation.** From the repo root:
   - `git tag --list 'v2*'` returns tags. If empty, run `git fetch --all --tags`. Builds fail without tags.
   - `ls v2/specs/azure-rest-api-specs/specification` shows content. If empty, run `git submodule init && git submodule update`. The generator will not run without the submodule.
4. **Next release identified.** Check the project's [GitHub milestones](https://github.com/Azure/azure-service-operator/milestones) for the next unreleased ASO version (e.g. `v2.21.0`). You need this for the `$supportedFrom:` directive.
5. **Task tool on PATH.** `which task` succeeds. If not, your environment is not set up — stop and ask the user to fix it. Never fall back to `./hack/tools/task`.
6. **Azure credentials (only needed for recording).** For **Step 6** (recording the CRUD test) and **Step 7** (recording the sample), `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and often `ENTRA_APP_ID` must be set (typically via a `test.env` file). If missing when you reach those steps, stop and ask the user — do not attempt to record without them.

## Workflow

### Step 1: Configure the code generator

Edit the group's config file under `v2/azure-arm/<group>.yaml` (e.g. `v2/azure-arm/synapse.yaml`). If the group doesn't exist, create the file and add a line under `imports:` in `v2/azure-arm.yaml`:

```yaml
imports:
  # ...alphabetical...
  synapse: azure-arm/synapse.yaml
```

Within the group file, add the resource under the correct version, keeping versions in numerical order. Per-group files use `groupModelConfiguration:` at the top level (the group name is implicit in the filename); versions nest directly under it:

```yaml
groupModelConfiguration:
  2021-06-01:
    Workspace:
      $exportAs: Workspace
      $supportedFrom: v2.21.0   # the next unreleased ASO version from Prerequisites
```

Notes:

- `$exportAs` renames the Kubernetes type. Use it to shorten the Azure name when the group already disambiguates (e.g. `SynapseWorkspace` → `Workspace`).
- **Never** run a YAML formatter over `azure-arm.yaml` or its imports. A prettification diff makes new-resource PRs unreviewable.

**Done when:** the resource, its `$exportAs`, and its `$supportedFrom` are present under the correct group + version, and (if the group file is new) it is imported from `v2/azure-arm.yaml`.

### Step 2: Iterate the generator until it succeeds

Run:

```bash
task generator:quick-checks
```

The first run is **expected to fail** with configuration errors. This is the working loop of the whole skill — the generator produces one error class at a time, you fix the config, you re-run. Do not treat the first failure as a real problem; treat it as the next instruction.

Continue looping until `task generator:quick-checks` exits cleanly. See **Troubleshooting → Generator errors** below for the common error classes and their fixes.

**Done when:** `task generator:quick-checks` succeeds with no errors, AND new files exist under `v2/api/<group>/<version>/` for your resource.

### Step 3: Review the generated code

Open `v2/api/<group>/<version>/<kind>_types_gen.go` (also skim `structure.txt` in the same folder for a tree view). Check the generated `Spec` and `Status` for four problems the OpenAPI spec often gets wrong. Every one you find sends you back to Step 1/2 with new config.

- **Properties on Spec that should be read-only.** Any field the user cannot legally set (provisioning state, IDs generated by Azure, etc.) must be removed from Spec via a `remove: true` directive in the `typeTransformers` section of `azure-arm.yaml`.
- **Passwords/keys typed as `string`.** Anything a user supplies as a secret must become a `genruntime.SecretReference`. Mark it in config with `$importSecretMode: required` on the property (legal values: `required`, `optional`, `never`). The Azure spec usually annotates these with `x-ms-secret` and the generator handles them automatically, but not always. **Do not use the deprecated `$isSecret: true`** — it still parses for backward compatibility but is being removed.
- **Azure-generated secrets on Status.** Keys, connection strings, endpoints that Azure produces on creation are **not** auto-detected. If your resource has any (check ARM docs), declare them with `$azureGeneratedSecrets:` on the resource. This also requires a hand-written `KubernetesSecretExporter` extension under `v2/api/<group>/customizations/`.
- **Runtime-configurable properties.** IDs like `PrincipalID`/`TenantID` that another resource produces at runtime should use `$importConfigMapMode: optional` (or `required`) so users can wire them from a ConfigMap.

For any change, edit `azure-arm.yaml` (or the group import), re-run Step 2, and re-review.

**Done when:** every property on Spec is genuinely user-settable; every user-supplied secret is a `SecretReference`; every Azure-generated secret is declared; and no runtime-only value is a hard-coded string.

### Step 4 (optional): Implement extensions

Most resources need no extension. Skip this step unless testing later reveals odd behaviour (`BadRequest` treated as permanent when it should retry, missing post-creation checks, secret export, etc.).

Extensions live in `v2/api/<group>/customizations/`. Available extension points (`ARMResourceModifier`, `Claimer`, `Deleter`, `ErrorClassifier`, `Importer`, `KubernetesSecretExporter`, `PostReconciliationChecker`, `PreReconciliationChecker`, `SuccessfulCreationHandler`) are documented in [`docs/hugo/content/contributing/add-a-new-code-generated-resource/implement-extensions.md`](../../../docs/hugo/content/contributing/add-a-new-code-generated-resource/implement-extensions.md).

**Done when:** either no extension is required, or the extension is written and compiles cleanly under `task controller:quick-checks`.

### Step 5: Write a CRUD test

Every new resource needs a hand-written CRUD test — even new versions of existing resources. Copy an existing test from `v2/internal/controllers/` that targets a similar resource. The [README in that folder](../../../v2/internal/controllers/README.md#file-naming) defines the naming convention as `<group>_<subject>_<scenario>_<version>_test.go` (e.g. `batch_account_crud_v1api20210101_test.go`, `servicebus_namespace_standardSkuCrud_v1api20211101_test.go`) — **always follow this pattern for new tests**, even though some existing tests do not.

The test must:

- Create the resource (plus any required dependencies) and wait for `Ready`.
- Verify at least one round-trip property was set as expected.
- Delete the resource and confirm cleanup.
- Use `CreateResourcesAndWait` (plural) to create dependent resources **together**, not one-by-one. Sequential creation hides ordering bugs the operator has to handle in a real cluster.
- Prefer parallel subtests (`RunParallelSubtests`) for complex independent scenarios — the operator is concurrent in production, so tests should be too. **Do not use a subtest for simple CRUD scenario**; it adds unnecessary complexity, slows the test, and forces sequential creation of resources.

**Done when:** the test file compiles (`task controller:quick-checks` passes) and reads like the CRUD tests already in the folder.

### Step 6: Record the CRUD test

**REQUIRED SUB-SKILL:** Use `testing-aso-recordings` to run the test and produce a recording; use `diagnosing-vcr-failures` if the recording fails.

Recording requires the Azure credentials from Prerequisites #6. If missing, stop and ask the user — you cannot record without them.

Run (substituting your test's actual `Test_...` function name):

```bash
source test.env && TIMEOUT=60m TEST_FILTER=<YourTestFunctionName> task controller:test-controllers
```

The `TEST_FILTER` value is the Go test function name you wrote in Step 5, not a fixed pattern — function names in this folder vary widely (`Test_Workspace_BigDataPool`, `Test_ContainerRegistry_Registry_20230701_CRUD`, etc.). Read the top of your test file to confirm the exact name.

Then re-run the same command **without** `source test.env` to verify playback works from the recording.

**Done when:** a new recording file exists under `v2/internal/controllers/recordings/`, and the test passes in playback mode.

### Step 7: Create and record a sample

Every resource must ship with a sample under `v2/samples/<group>/<version>/`. CI enforces this. The sample must exercise a working configuration, including all required dependencies:

- Dependencies **in the same group** go in the same directory.
- Dependencies **from other groups** go under a `refs/` subdirectory. Look at existing samples for the exact pattern.

Once the YAML is in place, record with the samples suite:

```bash
source test.env && TIMEOUT=60m TEST_FILTER="Test_Samples_CreationAndDeletion/Test_<GroupTitleCased>_v<version>_CreationAndDeletion" task controller:test-samples
```

The group segment is title-cased using `cases.Title(language.English)` — only the first letter of each segment changes (`synapse` → `Synapse`, `appconfiguration` → `Appconfiguration`). For dotted groups like `network.frontdoor`, dots are **dropped** and the title-cased segments are concatenated with no separator: `network.frontdoor` → `NetworkFrontdoor`, giving `Test_NetworkFrontdoor_v<version>_CreationAndDeletion`. See `getTestName` in `v2/internal/testsamples/samples_test.go` for the definitive rule.

Verify playback the same way as Step 6. 

**Done when:** the sample YAML(s) exist under `v2/samples/<group>/<version>/`, a recording exists under `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/`, and playback passes.

### Step 8: Final verification

The full CI is too slow to run in one shot. Run each check separately, in order. **If any step fails, fix the issue and restart Step 8 from the top** — a downstream fix can invalidate an upstream check.

1. `task format-code` — format Go and YAML.
2. `task generator:quick-checks` — regenerate code, run generator unit tests.
3. `task controller:quick-checks` — build the controller, run its unit tests.
4. `task asoctl:quick-checks` — build asoctl, run its unit tests.
5. `task doc:crd-api` — regenerate reference docs (if changed, commit the diff).

**Done when:** all five commands exit cleanly, `git status` shows only the changes you intended, and the PR includes all the file categories listed in **Postrequisites** below.

## Troubleshooting

### Generator errors

The generator surfaces one problem at a time. Read the exact error text — it usually tells you which property or type needs configuration.

- **"looks like a resource reference but was not labelled as one"** — the generator's heuristics flagged an ID-shaped property. Decide whether it points at another ARM resource. If yes, add `$referenceType: arm` under that property. If no, add `$referenceType: simple`. The `arm` form is swapped for `genruntime.ResourceReference` (which supports both in-cluster and Azure references); `simple` leaves it as a plain string.

- **"$exportAs specified for type X but not consumed"** — configuration you added had no effect. Usually a typo: the error prints `did you mean ...?` with the nearest match. (Note: `$export: true` is now rejected up front with a different message — `$export is deprecated, use $exportAs instead` — so if you see that, just rename the directive.) If your target is a preview version, you also need a `typeFilters` entry — preview versions are pruned by default. Add an `include` filter **before** the `prune` filter, and give a reason:

  ```yaml
  - action: include
    group: keyvault
    version: v*20210401preview
    because: We want to support keyvault which is only available in preview.
  ```

- **"type X not seen"** — the submodule `v2/specs/azure-rest-api-specs` is likely stale. Run `git submodule update --init --recursive`. If that fixes it, submit the submodule bump as a **separate PR** — submodule updates drag in unrelated doc changes and should not ride along with a new resource.

- **Deeper debugging** — use `--debug <group>` (or `--debug <group>/<version>`, semicolon-separated for multiple) to dump per-stage internal state to a temp directory. The volume is high; use it only when normal errors don't point at the answer. Documented in [`run-the-code-generator.md`](../../../docs/hugo/content/contributing/add-a-new-code-generated-resource/run-the-code-generator.md#debugging).

### Test recording issues

Delegate diagnosis to the `diagnosing-vcr-failures` skill. Common headline problems the skill handles:

- Missing HTTP interaction → recording is stale; delete and re-record.
- Body hash mismatch on a freshly recorded run → non-deterministic serialisation; do **not** re-record in a loop, report to the user.
- Timeout while recording → some resources need 60–90 min; extend with `TIMEOUT=90m`.
- Region capacity / quota errors → pick another region in the sample YAML, delete the partial recording, retry.

**Never edit recording YAML by hand.** They are machine-generated and any manual edit will drift from what the runtime produces.

### Live/replay drift

If a test passes recording but fails in playback (or vice versa), the recording is not the fix — the code is producing different HTTP output across runs. This is a systemic issue. Report to the user with the exact request that diverged; do not keep re-recording.

## Guidance and gotchas

- **Never edit `*_gen.go` or `*_gen_test.go` files by hand.** They are overwritten on the next generator run. Behaviour changes must go through `azure-arm.yaml` (for shape) or `customizations/` (for extensions).
- **Never edit files under `docs/hugo/content/reference/`.** They are generated by `task doc:crd-api`.
- **Never prettify or reformat `azure-arm.yaml` or its imports.** A separate reformat PR is welcome; don't bundle it with a resource.
- **Progression is not linear.** Writing the test frequently sends you back to Step 1 to fix a config problem you couldn't see from the generator output alone. Expect at least one round-trip through Steps 1–5.
- **Split commits sensibly.** Reviewers appreciate: (a) config changes to `azure-arm.yaml` + group import, (b) generated Go + registration changes, (c) generated docs, (d) hand-written test/sample/extension changes. Making each of these a separate commit makes review dramatically easier.

## Postrequisites

Before declaring the resource done, confirm the PR includes **all** of the following. Missing any of them is a rework signal, not a nice-to-have.

- Configuration changes to `v2/azure-arm.yaml` and/or `v2/azure-arm/<group>.yaml`.
- Generated code under `v2/api/<group>/<version>/` (including `_gen.go` files, `zz_generated.deepcopy.go`, and `structure.txt`).
- Registration changes in `v2/internal/controllers/controller_resources_gen.go`.
- Generated documentation under `docs/hugo/content/reference/`.
- Any hand-written extensions under `v2/api/<group>/customizations/` (if Step 4 applied).
- A hand-written CRUD test under `v2/internal/controllers/` plus its recording under `v2/internal/controllers/recordings/`.
- Sample YAML(s) under `v2/samples/<group>/<version>/` plus a recording under `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/`.

Report back to the user with a table listing what was added, which tests were recorded, and any generator config decisions that required judgement (secret detection, read-only properties, extensions).

## References

- Canonical narrative documentation (with screenshots): [`docs/hugo/content/contributing/add-a-new-code-generated-resource/`](../../../docs/hugo/content/contributing/add-a-new-code-generated-resource/_index.md)
- **REQUIRED SUB-SKILL** for running and recording tests: `testing-aso-recordings`
- **REQUIRED SUB-SKILL** for diagnosing recording failures: `diagnosing-vcr-failures`
- Full list of `azure-arm.yaml` configuration directives: see the large comment at the end of `v2/azure-arm.yaml`
