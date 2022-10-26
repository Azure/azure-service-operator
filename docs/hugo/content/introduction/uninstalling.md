---
title: "Uninstalling"
---

## Before you uninstall

Before you uninstall Azure Service Operator, ensure that all of the resources managed by it are deleted from Kubernetes.

Use the [detach-on-delete reconcile-policy annotation](https://azure.github.io/azure-service-operator/introduction/annotations/#serviceoperatorazurecomreconcile-policy) 
if you want to delete an ASO resource from Kubernetes but retain it in Azure.

You can check if all ASO resources have been deleted by using: `kubectl api-resources -o name | grep azure.com | paste -sd "," - | xargs kubectl get -A`

## Uninstalling

We recommend that you uninstall ASO using the same mechanism you used to install it.

### Uninstalling with Helm

> **Warning**: This command will also remove installed ASO CRDs. As mentioned [above](#before-you-uninstall) it is strongly recommended that you
> ensure that there are no remaining ASO resources in the cluster before you run this command.

```bash
helm --namespace azureserviceoperator-system delete asov2
kubectl delete namespace azureserviceoperator-system
```

### Uninstalling with kubectl

> **Warning**: This command will also remove installed ASO CRDs. As mentioned [above](#before-you-uninstall) it is strongly recommended that you
> ensure that there are no remaining ASO resources in the cluster before you run this command.

```bash
kubectl delete -f https://github.com/Azure/azure-service-operator/releases/download/v2.0.0-beta.3/azureserviceoperator_v2.0.0-beta.3.yaml
```
