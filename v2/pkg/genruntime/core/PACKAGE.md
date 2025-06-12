# core

## Overview

The core package provides essential types and error handling primitives used across the Azure Service Operator (ASO) codebase. It serves as the foundation for error classification, expression-based functionality, and standardized error reporting used by both the generic reconciler and generated code.

**Error classification**: This package defines the error classification system used throughout ASO, categorizing errors as retryable or fatal to enable intelligent retry behaviors during reconciliation operations.

**Expression support**: Includes the DestinationExpression type which is a key building block for exporting Azure resource properties into ConfigMaps and Secrets through CEL expressions, providing a declarative way to map resource fields to Kubernetes objects.

**Error types**: Implements a rich set of error types tailored for Kubernetes and Azure operations, including specialized errors for reference resolution, ownership conflicts, and cloud-specific error details that provide context about Azure API responses.

## Testing

The package is tested using Go's standard testing framework:

* Unit tests for error behaviors and classifications
* Tests for proper error wrapping and unwrapping
* Testing of error type detection and conversion utilities

Tests can be run using:

```bash
go test ./v2/pkg/genruntime/core/...
```

## Related packages

* **genruntime**: The parent package that uses core types as building blocks for resource definitions and operations.
* **genruntime/conditions**: Uses error types from core to create appropriate resource conditions based on error classifications.
* **genruntime/retry**: Works closely with core's error classification to determine appropriate retry strategies.
* **internal/reconcilers**: Depends on core's error types to handle and classify errors during reconciliation.
* **internal/genericarmclient**: Uses CloudErrorDetails from core to standardize ARM API error reporting.
