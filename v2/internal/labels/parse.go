/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package labels

import (
	"strings"

	"github.com/rotisserie/eris"
	"k8s.io/apimachinery/pkg/util/validation"

	"github.com/Azure/azure-service-operator/v2/internal/annotations"
)

// Separators are the characters accepted between individual labels by ParseMap. ',' is the canonical
// separator for label lists, but ';' is accepted too so that label flags can be written the same way as
// flags which take a ';' delimited list.
const Separators = ",;"

type Label struct {
	Key   string
	Value string
}

// ParseMap parses a delimited string of key=value labels into a map, validating each label against the
// Kubernetes label syntax rules. Entries may be separated by any of the characters in Separators, and
// blank entries are ignored. Whitespace surrounding each key and value is trimmed, so that lists may be
// written with spaces after the separators. An error is returned if a label is malformed, invalid, or
// specified twice.
//
// Unlike Parse, this applies label validation rather than the more permissive annotation rules.
func ParseMap(value string) (map[string]string, error) {
	result := make(map[string]string)
	if value == "" {
		return result, nil
	}

	assignments := strings.FieldsFunc(value, func(r rune) bool {
		return strings.ContainsRune(Separators, r)
	})

	for _, assignment := range assignments {
		if strings.TrimSpace(assignment) == "" {
			continue
		}

		key, labelValue, found := strings.Cut(assignment, "=")
		key = strings.TrimSpace(key)
		labelValue = strings.TrimSpace(labelValue)
		if !found || key == "" {
			return nil, eris.Errorf("label %q must be in key=value form", assignment)
		}
		if errs := validation.IsQualifiedName(key); len(errs) > 0 {
			return nil, eris.Errorf("invalid label key %q: %s", key, strings.Join(errs, "; "))
		}
		if errs := validation.IsValidLabelValue(labelValue); len(errs) > 0 {
			return nil, eris.Errorf("invalid value for label %q: %s", key, strings.Join(errs, "; "))
		}
		if _, exists := result[key]; exists {
			return nil, eris.Errorf("label %q was specified more than once", key)
		}
		result[key] = labelValue
	}

	return result, nil
}

// Parse parses a label. Amazingly there doesn't seem to be a function in client-go or similar that does this.
// There does exist an apimachinery labels.Parse but it parses label selectors not labels themselves.
//
// Note that this currently applies annotation validation rules, which are more permissive than the rules
// Kubernetes applies to labels. Prefer ParseMap where strict label validation is wanted.
func Parse(s string) (Label, error) {
	// Currently the label restrictions are exactly the same as annotations,
	// so we can just re-use annotation parse here
	annotation, err := annotations.Parse(s)
	if err != nil {
		return Label{}, err
	}

	return Label{
		Key:   annotation.Key,
		Value: annotation.Value,
	}, nil
}

// ParseAll parses all the given labels and returns a collection of parsed labels
func ParseAll(labels []string) ([]Label, error) {
	result := make([]Label, 0, len(labels))

	for _, label := range labels {
		parsed, err := Parse(label)
		if err != nil {
			return nil, eris.Wrapf(err, "failed parsing %s", label)
		}
		result = append(result, parsed)
	}

	return result, nil
}
