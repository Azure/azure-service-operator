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
}

var _ fmt.Stringer = &TypeMatcher{}

// Initialize initializes the type matcher
func (typeMatcher *TypeMatcher) Initialize() error {
	typeMatcher.groupRegex = createGlobbingRegex(typeMatcher.Group)
	typeMatcher.versionRegex = createGlobbingRegex(typeMatcher.Version)
	typeMatcher.nameRegex = createGlobbingRegex(typeMatcher.Name)

	return nil
}

func (typeMatcher *TypeMatcher) groupMatches(schema string) bool {
	return typeMatcher.matches(typeMatcher.Group, &typeMatcher.groupRegex, schema)
}

func (typeMatcher *TypeMatcher) versionMatches(version string) bool {
	return typeMatcher.matches(typeMatcher.Version, &typeMatcher.versionRegex, version)
}

func (typeMatcher *TypeMatcher) nameMatches(name string) bool {
	return typeMatcher.matches(typeMatcher.Name, &typeMatcher.nameRegex, name)
}

func (typeMatcher *TypeMatcher) matches(glob string, regex **regexp.Regexp, name string) bool {
	if glob == "" {
		return true
	}

	if *regex == nil {
		*regex = createGlobbingRegex(glob)
	}

	return (*regex).MatchString(name)
}

// AppliesToType indicates whether this filter should be applied to the supplied type definition
func (typeMatcher *TypeMatcher) AppliesToType(typeName astmodel.TypeName) bool {
	if localRef, ok := typeName.PackageReference.AsLocalPackage(); ok {
		group := localRef.Group()
		version := localRef.Version()

		result := typeMatcher.groupMatches(group) &&
			typeMatcher.versionMatches(version) &&
			typeMatcher.nameMatches(typeName.Name())

		return result
	}

	// Never match external references
	return false
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
