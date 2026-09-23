# HCP OpenShift integration test

`Test_RedHatOpenShift_HcpOpenShiftCluster_v20260901preview_CRUD` uses ASO's
controller envtest and Azure record/replay framework. It submits a cluster,
node pool, and external-auth resource together, verifies their status, updates
each resource, deletes the external-auth and cluster, and checks that all
three resources are absent in Azure.

The external-auth issuer URL and client ID need not identify a real Entra ID
app registration. ARM's own "maximum fuzz" example for this resource (see
`specs/azure-rest-api-specs/.../ExternalAuths_CreateOrUpdate_MaximumSet_Gen.json`
for the `2026-09-01-preview` version) uses arbitrary non-GUID strings for
`clientId` and an unrelated `https` URL for the issuer, and the generated
schema enforces no stricter pattern. This was confirmed live: the recording
below uses `https://login.microsoftonline.com/<tenant>/v2.0` (always a live
endpoint for any tenant) as the issuer and the plain string
`aso-hcp-test-client` as the client ID, and ARM accepted both without
complaint. Treat both as opaque configuration handed to the cluster's own
OIDC setup, not something ARM validates against Entra/Graph at creation time.

Deletion issues `DeleteResourceAndWait` on the external-auth and the cluster,
but not the node pool: Azure rejects an explicit DELETE of a cluster's last
remaining node pool while the cluster still exists (`Forbidden: The last node
pool can not be deleted from a cluster.`), so ARM's cascade from deleting the
cluster is relied on instead, with all three ARM resources confirmed gone
afterward.

The fixtures follow the CAPZ `mv1-tests-stage/aro.yaml` scenario: OpenShift
4.20 / node pool 4.20.17, the experimental SingleReplica/Minimal tags, and Entra
external authentication with the `oid` username claim. The test creates its own
resource group, VNet, NSG, worker subnet, delegated integration subnet, 13 managed
identities, Key Vault, and 32 role assignments. It exports identity principal IDs
to ConfigMaps and uses them in role assignments, just like the CAPZ manifests.

The KMS key is created through the repository's existing recorded ARM helper
because ASO does not yet expose a Key Vault key resource. All generated names use
the standard deterministic test namer. Existing CAPZ resources are never adopted.
The cluster and node pool use Kubernetes references to the new prerequisites,
including both maps of operator identities.

## Recording

Use an Azure subscription with access to the `2026-09-01-preview` HCP API, the
required HCP roles, and permission to create the resources and role assignments
above. No Entra ID app registration is needed for the external-auth issuer or
client ID (see above); the checked-in example values work for recording as-is.

Copy `config.example.json` outside the repository and fill in its five values:
location, cluster version, node-pool version, issuer URL, and client ID. The
first three must describe a supported HCP deployment; the issuer URL should
point at your tenant's `v2.0` endpoint (`https://login.microsoftonline.com/<tenant>/v2.0`)
and the client ID can be any non-empty string. Do not put tokens or client
secrets in this file. Set `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and
configure Azure authentication as for other ASO integration tests.

### Global test credential

This test uses `globalTestContext.ForTest(t)` and ASO's shared test credential
provider. The envtest controller receives that credential as its global default;
the HCP resources do not configure namespace or per-resource credentials.
The test does not read CAPZ `credentials.yaml` or create an `aso-credential` Secret.

Configure the shared test identity in the process environment before running
the task. For service-principal authentication, this includes `AZURE_CLIENT_ID`
and `AZURE_CLIENT_SECRET` (or the certificate settings supported by Azure's
environment credential), alongside subscription and tenant IDs. Keep those
values in the existing local credential configuration or protected CI secrets.
For example, load an existing shell-compatible environment file locally:

```bash
set -a
source /absolute/path/to/shared-aso-test.env
set +a
```

ASO's test credential chain tries Azure CLI authentication first, then environment
authentication. If using Azure CLI, its active login must be the intended shared
test identity. Supplying service-principal environment variables does not override
a successful CLI login.

The `clientID` in `config.example.json` is the external-auth client ID; it is
an arbitrary string, not a real Entra ID application, and is separate from the
global Azure credential used to provision test resources. Playback uses
recorded responses and requires neither credential.

From the repository root:

```bash
ARO_HCP_TEST_CONFIG=/absolute/path/hcp-test-config.json \
  TEST_FILTER='^Test_RedHatOpenShift_HcpOpenShiftCluster_v20260901preview_CRUD$' \
  TIMEOUT=120m task controller:test-controllers
```

This creates live Azure resources, scales the node pool from two to three
workers, and updates the external-auth claim mapping prefix. Allow cleanup to
finish; a full run (create, patch, delete, resource-group cleanup) takes
roughly an hour. The existing per-operation recording timeout still applies;
the command above extends the overall Go test timeout.

The recorder replaces the five configuration values with the checked-in example values,
in addition to standard ASO redactions. Inspect the resulting cassette for
environment-specific information before committing it:

`v2/internal/controllers/recordings/Test_RedHatOpenShift_HcpOpenShiftCluster_v20260901preview_CRUD.yaml`

A cassette is checked in, so CI replays this test automatically. Without a
cassette or `ARO_HCP_TEST_CONFIG`, the test would explicitly skip instead of
silently passing; a skip is not evidence that the HCP lifecycle works. To
re-record, move the existing cassette aside before running the command above.

## Playback

With the cassette present, no Azure credentials or configuration file are needed:

```bash
TEST_FILTER='^Test_RedHatOpenShift_HcpOpenShiftCluster_v20260901preview_CRUD$' \
  task controller:test-controllers
```

Playback always uses `config.example.json`, even if `ARO_HCP_TEST_CONFIG` is set.
Keep the example values and resource fixture synchronized with the recording.

## Adding tests for other resources

ASO normally adds a Go CRUD test under `v2/internal/controllers`, sample manifests
under `v2/samples/<group>/<version>`, and recorded Azure interactions for both.
Examples in this repository's history are [Cassandra support, PR #5175](https://github.com/Azure/azure-service-operator/pull/5175)
and [Communication Services support, PR #5263](https://github.com/Azure/azure-service-operator/pull/5263).

The sample suite discovers version directories automatically; it does not need
a new Go test for each resource. Dependencies from other resource groups belong
in the sample directory's `refs/` folder. Run a specific directory with:

```bash
TEST_FILTER='Test_Samples_CreationAndDeletion/Test_<Group>_<version>_CreationAndDeletion' \
  task controller:test-samples
```

The HCP sample manifests under `v2/samples/redhatopenshift/v20260901preview` use
generic placeholder values (subscription `00000000-0000-0000-0000-000000000000`,
`aso-sample-*` names), matching the rest of the sample suite, but they still
reference external prerequisites (VNet, NSG, managed identities, Key Vault) that
don't exist within the sample directory itself. `v2/internal/testcommon/samples_tester.go`
therefore excludes `/redhatopenshift/v20260901preview/` specifically (the older
`/redhatopenshift/v1api/` directory remains separately excluded because it requires
service-principal creation).
This dedicated controller test supplies lifecycle coverage once recorded;
automatic sample coverage still requires adding self-contained prerequisites
under the sample directory's `refs/` folder, removing the `v20260901preview`
exclusion above, and recording `Test_RedHatOpenShift_v20260901preview_CreationAndDeletion`.
