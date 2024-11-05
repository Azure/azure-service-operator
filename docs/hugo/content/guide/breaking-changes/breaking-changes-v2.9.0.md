---
title: "v2.9.0 Breaking Changes"
linkTitle: "v2.9.0"
weight: -35  # This should be 5 lower than the previous breaking change document
---

# DelegatedManagedIdentityResourceId is now a secret

The RoleAssignment property `.spec.delegatedManagedIdentityResourceId` has been changed from a string to a 
SecretReference and renamed to `.spec.delegatedManagedIdentityResourceReference`.

We try to avoid breaking changes, but in this case, allowing raw passwords in the spec is a security 
problem and as such we've decided to make a break to correct this issue.

**Action required:** If the `authorization.azure.com/RoleAssignment` resource is used in your cluster, and if you are 
using the `DelegatedManagedIdentityResourceId` property, you'll need to change your resource to use the new 
`DelegatedManagedIdentityResourceReference` property, pulling the value from a secret.

# ContainerService version v1api20230202preview has been deleted

The AKS preview version `2023-02-02-preview` has been deprecated by AKS and is no longer supported. We've removed it 
from the operator to prevent unrecoverable errors from occurring.

**Action required:** 

- If you are using this specific version, change your resources to use a different version _before_ you upgrade to v2.9.0.
- Check if the `trustedaccessrolebindings.containerservice.azure.com` CRD is installed. If it is, check if the
  `status.storedVersions` field contains the `v1api20230202previewstorage` API. 
  - If it doesn't, no action is needed.
  - If it **does**, download the latest [experimental](https://github.com/Azure/azure-service-operator/releases/tag/experimental)
    or v2.11.0+ asoctl and run `asoctl clean crds`.
