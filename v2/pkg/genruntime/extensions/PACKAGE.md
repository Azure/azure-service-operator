# extensions

## Overview

The extensions package provides a comprehensive set of extension points that allow customization of the Azure Service Operator (ASO) core reconciliation behavior. It enables resource-specific behaviors to be implemented without modifying the generic reconciler itself, providing a clean separation of concerns.

**Resource reconciliation customization**: This package defines interfaces like `PreReconciliationChecker` and `PostReconciliationChecker` that allow resources to implement custom logic that executes before and after the standard reconciliation process, enabling prerequisite verification and post-deployment operations.

**ARM resource modification**: The package provides the `ARMResourceModifier` extension point for customizing the ARM payloads sent to Azure API, allowing resource-specific adjustments before deployment that aren't covered by the standard conversion process.

**Error classification**: Through the `ErrorClassifier` interface, resources can customize how ARM API errors are processed, enabling resource-specific error handling strategies such as determining whether certain errors should be treated as retryable or fatal.

**Lifecycle customization**: Includes extension points like `Deleter` and `Claimer` that allow resources to customize the deletion process and resource claiming mechanisms respectively, adapting them to resource-specific requirements.

## Testing

The package is tested using Go's standard testing framework:

* Unit tests for all extension point functionality
* Tests for the extension point creation helpers
* Tests verifying extension behavior in different scenarios

Tests can be run using:

```bash
go test ./v2/pkg/genruntime/extensions/...
```

## Related packages

* **genruntime**: The parent package that defines the ResourceExtension interface that forms the foundation of the extension system.
* **genruntime/core**: Provides error types and other primitives used by extension implementations.
* **internal/reconcilers**: Contains the generic reconciler that uses the extension points to customize reconciliation behavior.
* **api/**/customizations**: Contains the actual resource-specific implementations of the extension interfaces.
* **tools/generator**: Generates extension scaffolding as part of the code generation pipeline.
