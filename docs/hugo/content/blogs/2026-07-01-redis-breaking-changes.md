---
title: "Breaking Changes for Redis Enterprise"
date: 2026-07-01
description: "Upcoming breaking changes for Redis Enterprise"
type: blog
---

If you manage `RedisEnterprise` or `RedisEnterpriseDatabase` resources with ASO, this post is for you. In **ASO v2.21** (scheduled for August 2026), we will remove the `v1api20210301` and `v1api20230701` versions of both resources. You'll need to move your YAML to a newer API version before upgrading.

This change is driven by an upstream Azure decision: _Azure Cache for Redis Enterprise_ is being retired in favour of _Azure Managed Redis_, and the older ARM API versions cannot manage the new product.

## Why this is happening

Microsoft has [announced the retirement](https://learn.microsoft.com/en-us/azure/azure-cache-for-redis/retirement-faq#questions-on-the-enterprise-tier-of-azure-cache-for-redis-retirement) of Azure Cache for Redis Enterprise and Enterprise Flash tiers on **March 31, 2027**. All Enterprise instances will be disabled on April 1, 2027. The recommended migration path is to [Azure Managed Redis](https://learn.microsoft.com/en-us/azure/redis/migrate/migrate-redis-enterprise-overview), an Azure first-party service built on the same Redis Enterprise software stack.

At the ARM and Bicep level, the difference between the two products is largely a SKU change. However, the SKUs required for Azure Managed Redis are only present in newer versions of the `Microsoft.Cache/redisEnterprise` API. The older versions ASO ships (`2021-03-01` and `2023-07-01`) cannot create or manage Azure Managed Redis instances at all, and would leave you stuck on a retiring product.

Removing these versions from ASO ensures that:

- Existing users are prompted to move to a supported API version well ahead of the Azure retirement deadline.
- All ASO-managed `RedisEnterprise` resources can be migrated to Azure Managed Redis without changing the resource `Kind` or losing ownership of dependent resources.
- We're not implicitly directing new users to a retiring SKU.

## What you need to do

If you have `RedisEnterprise` or `RedisEnterpriseDatabase` resources using `cache.azure.com/v1api20210301` or `cache.azure.com/v1api20230701` in your cluster, you must update them to a newer version before upgrading to ASO v2.21.

We recommend `v20250401` for new and migrated resources.

{{% alert title="Note" %}}
Changing the ASO `apiVersion` only updates how ASO represents the resource. To actually move your data to Azure Managed Redis, follow the [Azure migration guide](https://learn.microsoft.com/en-us/azure/redis/migrate/migrate-redis-enterprise-overview).
{{% /alert %}}

## Choosing an Azure Managed Redis SKU

Azure Managed Redis replaces the `Enterprise_` and `EnterpriseFlash_` SKUs with a simpler structure based on memory and performance tier. See the Azure documentation for the full SKU list and sizing guidance:

- [What is Azure Managed Redis?](https://learn.microsoft.com/en-us/azure/redis/overview)
- [Azure Managed Redis architecture](https://learn.microsoft.com/en-us/azure/redis/architecture)
- [Migrate from Redis Enterprise to Azure Managed Redis](https://learn.microsoft.com/en-us/azure/redis/migrate/migrate-redis-enterprise-overview)

Once you've chosen an appropriate SKU, set it in the `spec.sku.name` field of your `RedisEnterprise` resource.
