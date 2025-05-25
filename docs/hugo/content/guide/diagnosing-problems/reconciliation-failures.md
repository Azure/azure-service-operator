---
title: Reconciliation Failures
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

Sometimes you able to successfully create a resource in your Kubernetes cluster, but Azure Service Operator (ASO) experiences issues when trying to create that resource in Azure.

## When to apply this TSG

This TSG applies when review of ASO logs shows

* creation of the resource fails in Azure; or
* repeated reconcile failures followed by successful resource creation.

### Creation fails in Azure

The most usual case where resource creation fails in Azure is an error in the requested configuration of the resource.

These errors can be seen by retrieving the condition of the resource using `kubectl`:

```bash
kubectl TBC
```

or by finding details in the ASO logs. Log errors are always logged in full, and will look similar to this example:

```text
I2025-05-25T21:13:41Z] cache_redis "msg"="Error during Delete" "err"="deleting resource "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-fanmqx/providers/Microsoft.Cache/redis/asotest-redis2-vscggl": DELETE https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-fanmqx/providers/Microsoft.Cache/redis/asotest-redis2-vscggl
    --------------------------------------------------------------------------------
    RESPONSE 400: 400 Bad Request
    ERROR CODE: BadRequest
    --------------------------------------------------------------------------------
    {
      "error": {
        "code": "BadRequest",
        "message": "Cannot delete 'asotest-redis2-vscggl' since it has linked servers associated with it. Please remove the linked server(s) and then try deleting the cache.\r\nRequestID=8e55c6c4-ca68-477c-bcb8-db417aae80f3",
        "target": null
      }
    }
    --------------------------------------------------------------------------------
    " name="asotest-redis2-vscggl" namespace="aso-test-cache-redis-20230801-crud" azureName="asotest-redis2-vscggl" action="BeginDelete"
```

In most cases the error will be clearly shown and it will often be clear which correction is needed. 

If the error is not obvious (e.g. the response code indicates a server error), consider an online search for the error to see how others have handled it. Or, you can contacting Azure directly for support.

If you see an authentication issue (e.g. `401 Unauthorized` or `403 Forbidden`), this is likely an issue with Entra identity used by ASO - it may not have sufficient privileges to create the resource.

If you believe it's an ASO problem - where ASO is constructing the ARM payload incorrectly, please create a [GitHub issue](https://github.com/Azure/azure-service-operator/issues).

### Repeated failures followed by success

Check the ASO logs for each attempt to reconcile the resource. If there are repeated attempts to reconcile the resource that fail for the same reason, and those attempts are then followed by success, we may want to modify the behaviour of ASO by implementing the PreReconcileCheck interface.

For example, any attempt to PUT a `containerservice ManagedCluster` will automatically fail if the `provisioningState` of the `ManagedCluster` is in a transient state (e.g. `Creating` or `Updating`). To avoid polluting the ASO logs with a large number of failures, we have implemented the `PreReconcileCheck` interface for `ManagedCluster` resources, blocking reconcile attempts until the cluster reaches a stable state (e.g. `Successful` or `Failed`).

If there is a clear precondition that should block creation or updates of a specific resource, implementing the `PreReconcileCheck` interface can reduce noise in the logs and help ASO play nicely with others.

## Actions

These actions may be useful. See above for context on when each action is appropriate.

### Verify the credentials used by ASO

Review the [ASO documentation on authentication](https://azure.github.io/azure-service-operator/guide/authentication/) and verify both that ASO is using the expected identity, and that the identity has the correct permissions.

Different kinds of resources require different permissions - for example, connecting a ManagedCluster to a ContainerRegistry requires **Owner** permissions, not simply **Contributor** permissions. It's therefore important to check whether the specific resource type requires any special permissions.

### Request implementation of the PreReconcileCheck interface

Modify the behaviour of the resource by implementing the PreReconcileCheck interface to add any required preconditions that might block successful reconciliation.

As ASO is an open source project that accepts contributions, we'd encourage you to consider contributing this yourself.

If this is not an option (for any reason), create a [GitHub issue](https://github.com/Azure/azure-service-operator/issues) requesting the change be made.

If you don't want to create the issue yourself (e.g. because you don't want to be publicly identified), reach out via Azure support channels and ask them to contact the Azure Kubernetes Service (AKS) team team on your behalf.

### Contact Azure for support

If you need to contact Azure for support, keep in mind your initial contacts won't be familiar with Azure Service Operator and they may have some questions..

Ensure you provide all the following information:

* **Full error details** - from either the `condition` on the resource, or from ASO logs. Ensure you do not truncate any details.
* **Exact timestamp** - the instant, in UTC, so they can review internal logs if required. One place to get this is the ASO logs.
* **Resource ID** - the full ARM identifier of the resource. Available via the Azure Portal, or from the `serviceoperator.azure.com/resource-id` annotation on the resource.

It may also be helpful to provide some background on ASO, including the fact it works through the normal ARM API, just like any other client.

## Background

### Resource and property naming

Azure has strong standards for resource and property naming, but they differ slightly from the conventions used in the Kubernetes ecosystem. In order to work well with other Kubernetes tools, the names of resources and properties on ASO CRDs may differ subtly from the Azure originals.

**Resource names** are typically plural in Azure (e.g. Subscriptions, ManageClusters, or VirtualMachineScaleSets) but singular in ASO (Subscription, ManagedCluster, or VirtualMachineScaleSet)

**References** to other ARM resources often have the suffix `ID` in Azure, though this isn't 100% consistent. ASO requires the ability to refer to resources either within the cluster (via `Group`, `Kind`, and `Name`) or within Azure (via `ARM ID`), so these properties are renamed to have a consistent `Reference` suffix.

### The PreReconcileCheck interface

The generic reconciler used by ASO automatically applies some preconditions before trying to create a resource in Azure.

For example, any owning resource, must already exist. Ditto for any secrets or configmaps required to construct the resource payload.

The interface `PreReconcileCheck` provides an extension point, allowing this behaviour to be customized where needed, typically by checking for additional preconditions that need to be satisfied before the resource may be created or updated.
