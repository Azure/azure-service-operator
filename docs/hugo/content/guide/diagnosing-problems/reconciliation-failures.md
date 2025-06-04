---
title: Reconciliation Failures
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

Sometimes you are able to successfully create a resource in your Kubernetes cluster, but Azure Service Operator (ASO) experiences issues when trying to create that resource in Azure.

**How to use this TSG**: Review [_When to apply this TSG_](#when-to-apply-this-tsg) to determine whether this TSG is relevant. Each subsection recommends one or more actions to take, these are detailed below under [_Actions_](#actions).

# When to apply this TSG

This TSG applies when review of ASO logs shows:

* creation of the resource fails in Azure; or
* repeated reconcile failures followed by successful resource creation.

Each of these conditions is discussed in detail, below.

## Creation fails in Azure

The most usual case where resource creation fails in Azure is an error in the requested configuration of the resource.

These errors can be seen by retrieving the condition of the resource using `kubectl describe`:

```bash
$ kubectl describe resourcegroups/aso-sample-rg -n default
Name:         aso-sample-rg
Namespace:    default
Labels:       <none>
Annotations:  serviceoperator.azure.com/operator-namespace: azureserviceoperator-system
              serviceoperator.azure.com/resource-id: /subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/resourceGroups/aso-sample-rg
API Version:  resources.azure.com/v1beta20200601
Kind:         ResourceGroup
Metadata:
  ... elided ...
Spec:
  ... elided ...
Status:
  Conditions:
    Last Transition Time:  2023-08-31T01:25:58Z
    Observed Generation:   1
    Reason:                Succeeded
    Status:                True
    Type:                  Ready
  ... elided ...
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

Errors shown in this way come from Azure Resource Manager (ARM) and are not specific to ASO. They are the same errors you would see if you were using the Azure CLI or Azure Portal, and can be resolved using the same research techniques.

### Potential actions

* **Fix the problem**: Generic Azure errors are often self-explanatory, or they can be resolved by reference to documentation, or from an online search, allowing you to fix the issue yourself.

* [**Contact the Azure product group for support**](#contact-the-azure-product-group-for-support): If the cause is not obvious (e.g. the response code indicates a server error), consider contacting the appropriate product group directly for support.

* [**Verify the credentials used by ASO**](#verify-the-credentials-used-by-aso): If the error indicates an authentication issue, verify that ASO is using the expected identity and that the identity has the correct permissions.

## Repeated failures followed by success

Check ASO logs or events for each attempt to reconcile the resource and take note of the reasons for failure.

A sequence of near-identical failures terminated by an eventual success may indicate that creation of the resource was not possible until some precondition was satisfied.

In some cases (e.g. quota issues), the precondition is external and requires intervention (e.g. a quota request) before ASO can successfully create the resource. Based on the error message, you may be able to _fix the underlying cause_ yourself, or you may need to contact the relevant Azure product group for support.

In other cases, you might conclude that ASO is simply trying to create the resource too early. You may want to [_implement a new precondition check_](#implement-a-new-precondition-check) yourself, or [_request implementation of that check_](#request-implementation-of-a-new-precondition-check) from the ASO team by creating an [issue on GitHub](https://github.com/Azure/azure-service-operator/issues).

### Potential actions

* **Fix the underlying cause**: Many errors require you to take action (e.g. quota issues require a request for more quota). In these cases, the problem will automaticaly resolve once the underlying issue is resolved.

* [**Implement a new precondition check**](#implement-a-new-precondition-check): It's straightforward to modify ASO to respect a new precondition, delaying creation until that precondition is satisfied. ASO is an open source project that accepts contributions, so we encourage you to consider contributing changes yourself.

* [**Request implementation of a new precondition check**](#request-implementation-of-a-new-precondition-check): If contribution is not an option (for any reason), please create a [GitHub Issue](https://github.com/Azure/azure-service-operator/issues) requesting the ASO team implement the required check.

# Actions

This section identifies specific actions you can take to resolve issues. For discussion of which action should be applied, see above.

## Contact the Azure product group for support

For generic Azure errors, you should contact the Azure product group responsible for the resource type being used. Keep in mind that they won't be familiar with Azure Service Operator - some teams will bounce calls back to you if they don't recognize the client.

Ensure you provide all the following information:

* **Full error details** - from either the `condition` on the resource, or from ASO logs. Ensure you do not truncate any details.
* **Exact timestamp** - the instant, in UTC, of the error so they can review internal logs if required. This is available both on the resource condition, and in the ASO logs.
* **Resource ID** - the full ARM identifier of the resource. Usually provided by the customer, but available from the `serviceoperator.azure.com/resource-id` annotation on the resource if not.

It may also be helpful to provide some background on ASO, including the fact it works through the normal ARM REST API, just like any other client.

## Verify the credentials used by ASO

For authentication errors logged by ASO, it's possible that ASO is using the wrong identity, or that the identity does not have the correct permissions to create the resource.

Review the [ASO documentation on authentication](https://azure.github.io/azure-service-operator/guide/authentication/) and verify both that ASO is using the expected identity, and that the identity has the correct permissions.

Different kinds of resources require different permissions - for example, connecting a ManagedCluster to a ContainerRegistry requires **Owner** permissions, not simply **Contributor** permissions. It's therefore important to check whether the specific resource type requires any special permissions.

## Implement a new precondition check

When a new resource precondition is identified, the fastest route to resolution is for you to implement the change yourself. ASO is an open source project that accepts contributions, and we encourage you to consider contributing enhancements.

The ASO team is always available to advise on how to successfully implement new features, such as a precondition check.

Moreover, external PRs take priority over other work, so if you implement the change yourself, it is likely to be merged and released more quickly than if the ASO team implements it.

A good place to start is by creating a [GitHub issue](https://github.com/Azure/azure-service-operator/issues) to discuss the proposed change.

## Request implementation of a new precondition check

When a new resource precondition is identified, we can often implement the change on your behalf. This is a good option if you choose (for any reason) not to contribute a new precondition check yourself.

You should create a [GitHub issue](https://github.com/Azure/azure-service-operator/issues) requesting the change be made.

Please ensure the issue contains as much information as possible about the requested change, remembering that the ASO team may not be intimately familiar with the specific resource type or the precondition that should be checked.

New feature requests go into the ASO backlog, and are prioritized by the team alongside existing issues.

# Background

## Preconditions for resource creation

ASO already checks for a number of preconditions before attempting to create or update resources in Azure. For example, it will not attempt to create a resource if the owning resource does not exist, or if a required secret or configmap is missing.

For some resources, we've found it useful to introduce additional preconditions. For example, any attempt to PUT a `containerservice ManagedCluster` will automatically fail if the `provisioningState` of the `ManagedCluster` is in a transient state (e.g. `Creating` or `Updating`). To avoid polluting the ASO logs with a large number of failures, we block reconcile attempts until the cluster reaches a stable state (e.g. `Successful` or `Failed`).

Similar checks can be added for any resource, as needed, by implementing Go interface `PreReconcileCheck`. This interface provides an extension point, allowing resource behaviour to be customized where needed. See [PR#2686](https://github.com/Azure/azure-service-operator/pull/2686) for an example of this in action.
