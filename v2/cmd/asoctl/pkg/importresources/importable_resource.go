/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importreporter"
)

// ImportableResource is an interface that wraps a Kubernetes resource that can be imported.
// Different implementations of this interface will be used for different types of resources.
type ImportableResource interface {
	// GroupKind returns the GroupKind of the resource being imported.
	// (empty if the GK can't be determined)
	GroupKind() schema.GroupKind

	// Name is a human-readable identifier for this resource
	Name() string

	// ID is a unique identifier for this resource.
	// The ID of a resource must unique within the import operation; the easiest way to achieve this is
	// to make it globally unique.
	ID() string

	// Import does the actual import, updating the Spec on the wrapped resource.
	// ctx allows for cancellation of the import.
	// log allows information about progress to be reported
	Import(
		ctx context.Context,
		reporter importreporter.Interface,
		factory *importFactory,
		log logr.Logger,
	) (ImportResourceResult, error)
}
