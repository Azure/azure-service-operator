---
title: Workload Identity isolation
linktitle: Workload Identity isolation
---

## Workload Identity basics

Usually when an operator supports managed identity via workload identity, the operator has a single identity (`ServiceAccount`) 
and that `ServiceAccount` is federated with Azure via a `FederatedIdentityCredential` so that it can exchange its token with Azure for
the ManagedIdentity token, which is then used to authenticate all calls to Azure.

In a multitenant environment, resources in each namespace are still reconciled by a single operator pod with a single `ServiceAccount`.
When using workload identity in each namespace, that `ServiceAccount` is granted a `FederatedIdentityCredential` that allows it to 
exchange its token for the managed identity defined in the `aso-credential` of that namespace.

## The risk with shared trust

Workload Identity credential secrets stored in `aso-credential` or other secrets in the namespace contain 
the fields `AZURE_TENANT_ID`, `AZURE_SUBSCRIPTION_ID`, and `AZURE_CLIENT_ID`. These fields are not secrets/passwords and as such they can
be discovered via the Azure API or may be logged or otherwise visible for attackers to see. 

The consequence of this is that a user who can create credential secrets and ASO resources in their own namespace can 
copy those identifiers and direct ASO to use that identity. They do not need to read another namespace's secret.
[Azure/azure-service-operator#4810](https://github.com/Azure/azure-service-operator/issues/4810) describes the risk in more detail.

## Relaxed vs. strict mode

* **`relaxed` (default):** Namespace-scoped and per-resource Workload Identity credentials use the controller's
  `ServiceAccount` token. Setup is simpler, but credential scopes alone do not prevent the cross-namespace identity reuse
  described above.
* **`strict`:** ASO requests a token for a `ServiceAccount` in the credential secret's namespace, rather than using the
  controller's token. The Azure identity's `FederatedIdentityCredential` must trust that namespace-specific `ServiceAccount`.
  Copying identifiers to another namespace is insufficient: its `ServiceAccount` subject does not match the trusted subject.

## Choosing strict mode

Use strict mode for mutually untrusted namespaces. Follow these steps in the [credential format guide]( {{< relref "credential-format" >}} ):

1. [Set Helm's `workloadIdentityAuthMode` to `strict`]( {{< relref "credential-format#configure-trust" >}} )
   using the **Managed Identity (strict)** tab.
2. [Create a ServiceAccount in each tenant namespace]( {{< relref "credential-format#create-the-serviceaccount-strict-mode-only" >}} ).
3. [Grant ASO token-creation permissions]( {{< relref "credential-format#create-the-serviceaccount-strict-mode-only" >}} )
   if using a custom ServiceAccount name; Helm handles the default name.
4. [Configure Azure federated trust for each namespace's ServiceAccount]( {{< relref "credential-format#configure-trust" >}} )
   using the **Managed Identity (strict)** tab.
5. [Create or update the namespace-scoped or per-resource credential secret]( {{< relref "credential-format#create-the-secret" >}} ).

{{% alert title="Note" %}}
Strict mode does not change the global credential or other (non-workload identity) authentication types, so avoid a privileged global credential
available to all namespaces.
{{% /alert %}}
