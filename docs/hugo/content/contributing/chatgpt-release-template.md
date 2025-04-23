---
title: A simple ChatGPT release template
---

## About

This can be used to help format our release notes once you've got them somewhat organized by category.

We can improve this template based on our experience with it. Starting now with something basic.

## Prompt
Please help me format the upcoming release notes for this repository. See below for how each sections text should be transformed.

Example input for "New resources" section:

```
### New resources

* Add support for ApplicationGatewayWebApplicationFirewallPolicies by @matthchr in https://github.com/Azure/azure-service-operator/pull/4238
* Add support for 20240401 AML resources by @super-harsh in https://github.com/Azure/azure-service-operator/pull/4237
* Add support for KubernetesConfiguration/FluxConfiguration by @super-harsh in https://github.com/Azure/azure-service-operator/pull/4275
* Add support for v1api20230801/Cache.Redis by @super-harsh in https://github.com/Azure/azure-service-operator/pull/4287
* Add support for network DnsForwardingRuleSetsVirtualNetworkLink VirutalNetworkLinks by @matthchr in https://github.com/Azure/azure-service-operator/pull/4345
* Feature/application security group by @nishant221 in https://github.com/Azure/azure-service-operator/pull/4342
* Add support for AML Registries by @super-harsh in https://github.com/Azure/azure-service-operator/pull/4339
```

Example output for "New resources" section:
```
### New resources

* Add support for new machinelearningservices API version v1api20240401 (#4237)
* Add support for new machinelearningservices Registry resource (#4339)
* Add support for new kubernetesconfiguration FluxConfiguration resource (#4275)
* Add support for new cache API version v1api20230801 (#4287)
* Add support for new network DnsForwardingRuleSetsVirtualNetworkLink resource (#4345)
* Add support for new network ApplicationSecurityGroup group resource (#4342)
* Add support for new network ApplicationGatewayWebApplicationFirewallPolicies resource (#4238)
```

For the new resources section, please ensure that each entry refers to the resource by its Kubernetes group. The Kubernetes group can
be found by looking at the PR in question. There will be changes in the v2/api/$group folder, for example https://github.com/Azure/azure-service-operator/pull/4237
changes the v2/api/machinelearningservices folder, so the group is machinelearningservices, and the entry in the release notes should say
machinelearningservices rather than AML.

Example input for "Features" section:

```
### Features

* Allow the data encryption keys for PostgreSQL Flexible server to be configured by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4311
* feat(helm): allow to use existing service account by @t3mi in https://github.com/Azure/azure-service-operator/pull/4211
* feat(helm): allow hardcodes to be configurable by @t3mi in https://github.com/Azure/azure-service-operator/pull/4207
* Remove pre-upgrade check by @matthchr in https://github.com/Azure/azure-service-operator/pull/4217
* Update dependencies by @matthchr in https://github.com/Azure/azure-service-operator/pull/4233
* Improve security context restrictions by @matthchr in https://github.com/Azure/azure-service-operator/pull/4242
* Add some validation of crd-pattern to asoctl export template by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4245
* Release lease when the pod is terminated by @matthchr in https://github.com/Azure/azure-service-operator/pull/4250
* Allow reuse of import from asoctl by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4243
* Allow number of simultaneous workers to be configured by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4257
* Reduce noise when importing PostgreSQL Flexible Server Configurations by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4279
* Reduce noise when importing MySQL Flexible Server Configurations by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4282
* Name spec and status for the associated resource by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4341
* Built-in role definitions should be skiped by asoctl by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4263
```

Example output for "Features" section:
```
### Features

* Allow the data encryption keys for PostgreSQL Flexible server to be configured via ConfigMap (#4311)
* Allow significantly more Helm chart configuration (#4211, #4207)
* Remove pre-upgrade check (#4217)
* Updated numerous dependencies
* Improve controller pod security context restrictions (#4242)
* Release lease when the pod is terminated for faster pod shutdown (#4250)
* Improve the spec and status names for some resources (breaking for Go pkg consumers) (#4341)
* asoctl: Add some validation of crd-pattern to `export template` command (#4245)
* asoctl: Allow reuse of import command via Go code (#4243)
* asoctl: Allow number of simultaneous workers to be configured (#4257)
* asoctl: Reduce noise when importing PostgreSQL Flexible Server Configurations (#4279)
* asoctl: Reduce noise when importing MySQL Flexible Server Configurations (#4282)
* asoctl: Built-in role definitions should be skiped by import command (#4263)
```

Example input for "Bug fixes" section:

```
### Bug fixes

* Fix panic in log by @matthchr in https://github.com/Azure/azure-service-operator/pull/4249
* Ensure resource names are always K8s compliant by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4244
* Allow importing of FrontDoor resources by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4262
* Fix deadlock and ux bugs in asoctl by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4270
* Update ASO's NetworkPolicy to allow egress to SqlServer's default port by @super-harsh in https://github.com/Azure/azure-service-operator/pull/4283
* Retry if/when ScheduledQueryRules encounter BadRequest by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4312
* Relax enum restrictions on Storage bypass parameter by @matthchr in https://github.com/Azure/azure-service-operator/pull/4328
* Fix bug where some ConfigMap and Secret references could be missed by @matthchr in https://github.com/Azure/azure-service-operator/pull/4330
```

Example output for "Bug fixes" section:
```
### Bug fixes

* Fix bug where a log could cause a panic (#4249)
* Fix bug where ASO's NetworkPolicy didn't allow egress to SqlServer's default port (#4283)
* Fix bug where ScheduledQueryRules could get stuck reconciling when encountering a BadRequest (#4312)
* Fix bug where storage account bypass parameter did not accept all allowed values (#4328)
* Fix bug where some ConfigMap and Secret references could fail to be exported (#4330)
* asoctl: Fix bug where resource names could be generated which where not valid in Kubernetes (#4244)
* asoctl: Fix bug where FrontDoor resources couldn't be imported (#4262)
* asoctl: Fix deadlock and ux bugs (#4270)
```

Example input for "Documentation" section:

```
### Documentation

* Add missing SQL User to documentation index by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4230
* add docs for non-podidentity managed identity by @nojnhuh in https://github.com/Azure/azure-service-operator/pull/4227
* Add some additional documentation about metrics by @matthchr in https://github.com/Azure/azure-service-operator/pull/4273
* Update redis migration docs with correct redis key by @bwiens in https://github.com/Azure/azure-service-operator/pull/4297
* Add ArgoCD FAQ item by @theunrepentantgeek in https://github.com/Azure/azure-service-operator/pull/4309
```

Example output for "Documentation" section:
```
### Documentation

* Add missing SQL User documentation (#4230)
* Add docs for non-podidentity managed identity (#4227)
* Add some additional documentation about the controller metrics endpoint (#4273)
* Update redis migration docs with correct redis key (#4297)
* Add ArgoCD FAQ item (#4309)
```

Please format this document according to the above rules, and return the output as raw (unrendered) markdown:
```
Initial release notes here
```
