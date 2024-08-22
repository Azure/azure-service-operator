---
title: "Best Practices"
linkTitle: "Best Practices"
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

## Managing multiple copies of the same resource

### Transferring resources from one cluster to another

There are two important tenets to remember when transferring resources between clusters:
1. Don't accidentally delete the resources in Azure during the transfer.
2. Don't have two instances of ASO fighting to reconcile the same resource to different states.

If you want to migrate all of your ASO resources from cluster A to cluster B, we recommend the following
pattern:

1. Annotate the resources in cluster A with
   [serviceoperator.azure.com/reconcile-policy: skip]( {{< relref "annotations#serviceoperatorazurecomreconcile-policy" >}} ).
   This prevents ASO in that cluster from updating or deleting those resources.
2. Ensure that cluster B has ASO installed.
3. `kubectl apply` the resources into cluster B. We strongly recommend an infrastructure-as-code approach where you
   keep your original/goal-state ASO YAMLs around.
    - Ensure these resources do not have the `serviceoperator.azure.com/reconcile-policy: skip` annotation set.
4. Delete the resources in cluster A. Note that because of the `skip` annotation, this will not delete the backing
   Azure resources.

### Transferring resources from one namespace to another

See [above](#transferring-resources-from-one-cluster-to-another). The process is
the same for moving between namespaces.

## Common cluster architectures

It's easiest to reason about ASO and what it is managing in Azure if you set up a simple mapping between
Kubernetes entities (clusters and namespaces) and the Azure resources being managed (subscriptions and resource groups).
These architectures are often used with GitOps tools such as Flux or ArgoCD.

See also:
- [Reducing access]({{< relref "reducing-access" >}}).
- [Security best practices]({{< relref "security" >}}).

### Cluster per environment

Environments like `dev`, `test`, and `prod` each have dedicated clusters with their own copy of ASO installed. ASO is
configured with a [global credential]({{< relref "credential-scope#global-scope" >}}) that has permissions to manage the
environment in question.

### Namespace per environment

Environments like `dev`, `test`, and `prod` each have dedicated namespaces within your Kubernetes cluster. Each namespace
has an `aso-credential` with permissions to manage the environment in question.

A variant on this is **namespace per developer**, where each developer gets their own `dev-alice` or `dev-bob` namespace,
rather than the whole team sharing a single `dev` namespace. Each `dev` namespace can point to a separate dev
subscription or share the same dev subscription. 

## AKS Best Practices

### CRDPattern

As described in the [CRD Management]({{< relref "crd-management" >}}) section, avoid installing a large number of CRDs 
into a Free tier AKS cluster.

### Affinity

If ASO is critical to your cluster/service operation, consider forcing the ASO pods to the `System` node pool which can
often be less CPU/memory loaded than the `User` pools. This can be done with the following modification to the Helm
`values.yaml` prior to installation:
```yaml
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
        - matchExpressions:
          - key: kubernetes.azure.com/mode
            operator: In
            values:
              - system
```
