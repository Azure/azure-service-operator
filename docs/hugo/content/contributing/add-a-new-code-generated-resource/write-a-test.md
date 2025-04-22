---
title: Write a test for the resource
linktitle: Write a test
weight: 50
---

Every new resource should have a handwritten test as there is always the possibility of a resource behaving in a subtly different way than expected. This also applies to _new versions_ of existing resources.

These handwritten tests should validate that the controller can create, update, and delete the resource. Often, they also validate that related resources can be created at the same time.
All these tests live in the [`v2/internal/controllers`](https://github.com/Azure/azure-service-operator/tree/main/v2/internal/controllers). There's a readme in that folder that describes the test structure and naming conventions.

For each test there's a [recording]({{< relref "testing">}}#recordreplay) in the `recordings` folder that records the HTTP interactions between the operator and Azure allowing for rapid replay for future runs. We create these recordings using [go-vcr](https://github.com/dnaeon/go-vcr).

One good approach is to start from an [existing test](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/controllers/documentdb_mongodb_crud_v20231115_test.go) and modify it to work for your resource. It can also be helpful to refer to existing ARM template examples for the resource in the [ARM templates GitHub repo](https://github.com/Azure/azure-quickstart-templates) to get an idea of how the resource is expected to behave.

## Record the test passing

When the test passes, you'll find a new recording file in the `recordings` folder. When you modify your test, you'll need to delete the recording file so that the test runs against the real Azure API.

Once you're finished, include the recording file in your PR. This is important because it allows other developers to run the test without needing to make real API calls to Azure.

See [the code generator test README]({{< relref "testing" >}}/#running-envtest-integration-tests) for more information on how to run tests and record their HTTP interactions to allow replay.

## Consider removing old tests

Each of these integration tests can take up to a few minutes to run from recordings, and we're starting to see issues due to the large number of resources we're testing.

If you are adding a new version for resource that already has _several_ supported versions, we _may_ request that you remove tests for older versions of the resource as a part of your PR.

As an absolute minimum, we want to have tests for

* the latest `stable` version of the resource;
* the prior `stable` version of the resource; and
* the latest `preview` version of the resource.

In most cases you won't need to worry about this.

### Prefer CreateResourcesAndWait

If your rest needs to create multiple resources, prefer using `CreateResourcesAndWait` (note the plural `Resources`) to create all the resources at once, rather than calling `CreateResourceAndWait` (singular `Resource`) once for each resource in turn.

When a user applies a YAML file containing multiple resources, ASO will be expected to handle creating all the resources in the correct order. Simluating this in the test by calling `CreateResourcesAndWait` to create all the resources at once is a good way to ensure that the test is realistic.

### Prefer parallel tests

Operation of ASO in a cluster is inherently concurrent, with a lot going on and an expectation that the state of the cluster will converge to the desired state.

It can be useful to break independent subtests up into parallel tests to simulate this concurrency. Check out the use of `RunParallelSubtests` in existing tests to see how this is done.

### Consider extensions

Odd behaviour in your test, this _may_ indicate you need to [implement an extension]({{< relref "implement-extensions" >}}) to customize the behaviour of the resource.

{{% alert title="Reach out" %}}
If you think this may be applicable, please reach out to us for help. We can help you identify the right extension point and how to implement it. In most cases, extensions aren't needed, but they are essential for some resources.
{{% /alert %}}

----

With a successful test demonstrating the resource working, it's time to [create a sample]({{< relref "create-a-sample" >}}) to show other users how to use the resource.
