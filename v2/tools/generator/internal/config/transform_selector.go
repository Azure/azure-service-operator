/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"strings"
	"sync"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TransformSelector is used to select a type for transformation
type TransformSelector struct {
	Group        FieldMatcher           `yaml:",omitempty"`
	Version      FieldMatcher           `yaml:"version,omitempty"`
	Name         FieldMatcher           `yaml:",omitempty"`
	Optional     bool                   `yaml:",omitempty"`
	Object       bool                   `yaml:",omitempty"`
	Map          *MapSelector           `yaml:",omitempty"`
	lock         sync.Mutex             // protects appliesCache
	appliesCache map[astmodel.Type]bool // cache for the results of AppliesToType()
}

type MapSelector struct {
	Key   TransformSelector `yaml:",omitempty"`
	Value TransformSelector `yaml:",omitempty"`
}

func (ts *TransformSelector) AppliesToType(t astmodel.Type) bool {
	ts.lock.Lock()
	defer ts.lock.Unlock()

	if ts.appliesCache == nil {
		ts.appliesCache = make(map[astmodel.Type]bool)
	}

	if result, ok := ts.appliesCache[t]; ok {
		return result
	}

	result := ts.appliesToType(t)
	ts.appliesCache[t] = result
	return result
}

func (ts *TransformSelector) appliesToType(t astmodel.Type) bool {
	if ts == nil {
		return true
	}

	inspect := t
	if ts.Optional {
		// Need optional
		opt, ok := astmodel.AsOptionalType(inspect)
		if !ok {
			// but don't have optional
			return false
		}

		inspect = opt.Element()
	}

	if ts.Object {
		if _, ok := astmodel.AsObjectType(inspect); !ok {
			return false
		}
	}

	if ts.Name.IsRestrictive() {
		if ts.Group.IsRestrictive() || ts.Version.IsRestrictive() {
			// Expecting TypeName
			if tn, ok := astmodel.AsInternalTypeName(inspect); ok {
				return ts.appliesToTypeName(tn)
			}

			return false
		}

		// Expecting primitive type
		if pt, ok := astmodel.AsPrimitiveType(inspect); ok {
			return ts.appliesToPrimitiveType(pt)
		}

		return false
	}

	if ts.Map != nil {
		// Expecting map type
		if mp, ok := astmodel.AsMapType(inspect); ok {
			return ts.appliesToMapType(mp)
		}

		return false
	}

	return true
}

func (ts *TransformSelector) appliesToTypeName(tn astmodel.InternalTypeName) bool {
	if !ts.Name.Matches(tn.Name()).Matched {
		// No match on name
		return false
	}

	grp, ver := tn.InternalPackageReference().GroupVersion()

	if ts.Group.IsRestrictive() {
		if !ts.Group.Matches(grp).Matched {
			// No match on group
			return false
		}
	}

	if ts.Version.IsRestrictive() {
		// Need to handle both full (v1beta20200101) and API (2020-01-01) formats
		switch ref := tn.PackageReference().(type) {
		case astmodel.LocalPackageReference:
			if !ref.HasAPIVersion(ts.Version.String()) && !ts.Version.Matches(ver).Matched {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func (ts *TransformSelector) appliesToPrimitiveType(pt *astmodel.PrimitiveType) bool {
	if ts.Name.Matches(pt.Name()).Matched {
		return true
	}

	// Special case, allowing config to use `any` as a synonym for `interface{}`
	if strings.EqualFold(ts.Name.String(), "any") &&
		strings.EqualFold(pt.Name(), "interface{}") {
		return true
	}

	return false
}

func (ts *TransformSelector) appliesToMapType(mp *astmodel.MapType) bool {
	return ts.Map.Key.AppliesToType(mp.KeyType()) &&
		ts.Map.Value.AppliesToType(mp.ValueType())
}
