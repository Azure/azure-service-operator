# Rapid Tests for Five Small Resource Groups Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace gopter-generated property tests with rapid-generated property tests for `monitor`, `network.frontdoor`, `quota`, `redhatopenshift`, and `resources`.

**Architecture:** The generator already selects the property-test implementation through `gopterGroups`. Remove the five groups from that set, protect the selection with a focused unit test, and regenerate their checked-in API tests through the existing pipeline.

**Tech Stack:** Go 1.25+, Task, ASO code generator, Gomega, pgregory/rapid

## Global Constraints

- Migrate exactly `monitor`, `network.frontdoor`, `quota`, `redhatopenshift`, and `resources`.
- Do not change generator architecture, public APIs, or operator runtime behavior.
- Keep the migration and regenerated output in one focused implementation commit.
- Generated `*_gen_test.go` files must be updated only by the generator.

---

### Task 1: Gate and Generate Rapid Tests

**Files:**
- Create: `v2/tools/generator/internal/testcases/rapid_migration_test.go`
- Modify: `v2/tools/generator/internal/testcases/rapid_migration.go:12-58`
- Modify (generated): `v2/api/monitor/v1api20230403/account_types_gen_test.go`
- Modify (generated): `v2/api/monitor/v1api20230403/arm/account_spec_types_gen_test.go`
- Modify (generated): `v2/api/monitor/v1api20230403/arm/account_status_types_gen_test.go`
- Modify (generated): `v2/api/monitor/v1api20230403/storage/account_types_gen_test.go`
- Modify (generated): `v2/api/network.frontdoor/v1api20220501/web_application_firewall_policy_types_gen_test.go`
- Modify (generated): `v2/api/network.frontdoor/v1api20220501/arm/web_application_firewall_policy_spec_types_gen_test.go`
- Modify (generated): `v2/api/network.frontdoor/v1api20220501/arm/web_application_firewall_policy_status_types_gen_test.go`
- Modify (generated): `v2/api/network.frontdoor/v1api20220501/storage/web_application_firewall_policy_types_gen_test.go`
- Modify (generated): `v2/api/quota/v1api20250901/quota_types_gen_test.go`
- Modify (generated): `v2/api/quota/v1api20250901/arm/quota_spec_types_gen_test.go`
- Modify (generated): `v2/api/quota/v1api20250901/arm/quota_status_types_gen_test.go`
- Modify (generated): `v2/api/quota/v1api20250901/storage/quota_types_gen_test.go`
- Modify (generated): `v2/api/redhatopenshift/v1api20231122/open_shift_cluster_types_gen_test.go`
- Modify (generated): `v2/api/redhatopenshift/v1api20231122/arm/open_shift_cluster_spec_types_gen_test.go`
- Modify (generated): `v2/api/redhatopenshift/v1api20231122/arm/open_shift_cluster_status_types_gen_test.go`
- Modify (generated): `v2/api/redhatopenshift/v1api20231122/storage/open_shift_cluster_types_gen_test.go`
- Modify (generated): `v2/api/resources/v1api20200601/resource_group_types_gen_test.go`
- Modify (generated): `v2/api/resources/v1api20200601/arm/resource_group_spec_types_gen_test.go`
- Modify (generated): `v2/api/resources/v1api20200601/arm/resource_group_status_types_gen_test.go`
- Modify (generated): `v2/api/resources/v1api20200601/storage/resource_group_types_gen_test.go`

**Interfaces:**
- Consumes: `UseRapidForGroup(group string) bool` and `UseGopterForGroup(group string) bool`
- Produces: rapid-generated serialization, assignment, and conversion tests for the five groups

- [ ] **Step 1: Write the failing migration test**

```go
/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRapidMigration_SelectedSmallGroupsUseRapid(t *testing.T) {
	t.Parallel()

	groups := []string{
		"monitor",
		"network.frontdoor",
		"quota",
		"redhatopenshift",
		"resources",
	}

	for _, group := range groups {
		group := group
		t.Run(group, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(UseRapidForGroup(group)).To(BeTrue())
			g.Expect(UseGopterForGroup(group)).To(BeFalse())
		})
	}
}
```

- [ ] **Step 2: Run the focused test and confirm the current gate fails**

Run:

```bash
TEST_FILTER=TestRapidMigration_SelectedSmallGroupsUseRapid ./hack/tools/task generator:unit-tests
```

Expected: FAIL for all five subtests because each group is still present in `gopterGroups`.

- [ ] **Step 3: Remove the five groups from the gopter migration set**

Delete these exact entries from `gopterGroups`:

```go
	"monitor",
	"network.frontdoor",
	"quota",
	"redhatopenshift",
	"resources",
```

- [ ] **Step 4: Run the focused test and confirm the rapid gate passes**

Run:

```bash
TEST_FILTER=TestRapidMigration_SelectedSmallGroupsUseRapid ./hack/tools/task generator:unit-tests
```

Expected: PASS for the parent test and all five subtests.

- [ ] **Step 5: Regenerate checked-in API tests**

Run:

```bash
./hack/tools/task controller:generate-types
```

Expected: the 20 listed `*_gen_test.go` files are rewritten to use rapid; no handwritten API or runtime files change.

- [ ] **Step 6: Verify the generated migration**

Run:

```bash
rg 'pgregory.net/rapid' \
  v2/api/monitor \
  v2/api/network.frontdoor \
  v2/api/quota \
  v2/api/redhatopenshift \
  v2/api/resources

if rg 'leanovate/gopter|gopter\.' \
  v2/api/monitor \
  v2/api/network.frontdoor \
  v2/api/quota \
  v2/api/redhatopenshift \
  v2/api/resources; then
  exit 1
fi
```

Expected: the first command finds rapid imports and generated calls; the second command produces no matches.

- [ ] **Step 7: Format the implementation**

Run:

```bash
./hack/tools/task format-code
```

Expected: PASS and any formatting changes are limited to Task 1 files.

- [ ] **Step 8: Review and commit the implementation**

Run:

```bash
git diff --check
git status --short
git diff --stat
git add \
  v2/tools/generator/internal/testcases/rapid_migration.go \
  v2/tools/generator/internal/testcases/rapid_migration_test.go \
  v2/api/monitor \
  v2/api/network.frontdoor \
  v2/api/quota \
  v2/api/redhatopenshift \
  v2/api/resources
git commit -m "test: enable rapid tests for five small groups" \
  -m "Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>"
```

Expected: one implementation commit containing only the migration gate, its unit test, and regenerated tests.

### Task 2: Validate and Open the Pull Request

**Files:**
- Verify: all files changed by Task 1
- Verify: `docs/superpowers/specs/2026-08-06-rapid-tests-small-groups-design.md`
- Verify: `docs/superpowers/plans/2026-08-06-rapid-tests-small-groups.md`

**Interfaces:**
- Consumes: the committed five-group migration from Task 1
- Produces: a pushed branch and GitHub pull request

- [ ] **Step 1: Run generator checks**

Run:

```bash
./hack/tools/task generator:quick-checks
```

Expected: PASS.

- [ ] **Step 2: Run controller checks**

Run:

```bash
./hack/tools/task controller:quick-checks
```

Expected: PASS.

- [ ] **Step 3: Run asoctl checks**

Run:

```bash
./hack/tools/task asoctl:quick-checks
```

Expected: PASS.

- [ ] **Step 4: Confirm the branch is ready**

Run:

```bash
git status --short
git log --oneline origin/main..HEAD
git diff --check origin/main...HEAD
```

Expected: a clean worktree, the design/plan commits plus one focused implementation commit, and no whitespace errors.

- [ ] **Step 5: Push and open the pull request**

Run:

```bash
git push --set-upstream origin HEAD
gh pr create \
  --base main \
  --title "Enable rapid tests for five small resource groups" \
  --body "$(cat <<'EOF'
## Summary

- migrate monitor, network.frontdoor, quota, redhatopenshift, and resources property tests from gopter to rapid
- add regression coverage for the migration gate
- regenerate checked-in API property tests

## Validation

- ./hack/tools/task generator:quick-checks
- ./hack/tools/task controller:quick-checks
- ./hack/tools/task asoctl:quick-checks
EOF
)"
```

Expected: GitHub returns the URL of a new pull request targeting `main`.
