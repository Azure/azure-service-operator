---
title: Resource Versioning
linktitle: Resource Versioning
weight: -1
---

## How ASO versions resources

Each ASO resource version corresponds directly to an Azure API version. For example, `resources.azure.com/v1api20200601`
maps to the `2020-06-01` version of the Azure Resource Manager API. When multiple API versions are available for a
resource, ASO offers a separate CRD version for each one.

Internally, Kubernetes stores every resource using a single canonical **storage version**, regardless of which API
version you use when creating or updating it. ASO handles conversion between versions automatically, so you can
interact with a resource using any supported version and ASO takes care of the rest.

## Which API version does ASO use when talking to Azure?

ASO tracks the API version you last used when applying the resource and uses that version when communicating
with Azure. This matters because different Azure API versions can have different behavior and different required fields.

For example, if you create a resource using `v1api20200601`, ASO sends requests to Azure using the `2020-06-01` API.
If you later re-apply the resource using `v1api20220601`, ASO switches to using the `2022-06-01` API going forward.

## How to change the API version of a resource

To move a resource to a newer (or different) API version, update the `apiVersion` in your YAML manifest and apply it
with `kubectl apply`.

For example, to upgrade a `ResourceGroup` from `v1api20200601` to `v1api20240301`:

```yaml
apiVersion: resources.azure.com/v1api20240301
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
spec:
  location: westus2
```

```bash
$ kubectl apply -f resourcegroup.yaml
```

After the apply, ASO begins using the new API version when communicating with Azure.

{{% alert title="Note" %}}
You cannot change the `apiVersion` of a resource using `kubectl edit`. Kubernetes does not allow modifications to
`apiVersion` or `kind` through edit. You must use `kubectl apply` (or similar) with an updated manifest instead.
{{% /alert %}}
