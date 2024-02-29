---
title: 2024-02 Upstream Deletion
---

What do we do if/when a resource supported by ASO v2 is deleted upstream - removed from the Azure OpenAPI repository?

## Context

In July 2022, we encountered a problem when the preview versions of the `Microsoft.Auithorization` service were deleted from the [azure-resource-manager-schemas](https://github.com/Azure/azure-resource-manager-schemas) repository that ASO consumed at the time. This orphaned the `RoleAssignment` resource in ASO, which would have been a breaking change for our users.

Since that time, ASO has migrated to depend soley on the [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs) repo (see [PR#2323](https://github.com/Azure/azure-service-operator/pull/2323), merged December 2022) so the problem with `RoleAssignment` went away.

However, we decided that this question still needed to be considered, as it may come up again in the future and we need to have a plan in place for how to handle it.

Fortunately, the `azure-rest-api-specs` is very tightly controlled and is used as the source of truth for code generation of client APIs across multiple languages (including Python, Java, C# and Go) so it's vanishingly unlikely for a resource version to be removed without warning.

### Scenario A: Deletion of preview version

The most likely scenario is that an old preview version of a service is removed. Preview versions are only guaranteed to be supported for 90 days after creation, though in practice they are retained much longer.

> When releasing a new preview, the service team may completely retire any previous preview versions after giving customers at least 90 days to upgrade their code.
  -- [Microsoft Azure REST API Guidelines](https://github.com/microsoft/api-guidelines/blob/vNext/azure/Guidelines.md), retrieved 23 Feb 2024

In this situation, the resource provider will stop accepting the old version of the resource, requiring all users to update to a newer api version, including users of ASO.

ASO can support migration by including a later version of the resource in a new release.

If the new ASO release is available before the deprecation/deletion date, users can update their clusters to the new version, and then update their resources to the new version at any time. As long as they complete the update prior to deletion, no downtime will occur.

If the new ASO release is not available until after the depredcation/deletion date, users will need to export/recreate their resources as a part of the upgrade process.

In either case, the removal of the old version from ASO will be a breaking change that must be documented on the breaking release, and should be documented on earlier releases if possible.

### Scenario: Deletion of GA version

The deletion of a GA version of a resource is much rarer event, and requires a much longer notice and deprecation period.

Either the resource provider will stop accepting the old version of the resource, or it will be entirely shut down.

ASO should be aware of this change well in advance, and should be able to support migration by including a later version of the resource in a new release. We should also document the upcoming deprecation so that users know they need to migrate.

### Scenario: Deletion of an entire resource

Even rarer, this can happen when a service is retired or replaced. ASO should be aware of this change well in advance. Where possible, we should provide a migration path to the new service, and document the upcoming deprecation so that users know they need to migrate.

## Decision

Deletion inevitably means that attempts to use the API will fail because the backing resource provider will no longer accept the requests.

This is a breaking change no matter what ASO does, even if we found a way to work around our current inability to regenerate the resource.

ASO should work to be aware of upcoming deprecations, to document them where ASO users may be affected, and to provide migration paths to newer versions where possible.

We may want to consider building upgrade tooling into `asoctl` to make migration smoother.
