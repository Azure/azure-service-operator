---
title: Customize the behaviour of generated resource
linktitle: Customize behaviour
weight: 40
---

In an ideal world, the code generated resource will work out of the box with no changes. However, in practice, this is not always the case.

There are some variations in the way different product groups implement their resources. To support these variations, there are a number of available extension points for resources that allow the behaviour to be modified. 

For example, some resources will return a `BadRequest` error when one prerequisite is not met. Normally, ASO will treat `BadRequest` as a permanent failure, but this can be reclassified as a warning by using the [`ErrorClassifier`](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/error_classifier.go#L19) extension point.

Other extension points allow blocking reconciliation attempts until a prerequisite is met, additional API calls after a successful reconciliation, or even modifying the resource before it is sent to Azure.

For most resource extensions, any implementation will go in the `customizations` folder generated for each resource group.

It's often unclear which extensions are needed until you start testing the resource. The best approach is to start with the generated resource, and then add extensions as needed.
Reach out to us if you want some guidance on which (if any) extension points are appropriate for your resource.

----

Whether or not you've implemented any extensions, our next step is to verify that the resource works as expected by [writing a test]({{< relref "write-a-test" >}}).
