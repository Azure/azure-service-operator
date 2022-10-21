---
title: FAQ
---
## Frequently Asked Questions

### Does ASO help with Disaster Recovery (DR) of resources in Azure?

No. If the Azure resource supports DR then you can configure it through ASO. 
If the underlying Azure Resource doesn't support DR (or the story is more complicated/manual), then you cannot currently configure it through ASO.

### How can I protect against accidentally deleting an important resource?

1. You can set [serviceoperator.azure.com/reconcile-policy: detach-on-delete](https://azure.github.io/azure-service-operator/introduction/annotations/#serviceoperatorazurecomreconcile-policy). This will allow the resource to be deleted in k8s but not delete the underlying resource in Azure.
2. You can use a project like https://github.com/petrkotas/k8s-object-lock to protect the resources you're worried about. Note: That project is not owned/sponsored by Microsoft.
3. You can manually add a finalizer to the resource which will not be removed except manually by you when ready to delete the resource, see [this](https://kubernetes.io/blog/2021/05/14/using-finalizers-to-control-deletion/)

There's also a proposal for [more general upstream support](https://github.com/kubernetes/kubernetes/issues/10179) on this topic, although there hasn't been movement on it in a while.

### What is the best practice for transferring ASO resources from one cluster to another?

There are two important tenets to remember when transferring resources between clusters:
1. Don't accidentally delete the resources in Azure during the transfer.
2. Don't have two instances of ASO fighting to reconcile the same resource to different states.

Let's say that you want to migrate all of your ASO resources from cluster A to cluster B. We recommend the following pattern:

1. Annotate the resources in cluster A with [serviceoperator.azure.com/reconcile-policy: skip](https://azure.github.io/azure-service-operator/introduction/annotations/#serviceoperatorazurecomreconcile-policy). This prevents ASO in that cluster from updating or deleting those resources.
2. Ensure that cluster B has ASO installed.
3. `kubectl apply` the resources into cluster B. We strongly recommend an infrastructure-as-code approach where you keep your original/goal-state ASO YAMLs around.
4. Delete the resources in cluster A. Note that because of the `skip` annotation, this will not delete the backing Azure resources.

### What is the best practice for transferring ASO resources from one namespace to another? 

See [above](#what-is-the-best-practice-for-transferring-aso-resources-from-one-cluster-to-another). The process is the same for moving between namespaces.

### Can I run ASO in active-active mode?

This is where two different ASO instances manage the same resource in Azure.

This _can_ be done but is not recommended. The main risk here is the goal state between the two instances of the same resource differing, causing thrashing in Azure
as each operator instance tries to drive to its goal. If you take great care to ensure that the goal state between the two clusters cannot differ, then
active-active can be done.

We instead recommend an active-passive approach, where in 1 cluster the resources are created/managed as normal, and in the other cluster the resources are just watched.
This can be accomplished with the [serviceoperator.azure.com/reconcile-policy: skip](https://azure.github.io/azure-service-operator/introduction/annotations/#serviceoperatorazurecomreconcile-policy) 
annotation used in the second cluster. In the case of a DR event, automation or manual action can remove the `skip` annotation in the passive cluster, turning it into active mode.

### Can ASO be used with IAC/GitOps tools?

Yes! We strongly recommend using something like [fluxcd](https://fluxcd.io/) or [argocd](https://argo-cd.readthedocs.io/en/stable/) with ASO.

If using argocd, make sure to **avoid** the `SyncPolicy Replace=true`, as that removes finalizers and annotations added by the operator whenever resources are re-applied.
ASO relies on the finalizer and annotations it adds being left alone to function properly. If they are unexpectedly removed the operator may not behave as expected.
