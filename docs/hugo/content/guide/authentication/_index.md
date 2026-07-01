---
title: "Authentication in Azure Service Operator"
linkTitle: "Authentication"
layout: single
weight: 1 # This is the default weight if you just want to be ordered alphabetically
cascade:
- type: docs
- render: always
---

There are two key topics surrounding authentication in Azure Service Operator: The type of credential, and the credential scope.

## Credential type

Azure Service Operator supports five different styles of authentication today.

1. [Recommended for production] [Azure Workload Identity]( {{< relref "credential-format#managed-identity-via-workload-identity" >}} ) (OIDC + Managed Identity or Service Principal)
2. [Service Principal using a Client Secret]( {{< relref "credential-format#service-principal-using-a-client-secret" >}} )
3. [Service Principal using a Client Certificate]( {{< relref "credential-format#service-principal-using-a-client-certificate" >}} )
4. [Deprecated] [aad-pod-identity authentication (Managed Identity)]( {{< relref "credential-format#deprecated-managed-identity-aad-pod-identity" >}} )
5. [User Assigned Identity Credentials]( {{< relref "credential-format#user-assigned-identity-credentials" >}} )

## Credential scope

Each supported credential type can be specified at one of three supported scopes:

1. [Not recommended] [Global]( {{< relref "credential-scope#global-scope" >}} ) - The credential applies to all resources managed by ASO.
2. [Namespace]( {{< relref "credential-scope#namespace-scope" >}} ) - The credential applies to all resources managed by ASO in that namespace.
3. [Resource]( {{< relref "credential-scope#resource-scope" >}} ) - The credential applies to only the specific resource it is referenced on.

When presented with multiple credential choices, the operator chooses the most specific one:
_resource scope_ takes precedence over _namespace scope_ which takes precedence over _global scope_.

> **Warning:** The operator identity is used to access the global, namespace, and resource scoped secrets. A user with
> access to create ASO resources but without Kubernetes secret read permissions can still direct ASO to use a secret the
> user cannot read. The user cannot access the contents of the secret, but they can manage resources in Azure
> via the identity the secret represents.
>
> The namespace is the security boundary. ASO will not allow users to read secrets from other namespaces. We recommend
> using separate namespaces for separate environments (dev, test, prod, etc) for this reason

## Using ASO to create its own credentials

It is possible to use ASO to create its own [namespace-scoped]( {{< relref "credential-scope#namespace-scope" >}} ) or 
[resource-scoped]({{< relref "credential-scope#resource-scope" >}}) credentials.

Achieving this requires two things:

1. A namespace for managing credentials, complete with `aso-credential` (namespace-scoped) or [global credential (not recommended)]( {{< relref "credential-scope#global-scope" >}} )
   that has permission to create `managedidentity.azure.com/UserAssignedIdentity` and `authorization.azure.com/RoleAssignment` resources. This namespace will be used
   to create/manage the credentials.
2. An operator such as [kubernetes-reflector](https://github.com/emberstack/kubernetes-reflector), used to mirror the credential into the target namespace, which will use the credential

Here's an example where the management namespace creates a `resources.azure.com/ResourceGroup`, `managedidentity.azure.com/UserAssignedIdentity` and `authorization.azure.com/RoleAssignment` 
granting the identity Contributor on the resource group.

{{% alert title="Note" %}}
ASO does not support creating secrets in namespaces other than the namespace the corresponding resource is in.
{{% /alert %}}

{{% alert title="Note" %}}
The secret written by `secretExpressions` must **not** be named `aso-credential` in the management namespace, as that
would conflict with the existing namespace-scoped credential that ASO uses to manage resources there.
Use a different name (e.g. `sample-identity-credential`) and have reflector mirror it into the target namespace
under the name `aso-credential`.
{{% /alert %}}

```yaml
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: management
spec:
  location: westcentralus
---
apiVersion: managedidentity.azure.com/v1api20230131
kind: UserAssignedIdentity
metadata:
  name: sample-identity
  namespace: management
spec:
  location: westcentralus
  owner:
    name: aso-sample-rg
  operatorSpec:
    secretExpressions:
      - name: sample-identity-credential
        key: AZURE_SUBSCRIPTION_ID
        value: "aso.parseResourceId(self.status.id).subscriptionId"
        annotations:
          reflector.v1.k8s.emberstack.com/reflection-allowed: '"true"'
          reflector.v1.k8s.emberstack.com/reflection-allowed-namespaces: '"app-namespace"'
      - name: sample-identity-credential
        key: AZURE_TENANT_ID
        value: "self.status.tenantId"
      - name: sample-identity-credential
        key: AZURE_CLIENT_ID
        value: "self.status.clientId"
      - name: sample-identity-credential
        key: USE_WORKLOAD_IDENTITY_AUTH
        value: "string(true)"
---
apiVersion: authorization.azure.com/v1api20220401
kind: RoleAssignment
metadata:
  name: sample-identity-contributor
  namespace: management
spec:
  owner:
    name: aso-sample-rg
    group: resources.azure.com
    kind: ResourceGroup
  principalIdFromConfig:
    name: sample-identity-settings
    key: principalId
  roleDefinitionReference:
    wellKnownName: Contributor
```

Once the resources are provisioned, a user in `app-namespace` can mirror the credential by creating a stub secret 
with the `reflects` annotation:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: aso-credential
  namespace: app-namespace
  annotations:
    reflector.v1.k8s.emberstack.com/reflects: "management/sample-identity-credential"
```

[kubernetes-reflector](https://github.com/emberstack/kubernetes-reflector) will automatically populate this secret
with the contents of `sample-identity-credential` from the `management` namespace. Since the destination secret is 
named `aso-credential`, ASO will use it as the namespace-scoped credential for resources in `app-namespace`, 
allowing them to manage Azure resources using workload identity.

## Using multiple operators with a single credential per operator

> **This mode is not recommended unless you _really_ need it**

ASO also supports installing multiple instances of the operator alongside one another, each
configured with different credentials.

This option exists for the most security conscious customers. 

Advantages:
* Each operator pod only has access to a single credential, reducing risk if one pod is somehow compromised.

Disadvantages:
* Significantly harder to orchestrate ASO upgrades.
* More kube-apiserver load, as there will be multiple operators running and watching/reconciling resources.

For more details about this approach, see [multitenant deployment]( {{< relref "multitenant-deployment" >}} )
