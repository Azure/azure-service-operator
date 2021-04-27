# Multitenancy approach

This document presents our proposed approach for implementing multitenancy support in Azure Service Operator.

## Overview

The operator needs to support creating Azure resources using multiple credentials, managed by different groups, within one cluster. There are (at least) two ways we could handle multiple credentials in ASO:

1. Allow running multiple operators in one cluster each with its own credentials and set of namespaces to watch for Azure resources.
2. Keep the operator cluster-global but add a credential lookup when reconciling so that the identity used to communicate with Azure can be picked on a per-namespace (or per-resource) basis. A proof of concept of this approach can be seen in [PR #1206](https://github.com/Azure/azure-service-operator/pull/1206).

Initially we'll be working on option 1, but in the long term we expect to support both options for different use cases.

Option 1 will be enabled by changing the operator configuration to allow specifying a set of the cluster's namespaces to watch for resources to manage, and then running a deployment of the ASO operator per tenant.
Each deployment of the operator will run in a separate namespace (rather than the current default `azureoperator-system`) and read its configuration from a secret named `azureoperatorsettings` in that namespace.
If the operator is configured to use managed identity (`AZURE_USE_MI=1`), then the `AzureIdentity` and `AzureIdentityBinding` will be found in that namespace as well.
Each instance of the operator will use the service account defined in its namespace.
(This is all the same as the current behaviour of the operator.)

A new configuration attribute `AZURE_TARGET_NAMESPACES` will control which namespaces that operator will monitor.
This should be a comma-separated list of names, which could include the namespace containing the operator deployment (but doesn't have to).
Having no namespaces in the list or omitting the setting from the configuration will result in the operator watching all namespaces for ASO resources as it does now.

To make it clear which operator is reconciling a given resource, it will add a `azure.microsoft.com/operator-namespace` annotation when updating the status.
This will also allow one instance of ASO to detect when it is fighting with another over a resource (because they're both configured to manage that namespace) and highlight the potential problem in logging.

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

An earlier plan for implementing multitenancy in ASO was to keep the operator deployment as it is now (one per cluster) but add flexibility in how the credentials would be looked up - this is option 2 above.
A proof of concept to find credentials based on a resource's namespace can be seen in [PR #1206](https://github.com/Azure/azure-service-operator/pull/1206).
This would be simpler in terms of deployments, but a global operator with access to different tenants' credentials carries the risk of applying the wrong credentials when reconciling an Azure resource, a risk that may be unacceptable for users who are very security conscious.
We have chosen to implement the operator-per-tenant architecture first to enable the most security conscious customers.
Along the way we will hopefully be able to lay the foundation for other multitenancy configurations as well.

Once that's ready we'll come back to implementing more flexible credential matching.
There are also potential ASO users who would benefit from being able to assign credentials directly to resources, which can be implemented as an extension of the per-namespace credential handling.

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
