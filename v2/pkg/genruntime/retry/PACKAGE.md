# retry

## Overview

The retry package provides classification and behavior definitions for retry operations in the Azure Service Operator (ASO). It establishes a standardized approach for determining how to handle transient failures when interacting with Azure resources.

**Retry classification**: Defines a set of retry classifications (None, Fast, Slow, VerySlow) that represent different retry strategies based on the nature of errors encountered during resource reconciliation.

**Failure handling**: Works in conjunction with the error classification system to determine appropriate retry behaviors for different types of failures, enabling more resilient resource management.

**Reconciliation pacing**: The classifications help control the pace of reconciliation attempts, preventing unnecessary load on both the Kubernetes cluster and Azure APIs during periods of temporary failure.

## Testing

The package is used and tested through its integration with other components:

* Integration tests verify the behavior of the retry system in various error scenarios
* Testing of retry behaviors with different error classifications
* Verification that retry classifications properly translate to reconciliation intervals

Tests demonstrate the appropriate behavior of different retry classifications in the broader ASO context.

## Related packages

* **genruntime/core**: Defines the error types that are associated with retry classifications.
* **genruntime/extensions**: Contains the ErrorClassifier extension point that can affect retry behavior.
* **internal/reconcilers**: Uses the retry classifications to determine appropriate requeue intervals during reconciliation.
* **internal/controllers**: Implements controller behavior that respects the retry classifications.
