---
title: '2022-09: Support For Multiple Credentials Under Global Operator'
---

## Context

The operator needs to support creating Azure resources using multiple credentials, managed by different groups, within one cluster by a single global instance of operator. 
The idea is to support the above by having **Namespaced secret**, **Per-resource-group secret** and Per-resource secret. 

### Namespaced Secret

Namespaced secret holds credentials used for the resources in a namespace.

### Per-resource-group Secret

Per-resource-group secret holds credential used for the resource under a resource group.

### Per-resource Secret

Per-resource-group secret holds credential used for the resource separately.

## Credential selection hierarchy

We'll be using a pattern to determine which of the above secrets we use for operation on a resource. As in the below flow chart, if a resource is applied to the operator and the resource exists, operator would fetch the credentials from
cache and perform the operation. If its a new resource, operator would have to go through the secret selection hierarchy. Where, operator would first check if Per-resource or per-resource-group secret exists, if not, then will check for 
the namespaced secret. If any of the above is provided, operator would perform actions according to the **options** below. If none is provided, we'll fallback to use the global credential for that resource. 

![hierarchy](images/adr-2022-09-multiple-credential-operator.png) 

## Options

In all the proposed solutions, we would need to cache the credentials to be used for resources when we reconcile and a mechanism to identify the credentials used for a resource. 

### Option 1: Secrets lookup

In secrets lookup option, we would need a fixed secret pattern like below. Which then will be looked up for each time a new resource is created. Operator would look up for the secret pattern for resource group first, then namespace. If secret exists, cache the credentials for the resource to be used for the next time or use the default global.

NOTE: Annotation/condition will be added on the resource while creation which will have information about the secret name used for creation/reconcile.

**Fixed secret name pattern:**
```
Namespaced secret -> aso-credential (in the namespace)
Per-resource-group secret -> aso-credential-{rg-name}
Per-resource secret -> aso-credential-{resource-name}
```

**Pros:**
1. Users don't have to specify the configuration for each resource explicitly

**Cons:**
1. Not very flexible, as users don't have much control for which resource they want to use which credential

### Option 2: Configuration using annotations

In Configuration using annotations option, we'll use annotations like below on the resource to fetch credentials from a secret and cache it. Will default to global if secret with name not found or no annotation provided.

```
annotations: 
  serviceoperator.azure.com/credential-from: any-secret
```

**Pros:**
1. Flexible, as users have the control for which resource they want to use which credential

**Cons:**
1. Will have to specify annotation on each resource explicitly

### Option 3: Secret lookup + configuration using annotations

Here, we would use a fixed pattern for namespaces secrets + annotations for per-resource-group and per-resource secrets. Operator would look up for the annotation(`serviceoperator.azure.com/credential-from`) on resource created, if annotation found, then use credentials from secret in annotation. If not, then look for the namespaced secret and use the namespaced credential.If none provided, operator shall continue with the global credential.

NOTE: If namespaced secret or global secret is used, operator would inject `serviceoperator.azure.com/credential-from` annotation with the secret name value. 

**Fixed namespaced secret name pattern:**
```
Namespaced secret -> aso-credential (in the namespace)
```

**Annotation for per-resource-group and per-resource secret**
```
annotations: 
  serviceoperator.azure.com/credential-from: any-secret
```

**Pros:**
1. Users don't have to specify the configuration for each resource explicitly if using a namespaced secret
2. Users get the flexibility of over-riding the credential they want to use as per resource and as per resource group

**Cons:**
1. Could be complicated/confusing for users to understand the usage and credentials being used

### Option 4: Secret lookup + configuration using global secret

This option, relates to the option 3. However, the change is how we configure to use the per-resource-group and per-resource secret. We can have a unique slice property(CREDENTIAL_FROM) in the global secret where admin users can specify the secrets they want to load on the startup.
Then only these loaded secrets would be used further for per-resource and per-resource-group. This option would restrict the other non-admin users(who don't have access to the global secret) to add their credential. 

**Pros:**
1. Users don't have to specify the configuration for each resource explicitly if using a namespaced secret
2. Security conscious, not anyone in the cluster can add to use their credentials.

**Cons:**
1. Could be complicated/confusing for users to understand the usage and credentials being used
2. Not very flexible, as pod will have to startup again if a new secret needs to be loaded

## Handling for the failure cases and logging

- Failure on loading a secret would result in failure on resource creation
- Failure on authenticating the credentials would result in failure on resource creation
- Logging should be clear about which credential is being used for a resource
- Logging should be clear about the credential load/authentication failure

## Milestones

### Support for namespaced secrets

We aim to do the initial work needed in the operator to support this type of multitenancy and add support for namespaced secrets.

**Timelines:** TBD

### Support for per-resource-group secrets

We aim to support per-resource-group secrets under this milestone and configuration work required.

**Timelines:** TBD

### Support for per-resource secrets

We aim to support per-resource secrets under this milestone and configuration work required.

**Timelines:** TBD

### Support for credential cache

TBD

## Open questions
1. Do we need to support both per-resource and per-resource-group credentials?
2. Do we want a serviceoperator.azure.com/subscription-id annotation that we write? We can use this to protect against users accidentally updating credential to move resources between subscriptions.
3. Spell out what fields are in aso-credential. clientid/subscriptionid/tenantId/etc?
4. How are we going to make this work for Managed Identity?
5. Replace the above-mentioned(in options) annotations with spec property?

## Consequences

TBC

## Experience Report

TBC

## References

TBC
