# ARM Conversion Package

## Overview

The `armconversion` package provides specialized builders for generating the functions that convert between Kubernetes resource types and their corresponding Azure Resource Manager (ARM) representations. It handles the complex mapping of properties between these different representations, supporting the bidirectional translation required for Azure Service Operator resources.

**ARM conversion function builders**: The package implements specialized builders for generating the `ConvertToARM` and `PopulateFromARM` functions required by each resource. These builders handle the complexities of property mapping, flattening, nesting, and type transformations needed when converting between Kubernetes and ARM formats.

**Property conversion handlers**: The package defines a chain of specialized handlers for different property conversion scenarios. Each handler addresses specific patterns like reference properties, flattened properties, name transformations, and specialized collections like user-assigned identities.

**Advanced transformation support**: The package provides mechanisms for property promotion and demotion when moving between formats. This allows for restructuring object hierarchies when the Kubernetes and ARM representations have different nesting patterns, enabling attributes to move up or down in the object tree during conversion.

**Interface implementation generation**: The package generates the code necessary to implement the `ARMTransformer` and `FromARMConverter` interfaces from the genruntime package. This ensures that all generated resources can seamlessly interact with the ARM client infrastructure.

## Testing

The package includes unit tests for the ARM conversion builders. To test the armconversion package:

```bash
cd v2/tools/generator/internal/armconversion
go test ./...
```

The testing approach includes:

* **Unit tests**: The package uses the standard Go testing library with Gomega for assertions
* **Function generation testing**: Tests verify that the generated ARM conversion functions handle all required property conversion scenarios
* **Edge case verification**: Tests for special cases like flattened properties, deep nesting, and resource references

No special environment variables or setup is required to run these tests.

## Related packages

- **`v2/tools/generator/internal/astmodel`**: Provides the type system model that the armconversion package uses to understand and manipulate the structure of resources. It serves as the foundation for understanding type relationships and property definitions.

- **`v2/tools/generator/internal/astbuilder`**: Used by armconversion builders to generate the actual Go code for the ARM conversion functions. The armconversion package decides what code to generate while astbuilder handles the mechanics of AST construction.

- **`v2/pkg/genruntime`**: Defines the interfaces (`ARMTransformer`, `FromARMConverter`) that the armconversion package generates implementations for. The generated code enables resources to interact with the Azure SDK through the genruntime abstractions.
