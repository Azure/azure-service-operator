---
title: "Authentication in Azure Service Operator"
linkTitle: "Authentication"
layout: single
cascade:
- type: docs
- render: always
---

There are two key topics surrounding authentication in Azure Service Operator: The type of credential, and the credential scope.

## Credential type

Azure Service Operator supports four different styles of authentication today.

1. [Service Principal using a Client Secret](./credential-format#service-principal-using-a-client-secret)
2. [Service Principal using a Client Certificate](./credential-format#service-principal-using-a-client-certificate)
3. [Azure-Workload-Identity authentication](./credential-format#azure-workload-identity) (OIDC + Managed Identity or Service Principal)
4. [Deprecated] [aad-pod-identity authentication (Managed Identity)](./credential-format#deprecated-managed-identity--aad-pod-identity-)

## Credential scope

Each supported credential type can be specified at one of three supported scopes:

1. [Global](./credential-scope#global-scope) - The credential applies to all resources managed by ASO.
2. [Namespace](./credential-scope#namespace-scope) - The credential applies to all resources managed by ASO in that namespace.
3. [Resource](./credential-scope#resource-scope) - The credential applies to only the specific resource it is referenced on.

When presented with multiple credential choices, the operator chooses the most specific one:
_resource scope_ takes precedence over _namespace scope_ which takes precedence over _global scope_.

> **Warning:** The operator identity is used to access the global, namespace, and resource scoped secrets. A user with
> access to create ASO resources but without Kubernetes secret read permissions can still direct ASO to use a secret the
> user cannot read. The user cannot access the contents of the secret, but they can manage resources in Azure
> via the identity the secret represents.
>
> The namespace is the security boundary. ASO will not allow users to read secrets from other namespaces. We recommend
> using separate namespaces for separate environments (dev, test, prod, etc) for this reason

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

For more details about this approach, see [multitenant deployment](./multitenant-deployment)