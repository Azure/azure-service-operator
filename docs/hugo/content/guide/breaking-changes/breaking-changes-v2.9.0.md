---
title: "v2.8.0 Breaking Changes"
linkTitle: "v2.8.0"
weight: -25  # This should be 5 lower than the previous breaking change document
---

# DelegatedManagedIdentityResourceId is now a secret

The RoleAssignment property `.spec.DelegatedManagedIdentityResourceId` has been changed from a string to a SecretReference and renamed to `.spec.DelegatedManagedIdentityResourceReference`.

We try to avoid breaking changes, but in this case, allowing raw passwords in the spec is a security 
problem and as such we've decided to make a break to correct this issue.

**Action required:** If the `authorization.azure.com/RoleAssignment` resource is used in your cluster, and if you are using the `DelegatedManagedIdentityResourceId` property, you'll need to change your resource to use the new `DelegatedManagedIdentityResourceReference` property, pulling the value from a secret.

# ContainerService version v1api20230202preview has been deleted

The AKS preview version `2023-02-02-preview` has been deprecated by AKS and is no longer supported. We've removed it from the operator to prevent unrecoverable errors from occurring.

**Action required:** If you are using this specific version, change your resources to use a different version _before_ you upgrade to v2.9.0. 

