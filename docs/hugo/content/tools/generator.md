---
title: "generator"
linkTitle: "generator"
weight: 20
menu:
  main:
    weight: 20
layout: single
cascade:
- type: docs
- render: always
description: Azure Service Operator Code Generator
---

``` bash
$ generator
aso-gen provides a cmdline interface for generating Azure Service Operator types from Azure deployment template schema

Usage:
  aso-gen [command]

Available Commands:
  completion    Generate the autocompletion script for the specified shell
  gen-kustomize generate K8s Kustomize file in the spirit of Kubebuilder, based on the specified config folder
  gen-types     generate K8s resources from Azure deployment template schema
  help          Help about any command
  version       Display version information
```

