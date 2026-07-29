# Container Instance Hybrid Versioning Design

## Scope

Migrate the `containerinstance` resource group from legacy to hybrid versioning in
the next unreleased ASO release, `v2.21.0`.

## Design

The generator will emit the new `containerinstance.azure.com/v20211001` API while
retaining `containerinstance.azure.com/v1api20211001` for backward compatibility.
The migration release is registered so generated documentation reports the new API
as available in `v2.21.0` and does not deprecate the legacy API before that release.

The existing ContainerGroup sample is copied unchanged except for its filename,
directory, and `apiVersion`. The existing CRUD test already has a version-neutral
test name, so its import changes to the new API without renaming its recording.
Sample recordings are created for the new sample version.

## Validation

Generate types, record the targeted Container Instance samples using configured
Azure credentials, then run formatting, generator checks, controller checks, and
CRD API documentation generation.
