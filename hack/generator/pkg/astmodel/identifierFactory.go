/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"regexp"
	"strings"
	"unicode"
)

var filterRegex = regexp.MustCompile("[$@._-]")

// IdentifierFactory is a factory for creating Go identifiers from Json schema names
type IdentifierFactory interface {
	CreateIdentifier(name string) string
	CreatePackageNameFromVersion(version string) string
	CreateGroupName(name string) string
}

// identifierFactory is an implementation of the IdentifierFactory interface
type identifierFactory struct {
	renames map[string]string
}

// assert the implementation exists
var _ IdentifierFactory = (*identifierFactory)(nil)

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

func (factory *identifierFactory) CreatePackageNameFromVersion(version string) string {
	return "v" + sanitizePackageName(version)
}

func (factory *identifierFactory) CreateGroupName(group string) string {
	return strings.ToLower(group)
}

// sanitizePackageName removes all non-alphanum characters and converts to lower case
func sanitizePackageName(input string) string {
	var builder []rune

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder = append(builder, unicode.ToLower(rune(r)))
		}
	}

	return string(builder)
}
