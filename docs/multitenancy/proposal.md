# Multitenancy approach

This document presents our proposed approach for implementing multitenancy support in Azure Service Operator.

## Overview

The operator needs to support creating Azure resources using multiple credentials, managed by different groups, within one cluster.
This will be enabled by changing the operator configuration to allow specifying a set of the cluster's namespaces to watch for resources to manage, and then running a deployment of the ASO operator per tenant.

Each deployment of the operator will run in a separate namespace (rather than the current default `azureoperator-system`) and read its configuration from a secret named `azureoperatorsettings` in that namespace.
If the operator is configured to use managed identity (`AZURE_USE_MI=1`), then the `AzureIdentity` and `AzureIdentityBinding` will be found in that namespace as well.
Each instance of the operator will use the service account defined in its namespace.
(This is all the same as the current behaviour of the operator.)

A new configuration attribute `AZURE_TARGET_NAMESPACES` will control which namespaces that operator will monitor.
This should be a comma-separated list of names, which could include the namespace containing the operator deployment (but doesn't have to).
Having no namespaces in the list or omitting the setting from the configuration will result in the operator watching all namespaces for ASO resources as it does now.

// TODO: How to ensure that each namespace is managed by at most one operator? Having different operators (with potentially different credentials) fighting over a single resource would be messy. Similarly having both a global ASO and per-namespace instances would cause problems.

### CRDs and webhooks

`CustomResourceDefinitions`, `ValidatingWebhookConfigurations` and `MutatingWebhookConfigurations` are cluster-wide so they can't be managed on a per-operator basis.
In a multitenant configuration installing and managing the CRDs and webhooks will be a separate step from deploying operators (and will require cluster admin privileges).

The conversion webhooks in the CRDs and the validating and mutating webhook configurations that are defined for different ASO resources will all point to the `azureoperator-webhook-service` service deployed in the `azureoperator-system` namespace (by default).
The service will be backed by a deployment of the normal ASO operator running in a restricted webhook-only mode - it won't be watching any namespaces for Azure resources and won't have access to any Azure credentials.
This mode will be triggered by the configuration setting `AZURE_WEBHOOK_ONLY=1`.
With this setting no other configuration values will be required, and no resource watches will be started - the operator will just respond to webhook requests from the API server.

### RBAC

At the moment installing a global instance of ASO creates a `ClusterRole` defining all of the permissions the operator requires on the different Azure resources (`azureoperator-manager-role`), and then a `ClusterRoleBinding` granting those permissions for resources in all namespaces (`azureoperator-manager-rolebinding`).
A multitenant deployment won't have the global `ClusterRoleBinding`; instead each managed namespace N will need to contain a `RoleBinding` connecting the global `ClusterRole` with the service account of the operator that will manage N's resources.

### Advantages and constraints

The primary benefit of this architecture is that it enables different tenants to manage their operators independently along with their credentials and resources.
Additionally in the case of operators using managed identities it removes the need to replicate `AzureIdentity` and `AzureIdentityBinding` resources between namespaces to ensure that the pod identity machinery configures the node identities correctly for the operator.

The biggest downside (other than complexity) is that the different operator installations will be separately managed and we can't assume that they will all be upgraded at the same time.
The CRDs and webhook service deployment will need to be upgraded by the cluster administrator before any of the tenant operators can be upgraded.
Conversely if there's a release of the operator that removes an old version of a resource definition, the cluster administrator will need to ensure that all of the tenant operators and resources have been upgraded beyond that version before the core CRDs and webhook service are upgraded.
The [storage-version-migrator operator](https://github.com/kubernetes-sigs/kube-storage-version-migrator) might be useful for this.

### Alternatives

An earlier plan for implementing multitenancy in ASO was to keep the operator deployment as it is now (one per cluster) but add flexibility in how the credentials would be looked up.
A proof of concept for this approach can be seen in [PR #1206](https://github.com/Azure/azure-service-operator/pull/1206).
This would be simpler in terms of deployments, but a global operator with access to different tenants' credentials carries the risk of applying the wrong credentials when reconciling an Azure resource.
The per-tenant architecture has been chosen to reduce that risk, since any given operator will only have access to the credentials it's configured with.

## Upgrading

As discussed above, the CRDs and webhook service will need to be upgraded to a given version before any of the per-tenant operators can be upgraded.
If a tenant operator is upgraded beyond the global version, the new operator pods will crash complaining about an unknown custom resource version until the deployment image version is rolled back.
In the longer term it would be useful to have tool support for upgrading a tenant operator that would prevent upgrading beyond the version of the ASO CRDs and webhook service.

We haven't needed to remove an old version of a resource from its definition yet, but once code-generated resources are incorporated into the operator we anticipate needing to do this.
At the moment the best way to manage that is to ensure that all the tenant operator's resources are migrated to the current storage version after an upgrade using an approach like the storage-version-migrator operator.
Checking that the resources have all been migrated to the latest version would ensure that it is safe to upgrade the CRDs removing the old version.
Unfortunately that project is still in a prerelease stage of development.
(This open problem also affects the current cluster-global operator.)

## Roadmap

### Initial per-namespace changes

### Deployment tooling

### Upgrade support
