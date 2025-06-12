# configmaps

## Overview

The configmaps package provides functionality for managing ConfigMap resources in Kubernetes for the Azure Service Operator (ASO). It enables the collection, validation, and processing of ConfigMap data, facilitating the export of Azure resource properties as Kubernetes ConfigMaps.

**Collector pattern**: The package implements a Collector type that aggregates ConfigMap values from different sources, merges them when appropriate, and produces properly structured Kubernetes ConfigMap objects ready for application to a cluster.

**Validation utilities**: The package includes comprehensive validation logic that ensures ConfigMap destinations don't conflict with each other, verifies CEL expressions used for dynamic ConfigMap generation, and validates that keys are appropriately specified based on the output type of expressions.

**ConfigMap exporting**: Facilitates the export of Azure resource properties to Kubernetes ConfigMaps through the Exporter interface, which defines methods for providing destination expressions that specify how to generate ConfigMap values.

## Testing

The package is tested using Go's standard testing framework with:

* Gomega for assertions to verify ConfigMap generation, validation, and collection behavior
* Tests for various validation scenarios including key collisions, type validations, and expression evaluation
* Test coverage for edge cases like nil elements, empty lists, and case sensitivity

Tests can be run using:

```bash
go test ./v2/pkg/genruntime/configmaps/...
```

## Related packages

* **genruntime**: The parent package that defines the core types like `ConfigMapReference` and `ConfigMapDestination` used throughout the configmaps package.
* **genruntime/core**: Provides shared declarations used by the configmaps package, particularly the `DestinationExpression` type for dynamic ConfigMap generation.
* **genruntime/merger**: Uses the configmaps package to merge ConfigMaps from multiple sources during reconciliation.
* **functions**: Contains code generation functions that generate ConfigMap exporting functionality for ASO resources.
* **internal/resolver**: Implements the ConfigMapResolver which resolves ConfigMap references at runtime during resource reconciliation.
