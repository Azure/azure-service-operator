/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TypeMatcher contains basic functionality for a filter
type TypeMatcher struct {
	Group   FieldMatcher `yaml:",omitempty"` // Filter matching types by group
	Version FieldMatcher `yaml:",omitempty"` // Filter matching types by version
	Name    FieldMatcher `yaml:",omitempty"` // Filter matching types by name
	// Because is used to articulate why the filter applied to a type (used to generate explanatory logs in debug mode)
	Because string
	// MatchRequired indicates if an error will be raised if this TypeMatcher doesn't match at least one type.
	// The default is true.
	MatchRequired *bool `yaml:"matchRequired,omitempty"`

	// matchedAnything is true if TypeMatcher matched anything
	matchedAnything bool
}

var _ fmt.Stringer = &TypeMatcher{}

// Initialize initializes the type matcher
func (t *TypeMatcher) Initialize() error {
	// Default MatchRequired
	if t.MatchRequired == nil {
		temp := true
		t.MatchRequired = &temp
	}

	return nil
}

// AppliesToType indicates whether this filter should be applied to the supplied type definition
func (t *TypeMatcher) AppliesToType(typeName astmodel.TypeName) bool {
	group, version, ok := typeName.PackageReference.TryGroupVersion()
	if !ok {
		// Never match external references
		return false
	}

	result := t.Group.Matches(group).Matched &&
		t.Version.Matches(version).Matched &&
		t.Name.Matches(typeName.Name()).Matched

	// Track this match, so we can later report if we didn't match anything
	if result {
		t.matchedAnything = true
	}

	return result
}

// RequiredTypesWereMatched returns an error if no matches were made
func (t *TypeMatcher) RequiredTypesWereMatched() error {
	if *t.MatchRequired {
		return t.WasMatched()
	}

	return nil
}

// WasMatched returns nil if this matcher ever matched at least one type, otherwise a diagnostic error
func (t *TypeMatcher) WasMatched() error {
	if t.matchedAnything {
		// Matched at least one type
		return nil
	}

	if err := t.Group.WasMatched(); err != nil {
		return errors.Wrapf(
			err,
			"type matcher [%s] matched no types; every group was excluded",
			t.String())
	}

	if err := t.Version.WasMatched(); err != nil {
		return errors.Wrapf(
			err,
			"type matcher [%s] matched no types; groups matched, but every version was excluded",
			t.String())
	}

	if err := t.Name.WasMatched(); err != nil {
		return errors.Wrapf(
			err,
			"type matcher [%s] matched no types; groups and versions matched, but every type was excluded",
			t.String())
	}

	// Don't expect this case to ever be used, but be informative anyway
	return errors.Errorf("Type matcher [%s] matched no types", t.String())
}

// String returns a description of this filter
func (t *TypeMatcher) String() string {
	var result strings.Builder
	var spacer string
	if t.Group.IsRestrictive() {
		result.WriteString(fmt.Sprintf("Group: %q", t.Group.String()))
		spacer = "; "
	}

	if t.Version.IsRestrictive() {
		result.WriteString(fmt.Sprintf("%sVersion: %q", spacer, t.Version.String()))
		spacer = "; "
	}

	if t.Name.IsRestrictive() {
		result.WriteString(fmt.Sprintf("%sName: %q", spacer, t.Name.String()))
		spacer = "; "
	}

	if t.Because != "" {
		result.WriteString(fmt.Sprintf("%sBecause: %q", spacer, t.Because))
	}

	return result.String()
}
