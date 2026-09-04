---
title: Security Best Practices
linktitle: Security
---

## Securing ASO in your cluster

ASO has 3 levers that allow you to manage access to your Azure resources:

1. Controlling [which CRDs are installed]({{< relref "crd-management" >}}).
2. Controlling the Azure identities used by ASO at each [scope]({{< relref "credential-scope" >}}), 
   including the Azure RBAC permissions assigned to those identities.
3. Controlling the Kubernetes identities that use the cluster and their Kubernetes RBAC permissions.

We recommend making use of all 3 of these levers to fully secure a cluster running ASO.

## Dos and Don'ts

> ✅ DO adopt this pattern.
> 
> ⛔ DO NOT adopt this pattern.

### General guidance

✅ DO use [Azure Workload Identity]( {{< relref "credential-format#managed-identity-via-workload-identity" >}} ) for all
credentials. Other supported identity types are called out
[in the authentication documentation]( {{< relref "authentication#credential-type" >}} ).

✅ DO use namespace-scoped ASO credentials, rather than global scope. Note that the global scope credential is _optional_ 
and may be omitted when installing ASO.

✅ DO follow the [principle of least privilege](https://learn.microsoft.com/entra/identity/role-based-access-control/best-practices#1-apply-principle-of-least-privilege)
when assigning roles to identities which will be used by ASO. Remember, users with access to the namespace the 
ASO credential is in can do everything that credential can do via ASO. This means that if users in namespace `a` 
are supposed to have broad permissions only to resources in resourceGroup `a`, then the ASO 
identity for namespace `a` should have **Contributor** only on resourceGroup `a` and not the whole 
subscription. See [reducing access]( {{< relref "reducing-access" >}} ) for more details on managing Azure access.

✅ DO restrict access to sensitive namespaces in the cluster. On AKS, you can use a combination of
[AAD (now Entra) integration](https://learn.microsoft.com/en-us/azure/aks/enable-authentication-microsoft-entra-id),
[disabling local users](https://learn.microsoft.com/en-us/azure/aks/manage-local-accounts-managed-azure-ad), and
defining [JIT/Conditional access policies](https://learn.microsoft.com/en-us/azure/aks/access-control-managed-azure-ad).
We strongly recommend setting up conditional access policies for sensitive namespaces such as `production`. Doubly so
if the ASO credential for that namespace has broad scope.

✅ DO use tools like ArgoCD or Flux to perform code review of changes to ASO CRs before allowing them to be
merged and applied.

✅ DO only install the ASO CRDs you need, no more.

⛔ DO NOT install the `RoleAssignment` CRD if you don't need it. This CRD can enable escalation of privilege if not
used carefully. If using the `RoleAssignment`, follow the other DOs in this guide to do it safely.

## An Example Setup

A possible setup might be a `dev` namespace set up as a development environment pointing at a development subscription, 
and a `prod` namespace set up as a production env. The `dev` namespace might point to a development subscription 
and the `prod` namespace to a production subscription.

`dev` namespace has Azure credentials which are contributor on the `dev` subscription, 
same for `prod` for the prod sub. Developers in `dev` are members of an Azure AD group with roles that give 
access to CRUD ASO CRDs and other Kubernetes resources (Pods, etc) in the `dev` namespace, but _not_ the production 
namespace.

This means that developers can do basically whatever they want in the dev namespace, including assign roles to 
themselves at the Azure level in the dev subscription.

`prod` namespace _also_ has an Azure AD group with roles that give access to CRUD ASO CRDs and other 
Kubernetes resources, but that group is by default empty. 
Users can use [JIT/Conditional access policies](https://learn.microsoft.com/azure/aks/access-control-managed-azure-ad)
to escalate into that group. This means that by default, nobody can do anything in the `prod` namespace to 
either the Kubernetes resources (Pods, etc) or the Azure resources via ASO or the portal.

When a user needs ad-hoc access to the `prod` namespace they can go through the JIT process to get 
access to the `prod` namespace. Standard deployments to `prod` should be done through a CI/CD tool 
like Argo or Flux. This has the advantage of ensuring that proposed changes to `prod` need to first meet the merge bar 
(pass through code review and other processes) to make it into the repo before Argo/Flux will deploy them to `prod`. 
The conditional access/JIT is a break-glass used rarely.

Note that the above is just _one_ way to lay things out. The same ideas can be applied to a `dev` and `prod` resource group 
within a single sub and also to other more complex topologies. `test` or `int` can be added in the middle with a more 
locked down set of rules than dev but less locked down than prod (or maybe `test` and `prod` have very similar 
lockdowns to force the same procedures across both).
