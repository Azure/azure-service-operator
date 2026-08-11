# Container Registry Cache Rule Recordings Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make the new Container Registry cache rule controller test and sample valid against the current Azure API, then create playback-safe VCR recordings for both.

**Architecture:** Keep the cache rules unauthenticated, but replace the now-rejected Docker Hub source with the existing public MCR fixture `mcr.microsoft.com/azuredocs/aci-helloworld`. Use identical source and target repository values in the Go CRUD test and YAML sample, then validate each cassette first by live recording and then by credential-free playback.

**Tech Stack:** Go, Kubernetes YAML, Azure Service Operator controller tests, Task, go-vcr

## Global Constraints

- Keep the cache rule unauthenticated.
- Use the same source and target repository values in both fixtures.
- Do not change generated APIs or controller behavior.
- Never edit recording YAML by hand.
- Do not interrupt live recording runs.

---

### Task 1: Update the Cache Rule Test Fixtures

**Files:**
- Modify: `v2/internal/controllers/crd_containerregistry_cache_rule_20250401_test.go:44-45`
- Modify: `v2/samples/containerregistry/v20250401/v20250401_registrycacherule.yaml:9-10`
- Create: `v2/samples/containerregistry/v20250401/refs/v1api20230701_registry.yaml`

**Interfaces:**
- Consumes: `RegistryCacheRule_Spec.SourceRepository` and `RegistryCacheRule_Spec.TargetRepository`
- Produces: Matching controller and sample fixtures accepted by Azure without credentials

- [ ] **Step 1: Preserve the observed failing baseline**

Confirm `reports/test-controllers.log` contains the live Azure failure:

```bash
grep -n "EnforceCacheRuleAuthentication" reports/test-controllers.log
```

Expected: at least one line stating that Docker cache rules require credentials.

- [ ] **Step 2: Update the controller test**

Replace the repository fields in `Test_ContainerRegistry_CacheRule_20250401_CRUD` with:

```go
SourceRepository: to.Ptr("mcr.microsoft.com/azuredocs/aci-helloworld"),
TargetRepository: to.Ptr("cached-mcr/aci-helloworld"),
```

- [ ] **Step 3: Update the sample**

Replace the repository fields in `v20250401_registrycacherule.yaml` with:

```yaml
  sourceRepository: mcr.microsoft.com/azuredocs/aci-helloworld
  targetRepository: cached-mcr/aci-helloworld
```

- [ ] **Step 4: Add the sample registry dependency**

Create `v2/samples/containerregistry/v20250401/refs/v1api20230701_registry.yaml`:

```yaml
apiVersion: containerregistry.azure.com/v1api20230701
kind: Registry
metadata:
  name: sampleregistry
  namespace: default
spec:
  location: westcentralus
  owner:
    name: aso-sample-rg
  publicNetworkAccess: Enabled
  sku:
    name: Premium
  zoneRedundancy: Disabled
```

- [ ] **Step 5: Check formatting and fixture consistency**

Run:

```bash
task format-code
git diff --check
grep -R "mcr.microsoft.com/azuredocs/aci-helloworld\|cached-mcr/aci-helloworld" \
  v2/internal/controllers/crd_containerregistry_cache_rule_20250401_test.go \
  v2/samples/containerregistry/v20250401/v20250401_registrycacherule.yaml
grep -R "name: sampleregistry\|name: Premium" \
  v2/samples/containerregistry/v20250401
```

Expected: `git diff --check` exits successfully, each repository value appears
in both fixtures, and the sample version contains its Premium registry owner.

- [ ] **Step 6: Commit the fixture correction**

```bash
git add \
  v2/internal/controllers/crd_containerregistry_cache_rule_20250401_test.go \
  v2/samples/containerregistry/v20250401/v20250401_registrycacherule.yaml \
  v2/samples/containerregistry/v20250401/refs/v1api20230701_registry.yaml
git commit -m "test: use public MCR source for cache rules" \
  -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

### Task 2: Record and Replay the Controller CRUD Test

**Files:**
- Create: `v2/internal/controllers/recordings/Test_ContainerRegistry_CacheRule_20250401_CRUD.yaml`
- Inspect: `reports/test-controllers.log`

**Interfaces:**
- Consumes: `Test_ContainerRegistry_CacheRule_20250401_CRUD`
- Produces: A deterministic controller VCR cassette

- [ ] **Step 1: Run the live controller test**

Run asynchronously and allow it to finish:

```bash
source test.env &&
TIMEOUT=60m \
TEST_FILTER="Test_ContainerRegistry_CacheRule_20250401_CRUD" \
task controller:test-controllers
```

Expected: command exits successfully and creates
`v2/internal/controllers/recordings/Test_ContainerRegistry_CacheRule_20250401_CRUD.yaml`.

- [ ] **Step 2: Confirm the recording run succeeded**

```bash
grep "FAIL:" reports/test-controllers.log || echo "No failures found"
test -f v2/internal/controllers/recordings/Test_ContainerRegistry_CacheRule_20250401_CRUD.yaml
```

Expected: `No failures found` and `test -f` exits successfully.

- [ ] **Step 3: Replay without Azure credentials**

```bash
TIMEOUT=60m \
TEST_FILTER="Test_ContainerRegistry_CacheRule_20250401_CRUD" \
task controller:test-controllers
```

Expected: command exits successfully using the new cassette.

- [ ] **Step 4: Confirm playback succeeded**

```bash
grep "FAIL:" reports/test-controllers.log || echo "No failures found"
```

Expected: `No failures found`.

### Task 3: Record and Replay the Sample Test

**Files:**
- Create: `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion.yaml`
- Inspect: `reports/test-samples.log`

**Interfaces:**
- Consumes: `v2/samples/containerregistry/v20250401/v20250401_registrycacherule.yaml`
- Produces: A deterministic sample VCR cassette for `Test_Containerregistry_v20250401_CreationAndDeletion`

- [ ] **Step 1: Run the live sample subtest**

Run asynchronously and allow it to finish:

```bash
source test.env &&
TIMEOUT=60m \
TEST_FILTER="Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion" \
task controller:test-samples
```

Expected: command exits successfully and creates
`v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion.yaml`.

- [ ] **Step 2: Confirm the recording run succeeded**

```bash
grep "FAIL:" reports/test-samples.log || echo "No failures found"
test -f v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion.yaml
```

Expected: `No failures found` and `test -f` exits successfully.

- [ ] **Step 3: Replay without Azure credentials**

```bash
TIMEOUT=60m \
TEST_FILTER="Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion" \
task controller:test-samples
```

Expected: command exits successfully using the new cassette.

- [ ] **Step 4: Confirm playback succeeded**

```bash
grep "FAIL:" reports/test-samples.log || echo "No failures found"
```

Expected: `No failures found`.

### Task 4: Verify and Commit the Recordings

**Files:**
- Inspect: `v2/internal/controllers/recordings/Test_ContainerRegistry_CacheRule_20250401_CRUD.yaml`
- Inspect: `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion.yaml`

**Interfaces:**
- Consumes: Both newly generated cassettes and playback results
- Produces: Persistent recordings ready for review

- [ ] **Step 1: Verify both generated files and repository hygiene**

```bash
test -s v2/internal/controllers/recordings/Test_ContainerRegistry_CacheRule_20250401_CRUD.yaml
test -s v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion.yaml
git diff --check
git status --short
```

Expected: both `test -s` commands and `git diff --check` succeed; `git status` shows the two new recordings plus any previously existing branch changes.

- [ ] **Step 2: Commit only the recordings**

```bash
git add \
  v2/internal/controllers/recordings/Test_ContainerRegistry_CacheRule_20250401_CRUD.yaml \
  v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Containerregistry_v20250401_CreationAndDeletion.yaml
git commit -m "test: record container registry cache rules" \
  -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```
