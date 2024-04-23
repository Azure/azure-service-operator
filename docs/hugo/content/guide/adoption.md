---
title: Adopting existing Azure resources
linktitle: Adoption
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

## What is adoption?

Adoption is when Azure Service Operator starts managing a resource which already exists in Azure but was 
not previously being managed by ASO.

**ASO adopts resources by default.** For example, if you create a Storage Account named "mystorageaccount" in resource
group "myrg" in subscription "dev1" via ASO, but a storage account with that name _already exists_ in that resource
group in that subscription, then ASO will automatically update the existing storage account configuration to match
what you specified in the `StorageAccount` resource in your cluster. Subsequent changes to the `StorageAccount` resource will be
propagated to Azure as normal, and changes to the storage account via ad-hoc mechanisms such as the portal
will be lost as ASO reconciles and manages the resource.

For more details on adoption and why ASO chose to do what it does, see
[the original adoption design]( {{< relref "adr-2023-02-adoption-policy" >}} ).

## Types of adoption.

- Adopt and manage: This is the default experience, when ASO finds an existing resource in Azure it immediately
adopts it and starts managing it. Deleting the resource in ASO will delete the underlying resource in Azure 
as well.
- Adopt and manage, but do not delete: Configured by setting the 
[serviceoperator.azure.com/reconcile-policy]( {{< relref "annotations#serviceoperatorazurecomreconcile-policy" >}} )
to `detach-on-delete` before creating the resource.

## Helpful tools

If you want to take an existing Azure resource and adopt it _as is_ into ASO, you can painstakingly recreate its
existing configuration in ASO, or you can use the
[`asoctl import azure-resource`]( {{< relref "asoctl#import-azure-resource" >}} ) command.