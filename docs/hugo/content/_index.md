---
title: Azure Service Operator v2
type: docs
description: "Manage your Azure resources from within your Kubernetes cluster."
---

<img src="https://azure.github.io/azure-service-operator/favicons/favicon-128.png" style="float:left; margin: -8px 8px 8px 0px;"/>Azure Service Operator (ASO) allows you to deploy and maintain a wide variety of Azure Resources using the Kubernetes tooling you already know and use.

Instead of deploying and managing your Azure resources separately from your Kubernetes application, ASO allows you to manage them together, automatically configuring your application as needed. For example, ASO can set up your [Redis Cache](https://azure.github.io/azure-service-operator/reference/cache/) or [PostgreSQL database server](https://azure.github.io/azure-service-operator/reference/dbforpostgresql/) and then configure your Kubernetes application to use them.

## Project Status

This project is stable. We follow the [Kubernetes definition of stable](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/#feature-stages).

## Why use Azure Service Operator v2?

- **K8s Native:** we provide CRDs and Golang API structures to deploy and manage Azure resources through Kubernetes.
- **Azure Native:** our CRDs understand Azure resource lifecycle and model it using K8s garbage collection via ownership references.
- **Cloud Scale:** we generate K8s CRDs from Azure Resource Manager schemas to move as fast as Azure.
- **Async Reconciliation:** we don't block on resource creation.

## What resources does ASO v2 support?

ASO supports more than 150 different Azure resources, with more added every release. See the full list of [supported resources](https://azure.github.io/azure-service-operator/reference/).

## Contact us

If you've got a question, a problem, a request, or just want to chat, here are two ways to get in touch:

* Log an issue on our [GitHub repository](https://github.com/azure/azure-service-operator).
* Join us over on the [Kubernetes Slack](https://kubernetes.slack.com) in the [#azure-service-operator](https://kubernetes.slack.com/archives/C046DEVLAQM) channel. (Don't have an account? [Sign up here](https://slack.k8s.io/).)

## How to contribute

To get started developing or contributing to the project, follow the instructions in the [contributing guide](https://azure.github.io/azure-service-operator/contributing/).
