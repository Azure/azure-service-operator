---
title: Insights Supported Resources
linktitle: Insights
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `insights.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource           | ARM Version | CRD Version   | Supported From | Sample |
|--------------------|-------------|---------------|----------------|--------|
| ScheduledQueryRule | 2023-12-01  | v1api20231201 | v2.12.0        | -      |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                               | ARM Version        | CRD Version          | Supported From | Sample                                                                                                                                                 |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|----------------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| [ActionGroup](https://azure.github.io/azure-service-operator/reference/insights/v1api20230101/#insights.azure.com/v1api20230101.ActionGroup)                           | 2023-01-01         | v1api20230101        | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20230101_actiongroup.yaml)                             |
| [AutoscaleSetting](https://azure.github.io/azure-service-operator/reference/insights/v1api20221001/#insights.azure.com/v1api20221001.AutoscaleSetting)                 | 2022-10-01         | v1api20221001        | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20221001_autoscalesetting.yaml)                        |
| [Component](https://azure.github.io/azure-service-operator/reference/insights/v1api20200202/#insights.azure.com/v1api20200202.Component)                               | 2020-02-02         | v1api20200202        | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20200202_component.yaml)                               |
| [DiagnosticSetting](https://azure.github.io/azure-service-operator/reference/insights/v1api20210501preview/#insights.azure.com/v1api20210501preview.DiagnosticSetting) | 2021-05-01-preview | v1api20210501preview | v2.11.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api20210501preview/v1api20210501preview_diagnosticsetting.yaml) |
| [MetricAlert](https://azure.github.io/azure-service-operator/reference/insights/v1api20180301/#insights.azure.com/v1api20180301.MetricAlert)                           | 2018-03-01         | v1api20180301        | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20180301_metricalert.yaml)                             |
| [ScheduledQueryRule](https://azure.github.io/azure-service-operator/reference/insights/v1api20220615/#insights.azure.com/v1api20220615.ScheduledQueryRule)             | 2022-06-15         | v1api20220615        | v2.4.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20220615_scheduledqueryrule.yaml)                      |
| [Webtest](https://azure.github.io/azure-service-operator/reference/insights/v1api20220615/#insights.azure.com/v1api20220615.Webtest)                                   | 2022-06-15         | v1api20220615        | v2.7.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20220615_webtest.yaml)                                 |
| [Webtest](https://azure.github.io/azure-service-operator/reference/insights/v1api20180501preview/#insights.azure.com/v1api20180501preview.Webtest)                     | 2018-05-01-preview | v1api20180501preview | v2.0.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/insights/v1api/v1api20180501preview_webtest.yaml)                          |

