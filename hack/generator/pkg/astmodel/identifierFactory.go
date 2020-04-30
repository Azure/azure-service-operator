/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"regexp"
	"strings"
)

var filterRegex *regexp.Regexp

// IdentifierFactory is a factory for creating Go identifiers from Json schema names
type IdentifierFactory interface {
	CreateIdentifier(name string) string
}

// identifierFactory is an implementation of the IdentifierFactory interface
type identifierFactory struct {
	renames map[string]string
}

// NewIdentifierFactory creates an IdentifierFactory ready for use
func NewIdentifierFactory() IdentifierFactory {
	return &identifierFactory{
		renames: createRenames(),
	}
}

// CreateIdentifier returns a valid Go public identifier
func (factory *identifierFactory) CreateIdentifier(name string) string {
	if identifier, ok := factory.renames[name]; ok {
		return identifier
	}

	// replace with spaces so titlecasing works nicely
	clean := filterRegex.ReplaceAllLiteralString(name, " ")

	titled := strings.Title(clean)
	result := strings.ReplaceAll(titled, " ", "")
	return result
}

func createRenames() map[string]string {
	return map[string]string{
		"$schema": "Schema",
	}
}

func init() {
	filterRegex = regexp.MustCompile("[$@._-]")
}
