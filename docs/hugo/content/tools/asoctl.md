---
title: "asoctl"
linkTitle: "asoctl"
weight: 10
menu:
  main:
    weight: 10
layout: single
cascade:
- type: docs
- render: always
description: Azure Service Operator Controller
---

``` bash
$ asoctl
asoctl provides a cmdline interface for working with Azure Service Operator

Usage:
  asoctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  crd         Custom Resource Definition (CRD) related actions
  help        Help about any command
  import      imports ARM resources as YAML resource definitions
  version     Display version information
```

## Clean CRDs

This command can be used to prepare ASOv2 `v1alpha1api` CustomResources for upgrading to a version `v1api` for `2.0.0` release. 
It ensures that any ASOv2 `v1alpha1api` deprecated version resources that may have been stored in etcd get migrated to `v1api`. 

```bash
$ asoctl crd clean --help
Clean deprecated CRD versions from cluster

Usage:
  asoctl crd clean [flags]

Flags:
      --dry-run   
  -h, --help      help for clean

Global Flags:
      --verbose   Enable verbose logging
```

`--dry-run` flag outputs about CRDs to be updated and CRs to be migrated. It **does not** modify any CRD and CRs.

Using `asoctl crd clean` is an important step if `v1alpha1api` resources are present in the cluster. If not used correctly, operator pod would output log like:

```
"msg"="failed to apply CRDs" "error"="failed to apply CRD storageaccountsqueueservicesqueues.storage.azure.com: CustomResourceDefinition.apiextensions.k8s.io \"storageaccountsqueueservicesqueues.storage.azure.com\" is invalid: status.storedVersions[0]: Invalid value: \"v1alpha1api20210401storage\": must appear in spec.versions" 
```

### Steps for migration using `asoctl crd clean`:

- Ensure the current ASOv2 version is `beta.5`
- Run `asoctl crd clean`
- Upgrade ASOv2 to `2.0.0`

## Import Azure Resource

TBC
