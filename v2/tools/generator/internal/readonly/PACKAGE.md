# readonly

## Overview

The readonly package provides immutable (read-only) data structures for the Azure Service Operator code generator. It offers a type-safe way to create and use immutable maps and other collections that prevent unintended modifications after creation.

**Immutable collections**: The package implements a readonly Map type that wraps standard Go maps, providing immutable access to their contents while preventing accidental modifications after creation. This is important for ensuring data integrity throughout the code generation pipeline.

**Type safety**: Using Go generics, the package provides type-safe implementations that work with any comparable key type and any value type, making it widely applicable across the codebase.

**Performance considerations**: The package offers both safe and unsafe creation methods, allowing callers to choose between defensive copying for safety or direct wrapping for performance in contexts where the source map won't be modified elsewhere.

## Testing

The package uses standard Go testing with:

* Gomega for assertions to verify immutability behavior
* Unit tests that validate operations work as expected and confirm immutability guarantees

Tests can be run using:

```bash
go test ./v2/tools/generator/internal/readonly/...
```

## Related packages

* **astmodel**: Uses readonly collections to manage type definitions and references securely.
* **codegen**: Leverages immutable data structures from readonly when processing code generation pipelines.
* **config**: Stores configuration using immutable collections to prevent unexpected modifications.
* **jsonast**: Uses readonly collections when parsing and manipulating JSON schema structures.
