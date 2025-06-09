# Genruntime Package

## Overview

The `genruntime` package provides the runtime foundation for Azure Service Operator (ASO), defining essential interfaces and types used by generated code and reconcilers to interact with Azure resources. This package enables consistent behavior across all generated controllers regardless of the Azure service they target.

**Resource abstractions**: The package defines key abstractions like `KubernetesResource`, `ARMResource`, and `ARMMetaObject` that represent Azure resources in various contexts. These abstractions provide a common interface for working with resources across the operator's codebase, ensuring consistent behavior regardless of the specific Azure service.

**Reconciliation primitives**: The package includes the `Reconciler` interface and related types that define how resources are created, updated, and deleted in Azure. This reconciliation system handles the translation of Kubernetes resource states to Azure resource states, including status updates and error reporting.

**Resource identity and ownership**: The package provides functionality for managing resource IDs, references, and ownership relationships between resources. This includes helper functions for extracting and parsing ARM resource IDs and enforcing ownership semantics between related resources.

**Extension points**: The package defines interfaces like `ResourceExtension`, `KubernetesSecretExporter`, and `KubernetesConfigExporter` that allow resources to extend their behavior to customize reconciliation. These extension points enable custom handling of secrets, configuration, and other resource-specific requirements.

**Validation and defaulting**: The package includes interfaces and helpers for resource validation and defaulting, ensuring that resources are properly configured before being applied to Azure. This includes support for both code-generated and custom validation rules.

**Type conversion**: The package provides utilities for converting between different versions of resources and between Kubernetes resources and their ARM representations. This includes support for handling SDK-specific type conversions and property transformations.

## Testing

The genruntime package includes several test files to verify its functionality. To run the tests:

```bash
cd v2/pkg/genruntime
go test ./...
```

The tests validate the core functionality of the package, including resource references, property bags, and admission controls. Testing approach includes:

* **Unit tests**: Uses the standard Go testing library with Gomega assertions (github.com/onsi/gomega)
* **Test suite**: A dedicated test suite in the `test` subdirectory provides integration testing tools for resources built on top of the genruntime package
* **Golden files**: Some tests use golden files to validate output consistency
* **Mock implementations**: For testing interface implementations without requiring actual Azure resources

For more complex testing scenarios, the package works with the test helpers in the internal testing packages to provide mocking and verification of Azure API interactions. No special environment variables are required for basic tests, but integration tests may require Azure credentials.

## Related packages

- **`v2/tools/generator/internal/astmodel`**: Uses the interfaces and types defined in genruntime as targets for code generation. The types defined in genruntime are used as base types and interfaces that generated code must implement.

- **`v2/tools/generator/internal/codegen`**: Generates code that implements the interfaces defined in genruntime. The codegen package produces Go code that satisfies the contracts specified by genruntime interfaces.

- **`v2/pkg/genruntime/core`**: Contains core error types and fundamental declarations used by genruntime and the generic reconciler. This is a minimal dependency package to avoid circular dependencies.

- **`v2/pkg/genruntime/registration`**: Provides functionality for registering resource types with the operator. This package uses the types defined in genruntime to create registrations for reconcilers and controllers.

- **`v2/pkg/genruntime/conditions`**: Defines condition types and helpers used to track resource state in Kubernetes status objects. These conditions follow Kubernetes conventions for reporting resource status.

- **`v2/pkg/genruntime/secrets`** and **`v2/pkg/genruntime/configmaps`**: Provide specialized functionality for working with Kubernetes secrets and configmaps, particularly for exporting sensitive information from Azure resources.
