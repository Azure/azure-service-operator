/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"errors"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// ExportConfiguration is used to control which types get generated
type ExportConfiguration struct {
	// Base URL for the JSON schema to generate
	SchemaURL string
	// Filters used to control which types are included
	TypeFilters []*TypeFilter
}

// ShouldExportResult is returned by ShouldExport to indicate whether the supplied type should be exported
type ShouldExportResult string

const (
	// Export indicates the specified type should be exported to disk
	Export ShouldExportResult = "export"
	// Skip indicates the specified type should be skpped and not exported
	Skip ShouldExportResult = "skip"
)

// NewExportConfiguration is a convenience factory for ExportConfiguration
func NewExportConfiguration(filters ...*TypeFilter) *ExportConfiguration {
	result := ExportConfiguration{
		TypeFilters: filters,
	}

	return &result
}

// Validate checks our configuration for common issues
func (config *ExportConfiguration) Validate() error {
	if config.SchemaURL == "" {
		return errors.New("SchemaURL missing")
	}

	return nil
}

// ShouldExport tests for whether a given struct should be exported
// Returns a result indicating whether export should occur as well as a reason for logging
func (config *ExportConfiguration) ShouldExport(definition *astmodel.StructDefinition) (result ShouldExportResult, because string) {
	for _, f := range config.TypeFilters {
		if f.AppliesToType(definition) {
			if f.Action == ExcludeType {
				return Skip, f.Because
			} else if f.Action == IncludeType {
				return Export, f.Because
			}
		}
	}

	// By default we export all types
	return Export, ""
}
