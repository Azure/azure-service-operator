---
title: "v2.1.0 Breaking Changes"
linkTitle: "v2.1.0"
weight: 80
---

## The operator no longer installs CRDs by default

**Action required:** When installing ASO for the first time, you must now specify `crdPattern` (for Helm) or `--crd-patterns` 
(in operator pod cmdline for raw YAML) to select the subset of CRDs you would like to install.

When upgrading ASO, existing CRDs will be automatically updated to the new version but new CRDs added in that release 
will not automatically be installed. 
This means that when upgrading the operator, if you don't want to use any CRDs newly added in that release you don't 
need to do anything.

**Action required:** When upgrading ASO, if you want to install new CRDs (for example CRDs just added in the version of 
ASO you are upgrading to) you must specify `crdPattern` (Helm) or `--crd-patterns` (YAML) to install the CRDs. 
For example: if you do want to use a newly added CRD (such as `network.azure.com/bastionHosts` mentioned
below), you would need to specify `crdPatterns=network.azure.com/*` when performing the upgrade.

See [CRD management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for more details 
about this change and why it was made.

## `serviceoperator.azure.com/credential-from` no longer supports cross namespace secret references

This was never documented as supported but worked unintentionally. The feature now works as it was always documented: 
allowing references to secrets only if the secret is in the same namespace as the resource itself.

This was a security issue which we had to close.

See [#2919](https://github.com/Azure/azure-service-operator/pull/2919)

## Upcoming breaking changes

### AKS ManagedClusterServicePrincipalProfile.Secret will change from `string` to `genruntime.SecretReference`

We realized that this field contains a secret and as such _should not be specified_. Secrets should not appear in plain
text in CRs. We will be making a breaking change in 2.2.0 to resolve this issue.

**In the meantime:** We strongly recommend you use managed identity (the default) for your clusters.