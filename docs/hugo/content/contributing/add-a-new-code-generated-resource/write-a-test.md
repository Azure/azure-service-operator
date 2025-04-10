---
title: Write a test for the resource
linktitle: Write a test
weight: 40
---

Every new resource should have a handwritten test as there is always the possibility of a resource behaving in a subtly different way than expected. This also applies to _new versions_ of existing resources.

All these tests live in the [`v2/internal/controllers`](https://github.com/Azure/azure-service-operator/tree/main/v2/internal/controllers). There's a readme in that folder that describes the test structure and naming conventions.

For each test there's a recording in the `recordings` folder that records the HTTP interactions between the test and Azure allowing for rapid replay for future runs.

One good approach is to start from an [existing test](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/documentdb_mongodb_crud_v20231115_test.go) and modify it to work for your resource. It can also be helpful to refer to examples in the [ARM templates GitHub repo](https://github.com/Azure/azure-quickstart-templates).

## Resource extensions

There are some variations in the way different product groups implement their resources. To support these variations, there are a number of available extension points for resources that allow the behaviour to be modified. 

For example, some resources will return a `BadRequest` error when one prerequisite is not met. Normally, ASO will treat `BadRequest` as a permanent failure, but this can be reclassified as a warning by using the [`ErrorClassifier`](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/error_classifier.go#L19) extension point.

Other extension points allow blocking reconciliation attempts until a prerequisite is met, additional API calls after a successful reconciliation, or even modifying the resource before it is sent to Azure.

For most resource extensions, any implementation will go in the `customizations` folder generated for each resource group.

Reach out to us if you want some guidance on which (if any) extension points are appropriate for your resource.

## Record the test passing

When the test passes, you'll find a new recording file in the `recordings` folder. When you modify your test, you'll need to delete the recording file so that the test runs against the real Azure API. 

Once you're finished, include the recording file in your PR. This is important because it allows other developers to run the test without needing to make real API calls to Azure.

See [the code generator test README](../testing/#running-envtest-integration-tests) for more information on how to run tests and record their HTTP interactions to allow replay.

## Test volumes

These integration tests can take a few minutes to run from recordings, and we're starting to see issues due to the large number of resources we're testing. 

If you are adding a new version of an already supported resource, we may request that you remove tests for older versions of the resource as a part of your PR. 

As an absolute minimum, we want to have tests for

* the latest `stable` version of the resource;
* the prior `stable` version of the resource; and
* the latest `preview` version of the resource.

----

With a successful test demonstrating the resource working, it's time to [create a sample]({{< relref "create-a-sample" >}}) to show other users how to use the resource.
