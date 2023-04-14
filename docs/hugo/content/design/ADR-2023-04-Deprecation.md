---
title: '2023-04: Resource Deprecation'
---

## Context

As succinctly raised by `@sakthi-vetrivel` in (#2070)[https://github.com/Azure/azure-service-operator/issues/2070]:

> Today, there's no change to the support for an Azure Service Operator CRD when the underlying Azure service/offering is being deprecated. How should we communicate these deprecations and what timelines should we follow? How will we keep our catalog up to date?

Starting in Kubernetes v1.19, it's possible to flag individual versions of a custom resource as (deprecated)[https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/#version-deprecation]. This results in the follow behaviour:

> When API requests to a deprecated version of that resource are made, a warning message is 
> returned in the API response as a header. The warning message for each deprecated version of the 
> resource can be customized if desired.

## Decision

To make deprecations visible, we will modify our ObjectModelConfiguration to allow flagging a particular resource (or resource version) as deprecated (including a message), with this flowing through into the generated code and our custom resource definitions. 

When we become aware of resource or version deprecation, we'll do the following:

* Modify our configuration to indicate deprecation of the resource (or resource version) and the next release of Azure Service Operator will include that information.
* Include details on the deprecation in our release notes for that release of ASO
* List the deprecation as a future breaking change

After the Azure deprecation has taken effect (that is, the resource or version is no longer deployable), we'll remove it from the following ASO release, with the breaking change noted in both the release notes and on our website.

## Status

TBC


## Consequences

TBC


## Experience Report

TBC


## References

TBC

