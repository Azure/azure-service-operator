/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package labels

import (
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/annotations"
)

type Label struct {
	Key   string
	Value string
}

// Parse parses a label. Amazingly there doesn't seem to be a function in client-go or similar that does this.
// There does exist an apimachinery labels.Parse but it parses label selectors not labels themselves.
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
