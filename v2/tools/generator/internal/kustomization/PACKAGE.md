# kustomization

## Overview

The kustomization package provides utilities for creating and managing Kubernetes Kustomization resources for the Azure Service Operator. It handles generation of kustomization.yaml files and conversion patch files used for enabling conversion webhooks in Custom Resource Definitions (CRDs).

**Resource Definitions**: This package provides functionality to load and parse Kubernetes resource definitions from YAML files, with specific focus on extracting metadata such as resource names needed for creating conversion patches.

**CRD Kustomization**: The package generates and manages kustomization.yaml files used to combine CRD resources with their corresponding webhook conversion patches, providing a clean way to apply all related resources together.

**Conversion Patches**: Automatically generates conversion webhook patches for CRDs to support conversion between different versions of the same resource type, ensuring compatibility across API versions.

## Testing

The package uses standard Go testing with:

* Gomega for assertions to validate the generated YAML content
* Golden file testing for validating file creation and structure

Tests can be run using the standard Go test command:

```bash
go test ./v2/tools/generator/internal/kustomization/...
```

## Related packages

* **generator/gen_kustomize.go**: Uses this package to implement the gen-kustomize command-line functionality for the ASO code generator.
* **config**: Dependent on the kustomization package for generating CRD configurations.
* **astmodel**: Provides type definitions that kustomization uses when handling CRDs.
* **readonly**: Provides immutable data structures used for safely sharing configuration data.
