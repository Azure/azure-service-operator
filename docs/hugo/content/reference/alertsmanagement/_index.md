---
title: AlertsManagement Supported Resources
linktitle: AlertsManagement
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `alertsmanagement.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource                                                                                                                                                                     | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                                |
|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|---------------|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| [PrometheusRuleGroup](https://azure.github.io/azure-service-operator/reference/alertsmanagement/v1api20230301/#alertsmanagement.azure.com/v1api20230301.PrometheusRuleGroup) | 2023-03-01  | v1api20230301 | v2.8.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/alertsmanagement/v1api20230301/v1api20230301_prometheusrulegroup.yaml)    |
| SmartDetectorAlertRule                                                                                                                                                       | 2021-04-01  | v1api20210401 | v2.11.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/alertsmanagement/v1api20210401/v1api20210401_smartdetectoralertrule.yaml) |

