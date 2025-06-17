# AST Model Package

## Overview

The `astmodel` package is the core of the Azure Service Operator code generator, providing a rich semantic object model for defining Go data types, functions, and methods. It's an abstraction layer over the raw Go AST (Abstract Syntax Tree), allowing for structured representation of code to be generated. This package defines the semantic model that drives the generation of all Kubernetes controllers for Azure resources.

**Type system abstractions**: The package defines a comprehensive type system with the `Type` interface at its core. It offers implementations for all common Go types such as primitives, arrays, maps, and more complex types like objects and enums. This abstraction allows the generator to work with strongly typed representations rather than raw syntax trees.

**Package and import management**: The package includes robust systems for handling Go package references, imports, and naming conflicts. It manages the complexities of cross-package references and import aliases automatically, which is essential when generating code that spans multiple packages.

**Code generation context**: The `CodeGenerationContext` provides a unified environment for generating Go code, tracking imported packages and used references. This context ensures consistency across all generated files and manages the interdependencies between different parts of the generated codebase.

**Resource type modeling**: The package offers specialized support for Kubernetes resource types, including custom resource definitions (CRDs) with their specifications, statuses, and validation rules - a key requirement for Azure Service Operator's Azure resource controllers.

## Testing

The package includes comprehensive tests for all types and functionality. To run the tests:

```bash
cd v2/tools/generator/internal/astmodel
go test ./...
```

The testing approach includes:

* **Unit tests**: The package uses the standard Go testing library with Gomega for assertions
* **Type system testing**: Tests validate the behavior of the type system abstractions
* **Code generation testing**: Tests verify that types can be correctly transformed into Go ASTs
* **Package management testing**: Tests ensure proper handling of imports, references, and naming

No special environment variables or setup is required to run these tests.

## Related packages

- **`v2/tools/generator/internal/astbuilder`**: Provides intention-revealing utility methods for creating and manipulating Go ASTs. While `astmodel` defines what to generate, `astbuilder` handles the low-level AST construction.

- **`v2/tools/generator/internal/jsonast`**: Transforms JSON schemas into `astmodel` types. It uses the abstractions provided by `astmodel` to represent the schema's types and structures in the Go type system.

- **`v2/tools/generator/internal/functions`**: Contains implementations of the `astmodel.Function` interface for various common code generation tasks. These functions are used to generate the method implementations for resources.

- **`v2/tools/generator/internal/testcases`**: Implements the `astmodel.TestCase` interface to generate test cases for verifying the functionality of generated code.
