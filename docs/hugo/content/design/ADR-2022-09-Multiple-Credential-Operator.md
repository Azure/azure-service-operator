---
title: '2022-09: Support For Multiple Credentials Under Global Operator'
---

## Context

Some users of Azure Service Operator want to have a single instance that manages Azure resources across multiple subscriptions. This is useful, for example, if the organisation has a large cluster used by multiple teams, each team having their own Azure subscription.
To support this, ASO will introduce support for multiple credentials, each associated with a specific scope:

**Global Secret(aso-controller-settings)** is a currently existing one and is used as default for all Azure Resources

**Namespaced Secrets** will be used for all Azure Resources within a specific cluster namespace

**Per-resource-group Secrets** will be used for all Azure Resources within a specific Resource Group

**Per-resource Secrets** will be used for specific resources

## Credential selection hierarchy

We'll be using a pattern to determine which of the above secrets we use for operation on a resource. As in the below flow chart, if a resource is applied to the operator and the resource exists, operator would fetch and apply credentials in a following manner. If it's a new resource, operator would have to go through the secret selection hierarchy. Where, operator would first check if Per-resource or per-resource-group secret exists, if not, then will check for the namespaced secret. If any of the above is provided, operator would perform actions according to the **options** below. If none is provided, we'll fallback to use the global credential for that resource, which is the default cluster scoped secret(aso-controller-settings) in ASO namespace used today. 

![hierarchy](images/adr-2022-09-multiple-credential-operator.png) 

## Options

In all the proposed solutions: 
1. Operator would cache(in-memory) the credentials to be used for resources when we reconcile and a mechanism to identify the credentials used for a resource. 
2. Operator would add `serviceoperator.azure.com/subscription: <SUBSCRIPTION_ID>` annotation on each resource, which would be checked with the current credentials on each reconcile. This would prevent users from updating a credential and causing their resource to be created in another subscription without deleting it from the first subscription.

### Option 1: Fixed secret names

In secrets lookup option, fixed naming patterns would be used to look up a secret each time a new resource is created. Operator would first use the naming pattern for the resource group, then for the cluster namespace. If a secret is found with the expected name, the credentials for the resource would be cached for use the next time. If no secret was found matching any patter, the default global secret would be used.

NOTE: Annotation/condition will be added on the resource while creation which will have information about the secret name used for creation/reconcile.

**Fixed secret name pattern:**

| Secret Type        | Pattern                        |
|--------------------|--------------------------------|
| Namespaced         | aso-credential                 |
| Per-resource-group | aso-credential-{rg-name}       | 
| Per-resource       | aso-credential-{resource-name} |


**Pros:**
1. Users don't have to specify the configuration for each resource explicitly

**Cons:**
1. Not very flexible, as users don't have much control for which resource they want to use which credential
2. Can not re-use a secret across Resource Groups. Will have to specify secrets individually for each Resource Group, even if they are same
3. Not friendly to apply-at-once code, as there could be a race while applying secret and resource at the same time

### Option 2: Configuration using annotations

In Configuration using annotations option, we'll use annotations like below on the resource to fetch credentials from a secret and cache it. If secret is not found, operator would block the creation and set the condition to be `Ready condition == false` with `severity == Warning`, so that operator would keepo retrying to fetch the credentials until created. 

```
annotations: 
  serviceoperator.azure.com/credential-from: any-secret
```

**Pros:**
1. Flexible, as users have the control for which resource they want to use which credential

**Cons:**
1. Will have to specify annotation on every resource explicitly, involving lots of repetition

### Option 3: Fixed secret name for namespace + configuration using annotations

Here, we would use a fixed pattern for namespaces secrets + annotations for per-resource-group and per-resource secrets. Operator would look up for the annotation(`serviceoperator.azure.com/credential-from`) on resource created, if annotation found, then use credentials from secret in annotation. If not, then look for the namespaced secret and use the namespaced credential.If none provided, operator shall continue with the global credential.

NOTE: If namespaced secret is used, operator would inject `serviceoperator.azure.com/credential-from` annotation with the secret name value. This would be useful for the operator on each reconcile to know which secret should be used for auth and help users to figure out which credential is used for a resource.

**Fixed namespaced secret name pattern:**

| Secret Type        | Pattern                                                  | Notes                                                                |
|--------------------|----------------------------------------------------------|----------------------------------------------------------------------|
| Namespaced         | aso-credential                                           | Created in the Cluster namespace                                     |
| Per-resource-group | serviceoperator.azure.com/credential-from: `SECRET_NAME` | Annotation on resource mentioning the secret to use for the resource |
| Per-resource       | serviceoperator.azure.com/credential-from: `SECRET_NAME` | Annotation on resource mentioning the secret to use for the resource | 

**Pros:**
1. Users don't have to specify the configuration for each resource explicitly if using a namespaced secret
2. Users get the flexibility of overriding the credential they want to use as per resource and as per resource group

**Cons:**
1. Could be complicated/confusing for users to understand the usage and credentials being used. This can be mitigated to an extent by having `serviceoperator.azure.com/credential-from` annoation on each resource using credential other than global.  

### Option 4: Fixed secret name for namespace + configuration using global secret

We'd use Fixed secret name for namespace + configuration using global secret for per-resource-group and per-resource secret. We can add another list(CREDENTIAL_FROM) in the global secret where admin users can specify the secrets they want to load on the startup.
Then only these loaded secrets would be used further for per-resource and per-resource-group. This option would restrict the other non-admin users(who don't have access to the global secret) to add their credential. 

**Pros:**
1. Users don't have to specify the configuration for each resource explicitly if using a namespaced secret
2. Security conscious, not anyone in the cluster can add to use their credentials.

**Cons:**
1. Could be complicated/confusing for users to understand the usage and credentials being used
2. Not very flexible, as pod will have to startup again if a new secret needs to be loaded and not anyone can add secrets rather than the admin

## Decision

TBD

## Handling for the failure cases and logging

- Failure on loading a secret would result in failure on resource creation
- Failure on authenticating the credentials would result in failure on resource creation
- Logging should be clear about which credential is being used for a resource
- Logging should be clear about the credential load/authentication failure
- Errors can be reported through the standard `Condition` mechanism

## Milestones

1. Namespaced secrets: TBD
2. per-ResourceGroup secrets: TBD
3. per-Resource secrets: Will not be implemented unless there is strong user demand. We believe that per-ResourceGroup secrets should be sufficiently granular.

## Open questions
1. Spell out what fields are in aso-credential. clientid/subscriptionid/tenantId/etc?
2. How are we going to make this work for Managed Identity?

## Consequences

TBC

## Experience Report

TBC

## References

TBC
