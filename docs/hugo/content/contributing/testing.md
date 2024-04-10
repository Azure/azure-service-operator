---
title: Testing
---

## Running integration tests

Basic use: run `task controller:test-integration-envtest`.

### Record/replay

The task `controller:test-integration-envtest` runs the tests in a record/replay mode by default, so that it does not 
touch any live Azure resources. (This uses the [go-vcr](https://github.com/dnaeon/go-vcr) library.) If you change the controller or other code in 
such a way that the required requests/responses from ARM change, you will need to update the recordings.

To do this, delete the recordings for the failing tests (under `{test-dir}/recordings/{test-name}.yaml`), and re-run 
`controller:test-integration-envtest`. If the test passes, a new recording will be saved, which you can commit to 
include with your change. All authentication and subscription information is removed from the recording.

To run the test and produce a new recording you will need to have set the required authentication environment variables 
for an Azure Service Principal: `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`. 
This Service Principal will need access to the subscription to create and delete resources.

A few tests also need the `TEST_BILLING_ID` variable set to a valid Azure Billing ID when running in record mode. 
In replay mode this variable is never required. Note that the billing ID is redacted from all recording files so that 
the resulting file can be replayed by anybody, even somebody who does not know the Billing ID the test was recorded with.

Some Azure resources take longer to provision or delete than the default test timeout of 15m. To change the timeout, 
set `TIMEOUT` to a suitable value when running task. For example, to give your test a 60m timeout, use:

``` bash
TIMEOUT=60m task controller:test-integration-envtest
```

If you need to create a new Azure Service Principal, run the following commands:

```console
$ az login
… follow the instructions …
$ az account set --subscription {the subscription ID you would like to use}
Creating a role assignment under the scope of "/subscriptions/{subscription ID you chose}"
…
$ az ad sp create-for-rbac --role contributor --name {the name you would like to use}
{
  "appId": "…",
  "displayName": "{name you chose}",
  "name": "{name you chose}",
  "password": "…",
  "tenant": "…"
}
```
The output contains `appId` (`AZURE_CLIENT_ID`), `password` (`AZURE_CLIENT_SECRET`), and `tenant` (`AZURE_TENANT_ID`). 
Store these somewhere safe as the password cannot be viewed again, only reset. The Service Principal will be created as 
a “contributor” to your subscription which means it can create and delete resources, so 
**ensure you keep the secrets secure**.

### Running live tests

If you want to skip all recordings and run all tests directly against live Azure resources, you can use the 
`controller:test-integration-envtest-live` task. This will also require you to set the authentication environment 
variables, as detailed above.

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
