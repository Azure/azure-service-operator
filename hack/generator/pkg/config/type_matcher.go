/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// TypeMatcher contains basic functionality for a filter
type TypeMatcher struct {
	// Group is a wildcard matching specifier for which groups are selected by this filter
	Group      string `yaml:",omitempty"`
	groupRegex *regexp.Regexp
	// Version is a wildcard matching specifier for which types are selected by this filter
	Version      string `yaml:",omitempty"`
	versionRegex *regexp.Regexp
	// Name is a wildcard matching specifier for which types are selected by this filter
	Name      string `yaml:",omitempty"`
	nameRegex *regexp.Regexp
	// Because is used to articulate why the filter applied to a type (used to generate explanatory logs in debug mode)
	Because string

	// MatchedTypes is the set of all type names matched by this matcher
	matchedTypes astmodel.TypeNameSet
}

var _ fmt.Stringer = &TypeMatcher{}

// Initialize initializes the type matcher
func (t *TypeMatcher) Initialize() error {
	t.groupRegex = createGlobbingRegex(t.Group)
	t.versionRegex = createGlobbingRegex(t.Version)
	t.nameRegex = createGlobbingRegex(t.Name)
	t.matchedTypes = make(astmodel.TypeNameSet)

	return nil
}

func (t *TypeMatcher) groupMatches(schema string) bool {
	return t.matches(t.Group, &t.groupRegex, schema)
}

func (t *TypeMatcher) versionMatches(version string) bool {
	return t.matches(t.Version, &t.versionRegex, version)
}

func (t *TypeMatcher) nameMatches(name string) bool {
	return t.matches(t.Name, &t.nameRegex, name)
}

func (t *TypeMatcher) matches(glob string, regex **regexp.Regexp, name string) bool {
	if glob == "" {
		return true
	}

	if *regex == nil {
		*regex = createGlobbingRegex(glob)
	}

	return (*regex).MatchString(name)
}

// AppliesToType indicates whether this filter should be applied to the supplied type definition
func (t *TypeMatcher) AppliesToType(typeName astmodel.TypeName) bool {
	if localRef, ok := typeName.PackageReference.AsLocalPackage(); ok {
		group := localRef.Group()
		version := localRef.Version()

		result := t.groupMatches(group) &&
			t.versionMatches(version) &&
			t.nameMatches(typeName.Name())

		// Track this match so we can later report if we didn't match anything
		if result {
			t.matchedTypes = t.matchedTypes.Add(typeName)
		}

		return result
	}

	// Never match external references
	return false
}

// HasMatches returns true if this matcher has ever matched at least 1 type
func (t *TypeMatcher) HasMatches() bool {
	return len(t.matchedTypes) > 0
}

// String returns a description of this filter
func (t *TypeMatcher) String() string {
	var result strings.Builder
	var spacer string
	if t.Group != "" {
		result.WriteString(fmt.Sprintf("Group: %q", t.Group))
		spacer = "; "
	}

	if t.Name != "" {
		result.WriteString(fmt.Sprintf("%sName: %q", spacer, t.Name))
		spacer = "; "
	}

	if t.Version != "" {
		result.WriteString(fmt.Sprintf("%sVersion: %q", spacer, t.Version))
	}

	return result.String()
}

// createGlobbingRegex creates a regex that does globbing of names
// * and ? have their usual (DOS style) meanings as wildcards
// Multiple wildcards can be separated with semicolons
func createGlobbingRegex(globbing string) *regexp.Regexp {
	if globbing == "" {
		// nil here as "" is fast-tracked elsewhere
		return nil
	}

	var regexes []string
	for _, glob := range strings.Split(globbing, ";") {
		g := regexp.QuoteMeta(glob)
		g = strings.ReplaceAll(g, "\\*", ".*")
		g = strings.ReplaceAll(g, "\\?", ".")
		g = "(^" + g + "$)"
		regexes = append(regexes, g)
	}

	// (?i) forces case insensitive matches
	regex := "(?i)" + strings.Join(regexes, "|")
	return regexp.MustCompile(regex)
}
