---
title: Diagnosing problems
---
## Common mistakes

### ASO controller pod in Status CreateContainerConfigError
```
$ kubectl get pods -n azureserviceoperator-system
NAME                                                       READY   STATUS                       RESTARTS   AGE
azureserviceoperator-controller-manager-69cbccd645-4s5wz   1/2     CreateContainerConfigError   0          7m49s
```

Very likely that you forgot to create the `aso-controller-settings` secret. This secret must be in
the same namespace as the pod. 

You can confirm this with `kubectl describe pod -n azureserviceoperator-system --selector control-plane=controller-manager`.
Look for the "Error: secret "aso-controller-settings" not found" event in the `describe` output:

```
Events:
  Type     Reason     Age                From               Message
  ----     ------     ----               ----               -------
  ...
  Warning  Failed     36s (x7 over 99s)  kubelet            Error: secret "aso-controller-settings" not found
  ...
```

## Problems with resources

### Resource with no Ready condition set
When resources are first created a ready condition should appear quickly, indicating that the resource is being reconciled:
```
$ kubectl get resourcegroups.resources.azure.com 
NAME            READY     SEVERITY   REASON          MESSAGE
aso-sample-rg   False     Info       Reconciling     The resource is in the process of being reconciled by the operator   
```
if this isn't happening then check the [controller logs](#getting-aso-controller-pod-logs).

### Resource stuck deleting 

This presents slightly differently for different resources, some examples are:

* [#2478](https://github.com/Azure/azure-service-operator/issues/2478)
* [#2586](https://github.com/Azure/azure-service-operator/issues/2586)
* [#2607](https://github.com/Azure/azure-service-operator/issues/2607)

For example, you might see something like this:
```
deleting resource "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dev-rg/providers/Microsoft.KeyVault/vaults/kvname/providers/Microsoft.Authorization/roleAssignments/kv-role-assignement3": DELETE https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dev-rg/providers/Microsoft.KeyVault/vaults/kvname/providers/Microsoft.Authorization/roleAssignments/kv-role-assignement3
--------------------------------------------------------------------------------
RESPONSE 400: 400 Bad Request
ERROR CODE: InvalidRoleAssignmentId
--------------------------------------------------------------------------------
{
    "error": {
        "code": "InvalidRoleAssignmentId",
        message": "The role assignment ID 'kv-role-assignement3' is not valid. The role assignment ID must be a GUID."
    }
}
--------------------------------------------------------------------------------
```

This can happen because the resource was created with an invalid name, and when ASO is trying to delete it,
it cannot delete the resource because the name is invalid.

_Usually_, ASO will prevent this situation from happening by blocking the original apply that attempts to create the resource, 
but from time to time that protection may be imperfect.

If you see this problem, the resource wasn't ever created successfully in Azure and so it is safe to instruct ASO to 
skip deletion of the Azure resource. This can be done by adding the `serviceoperator.azure.com/reconcile-policy: skip` 
annotation to the resource in your cluster.

## Getting ASO controller pod logs
The last stop when investigating most issues is to look at the ASO pod logs. We expect that
most resource issues can be resolved using the resources .status.conditions without resorting to 
pod logs. Setup issues on the other hand may requiring digging into the ASO controller pod logs.

Assuming that ASO is installed into the default namespace, you can look at the logs for the controller
with the following command: 
`kubectl logs -n azureserviceoperator-system --selector control-plane=controller-manager --container manager`

For example, here's the log from an ASO controller that was launched with some required CRDs missing:
> E0302 21:54:54.260693       1 deleg.go:144] setup "msg"="problem running manager" "error"="failed to wait for registry caches to sync: no matches for kind \"Registry\" in version \"containerregistry.azure.com/v1alpha1api20210901storage\""

### Log levels
To configure the log level for ASO, edit the `azureserviceoperator-controller-manager` deployment and 
set the `--v=2` parameter to a different log level.

See [levels.go](https://github.com/Azure/azure-service-operator/blob/main/v2/internal/logging/levels.go) for a list of levels that we log at. Other components used
by ASO (such as controller-runtime) may log at different levels, so there may be value in raising the log level
above the highest level we use. Be aware that this might produce a **lot** of logs.

### Logging aggregator
You may want to consider using a cluster wide logging aggregator such as [fluentd](https://www.fluentd.org/) (there are many other options).
Without one it may be difficult to diagnose past failures.

## Fetching controller runtime and ASO metrics

Follow the metrics [documentation](https://github.com/Azure/azure-service-operator/blob/main/docs/hugo/content/guide/metrics.md) for more information on how to fetch, configure and understand ASO and controller runtime metrics. 