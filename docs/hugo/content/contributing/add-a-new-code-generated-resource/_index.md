---
title: Adding a new code generated resource to ASO v2
linktitle: Add a resource
---

Want to add a new code-generated resource to Azure Service Operator v2? You're in the right place - here's your step by step guide. We've broken the process down into a series of steps, and we'll walk you through each one.

[**Before you begin**]({{< relref "before-you-begin" >}}) sets the context for the rest of the process. You begin by identifying the resource you want to add and preparing your development environment.

[**Running the code generator**]({{ relref "running-the-code-generator" }}) details how to configure and run the code generator that writes 95% of the code for you. 

[**Review the generated resource**]({{< relref "review-the-generated-resource" >}}) gives you an overview of the generated code and how to identify common issues. 




----------


This document discusses how to add a new resource to the ASO v2 code generation configuration. Check out
[this PR](https://github.com/Azure/azure-service-operator/pull/2860) if you'd like to see what the end product looks like.




## Write a CRUD test for the resource

The best way to do this is to start from an [existing test](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/documentdb_mongodb_crud_v20231115_test.go) and modify it to work for your resource. It can also be helpful to refer to examples in the [ARM templates GitHub repo](https://github.com/Azure/azure-quickstart-templates).

Every new resource should have a handwritten test as there is always the possibility that the way a particular resource provider behaves will change with a new version. 

Given that we don't want to have to maintain tests for every version of every resource, and each additional test makes our CI test suite take 
longer, consider removing tests for older versions of resources when we add tests for newer versions. This is a judgment call, and we recommend 
discussion with the team first.

As an absolute minimum, we want to have tests for

* the latest `stable` version of the resource;
* the prior `stable` version of the resource; and
* the latest `preview` version of the resource.

These tests live in the [`v2/internal/controllers`](https://github.com/Azure/azure-service-operator/tree/main/v2/internal/controllers) folder and should follow the following naming convention:

``` 
<group>_<subject>_<scenario>_<version>_test.go
```

More information on the naming convention can be found in that folders [README](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/README.md).

### Record the test passing

See [the code generator test README](../testing/#running-envtest-integration-tests) for how to run tests and record their HTTP interactions to allow replay.

## Add a new sample

The samples are located in the [samples directory](https://github.com/Azure/azure-service-operator/blob/main/v2/samples). There should be at least one sample for each kind of supported resource. These currently need to be added manually. It's possible in the future we will automatically generate samples similar to how we automatically generate CRDs and types, but that doesn't happen today.

## Run test for added sample and commit the recording

The added new sample needs to be tested and recorded.

If a recording for the test already exists, delete it.  
Look in the [recordings directory](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/recordings/Test_Samples_CreationAndDeletion) for a file with the same name as your new test.  
Typically these are named `Test_<GROUP>_<VERSION_PREFIX>_CreationAndDeletion.yaml`.
For example, if we're adding sample for NetworkSecurityGroup resource, check for `Test_Network_v1beta_CreationAndDeletion.yaml`

Run the test and record it:

``` bash
TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```

Some Azure resources take longer to provision or delete than the default test timeout of 15m, so you may need to add the `TIMEOUT` environment variable to the command above. For example, to give your test a 60m timeout, use:

``` bash
TIMEOUT=60m TEST_FILTER=Test_Samples_CreationAndDeletion task controller:test-integration-envtest
```

## Send a PR

You're all done!
