# Code Generator Overview

Core to Azure Service Operator (ASO) v2 is our code generator. This consumes ARM JSON Schema and Swagger specifications and generates the code for each desired resource.

## Code Structure

The key packages used to structure the code of the generator are as follows:

| Package          | Content                                                                                                                                               |
| ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `astmodel`         | Short for _Abstract Syntax Tree Model_, this package contains the core data types for defining the Go functions, data types, and methods we generate. |
| `functions`        | Support for generation of individual functions, based on the `astmodel.Functions` interface                                                           |
| `interfaces`       | Support for generation of interface implementations, based on the `astmodel.InterfaceImplementer` interface                                           |
| `testcases`        | Support for generation of test cases (used to verify our generated code works as expected), based on the `astmodel.TestCase` interface                |
| `astbuilder`       | Intention revealing utility methods for creating the underlying Go abstract syntax tree we serialize as the last step of generation.                  |
| `codegen`          | The core processing pipeline of the code generator                                                                                                    |
| `codegen/pipeline` | Individual pipeline stages that are composed to form the code generator itself.                                                                       |
| `test`             | Support methods to make writing tests easier                                                                                                          |

### Object Model

At the core, the code generator works with a rich semantic object model that captures the structure of the resources (and related types) we are generating.

Underpinning this object model is the `astmodel` package. One of the core `astmodel` interfaces is `Type`, representing a specific data type. Implementations of `Type` fall into a few separate groups:

* Simple Go types, including `PrimitiveType`, `ArrayType`, and `MapType`.
* Complex types with internal structure, most notably `ObjectType` and `ResourceType`.
* Wrapper types that provide additional semantics, including `OptionalType`, `ErroredType`, `FlaggedType` and `ValidatedType`.
* Schema types, including `AllOfType` and `OneOfType`.

(This list is not exhaustive.)

Usefully, there is also `TypeName` which is both a type in itself and an indirect reference to a type.

When a `Type` is given a name, it becomes a `TypeDefinition` and can be emitted as a Go type definition.

Both `ResourceType` and `ObjectType` act as containers, each implementing `PropertyContainer`, `FunctionContainer` and `TestCaseContainer`. These do pretty much what you'd expect from the names, though the implementations may differ between `ResourceType` and `ObjectType`. For example, where an `ObjectType` implements `PropertyContainer` by providing support for an arbitrary set of properties, `ResourceType` has only `Spec` and `Status`.

Most implementations of `astmodel.Function` are found in the **functions** package. New function implementations should go here; existing implementations are slowly being relocated.

All implementations of `astmodel.TestCase` are found in the **testcases** package. New test case implementations should go here.

### Generator Pipeline

The code generator itself, found in the package `codegen`, is structured as a pipeline, each `pipeline.Stage` transforming our object model incrementally.

One reason for this is to allow the creation of multiple pipelines (currently targeting Azure, Crossplane, and testing), each sharing the majority of their implementation. Another reason is to allow individual pipeline stages to be tested in isolation, though not all existing pipeline stages have tests in this form. New stage implementations should have isolated tests where possible.

Pipeline stages are defined in the sub-package `codegen/pipeline`. Each has a factory method that returns a `pipeline.Stage` instance. New stages should use the `MakeStage()` function, but older stages that predate a structural change use a `MakeLegacyStage()` factory instead.
