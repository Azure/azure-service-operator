# registration

## Overview

The registration package provides foundational types and utilities for registering controller-runtime resources within the Azure Service Operator (ASO). It defines the structures necessary for configuring custom resource definitions, controllers, and webhooks in a Kubernetes operator.

**Controller registration**: This package defines the `StorageType` structure, which encapsulates all the information needed to register a reconciler with the controller-runtime manager, including the object type, indexes, watches, and predicates.

**Field indexing**: Provides the Index type that supports efficient field indexing to enable fast lookup of resources based on arbitrary fields, which is essential for implementing cross-resource references and ownership relationships.

**Watch configuration**: Offers the Watch type that defines additional resources whose changes should trigger reconciliation of the primary resource, enabling proper handling of dependent resources.

**Webhook registration**: Through the KnownType structure, supports registration of validation and defaulting webhooks that handle custom resource validation and defaulting operations.

## Testing

The package is tested through integration with other ASO components:

* Controller registration tests verify that resources are properly registered with the controller-runtime manager
* Index registration tests ensure that custom field indexes work correctly
* Watch tests validate that dependent resource changes trigger appropriate reconciliations

Tests can be run as part of the broader ASO test suite.

## Related packages

* **genruntime**: The parent package that defines the Reconciler interface which is referenced by StorageType.
* **internal/controllers**: Uses the registration package to register all resource controllers with the controller-runtime manager.
* **internal/webhooks**: Implements validation and defaulting webhooks that are registered using the structures in this package.
* **v2/cmd/controller**: The main controller application that leverages this package for registering all resources.
