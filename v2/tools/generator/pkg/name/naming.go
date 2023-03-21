/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package name

import (
	"strings"
	"unicode"

	"github.com/gobuffalo/flect"
)

// flectOverrides is an ordered list of overrides for flect pluralization
// Earlier overrides take precedence over later ones
var flectOverrides = []struct {
	single string
	plural string
}{
	{"Redis", "Redis"},
	{"redis", "redis"},
	{"FIPS", "FIPS"},
	{"ID", "IDs"},
	{"Id", "Ids"},
	{"IP", "IPs"},
	{"knownAs", "knownAs"},
	{"ssh", "sshs"},
	{"SSH", "SSHs"},
}

// Singularize returns the singular form of the given name
func Singularize(name string) string {
	// work around bug in flect: https://github.com/Azure/azure-service-operator/issues/1454
	for _, o := range flectOverrides {
		if strings.HasSuffix(name, o.plural) {
			return name[0:len(name)-len(o.plural)] + o.single
		}
	}

	singular := flect.Singularize(name)

	// If name starts with an uppercase letter, ensure the result does too
	return fixCasing(name, singular)
}

// Pluralize returns the plural form of the given name
func Pluralize(name string) string {
	// work around bug in flect: https://github.com/Azure/azure-service-operator/issues/1454
	for _, o := range flectOverrides {
		plural := o.plural
		single := o.single
		if strings.HasSuffix(name, single) {
			return name[0:len(name)-len(single)] + plural
		}

		if strings.HasSuffix(name, plural) {
			return name
		}
	}

	plural := flect.Pluralize(name)

	// If name starts with an uppercase letter, ensure the result does too
	return fixCasing(name, plural)
}

func fixCasing(original string, result string) string {
	if len(original) == 0 || len(result) == 0 {
		return result
	}

	r := []rune(result)
	first := rune(original[0])
	if unicode.IsUpper(first) {
		r[0] = unicode.ToUpper(r[0])
	} else {
		r[0] = unicode.ToLower(r[0])
	}

	return string(r)
}
