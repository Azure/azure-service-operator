# RegistryCacheRule 2025-11-01 Migration Design

## Goal

Move ASO support for `RegistryCacheRule` exclusively from Azure API version
`2025-04-01` to stable version `2025-11-01`.

## Configuration and Generation

Move the `CacheRuleProperties` reference customization and
`Registries_CacheRule` export configuration in
`v2/azure-arm/containerregistry.yaml` to `2025-11-01`. Run the ASO code
generator and accept the generated replacement of the `v20250401` API,
storage, ARM, webhook, controller registration, and CRD artifacts with
`v20251101` equivalents.

The existing `Registry` dependency remains on `v1api20230701`; this migration
does not add or change Registry API support.

## Tests and Samples

Rename and update the RegistryCacheRule CRUD test to use `v20251101`, while
preserving its current behavior and representative properties. Move the sample
to `v2/samples/containerregistry/v20251101`, update its API version, and retain
the existing `v1api20230701` Registry reference sample.

Delete only the obsolete RegistryCacheRule CRUD recording and containerregistry
`v20250401` sample recording. Re-record the migrated CRUD test and the
containerregistry `v20251101` sample creation/deletion test against Azure.

## Validation

Run formatting, generator quick checks, and focused controller integration
tests for the RegistryCacheRule CRUD test and the `v20251101` containerregistry
sample. Confirm generated and handwritten files no longer reference the
RegistryCacheRule `v20250401` API.
