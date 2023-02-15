/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// resourceImporter represents an importer that's been created for a specific resource
type resourceImporter interface {
	Import(ctx context.Context) (*resourceImportResult, error)
}

// resourceImportResult represents the result of an import operation
// TODO: extend this to include child objects that need to be imported
type resourceImportResult struct {
	// Object is the resource that has been imported
	Object genruntime.MetaObject
}
