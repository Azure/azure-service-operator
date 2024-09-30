---
title: FAQ
weight: -2
---
## Frequently Asked Questions

### What is the release cadence?

We ship updates to ASO as needed, with an eye towards releasing every 1-2 months. If there are urgent fixes, a release may happen
more quickly than that. If there haven't been any major changes (or there are ongoing major changes that are taking a long time) a 
release may happen more slowly. For an up-to-date plan check the 
[milestone tracker](https://github.com/Azure/azure-service-operator/milestones).

### How are CVEs dealt with?

The ASO controller container is built on [distroless](https://github.com/GoogleContainerTools/distroless), and as such
has a relatively minimal surface area. Most CVEs that impact ASO are related to the Golang packages used by 
the controller.

We scan for new CVEs weekly using [Trivy](https://github.com/aquasecurity/trivy) and also get proactive updates via 
[Dependabot](https://docs.github.com/code-security/dependabot/dependabot-alerts/about-dependabot-alerts).

CVEs are triaged according to their severity (Low, Moderate, High, Critical) and whether they are exploitable in ASO. 
Low and moderate severity issues will be fixed in the next minor release of ASO, high and critical severity CVEs that 
can be exploited in ASO will have a patch released for them.

Note that we cannot patch CVEs for which there is no upstream fix. Only once an upstream fix has been released will ASO
fix the CVE.

Fixes are _not_ backported to older versions of ASO. If you're running v2.5.0 and a CVE is fixed in v2.7.0, you must
upgrade to v2.7.0 to get the fix.

### What is the support model?

Azure Service Operator is an officially supported Microsoft OSS product. Currently, support is done via GitHub issues or
the `azure-service-operator` channel of the [Kubernetes Slack](https://kubernetes.slack.com/). There are plans to integrate ASO into
AKS as an addon after ASO has officially gone GA. At that point, support for ASO as an AKS addon would be accessed by raising an Azure 
support ticket.

### Does ASO help with Disaster Recovery (DR) of resources in Azure?

No. If the Azure resource supports DR then you can configure it through ASO. 
If the underlying Azure Resource doesn't support DR (or the story is more complicated/manual), then you cannot currently configure it through ASO.

### How can I protect against accidentally deleting an important resource?

1. You can set [serviceoperator.azure.com/reconcile-policy: detach-on-delete]( {{< relref "annotations#serviceoperatorazurecomreconcile-policy" >}}). This will allow the resource to be deleted in k8s but not delete the underlying resource in Azure.
2. You can use a project like https://github.com/petrkotas/k8s-object-lock to protect the resources you're worried about. Note: That project is not owned/sponsored by Microsoft.
3. You can manually add a finalizer to the resource which will not be removed except manually by you when ready to delete the resource, see [this](https://kubernetes.io/blog/2021/05/14/using-finalizers-to-control-deletion/)

There's also a proposal for [more general upstream support](https://github.com/kubernetes/kubernetes/issues/10179) on this topic, although there hasn't been movement on it in a while.

### What are some ASO best practices?

See [best practices]( {{< relref "best-practices" >}} ).

### Can I run ASO in active-active mode?

This is where two different ASO instances manage the same resource in Azure.

This _can_ be done but is not recommended. The main risk here is the goal state between the two instances of the same resource differing, causing thrashing in Azure
as each operator instance tries to drive to its goal. If you take great care to ensure that the goal state between the two clusters cannot differ, then
active-active can be done.

We instead recommend an active-passive approach, where in 1 cluster the resources are created/managed as normal, and in the other cluster the resources are just watched.
This can be accomplished with the [serviceoperator.azure.com/reconcile-policy: skip]( {{< relref "annotations#serviceoperatorazurecomreconcile-policy" >}} ) 
annotation used in the second cluster. In the case of a DR event, automation or manual action can remove the `skip` annotation in the passive cluster, turning it into active mode.

### Can ASO be used with IAC/GitOps tools?

Yes! We strongly recommend using something like [fluxcd](https://fluxcd.io/) or [argocd](https://argo-cd.readthedocs.io/en/stable/) with ASO.

If using argocd, make sure to **avoid** the `SyncPolicy Replace=true`, as that removes finalizers and annotations added by the operator whenever resources are re-applied.
ASO relies on the finalizer and annotations it adds being left alone to function properly. If they are unexpectedly removed the operator may not behave as expected.

### What's the difference between ASO and Crossplane.io?

There are a lot of similarities between ASO and Crossplane. They do similar things and have similar audiences. You can see some of this discussed [here](https://github.com/Azure/azure-service-operator/issues/1190).

**Today** primary differences are:
* ASO is officially maintained by Microsoft, while Crossplane Azure is community maintained.
* ASO focuses on simplicity. It doesn't offer any of the higher level abstractions that Crossplane does. ASO is not and will not ever be multi-cloud.
* The code generator we use to generate ASO resources is higher fidelity than the one that Crossplane uses. As a result, there are places where ASO resources are easier to use. 
  One example of this is references between resources such as linking a VMSS to a VNET. In Crossplane you do this by specifying the raw ARM ID. In ASO, you can specify the raw ARM ID but you can also specify a reference to the corresponding resource in Kubernetes (with its Kubernetes name) and ASO translates that into an ARM ID under the hood so that you don’t have to. This makes managing graphs of interlinked resources easier.
  
We would like to share our code-generator with Crossplane, as it’s higher fidelity than Terrajet (the codegenerator Crossplane uses to generate resources) for Azure resources. 
Right now our focus is on getting ASO to GA, after which we will hopefully have more time to invest in that.

### Can I configure how often ASO re-syncs to Azure when there have been no changes?

Yes, using the `azureSyncPeriod` argument in Helm's values.yaml, or using the `AZURE_SYNC_PERIOD`
in the `aso-controller-settings` secret. This value is a string with format like: `15m`, `1h`, or `24h`.

After changing this value, you must restart the `azureserviceoperator-controller-manager` pod in order for it to take effect
if the pod is already running.

Be careful setting this value too low as it can produce a lot of calls to Azure.

### I'm seeing Subscription throttling, what can I do?

ASO puts some steady load on your subscription due to re-reconciling resources periodically to ensure that
there is no drift from the desired goal state. The rate at which this syncing occurs is set by the 
[AZURE_SYNC_PERIOD]({{< relref "aso-controller-settings-options#azure_sync_period" >}}) (`azureSyncPeriod` in Helm)
Prior to ASO-beta.4, ASO's default `azureSyncPeriod` was 15m. It was changed to 1h in ASO-beta.4.

When `azureSyncPeriod` is up for a particular resource, a new PUT is issued to the resource RP to correct any drift from 
the goal state defined in ASO. There has been discussion about changing to do diffing locally to reduce requests to Azure, 
see [#1491](https://github.com/Azure/azure-service-operator/issues/1491).

You can estimate the maximum idle request rate of ASO based on the configured `azureSyncPeriod` and the number of 
resources being managed. The rough formula is: `numResources * 60/azureSyncPeriod(in minutes) = requestsPerHour`

For example:

| azureSyncPeriod | Number of resources | Requests / hour |
|-----------------|---------------------|-----------------|
| 15m             | 300                 | 1200            |
| 15m             | 1000                | 4000            |
| 1h              | 1200                | 1200            |
| 24h             | 28800               | 1200            |

### Why doesn't ASO support exporting/importing all data from configmap/secret?

For configuration management details, where the values are statically known prior to deployment, 
we strongly recommend people using existing tools such as Helm or Kustomize (or whatever other templating engine you love).
The secret and config map features in ASO are generally intended for dynamic values that aren't known prior to 
deployment, such as service endpoints and API keys generated by Azure.

### Why doesn't ASO have built-in templating?

The problem of how to template Kubernetes resources has already been solved a number of ways by projects such
as Kustomize and Helm (among others). ASO composes with those projects, and others like them, rather than trying
to build our own templating.

### Should I use user-assigned or system-assigned identity?

We don't take a position on whether it's universally better to deploy ASO using a user-assigned or system-assigned managed identity because the correct choice for you depends on your own context.

If you haven't already read it, Azure has a good [best practices for managed identity guide](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/managed-identity-best-practice-recommendations) that may be useful.

### When using Workload Identity, how can I easily inject the ASO created User Managed Identity details onto the service account?

The [workload identity documentation](https://azure.github.io/azure-workload-identity/docs/topics/service-account-labels-and-annotations.html#service-account)
suggests that you need to set the `azure.workload.identity/client-id` annotation on the ServiceAccount. 
This is not actually required! Setting that annotation instructs the Workload Identity webhook to inject the `AZURE_CLIENT_ID`
environment variable into the pods on which the ServiceAccount is used.

If you've created your user managed identity with ASO, it's easier to just do that injection yourself by using the 
`operatorSpec.configMaps` feature of the identity:

Identity:
```yaml
operatorSpec:
  configMaps:
    tenantId:
      name: identity-details
      key: tenantId
    clientId:
      name: identity-details
      key: clientId
```

and

Pod:
```yaml
env:
  - name: AZURE_CLIENT_ID
    valueFrom:
      configMapKeyRef:
        key: clientId
        name: identity-details
```

### How can I feed the output of one resource into a parameter for the next?

The answer changes a bit depending on the nature of the parameter.

For resource ownership: Set the `spec.owner.name` field of the dependent resource to be "owned by" the owning resource.
This will inform ASO that the owning resource needs to be deployed first.

For cross-resource-relationships: A resource referring to another resource will have a field like 
[diskEncryptionSetReference](https://azure.github.io/azure-service-operator/reference/containerservice/v1api20231001/#containerservice.azure.com/v1api20231001.ManagedCluster_Spec).
Set the reference to point to the resource you want:
```yaml
diskEncryptionSetReference:
  group: compute.azure.com
  kind: DiskEncryptionSet
  name: my-disk-encryption-set
```

For other fields: Not every field can be exported/imported. ASO does not have general purpose DAG support, nor does it have a
general-purpose templating language to describe such relationships. Instead, important properties can be imported/exported
from `ConfigMaps` or `Secrets`. See [setting UMI details on pods](#when-using-workload-identity-how-can-i-easily-inject-the-aso-created-user-managed-identity-details-onto-the-service-account)
above for one example of this. Another example can be found in 
[the authorization samples](https://github.com/Azure/azure-service-operator/blob/fe248787385af1b7b813e7bc2c8dbc595b8b4006/v2/samples/authorization/v1api20220401/v1api20220401_roleassignment.yaml#L12-L15), 
which reads the `principalId` from the `ConfigMap` written by 
[the associated identity](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/authorization/v1api20220401/refs/v1api20181130_userassignedidentity.yaml).

### What should I know about RoleAssignment naming in ASO?

The `RoleAssignment` resource is required to have a name in Azure which is a UUID. Since it's difficult to make a 
UUID for each assignment, ASO attempts to generate a unique UUID automatically if one isn't specified. ASO does this
by generating a UUID (v5) from a seed string with the following fields and setting it as the AzureName if one 
has not been specified:
1. name
2. namespace
3. owner.name
4. owner.namespace
5. owner.group
6. owner.kind

For the most part this should just work without you needing to worry about it. If you're running multiple ASO clusters
targeting different subscriptions in the same tenant you may need to be careful to avoid situations where
two resources (one in subscription A, one in subscription B) both have the same values for the fields called out above.
If this happens you may see an error like:

```
 (RoleAssignmentUpdateNotPermitted) Tenant ID, application ID, principal ID, and scope are not allowed to be updated.
 Code: RoleAssignmentUpdateNotPermitted
 Message: Tenant ID, application ID, principal ID, and scope are not allowed to be updated.
```

If you see this error, you can work around it by manually specifying a UUID for AzureName, or deleting the 
`RoleAssignment` which hit the error and creating it again with a different name.

If you're planning to move management of a `RoleAssignment` from one namespace to another (by setting the
reconcile-policy: skip on the old one, deleting it, and then creating the `RoleAssignment` in a different namespace 
allowing it to adopt the existing resource in Azure) you must manually specify the AzureName
of the `RoleAssignment` as the original UUID. Otherwise, the UUID defaulting algorithm will choose a different UUID since
the namespace has changed.

### How can I import existing Azure resources into ASO?

See [Annotations understood by the operator]({{< relref "annotations#serviceoperatorazurecomreconcile-policy" >}}) for
details about how to control whether the operator modifies Azure resources or just watches them.

There are a few options for importing resources into your cluster:
* If you're looking to import a large number of Azure resources you can use [asoctl]( {{< relref "tools/asoctl" >}}).
* If you're looking to import a small number of resources, you can also manually create the resources in your cluster 
  yourself and apply them. As long as the resource name, type and subscription are the same as the existing Azure 
  resource, ASO will automatically adopt the resource. Make sure to use the `reconcile-policy` you want.

### How do I get ArgoCD to show accurate health colours for ASO resources?

By default ArgoCD will show _green_ for all ASO resources regardless of their actual health. In [#4258](https://github.com/Azure/azure-service-operator/issues/4258), [@neil-199](https://github.com/neil-119) reported [this sample script](https://argo-cd.readthedocs.io/en/stable/operator-manual/health/#way-1-define-a-custom-health-check-in-argocd-cm-configmap) worked to show accurate health information based on the [Conditions]( {{< relref "conditions" >}}) already published by ASO.

``` yaml
data:
  resource.customizations: |
    "*.azure.com/*":
      health.lua: |
        hs = {}
        if obj.status ~= nil then
          if obj.status.conditions ~= nil then
            for i, condition in ipairs(obj.status.conditions) do
              if condition.type == "Ready" and condition.status == "False" then
                hs.status = "Degraded"
                hs.message = condition.message
                return hs
              end
              if condition.type == "Ready" and condition.status == "True" then
                hs.status = "Healthy"
                hs.message = condition.message
                return hs
              end
            end
          end
        end

        hs.status = "Progressing"
        hs.message = "Waiting for certificate"
        return hs
```
