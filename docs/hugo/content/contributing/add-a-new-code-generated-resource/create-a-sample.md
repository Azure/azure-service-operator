---
title: Create a sample to demonstrate the resource
linktitle: Create a sample
weight: 60
---

We mandate that every resource has a sample that demonstrates how to use it - and enforce this with checks in our CI pipeline. Good examples help other ASO users understand how to use each resource. 

To ensure that each sample is correct, we run them against a real cluster, and create the matching Azure resources - allowing us to catch any typos or errors before they trip up users in the wild.

The samples are located in the [samples directory](https://github.com/Azure/azure-service-operator/blob/main/v2/samples). There should be at least one sample for each kind of supported resource. These currently need to be added manually. It's possible in the future we will automatically generate samples similar to how we automatically generate CRDs and types, but that doesn't happen today.

If you're adding a new resource to an existing group, you may find it easier to add your new sample to an existing sample directory. 

Each new sample added needs to be tested, and a recording created. This is done in the same way as the recording we made for the resource test in the previous step.

If you need to re-record the test run for a sample (say, if you've changed the contents of the sample directory), you can do so by deleting the recording file(s) from the [recordings directory](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/recordings/Test_Samples_CreationAndDeletion).

Typically these are named `Test_<GROUP>_<VERSION_PREFIX>_CreationAndDeletion.yaml`.
For example, if we're adding sample for the `v1api20240101` version of `NetworkSecurityGroup` resource from the `network` group, check for `Test_Network_v1api20240101_CreationAndDeletion.yaml`

Run the test and record it using `task`:

```bash
TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```

This will run ALL the sample tests, and create recordings for any that are missing.


Some Azure resources take longer to provision or delete than the default test timeout of 15m, so you may need to add the `TIMEOUT` environment variable to the command above. 

For example, to give your test a 60m timeout, use:

```bash
TIMEOUT=60m TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```

Test replay runs at a faster rate than real Azure API calls, so the TIMEOUT is only needed for recording.

If you want to rerun just a single sample test, extend `TEST_FILTER` to include the name of the test you want to run. 

```bash
TEST_FILTER=Test_Samples_CreationAndDeletion/Test_Network_v1api20240101_CreationAndDeletion task controller:test-integration-envtest
```

Tip: If you're rerecording an existing sample, the name of the recording file is the name of the test (with `.yaml` appended). Copying that filename into the `TEST_FILTER` variable is an easy way to ensure you have the right name.


Once you have a successful recording, _test it_ by running the same command again. It should run quickly and pass. If that doesn't happen, you may need to delete the recording file again and re-record it. If problems persist, contact us for help.

---

With a sample created and tested, you're almost ready to submit your PR. Before you do, there are a few [final checks]({{< relref "final-checks" >}}) to make sure everything is working as expected.