---
name: diagnosing-vcr-failures
description: Use when ASO controller CRUD tests or sample tests fail during playback or recording, to diagnose the failure and re-record if needed.
---

# Diagnose and Fix go-vcr Test Failures

This skill covers diagnosing failures in ASO's go-vcr recording/replay tests (both controller CRUD and sample tests) and re-recording when needed.

For running tests and understanding the test suites, see the **testing-aso-recordings** skill.

## Quick Reference

|                    | Controllers                           | Samples                                                                |
| ------------------ | ------------------------------------- | ---------------------------------------------------------------------- |
| **Task command**   | `controller:test-controllers`         | `controller:test-samples`                                              |
| **Log file**       | `reports/test-controllers.log`        | `reports/test-samples.log`                                             |
| **Recordings dir** | `v2/internal/controllers/recordings/` | `v2/internal/testsamples/recordings/Test_Samples_CreationAndDeletion/` |

## Identifying Failures

Logs overflow the terminal buffer. Always use the log file (see table above).

- Search for `FAIL:` (trailing colon avoids false positives).
- Initial failures often cascade — identify the **first** problem.
- Compare the error message against the known failure modes below.

## Known Failure Modes

**IMPORTANT:** Before deleting a recording and re-recording, always investigate WHY the test failed. Blindly re-recording wastes time (30–60 min) and may not fix the problem. If you can't determine the root cause, report to the user and ask for guidance.

### Missing interaction

**Signature:** go-vcr error containing "no responses recorded for this method/URL".

The test made an HTTP request that doesn't exist in the recording.

- If recording is **unmodified** (`git status` shows clean) → the recording is stale. Delete it and re-record.
- If recording was **already updated** (you just re-recorded or it shows modified in git) → the test or controller is making an unexpected request. Investigate what changed; report to user.

### Body hash mismatch

**Signature:** go-vcr error containing "body mismatch" and/or logs showing "Request body hash header mismatch".

The test found the right URL in the recording but the request body differs.

- Recording **unmodified** → stale, re-record.
- Recording **just created** (you literally just recorded it) → non-deterministic serialization. The request body is being serialized differently across runs. This is a **systemic issue**. Do NOT keep re-recording — report to user.

### Test assertion failure

**Signature:** assertion error from test code (e.g. "Expected true to be false", "Expected ... to have length"), NOT a go-vcr error.

- **During recording** (live Azure) → likely a real bug in the controller or test code. Report to user.
- **During playback** → typically means the replay roundtripper is returning stale cached responses that don't reflect the current resource state. The recording may need re-creating, but if not, report the specific assertion and context to user.

### Timeout while recording

**Signature:** test fails with a context deadline or timeout during a live Azure recording run.

Some resources take 30+ min to provision (ApiManagement, Kusto, etc.) and the default test timeout may be insufficient.

- Extend timeout: add `TIMEOUT=90m` to the command line.
- If the test still fails with a longer timeout, investigate further (e.g., is the resource stuck provisioning?) and report to user.

### Lack of capacity or quota

**Signature:** `ServiceUnavailable` with a message about "high demand" in a region, or quota/capacity exceeded errors during recording.

1. Check the sample YAML for `location`/`locationName` fields.
2. Check what regions other samples in the same resource group use — prefer consistency.
3. Update the sample YAML to use a different Azure region (e.g. `eastus` → `australiaeast` or `westus2`).
4. Delete the failed/partial recording and re-run.
5. If the new region also fails, try another region and repeat.

## Re-Recording Workflow

### Step 1: Delete the stale recording

Delete the recording file for the failing test from the appropriate recordings directory (see Quick Reference table).

### Step 2: Record

Run as a **background terminal**:

**Controllers:**

```bash
source test.env && TIMEOUT=60m TEST_FILTER="<your-test-here>" task controller:test-controllers
```

**Samples** (note the `Test_Samples_CreationAndDeletion/` prefix is required):

```bash
source test.env && TIMEOUT=60m TEST_FILTER="Test_Samples_CreationAndDeletion/<your-test-here>" task controller:test-samples
```

Example: to re-record `Test_Redis_v1api20230801_CreationAndDeletion`:

```bash
source test.env && TIMEOUT=60m TEST_FILTER="Test_Samples_CreationAndDeletion/Test_Redis_v1api20230801_CreationAndDeletion" task controller:test-samples
```

For slow-provisioning resources, use `TIMEOUT=90m`. Record one test at a time to isolate problems.

### Step 3: Verify playback

Run the **same command without `source test.env`**:

```bash
TIMEOUT=60m TEST_FILTER="<your-test-here>" task controller:<SUITE>
```

Playback takes 1–3 min. If playback fails but recording succeeded, this is a systemic VCR issue (e.g. non-deterministic request serialization) — do NOT keep re-recording. Investigate and report.

### Step 4: Handle failed recording attempts

When a recording attempt fails, watch for:

1. **Azure debris** — failed runs can leave resources in Azure (resource groups, Entra security groups, etc.) that weren't cleaned up. On retry, the test may adopt these pre-existing resources instead of creating new ones, causing unexpected behavior. Wait a few minutes for Azure cleanup, delete the partial recording, and try again. Sometimes the quickest fix is to run a live test (which will delete everything in the resource group) and then re-record.

### Step 5: Run the full suite

Once updated recordings pass individually, **re-run the full test suite** to confirm all tests pass. A test run exits on first failure, leaving later tests unexecuted. The full re-run confirms all problems are resolved and there are no hidden issues.

If new failures appear, loop back to the top and diagnose them.

## Guidance

- **DO NOT interrupt a test run**, even if slow. Interrupting can leave Azure debris that pollutes the next recording attempt.
- **NEVER modify recording files by hand.** They are machine-generated and must only be created by running the tests.
- **DO NOT analyze recording YAML files** to debug mismatches. They're large and machine-generated — manual inspection is not productive. Delete and re-record instead.
- **ALWAYS use `task` directly** — never use `./hack/tools/task`. If `task` is not on the PATH, your environment is not set up correctly for testing. Stop and ask the user to fix your environment.
- **Monitoring:** Run as background terminals and use `await_terminal`. DO NOT poll logs in a loop. Timeouts:
  - Playback (single): 5 min (300000ms)
  - Recording (single): 10 min (600000ms), then repeat — can take 20–60+ min
  - If `await_terminal` times out, just call it again.

## Postrequisites

When finished, provide a table to the user showing which tests were re-recorded and why.
