/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"regexp"
	"strings"
	"sync"
	"unicode"

	"github.com/Azure/azure-service-operator/v2/internal/set"
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
	CreateGroupName(name string) string
	// CreateEnumIdentifier generates the canonical name for an enumeration
	CreateEnumIdentifier(namehint string) string
	// CreateLocal creates a local variable name
	CreateLocal(name string) string
	// CreateReceiver creates a name for a method receiver
	CreateReceiver(name string) string
}

// identifierFactory is an implementation of the IdentifierFactory interface
type identifierFactory struct {
	renames                   map[string]string
	reservedWords             map[string]string
	forbiddenReceiverSuffixes set.Set[string]

	idCache       idCache
	receiverCache map[string]string
	rwLock        sync.RWMutex
}

type idCacheKey struct {
	name       string
	visibility Visibility
}

type idCache map[idCacheKey]string

// assert the implementation exists
var _ IdentifierFactory = (*identifierFactory)(nil)

// NewIdentifierFactory creates an IdentifierFactory ready for use
func NewIdentifierFactory() IdentifierFactory {
	return &identifierFactory{
		renames:                   createRenames(),
		reservedWords:             createReservedWords(),
		idCache:                   make(idCache),
		receiverCache:             make(map[string]string),
		forbiddenReceiverSuffixes: createForbiddenReceiverSuffixes(),
	}
}

// CreateIdentifier returns a valid Go public identifier
func (factory *identifierFactory) CreateIdentifier(name string, visibility Visibility) string {
	cacheKey := idCacheKey{name, visibility}
	factory.rwLock.RLock()
	cached, ok := factory.idCache[cacheKey]
	factory.rwLock.RUnlock()
	if ok {
		return cached
	}

	result := factory.createIdentifierUncached(name, visibility)
	factory.rwLock.Lock()
	factory.idCache[cacheKey] = result
	factory.rwLock.Unlock()
	return result
}

func (factory *identifierFactory) createIdentifierUncached(name string, visibility Visibility) string {
	if identifier, ok := factory.renames[name]; ok {
		// Case the first character according to visibility
		r := []rune(identifier)
		if visibility == NotExported {
			r[0] = unicode.ToLower(r[0])
		} else {
			r[0] = unicode.ToUpper(r[0])
		}

		return string(r)
	}

	// replace non-word characters with spaces so title-casing works nicely
	parts := strings.Split(name, "_")
	for ix := range parts {
		clean := filterRegex.ReplaceAllLiteralString(parts[ix], " ")

		cleanWords := sliceIntoWords(clean)
		var caseCorrectedWords []string
		for i, word := range cleanWords {
			// Only alter the first word of the first part:
			if visibility == NotExported && i == 0 && ix == 0 {
				caseCorrectedWords = append(caseCorrectedWords, strings.ToLower(word))
			} else {
				// Disable lint: the suggested "replacement" for this in /x/cases has fundamental
				// differences in how it works (e.g. 'JSON' becomes 'Json'; we don’t want that).
				// Furthermore, the cases (ha) that it "fixes" are not relevant to us
				// (something about better handling of various punctuation characters;
				// our words are punctuation-free).
				//nolint:staticcheck
				caseCorrectedWords = append(caseCorrectedWords, strings.Title(word))
			}
		}

		parts[ix] = strings.Join(caseCorrectedWords, "")
	}

	result := strings.Join(parts, "_")

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

// CreateLocal creates a local variable identifier
func (factory *identifierFactory) CreateLocal(name string) string {
	return factory.CreateIdentifier(name, NotExported)
}

// CreateReceiver creates an identifier for a method receiver
func (factory *identifierFactory) CreateReceiver(name string) string {
	// Check the cache first
	factory.rwLock.RLock()
	result, found := factory.receiverCache[name]
	factory.rwLock.RUnlock()

	if found {
		return result
	}

	// Convert to a sequence of words
	clean := filterRegex.ReplaceAllLiteralString(name, " ")
	words := sliceIntoWords(clean)

	// Remove forbidden suffix words from the end
	for {
		if len(words) == 1 {
			break
		}

		last := len(words) - 1
		if !factory.forbiddenReceiverSuffixes.Contains(words[last]) {
			break
		}

		words = words[:last]
	}

	base := words[len(words)-1]

	// Prefix with a qualifying term if one is available,
	// AND either base is a reserved word, or it is too short (3 characters or less)
	if len(words) > 1 {
		if _, found := factory.reservedWords[strings.ToLower(base)]; found || len(base) <= 3 {
			base = words[len(words)-2] + base
		}
	}

	result = factory.CreateLocal(base)

	factory.rwLock.Lock()
	factory.receiverCache[name] = result
	factory.rwLock.Unlock()

	return result
}

func createRenames() map[string]string {
	return map[string]string{
		"$schema":    "Schema",
		"*":          "Star", // This happens mostly in enums
		"apiVersion": "APIVersion",
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

// createForbiddenReceiverSuffixes creates a case-sensitive list of words we don't want to use as receiver names
func createForbiddenReceiverSuffixes() set.Set[string] {
	return set.Make("STATUS", "Spec", "ARM")
}

func (factory *identifierFactory) CreateGroupName(group string) string {
	return strings.TrimPrefix(strings.ToLower(group), "microsoft.")
}

func (factory *identifierFactory) CreateEnumIdentifier(namehint string) string {
	return factory.CreateIdentifier(namehint, Exported)
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

// sliceIntoWords splits the provided identifier into a slice of individual words.
// A word is defined by one of the following:
//   1. A space ("a test" becomes "a" and "test")
//   2. A transition between lowercase and uppercase ("aWord" becomes "a" and "Word")
//   3. A transition between multiple uppercase letters and a lowercase letter ("XMLDocument" becomes "XML" and "Document")
//   4. A transition between a letter and a digit ("book12" becomes "book" and "12")
//   5. A transition between a digit and a letter ("12monkeys" becomes "12" and "monkeys")
func sliceIntoWords(identifier string) []string {
	// Trim any leading and trailing spaces to make our life easier later
	identifier = strings.Trim(identifier, " ")

	var result []string
	chars := []rune(identifier)
	lastStart := 0
	for i := range chars {
		preceedingLower := i > 0 && unicode.IsLower(chars[i-1])
		preceedingDigit := i > 0 && unicode.IsDigit(chars[i-1])
		succeedingLower := i+1 < len(chars) && unicode.IsLower(chars[i+1]) // This case is for handling acronyms like XMLDocument
		isSpace := unicode.IsSpace(chars[i])
		foundUpper := unicode.IsUpper(chars[i])
		foundDigit := unicode.IsDigit(chars[i])
		caseTransition := foundUpper && (preceedingLower || succeedingLower)
		digitTransition := (foundDigit && !preceedingDigit) || (!foundDigit && preceedingDigit)
		if isSpace {
			r := string(chars[lastStart:i])
			r = strings.Trim(r, " ")
			// If r is entirely spaces... just don't append anything
			if len(r) != 0 {
				result = append(result, r)
			}
			lastStart = i + 1 // skip the space
		} else if i > lastStart && (caseTransition || digitTransition) {
			result = append(result, string(chars[lastStart:i]))
			lastStart = i
		}
	}

	if lastStart < len(chars) {
		result = append(result, string(chars[lastStart:]))
	}

	return result
}
