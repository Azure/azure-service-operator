---
title: Confused Deputy Problem
linktitle: Confused Deputy
---

A confused deputy vulnerability occurs when a service with elevated permissions is induced to use those permissions
in a way that the requester is not authorized to perform.

Azure Service Operator (ASO) is designed to manage Azure resources on behalf of Kubernetes users. Organizations commonly
use ASO as a control plane: users who do not have direct Azure access can manage an approved set of Azure resources
through Kubernetes, using their Kubernetes cluster permissions.

This delegation is expected behavior. It becomes a security issue if a user can cross an intended authorization
boundary—for example, by using a credential that should not apply to their resource or by acting through a namespace
where they do not have the required Kubernetes permissions.

{{% alert title="Important" color="warning" %}}
When reporting a suspected confused deputy vulnerability, describe the authorization boundary that ASO allowed the
user to cross. If you are uncertain whether behavior is expected, report it so that we can assess it.
{{% /alert %}}

## The ASO Security Model

ASO relies on Kubernetes RBAC to control who can create or modify ASO resources. For each resource, ASO selects an Azure
credential according to its [credential scope]({{< relref "credential-scope" >}}).

A user who can create or modify ASO resources can cause ASO to perform supported Azure operations permitted by the
selected credential. The user may be able to exercise those permissions through ASO without being able to read or use
the credential directly.

Treat write access to ASO resources as delegated access to Azure. Restrict that access in every namespace ASO watches,
and grant each Azure identity only the permissions required for resources managed from that namespace.

For more information, see:

- [Security best practices]({{< relref "security" >}})
- [Reducing access]({{< relref "reducing-access" >}})
- [Best Practices for Using Azure Service Operator]({{< relref "best-practices" >}})
