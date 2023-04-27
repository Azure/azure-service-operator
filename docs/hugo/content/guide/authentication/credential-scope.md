---
title: Credential Scope
linktitle: Credential Scope
---

When presented with multiple credential choices, the operator chooses the most specific one:
_resource scope_ takes precedence over _namespace scope_ which takes precedence over _global scope_.

## Global scope

The global credential resides in the `aso-controller-settings` secret deployed as part of operator deployment in 
operator's namespace.
This is the scope configured when configuring credentials via the Helm chart installation.

When updating the `aso-controller-settings` secret, you must restart the operator pod for the updated secret to 
take effect.

> **Note**: The `aso-controller-settings` secret must exist, but the actual global credential portion of the secret
> **is optional**. If you do not wish to use a global credential, you must still provide the secret with empty values
> for the `AZURE_SUBSCRIPTION_ID`, and `AZURE_CLIENT_ID`, and `AZURE_TENANT_ID` strings.

## Namespace scope

A secret named `aso-credential` in a Kubernetes namespace configures the credential for all resources in that namespace,
overriding the global credential. This credential can be in a different tenant or subscription than the global credential.

## Resource scope

A secret with any name can be referenced by the `serviceoperator.azure.com/credential-from` annotation.
Create this annotation on each resource to configure the credential used for that resource. 
The secret containing the credential must be in the same Kubernetes namespace as the resource but may be in a different
tenant or subscription than both the global credential and per-namespace credential.

Note that multiple resources may refer to the same secret.
