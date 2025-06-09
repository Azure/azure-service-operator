# reporting

## Overview

The reporting package provides utilities for generating human-readable reports and structured documentation from the ASO code generator's internal type representations. It supports creating formatted tables, hierarchical structure reports, and type catalog documentation.

**Structured tables**: The package offers both sparse table and Markdown table implementations, providing flexible ways to generate tabular data with automatic column width calculation and alignment for clean output.

**Type catalog reports**: Specialized reporting tools that transform complex type definition sets into readable documentation, essential for generating API reference documentation and understanding the relationships between types.

**Hierarchical structure reports**: Generates tree-like visualizations of nested structures, making it easier to understand complex type hierarchies and relationships between components.

**Markdown generation**: Provides utilities for generating well-formatted Markdown output that can be used in documentation files, GitHub README files, and other documentation systems.

## Testing

The package uses standard Go testing with:

* Gomega for assertions to validate the generated report structures
* Golden file testing for comparing generated output with expected content
* Unit tests for individual formatting and rendering functions

Tests can be run using:

```bash
go test ./v2/tools/generator/internal/reporting/...
```

## Related packages

* **astmodel**: Provides the type definitions and structures that reporting uses to generate documentation.
* **codegen**: Uses reporting to generate documentation during the code generation process.
* **test**: Uses reporting capabilities to format and present test results and failure information.
* **generator**: The main generator package integrates with reporting to create documentation artifacts alongside generated code.
