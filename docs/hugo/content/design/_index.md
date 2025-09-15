---
title: "Design & Specifications"
linktitle: "Design"
weight: 60
menu:
  main:
    weight: 60
layout: single
cascade:
- type: docs
description: "Design discussions, decisions, and specifications for Azure Service Operator v2"
---
Significant changes to Azure Service Operator should be captured by an ADR (architecture design record) in this
folder. ADRs are listed reverse chronologically by year. [See below](#format) for the structure each ADR uses.

## In progress

These ADRs reflect work that is being discussed, has been proposed, or is currently in progress.

### 2025

| Title                                                                       | Description                                                                                                                                                 |
| --------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [**Version Priority**]({{< relref "ADR-2025-05-Version-Priority" >}})       | A design to address the problems with version priority in Azure Service Operator, ensuring that the most recent resource version is selected by Kubernetes. |
| [**Improving ARM References**]({{< relref "ADR-2025-01-ARM-References" >}}) | Options for remedying the problems we face when we fail to identify a property as an ARM reference prior to releasing ASO.                                  |

### 2024

| Title                                                                                              | Description                                                                                                  |
| -------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------ |
| [**OperatorSpec Custom Properties**]({{< relref "ADR-2024-05-OperatorSpec-Custom-Properties" >}})  | Injecting additional properties into OperatorSpec for additional flexibility and extensibility of resources. |
| [**RoleAssignments UUID Generation**]({{< relref "ADR-2024-04-RoleAssignment-UUID-Generation" >}}) | Allowing users more control over how the GUIDs used for naming RoleAssignments are generated.                |
| [**Upstream Deletion**]({{< relref "ADR-2024-02-Upstream-Deletion" >}})                            | When might deletion of an upstream resource occur and how will ASO handle it.                                |

### 2023

| Title                                                                   | Description                                                                                                                                                                                                                                                                    |
| ----------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [**Adoption Policy**]({{< relref "ADR-2023-02-Adoption-Policy" >}})     | Understand the policy for adopting pre-existing Azure resources in Kubernetes through the Azure Service Operator, including the principles guiding resource adoption and the user's ability to control operator actions.                                                       |
| [**Deprecation**]({{< relref "ADR-2023-04-Deprecation" >}})             | Understand the policy for handling Azure resource deprecation in Azure Service Operator, including the process for flagging deprecated resources, communicating changes, and removing unsupported resources.                                                                   |
| [**Patch Collections**]({{< relref "ADR-2023-04-Patch-Collections" >}}) | Address the issue of resource drift in Azure Service Operator due to Azure Resource Providers' PATCH-like behavior on PUT, and discuss potential solutions such as optional properties and explicit nulls to ensure the custom resource represents the complete desired state. |

### 2022

| Title                                                                 | Description                                                                                                                                                                                                                                                                      |
| --------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [**Change Detection**]({{< relref "ADR-2022-11-Change-Detection" >}}) | The challenges of change detection in Azure Service Operator, discussing the limitations of relying on Azure Resource Manager for goal state comparison, the impact of ARM throttling, and a proposal for ASO to perform its own goal state comparison to mitigate these issues. |

### Older

| Title                                                                               | Description                                                                                                                                                                                                                                          |
| ----------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [**CrossPlane**]({{< relref "crossplane" >}})                                       | Discusses the intricacies of code generation for Crossplane, including static "special" properties, cross resource references, and the challenges associated with generating these references.                                                       |
| [**Custom validation and defaulting**]({{< relref "annotations" >}})                | Delve into the proposal for custom validation and defaulting for code-generated resources, discussing the limitations of controller-runtime interfaces and suggesting a structure for autogenerated webhooks that allows for custom implementations. |
| [**Improving the Reconciler interface**]({{< relref "reconcile-interface" >}})      | Explore a proposal for improving the Reconciler interface in controller-runtime, addressing issues of code duplication and potential bugs due to differences in checks for Create vs Delete operations.                                              |
| [**Type References and Ownership**]({{< relref "type-references-and-ownership" >}}) | type references and ownership in Azure Service Operator, discussing goals such as idiomatic expression of Azure resource relationships, automatic ownership, garbage collection, and interaction with non-Kubernetes managed Azure resources.        |

## Completed Changes

These ADRs reflect work that has been completed.

### 2023

| Title                                                                                            | Description                                                                                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [**Conversion Augmentation**]({{< relref "ADR-2023-01-Conversion-Augmentation" >}})              | Learn about the need for conversion augmentation in Kubernetes resource versions, the role of code generation, and the strategies for augmenting these conversions with hand-written code to ensure backward and forward compatibility.                    |
| [**Goal Seeking KeyVaults**]({{< relref "ADR-2023-07-Goal-Seeking-KeyVaults" >}})                | Understand the challenges posed by Azure KeyVault's soft-delete feature in a goal-seeking system like Azure Service Operator, and the need for a solution that prevents resources from entering permanently degraded states without manual intervention.   |
| [**Helm Chart Size Limitations**]({{< relref "ADR-2023-02-Helm-Chart-Size-Limitations" >}})      | Address the Helm chart size limitations causing problems for Azure Service Operator, discussing the implications of Kubernetes secret size limits, the impact on chart installation, and potential workarounds such as splitting the Helm chart.           |
| [**Installing Only Selected CRDs**]({{< relref "ADR-2023-05-Installing-Only-Selected-CRDs" >}})  | Understand the need for selective CRD installation in Azure Service Operator to mitigate kube-apiserver memory usage, and the requirements for a seamless upgrade experience and default configuration that fits into minimal memory environments          |
| [**Package References**]({{< relref "ADR-2023-07-PackageReferences" >}})                         | Consider the proposal for introducing a SubPackageReference in the Azure Service Operator code generator, discussing the implications for code that consumes PackageReference and the potential impact on namespace pollution and code organization.       |
| [**Post-Reconciliation Extensions**]({{< relref "ADR-2023-03-Post-Reconciliation-Extension" >}}) | Understand the introduction of the PostReconciliationChecker extension in Azure Service Operator to handle post-reconciliation operations, allowing for custom conditions to be set on resources based on the status received from Azure Resource Manager. |
| [**Skipping Properties**]({{< relref "ADR-2023-09-Skipping-Properties" >}})                      | A design to address the problems that occur when a property skips multiple resource versions and has a different shape when it returns.                                                                                                                    |

### 2022

| Title                                                                                                                         | Description                                                                                                                                                                                                                                                              |
| ----------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [**Backward Compatibility**]({{< relref "ADR-2022-02-Backward-Compatibility" >}})                                             | As a part of planning for our first beta release consider the implications of modifying the version prefix for code-generated resources in Azure Service Operator.                                                                                                       |
| [**Evil Discriminator**]({{< relref "ADR-2022-08-Evil-Discriminator" >}})                                                     | Understand the challenges posed by the 'evil discriminator' in Swagger/OpenAPI, a term coined by the Azure Service Operator team to describe the differences in defining polymorphic API payloads in Swagger compared to JSON Schema.                                    |
| [**Muliple Credential Operator**]({{< relref "ADR-2022-09-Multiple-Credential-Operator" >}})                                  | proposed support for multiple credentials under a global operator in Azure Service Operator, enabling management of Azure resources across multiple subscriptions, with a detailed explanation of the credential selection hierarchy and caching mechanism.              |
| [**Reading Status properties from other ASO resources**]({{< relref "ADR-2022-09-Reading-Status-Properties-Of-Resources" >}}) | discussion of various ways to allow one ASO resource to consume information generated by another, with details on how we decided to use configmaps to do this.                                                                                                           |
| [**Reconciler Extensions**]({{< relref "ADR-2022-01-Reconciler-Extensions" >}})                                               | Understand the need for reconciler extensions in Azure Service Operator to manage inconsistencies in Azure Resource Providers' behavior, and the proposal for precise, per-resource customization without duplicating existing functionality or risking inconsistencies. |
| [**Resource Import**]({{< relref "ADR-2022-11-Resource-Import" >}})                                                           | Review the strategies for migrating from Azure Service Operator v1 to v2, including version integration and resource import job, each with its pros, cons, and potential blockers.                                                                                       |
| [**Reconciliation Extensions**]({{< relref "ADR-2022-12-Reconciliation-Extensions" >}})                                       | Understand the need for reconciliation extensions in Azure Service Operator to handle resources in transient states that preclude successful PUT operations, and the proposal for an extension point to modify the reconciliation flow accordingly.                      |

### 2021

| Title                                                                                   | Description                                                                                                                                                                                                                                                   |
| --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [**API Version Recovery**]({{< relref "ADR-2021-06-API-Version-Recovery" >}})           | Identity the challenges of API version recovery in Azure Service Operator, discussing the importance of using the correct Azure API version for expected behavior, and the need to pivot back to the original version for generating the correct ARM payload. |
| [**Committing Generated Code**]({{< relref "ADR-2021-08-Committing-Generated-Code" >}}) | Uncover the reasons and implications of committing generated code in the Azure Service Operator, discussing the benefits of traceability, easier debugging, and the ability to review generated changes.                                                      |
| [**Property Conversions**]({{< relref "ADR-2020-04-Code-Generation" >}})                | Explore the need for codegen conversion routines in handling property conversions across different versions of a resource, with examples of how primitive types can evolve over time.                                                                         |

### 2020

| Title                                                                                         | Description                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [**Abstract Syntax Tree Library Choice**]({{< relref "ADR-2020-11-Abstract-Syntax-Trees" >}}) | Background on the decision to adopt the `dst` (Decorated Syntax Tree) library in the Azure Service Operator Code Generator, addressing limitations of the standard Go `ast`` library such as unexpected comment placements, non-compliance with go fmt standards, and poor control over whitespace. |
| [**Pipeline Architecture**]({{< relref "ADR-2020-07-Pipeline-Architecture" >}})               | The pipeline architecture of the ASO code generator, its design, and how it contributes to the overall functionality and efficiency of the system.                                                                                                                                                  |
| [**Why Code Generation?**]({{< relref "ADR-2020-04-Code-Generation" >}})                      | Rationale and impact of adopting code generation in the Azure Service Operator, including the benefits of full feature coverage, reduced overhead for new resources, and consistent behavior across all resources.                                                                                  |

### Earlier

| Title                                                                           | Description                                                                                                                                                                                                                                                    |
| ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [**Clarifying Object Structure**]({{< relref "clarifying-object-structure" >}}) | Dive into the challenges of differentiating Azure-specific properties from operator-specific properties in the object structure of resources, with a focus on potential solutions for Status and upcoming properties like SecretConfiguration and Credentials. |
| [**Managing dataplane secrets**]({{< relref "secrets" >}})                      | Managing dataplane secrets in Azure Service Operator, including the drawbacks of auto-generating secrets and the implications for secret rollover, control, security, resource adoption, and GitOps compatibility.                                             |
| [**Reporting Resource Status**]({{< relref "resource-states" >}})               | A proposal for reporting resource status in Azure Service Operator, discussing the distinction between operator status and Azure resource state, and the current state and limitations of both handcrafted and auto-generated resources.                       |

## Format

Each ADR should contain the following sections:

* Context: Context for the problem. What are related issues/problems? What is the history?
* Requirements: A short list of goals
* Options optional section containing 1 or more options, w/ pros + cons. Can omit if there is only one real option.
  * OR Category + Options: For more complex ADRs where there are multiple facets to the design that are somewhat
    orthogonal, you can instead give categories and then define options within those categories
* Decision: Which option did you choose and why?
  * FAQ: Can optionally include an FAQ in this section answering common questions about the decision that weren't
    already answered in the options section.
* Status
* Consequences
* Experience Report
* References

Ideally, each ADR should run to one or two pages in length. You can use the
[simple template]({{< relref "ADR-(YEAR)-(MONTH)-(TITLE)" >}}).

For complex changes that require more detail, the ADR should provide an overview. Additional detail (say, to list
alternative solutions, detail complex edge cases, or present detailed diagrams) should be captured
in a separate design document linked from the ADR.

For background information, check out [this Cognitect blog entry](https://www.cognitect.com/blog/2011/11/15/documenting-architecture-decisions).

ADR documents should be updated over time to keep them relevant, typically by updating the *Experience Report* section.
