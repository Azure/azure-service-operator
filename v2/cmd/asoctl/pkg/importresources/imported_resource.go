/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// ImportedResource is an interface that wraps a Kubernetes resource that has been imported.
type ImportedResource interface {
	// GroupKind returns the GroupKind of the resource that was imported
	GroupKind() schema.GroupKind

	// Name is a human-readable identifier for this resource
	Name() string

	// ID is a unique identifier for this resource.
	// The ID of a resource must unique within the import operation; the easiest way to achieve this is
	// to make it globally unique.
	ID() string

	// Resource returns the actual resource that has been imported.
	Resource() genruntime.MetaObject
}
