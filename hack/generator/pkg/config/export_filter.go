/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

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
	Action      ExportFilterAction
	TypeMatcher `yaml:",inline"`
}
