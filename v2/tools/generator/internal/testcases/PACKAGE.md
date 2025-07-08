# testcases

## Overview

The testcases package provides specialized test case generators for the Azure Service Operator code generator. It focuses on creating comprehensive test cases that validate critical aspects of generated code such as property assignments, JSON serialization, and resource conversions.

**Property assignment testing**: Generates test cases that verify lossless property transfers between related objects in the conversion graph, ensuring data integrity is maintained during conversion operations.

**JSON serialization validation**: Creates test cases that validate objects can be correctly serialized to JSON and deserialized back without data loss, crucial for ensuring proper persistence and API compatibility.

**Resource conversion testing**: Implements multi-step conversion test cases that verify resources can be converted to hub resources and back again without data loss, essential for maintaining compatibility across API versions.

**Test code generation**: Uses AST manipulation to dynamically generate readable, maintainable test code that comprehensively exercises generated types and their operations.

## Testing

The package uses standard Go testing with:

* Gomega for assertions to validate the test case generation logic
* Golden file testing that compares generated test code against expected outputs
* Property-based testing approaches for comprehensive coverage of edge cases

Tests can be run using:

```bash
go test ./v2/tools/generator/internal/testcases/...
```

## Related packages

* **astmodel**: Provides the type definitions that testcases uses as subjects for test generation.
* **astbuilder**: Used to construct AST nodes that form the generated test code.
* **functions**: Works with various function implementations that testcases needs to test, such as property assignment and resource conversion functions.
* **conversions**: Provides conversion logic that testcases validates through generated tests.
* **test**: Supplies foundational testing utilities that testcases builds upon for specific test scenarios.
