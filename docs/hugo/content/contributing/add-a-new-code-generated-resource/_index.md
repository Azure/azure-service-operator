---
title: Adding a new code generated resource to ASO v2
linktitle: Add a resource
layout: single
---

Want to add a new resource to Azure Service Operator v2? You're in the right place - here's your step by step guide.

## Quick Start

Adding a new resource is a multi-step process driven by our code generator. At a high level, the process is:

1.  **Configure**: Edit `v2/azure-arm.yaml` to add the new resource.
2.  **Generate**: Run `task generator:quick-checks` to generate the Go code. This will likely fail.
3.  **Fix & Repeat**: Fix errors reported by the generator by adding configuration to `v2/azure-arm.yaml`, then re-run the generator. Repeat until generation succeeds.
4.  **Test**: Write an integration test to verify the resource can be created, updated, and deleted.
5.  **Sample**: Create a sample YAML to demonstrate how to use the resource.
6.  **Verify**: Run `task ci` to ensure all checks pass before creating a pull request.

## Step by step guide

By the end of this process, you'll have created a new resource, written a test to verify it works, and created a sample to demonstrate how to use it. This sounds like a lot of work, but most of it is done for you by our code generator.

1. [**Before you begin**]({{< relref "before-you-begin" >}}) sets the context for the rest of the process. You begin by identifying the resource you want to add and preparing your development environment.

2. [**Generating code**]({{< relref "run-the-code-generator" >}}) details how to configure and run the code generator that writes 95% of the code for you.

3. [**Review the generated resource**]({{< relref "review-the-generated-resource" >}}) gives you an overview of the generated code and how to identify common issues.

4. [**Implement extensions**]({{< relref "implement-extensions" >}}) explains how to customize the generated resource to behave as expected. Most resources will work out of the box, but some may require additional customization.

5. [**Write a test for the resource**]({{< relref "write-a-test" >}}) explains how to write a test to verify the resource can be created, updated, and deleted. This ensures that the resource works as expected and helps catch any issues early on.

6. [**Create a sample**]({{< relref "create-a-sample" >}}) shows you how to create a sample that demonstrates how to use the resource, and how to verify the sample works by recording it in use.

7. [**Final checks**]({{< relref "final-checks" >}}) walks you through the final checks to ensure everything is working as expected before you create your pull request (PR).

We usually find new contributors are able to work through this process over the course of a few days. Experienced contributors (we count ourselves in this group) have completed some resources in just a few hours.

Progression isn't necessarly linear through these steps. It's common, for example, to identify issues while writing a test that require changing the configuration of the code generator, or writing a customization.

If you run into any issues, create a draft PR with your progress to this point and we're happy to help you move forward.

[Let's get started!]({{< relref "before-you-begin" >}})
