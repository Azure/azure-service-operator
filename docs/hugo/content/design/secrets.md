# Azure Service Operator support for managing data plane secrets

## What secrets are we talking about?

<!-- TODO: Are there secrets that aren't associated with a "data plane"? Maybe direct KeyVault secrets? -->
The secrets discussed in this document are associated with accessing the data plane of various services.

Think: Accessing a `StorageAccount` via Shared Key, accessing a `MySQLServer` by admin `Username` and `Password`, or accessing a VM
by `SSHKey`.

_Sometimes_ these secrets may be used by other CRDs managed by ASO, as would be the case for a `MySQLUser` CRD, but often the consumers of these
secrets are the users applications.

## Goals

### ASO should not be generating or managing secrets on the users behalf

ASO v1 generated secrets on the users behalf in some cases. The users had access to the generated secret as ASO wrote it into a Kubernetes (or KeyVault) 
secret.

There are problems with this approach:
1. It makes performing secret rollover more difficult. If the user had specified the secret, rollover is as easy as modifying the specified secret, which 
   triggers a reconcile, which triggers a PUT to the Azure Resource updating the secret. If the operator generated the secret then the user must somehow 
   issue an instruction to us to perform an action to roll the secret, which doesn't fit as well into the goal-seeking paradigm Kubernetes prefers.
2. It takes control away from the user. What if the secret the operator generates doesn't comply with a particular organizations complexity requirements?
3. It requires that we do a _very_ good job generating cryptographically secure passwords/secrets. This could easily become a can of worms.
4. It hinders adoption of existing resources. If the operator expects to always generate the secret for a SQL DB, but the user wants to import a SQL DB 
   they've already created through some other mechanism, then they are at an impasse as there is no (easy) way for them to provide ASO with the secret.
5. It doesn't work well with GitOps. A number of customers have expressed the desire to move resources between namespaces. In theory this is easy - 
   just create the exact same resources in a different namespace (pointing at the same Azure resources), mark the old resources as `skip-deletion` (so that
   the backing Azure resources are not deleted), and delete the original namespace. If the definition of the secret isn't part of the users GitOps flow this 
   becomes more difficult as the first namespace has  the secret (which was created by ASO) and that secret must be cloned manually to the second namespace,
   otherwise the creation of the resource in the second namespace will attempt to generate a new secret.

We could investigate allowing both automatically generated secrets (the default) and user specified secrets (opt-in). That avenue results in the operator 
shouldering all of the complexity burden of both though. 

Given ASO's approach as a low level toolkit providing the ability to create Azure resources, at least 
at this time we should avoid the added complexity of generating user secrets.

### Prefer Managed Identity

Managing secrets is hard. We should prefer AAD Authentication and Managed Identity where possible.

What this means in practice is that secrets should not be retrieved **unless the user has specifically asked for them to be**.

This has a number of advantages:
1. If the user _is_ using Managed Identity, they don't have a secret leak waiting to happen created in their namespace which they didn't ask about 
   (and might not even know about).
2. If the user _is_ using Managed Identity, many services support ways to instruct the `ListKeys` (or equivalent) API to error when called, as a way to
   force all callers to use an AAD Identity. If the operator is not calling those APIs unless instructed, we can make their failure fatal - the user must
   remove the bit of the `Spec` instructing us to call that API for their resource to reconcile successfully.

We'll dig more into how we might accomplish this a bit later.

### Work well with GitOps

Resources should strive to play well with GitOps. This means that when a new resource is deployed it is capable of adopting an already existing ARM resource
if one exists. This also must apply to secrets. As mentioned above one of the reasons that 
[we do not want to be generating secrets on the users behalf](#aso-should-not-be-generating-or-managing-secrets-on-the-users-behalf) is that it makes redeploy
and resource adoption hard because the user has no configuration representing the generated secret.

## Kinds of secrets

There are two main types of secrets we need to consider.

**User provided secrets**: Secrets provided by the user at resource creation time, and never returned by Azure.

**Azure generated secrets**: Secrets created by Azure, and returned in a special `GetMeTheSecrets` API call.

### Sample resources that have secrets

Below is a table containing a sampling of resources with secrets that ASO already supports or has a plan to support in the near future.

| CRD                        | User provided secrets | Azure generated secrets | AAD/Managed Identity Support | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| -------------------------- | --------------------- | ----------------------- | ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| VirtualMachineScaleSet     | ✔️                     | ❌                       | ❌                            | `Username` and `Password` (or SSH `PublicKey`). Can be modified by subsequent PUT. <!-- TODO: Validate this claim -->                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| VirtualMachine             | ✔️                     | ❌                       | ❌                            | `Username` and `Password` (or SSH `PublicKey`). Can be modified by subsequent PUT. <!-- TODO: Validate this claim -->                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| PostgreSQL FlexibleServer  | ✔️                     | ❌                       | ✔️                            | `AdministratorLogin` and `AdministratorLoginPassword`. Must have even if using AAD. <!-- TODO: Can this be modified by PUT? -->                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| MySQL FlexibleServer       | ✔️                     | ❌                       | ✔️                            | `AdministratorLogin` and `AdministratorLoginPassword`. Must have even if using AAD. <!-- TODO: Can this be modified by PUT? -->                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| StorageAccount             | ❌                     | ✔️                       | ✔️                            | [List Keys API](https://docs.microsoft.com/rest/api/storagerp/storage-accounts/list-keys) and [Regenerate Keys API](https://docs.microsoft.com/en-us/rest/api/storagerp/storage-accounts/regenerate-key). AAD+RBAC (blob/table only?) [Authorizing Access with Active Directory](https://docs.microsoft.com/en-us/azure/storage/blobs/authorize-access-azure-active-directory).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| CosmosDB DatabaseAccount   | ❌                     | ✔️                       | ✔️                            | [List Keys API](https://docs.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2021-04-15/database-accounts/list-keys), [List Read Only Keys](https://docs.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2021-04-15/database-accounts/list-read-only-keys) and [Regenerate Key API](https://docs.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2021-04-15/database-accounts/regenerate-key). For AAD+RBAC (supported by SQL only?), see [Disabling Local Auth](https://docs.microsoft.com/en-us/azure/cosmos-db/how-to-setup-rbac#disable-local-auth), [Create Role Assignment API](https://docs.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2021-04-15/sql-resources/create-update-sql-role-assignment), [Create Role Definition API](https://docs.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2021-04-15/sql-resources/create-update-sql-role-definition). [Built-in Role Definitions](https://docs.microsoft.com/en-us/azure/cosmos-db/how-to-setup-rbac#built-in-role-definitions). |
| EventHubAuthorizationRules | ❌                     | ✔️                       | ❌                            | [List Keys API](https://docs.microsoft.com/en-us/rest/api/eventhub/stable/authorization-rules-event-hubs/list-keys). There are default authorization rules created, such as `RootManageSharedAccessKey`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| Redis                      | ❌                     | ✔️                       | ❌                            | [List Keys API](https://docs.microsoft.com/en-us/rest/api/redis/redis/list-keys). [Regenerate Key API](https://docs.microsoft.com/en-us/rest/api/redis/redis/regenerate-key). <!-- AAD? -->                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |


<!-- TODO: ServiceBus namespace? -->

### Other kinds of secrets in Azure:

There are a few other types of secrets in Azure in addition to the two main ones discussed above.

1. Get once Azure generated secrets: These are secrets created by Azure and _only returned once_, usually as the response to a POST.
   These don't fit cleanly into the table above because they are a POST action on a parent resource these are not in fact resources themselves.
   1. Application Insights Component APIKey: This is a POST to the Component/ApiKey URL.
   2. KeyVault Key: [Create Key API](https://docs.microsoft.com/en-us/rest/api/keyvault/create-key/create-key).
2. "Secrets" created by Azure, and returned in the GET: These "secrets" are not really secrets per-se, but ASOv1 classifies them as secrets.
   1. `ApplicationInsights` `InstrumentationKey` or `ConnectionString`? 
      See [here](https://docs.microsoft.com/en-us/rest/api/application-insights/components/create-or-update#applicationinsightscomponent).
3. Short-lived "tokens".
   1. `StorageAccount` SAS
   2. `CosmosDB` ResourceToken.


## Other Operators

A quick look at what other operators are doing with regard to secrets.

### Crossplane

**Azure generated secrets**: A 
[Connection secret](https://doc.crds.dev/github.com/crossplane/provider-azure/database.azure.crossplane.io/CosmosDBAccount/v1alpha3@v0.17.0#spec-writeConnectionSecretToRef)
stores information needed to connect to the resource, including keys generated by Azure 
(see for example [Storage Account](https://github.com/crossplane/provider-azure/blob/faab5b58ea2cef6a2d1afbebb3c1a8943e72248e/pkg/controller/storage/account/account.go#L347)).
This is a [standard pattern](https://github.com/crossplane/crossplane-runtime/blob/6a7a44ac50aa1caca20a3cb5c215e6c03dc2b58e/pkg/resource/resource.go#L131)
used across all of the providers.

The destination of the `writeConnectionSecretToRef` is currently always a Kubernetes secret, but there is an [open issue](https://github.com/crossplane/crossplane/issues/2366) 
requesting pluggable secret stores.

**User provided secrets** seem to be 
[automatically generated](https://github.com/crossplane/provider-azure/blob/f37af6dd4f9d7d10a14caf8270d38078aee06bbe/pkg/controller/database/mysqlserver/managed.go#L157)
by Crossplane.

### ACK

**User provided secrets** are provided via a [SecretKeyRef](https://github.com/aws-controllers-k8s/runtime/blob/main/apis/core/v1alpha1/secret.go). 
This allows cross-namespace references to a secret. The specific `Key` of the secret is selected with a `Key string` parameter.

It looks like there may [not currently be support](https://github.com/aws-controllers-k8s/community/issues/700) for key updates/rollover.

**AWS generated secrets** I can't find any examples of. They do not seem to classify "endpoints" as a secret, as shown by
[opensearchservice endpoint](https://aws-controllers-k8s.github.io/community/reference/opensearchservice/v1alpha1/domain/) and
[cluster endpoint](https://aws-controllers-k8s.github.io/community/reference/eks/v1alpha1/cluster/).

## Proposal

### User specified secrets

User specified secrets will be detected and transformed from `string` to a `SecretRef`:

```go
type SecretRef struct {
    // SecretName is the name of the secret. The secret must be in the same namespace as the resource.
    SecretName string

    // Key is the key in the secret to use.
    Key string
}
```

Detection will be done with a combination of:
1. Additions to ASO's configuration file to flag particular fields as secret (probably in the `ObjectModelConfiguration` section).
2. Using the `"format": "password"` data from Swagger, such as that used by 
   [MySQL Flexible Server](https://github.com/Azure/azure-rest-api-specs/blob/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/mysql.json#L1662). 
   Note that not all specs have this, for example
   [VMSS does not](https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/compute.json#L9397).
3. Using the `x-ms-secret` annotation, which I assume has the same meaning as `"format": "password"` although it's not actually exactly documented anywhere.

This is a place where we can push changes upstream to flag things as passwords if they're not being flagged.

#### Lifecycle
Since these secrets were created _by the user_, the user owns the lifecycle of these secrets. They could be using the same secret across many resources, or 
intending to use this secret on a resource they have not created yet. As such, the lifecycle of these secrets must be controlled by the user.

Rollover will be supported by triggering events on the associated resource when the secret is modified. Since multiple custom resources might be using 
the same secret, this could trigger reconciles on multiple resources. An existing pattern has been established for this in 
[mysqlserver_controller.go](https://github.com/Azure/azure-service-operator/blob/main/controllers/mysqlserver_controller.go#L51) of ASO v1.

#### Open questions

**Should we support KeyVault inputs?**

I think we can get away without it at least initially... but maybe we need to fix up the shape of the `SecretRef` a bit to more easily support it in the future?

```yaml
  ...
  secret:
    kubernetes:
      name: foo
      key: bar
    keyVault:
      url: ??
```

Alternatively, we can keep this focused on Kubernetes secrets and if we want to expand to KeyVault secrets in the future just that _that_ on as a separate property
so that the base (Kubernetes) case is always clean:
```yaml
  secret:
    name: foo
    key: bar
    keyVault:
      url: ??
```

Kubernetes has a proposal out for [unions](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1027-api-unions). In this proposal they
suggest using a `discriminator` field, but allow for either of the shapes proposed above.

On the discriminator, they say:
> The value of the discriminator is going to be set automatically by the apiserver when a new field is changed in the union. It will be set to the value of the 
> fields-to-discriminateBy for that specific field.
> When the value of the discriminator is explicitly changed by the client, it will be interpreted as an intention to clear all the other fields. See section below.

See more details about what they see the `discriminator` doing in 
[normalizing on updates](https://github.com/kubernetes/enhancements/tree/master/keps/sig-api-machinery/1027-api-unions#normalizing-on-updates).

Since their proposal suggests that `discriminator` is optional, _and_ as far as I know it is not supported by `kubebuilder` yet, I suggest we don't add one
for now. In any case it's not totally clear to me that we really need the value that it is adding, and it seems to add a significant amount of update complexity.

We need to decide which of the above shapes we like more, and also if we do or do not want a `secretType` discriminator.

Conclusion: TBD

**Is an `AccountName` a secret?**

If the service is returning it in the resource `GET`, then strictly speaking it is not a secret. Since the field isn't secret, we will not transform it
to be a `SecretRef`. This is for two reasons:
1. We won't be able to automate it, since as far as the OpenAPI specification is concerned it isn't a secret.
2. The primary reason to classify this as a secret would be so that users could then inject the value from their secret into a pod. There is a workaround
   for this though since the user could (using Kustomize or something) do this already I think.

Conclusion: TBD

**Do we allow reading a secret from a namespace where the resource isn't?**

There were some requests for this in ASOv1, see [this](https://github.com/Azure/azure-service-operator/issues/1396) for example.

This has security implications, so initially at least the answer should be no. See: https://github.com/kubernetes/community/pull/5455.

Conclusion: TBD

**What's the difference between `x-ms-secret` and `format: password`?**

Unclear currently. I have a question out to the Swagger team. Current plan is to just treat them the same.

Conclusion: TBD

### Azure generated secrets

Azure generated secrets will be _optionally_ downloaded to a Kubernetes or KeyVault secret. Users instruct the operator to download
the secrets associated with a resource by supplying a `SecretDestination` in the `Spec`.

The optionality of this step is key for [preferring managed identity](#prefer-managed-identity), as it allows those in AAD/Managed identity cases to avoid
secrets they don't want/won't use being put into their namespace.

```go
type SecretDestination struct {
    // Secret includes details for where to store this secret in Kubernetes.
    Kubernetes *KubernetesSecretDestination  // TODO: We could call this "Secret" instead

    // KeyVault includes details for where to store this secret in KeyVault.
    KeyVault *KeyVaultSecretDestination
}

type KubernetesSecretDestination struct {
    // Name is the name of the secret.
    Name string

    // Annotations will be added to the created secrets annotations.
    Annotations map[string]string
}

type KeyVaultSecretDestination struct {
    // Reference refers to the KeyVault to write the secret to.
    Reference ResourceReference

    // Name is the name of the secret to create in KeyVault.
    Name string
}
```

Using this in the YAML would look like this, for an example CosmosDB `DatabaseAccount`:

```yaml
spec:
  # Other spec fields elided...
  forOperator:
    secrets:
      # Save the read-write keys (and endpoints) into a kubernetes secret called "my-secret" and a KeyVault secret called "my-secret".
      keyDestination:
        kubernetes:
          name: my-secret
        keyVault:
          reference:
            armId: /subscriptions/.../resourceGroups/.../providers/Microsoft.KeyVault/vaults/asokeyvault
          name: my-secret
      # Save the read only keys just to a Kubernetes secret called "my-readonly-secret"
      readOnlyKeyDestination:
        kubernetes:
          name: my-readonly-secret
```

Note that some resources (like CosmosDB `DatabaseAccount`) have multiple kinds of secrets. There might be a `PrimaryKey`, `SecondaryKey`, `ReadOnlyPrimaryKey`, 
and `ReadOnlySecondaryKey`. At the very least we need to support putting the main keys and the readonly keys into two different secrets. To accomplish this,
we can define multiple logical secret groupings in the ASO config and a corresponding `destination` property will be created for each of them. You can see this
done above in the sample.

Some fields such as `Endpoint` (or `AccountName` for other databases) are not _really_ secret, but should be included in these logical secret groups anyway so
that they're easier for users to inject into pods.

Azure generated secrets are more difficult to automatically detect. Often there is a `ListKeys` or `GetKey` API for the resource in question, but nothing on 
the resource itself indicates that it has secrets automatically generated by Azure. For resources like these we will add a flag in the ASO configuration to
generate the appropriate structures in the resource. In the future we can investigate detecting this from the Swagger (probably after we move to Swagger
as the single source of truth) by introspecting "other API calls" on the resource in question.

#### Hooks required
In addition to the configuration required to generate types with the right shape, we will also need a way to hook into the reconcile process and actually make
the right `ListKeys` or `GetKeys` call. To support rollover we would need a different hook as well.

This is a relatively involved topic so not designing it all here. As a starting point, resources manually implementing the following interface would
get us what we need. Issue [#1978](https://github.com/Azure/azure-service-operator/issues/1978) is tracking this request in more detail.

```go
type ARMDetails struct {
    Endpoint       string
    SubscriptionID string
    Creds          azcore.TokenCredential
    HttpClient     *http.Client
}

type ReconcileDetails struct {
    log                logr.Logger
    recorder           record.EventRecorder
    ARMDetails         *ARMDetails
    KubeClient         *kubeclient.Client
    ResourceResolver   *genruntime.Resolver
    PositiveConditions *conditions.PositiveConditionBuilder
}

type BeforeReconcileOverride interface {
    // BeforeCreateOrUpdate runs before sending the resource to Azure
    BeforeCreateOrUpdate(ctx context.Context, details *ReconcileDetails) (ctrl.Result, error)
}

type AfterReconcileOverride interface {
    // AfterCreateOrUpdate runs after the resource has been sent to Azure and finished creating, either successfully or with a fatal error.
    // You can determine the current state by examining the Ready condition on obj.
    AfterCreateOrUpdate(ctx context.Context, details *ReconcileDetails) (ctrl.Result, error)
}
```

Resources with Azure generated secrets would manually implement `AfterCreateOrUpdate`, which would then:
* Check that the resource is in a `Ready` state.
* Cast the provided `obj` to the expected type. This should be guaranteed safe because the method won't have been called unless its implemented.
* Check the `forOperator` section of the spec to determine if any keys should be written. This includes determining the destination secret name and any
  annotations or other properties that should be written or updated on that secret. Also includes ensuring that the secret in question (if it exists)
  is owned by the resource in question.
* Make the required `GetKeys` or `ListKeys` call to Azure.
* Create or update the secret.

#### Lifecycle
These are secrets created by the operator (usually after calling some `ListKeys` type API). Since the these secrets are by definition specific to the resource 
that created them, their ownership in Kubernetes will be set to the resource that created them. When the owning resource is deleted, the created secret will 
also be deleted.

#### Open questions

**Should `Endpoint`/`AccountName` type "secrets" really be put into the secret, or no?**

Putting them into the secret makes injecting them into pods easier, which is what people are going to want to do with these values. ASOv1 classifies all
of these sorts of things as secrets.

I think we should put these into the generated secrets... although this does somewhat conflict with my stance on user specified `accountName`, etc where
I had said to not classify them as secrets.

Conclusion: TBD

**How does key rollover work for these types of secrets?**

We don't support this at all initially. This is a somewhat advanced scenario that doesn't fit well into the Kubernetes resource model anyway
because rollover is more of an action and less of an actual resource. The user can roll their secrets using the `az cli` (or other tooling) and 
then either wait for a reconcile to occur naturally or force one to refresh the secrets locally in the cluster.

When we do decide to support this, we should be able to do so as a `Job`-esque resource that runs to completion (and somehow triggers a re-reconcile
on the parent resource type).

Conclusion: TBD

**How do deal with soft-delete and purge protection?**

If we support KeyVault, we have to deal with the soft-delete + purge awkwardness that deleting and recreating brings. Some of this we dodge by giving control 
to the user and expecting them to provide us with the name of a secret that's going to work. We may also need a flag per-secret (or global at the operator 
level?) for if we should purge the secret when we delete. Something like `purgeOnDelete`.

The main downside of a global flag controlling the behavior of purge or delete is that it doesn't give fine-grain control, yet we are proposing to give 
fine-grained control over what vault to put secrets in. This seems like a discrepancy. This seems to suggest we just put this option on every KeyVault reference:

```yaml
spec:
  # Other spec fields elided...
  forOperator:
    secrets:
      # Save the read-write keys (and endpoints) into a kubernetes secret called "my-secret" and a KeyVault secret called "my-secret".
      keyDestination:
        keyVault:
          reference:
            armId: /subscriptions/.../resourceGroups/.../providers/Microsoft.KeyVault/vaults/asokeyvault
          name: my-secret
          purge: false  # Optional: defaults to false
          delete: false # Optional: defaults to true
```

Conclusion: TBD

**Do we need to support "flatten" for KeyVault secrets like ASO v1 does?**

ASO v1 flatten determines if a secret is written as a single KeyVault secret containing a JSON blob with multiple entries, or as a bunch of individual secrets.
We don't have  this problem when writing Kubernetes secrets because they're somewhat hierarchical - you have a secret called `Foo` and you can refer to the 
the `Foo.mysecret` field within that secret when using it on a pod, etc.
KeyVault secrets don't have that structure. That might be OK though, as if users are using something like the 
[KeyVault secret store CSI driver](https://github.com/Azure/secrets-store-csi-driver-provider-azure) they could import the entire JSON blob and then extract 
the fields they need?

See the [alternate design](#alternative-design) that removes the need for any notion of flattening, which also resolves this.

Conclusion: TBD

**With the proposed design, there is no static understanding of the fields that will be present in the secret**

The fact that you will get 3 fields (`primaryKey`, `secondaryKey`, `endpoint`) in your CosmosDB `DatabaseAccount` secret is totally opaque to users.
Do we just need to document this? See proposal below for a possible alternative that solves this discovery problem.

Conclusion: TBD

**How do we ensure that we aren't overwriting secrets that we don't own?**

The operator must support updating the secret as keys can change or be rolled over. On the other hand we should present a clear error if the user has 
accidentally pointed two resources at the same secret. The first resource should be unaffected and work normally while the second resource should encounter
a reconcile error stating that the secret in question already exists.

This applies to either Kubernetes or KeyVault secrets.

Proposed solution for Kubernetes secrets is to issue a `GET` and check the `Owner` field for Kubernetes secrets before issuing an update. `resourceVersion`
should ensure that we're protected from any races where some external entity deletes the secret and another resource creates it between our `GET` and `PUT`.

Proposed solution for KeyVault secrets is to label the secret with a key uniquely identifying the operator and resource (GVK + namespace + name) which it
corresponds to. The operator will then issue a `GET` prior to attempting to update the secret to ensure that it owns the secret. Unfortunately, as far as I can
tell KeyVault secrets don't support `Etag` so there's a possible data race here that we can't avoid...

Conclusion: TBD

**What happens when a resource is updated to remove the `secrets` entry?**

If a resource is created and given a `forOperator.secrets` field that instructs it to create a secret, and then is updated to remove this secrets entry,
should we delete the secret we previously created? It feels like we _should_, but when the operator is presented with the updated object during that second
reconcile it doesn't know that it ever had a secret.

Some possible solutions:

1. We just don't delete the secrets when you do this and expect users to do it themselves if they want things cleaned up.
2. For k8s secrets, we could probably look through all of the secrets in the resources namespace and see if any of them are owned by us. If there are some
   and we don't have any `forOperator.secrets` we delete them. The downside here is that's a lot of overhead to do for all reconciles. This also assumes that
   the operator has Secret `Read` permissions in the namespaces in question - which it's possible that it doesn't (it may never have created any secrets and
   the users may know that it won't and so have denied it permission). This also obviously doesn't work for KeyVault secrets as we have no idea which KeyVault
   the secrets might even be in.
3. When we create a secret we store information about it in `Status` or `Annotations` and then use that information to clean up after ourselves. It's a bit 
   cleaner to store it in `Status` but technically that can be lost, whereas `Annotations` won't be. This should work for both KeyVault and k8s secrets and 
   while it's a bit icky I think gives the best experience.

Conclusion: TBD

**Should we drop the `Destination` suffix in the field names?**

For example, the proposal for CosmosDB `DatabaseAccount` was:
```yaml
spec:
  # Other spec fields elided...
  forOperator:
    secrets:
      keyDestination:
        # stuff elided...
      readOnlyKeyDestination:
        # stuff elided...
```

```yaml
spec:
  # Other spec fields elided...
  forOperator:
    secrets:
      key:
        # stuff elided...
      readOnlyKey:
        # stuff elided...
```

Conclusion: TBD

**Do we allow writing a secret to a namespace where the resource isn't?**

No, see: https://github.com/kubernetes/community/pull/5455

Conclusion: TBD

**Do we support writing the same secret to multiple destinations?**

If/when we add support for writing secrets to KeyVault rather than to Kubernetes secrets, would we also enable support for writing the same secret to _both_
KeyVault and Kubernetes? This would allow the user to specify a `spec` like:
```yaml
      # Save the read-write keys (and endpoints) into a kubernetes secret called "my-secret" and a KeyVault secret called "my-secret".
      keyDestination:
        kubernetes:
          name: my-secret
        keyVault:
          reference:
            armId: /subscriptions/.../resourceGroups/.../providers/Microsoft.KeyVault/vaults/asokeyvault
          name: my-secret
```

I think we should.

Conclusion: TBD

**What client should we use for issuing the ListKeys/GetKeys request?**

We don't automatically generate methods for performing these calls. It seems easiest to just use the corresponding Azure SDK for this.
This does mean that we're taking a dependency on the SDK where we didn't have one before, but the overall footprint of usage is pretty small
and it's going to be easier to do that than it is to write the methods ourselves.

Conclusion: TBD

**What API version should we use for issuing the ListKeys/GetKeys request?**

There are a few options here. The custom hooks could be applied to either the customer facing resource version _or_ to the internal storage version.
Applying the hooks to the storage version is easier, but means that the same hook will be run regardless of which customer 

#### Alternative design

Would we prefer an alternative design where each individual field that goes into a secret is called out explicitly? It's noisier but more flexible and also 
probably more discoverable, at least in theory? Here's what it might look like for CosmosDB `DatabaseAccount`

```yaml
  # Other spec fields elided...
  forOperator:
    secrets:
      primaryKeyDestination:
        kubernetes:
            name: my-secret
        keyVault:
          reference:
          armId: /subscriptions/.../resourceGroups/.../providers/Microsoft.KeyVault/vaults/asokeyvault
          name: my-primary-key
      secondaryKeyDestination:
        kubernetes:
            name: my-secret
        keyVault:
          reference:
          armId: /subscriptions/.../resourceGroups/.../providers/Microsoft.KeyVault/vaults/asokeyvault
          name: my-secondary-key
      readOnlyPrimaryKeyDestination:
        kubernetes:
          name: my-readonly-secret
      readOnlySecondaryKeyDestination:
        kubernetes:
          name: my-readonly-secret
      endpointDestination:
        kubernetes:
          name: my-readonly-secret
      # Save the read only keys just to a Kubernetes secret called "my-readonly-secret"
      readOnlyKeyDestination:
        kubernetes:
          name: my-readonly-secret
```

**Note:** the above is probably the "worst case" in terms of complexity. Most resources have fewer kinds of secrets and usually you wouldn't be writing 
everything to two different places.

This has some advantages:
1. It's more obvious what fields you're going to end up with in your secret.
2. It allows writing different parts of the secret to different secrets/locations - which doesn't seem like it's all that interesting?
3. It allows you to avoid writing certain parts of the secret, while still writing others. This is in contrast to the original proposals prebuilt groupings
   that cannot be changed.
4. It "solves" the KeyVault flattening question, as everything is always flattened in KeyVault and you have to name the secrets individually. 
   This makes for easier consumption. The resulting key names are more obvious too. Supporting flattening in the other design would require us to generate
   the names (ASOv1 does "<user-provided-name>+<secret-name>", which works pretty well but again isn't all that discoverable).

There are also some disadvantages:
1. Allowing multiple secrets means that if we support some additional details such as annotations as 
   [requested by customers](https://github.com/Azure/azure-service-operator/issues/1398), the common use case may result in customers duplicating the same
   annotations (although possibly there would be no difference if they didn't, provided they were writing to the same secret for all keys).

Personal note: I _think_ I like this one more.

Conclusion: TBD

### Other kinds of secrets

See [other kinds of secrets in Azure](#other-kinds-of-secrets-in-azure) for examples of each of these types of secrets.

**Get once Azure generated secrets**: These are secrets created by Azure and _only returned once_, usually as the response to a POST.

This pattern doesn't seem to be very common. The current plan is to not support resources with this pattern. If we did need to support
resources with this pattern there would be no way to avoid violating the [work well with GitOps](#work-well-with-gitops) goal, as when
deploying this resource there would be no way to adopt secrets associated with it (they're GET-once). That might be ok for particular
kinds of resources provided we document it though. 

If there are large user requests for this sort of secret management we can support it similar to the `ListKeys` cases except that 
the resources won't be movable without also cloning the secret.

**"Secrets" created by Azure, and returned in the `GET`**: This is things like `InstrumentationKey`, or server endpoint URLs.

Since these are returned in a `GET` they are not secret and are already being shown on Status. We will also manually classify some of them to be 
included in secrets we write. That will include things like `endpoint`. This means that `endpoint` for a SQL Server would show up in two places,
in the status _and_ possibly in the secret written by the operator (assuming that the user has instructed us to write the endpoint someplace).

The main reason for writing this into two places is:
1. We already have it in the `Status` and it doesn't make sense to remove it as it's part of the resource payload from Azure.
2. We want the ability to put it into a Secret so that if users want to inject it into their pods as an environment variable it's easy to do so.
   There is no good way to inject environment variables from CRD Status's.

**Short-lived "tokens"**: Like Storage SAS.

We will not support these sorts of secrets.


### A special note on KeyVault
We have an [open issue](https://github.com/Azure/azure-service-operator/issues/1894) asking for support to manage KeyVault secrets via ASO.

We have to be _very_ careful about support for KeyVault secrets (or certs, keys, etc), as the whole point of KeyVault is as a secure place to store your 
secrets.
If you're instead creating those secrets via the operator then you have by definition also located the secret in Kubernetes, which somewhat defeats the purpose 
from a security perspective. 

There might be cases where this makes sense if there are APIs that require a KeyVault secret to be provided to them via ARM ID and so KeyVault isn't the 
medium of secure storage so much as it is a medium of secret transfer.

Until we run into such scenarios **we should avoid implementing any of the KeyVault key/secret/certificate/etc APIs**.

## Integrations

If/when we support storing secrets in KeyVault, we can create some demos showing integration with the 
[KeyVault secret store csi driver](https://github.com/Azure/secrets-store-csi-driver-provider-azure).

## Supported secret stores

P0: Kubernetes Secrets

P1: KeyVault - this seems more interesting for the "secrets from Azure" case, since at least right now you cannot create KeyVault secrets through
ASO. That limitation means if you wanted to use KeyVault for input secrets you must have already pre-created the secrets before deploying via the operator.

## Implementation plan

There are effectively two parallel features here:
1. Input secrets (reading secrets from a store and supplying those secrets to Azure)
2. Output secrets (reading secrets from Azure and storing them in a secret store)

### Input secrets

1. `SecretRef` implementation.
2. Code generator changes to detect `format: password` or `x-ms-secret` and transform properties appropriately.
3. Reflector library to generically crawl resources and find secret refs.
4. `SecretReader` library to get secrets from a collection of `SecretRef`'s. This should be expandable to support KeyVault secret refs in 
   the future if we decide to add them. Should probably look something like this. Note that this  related to the `SecretWriter` below in
   the [output secrets](#output-secrets) section:
    ```go
    // TODO: these strings may need to be []byte
    // Note: these interfaces are acting on collections only so that they can be more efficient, as multiple SecretRef's or secret values may be read from or written to 
    // a single secret.

    type SecretReader interface {
        GetSecrets(ctx context.Context, refs []SecretRef) (map[SecretRef]string, error)
    }
    ```
5. ARM conversion changes to take a `ResolvedSecrets` parameter in addition to the `ResolvedReferences` it takes now.
   This will probably require a bit of fussing with the types passed to the ARM conversion methods and the conversion methods themselves.
6. Reconciler changes to perform secret reading in addition to reference resolving.
7. Testing: Modify existing tests passing secrets in plain text as part of the spec to use the new mechanism instead.
8. [Stretch goal] Support for secret rollover:
   1. Hook to add field indexers for specific resources (see what we're doing for MySQL in ASOv1 today).
   2. Hook to control custom additional watches used for monitoring changes to secrets (see what we're doing for MySQL in ASOv1 today).

### Output secrets

1. `operatorSpec`/`operatorStatus` preparatory work. See [#1612](https://github.com/Azure/azure-service-operator/issues/1612).
2. Add hooks to controller allowing handcrafted per-resource customization, see [#1978](https://github.com/Azure/azure-service-operator/issues/1978).
3. Update azure-arm configuration to allow for additional AzureGeneratedSecret properties to be defined. These properties will be rendered into the 
   `operatorSpec` as `SecretDestination`'s and subsequently read by the resource specific hooks in order to determine what (if any) `ListKeys` APIs to call
   and where to store the results.
4. Implement the `SecretWriter` interface described below (or something like it). Note that this is related to the `SecretReader` above in 
   the [input secrets](#input-secrets) section.
    ```go
    // TODO: these strings may need to be []byte
    // Note: these interfaces are acting on collections only so that they can be more efficient, as multiple SecretRef's or secret values may be read from or written to 
    // a single secret.
    type DestinationValuePair struct {
        Value string
        Destination SecretDestination
    }

    type SecretWriter interface {
        // This will perform ownership checks and return an error if an attempt is made to update a secret that is not owned by the operator
        SetSecrets(ctx context.Context, owner MetaObject, secrets []DestinationValuePair) error
    }
    ```
5. Use customization hooks to implement `GetKeys` or `ListKeys` for each applicable resource, utilizing the `SecretWriter` to write the secrets to their 
   destination.

## Testing

### Unit testing

All of the library code should have unit tests written:
1. Reflector to discover `SecretRef`.
2. `SecretReader`/`SecretWriter`

### End to end testing

The existing EnvTest tests can be expanded to test secret management. We will need to make sure that we redact the keys returned by `ListKeys`. Tests for resources which 
take secrets will need the secrets created beforehand and passed to the resource (just as the customer would have to do).

## Related issues
1. [Secrets created by ASO should have configurable annotations](https://github.com/Azure/azure-service-operator/issues/1398)
2. [Path to secrets created by ASO should be in the status](https://github.com/Azure/azure-service-operator/issues/1318) - maybe not needed if we're making the user tell us where to put it?
3. [Secrets deletion should be required as part of resource deletion](https://github.com/Azure/azure-service-operator/issues/1280)
4. [Add option to create and manage keyvault secrets through ASO](https://github.com/Azure/azure-service-operator/issues/1894)
