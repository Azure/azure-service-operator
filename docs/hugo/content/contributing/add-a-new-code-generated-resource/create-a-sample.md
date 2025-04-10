---
title: Create a sample to demonstrate the resource
linktitle: Create a sample
weight: 50
---

We mandate that every resource has a sample that demonstrates how to use it - and enforce this with checks in our CI pipeline. Good examples help other ASO users understand how to use each resource. 

To ensure that each sample is correct, we run them - allowing us to catch any typos or errors before they trip up users in the wild.

The samples are located in the [samples directory](https://github.com/Azure/azure-service-operator/blob/main/v2/samples). There should be at least one sample for each kind of supported resource. These currently need to be added manually. It's possible in the future we will automatically generate samples similar to how we automatically generate CRDs and types, but that doesn't happen today.

The added new sample needs to be tested and recorded.

If a recording for the test already exists, delete it. Look in the [recordings directory](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/recordings/Test_Samples_CreationAndDeletion) for a file with the same name as your new test.  

Typically these are named `Test_<GROUP>_<VERSION_PREFIX>_CreationAndDeletion.yaml`.
For example, if we're adding sample for NetworkSecurityGroup resource, check for `Test_Network_v1beta_CreationAndDeletion.yaml`

Run the test and record it using `task`:

``` bash
TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```

Some Azure resources take longer to provision or delete than the default test timeout of 15m, so you may need to add the `TIMEOUT` environment variable to the command above. 

For example, to give your test a 60m timeout, use:

``` bash
TIMEOUT=60m TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```

Test replay runs at a faster rate than real Azure API calls, so the TIMEOUT is only needed for recording.

Once you have a successful recording, _test it_ by running the same command again. It should run quickly and pass. If that doesn't happen, you may need to delete the recording file again and re-record it. If problems persist, contact us for help.

