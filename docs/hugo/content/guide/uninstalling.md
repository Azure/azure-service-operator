---
title: "Uninstalling"
weight: 10
---

## Before you uninstall

Before you uninstall Azure Service Operator, ensure that all of the resources managed by it are deleted from Kubernetes.

Use the [detach-on-delete reconcile-policy annotation]( {{< relref "annotations#serviceoperatorazurecomreconcile-policy" >}} )
if you want to delete an ASO resource from Kubernetes but retain it in Azure.

You can check if all ASO resources have been deleted:

``` bash
kubectl api-resources -o name | grep azure.com | paste -sd "," - | xargs kubectl get -A
``````

## Uninstalling

We recommend that you uninstall ASO using the same mechanism you used to install it.

### Uninstalling with Helm

As mentioned [above](#before-you-uninstall) it is strongly recommended that you
ensure that there are no remaining ASO resources in the cluster before you run this command.

``` bash
helm --namespace azureserviceoperator-system delete asov2
kubectl delete namespace azureserviceoperator-system
```

Note that Helm does not remove CRDs, due to the potential for user data-loss. The [Helm documentation](https://helm.sh/docs/chart_best_practices/custom_resource_definitions/) reads:

> There is no support at this time for upgrading or deleting CRDs using Helm. This was an explicit decision after much community discussion due to the danger for unintentional data loss.

-- Retrieved on Monday, 16 October 2023

### Uninstalling with kubectl

Use the same URL as you used during installation.

```bash
kubectl delete -f https://github.com/Azure/azure-service-operator/releases/download/v2.3.0/azureserviceoperator_v2.3.0.yaml
```
