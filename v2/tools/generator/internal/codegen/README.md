# Codegen Package

The `codegen` package serves as the core code generation engine for the Azure Service Operator (ASO). It orchestrates the generation of Kubernetes resource definitions and supporting code from Azure ARM templates and schema definitions through a configurable pipeline architecture.

## Overview

The `codegen` package provides a flexible and extensible pipeline-based architecture for transforming Azure Resource Manager (ARM) API definitions into Kubernetes Custom Resource Definitions (CRDs) and their corresponding Go types. It coordinates the end-to-end generation process, from schema loading through type transformation to Go file output.

**Pipeline architecture** forms the foundation of the code generation system. The generation process is broken down into discrete stages, each handling a specific transformation task. Stages can be included or excluded based on the target platform (ARM or Crossplane), allowing flexible code generation for different purposes.

**Code Generator** serves as the central coordinator that configures and executes the pipeline. It loads configuration from YAML files, sets up the appropriate pipeline stages based on the target platform, and processes the input schema to generate output code.

**File generation** capabilities include producing Go source files, test files, and supporting documentation. The system handles package structure, imports, and proper formatting for generated code.

## Testing

Testing of the codegen package uses several approaches to ensure correctness:

* Gomega is used as the primary assertion library for unit tests
* Golden file testing compares generated output against expected results
* Pipeline stage tests verify the behavior of individual pipeline stages
* Configuration loading tests ensure proper parsing and validation of input configuration

## Related packages

* **astmodel** - Provides the abstract syntax tree model used for code generation. The codegen package uses these AST models to represent and transform Go code structures.
* **astbuilder** - Offers helper functions for building Go AST nodes. The codegen package uses this to construct code fragments.
* **config** - Provides configuration structures used by the codegen package to control code generation behavior.
* **pipeline** - Contains the individual pipeline stages used in the code generation process. The codegen package assembles and executes these stages.
* **reporting** - Provides logging and reporting capabilities used to track the code generation process and capture metrics.
* **test** - Contains testing utilities used by the codegen package and other generator components.
