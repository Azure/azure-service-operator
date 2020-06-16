/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"errors"
	"fmt"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// Configuration is used to control which types get generated
type Configuration struct {
	// Base URL for the JSON schema to generate
	SchemaURL string `yaml:"schemaUrl"`
	// The folder where the code should be generated
	OutputPath string `yaml:"outputPath"`
	// Filters used to control which types are exported
	ExportFilters []*ExportFilter `yaml:"exportFilters"`
	// Filters used to control which types are created from the JSON schema
	TypeFilters []*TypeFilter `yaml:"typeFilters"`
	// Transformers used to remap types
	TypeTransformers []*TypeTransformer `yaml:"typeTransformers"`
}

// ShouldExportResult is returned by ShouldExport to indicate whether the supplied type should be exported
type ShouldExportResult string

const (
	// Export indicates the specified type should be exported to disk
	Export ShouldExportResult = "export"
	// Skip indicates the specified type should be skipped and not exported
	Skip ShouldExportResult = "skip"
)

// ShouldPruneResult is returned by ShouldPrune to indicate whether the supplied type should be exported
type ShouldPruneResult string

const (
	// Include indicates the specified type should be included in the type graph
	Include ShouldPruneResult = "include"
	// Prune indicates the type (and all types only referenced by it) should be pruned from the type graph
	Prune ShouldPruneResult = "prune"
)

// NewConfiguration is a convenience factory for Configuration
func NewConfiguration() *Configuration {
	result := Configuration{}
	return &result
}

/// WithExportFilters adds the provided ExportFilters to the configurations collection of ExportFilters
func (config *Configuration) WithExportFilters(filters ...*ExportFilter) *Configuration {
	result := *config
	result.ExportFilters = append(result.ExportFilters, filters...)

	return &result
}

// Initialize checks for common errors and initializes structures inside the configuration
// which need additional setup after json deserialization
func (config *Configuration) Initialize() error {
	if config.SchemaURL == "" {
		return errors.New("SchemaURL missing")
	}

	if config.OutputPath == "" {
		// Default to an apis folder in the current directory if not specified
		config.OutputPath = "apis"
	}

	var errs []error
	for _, filter := range config.ExportFilters {
		err := filter.Initialize()
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, filter := range config.TypeFilters {
		err := filter.Initialize()
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, transformer := range config.TypeTransformers {
		err := transformer.Initialize()
		if err != nil {
			errs = append(errs, err)
		}
	}

	return kerrors.NewAggregate(errs)
}

// ShouldExport tests for whether a given type should be exported as Go code
// Returns a result indicating whether export should occur as well as a reason for logging
func (config *Configuration) ShouldExport(typeName *astmodel.TypeName) (result ShouldExportResult, because string) {
	for _, f := range config.ExportFilters {
		if f.AppliesToType(typeName) {
			switch f.Action {
			case ExportFilterExclude:
				return Skip, f.Because
			case ExportFilterInclude:
				return Export, f.Because
			default:
				panic(fmt.Errorf("unknown exportfilter directive: %s", f.Action))
			}
		}
	}

	// By default we export all types
	return Export, ""
}

// ShouldPrune tests for whether a given type should be extracted from the JSON schema or pruned
func (config *Configuration) ShouldPrune(typeName *astmodel.TypeName) (result ShouldPruneResult, because string) {
	for _, f := range config.TypeFilters {
		if f.AppliesToType(typeName) {
			switch f.Action {
			case TypeFilterPrune:
				return Prune, f.Because
			case TypeFilterInclude:
				return Include, f.Because
			default:
				panic(fmt.Errorf("unknown typefilter directive: %s", f.Action))
			}
		}
	}

	// By default we include all types
	return Include, ""
}

// TransformType uses the configured type transformers to transform a type name (reference) to a different type.
// If no transformation is performed, nil is returned
func (config *Configuration) TransformType(name *astmodel.TypeName) (astmodel.Type, string) {
	for _, transformer := range config.TypeTransformers {
		result := transformer.TransformTypeName(name)
		if result != nil {
			return result, transformer.Because
		}
	}

	// No matches, return nil
	return nil, ""
}
