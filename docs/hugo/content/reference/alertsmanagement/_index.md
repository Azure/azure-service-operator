---
title: AlertsManagement Supported Resources
linktitle: AlertsManagement
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `alertsmanagement.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource            | ARM Version | CRD Version   | Supported From | Sample                                                                                                                                             |
|---------------------|-------------|---------------|----------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| PrometheusRuleGroup | 2023-03-01  | v1api20230301 | v2.8.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/alertsmanagement/v1api20230301/v1api20230301_prometheusrulegroup.yaml) |

