/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// FieldMatcher allows a StringMatcher to be deserialized from YAML
type FieldMatcher struct {
	actual StringMatcher
}

var _ StringMatcher = &FieldMatcher{}

func NewFieldMatcher(field string) FieldMatcher {
	return FieldMatcher{
		actual: NewStringMatcher(field),
	}
}

func (dm *FieldMatcher) String() string {
	if dm.actual == nil {
		// No nested matcher
		return ""
	}

	return dm.actual.String()
}

func (dm *FieldMatcher) Matches(value string) bool {
	if dm.actual == nil {
		// No nested matcher
		return true
	}

	return dm.actual.Matches(value)
}

func (dm *FieldMatcher) WasMatched() error {
	if dm.actual == nil {
		// No nested matcher
		return nil
	}

	return dm.actual.WasMatched()
}

// IsRestrictive returns true if our nested matcher is present and restrictive, false otherwise
func (dm *FieldMatcher) IsRestrictive() bool {
	return dm.actual != nil && dm.actual.IsRestrictive()
}

// UnmarshalYAML populates our instance from the YAML.
// We expect just a single string, which we use create an actual StringMatcher
func (dm *FieldMatcher) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return errors.New("expected scalar value")
	}

	dm.actual = NewStringMatcher(value.Value)
	return nil
}
