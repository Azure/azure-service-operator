# Azure Service Operator Quota Sample

This directory contains the Quota sample for Azure Service Operator v2.

## Sample Structure

```
v2/samples/quota/v1api20250901/
├── v1api20250901_quota.yaml          # Main Quota sample
└── refs/
    └── v1api20200601_resourcegroup.yaml  # ResourceGroup dependency
```

## Sample Content

The Quota sample demonstrates how to create a quota for Azure resources using the standard ASO pattern:

```yaml
apiVersion: quota.azure.com/v1api20250901
kind: Quota
metadata:
  name: aso-sample-quota
  namespace: default
spec:
  owner:
    name: aso-sample-rg
    group: resources.azure.com
    kind: ResourceGroup
  properties:
    name:
      value: "cores"
    resourceType: "cores"
    limit:
      limitValue:
        limitObjectType: "LimitValue"
        limitType: "Independent"
        value: 100
```

## Key Features

- **Standard Reference Pattern**: Uses `name`/`group`/`kind` for owner references
- **Dependency Management**: ResourceGroup reference included in `refs/` folder  
- **ASO Conventions**: Follows established patterns used by other ASO samples
- **Test Ready**: Compatible with existing samples_tester infrastructure

## Testing

To test this sample (requires Azure credentials):

```bash
AZURE_SUBSCRIPTION_ID=84bee13f-1f0f-4ba4-9fed-240486dd4710 \
AZURE_TENANT_ID=<your-tenant-id> \
TEST_FILTER=Test_Samples_CreationAndDeletion/Test_Quota_v1api20250901_CreationAndDeletion \
task controller:test-integration-envtest
```

## Implementation Notes

This sample was updated to follow reviewer feedback and ASO best practices:

1. **Original Issue**: Used direct ARM references causing test failures
2. **Solution**: Changed to standard Kubernetes references with proper dependencies  
3. **Benefits**: Simpler, more maintainable, follows ASO conventions
