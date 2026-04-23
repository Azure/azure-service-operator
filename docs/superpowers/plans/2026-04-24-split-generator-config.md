# Split Generator Config Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Split the monolithic `azure-arm.yaml` into per-group config files loaded via an `imports` section, so each Azure service group has its own focused configuration file.

**Architecture:** New `GroupConfigurationFile` type in the config package holds per-group TypeFilters, TypeTransformers, and GroupModelConfiguration. `LoadConfiguration()` is extended to process an `imports` map, loading each group file, validating group consistency, and merging into the main configuration. A Go splitting tool extracts group-specific config from the existing YAML.

**Tech Stack:** Go, gopkg.in/yaml.v3, existing config package types

---

### Task 1: Create GroupConfigurationFile type with loading and validation

**Files:**
- Create: `v2/tools/generator/internal/config/group_configuration_file.go`
- Create: `v2/tools/generator/internal/config/group_configuration_file_test.go`
- Create: `v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/` (test fixtures)

This task creates the new `GroupConfigurationFile` struct and its loading function. The struct composes existing types.

**Key design decision:** `GroupModelConfiguration` is stored as a raw `*yaml.Node` during initial YAML parsing, then decoded into a properly initialized `*GroupConfiguration` (via `NewGroupConfiguration(name)`) when the group name is known. This mirrors the pattern in `ObjectModelConfiguration.UnmarshalYAML` (object_model_configuration.go:295-331) where `NewGroupConfiguration(lastID)` is called before `n.Decode(gc)`.

- [ ] **Step 1: Write test fixtures**

Create test YAML files in `v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/`:

File `WellFormed.yaml`:
```yaml
typeFilters:
  - action: include
    version: v*20210601
    because: "include this specific version"

typeTransformers:
  - name: SomeType_Spec
    property: ReadOnlyProp
    remove: true
    because: "read-only property"

groupModelConfiguration:
  $payloadType: explicitCollections
  2021-06-01:
    SomeResource:
      $exportAs: SomeResource
      $supportedFrom: v2.0.0
```

File `WellFormed_MinimalNoGroupModel.yaml`:
```yaml
typeFilters:
  - action: include
    version: v*20210601
    because: "include this specific version"
```

File `WellFormed_GroupModelOnly.yaml`:
```yaml
groupModelConfiguration:
  2021-06-01:
    SomeResource:
      $exportAs: SomeResource
      $supportedFrom: v2.0.0
```

- [ ] **Step 2: Write failing tests for loading**

In `group_configuration_file_test.go`, write tests following the existing pattern (see `configuration_private_test.go` for style — uses `t.Parallel()`, `NewGomegaWithT(t)`, subtests):

```go
package config

import (
    "testing"
    "github.com/onsi/gomega"
)

func TestGroupConfigurationFile_WhenYAMLWellFormed_LoadsCorrectly(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    gcf, err := loadGroupConfigurationFile(
        "testdata/TestGroupConfigurationFile/WellFormed.yaml",
        "testgroup",
    )
    g.Expect(err).ToNot(gomega.HaveOccurred())
    g.Expect(gcf).ToNot(gomega.BeNil())
    g.Expect(gcf.TypeFilters).To(gomega.HaveLen(1))
    g.Expect(gcf.Transformers).To(gomega.HaveLen(1))
    g.Expect(gcf.GroupModelConfiguration).ToNot(gomega.BeNil())
    g.Expect(gcf.GroupModelConfiguration.name).To(gomega.Equal("testgroup"))
}

func TestGroupConfigurationFile_WhenMinimal_LoadsCorrectly(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    gcf, err := loadGroupConfigurationFile(
        "testdata/TestGroupConfigurationFile/WellFormed_MinimalNoGroupModel.yaml",
        "testgroup",
    )
    g.Expect(err).ToNot(gomega.HaveOccurred())
    g.Expect(gcf.TypeFilters).To(gomega.HaveLen(1))
    g.Expect(gcf.Transformers).To(gomega.BeEmpty())
    g.Expect(gcf.GroupModelConfiguration).To(gomega.BeNil())
}

func TestGroupConfigurationFile_WhenGroupModelOnly_LoadsCorrectly(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    gcf, err := loadGroupConfigurationFile(
        "testdata/TestGroupConfigurationFile/WellFormed_GroupModelOnly.yaml",
        "testgroup",
    )
    g.Expect(err).ToNot(gomega.HaveOccurred())
    g.Expect(gcf.TypeFilters).To(gomega.BeEmpty())
    g.Expect(gcf.GroupModelConfiguration).ToNot(gomega.BeNil())
}
```

- [ ] **Step 3: Run tests to verify they fail**

Run: `cd v2 && go test ./tools/generator/internal/config/ -run TestGroupConfigurationFile -v`
Expected: FAIL (loadGroupConfigurationFile doesn't exist)

- [ ] **Step 4: Implement GroupConfigurationFile type and loading**

In `group_configuration_file.go`:

```go
package config

import (
    "os"

    "github.com/rotisserie/eris"
    "gopkg.in/yaml.v3"
)

// GroupConfigurationFile represents a per-group configuration file containing
// TypeFilters, TypeTransformers, and GroupModelConfiguration for a single group.
type GroupConfigurationFile struct {
    // Filters used to control which types from this group are included
    TypeFilters []*TypeFilter `yaml:"typeFilters,omitempty"`
    // Transformers used to remap types in this group
    Transformers []*TypeTransformer `yaml:"typeTransformers,omitempty"`
    // GroupModelConfiguration contains version/type/property configuration for this group.
    // This field is not directly YAML-decoded; it's handled via custom UnmarshalYAML
    // so that we can properly initialize GroupConfiguration with the group name.
    GroupModelConfiguration *GroupConfiguration `yaml:"-"`
    // groupModelNode stores the raw YAML for deferred decoding of GroupModelConfiguration
    groupModelNode *yaml.Node
}

// UnmarshalYAML implements custom YAML unmarshalling for GroupConfigurationFile.
// We need custom handling to capture the groupModelConfiguration node for deferred decoding.
func (gcf *GroupConfigurationFile) UnmarshalYAML(value *yaml.Node) error {
    if value.Kind != yaml.MappingNode {
        return eris.Errorf("expected mapping node for group configuration file, but found %s", value.Tag)
    }

    var lastKey string
    for i, n := range value.Content {
        if i%2 == 0 {
            lastKey = n.Value
            continue
        }

        switch lastKey {
        case "typeFilters":
            if err := n.Decode(&gcf.TypeFilters); err != nil {
                return eris.Wrap(err, "decoding typeFilters")
            }
        case "typeTransformers":
            if err := n.Decode(&gcf.Transformers); err != nil {
                return eris.Wrap(err, "decoding typeTransformers")
            }
        case "groupModelConfiguration":
            gcf.groupModelNode = n
        default:
            return eris.Errorf("unexpected field %q in group configuration file", lastKey)
        }
    }

    return nil
}

// loadGroupConfigurationFile loads a per-group configuration file from the given path.
// The groupName parameter is used to initialize the GroupConfiguration with the correct name.
func loadGroupConfigurationFile(path string, groupName string) (*GroupConfigurationFile, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, eris.Wrapf(err, "opening group configuration file %q", path)
    }
    defer f.Close()

    result := &GroupConfigurationFile{}
    decoder := yaml.NewDecoder(f)

    err = decoder.Decode(result)
    if err != nil {
        return nil, eris.Wrapf(err, "group configuration file %q is not valid YAML", path)
    }

    // Decode the groupModelConfiguration with a properly initialized GroupConfiguration
    if result.groupModelNode != nil {
        gc := NewGroupConfiguration(groupName)
        if err := result.groupModelNode.Decode(gc); err != nil {
            return nil, eris.Wrapf(err, "decoding groupModelConfiguration in %q", path)
        }
        result.GroupModelConfiguration = gc
    }

    return result, nil
}
```

- [ ] **Step 5: Run tests to verify they pass**

Run: `cd v2 && go test ./tools/generator/internal/config/ -run TestGroupConfigurationFile -v`
Expected: All 3 tests PASS

- [ ] **Step 6: Commit**

```bash
git add v2/tools/generator/internal/config/group_configuration_file.go \
        v2/tools/generator/internal/config/group_configuration_file_test.go \
        v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/
git commit -m "Add GroupConfigurationFile type with loading and tests

Introduces the per-group config file format for #4850.
Composes existing TypeFilter, TypeTransformer, and GroupConfiguration types.

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 2: Add validation logic (group consistency checks)

**Files:**
- Modify: `v2/tools/generator/internal/config/group_configuration_file.go`
- Modify: `v2/tools/generator/internal/config/group_configuration_file_test.go`
- Create: `v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/WrongGroup_Filter.yaml`
- Create: `v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/WrongGroup_Transformer.yaml`
- Create: `v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/MatchingGroup_Filter.yaml`
- Create: `v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/OmittedGroup_Filter.yaml`

Validation rules:
- If a TypeFilter/TypeTransformer has an explicit `group` that doesn't match the import key → error
- If `group` is omitted → auto-fill with import key
- If `group` matches the import key → no error

- [ ] **Step 1: Write test fixtures**

`WrongGroup_Filter.yaml`:
```yaml
typeFilters:
  - action: include
    group: wronggroup
    version: v*20210601
    because: "wrong group"
```

`WrongGroup_Transformer.yaml`:
```yaml
typeTransformers:
  - group: wronggroup
    name: SomeType
    remove: true
    because: "wrong group"
```

`MatchingGroup_Filter.yaml`:
```yaml
typeFilters:
  - action: include
    group: testgroup
    version: v*20210601
    because: "matching group"
```

`OmittedGroup_Filter.yaml`:
```yaml
typeFilters:
  - action: include
    version: v*20210601
    because: "omitted group should be auto-filled"
```

- [ ] **Step 2: Write failing tests**

```go
func TestGroupConfigurationFile_WhenFilterHasWrongGroup_ReturnsError(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    _, err := loadGroupConfigurationFile(
        "testdata/TestGroupConfigurationFile/WrongGroup_Filter.yaml",
        "testgroup",
    )
    g.Expect(err).To(gomega.HaveOccurred())
    g.Expect(err.Error()).To(gomega.ContainSubstring("wronggroup"))
}

func TestGroupConfigurationFile_WhenTransformerHasWrongGroup_ReturnsError(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    _, err := loadGroupConfigurationFile(
        "testdata/TestGroupConfigurationFile/WrongGroup_Transformer.yaml",
        "testgroup",
    )
    g.Expect(err).To(gomega.HaveOccurred())
    g.Expect(err.Error()).To(gomega.ContainSubstring("wronggroup"))
}

func TestGroupConfigurationFile_WhenFilterHasMatchingGroup_Succeeds(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    gcf, err := loadGroupConfigurationFile(
        "testdata/TestGroupConfigurationFile/MatchingGroup_Filter.yaml",
        "testgroup",
    )
    g.Expect(err).ToNot(gomega.HaveOccurred())
    g.Expect(gcf.TypeFilters).To(gomega.HaveLen(1))
}

func TestGroupConfigurationFile_WhenFilterGroupOmitted_AutoFillsGroup(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    gcf, err := loadGroupConfigurationFile(
        "testdata/TestGroupConfigurationFile/OmittedGroup_Filter.yaml",
        "testgroup",
    )
    g.Expect(err).ToNot(gomega.HaveOccurred())
    g.Expect(gcf.TypeFilters).To(gomega.HaveLen(1))
    g.Expect(gcf.TypeFilters[0].Group.IsRestrictive()).To(gomega.BeTrue())
}
```

- [ ] **Step 3: Run tests to verify they fail**

Run: `cd v2 && go test ./tools/generator/internal/config/ -run TestGroupConfigurationFile -v`
Expected: New validation tests FAIL (validation not yet implemented)

- [ ] **Step 4: Implement validation in loadGroupConfigurationFile**

Add a `validateAndFillGroup` method to `GroupConfigurationFile` in `group_configuration_file.go`:

```go
// validateAndFillGroup checks that all TypeFilters and TypeTransformers
// either have no explicit group or have a group matching the expected group name.
// Omitted groups are auto-filled with the expected group name.
func (gcf *GroupConfigurationFile) validateAndFillGroup(groupName string) error {
    for _, f := range gcf.TypeFilters {
        if f.Group.IsRestrictive() {
            if !f.Group.Matches(groupName).Matched {
                return eris.Errorf(
                    "type filter has group %q but is in configuration file for group %q",
                    f.Group.String(), groupName)
            }
        } else {
            f.Group = NewFieldMatcher(groupName)
        }
    }

    for _, t := range gcf.Transformers {
        if t.Group.IsRestrictive() {
            if !t.Group.Matches(groupName).Matched {
                return eris.Errorf(
                    "type transformer has group %q but is in configuration file for group %q",
                    t.Group.String(), groupName)
            }
        } else {
            t.Group = NewFieldMatcher(groupName)
        }
    }

    return nil
}
```

Call `validateAndFillGroup` from `loadGroupConfigurationFile`, after YAML decoding but before returning.

**Note:** Check that `FieldMatcher` has a `Matches(string)` method and a `String()` method. If not, use whatever API is available (check `field_matcher.go` and `type_matcher.go`). The implementer should verify the exact FieldMatcher API and adapt accordingly.

- [ ] **Step 5: Run tests to verify they pass**

Run: `cd v2 && go test ./tools/generator/internal/config/ -run TestGroupConfigurationFile -v`
Expected: All tests PASS

- [ ] **Step 6: Commit**

```bash
git add v2/tools/generator/internal/config/group_configuration_file.go \
        v2/tools/generator/internal/config/group_configuration_file_test.go \
        v2/tools/generator/internal/config/testdata/TestGroupConfigurationFile/
git commit -m "Add group validation and auto-fill for GroupConfigurationFile

Rejects explicit group mismatches, auto-fills omitted groups.

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 3: Add merge logic and integrate into LoadConfiguration

**Files:**
- Modify: `v2/tools/generator/internal/config/configuration.go` (add `Imports` field, extend `LoadConfiguration`)
- Modify: `v2/tools/generator/internal/config/group_configuration_file.go` (add merge method)
- Modify: `v2/tools/generator/internal/config/group_configuration_file_test.go`
- Create: test fixtures for integration tests

This task adds the `Imports` field to `Configuration`, implements the merge logic, and integrates import loading into `LoadConfiguration()`.

- [ ] **Step 1: Write test fixture for end-to-end import loading**

Create `v2/tools/generator/internal/config/testdata/TestLoadConfiguration/`:

File `main_with_imports.yaml` — a minimal main config that imports a group file:
```yaml
schemaRoot: "./schemas"
destinationGoModuleFile: "./go.mod"
imports:
  testgroup: "testgroup.yaml"
typeFilters:
  - action: prune
    version: '*preview'
    because: "global preview filter"
    matchRequired: false
typeTransformers:
  - name: "*"
    property: Tags
    ifType:
      map:
        key:
          name: string
        value:
          name: any
    target:
      map:
        key:
          name: string
        value:
          name: string
    because: "global tags transform"
    matchRequired: false
```

File `testgroup.yaml` in the same directory:
```yaml
typeFilters:
  - action: include
    version: v*20210601
    because: "include specific version"
    matchRequired: false

groupModelConfiguration:
  2021-06-01:
    SomeResource:
      $exportAs: SomeResource
      $supportedFrom: v2.0.0
```

- [ ] **Step 2: Write failing tests for merge ordering and integration**

```go
func TestGroupConfigurationFile_MergesTypeFilters_Prepended(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    // Create a Configuration with one global filter
    cfg := NewConfiguration()
    globalFilter := &TypeFilter{Action: TypeFilterPrune}
    cfg.TypeFilters = []*TypeFilter{globalFilter}

    // Create a GroupConfigurationFile with one group filter
    gcf := &GroupConfigurationFile{
        TypeFilters: []*TypeFilter{
            {Action: TypeFilterInclude},
        },
    }

    err := gcf.mergeInto(cfg, "testgroup")
    g.Expect(err).ToNot(gomega.HaveOccurred())
    g.Expect(cfg.TypeFilters).To(gomega.HaveLen(2))
    // Group filter should be prepended (first)
    g.Expect(cfg.TypeFilters[0].Action).To(gomega.Equal(TypeFilterInclude))
    // Global filter should be second
    g.Expect(cfg.TypeFilters[1].Action).To(gomega.Equal(TypeFilterPrune))
}

func TestGroupConfigurationFile_MergesTransformers_Appended(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    cfg := NewConfiguration()
    globalTransformer := &TypeTransformer{}
    globalTransformer.RenameTo = "GlobalRename"
    cfg.Transformers = []*TypeTransformer{globalTransformer}

    gcf := &GroupConfigurationFile{
        Transformers: []*TypeTransformer{
            {RenameTo: "GroupRename"},
        },
    }

    err := gcf.mergeInto(cfg, "testgroup")
    g.Expect(err).ToNot(gomega.HaveOccurred())
    g.Expect(cfg.Transformers).To(gomega.HaveLen(2))
    // Global transformer should be first
    g.Expect(cfg.Transformers[0].RenameTo).To(gomega.Equal("GlobalRename"))
    // Group transformer should be appended (second)
    g.Expect(cfg.Transformers[1].RenameTo).To(gomega.Equal("GroupRename"))
}

func TestGroupConfigurationFile_MergesGroupModelConfiguration(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    cfg := NewConfiguration()
    gc := NewGroupConfiguration("testgroup")
    gcf := &GroupConfigurationFile{
        GroupModelConfiguration: gc,
    }

    err := gcf.mergeInto(cfg, "testgroup")
    g.Expect(err).ToNot(gomega.HaveOccurred())
    // Verify the group was added to ObjectModelConfiguration
    g.Expect(cfg.ObjectModelConfiguration.groups).To(gomega.HaveKey("testgroup"))
}

func TestGroupConfigurationFile_DuplicateGroup_ReturnsError(t *testing.T) {
    t.Parallel()
    g := gomega.NewGomegaWithT(t)

    cfg := NewConfiguration()
    // Pre-add a group
    existing := NewGroupConfiguration("testgroup")
    err := cfg.ObjectModelConfiguration.addGroup("testgroup", existing)
    g.Expect(err).ToNot(gomega.HaveOccurred())

    // Try to merge another group with same name
    gc := NewGroupConfiguration("testgroup")
    gcf := &GroupConfigurationFile{
        GroupModelConfiguration: gc,
    }

    err = gcf.mergeInto(cfg, "testgroup")
    g.Expect(err).To(gomega.HaveOccurred())
    g.Expect(err.Error()).To(gomega.ContainSubstring("duplicate"))
}
```

- [ ] **Step 3: Run tests to verify they fail**

Run: `cd v2 && go test ./tools/generator/internal/config/ -run "TestGroupConfigurationFile_Merges|TestGroupConfigurationFile_Duplicate" -v`
Expected: FAIL (mergeInto doesn't exist)

- [ ] **Step 4: Implement merge logic**

Add to `group_configuration_file.go`:

```go
// mergeInto merges this group configuration file's contents into the given Configuration.
// TypeFilters are prepended, Transformers are appended, and GroupModelConfiguration
// is added to ObjectModelConfiguration.
func (gcf *GroupConfigurationFile) mergeInto(cfg *Configuration, groupName string) error {
    // Prepend group-specific TypeFilters before global filters
    if len(gcf.TypeFilters) > 0 {
        cfg.TypeFilters = append(gcf.TypeFilters, cfg.TypeFilters...)
    }

    // Append group-specific Transformers after global transformers
    if len(gcf.Transformers) > 0 {
        cfg.Transformers = append(cfg.Transformers, gcf.Transformers...)
    }

    // Add GroupModelConfiguration to ObjectModelConfiguration
    if gcf.GroupModelConfiguration != nil {
        if err := cfg.ObjectModelConfiguration.addGroup(groupName, gcf.GroupModelConfiguration); err != nil {
            return eris.Wrapf(err, "merging group %q", groupName)
        }
    }

    return nil
}
```

- [ ] **Step 5: Run tests to verify they pass**

Run: `cd v2 && go test ./tools/generator/internal/config/ -run "TestGroupConfigurationFile_Merges|TestGroupConfigurationFile_Duplicate" -v`
Expected: All PASS

- [ ] **Step 6: Add Imports field to Configuration and integrate into LoadConfiguration**

In `configuration.go`, add the `Imports` field to the `Configuration` struct (after `ObjectModelConfiguration`, around line 65):

```go
    // Imported group-specific configuration files
    Imports map[string]string `yaml:"imports,omitempty"`
```

Then add a new method `loadImports` and call it from `LoadConfiguration()` after YAML decode but before `initialize()`:

```go
// loadImports loads and merges all imported group configuration files
func (c *Configuration) loadImports(configDir string) error {
    if len(c.Imports) == 0 {
        return nil
    }

    // Sort keys for deterministic ordering.
    // Note: When multiple groups prepend TypeFilters, later groups (alphabetically)
    // end up earlier in the final slice. This is functionally harmless because
    // group-specific filters only match their own group.
    groups := make([]string, 0, len(c.Imports))
    for group := range c.Imports {
        groups = append(groups, group)
    }
    sort.Strings(groups)

    for _, group := range groups {
        filePath := c.Imports[group]
        // configDir may be relative; filepath.Join handles this correctly
        // as long as the working directory hasn't changed since LoadConfiguration was called.
        absPath := filepath.Join(configDir, filePath)

        gcf, err := loadGroupConfigurationFile(absPath, group)
        if err != nil {
            return eris.Wrapf(err, "loading import for group %q from %q", group, filePath)
        }

        if err := gcf.validateAndFillGroup(group); err != nil {
            return eris.Wrapf(err, "validating import for group %q", group)
        }

        if err := gcf.mergeInto(c, group); err != nil {
            return err
        }
    }

    return nil
}
```

In `LoadConfiguration()`, add the import loading call after YAML decode (around line 174):

```go
    // Load and merge any imported group configuration files
    configDir := filepath.Dir(configurationFile)
    if err := result.loadImports(configDir); err != nil {
        return nil, err
    }
```

Add `"sort"` to the imports if not already present.

- [ ] **Step 7: Run all config tests**

Run: `cd v2 && go test ./tools/generator/internal/config/ -v`
Expected: All tests PASS (both new and existing)

- [ ] **Step 8: Commit**

```bash
git add v2/tools/generator/internal/config/configuration.go \
        v2/tools/generator/internal/config/group_configuration_file.go \
        v2/tools/generator/internal/config/group_configuration_file_test.go \
        v2/tools/generator/internal/config/testdata/
git commit -m "Add merge logic and integrate imports into LoadConfiguration

TypeFilters prepend, Transformers append, GroupModelConfiguration merges.
Imports processed in sorted order for deterministic behavior.

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 4: Create the YAML splitting tool

**Files:**
- Create: `v2/tools/generator/internal/config/cmd/splitconfig/main.go` (temporary tool)

This task creates a Go program that reads `azure-arm.yaml` and produces:
1. Per-group YAML files in `v2/azure-arm/`
2. An updated `azure-arm.yaml` with group-specific content removed and `imports` section added

This is a one-time-use tool that can be deleted after the split is complete.

- [ ] **Step 1: Write the splitting tool**

Create `v2/tools/generator/internal/config/cmd/splitconfig/main.go`:

The tool should:
1. Parse `v2/azure-arm.yaml` using `yaml.Node` to preserve comments and ordering
2. For each group in `objectModelConfiguration`:
   - Create `v2/azure-arm/<group>.yaml`
   - Move the group's `objectModelConfiguration` entry → `groupModelConfiguration`
   - Move any `typeFilters` entries with matching `group:` field → `typeFilters`
   - Move any `typeTransformers` entries with matching `group:` field → `typeTransformers`
   - Strip the `group:` field from moved filters/transformers (since it's implied)
3. Update `v2/azure-arm.yaml`:
   - Add `imports:` section mapping each group to its file
   - Remove extracted entries from `typeFilters`, `typeTransformers`, `objectModelConfiguration`

**Important implementation notes:**
- Use `yaml.Node` tree manipulation (not struct-level) to preserve YAML comments
- The 45 groups to extract are listed in the spec
- TypeFilters with explicit `group:` — there are 23 of these
- TypeTransformers with explicit `group:` — there are only 2 of these
- Group names in typeFilter/typeTransformer `group:` fields may use semicolons for multiple groups (e.g., `group: "foo;bar"`) — only extract if the group field matches exactly one group

- [ ] **Step 2: Run the splitting tool**

```bash
cd v2 && go run ./tools/generator/internal/config/cmd/splitconfig/
```

- [ ] **Step 3: Verify the split files look correct**

Manually inspect a few group files (e.g., `v2/azure-arm/storage.yaml`, `v2/azure-arm/containerservice.yaml`) to verify:
- `groupModelConfiguration` has the correct version→type→property structure
- Any extracted `typeFilters` and `typeTransformers` are present
- The `group:` field has been stripped from extracted filters/transformers

Also verify `v2/azure-arm.yaml`:
- Has the `imports:` section
- `objectModelConfiguration:` section is empty or removed
- Group-specific `typeFilters` and `typeTransformers` have been removed
- Global filters and transformers remain

- [ ] **Step 4: Commit the split**

```bash
git add v2/azure-arm.yaml v2/azure-arm/
git commit -m "Split azure-arm.yaml into per-group configuration files

Extracts 45 groups into individual files under v2/azure-arm/.
Each group file contains typeFilters, typeTransformers, and
groupModelConfiguration for that group.

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

- [ ] **Step 5: Delete the splitting tool**

```bash
rm -rf v2/tools/generator/internal/config/cmd/splitconfig/
git add -A && git commit -m "Remove one-time splitting tool

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 5: Verify the split produces identical output

**Files:** None modified (verification only)

- [ ] **Step 1: Run code generation with split config**

Run: `cd v2 && ../hack/tools/task controller:generate-types`

Expected: Succeeds with no errors. The generated output should be identical to what the monolithic config produced.

- [ ] **Step 2: Verify no diff in generated code**

```bash
cd v2 && git diff --stat
```

Expected: No changes to generated files (only the config files and new Go code should differ from the base branch).

- [ ] **Step 3: If generated output differs, debug and fix**

If there are differences:
1. Check merge ordering (TypeFilters prepend, Transformers append)
2. Check that all groups were extracted correctly
3. Check that group-specific filters/transformers have correct group fields
4. Fix any issues and re-run

- [ ] **Step 4: Run generator quick-checks**

Run: `cd v2 && ../hack/tools/task generator:quick-checks`

Expected: All checks pass

- [ ] **Step 5: Run controller quick-checks**

Run: `cd v2 && ../hack/tools/task controller:quick-checks`

Expected: All checks pass

- [ ] **Step 6: Run asoctl quick-checks**

Run: `cd v2 && ../hack/tools/task asoctl:quick-checks`

Expected: All checks pass

- [ ] **Step 7: Final commit if any fixups were needed**

If any changes were needed during verification:
```bash
git add -A && git commit -m "Fix config split issues found during verification

Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```
