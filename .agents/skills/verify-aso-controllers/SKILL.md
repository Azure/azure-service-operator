---
name: verify-aso-controllers
description: Use when you need to verify that ASO CRUD tests of controllers are correct and functional.
---

# Verify ASO controller tests

When creating new resources, importing new resource versions, implementing or modifying resource extensions, or making other changes that change the sequence or content of HTTP requests made to Azure, we can need to replay and/or rerecord tests of our controllers.

## Workflow

### Run tests

Start by running all the tests:

```bash
./hack/tools/task controller:test-controllers
```

* If recordings for all tests exist, this will run in about 25m. Do not interrupt the test, just wait for it to complete.

* If any recordings are missing, each will result in a live test will automatically be run against Azure to create the missing recording. For most resources, this will take under 20m, but for some resources a recording may take 60m or more.

* If all the tests pass, then you're done and can stop working.
  
### Identify test failures

If the test run fails, look through the logs to identify which tests failed, and why. 

* Logs are likely to overflow the terminal buffer, so they get written to 'reports/test-controllers.log' as well. 
* Test failures can be identified by searching for "FAIL:" in the logs (the trailing colon is important to avoid false positives).
* For each tests, work out why the test failed. Note that initial failures tend to result in a cascade of following failures, so it's important to identify the first problem.

### Handling test failures

Here's a list of known failure modes and how to address them. Go through the newly identified test failures and make a TODO list of next actions based on this information.

**Missing interaction** - an error from go-vcr that the existing test recording is missing a required interaction. 

* If the test recording hasn't been updated (e.g. it does not show as modified to `git status`) then you should update the test by deleting the recording file and rerunning this specific test to create a new recording.
* If the test recording has already been updated, then you need to investigate further to identify the problem. Once you have a good idea of the problem, provide a detailed report to the user and ask for further instructions.

**Timeout while recording** - if a test fails with a timeout during recording, you may need to extend the time allowed for the recording by adjusting the `TIMEOUT` at the start of the test command. If the test still fails, investigate further to identify the problem and provide a detailed report to the user.

**Lack of capacity or quota** - if a test fails during recording with an error indicating that capacity or quota has been exceeded (e.g. `ServiceUnavailable` with a message about "high demand" in a region), take the following steps:

1. Look at the sample YAML that declares the resource (e.g. the `DatabaseAccount` sample) and note the `location`/`locationName` fields.
2. Check what regions other samples in the same resource group use — prefer consistency.
3. Update the sample YAML to use a different Azure region (e.g. change `eastus` to `australiaeast` or `westus2`).
4. Delete the failed/partial recording if one exists, and re-run the test to create a new recording.
5. If the new region also fails with a capacity error, try another region and repeat.

### Update a recording

Once you've identified a test that needs to be rerecorded, delete the recording file and run this command to create a new recording:

```bash
TIMEOUT=60m TEST_FILTER="<your-test-here>" ./hack/tools/task controller:test-controllers
```

Recordings should be updated one at a time to make it easier to identify problems. If you try to update multiple recordings at once, it will be harder to identify the source of any problems that arise.

After you have successfully updated a recording, run the test again to test that playback of that recording works correctly. The command to use is the same.

NEVER modify recording files by hand. They are machine-generated and must only be created by running the tests.

## Key information

Recordings of samples tests are found in `v2/internal/controllers/recordings/`, named for the test being executed.

The test source files are in `v2/internal/controllers/`. The test suite setup (envtest bootstrap, global test context) is in `v2/internal/controllers/suite_test.go`.

Test runs are CPU and time intensive, and are known to cause other processes to be memory or CPU starved. Be sure to keep detailed notes about your progress to allow you to pick up where you left off if this happens to you.

## Prerequisites

**Environment Variables:** You MUST have the following environment variables set to record the tests: `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and `ENTRA_APP_ID`. If any are missing, do not proceed because the tests will fail. Tell the user what is missing and stop the process.

## Guidance

DO NOT interrupt a test run, even if it looks like it's taking a long time. Interrupting a test run can cause problems with the recordings and make it harder to identify the source of any problems. (For example, you can end up leaving debris in Azure that will pollute the next recording attempt.)

When investigating test failures, be sure to look at the logs carefully to identify the specific error messages and any relevant context. This will help you to understand the problem and provide a detailed report to the user.

When updating a recording, be sure to only update one recording at a time to make it easier to identify any problems that arise. If you try to update multiple recordings at once, it will be harder to identify the source of any problems that arise.

DO NOT analyze recording YAML files to debug mismatches. Recording files are large and machine-generated — manual inspection is not productive. Instead, delete stale recordings and re-record.

DO NOT use `task` directly — always use `./hack/tools/task` to ensure the correct version is used.


## Postrequisites

When finished, provide a table to the user showing the tests that have been recorded, and why.
