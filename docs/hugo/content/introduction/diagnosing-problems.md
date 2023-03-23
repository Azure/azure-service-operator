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

Follow the metrics [documentation](https://github.com/Azure/azure-service-operator/blob/main/docs/hugo/content/introduction/metrics.md) for more information on how to fetch, configure and understand ASO and controller runtime metrics. 