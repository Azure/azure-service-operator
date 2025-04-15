---
title: Customize the behaviour of generated resource
linktitle: Customize behaviour
weight: 40
---

{{% alert title="Note" %}}
In most cases, a code generated resource will work out of the box with no changes, and there will be no need to customize its behaviour by implementing an extensions.

Unless you have a specific need, you can skip this step and move on to [writing a test]({{< relref "write-a-test" >}}).
{{% /alert %}}

In an ideal world, the code generated resource will work out of the box with no changes. However, in practice, this is not always the case.

There are some variations in the way different product groups implement their resources.

To support these variations, ASO has extension points for resources that allow the behaviour to be modified.

For example, some resource providers will return a `BadRequest` error when a prerequisite is not met. Normally, ASO will treat `BadRequest` as a permanent failure, but this can be reclassified as a warning by using the [`ErrorClassifier`](https://github.com/Azure/azure-service-operator/blob/main/v2/pkg/genruntime/extensions/error_classifier.go#L19) extension point.

Other extension points allow blocking reconciliation attempts until a prerequisite is met, additional API calls after a successful reconciliation, or even modifying the resource before it is sent to Azure.

For most resource extensions, any implementation will go in the `customizations` folder generated for each resource group.

It's often unclear which extensions are needed until you start testing the resource. The best approach is to start with the generated resource, and then add extensions as needed.
Reach out to us if you want some guidance on which (if any) extension points are appropriate for your resource.

## Available extension points

| Extension pont            | Purpose                                                                                                        | Example                                                                                                                                                                                                                                                                                                |
| ------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| ARMResourceModifier       | Provides a hook allowing resources to modify the payload that will be sent to ARM just before it is sent.      | Used by KeyVault [Vaults](https://github.com/Azure/azure-service-operator/blob/a6ed2d1d093ed0f0404e709a6b7b969e12d6605d/v2/api/keyvault/customizations/vault_extensions.go#L44) to allow specifying create modes like `createOrRecover`.                                                               |
| Claimer                   | Customize how the reconciler claims a resource                                                                 |                                                                                                                                                                                                                                                                                                        |
| Deleter                   | Customize how the reconciler deletes a resource                                                                |                                                                                                                                                                                                                                                                                                        |
| ErrorClassifier           | Customize how the reconciler reacts to specific errors returned by Azure                                       | Used by DocumentDB [SqlRoleAssignment](https://github.com/Azure/azure-service-operator/blob/176fee93ea4aa80f38615047371d501a819d618b/v2/api/documentdb/customizations/sql_role_assignment_extension_types.go#L25) to retry assignments that fail because the referenced identity is not yet available. |
| Importer                  | An optional interface that can be implemented by resource extensions to customize the import process in asoctl |                                                                                                                                                                                                                                                                                                        |
| KubernetesSecretExporter  | Implemented when a resource needs to export a Kubernetes secret                                                | Used by EventHub [NamespacesEventhubsAuthorizationRule](https://github.com/Azure/azure-service-operator/blob/176fee93ea4aa80f38615047371d501a819d618b/v2/api/eventhub/customizations/namespaces_eventhubs_authorization_rule_extension.go#L36) to publish requested secrets to the cluster.            |
| PostReconciliationChecker | Implemented when a resource needs to perform additional checks after reconciliation against Azure              |                                                                                                                                                                                                                                                                                                        |
| PreReconciliationChecker  | implemented by resources that want to do extra checks before proceeding with a full ARM reconcile.             |                                                                                                                                                                                                                                                                                                        |
| SuccessfulCreationHandler | Implemented by resources that need to perform additional actions after a successful creation                   |                                                                                                                                                                                                                                                                                                        |

----

Whether or not you've implemented any extensions, our next step is to verify that the resource works as expected by [writing a test]({{< relref "write-a-test" >}}).
