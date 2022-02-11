# Code Generator Overview

Core to Azure Service Operator (ASO) v2 is our code generator. This consumes ARM JSON Schema and Swagger specifications and generates code for each desired resource that works with our generic operator reconciler.

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

## Object Model

At the core, the code generator works with a rich semantic object model that captures the structure of the resources (and related types) we are generating.

Underpinning this object model is the `astmodel` package.

The interface `astmodel.Type` represents a specific data type. The most widely used implementations of `Type` fall into a few separate groups:

* Simple Go types, including `PrimitiveType`, `ArrayType`, and `MapType`.
* Complex types with internal structure, most notably `ObjectType` and `ResourceType`.
* Wrapper types that provide additional semantics, including `OptionalType`, `ErroredType`, `FlaggedType` and `ValidatedType`.

This list is not exhaustive; other implementations of `Type` are used within limited scopes. For example, `AllOfType` and `OneOfType` represent JSON Schema techniques for creating object definitions via composition.

Usefully, there is also `TypeName` which is both a type in itself and an indirect reference to a type defined elsewhere.

When a `Type` is given a `TypeName`, it becomes a `TypeDefinition` and can be emitted as the source code for a Go type definition. A set of many `TypeDefinition`, each with a unique name is a `Types`. 

Both `ResourceType` and `ObjectType` act as containers, each implementing `PropertyContainer`, `FunctionContainer`, and `TestCaseContainer`. These do pretty much what you'd expect from the names, though the implementations differ between `ResourceType` and `ObjectType`. For example, where an `ObjectType` implements `PropertyContainer` by providing support for an arbitrary set of properties, `ResourceType` has only `Spec` and `Status`.

Most implementations of `astmodel.Function` are found in the **functions** package. New function implementations should go here; existing implementations are slowly being relocated.

All implementations of `astmodel.TestCase` are found in the **testcases** package. New test case implementations should go here.

## Resources, Objects and other types

Each distinct resource is represented by a `ResourceType`. The `Spec` of each resource is an `ObjectType` containing a collection of `PropertyDefinition` values, along with implementations of the `Function`, and `TestCase` interfaces. The `Status` of a resource is a different `ObjectType`. 

Some properties capture primitive values (strings, integers, and so on), while others are themselves complex `ObjectType` definitions, forming a hierarchy of information. `MetaType` wrappers (including `OptionalType`, and `ValidatedType`) are used to add semantic information to both properties and to type definitions.

## Generator Pipeline

The code generator itself, found in the package `codegen`, is structured as a pipeline, composed of a series of stages that transform our object model incrementally. All the pipeline stages are found in the subpackage `codegen/pipeline`. 

One reason for this is to allow the creation of multiple pipelines (currently we have separate definitions targeting Azure, Crossplane, and for testing), each sharing the majority of their implementation. Another reason is to allow individual pipeline stages to be tested in isolation, though not all existing pipeline stages have tests in this form. New stage implementations should have isolated tests where possible.

Pipeline stages are instances of `pipeline.Stage`. Each has a factory method that returns a `pipeline.Stage` instance. In operation, each accepts a `pipeline.State` containing the current object model and transforms it into a new state, that is passed to the next stage in turn. If a stage returns an error, the pipeline run is aborted.

New stages should use the `MakeStage()` function. You'll see some older stages that predate a structural change use the deprecated `MakeLegacyStage()` factory instead; these older stages are slowly being migrated and `MakeLegacyStage()` will be deleted when this is complete.

## Code Generation

Code is generated as a Go abstract syntax tree fragments which are then serialized to files as compilable Go code, forming a part of the operator itself. 

To make generation easier, our `astbuilder` package contains a wide variety of helper methods that allow declarative construction of the required tree. We are using the `dst` package instead of the Go core `ast` package, as it provides better control of comments and whitespace in the final output.
