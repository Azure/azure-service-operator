# Cassandra Service Principal Sample Design

## Why this matters

The `documentdb/cassandra/v20260315` sample currently embeds a tenant-specific
object ID in two role assignments. Anyone running the sample in another tenant
must look up the Azure Cosmos DB service principal manually and edit both
manifests before deployment.

The sample should discover that object ID through ASO after Azure registers the
service principal for the Cassandra cluster.

## Design

Add a `entra.azure.com/v1` `ServicePrincipal` manifest directly to
`v2/samples/documentdb/cassandra/v20260315/`. It must not be placed under
`refs/`.

The resource will:

- use the globally known Cassandra application ID
  `a232010e-820c-4083-83bb-3ace5fc29d0b`;
- use `operatorSpec.creationMode: AdoptOnly`;
- export `status.entraID` to the `objectId` key in a ConfigMap named
  `cassandra-service-principal`.

Change both subnet `RoleAssignment` manifests to replace the literal
`principalId` with:

```yaml
principalIdFromConfig:
  name: cassandra-service-principal
  key: objectId
```

Keep `principalType: ServicePrincipal`, the existing owners, role definition,
and detach-on-delete policy unchanged.

## Reconciliation flow

The sample test submits refs and root-level samples together. Their ordering is
not deterministic, so the design must rely on normal controller reconciliation
rather than file ordering.

1. The Cassandra cluster begins provisioning.
2. Azure registers the Cassandra resource-provider service principal in the
   tenant.
3. The ASO `ServicePrincipal` resource retries adoption until the principal is
   available.
4. ASO writes the tenant-specific object ID to the ConfigMap.
5. Each role assignment resolves `principalIdFromConfig` and grants Network
   Contributor on its subnet.
6. Cassandra provisioning proceeds once the required subnet permissions exist.

Before the ConfigMap exists, the role assignments remain unready and retry.
Deleting the sample removes the Kubernetes `ServicePrincipal` object but, due
to `AdoptOnly`, does not delete Azure's tenant-owned service principal.

## Documentation

Update the sample README to remove the manual `az ad sp show` step. Explain
that deployment registers the service principal, ASO adopts it, and the
exported object ID supplies both role assignments.

## Validation

1. Run formatting and the smallest relevant repository checks for the YAML and
   documentation changes.
2. Remove the existing
   `Test_Cassandra_v20260315_CreationAndDeletion.yaml` cassette.
3. Record
   `Test_Samples_CreationAndDeletion/Test_Cassandra_v20260315_CreationAndDeletion`
   with `TIMEOUT=90m`, monitoring for authentication, policy, and authorization
   failures.
4. Run the same targeted test without credentials to verify playback.
5. Confirm the cassette was recreated and contains no unredacted
   tenant-specific identifiers.
