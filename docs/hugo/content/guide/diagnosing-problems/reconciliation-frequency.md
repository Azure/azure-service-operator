---
title: Reconciliation Frequencey
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

## ASO is reconciling too often

ASO will continue actively reconciling resources, even after they have been created. This is by design, as ASO is a desired-state operator. It will continue to check the state of resources and update them if they are not in the desired state.

### Action: change reconciliation interval

By default, ASO will try to reconcile each resource once an hour once it has been successfully deployed.

If it's necessary to change ASO so that it reconciles less frequently, set the [AZURE_SYNC_PERIOD](https://azure.github.io/azure-service-operator/guide/aso-controller-settings-options/) environment variable to a longer value.

## ASO doesn't reconcile often enough

_a.k.a. ASO is too slow to update resources_

ASO works by polling Azure Resource Manager (ARM). Recognition of changes requiring might be delayed until the next scheduled poll. This sometimes results in observed delays in ASO acting to achieve the desired goal state of resources.

### Action: change reconciliation interval

By default, ASO will try to reconcile each resource once an hour once it has been successfully deployed.

If it's necessary to change ASO so that it reconciles more frequently,set the [AZURE_SYNC_PERIOD](https://azure.github.io/azure-service-operator/guide/aso-controller-settings-options/) environment variable.

> Warning: Polling more frequently runs the risk of triggering ARM throttling that affects the entire subscription, not just ASO. Setting the reconciliation interval lower than 15m is not recommended.

## ASO keeps reconciling even after the resource has been created

_a.k.a. ASO keeps doing PUTs when nothing has changed_

If you keep a close eye on your Azure Operation logs, you will see that ASO continues to interact with ARM after resources have been successfully created.

It's a subtle difference, but _ASO is not a deployment tool_. You don't hand ASO a manfiest of resources for it to create.

ASO is a desired state tool. You give it a manifest describing the resources you require. Any that are missing, will be created. Any that are present but with the wrong configuration will be updated to comply with the desired state. And, ASO will continue to maintain that desired state for as long as it is running.

The current design of ASO is that it performs period reconciliation by doing a PUT of the resource, relying on ARM to determine if the resource has changed. If you ware interestd in ASO adopting a more sophisticated approach, where the status of the resource is compared to the desired state and a PUT performed only if a change is necessary, follow [#2811](https://github.com/Azure/azure-service-operator/issues/2811).

### Action: stop ASO from reconciling

If you want to stop ASO from reconciling a particular resource, use the  [`serviceoperator.azure.com/reconcile-policy`](https://azure.github.io/azure-service-operator/guide/annotations/) annotation to instruct ASO to stop reconciling the resource.

> Note that this will not stop ASO from monitoring the resource. It will still GET the state of the resource and make that available within the cluster as the `status` field of the resource.
