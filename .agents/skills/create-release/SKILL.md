---
name: create-release
description: Use when creating a new ASO release, including pre-checks, release notes generation, draft GitHub release creation, documentation updates, and ROADMAP changes.
---

# Create ASO Release

This skill automates the end-to-end process of creating a new Azure Service Operator release. It runs pre-checks, generates release notes, creates a draft GitHub release, updates documentation, and opens a docs PR.

## Determine the version

Determine the new version number by examining the latest release tag:
```bash
git describe --tags $(git rev-list --tags='v2*' --max-count=1) --match 'v2*'
```
If the latest tag is `v2.19.0`, the new version is `v2.20.0` (increment the minor version).

## Step 1: Pre-checks

All three pre-checks **MUST** pass before proceeding. If any check fails, **STOP** and report the failure. Do not continue with release creation.

### 1a. Pre-release upgrade tests

Check the [pre-release upgrade tests workflow](https://github.com/Azure/azure-service-operator/actions/workflows/pre-release-tests.yaml):
```bash
gh run list --workflow=pre-release-tests.yaml --limit=1 -R Azure/azure-service-operator --json status,conclusion,createdAt,url
```
- **Must be:** `conclusion: success`
- **Must be within:** the last 7 days

### 1b. Controller CVE scan

Check the [scan controller image workflow](https://github.com/Azure/azure-service-operator/actions/workflows/scan-controller-image.yaml):
```bash
gh run list --workflow=scan-controller-image.yaml --limit=1 -R Azure/azure-service-operator --json status,conclusion,createdAt,url
```
- **Must be:** `conclusion: success`
- **Must be within:** the last 7 days

### 1c. Experimental release

Check the [experimental release](https://github.com/Azure/azure-service-operator/releases/tag/experimental):
```bash
gh release view experimental -R Azure/azure-service-operator --json tagName,createdAt,isDraft,isPrerelease,assets
```
- **Must be:** published (not draft), all expected assets uploaded (9 assets: 5 asoctl binaries, CRD yaml, operator yaml, multitenant-cluster yaml, multitenant-tenant yaml)
- **Must be recent** (within a few days of HEAD)

If any pre-check fails, report which check failed, why, and provide the relevant URL. Do not proceed.

## Step 2: Generate release notes

Invoke the `create-release-notes` skill (defined in `.agents/skills/create-release-notes/SKILL.md`) to generate polished release notes. This involves:

1. Running `scripts/v2/generate-raw-changelog.sh` to get the raw changelog
2. Running `scripts/v2/generate-raw-changelog.sh --api-only` to identify API group changes
3. Formatting the changelog into sections: Breaking changes (if any), Upcoming breaking changes (if any), New resources, Features, Bug fixes, Documentation

Save the release notes to a temp file:
```bash
/tmp/aso-<VERSION>-release-notes.md
```

## Step 3: Create draft GitHub release

Use the `scripts/v2/create-release.sh` script to create a draft release:
```bash
scripts/v2/create-release.sh <VERSION> /tmp/aso-<VERSION>-release-notes.md
```

This creates a **draft** release on GitHub targeting the `main` branch. The release must be reviewed and published manually by a human.

## Step 4: Update documentation and ROADMAP

Create a local branch and make the following changes:

### 4a. Update `currentRelease` in config files

Update `supportedResourcesReport.currentRelease` to the new version in both:
- `v2/azure-arm.yaml`
- `hack/crossplane/azure-crossplane.yaml`

### 4b. Regenerate documentation

Run the code generators to update the resource documentation (moves resources from "Next Release" to "Released"):
```bash
./hack/tools/task controller:generate-types
./hack/tools/task crossplane:generate-types
./hack/tools/task doc:crd-api
```

**Note:** `controller:generate-types` takes ~5-7 minutes. `doc:crd-api` also runs `controller:generate-types` internally and takes ~10 minutes total. Set generous timeouts (900000ms+).

### 4c. Update ROADMAP.md

In `ROADMAP.md`:
1. **Add** the new version to the "Official Releases" table with today's date
2. **Remove** the new version from the "Release Cadence & Planning" table
3. **Add** a new future version to the end of the planning table (2 months after the last entry)

### 4d. Create the next milestone

Check if the milestone referenced by the new future version in the ROADMAP exists:
```bash
gh api repos/Azure/azure-service-operator/milestones --jq '.[] | "\(.number) \(.title)"'
```

If the milestone does not exist, create it following the pattern of existing milestones:
```bash
gh api repos/Azure/azure-service-operator/milestones -f title="v2.XX.0" -f description="" -f due_on="YYYY-MM-DDT00:00:00Z" -f state="open"
```

### 4e. Document breaking changes

If this release includes breaking changes, create or update `docs/hugo/content/guide/breaking-changes/breaking-changes-<VERSION>.md` (for example, `breaking-changes-v2.22.0.md`). Use the release notes' **Breaking changes** section, the changelog, and relevant release-planning issues to identify the changes.

Follow the pattern of the existing [v2.21.0 breaking changes page](../../../docs/hugo/content/guide/breaking-changes/breaking-changes-v2.21.0.md) and other recent pages in that directory:

- Set YAML frontmatter `title` to `"<VERSION> Breaking Changes"` and `linkTitle` to `"<VERSION>"`.
- Set `weight` to 5 lower than the most recent preceding breaking changes page (for example, `-85` after v2.21.0's `-80`). Preserve the explanatory comment used by existing pages.
- Add a `## Breaking changes` section describing each change's user-visible impact and the migration steps, configuration changes, or workarounds needed before upgrading. Include relevant PR or issue references.
- If upcoming changes also need to be announced, add a separate `## Future breaking changes` section, following the existing documentation convention. Clearly distinguish future plans or proposals from changes taking effect in this release.
- Keep the page consistent with the release notes. Do not overwrite an earlier release's page to document the new release.
- Update `docs/hugo/content/guide/breaking-changes/_index.md` if it still lists a change taking effect in this release as upcoming.

If there are no breaking changes in this release, skip creating a version-specific breaking changes page. Do not create an empty `## Breaking changes` section.

Include any breaking changes documentation updates in the documentation PR below. Run `./dev.sh -c 'task doc:frontmatter-check'` when adding or updating these pages.

### 4f. Commit and open PR

Follow the Conventional Commits guidance in [AGENTS.md](../../../AGENTS.md#detailed-instructions) and [CONTRIBUTING.md](../../../CONTRIBUTING.md#pull-request-titles-and-commit-messages) for both the documentation commit message and PR title.

1. Create a branch named `<user>/X.XX-docs` (e.g. `matthchr/2.20-docs`)
2. Commit all changes with message `docs: update docs for <major>.<minor>` (e.g. `docs: update docs for 2.20`)
3. Push the branch
4. Open a PR against `main` in `Azure/azure-service-operator` with:
   - **Title:** `docs: update docs for <major>.<minor>`
   - **Body:**
     ```
     - [x] this PR contains documentation
     - [ ] this PR contains tests
     - [ ] this PR contains YAML Samples
     ```

## Final Report

At the end of the run, produce a summary report including:

1. **Pre-check results:**
   - Pre-release upgrade tests: PASS/FAIL, date, [link to run]
   - Controller CVE scan: PASS/FAIL, date, [link to run]
   - Experimental release: PASS/FAIL, date, [link to release]

2. **Release:**
   - Draft release URL (if created)

3. **Documentation PR:**
   - PR URL (if created)
   - Breaking changes documentation path (if created or updated), or confirmation that this release has no breaking changes

4. **Remaining manual steps:**
   - Review and publish the draft release on GitHub
   - After publishing, verify the release GitHub Action completes successfully
   - After publishing, compare the released CRDs with the prior release to check for additional breaking changes and update the documentation if needed (see `docs/hugo/content/contributing/create-a-new-release.md`)
   - Perform smoke testing (see same doc)
   - Create a blog post
   - Post a message in the Slack channel
