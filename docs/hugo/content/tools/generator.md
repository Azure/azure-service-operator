---
title: "generator"
linkTitle: "generator"
weight: 20
layout: single
cascade:
- type: docs
- render: always
description: Azure Service Operator Code Generator
---

The Azure Service Operator code generator is useful for **contributors** to ASO. Users do not need to run the generator themselves.

For more details on how to contribute to ASO and use the generator, see 
[adding a new code generated resource]({{< relref "add-a-new-code-generated-resource" >}}).

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

