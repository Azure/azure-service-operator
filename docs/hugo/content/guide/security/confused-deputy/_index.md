---
title: Confused Deputy Problem
linktitle: Confused Deputy
---

Confused Deputy is a security vulnerability that occurs when a low-privilege user is able to trick a high-privilege
service into performing actions on their behalf that they would not normally be allowed to do.

The entire premise of Azure Service Operator (ASO) is that it is a high-privilege service that can perform actions
in Azure on behalf of users.

Many ASO users deliberately block their users from direct Azure access, and instead require them to use ASO to manage
Azure resources, using ASO as a control plane that gives their users a limited set of permissions.

{{% alert title="Warning" color="warning" %}}
Before you report a Confused Deputy vulnerability against ASO, please verify that you've identified something beyond
this expected behavior of ASO.
{{% /alert %}}

## The ASO Security Model

Azure Service Operator (ASO) relies on Kubernetes RBAC for access control. Any user who can create or modify ASO
resources in a namespace that ASO is configured to watch can use ASO to perform any action in Azure that the Azure
credential configured for that namespace permits.

This makes securing access to namespaces where ASO credentials are configured critical: users who can create or modify
ASO resources in a namespace effectively have the same Azure permissions as the credential stored in that namespace.

For more information, see:

- [Security best practices]({{< relref "security" >}})
- [Reducing access]({{< relref "reducing-access" >}})
- [Best Practices for Using Azure Service Operator]({{< relref "best-practices" >}})
