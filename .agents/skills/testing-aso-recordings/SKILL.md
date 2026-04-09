---
name: testing-aso-recordings
description: Use when you need to run ASO controller CRUD tests or sample tests, verify recordings play back correctly, or re-record after changes.
---

# Run ASO Recording-Based Tests

ASO uses [go-vcr](https://github.com/gophercloud/go-vcr) to record and replay HTTP interactions with Azure. Tests run in two modes:

- **Playback** — replays previously recorded HTTP interactions. Fast (minutes), no Azure credentials needed, deterministic.
- **Recording** — makes live HTTP calls to Azure, captures interactions to YAML files. Slow (20–60+ min), requires credentials, creates new recordings for future playback.

The test framework automatically chooses: if a recording file exists, it replays; if not, it records live. Both test suites use envtest (a local Kubernetes API server) to run the controller under test.

## Test Suites

|                            | Controllers                             | Samples                                                                |
| -------------------------- | --------------------------------------- | ---------------------------------------------------------------------- |
| **Task command**           | `controller:test-controllers`           | `controller:test-samples`                                              |
| **Log file**               | `reports/test-controllers.log`          | `reports/test-samples.log`                                             |
| **Test source**            | `v2/internal/controllers/`              | `v2/internal/testsamples/samples_test.go`                              |
| **Suite setup**            | `v2/internal/controllers/suite_test.go` | `v2/internal/testsamples/suite_test.go`                                |
| **Recordings dir**         | `v2/internal/controllers/recordings/`   | `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/` |
| **Playback (full suite)**  | 6–10 min                                | 5–10 min                                                               |
| **Playback (single test)** | 1–3 min                                 | 2–3 min                                                                |

**When to run which suite:**

- **Controllers** — after creating new resources, importing new resource versions, implementing or modifying resource extensions, or making changes that alter HTTP request sequences to Azure.
- **Samples** — after creating new samples, modifying existing samples, implementing resource extensions, or making changes that alter HTTP request sequences to Azure. Samples are YAML resource definitions found in `v2/samples/`, grouped by resource group.

### Test file naming

**Controller tests** follow the pattern `<group>_<subject>_<scenario>_<version>_test.go` (e.g. `crd_networking_virtualnetwork_v1api20201101_test.go`). Each file contains hand-written Go test functions.

**Sample tests** are generated dynamically — `samples_test.go` walks `v2/samples/` and creates a subtest per version directory. Test names follow the pattern `Test_<TitleCasedSubgroup>_<version>_CreationAndDeletion`. Example: `v2/samples/documentdb/mongodb/v1api20231115/` → `Test_Mongodb_v1api20231115_CreationAndDeletion`. (Only the last directory segment before the version is used.)

## Workflow

### Preflight check

Before running tests:

1. **Check environment variables.** Verify that `AZURE_SUBSCRIPTION_ID` and `AZURE_TENANT_ID` are set (or that the user has provided a `test.env` file to source). These are required for **recording** (live Azure calls). `ENTRA_APP_ID` may also be needed depending on the test — check whether the test references it before requiring it. If required variables are missing and recordings need to be created, do not proceed — tell the user what is missing and stop. **Playback runs do not need them.**
2. **Check whether a recording already exists.** Look in the appropriate recordings directory for a file matching the test name. If a recording exists, this will be a playback run (fast). If not, this will be a recording run (slow, live Azure calls).

### Run tests

Run all tests as a background terminal command:

```bash
source test.env && task controller:<SUITE>
```

Replace `<SUITE>` with `test-controllers` or `test-samples`. For playback-only runs (all recordings exist), omit `source test.env`.

Do NOT add `| tee` or other output redirection — the taskfile already pipes output to the log file.

For recording runs of slow resources, override the default Go test timeout:

```bash
source test.env && TIMEOUT=60m TEST_FILTER="<test>" task controller:test-controllers
```

Run the command in async mode. Monitor progress using `sleep` + `tail` on the log file (see **Monitoring** below).

- If all recordings exist, playback completes in the times shown in the Test Suites table. Do not interrupt.
- If recordings are missing, each triggers a live Azure recording. Most take 20–30 min; some take 60+ min.
- No visible output for several minutes is normal.

#### Running specific tests

Add `TEST_FILTER` to run a subset:

**Controllers:**

```bash
TEST_FILTER="<your-test-here>" task controller:test-controllers
```

**Samples:**

```bash
TEST_FILTER="Test_Samples_CreationAndDeletion/<your-test-here>" ./hack/tools/task controller:test-samples
```

### Verify success

After the test command finishes, confirm no failures in the log:

```bash
grep "FAIL:" reports/<LOG-FILE> || echo "No failures found"
```

Replace `<LOG-FILE>` with the appropriate log file from the Test Suites table.

- No failures → you're done.
- Failures found → use the **diagnosing-vcr-failures** skill to investigate and resolve.

## Guidance

- **DO NOT interrupt a test run**, even if slow. Interrupting can leave Azure debris that pollutes the next recording attempt.
- **Monitoring long-running tests:** Run the test command in async mode, then use `sleep <seconds> && tail -3 reports/<LOG-FILE>` to check progress at intervals. DO NOT poll in a tight loop — use 300-second (5 min) sleeps between checks. Key log markers to watch for:
  - Log Section messages — test reached a new phase 
  - `"MonitorDelete"` — test in cleanup
  - `"PASS"` / `"FAIL"` — test completed
  - `"saving ARM client recorder"` — recording was written to disk
- **Expected durations:**
  - Playback (full suite): 6–10 min
  - Playback (single test): 1–3 min
  - Recording (single test): 20–60+ min depending on resource type
  - Resource group deletion: 15–25 min for groups with nested resources (e.g., Cassandra clusters, AKS)
- **ALWAYS use `task` directly** — never use `./hack/tools/task`. If `task` is not on the PATH, your environment is not set up correctly for testing. Stop and ask the user to fix your environment.
- **Test runs are CPU/time intensive** and can starve other processes. Keep detailed notes about progress to allow recovery if this happens.

## Postrequisites

When finished, provide a table to the user showing the tests that have been run, their results, and any recordings that were created or updated.
