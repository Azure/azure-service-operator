# Cassandra Service Principal Sample Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make the Cassandra `v20260315` sample discover the tenant-local Cassandra service principal object ID through ASO instead of requiring a manually copied ID.

**Architecture:** Add an `AdoptOnly` Entra `ServicePrincipal` as a root-level sample resource and export its `status.entraID` to a ConfigMap. Both subnet role assignments consume that ConfigMap value through `principalIdFromConfig`; normal reconciliation handles the period between Cassandra registration and successful adoption.

**Tech Stack:** Kubernetes YAML, Azure Service Operator Entra and Authorization CRDs, Go envtest sample runner, go-vcr recordings.

## Global Constraints

- The `ServicePrincipal` manifest must be directly under `v2/samples/documentdb/cassandra/v20260315/`, not under `refs/`.
- Use application ID `a232010e-820c-4083-83bb-3ace5fc29d0b`.
- Use `operatorSpec.creationMode: AdoptOnly`.
- Export `status.entraID` to ConfigMap `cassandra-service-principal`, key `objectId`.
- Keep both role assignments' owners, `principalType`, role definition, and detach-on-delete policy unchanged.
- Never edit a recording by hand; delete it and regenerate it through the test.
- Do not interrupt a live recording unless an authentication, policy, or other infrastructure failure makes progress impossible.

---

### Task 1: Replace the manual Cassandra principal ID

**Files:**
- Create: `v2/samples/documentdb/cassandra/v20260315/v1_serviceprincipal.yaml`
- Modify: `v2/samples/documentdb/cassandra/v20260315/refs/v20220401_roleassignment_cassandra-dc-subnet.yaml`
- Modify: `v2/samples/documentdb/cassandra/v20260315/refs/v20220401_roleassignment_cassandra-mgmt-subnet.yaml`
- Modify: `v2/samples/documentdb/cassandra/v20260315/README.md`

**Interfaces:**
- Produces: ConfigMap `cassandra-service-principal` with key `objectId`.
- Consumes: The tenant service principal registered for application ID `a232010e-820c-4083-83bb-3ace5fc29d0b`.
- Supplies: `spec.principalIdFromConfig` for both Cassandra subnet `RoleAssignment` resources.

- [ ] **Step 1: Confirm the current manual dependency**

Run:

```bash
grep -R -n \
  -e 'principalId: e5007d2c-4b13-4a74-9b6a-605d99f03501' \
  -e 'az ad sp show --id a232010e-820c-4083-83bb-3ace5fc29d0b' \
  v2/samples/documentdb/cassandra/v20260315
```

Expected: both role assignments contain the literal principal ID, and the
README documents the manual lookup.

- [ ] **Step 2: Add the root-level ServicePrincipal sample**

Create `v2/samples/documentdb/cassandra/v20260315/v1_serviceprincipal.yaml`
with exactly:

```yaml
# Azure registers this service principal when Cassandra cluster provisioning begins.
# Adopt it and export its tenant-specific object ID for the subnet role assignments.
apiVersion: entra.azure.com/v1
kind: ServicePrincipal
metadata:
  name: cassandra-service-principal
  namespace: default
spec:
  appId: a232010e-820c-4083-83bb-3ace5fc29d0b
  operatorSpec:
    creationMode: AdoptOnly
    configmaps:
      entraID:
        name: cassandra-service-principal
        key: objectId
```

- [ ] **Step 3: Make the data-centre subnet role assignment consume the ConfigMap**

In
`v2/samples/documentdb/cassandra/v20260315/refs/v20220401_roleassignment_cassandra-dc-subnet.yaml`,
replace the tenant-specific comments and `principalId` field with:

```yaml
  principalIdFromConfig:
    name: cassandra-service-principal
    key: objectId
```

Leave this existing field unchanged:

```yaml
  principalType: ServicePrincipal
```

- [ ] **Step 4: Make the management subnet role assignment consume the ConfigMap**

In
`v2/samples/documentdb/cassandra/v20260315/refs/v20220401_roleassignment_cassandra-mgmt-subnet.yaml`,
replace the tenant-specific comments and `principalId` field with:

```yaml
  principalIdFromConfig:
    name: cassandra-service-principal
    key: objectId
```

Leave this existing field unchanged:

```yaml
  principalType: ServicePrincipal
```

- [ ] **Step 5: Replace the README's manual setup instructions**

Replace the manual lookup section with:

```markdown
# Tips

For your Cassandra resources to deploy correctly, the Azure Cosmos DB service
principal needs `Microsoft.Network/virtualNetworks/subnets/join/action` on both
subnets.

See https://learn.microsoft.com/en-us/azure/managed-instance-apache-cassandra/add-service-principal

Creating the Cassandra cluster registers the service principal in the tenant.
The `ServicePrincipal` sample adopts it and exports its tenant-specific object
ID to the `cassandra-service-principal` ConfigMap. Both `RoleAssignment`
resources read that value from the ConfigMap and grant Network Contributor on
their respective subnets.
```

- [ ] **Step 6: Check the manifest wiring**

Run:

```bash
grep -R -n \
  -e 'e5007d2c-4b13-4a74-9b6a-605d99f03501' \
  -e 'az ad sp show --id' \
  v2/samples/documentdb/cassandra/v20260315
```

Expected: no output.

Run:

```bash
grep -R -n \
  -e 'appId: a232010e-820c-4083-83bb-3ace5fc29d0b' \
  -e 'principalIdFromConfig:' \
  -e 'name: cassandra-service-principal' \
  -e 'key: objectId' \
  v2/samples/documentdb/cassandra/v20260315
```

Expected: one application ID, two `principalIdFromConfig` fields, and matching
ConfigMap producer and consumer names.

- [ ] **Step 7: Review and commit the sample change**

Run:

```bash
git diff --check
git diff -- \
  v2/samples/documentdb/cassandra/v20260315
```

Expected: no whitespace errors; the diff contains only the new
`ServicePrincipal`, the two ConfigMap-backed role assignments, and the README
update.

Commit:

```bash
git add \
  v2/samples/documentdb/cassandra/v20260315/v1_serviceprincipal.yaml \
  v2/samples/documentdb/cassandra/v20260315/refs/v20220401_roleassignment_cassandra-dc-subnet.yaml \
  v2/samples/documentdb/cassandra/v20260315/refs/v20220401_roleassignment_cassandra-mgmt-subnet.yaml \
  v2/samples/documentdb/cassandra/v20260315/README.md
git commit -m "Update Cassandra sample to discover service principal" \
  -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

### Task 2: Re-record and verify the Cassandra sample

**Files:**
- Regenerate: `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion.yaml`
- Inspect: `reports/test-samples.log`

**Interfaces:**
- Consumes: the sample manifests from Task 1 and credentials from `test.env`.
- Produces: a deterministic go-vcr cassette that passes both live recording and playback.

- [ ] **Step 1: Run recording preflight checks**

Run:

```bash
set -a
source test.env
set +a
export PATH="$PWD/hack/tools:$PATH"
test -n "$AZURE_SUBSCRIPTION_ID"
test -n "$AZURE_TENANT_ID"
test -n "$ENTRA_APP_ID"
test "$(az account show --query tenantId -o tsv)" = "$AZURE_TENANT_ID"
az account get-access-token --resource-type ms-graph --query expiresOn -o tsv >/dev/null
test "$(az ad app list --filter "appId eq '$ENTRA_APP_ID'" --query 'length(@)' -o tsv)" = "1"
command -v task
```

Expected: every command exits zero, `test.env` produces no authentication
warnings, and `task` resolves from `PATH`.

- [ ] **Step 2: Delete only the stale Cassandra cassette**

Delete:

```text
v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion.yaml
```

Confirm:

```bash
test ! -e v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion.yaml
```

Expected: exit code 0.

- [ ] **Step 3: Start the targeted live recording**

Run asynchronously:

```bash
set -a
source test.env
set +a
export PATH="$PWD/hack/tools:$PATH"
TIMEOUT=90m \
TEST_FILTER='Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion' \
task controller:test-samples
```

Expected: the command remains active while Azure provisions the cluster and
data centre.

- [ ] **Step 4: Monitor without interrupting normal provisioning**

At intervals of at least five minutes, run:

```bash
sleep 300
tail -30 reports/test-samples.log
grep -Ei \
  'FAIL:|AADSTS|policy|forbidden|authorization|denied|insufficient|permission|RequestDenied|Authorization_RequestDenied|token protection' \
  reports/test-samples.log | tail -40 || true
```

Expected: no authentication or policy failures. Messages that the adopted
service principal or ConfigMap is not yet available are transient only while
Azure registration completes.

If an authentication, policy, quota, or other infrastructure error appears,
stop the test promptly and diagnose that error before retrying. Otherwise,
allow the command to finish even if provisioning takes more than an hour.

- [ ] **Step 5: Verify recording success**

After the command exits, run:

```bash
grep 'FAIL:' reports/test-samples.log || echo 'No failures found'
test -s v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion.yaml
```

Expected:

```text
No failures found
```

The cassette exists and is non-empty.

- [ ] **Step 6: Verify playback**

Run without sourcing `test.env`:

```bash
export PATH="$PWD/hack/tools:$PATH"
TIMEOUT=90m \
TEST_FILTER='Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion' \
task controller:test-samples
```

After completion, run:

```bash
grep 'FAIL:' reports/test-samples.log || echo 'No failures found'
```

Expected:

```text
No failures found
```

- [ ] **Step 7: Check redaction and commit the cassette**

Load the current test values and verify they are absent without printing them:

```bash
set -a
source test.env
set +a
recording='v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion.yaml'
! grep -Fq "$AZURE_SUBSCRIPTION_ID" "$recording"
! grep -Fq "$AZURE_TENANT_ID" "$recording"
! grep -Fq "$ENTRA_APP_ID" "$recording"
git diff --check -- "$recording"
git status --short -- "$recording"
```

Expected: all raw tenant values are absent, no whitespace errors are reported,
and the cassette is shown as modified.

Commit:

```bash
git add v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion.yaml
git commit -m "Record Cassandra service principal sample" \
  -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```
