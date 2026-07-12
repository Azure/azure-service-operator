---
title: "v2.20.0 Breaking Changes"
linkTitle: "v2.20.0"
weight: -75  # This should be 5 lower than the previous breaking change document
---
## Future breaking changes

### Azure Cache for Redis Enterprise (v2.21)

Azure has retired _Azure Cache for Redis Enterprise_ in favour of _Azure Managed Redis_.

While the difference in the ARM/Bicep level is just another SKU, older resource versions cannot be used to create or manage the newer product, requiring all users to upgrade.

ASO will be removing versions `v1api20210301` and `v1api20230701` of `RedisEnterprise` and `RedisEnterpriseDatabase` in version v2.21, scheduled for August 2026.

Users are advised to move their resources to `v20250401` for new and migrated ASO resources as soon as practical, and to plan for migration to Azure Managed Redis well before the deadline of March 31, 2027.
