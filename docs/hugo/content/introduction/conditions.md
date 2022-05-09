---
title: Understanding Conditions
---
Each resource reports a `Ready` condition in `.status.conditions`. It is also visible when you examine
the resource with `kubectl get`, for example:

```
$ kubectl get resourcegroups.resources.azure.com 
NAME            READY     SEVERITY   REASON          MESSAGE
aso-sample-rg   False     Info       Reconciling     The resource is in the process of being reconciled by the operator   
```

If the condition's `status` is `True` (visible in the `READY` column via `kubectl get`) the resource is in
the goal state. No more reconciliations will be performed until a change is detected.

## Severity

When a conditions `status` is not `True`, it will include a `severity` and `message` detailing the problem (if any).
The possible severities are:

- **Info:** The resource is working as intended. It may be taking a while to get to the goal state.
  The `message` field describes what the operator is working on.
- **Warning:** There _may_ be a problem with the resource. The operator does not believe that
  this problem is fatal. Reconciliation will continue. Examine the `message` for more details.
- **Error:** There is a problem with the resource. The operator has given up reconciling this resource
  and requires you to make a change to correct the problem. See the `message` for specific details about
  the problem. The resource will stay in this state until user action is taken.
