---
title: Code Generator
---

Core to Azure Service Operator (ASO) v2 is our code generator. This consumes ARM JSON Schema and Swagger specifications and generates code for each desired resource that works with our generic operator reconciler.

## Code Structure

Key packages used to structure the code of the generator are as follows:

| Package                                                                                                                      | Content                                                                                                                                                                                                                                                   |
| ---------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [`astmodel`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/astmodel)                 | Short for _Abstract Syntax Tree Model_, this package contains the core data types for defining the Go functions, data types, and methods we generate.                                                                                                     |
| [`functions`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/functions)               | Support for generation of individual functions, based on the interface [`astmodel.Function](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/function.go#L13)                                               |
| [`interfaces`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/interfaces)             | Support for generation of interface implementations, based on the [`astmodel.InterfaceImplementer`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/interface_implementer.go#L19) interface                |
| [`testcases`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/testcases)               | Support for generation of test cases (used to verify our generated code works as expected), based on the [`astmodel.TestCase`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/test_case.go#L11) interface |
| [`astbuilder`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/astbuilder)             | Intention revealing utility methods for creating the underlying Go abstract syntax tree we serialize as the last step of generation.                                                                                                                      |
| [`codegen`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/codegen)                   | The core processing pipeline of the code generator                                                                                                                                                                                                        |
| [`codegen/pipeline`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/codegen/pipeline) | Individual pipeline stages that are composed to form the code generator itself.                                                                                                                                                                           |
| [`test`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/test)                         | Support methods to make writing tests easier                                                                                                                                                                                                              |

### Directory structure overview

In this diagram is shown the full directory structure of the ASO code generator, including all the packages named above.

![Overview](../images/aso-codegen-structure.svg)

The size of each dot reflects the size of the file; the legend in the corner shows the meaning of colour.

## Object Model

At the core, the code generator works with a rich semantic object model that captures the structure of the resources (and related types) we are generating.

Underpinning this object model is the [`astmodel`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/astmodel) package.

The interface [`astmodel.Type`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/type.go) represents a specific data type. The most widely used implementations of `Type` fall into a few separate groups:

* Simple Go types, including [`PrimitiveType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/primitive_type.go), [`ArrayType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/array_type.go), and [`MapType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/map_type.go).

* Complex types with internal structure, most notably [`ObjectType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/object_type.go) and [`ResourceType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/resource_type.go).

* Wrapper types that provide additional semantics, including [`OptionalType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/optional_type.go), [`ErroredType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/errored_type.go), [`FlaggedType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/flagged_type.go) and [`ValidatedType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/validated_type.go).

This list is not exhaustive; other implementations of `Type` are used within limited scopes. For example, [`AllOfType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/allof_type.go) and [`OneOfType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/oneof_type.go) represent JSON Schema techniques for creating object definitions via composition.

Usefully, there is also [`TypeName`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/type_name.go) which is both a type in itself and an indirect reference to a type defined elsewhere.

When a `Type` is given a `TypeName`, it becomes a [`TypeDefinition`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/type_definition.go) and can be emitted as the source code for a Go type definition. A set of many `TypeDefinition`, each with a unique name is a [`TypeDefinitionSet`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/type_definition_set.go). 

Both `ResourceType` and `ObjectType` act as containers, each implementing [`PropertyContainer`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/property_container.go), [`FunctionContainer`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/function_container.go), and embedding [`InterfaceImplementer`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/interface_implementer.go). These do pretty much what you'd expect from the names, though the implementations differ between `ResourceType` and `ObjectType`. For example, where an `ObjectType` implements `PropertyContainer` by providing support for an arbitrary set of properties, `ResourceType` has only `Spec` and `Status`.

Most implementations of `astmodel.Function` are found in the [**functions**](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/functions) package. New function implementations should go here; existing implementations are slowly being relocated.

All implementations of `astmodel.TestCase` are found in the [**testcases**](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/testcases) package. New test case implementations should go here.

## Resources, Objects and other types

Each distinct resource is represented by a `ResourceType`. The `Spec` of each resource is an `ObjectType` containing a collection of `PropertyDefinition` values, along with implementations of the `Function`, and `TestCase` interfaces. The `Status` of a resource is a different `ObjectType`. 

Some properties capture primitive values (strings, integers, and so on), while others are themselves complex `ObjectType` definitions, forming a hierarchy of information. [`MetaType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/meta_type.go) wrappers (including `OptionalType`, and `ValidatedType`) are used to add semantic information to both properties and to type definitions.

## Generator Pipeline

The code generator itself, found in the package [`codegen`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/codegen), is structured as a pipeline, composed of a series of stages that transform our object model incrementally. All the pipeline stages are found in the subpackage [`codegen/pipeline`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/codegen/pipeline). 

One reason for this is to allow the creation of multiple pipelines (currently we have separate definitions targeting *Azure*, *Crossplane*, and for testing) each sharing the majority of their implementation. Another reason is to allow individual pipeline stages to be tested in isolation, though not all existing pipeline stages have tests in this form. New stage implementations should have isolated tests where possible. The helper method [`RunTestPipeline()`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/codegen/pipeline/stage_test.go#L18) is useful here.

Pipeline stages are instances of [`pipeline.Stage`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/codegen/pipeline/stage.go). Each has a factory method that returns a `pipeline.Stage` instance. In operation, each accepts a [`pipeline.State`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/codegen/pipeline/state.go) containing the current object model and transforms it into a new state, that is passed to the next stage in turn. If a stage returns an error, the pipeline run is aborted. Each `State` instance is immutable, allowing comparison between states when debugging.

New stages should use the [`NewStage()`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/codegen/pipeline/stage.go#L38) function. You'll see some older stages that predate a structural change use the deprecated [`NewLegacyStage()`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/codegen/pipeline/stage.go#L52) factory instead; these older stages are slowly being migrated and `NewLegacyStage()` will be deleted when this is complete.

## Code Generation

Code is generated as a Go abstract syntax tree fragments which are then serialized to files as compilable Go code, forming a part of the operator itself. 

To make generation easier, our [`astbuilder`](https://github.com/Azure/azure-service-operator/tree/main/v2/tools/generator/internal/astbuilder) package contains a wide variety of helper methods that allow declarative construction of the required tree. We are using the `dst` package instead of the standard Go core `ast` package, as it provides better control of comments and whitespace in the final output.
