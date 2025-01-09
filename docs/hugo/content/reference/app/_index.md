---
title: App Supported Resources
linktitle: App
no_list: true
---
To install the CRDs for these resources, your ASO configuration must include `app.azure.com/*` as a one of the configured CRD patterns. See [CRD Management in ASO](https://azure.github.io/azure-service-operator/guide/crd-management/) for details on doing this for both [Helm](https://azure.github.io/azure-service-operator/guide/crd-management/#helm) and [YAML](https://azure.github.io/azure-service-operator/guide/crd-management/#yaml) based installations.

### Next Release

Development of these new resources is complete and they will be available in the next release of ASO.

| Resource     | ARM Version | CRD Version   | Supported From | Sample                                                                                                                         |
|--------------|-------------|---------------|----------------|--------------------------------------------------------------------------------------------------------------------------------|
| AuthConfig   | 2024-03-01  | v1api20240301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/app/v1api20240301/v1api20240301_authconfig.yaml)   |
| ContainerApp | 2024-03-01  | v1api20240301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/app/v1api20240301/v1api20240301_containerapp.yaml) |
| Job          | 2024-03-01  | v1api20240301 | v2.12.0        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/app/v1api20240301/v1api20240301_job.yaml)          |

### Released

These resource(s) are available for use in the current release of ASO. Different versions of a given resource reflect different versions of the Azure ARM API.

| Resource           | ARM Version | CRD Version   | Supported From | Sample                                                                                                                               |
|--------------------|-------------|---------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------|
| ManagedEnvironment | 2024-03-01  | v1api20240301 | 2.12.0         | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/samples/app/v1api20240301/v1api20240301_managedenvironment.yaml) |

