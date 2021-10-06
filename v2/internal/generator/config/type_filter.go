/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

// TypeFilterAction defines the possible actions that should happen for types matching the filter
type TypeFilterAction string

const (
	// TypeFilterInclude indicates that any type matched by the filter should be included in the type graph
	TypeFilterInclude TypeFilterAction = "include"
	// TypeFilterPrune indicates that any type matched by the filter, and any types only referenced by that type
	// should not be included in the type graph
	TypeFilterPrune TypeFilterAction = "prune"
)

// A TypeFilter is used to control which types should be included in the type graph when running the generator
type TypeFilter struct {
	Action      TypeFilterAction
	TypeMatcher `yaml:",inline"`
}
