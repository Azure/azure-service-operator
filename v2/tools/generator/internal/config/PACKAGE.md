# Config Package

## Overview

The `config` package provides the configuration system that controls the code generation process in Azure Service Operator. It defines a comprehensive hierarchy of structures that allow fine-grained specification of how types and properties should be processed during code generation from ARM schemas.

**Schema-based configuration**: The package implements a YAML-based configuration system that defines how JSON schemas and Swagger specifications are processed during code generation. This configuration allows for precise control over type filtering, property customization, and resource definition.

**Hierarchical configuration model**: The package implements a multi-level configuration hierarchy (ObjectModelConfiguration → GroupConfiguration → VersionConfiguration → TypeConfiguration → PropertyConfiguration) that maps to the structure of Azure resources, enabling granular control at each level.

**Type filtering and transformation**: Configuration options allow for selective generation of types based on criteria like name patterns, paths, or other attributes. This includes capabilities for renaming types, controlling exports, and defining relationships between resources.

**Code generation pipeline definition**: The package defines standard generation pipelines (Azure, Crossplane) with configurable stages that determine the sequence of transformations applied to types during code generation. This enables consistent generation across different target environments.

## Testing

The package includes comprehensive tests for configuration loading and application. To test the config package:

```bash
cd v2/tools/generator/internal/config
go test ./...
```

Tests verify that configuration is correctly loaded from YAML files and properly applied during type generation. The package uses the following testing approaches:

* **Golden file tests**: Located in the `testdata` directory, these tests validate that configuration correctly drives type generation according to expectations.
* **Unit tests**: The package uses the standard Go testing library with Gomega for assertions.
* **Visitor pattern testing**: Specialized tests for the configuration visitors, type matchers, and transformation selectors that ensure proper processing of configured options.

No special environment variables or setup is required to run these tests.

## Related packages

- **`v2/tools/generator/internal/astmodel`**: Consumes configuration from this package to determine which types to generate and how to modify them. The config package drives the creation and transformation of type definitions in the astmodel package.

- **`v2/tools/generator/internal/codegen`**: Orchestrates the overall code generation process based on configuration provided by this package. The codegen package uses the configuration to control which pipeline stages are executed and how they transform the types.

- **`v2/tools/generator/internal/jsonast`**: Uses configuration from this package to control how JSON schemas are transformed into Go AST structures. The configuration determines which schemas are processed and how their properties are mapped to Go types.
