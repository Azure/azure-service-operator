/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package vcr

import (
	"encoding/json"
	"regexp"
	"sort"
	"strings"
)

// policyTagRedactor removes the tags a tenant requires by policy on every resource group. Removing them
// keeps one tenant's policy out of the recording, and lets a recording made in such a tenant replay in one
// without the policy.
type policyTagRedactor struct {
	removals []*regexp.Regexp
}

// newPolicyTagRedactor returns nil when there are no tags to remove, which is the usual case and which
// makes hide a no-op.
func newPolicyTagRedactor(tags map[string]string) *policyTagRedactor {
	if len(tags) == 0 {
		return nil
	}

	keys := make([]string, 0, len(tags))
	for key := range tags {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	removals := make([]*regexp.Regexp, 0, 3*len(keys))
	for _, key := range keys {
		// Matching the value as well as the key leaves a tag a test sets for its own purposes alone, even
		// when the tenant's policy happens to mandate a tag of the same name. Both are matched as the
		// JSON they appear as, so that a value containing a quote or a brace still matches.
		for _, value := range jsonLiterals(tags[key]) {
			member := regexp.QuoteMeta(jsonLiteral(key)) + `\s*:\s*` + regexp.QuoteMeta(value)

			// Removing a member has to leave the object well formed whether it follows another member,
			// precedes one, or is the only one. Order matters: the last pattern alone would leave a
			// comma behind.
			removals = append(
				removals,
				regexp.MustCompile(`,\s*`+member),
				regexp.MustCompile(member+`\s*,`),
				regexp.MustCompile(member),
			)
		}
	}

	return &policyTagRedactor{
		removals: removals,
	}
}

// jsonLiteral returns s as it appears in JSON, quotes and all.
func jsonLiteral(s string) string {
	encoded, err := json.Marshal(s)
	if err != nil {
		// Marshalling a string cannot fail
		panic(err)
	}

	return string(encoded)
}

// jsonLiterals returns the forms s can take in JSON. Go's encoder escapes <, > and & by default, so a
// value containing one is written differently by the SDK than it is echoed back by ARM.
func jsonLiterals(s string) []string {
	escaped := jsonLiteral(s)

	var plain strings.Builder
	encoder := json.NewEncoder(&plain)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(s); err != nil {
		panic(err)
	}

	unescaped := strings.TrimSuffix(plain.String(), "\n")
	if unescaped == escaped {
		return []string{escaped}
	}

	return []string{escaped, unescaped}
}

// hide removes the policy tags from every tags object in s. Confining the removal to tags objects stops a
// tag key that happens to name another field from taking that field with it.
func (p *policyTagRedactor) hide(s string) string {
	if p == nil {
		return s
	}

	var result strings.Builder
	for {
		prefix := tagsObjectPrefix.FindStringIndex(s)
		if prefix == nil {
			break
		}

		open := prefix[1] - 1
		end, found := endOfJSONObject(s, open)
		if !found {
			break
		}

		result.WriteString(s[:open])
		result.WriteString(p.removeFrom(s[open:end]))
		s = s[end:]
	}

	result.WriteString(s)

	return result.String()
}

func (p *policyTagRedactor) removeFrom(object string) string {
	for _, removal := range p.removals {
		object = removal.ReplaceAllLiteralString(object, "")
	}

	return object
}

var tagsObjectPrefix = regexp.MustCompile(`"tags"\s*:\s*\{`)

// endOfJSONObject returns the index just past the brace closing the object that opens at open, ignoring
// any brace inside a string.
func endOfJSONObject(s string, open int) (int, bool) {
	depth := 0
	inString := false
	escaped := false

	for i := open; i < len(s); i++ {
		switch {
		case escaped:
			escaped = false
		case s[i] == '\\' && inString:
			escaped = true
		case s[i] == '"':
			inString = !inString
		case inString:
			// Braces inside a string don't nest the object
		case s[i] == '{':
			depth++
		case s[i] == '}':
			depth--
			if depth == 0 {
				return i + 1, true
			}
		}
	}

	return 0, false
}
