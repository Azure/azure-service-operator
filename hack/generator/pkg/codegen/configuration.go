/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"errors"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// Configuration is used to control which types get generated
type Configuration struct {
	// Base URL for the JSON schema to generate
	SchemaURL string
	// Filters used to control which types are included
	TypeFilters []*jsonast.TypeFilter
}

// ShouldExportResult is returned by ShouldExport to indicate whether the supplied type should be exported
type ShouldExportResult string

const (
	// Export indicates the specified type should be exported to disk
	Export ShouldExportResult = "export"
	// Skip indicates the specified type should be skpped and not exported
	Skip ShouldExportResult = "skip"
)

// NewConfiguration is a convenience factory for Configuration
func NewConfiguration(filters ...*jsonast.TypeFilter) *Configuration {
	result := Configuration{
		TypeFilters: filters,
	}

	return &result
}

// Validate checks our configuration for common issues
func (config *Configuration) Validate() error {
	if config.SchemaURL == "" {
		return errors.New("SchemaURL missing")
	}

	return nil
}

// ShouldExport tests for whether a given struct should be exported
// Returns a result indicating whether export should occur as well as a reason for logging
func (config *Configuration) ShouldExport(definition astmodel.Definition) (result ShouldExportResult, because string) {
	for _, f := range config.TypeFilters {
		if f.AppliesToType(definition) {
			if f.Action == jsonast.ExcludeType {
				return Skip, f.Because
			} else if f.Action == jsonast.IncludeType {
				return Export, f.Because
			}
		}
	}

	// By default we export all types
	return Export, ""
}
