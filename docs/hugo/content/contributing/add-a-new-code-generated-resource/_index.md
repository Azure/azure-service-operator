---
title: Adding a new code generated resource to ASO v2
linktitle: Add a resource
layout: single
---

Want to add a new code-generated resource to Azure Service Operator v2? You're in the right place - here's your step by step guide. We've broken the process down, and we'll walk you through each step.

* [**Before you begin**]({{< relref "before-you-begin" >}}) sets the context for the rest of the process. You begin by identifying the resource you want to add and preparing your development environment.

* [**Generating code**]({{< relref "run-the-code-generator" >}}) details how to configure and run the code generator that writes 95% of the code for you. 

* [**Review the generated resource**]({{< relref "review-the-generated-resource" >}}) gives you an overview of the generated code and how to identify common issues. 

* [**Write a test for the resource**]({{< relref "write-a-test" >}}) explains how to write a test to verify the resource can be created, updated, and deleted. This ensures that the resource works as expected and helps catch any issues early on.

* [**Create a sample**]({{< relref "create-a-sample" >}}) shows you how to create a sample that demonstrates how to use the resource, and how to verify the sample works by recording it in use.

* [**Final checks**]({{< relref "final-checks" >}}) walks you through the final checks to ensure everything is working as expected before you create your pull request (PR).

We usually find new contributors are able to work through this process over the course of a few days. Experienced contributors (we could ourselves in this group) have completed some resources in just a few hours. 

If you run into any issues, create a draft PR with your progress to this point and we're happy to help you move forward. 

