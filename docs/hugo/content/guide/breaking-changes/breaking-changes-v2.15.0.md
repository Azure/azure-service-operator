---
title: "v2.15.0 Breaking Changes"
linkTitle: "v2.15.0"
weight: -50  # This should be 5 lower than the previous breaking change document
---
## Breaking changes

Improvements to our code generator have resulted in some properties now being more tightly validated, with better alignment to the requirements of the associated Azure resources.

Some invalid resources that were previously accepted by ASO only to fail when reconciled will now be rejected up front. Valid resources will be unaffected.

Affected properties are:

| Group                   | Resource                    | Property                                   |
| ----------------------- | --------------------------- | ------------------------------------------ |
| kubernetesconfiguration | fluxConfiguration           | spec.ociRepository.url                     |
| machinelearningservices | workspacesConnection        | spec.properties.oAuth2.credentials.authUrl |
| managedidentity         | federatedIdentityCredential | spec.issuer                                |
