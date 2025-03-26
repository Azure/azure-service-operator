---
title: ARM Throttling
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

It was once fairly easy for Azure Service Operator (ASO) users to trigger Azure Resource Manager (ARM) throttling, but this is now much less common because ARM have lifted their service limits considerably.

Under the old limits, a single subscription was limited to 600 PUT requests per hour in a given region. With ASO configured to reconcile every 15m, it was possible to hit the rate limit as soon as the number of resources being reconciled closed in on 150.

Starting in 2024, ARM is [migrating to a new throttling model](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling#migrating-to-regional-throttling-and-token-bucket-algorithm) where the limits are much higher. By default, writes (PUTs) are now limited to 10 per second, with bursts up to 200 requests permitted. (Exact limits differ by subscription.)

## How do I know if I'm at risk of throttling?

Look at the approximate number of resources you plan to have ASO manage and the frequency of reconciliation to estimate the number of PUT requests ASO will generate each hour. Compare that with the [documented limits](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling#migrating-to-regional-throttling-and-token-bucket-algorithm).

### Example 1

With 200 resources and a 60m reconciliation period, ASO will be making only 200 PUT requests per hour.
Given the default limit is 6000 writes per hour, you're very unlikely to experience throttling.

### Example 2

If you have 2000 resources and you want to reconcile every 5m, ASO will be making 24,000 requests per hour - well over the 6000 limit. Throttling is very likely.

### Example 3

If you have 3000 resources and a 30m reconciliation period, ASO will be making 6000 requests per hour. This is right at the limit, so you may experience limited throttling.

## How do I know if ASO is being throttled?

Throttling can be observed either by inspecting the condition of an affected resource, or by looking at the ASO logs.

### Inspecting the resource condition

Use `kubectl describe` to look at the Status.Conditions of the resource. If the resource has experienced throttling, the `Message` property will give you details.

``` bash
kubectl describe resourcegroup.resources.azure.com/aso-sample-rg
```

### Inspecting the ASO logs

Assuming that ASO is installed into the default namespace, you can look at the logs for the controller with the following command:

``` bash
kubectl logs --namespace azureserviceoperator-system --selector control-plane=controller-manager --container manager
```

## How do I avoid throttling?

### Reconfigure ASO

The most direct control for this is to increase the reconciliation period for ASO by setting [`AZURE_SYNC_PERIOD`](https://azure.github.io/azure-service-operator/guide/aso-controller-settings-options/) to a longer period.

The ASO default is 60m, chosen to balance the need for up-to-date information with the need to avoid throttling. If you have a large number of resources, you may want to increase this further. We have seen customers successfully using ASO with values as high as 24h.

### Use a different subscription

If you have the flexibility to do so, consider sharding your resources across multiple subscriptions.
