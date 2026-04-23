# Split Generator Configuration File

**Issue:** [#4850](https://github.com/Azure/azure-service-operator/issues/4850)
**Status:** Design

## Problem

The generator configuration (`azure-arm.yaml`) is a monolithic 4600+ line YAML file. This is intimidating for humans and difficult for AI assistants to work with. Configuration for each Azure service group is scattered across three sections (typeFilters, typeTransformers, objectModelConfiguration).

## Approach

Split the monolithic config into per-group YAML files referenced via an `imports` section in the main config. Each group file contains TypeFilters, TypeTransformers, and GroupModelConfiguration for that group.

## Group Configuration File Format

A new `GroupConfigurationFile` struct represents per-group YAML files, composing existing types:

```go
type GroupConfigurationFile struct {
    TypeFilters             []*TypeFilter      `yaml:"typeFilters,omitempty"`
    Transformers            []*TypeTransformer `yaml:"typeTransformers,omitempty"`
    GroupModelConfiguration *GroupConfiguration `yaml:"groupModelConfiguration,omitempty"`
}
```

This reuses `*GroupConfiguration` directly, which already handles YAML unmarshalling of version entries and group-level properties like `$payloadType`.

Example group file (`azure-arm/containerservice.yaml`):

```yaml
typeFilters:
  - action: include
    version: v*20250301
    because: "Include this specific preview version"

typeTransformers:
  - name: ManagedCluster_Properties_Spec
    property: DnsPrefix
    remove: true
    because: "Read-only property"

groupModelConfiguration:
  $payloadType: explicitCollections
  2024-09-01:
    ManagedCluster:
      $exportAs: ManagedCluster
      $supportedFrom: v2.11.0
```

## Import Mechanism

Add an `Imports` field to `Configuration`:

```go
type Configuration struct {
    // ... existing fields ...
    Imports map[string]string `yaml:"imports,omitempty"` // group name → file path
}
```

Example in `azure-arm.yaml`:

```yaml
imports:
  containerservice: "azure-arm/containerservice.yaml"
  storage: "azure-arm/storage.yaml"
```

## Loading Flow

Within `LoadConfiguration()`, after loading the main config:

1. For each entry in `Imports`:
   - Resolve file path relative to main config file directory
   - Load and parse the `GroupConfigurationFile`
   - Validate: error if any TypeFilter/TypeTransformer has an explicit `group` that doesn't match the import key
   - Auto-fill: set `group` on TypeFilters/TypeTransformers where omitted
   - Merge into main config

## Merge Rules

- **TypeFilters**: Prepend group-specific filters before global filters (global filters like "prune preview" should run after group-specific includes)
- **TypeTransformers**: Append group-specific transformers after global transformers
- **GroupModelConfiguration**: Set the group name on the `GroupConfiguration` and add to `ObjectModelConfiguration.groups` keyed by the import group name. Use `NewGroupConfiguration()` to ensure proper initialization of internal fields (typoAdvisor, configurable fields).

## Validation & Error Cases

- Import file not found → error with clear path info
- Duplicate group (same group in both imports and main `objectModelConfiguration`) → error
- Explicit `group` on filter/transformer doesn't match import key → error
- Omitted `group` on filter/transformer → auto-filled with import key (not an error)
- Unknown YAML fields in group file → error (use `KnownFields(true)` for consistency with main config loading)

## File Organization After Split

- `v2/azure-arm.yaml` retains:
  - All global settings (schemaRoot, pipeline, paths, reports, etc.)
  - TypeFilters that are NOT group-specific (e.g., global `prune *preview`)
  - TypeTransformers that are NOT group-specific (e.g., global Tags transformer)
  - The `imports` section
  - Empty or removed `objectModelConfiguration`

- `v2/azure-arm/<group>.yaml` (one per group) contains:
  - TypeFilters with that group
  - TypeTransformers with that group
  - GroupModelConfiguration (the version→type→property hierarchy)

## Testing

### Unit Tests

1. **Load group file** — verify all fields parsed correctly
2. **Wrong group error** — explicit group mismatch rejected
3. **Omitted group auto-filled** — group set from import key
4. **TypeFilters prepended** — verify ordering
5. **Transformers appended** — verify ordering
6. **GroupModel merged** — verify integration into ObjectModelConfiguration
7. **Duplicate group error** — same group in imports and main config rejected
8. **End-to-end** — load main config with imports, verify complete config

### Integration Validation

Run `task controller:generate-types` to verify the split config produces identical generated code.

## Scope

### In Scope

- New `GroupConfigurationFile` type in config package
- Loading and merging logic in `LoadConfiguration()`
- Validation logic
- Unit tests
- Splitting all groups from `azure-arm.yaml` into per-group files
- Updating `azure-arm.yaml` with `imports` section

### Out of Scope

- Changing the code generator pipeline
- Modifying generated output
- Splitting by any axis other than group
