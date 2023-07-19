---
title: Insights Supported Resources
linktitle: Insights
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `insights.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                           | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                        |
|----------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|-------------------------------------------------------------------------------------------------------------------------------|
| [Component](https://azure.github.io/azure-service-operator/reference/insights/v1api20200202/#insights.azure.com/v1api20200202.Component)           | 2020-02-02         | v1api20200202        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20200202_component.yaml)      |
| [Webtest](https://azure.github.io/azure-service-operator/reference/insights/v1api20180501preview/#insights.azure.com/v1api20180501preview.Webtest) | 2018-05-01-preview | v1api20180501preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20180501preview_webtest.yaml) |

### Deprecated

These resource versions are deprecated and will be removed in an upcoming ASO release. Migration to newer versions is advised. See [Breaking Changes](https://azure.github.io/azure-service-operator/guide/breaking-changes/) for more information.

| Resource                                                                                                                                             | ARM Version        | CRD Version           | Supported From | Sample                                                                                                                          |
|------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|-----------------------|----------------|---------------------------------------------------------------------------------------------------------------------------------|
| [Component](https://azure.github.io/azure-service-operator/reference/insights/v1beta20200202/#insights.azure.com/v1beta20200202.Component)           | 2020-02-02         | v1beta20200202        | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1beta/v1beta20200202_component.yaml)      |
| [Webtest](https://azure.github.io/azure-service-operator/reference/insights/v1beta20180501preview/#insights.azure.com/v1beta20180501preview.Webtest) | 2018-05-01-preview | v1beta20180501preview | v2.0.0-beta.0  | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1beta/v1beta20180501preview_webtest.yaml) |

