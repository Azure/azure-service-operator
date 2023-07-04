---
title: "User’s Guide"
linktitle: "Guide"
weight: 10
menu:
  main:
    weight: 10
cascade:
- type: docs
description: "How to use Azure Service Operator v2 to manage your Azure resources from within your Kubernetes Cluster"
---

## Installation and setup


[Helm](../../#installation) is the preferred way to install and manage Azure Service Operator. 

If you prefer installing things manually you can do so [using YAML]( {{< relref "installing-from-yaml" >}}). 

When it comes time to [Upgrade]( {{< relref "upgrading" >}}) to the next release, be sure to check for any [breaking changes]( {{< relref "breaking-changes" >}}). 

If you need to, you can [remove ASO]( {{< relref "uninstalling" >}}) from your cluster.

## Configuration and monitoring

[Global configuration]( {{< relref "aso-controller-settings-options" >}}) of ASO requires a secret in the `azureserviceoperator-system` namespace called `aso-controller-settings`.

[Authentication]( {{< relref "authentication" >}}) can be controlled on a per-namespace


Each ASO resource has a standard [Condition]( {{< relref "conditions" >}}) revealing its current status.


The operation of ASO can be controlled by applying [annotations]( {{< relref "annotations" >}}) to your resources. 

[CRD Management]( {{< relref "crd-management" >}})



[FAQ]( {{< relref "frequently-asked-questions" >}})

[Diagnosing problems]( {{< relref "diagnosing-problems" >}})

[Handling secrets]( {{< relref "secrets" >}})
[Metrics]( {{< relref "metrics" >}})
[Sharing data through ConfigMaps]( {{< relref "configmaps" >}})
