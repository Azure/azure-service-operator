# merger

## Overview

The merger package provides utilities for combining multiple Kubernetes client objects into a consolidated set. It specializes in merging ConfigMaps and Secrets from different sources, ensuring keys and values are properly combined without conflicts.

**Object merging**: The package implements logic to consolidate client objects, with special handling for ConfigMaps and Secrets that need to be merged based on their name and key structure rather than replaced outright.

**Namespace validation**: Ensures that objects being merged belong to the same namespace, preventing unintended cross-namespace operations that would be disallowed in Kubernetes.

**Collector integration**: Leverages the collector patterns from the configmaps and secrets packages to efficiently merge objects with consistent behavior and error handling.

## Testing

The package is tested using Go's standard testing framework with:

* Gomega for assertions to verify merge behaviors
* Tests for successful merging of objects with common keys and names
* Tests for expected failures when trying to merge objects from different namespaces
* Tests for key collisions and proper error reporting

Tests can be run using:

```bash
go test ./v2/pkg/genruntime/merger/...
```

## Related packages

* **genruntime**: The parent package that defines the core types like ConfigMapDestination and SecretDestination used by the merger.
* **genruntime/configmaps**: Provides the Collector pattern used by merger to consolidate ConfigMap resources.
* **genruntime/secrets**: Provides the Collector pattern used by merger to consolidate Secret resources.
* **internal/reconcilers**: Uses the merger package to combine resources from different extension points during reconciliation.
* **genruntime/extensions**: Contains extension points that may produce resources to be merged.
