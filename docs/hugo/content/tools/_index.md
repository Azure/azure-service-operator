---
title: "Tools"
linkTitle: "Tools"
weight: 30
menu:
  main:
    weight: 30
layout: single
cascade:
- type: docs
- render: always
description: Additional tooling for Azure Service Operator
---

In addition to the ASO operator itself, there are other tools that are useful for working with ASO.

## asoctl

A command line swiss-army-knife that provides a number of useful commands for working with ASO. It can be used to:

* Clean your cluster of obsolete custom resources prior to an upgrade
* Convert an existing Azure resource into a custom-resource (nearly) ready for deployment

## generator

We have a sophisticated code generator that helps us add support for new resources (and new versions of existing resources). It consumes the Azure OpenAPI/Swagger specifications and writes most of the required code.

