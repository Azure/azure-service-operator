/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"regexp"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// TypeFilterAction defines the possible actions that should happen for types matching the filter
type TypeFilterAction string

const (
	// IncludeType indicates that any type matched by the filter should be exported to disk by the generator
	IncludeType TypeFilterAction = "include"
	// ExcludeType indicates that any type matched by the filter should be skipped and not exported
	ExcludeType TypeFilterAction = "exclude"
)

// A TypeFilter is used to control which types should be exported by the generator
type TypeFilter struct {
	Action TypeFilterAction
	// Schema is a wildcard matching specifier for which types are selected by this filter
	Schema      string `yaml:",omitempty"`
	schemaRegex *regexp.Regexp
	// Version is a wildcard matching specifier for which types are selected by this filter
	Version      string `yaml:",omitempty"`
	versionRegex *regexp.Regexp
	// Name is a wildcard matching specifier for which types are selected by this filter
	Name      string `yaml:",omitempty"`
	nameRegex *regexp.Regexp
	// Because is used to articulate why the filter applied to a type (used to generate explanatory logs in debug mode)
	Because string
}

// AppliesToType indicates whether this filter should be applied to the supplied type definition
func (filter *TypeFilter) AppliesToType(definition *astmodel.StructDefinition) bool {
	// TODO: also allow filtering on package group name
	result := filter.nameMatches(definition.Name()) &&
		filter.versionMatches(definition.StructReference.PackageReference.PackageName())

	return result
}

func (filter *TypeFilter) schemaMatches(schema string) bool {
	return filter.matches(filter.Schema, &filter.schemaRegex, schema)
}

func (filter *TypeFilter) versionMatches(version string) bool {
	return filter.matches(filter.Version, &filter.versionRegex, version)
}

func (filter *TypeFilter) nameMatches(name string) bool {
	return filter.matches(filter.Name, &filter.nameRegex, name)
}

func (filter *TypeFilter) matches(glob string, regex **regexp.Regexp, name string) bool {
	if glob == "" {
		return true
	}

	if *regex == nil {
		*regex = createGlobbingRegex(glob)
	}

	return (*regex).MatchString(name)
}

// create a regex that does globbing of names
// * and ? have their usual (DOS style) meanings as wildcards
func createGlobbingRegex(globbing string) *regexp.Regexp {
	g := regexp.QuoteMeta(globbing)
	g = strings.ReplaceAll(g, "\\*", ".*")
	g = strings.ReplaceAll(g, "\\?", ".")
	g = "^" + g + "$"
	return regexp.MustCompile(g)
}
