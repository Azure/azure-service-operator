---
name: verify-aso-samples
description: Use when you need to verify that ASO samples are correct and functional.
---

# Verify ASO samples

When creating new samples, modifying existing samples, implementing resource extensions, or making other changes that change the sequence or content of HTTP requests made to Azure, we can need to replay and/or rerecord tests of our samples.

## Workflow

### Run tests

Start by running all the tests:

```bash
task controller:test-integration-envtest
```

* If recordings for all tests exist, this will run in about 25m. Do not interrupt the test, just wait for it to complete.

* If any recordings are missing, each will result in a live test will automatically be run against Azure to create the missing recording. For most resources, this will take under 20m, but for some resources a recording may take 60m or more.

* If all the tests pass, then you're done and can stop working.

Variation: if you've been asked to run tests only for a specific group of resources, add the TEST_FILTER environment variable on the command line:

```bash
TEST_FILTER="Test_Samples_CreationAndDeletion/Test_<test-group-here>.*" task controller:test-integration-envtest
```

### Identify test failures

If the test run fails, look through the logs to identify which tests failed, and why. 

* Logs are likely to overflow the terminal buffer, so they get written to 'reports/test-samples.log' as well. 
* Test failures can be identified by searching for "FAIL:" in the logs (the trailing colon is important to avoid false positives).
* For each tests, work out why the test failed. Note that initial failures tend to result in a cascade of following failures, so it's important to identify the first problem.

### Handling test failures

Here's a list of known failure modes and how to address them. Go through the newly identified test failures and make a TODO list of next actions based on this information.

**Missing interaction** - an error from go-vcr that the existing test recording is missing a required interaction. 

* If the test recording hasn't been updated (e.g. it does not show as modified to `git status`) then you should update the test by deleting the recording file and rerunning this specific test to create a new recording.
* If the test recording has already been updated, then you need to investigate further to identify the problem. Once you have a good idea of the problem, provide a detailed report to the user and ask for further instructions.

**Timeout while recording** - if a test fails with a timeout during recording, you may need to extend the time allowed for the recording by adjusting the `TIMEOUT` at the start of the test command. If the test still fails, investigate further to identify the problem and provide a detailed report to the user.

**Lack of capacity or quota** - if a test fails during recording with an error indicating that capacity or quota has been exceeded, you will need to investigate further to identify the specific resource and region that is affected. Once you have this information, provide a detailed report to the user and ask for further instructions.

### Update a recording

Once you've identified a test that needs to be rerecorded, delete the recording file and run this command to create a new recording:

```bash
TIMEOUT=60m TEST_FILTER="Test_Samples_CreationAndDeletion/<your-test-here>" task controller:test-integration-envtest
```

Recordings should be updated one at a time to make it easier to identify problems. If you try to update multiple recordings at once, it will be harder to identify the source of any problems that arise.

After you have successfully updated a recording, run the test again to test that playback of that recording works correctly. The command to use is the same.


## Key information

Samples are found in the folder `v2/samples` in groups by resource group.

Recordings of samples tests are found in `v2/internals/testsamples/recordings`, named by group and version.

Test runs are CPU and time intensive, and are known to cause other processes to be memory or CPU starved. Be sure to keep detailed notes about your progress to allow you to pick up where you left off if this happens to you.

## Prerequisites

**Environment Variables:** You MUST have the following environment variables set to record the tests: `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, and `ENTRA_APPLICATION_ID`. If any are missing, do not proceed because the tests will fail. Tell the user what is missing and stop the process.


## Guidance

DO NOT interrupt a test run, even if it looks like it's taking a long time. Interrupting a test run can cause problems with the recordings and make it harder to identify the source of any problems. (For example, you can end up leaving debris in Azure that will pollute the next recording attempt.)

When investigating test failures, be sure to look at the logs carefully to identify the specific error messages and any relevant context. This will help you to understand the problem and provide a detailed report to the user.

When updating a recording, be sure to only update one recording at a time to make it easier to identify any problems that arise. If you try to update multiple recordings at once, it will be harder to identify the source of any problems that arise.


## Postrequisites

When finished, provide a table to the user showing the tests that have been recorded, and why.
