---
name: testing-aso-recordings
description: Use when you need to verify that ASO controller CRUD tests or sample tests are correct and functional, or when go-vcr recordings need to be replayed or re-recorded.
---

# Verify ASO Recording-Based Tests

ASO uses go-vcr recording/replay for two test suites: **controller CRUD tests** and **sample tests**. This skill covers running, debugging, and re-recording both.

## Test Suites

| | Controllers | Samples |
|---|---|---|
| **Task command** | `controller:test-controllers` | `controller:test-samples` |
| **Log file** | `reports/test-controllers.log` | `reports/test-samples.log` |
| **Test source** | `v2/internal/controllers/` | `v2/internal/testsamples/samples_test.go` |
| **Suite setup** | `v2/internal/controllers/suite_test.go` | `v2/internal/testsamples/suite_test.go` |
| **Recordings dir** | `v2/internal/controllers/recordings/` | `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/` |
| **Playback (full suite)** | 6–10 min | 5–10 min |
| **Playback (single test)** | 1–3 min | 2–3 min |

**When to run which suite:**

- **Controllers** — after creating new resources, importing new resource versions, implementing or modifying resource extensions, or making changes that alter HTTP request sequences to Azure.
- **Samples** — after creating new samples, modifying existing samples, implementing resource extensions, or making changes that alter HTTP request sequences to Azure. Samples are found in `v2/samples/`, grouped by resource group.

## Workflow

### Preflight check

Before running tests:

1. **Check environment variables.** Verify that `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and `ENTRA_APP_ID` are set (or that the user has provided a `test.env` file to source). These are required for **recording** new tests against live Azure. If any are missing and recordings may need to be created, do not proceed — tell the user what is missing and stop. Playback runs do not need them.

2. **Samples only — check for hardcoded test filters.** Verify there are no hardcoded test filters in `v2/internal/testsamples/samples_test.go`. The `Test_Samples_CreationAndDeletion` function walks the samples directory and dynamically creates subtests. Sometimes a developer temporarily adds a filter (e.g. an `if strings.Contains(testName, "...")` guard around the `t.Run` call) and forgets to remove it. If present, remove it so all sample tests run.

### Run tests

Run all tests as a background terminal command:

```bash
source test.env && ./hack/tools/task controller:<SUITE>
```

Replace `<SUITE>` with `test-controllers` or `test-samples`. If you're confident all recordings exist, omit `source test.env`.

Do NOT add `| tee` or other output redirection — the taskfile already pipes output to the log file.

Use `await_terminal` to wait for completion (see **Monitoring** in the Guidance section for timeout recommendations).

- If all recordings exist, playback completes in the times shown in the Test Suites table. Do not interrupt.
- If recordings are missing, each triggers a live Azure test. Most take 20–30 min; some take 60+ min.
- No visible output for several minutes is normal.

#### Running specific tests

Add `TEST_FILTER` to run a subset:

**Controllers:**
```bash
TEST_FILTER="<your-test-here>" ./hack/tools/task controller:test-controllers
```

**Samples:**
```bash
TEST_FILTER="Test_Samples_CreationAndDeletion/<your-test-here>" ./hack/tools/task controller:test-samples
```

**Sample test naming:** For a sample at `v2/samples/<group>/<subgroup>/<version>/`, the test name is `Test_<TitleCasedSubgroup>_<version>_CreationAndDeletion`. Example: `v2/samples/documentdb/mongodb/v1api20231115/` → `Test_Mongodb_v1api20231115_CreationAndDeletion`. (Only the last path segment before the version is used.)

### Verify success

After the test command finishes, confirm no failures in the log:

```bash
grep "FAIL:" reports/<LOG-FILE> || echo "No failures found"
```

Replace `<LOG-FILE>` with the appropriate log file from the Test Suites table.

- No failures → you're done.
- Failures found → proceed to the next section.

### Identify test failures

Look through the logs to identify which tests failed and why.

- Logs overflow the terminal buffer, so use the log file (see Test Suites table).
- Search for `FAIL:` (trailing colon avoids false positives).
- Initial failures often cascade — identify the **first** problem.

### Handling test failures

Known failure modes — go through failures and make a TODO list of next actions.

**IMPORTANT:** Before deleting a recording and re-recording, always investigate WHY the test failed. Blindly re-recording wastes time (30–60 min) and may not fix the problem. If you can't determine the root cause, report to the user and ask for guidance.

**Missing interaction** — go-vcr error: "no responses recorded for this method/URL".
- If recording is unmodified (`git status` shows clean) → delete recording and re-record.
- If recording was already updated → investigate further; report to user.

**Body hash mismatch** — go-vcr error: "body mismatch" / "Request body hash header mismatch". Different from missing interaction.
- Recording unmodified → stale, re-record.
- Recording just created → non-deterministic serialization. Do NOT keep re-recording. Investigate VCR matching/hashing code; report to user.

**Test assertion failure** — assertion error from test code (not go-vcr).
- During recording (live Azure) → likely real bug; report to user.
- During playback → stale cached responses; may need re-recording. Check if known replay issue first; report to user.

**Timeout while recording** — extend timeout with `TIMEOUT=90m` on the command line. Some resources (ApiManagement, Kusto) take 30+ min to provision. If still fails, investigate and report.

**Lack of capacity or quota** — `ServiceUnavailable` with "high demand" message:
1. Check the sample YAML for `location`/`locationName` fields.
2. Prefer consistency with other samples in the same resource group.
3. Change to a different Azure region (e.g. `eastus` → `australiaeast`).
4. Delete failed/partial recording and re-run.
5. If new region also fails, try another.

### Update a recording

Delete the recording file, then run as a **background terminal**:

**Controllers:**
```bash
source test.env && TIMEOUT=60m TEST_FILTER="<your-test-here>" ./hack/tools/task controller:test-controllers
```

**Samples:**
```bash
source test.env && TIMEOUT=60m TEST_FILTER="Test_Samples_CreationAndDeletion/<your-test-here>" ./hack/tools/task controller:test-samples
```

For slow-provisioning resources, use `TIMEOUT=90m`.

Update recordings one at a time. After a successful recording, **verify playback** by running the same command without `source test.env`. Playback takes 1–3 min. If playback fails but recording succeeded, this is a systemic VCR issue — do not keep re-recording; investigate and report.

NEVER modify recording files by hand. They are machine-generated.

#### Handling failed recording attempts

1. **Partial recording files** — a failed test may write a partial recording. Always delete it before retrying.

2. **Azure debris** — failed runs can leave resources in Azure that weren't cleaned up. On retry, the test may find pre-existing resources and behave unexpectedly. Wait a few minutes for Azure cleanup, delete the partial recording, and try again.

### Confirmation

Once updated recordings pass, **start over from the top** and run the full suite again. A test run exits on first failure, leaving later tests unexecuted. Re-running confirms all problems are resolved and there are no hidden issues.

## Guidance

- **DO NOT interrupt a test run**, even if slow. Interrupting can leave Azure debris that pollutes the next recording attempt.
- **Monitoring long-running tests:** Run as background terminals and use `await_terminal`. DO NOT poll logs with `tail`/`grep` in a loop. Recommended timeouts:
  - Playback (full suite): 10 min (600000ms)
  - Playback (single test): 5 min (300000ms)
  - Recording (single test): 10 min (600000ms), then repeat — recordings take 20–60+ min
  - Resource group deletion: 15–25 min for groups with nested resources
  - If `await_terminal` times out, just call it again — the terminal keeps running.
- **DO NOT analyze recording YAML files** to debug mismatches. They're large and machine-generated. Delete and re-record instead.
- **DO NOT use `task` directly** — always use `./hack/tools/task`.
- **Test runs are CPU/time intensive** and can starve other processes. Keep detailed notes about progress to allow recovery if this happens.

## Postrequisites

When finished, provide a table to the user showing the tests that have been recorded, and why.
