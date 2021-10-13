/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import "github.com/pkg/errors"

// ExportFilterAction defines the possible actions that should happen for types matching the filter
type ExportFilterAction string

const (
	// ExportFilterInclude indicates that any type matched by the filter should be exported to disk by the generator
	ExportFilterInclude ExportFilterAction = "include"

	// ExportFilterIncludeTransitive indicates that any type matched by the filter and any types referenced should be exported
	ExportFilterIncludeTransitive ExportFilterAction = "include-transitive"

	// ExportFilterExclude indicates that any type matched by the filter should not be exported as a struct.
	// This skips generation of types but may not prevent references to the type by other structs (which
	// will cause compiler errors).
	ExportFilterExclude ExportFilterAction = "exclude"
)

// A ExportFilter is used to control which types should be exported by the generator
type ExportFilter struct {
	Action ExportFilterAction
	// RenameTo indicates that the type matched by this export filter should be renamed to the
	// specified name. This may only be used with the include-transitive ExportFilterAction. If
	// the matcher matches more than a single type, an error is raised as we cannot rename multiple types
	// to the same name.
	RenameTo    string `yaml:"renameTo,omitempty"`
	TypeMatcher `yaml:",inline"`
}

// Initialize initializes the export filter
func (f *ExportFilter) Initialize() error {
	err := f.TypeMatcher.Initialize()
	if err != nil {
		return err
	}

	if f.RenameTo != "" && f.Action != ExportFilterIncludeTransitive {
		return errors.Errorf(
			"ExportFilter %s. RenameTo can only be specified on ExportFilters with Action %q",
			f.TypeMatcher.String(),
			ExportFilterIncludeTransitive)
	}

	return nil
}

// Error returns an error if the filter encountered an error, or nil if there was no error.
func (f *ExportFilter) Error() error {
	if !f.MatchedRequiredTypes() {
		return errors.Errorf("Export filter action: %q, target: %q matched no types", f.Action, f.String())
	}

	// No need to check that include-transitive is the action as that was already ensured above in Initialize()
	if f.RenameTo != "" && len(f.matchedTypes) != 1 {
		return errors.Errorf(
			"Export filter action: %q, target: %q has RenameTo specified and matched multiple types",
			f.Action,
			f.String())
	}

	return nil
}
