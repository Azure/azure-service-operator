/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package extensions

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Importer is an optional interface that can be implemented by resource extensions to customize the import process.
type Importer interface {
	// Import allows interception of the import process.
	// resource is the resource being imported.
	// next is a function to call to do the actual import.
	Import(
		rsrc genruntime.ImportableResource,
		next ImporterFunc,
	) (ImportResult, error)
}

// ImportResult is the result of doing an import.
type ImportResult struct {
	because string
}

// ImporterFunc is the signature of the function that does the actual import.
type ImporterFunc func(
	resource genruntime.ImportableResource,
) (ImportResult, error)

// NewImportSucceeded creates a new ImportResult with a resource that was imported successfully.
func NewImportSucceeded() ImportResult {
	return ImportResult{}
}

// NewImportSkipped creates a new ImportResult for a resource that was not imported.
func NewImportSkipped(because string) ImportResult {
	return ImportResult{
		because: because,
	}
}

// Skipped returns a reason and true if the import was skipped, empty string and false otherwise.
func (r ImportResult) Skipped() (string, bool) {
	return r.because, r.because != ""
}
