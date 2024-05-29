---
title: Testing
---

## Running integration tests

Basic use: run `task controller:test-integration-envtest`.

### Required variables

| Environment Variable  | Value                                                                                                   | Needed for CI (Github Actions) | Needed for local testing                     |
|-----------------------|---------------------------------------------------------------------------------------------------------|--------------------------------|----------------------------------------------|
| AZURE_SUBSCRIPTION_ID | The Azure Subscription ID                                                                               | Yes                            | Yes (when recording)                         |
| AZURE_TENANT_ID       | The Azure Tenant ID                                                                                     | Yes                            | Yes (when recording)                         |
| TEST_BILLING_ID       | The Azure billing ID                                                                                    | No                             | Yes (when recording SubscriptionAlias tests) |
| CODECOV_TOKEN         | The token to https://app.codecov.io/gh/Azure/azure-service-operator                                     | Yes                            | No                                           |
| GH_PAT                | GitHub PAT, used for PR automation                                                                      | Yes                            | No                                           |
| REGISTRY_LOGIN        | The Azure Container Registry to log in to (for az acr login --name {name})                              | Yes                            | No                                           |
| REGISTRY_PRERELEASE   | The path to the container prerelease registry (right now this isn't used)                               | No                             | No                                           |
| REGISTRY_PUBLIC       | The path to the container release registry, used in --tag "{REGISTRY_PUBLIC}/{CONTROLLER_DOCKER_IMAGE}" | No                             | No                                           |


### Record/replay

The task `controller:test-integration-envtest` runs the tests in a record/replay mode by default, so that it does not 
touch any live Azure resources. (This uses the [go-vcr](https://github.com/dnaeon/go-vcr) library.) If you change the controller or other code in 
such a way that the required requests/responses from ARM change, you will need to update the recordings.

To do this, delete the recordings for the failing tests (under `{test-dir}/recordings/{test-name}.yaml`), and re-run 
`controller:test-integration-envtest`. If the test passes, a new recording will be saved, which you can commit to 
include with your change. All authentication and subscription information is removed from the recording.

To run the test and produce a new recording you will need to have set the required authentication environment variables 
`AZURE_SUBSCRIPTION_ID` and `AZURE_TENANT_ID`, _and_ logged in via `az login` (or you just use the `task` commands 
mentioned below and it will prompt you to `az login` if needed for that specific command). 
Note that you must be `Owner` on the subscription to execute some tests in record mode.

A few tests also need the `TEST_BILLING_ID` environment variable set to a valid Azure Billing ID when running in record mode. 
In replay mode this variable is never required. Note that the billing ID is redacted from all recording files so that 
the resulting file can be replayed by anybody, even somebody who does not know the Billing ID the test was recorded with.

Some Azure resources take longer to provision or delete than the default test timeout of 15m. To change the timeout, 
set `TIMEOUT` to a suitable value when running task. For example, to give your test a 60m timeout, use:

``` bash
TIMEOUT=60m task controller:test-integration-envtest
```

### Running live tests

If you want to skip all recordings and run all tests directly against live Azure resources, you can use the 
`controller:test-integration-envtest-live` task. This will also require you to set the authentication environment 
variables and `az login`, as detailed above.

### Running a single test
By default `task controller:test-integration-envtest` and its variants run all tests. This is often undesirable 
as you may just be working on a single feature or test. In order to run a subset of tests, use the `TEST_FILTER`:

```bash
TEST_FILTER=<test_name_regex> task controller:test-integration-envtest
```

or, with a timeout:

```bash
TIMEOUT=10m TEST_FILTER=<test_name_regex> task controller:test-integration-envtest
```
