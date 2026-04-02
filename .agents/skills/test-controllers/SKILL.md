---
name: test-controllers
description: Use when you need to verify that ASO CRUD tests of controllers are correct and functional.
---

# Verify ASO controller tests

When creating new resources, importing new resource versions, implementing or modifying resource extensions, or making other changes that change the sequence or content of HTTP requests made to Azure, we can need to replay and/or rerecord tests of our controllers.

## Workflow

### Preflight check

Before running tests:

1. **Check environment variables.** Verify that `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and `ENTRA_APP_ID` are set (or that the user has provided a `test.env` file to source). These are required for recording new tests against live Azure. If any are missing and recordings may need to be created, do not proceed — tell the user what is missing and stop. Note: environment variables are only needed for **recording** (creating new recordings against live Azure). Playback runs do not need them.

### Run tests

Start by running all the tests as a background terminal command:

```bash
source test.env && ./hack/tools/task controller:test-controllers
```

(If you're confident all recordings exist and no new ones will be created, you can omit `source test.env`.)

* Do NOT add `| tee` or other output redirection — the taskfile already handles piping output to `reports/test-controllers.log`.

Use `await_terminal` to wait for completion (see the **Monitoring long-running tests** guidance below for timeout recommendations).

* If recordings for all tests exist, playback typically completes in about 6-10 minutes. Do not interrupt the test, just wait for it to complete.

* If any recordings are missing, each will result in a live test automatically run against Azure to create the missing recording. For most resources, this will take under 20m, but for some resources a recording may take 60m or more.

* The test command may produce no visible output for several minutes while tests run — this is normal.

### Verify success

After the test command finishes with exit code 0, always confirm there are no failures in the log:

```bash
grep "FAIL:" reports/test-controllers.log || echo "No failures found"
```

* If no failures are found, you're done and can stop working.
* If failures are found despite a zero exit code, proceed to the "Identify test failures" section below.
  
### Identify test failures

If the test run fails, look through the logs to identify which tests failed, and why. 

* Logs are likely to overflow the terminal buffer, so they get written to 'reports/test-controllers.log' as well. 
* Test failures can be identified by searching for "FAIL:" in the logs (the trailing colon is important to avoid false positives).
* For each tests, work out why the test failed. Note that initial failures tend to result in a cascade of following failures, so it's important to identify the first problem.

### Handling test failures

Here's a list of known failure modes and how to address them. Go through the newly identified test failures and make a TODO list of next actions based on this information.

IMPORTANT: Before deleting a recording and re-recording, always investigate WHY the test failed. Blindly re-recording wastes time (some recordings take 30-60 minutes) and may not fix the problem. If you can't determine the root cause from the logs, provide a detailed report to the user and ask for guidance.

**Missing interaction** - an error from go-vcr that the existing test recording is missing a required interaction (message contains "no responses recorded for this method/URL"). 

* If the test recording hasn't been updated (e.g. it does not show as modified to `git status`) then you should update the test by deleting the recording file and rerunning this specific test to create a new recording.
* If the test recording has already been updated, then you need to investigate further to identify the problem. Once you have a good idea of the problem, provide a detailed report to the user and ask for further instructions.

**Body hash mismatch** - an error from go-vcr that the existing test recording has an interaction at the right URL but with a different body hash (message contains "body mismatch" and logs show "Request body hash header mismatch"). This is different from a missing interaction.

* This indicates the request body content has changed between when the recording was made and when it's being replayed.
* If the recording hasn't been updated (`git status` shows it clean), the recording is stale and needs to be re-recorded.
* If the recording WAS just created (e.g. you just re-recorded it), this indicates a non-deterministic serialization problem — the request body is being serialized differently across runs. This is a systemic issue. Do NOT keep re-recording; instead, investigate the VCR matching/hashing code and report to the user.

**Test assertion failure** - a test fails with an assertion error from the test code itself (e.g. "Expected true to be false", "Expected ... to have length"), NOT a go-vcr error.

* During recording (live Azure): this may indicate a real bug in the controller or test code. Report to the user.
* During playback: this typically means the replay roundtripper is returning stale cached responses that don't reflect the current resource state. The recording may need to be re-created, but first check if it's a known issue with the replay infrastructure. Report the specific assertion and context to the user.

**Timeout while recording** - if a test fails with a timeout during recording, you may need to extend the time allowed for the recording by adjusting the `TIMEOUT` at the start of the test command (e.g. `TIMEOUT=90m`). Some resources like ApiManagement services take 30+ minutes to provision, and the internal test polling timeout is 30 minutes — use `TIMEOUT=90m` for such tests. If the test still fails with a longer timeout, investigate further to identify the problem and provide a detailed report to the user.

**Lack of capacity or quota** - if a test fails during recording with an error indicating that capacity or quota has been exceeded (e.g. `ServiceUnavailable` with a message about "high demand" in a region), take the following steps:

1. Look at the sample YAML that declares the resource (e.g. the `DatabaseAccount` sample) and note the `location`/`locationName` fields.
2. Check what regions other samples in the same resource group use — prefer consistency.
3. Update the sample YAML to use a different Azure region (e.g. change `eastus` to `australiaeast` or `westus2`).
4. Delete the failed/partial recording if one exists, and re-run the test to create a new recording.
5. If the new region also fails with a capacity error, try another region and repeat.

### Update a recording

Once you've identified a test that needs to be rerecorded, delete the recording file and run this command as a **background terminal** to create a new recording:

```bash
source test.env && TIMEOUT=60m TEST_FILTER="<your-test-here>" ./hack/tools/task controller:test-controllers
```

For tests involving slow-provisioning resources (e.g. ApiManagement, Kusto), use `TIMEOUT=90m`.

Recordings should be updated one at a time to make it easier to identify problems.

After you have successfully updated a recording, run the test again (same command, **without** `source test.env`) to verify that playback works correctly. Playback is much faster than recording — typically 1-3 minutes. If playback fails but recording succeeded, this indicates a systemic issue with the VCR infrastructure (e.g. non-deterministic request serialization), NOT a stale recording. Do not keep re-recording — investigate and report to the user.

NEVER modify recording files by hand. They are machine-generated and must only be created by running the tests.

#### Handling failed recording attempts

When a recording attempt fails (timeout, assertion error, Azure quota), be aware of two things:

1. **Partial recording files**: A failed test may write a partial recording file. Always delete this file before retrying: `rm v2/internal/controllers/recordings/<TestName>.yaml`

2. **Azure debris**: A failed recording run can leave resources in Azure (resource groups, Entra security groups, etc.) that were created but not cleaned up. On the next recording attempt, the test may find these pre-existing resources and adopt them instead of creating new ones, causing the test to fail with unexpected behavior. If a re-recording fails unexpectedly (e.g. the test skips creation steps or finds resources it didn't create), wait a few minutes for Azure cleanup to complete, delete the partial recording, and try again.

### Confirmation

Once you've updated one or more recordings and the tests pass, start over from the top of this workflow and run all the tests again to confirm that everything is working correctly. If you find new failures, investigate and address them as described above until you have a clean test run with no failures.

Why do this? A tests run winds up as soon as one test fails, leaving later tests unexecuted. By starting over and running all the tests, you can confirm that all problems are fully resolved and that there are no other hidden issues.

## Key information

Recordings of samples tests are found in `v2/internal/controllers/recordings/`, named for the test being executed.

The test source files are in `v2/internal/controllers/`. The test suite setup (envtest bootstrap, global test context) is in `v2/internal/controllers/suite_test.go`.

Test runs are CPU and time intensive, and are known to cause other processes to be memory or CPU starved. Be sure to keep detailed notes about your progress to allow you to pick up where you left off if this happens to you.

## Guidance

DO NOT interrupt a test run, even if it looks like it's taking a long time. Interrupting a test run can cause problems with the recordings and make it harder to identify the source of any problems. (For example, you can end up leaving debris in Azure that will pollute the next recording attempt.)

**Monitoring long-running tests:** Always run test commands as background terminals and use `await_terminal` with appropriate timeouts to wait for completion. DO NOT poll the log file with `tail`, `grep`, or similar commands in a loop — this wastes tool calls and clutters the conversation. Recommended `await_terminal` timeouts:
- Playback-only (full suite): 10 minutes (600000ms) — typically completes in 6-10 minutes
- Playback-only (single test): 5 minutes (300000ms) — typically completes in 1-3 minutes
- Recording (single test): 10 minutes (600000ms), then repeat — recordings routinely take 20-30 minutes, some up to 60+ minutes
- Resource group deletion during recording: 15-25 minutes for groups with nested resources (this is the slowest phase)

If `await_terminal` times out, just call it again — the terminal keeps running.

When investigating test failures, look at the logs carefully to identify the specific error messages and any relevant context.

DO NOT analyze recording YAML files to debug mismatches. Recording files are large and machine-generated — manual inspection is not productive. Instead, delete stale recordings and re-record.

DO NOT use `task` directly — always use `./hack/tools/task` to ensure the correct version is used.


## Postrequisites

When finished, provide a table to the user showing the tests that have been recorded, and why.
