/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package annotations

import (
	"strings"

	"github.com/rotisserie/eris"
)

type Annotation struct {
	Key   string
	Value string
}

// Parse parses an annotation. Amazingly there doesn't seem to be a function in client-go or similar that does this
func Parse(s string) (Annotation, error) {
	split := strings.Split(s, "=")
	if len(split) != 2 {
		return Annotation{}, eris.Errorf("%s must have two parts separated by '='", s)
	}

	// see https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set
	key := split[0]
	value := split[1]

	if len(key) == 0 {
		return Annotation{}, eris.New("key can't be length 0")
	}

	keySplit := strings.Split(key, "/")
	if len(keySplit) > 2 {
		return Annotation{}, eris.Errorf("key %s must contain only a single '/'", key)
	}

	var name string
	var prefix string
	if len(keySplit) == 1 {
		name = key
	} else {
		// Len == 2
		prefix = keySplit[0]
		name = keySplit[1]

	}

	if len(key) > 63 {
		return Annotation{}, eris.Errorf("name %s must be 63 characters or less", name)
	}

	if len(prefix) > 253 {
		return Annotation{}, eris.Errorf("prefix %s must be 253 characters or less", prefix)
	}

	// TODO: Could enforce character restrictions too but not bothering for now

	return Annotation{
		Key:   key,
		Value: value,
	}, nil
}

// ParseAll parses all the given annotations and returns a collection of parsed annotations
func ParseAll(annotations []string) ([]Annotation, error) {
	result := make([]Annotation, 0, len(annotations))

	for _, annotation := range annotations {
		parsed, err := Parse(annotation)
		if err != nil {
			return nil, eris.Wrapf(err, "failed parsing %s", annotation)
		}
		result = append(result, parsed)
	}

	return result, nil
}
