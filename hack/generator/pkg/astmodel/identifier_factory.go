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

// \W is all non-word characters (https://golang.org/pkg/regexp/syntax/)
var filterRegex = regexp.MustCompile(`[\W_]`)

type Visibility string

const (
	Exported    = Visibility("exported")
	NotExported = Visibility("notexported")
)

// IdentifierFactory is a factory for creating Go identifiers from Json schema names
type IdentifierFactory interface {
	CreateIdentifier(name string, visibility Visibility) string
	CreatePropertyName(propertyName string, visibility Visibility) PropertyName
	CreatePackageNameFromVersion(version string) string
	CreateGroupName(name string) string
	// CreateEnumIdentifier generates the canonical name for an enumeration
	CreateEnumIdentifier(namehint string) string
}

// identifierFactory is an implementation of the IdentifierFactory interface
type identifierFactory struct {
	renames       map[string]string
	reservedWords map[string]string
}

// assert the implementation exists
var _ IdentifierFactory = (*identifierFactory)(nil)

// NewIdentifierFactory creates an IdentifierFactory ready for use
func NewIdentifierFactory() IdentifierFactory {
	return &identifierFactory{
		renames:       createRenames(),
		reservedWords: createReservedWords(),
	}
}

// CreateIdentifier returns a valid Go public identifier
func (factory *identifierFactory) CreateIdentifier(name string, visibility Visibility) string {
	if identifier, ok := factory.renames[name]; ok {
		// Just lowercase the first character according to visibility
		r := []rune(identifier)
		if visibility == NotExported {
			r[0] = unicode.ToLower(r[0])
		} else {
			r[0] = unicode.ToUpper(r[0])
		}
		return string(r)
	}

	// replace with spaces so titlecasing works nicely
	clean := filterRegex.ReplaceAllLiteralString(name, " ")

	cleanWords := sliceIntoWords(clean)
	var caseCorrectedWords []string
	for i, word := range cleanWords {
		if visibility == NotExported && i == 0 {
			caseCorrectedWords = append(caseCorrectedWords, strings.ToLower(word))
		} else {
			caseCorrectedWords = append(caseCorrectedWords, strings.Title(word))
		}
	}

	result := strings.Join(caseCorrectedWords, "")

	if alternateWord, ok := factory.reservedWords[result]; ok {
		// This is a reserved word, we need to use an alternate word
		return alternateWord
	}

	return result
}

func (factory *identifierFactory) CreatePropertyName(propertyName string, visibility Visibility) PropertyName {
	id := factory.CreateIdentifier(propertyName, visibility)
	return PropertyName(id)
}

func createRenames() map[string]string {
	return map[string]string{
		"$schema": "Schema",
		"*":       "Star", // This happens mostly in enums
	}
}

// These are words reserved by go, along with our chosen substitutes
func createReservedWords() map[string]string {
	return map[string]string{
		"break":       "brk",
		"case":        "c",
		"chan":        "chn",
		"const":       "cnst",
		"continue":    "cont",
		"default":     "def",
		"defer":       "deferVar",
		"else":        "els",
		"fallthrough": "fallthrgh",
		"for":         "f",
		"func":        "funcVar",
		"go":          "g",
		"goto":        "gotoVar",
		"if":          "ifVar",
		"import":      "imp",
		"interface":   "iface",
		"map":         "m",
		"package":     "pkg",
		"range":       "rng",
		"return":      "ret",
		"select":      "sel",
		"struct":      "strct",
		"switch":      "sw",
		"type":        "typeVar",
		"var":         "v",
	}
}

func (factory *identifierFactory) CreatePackageNameFromVersion(version string) string {
	return "v" + sanitizePackageName(version)
}

func (factory *identifierFactory) CreateGroupName(group string) string {
	return strings.ToLower(group)
}

func (factory *identifierFactory) CreateEnumIdentifier(namehint string) string {
	return factory.CreateIdentifier(namehint, Exported)
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

// transformToSnakeCase transforms a string LikeThis to a snake-case string like_this
func transformToSnakeCase(input string) string {
	words := sliceIntoWords(input)

	// my kingdom for LINQ
	var lowerWords []string
	for _, word := range words {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}

	return strings.Join(lowerWords, "_")
}

func simplifyName(context string, name string) string {
	contextWords := sliceIntoWords(context)
	nameWords := sliceIntoWords(name)

	var result []string
	for _, w := range nameWords {
		found := false
		for i, c := range contextWords {
			if c == w {
				found = true
				contextWords[i] = ""
				break
			}
		}
		if !found {
			result = append(result, w)
		}
	}

	if len(result) == 0 {
		return name
	}

	return strings.Join(result, "")
}

func sliceIntoWords(identifier string) []string {
	// Trim any leading and trailing spaces to make our life easier later
	identifier = strings.Trim(identifier, " ")

	var result []string
	chars := []rune(identifier)
	lastStart := 0
	for i := range chars {
		preceedingLower := i > 0 && unicode.IsLower(chars[i-1])
		succeedingLower := i+1 < len(chars) && unicode.IsLower(chars[i+1])
		isSpace := unicode.IsSpace(chars[i])
		foundUpper := unicode.IsUpper(chars[i])
		if i > lastStart && foundUpper && (preceedingLower || succeedingLower) {
			result = append(result, string(chars[lastStart:i]))
			lastStart = i
		} else if isSpace {
			r := string(chars[lastStart:i])
			r = strings.Trim(r, " ")
			// If r is entirely spaces... just don't append anything
			if len(r) != 0 {
				result = append(result, r)
			}
			lastStart = i + 1 // skip the space
		}
	}

	if lastStart < len(chars) {
		result = append(result, string(chars[lastStart:]))
	}

	return result
}
