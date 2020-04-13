// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"log"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
)

func TestLabelsToTags(t *testing.T) {
	checks := []struct {
		Name string
		In   map[string]string
		Out  map[string]*string
	}{
		{
			Name: "normal labels",
			In: map[string]string{
				"age":  "null",
				"type": "null",
			},
			Out: map[string]*string{
				"age":  to.StringPtr("null"),
				"type": to.StringPtr("null"),
			},
		},
		{
			Name: "backslash labels",
			In: map[string]string{
				"age/date":   "null",
				"type\\kind": "null",
				"fun/\\zone": "null",
			},
			Out: map[string]*string{
				"age.date":  to.StringPtr("null"),
				"type.kind": to.StringPtr("null"),
				"fun..zone": to.StringPtr("null"),
			},
		},
		{
			Name: "greater than less than labels",
			In: map[string]string{
				"age>date":  "null",
				"type<kind": "null",
				"fun<>zone": "null",
			},
			Out: map[string]*string{
				"age.date":  to.StringPtr("null"),
				"type.kind": to.StringPtr("null"),
				"fun..zone": to.StringPtr("null"),
			},
		},
		{
			Name: "question mark labels",
			In: map[string]string{
				"age?date":  "null",
				"type?kind": "null",
				"fun??zone": "null",
			},
			Out: map[string]*string{
				"age.date":  to.StringPtr("null"),
				"type.kind": to.StringPtr("null"),
				"fun..zone": to.StringPtr("null"),
			},
		},
		{
			Name: "percent labels",
			In: map[string]string{
				"age%date": "null",
				"%":        "null",
			},
			Out: map[string]*string{
				"age.date": to.StringPtr("null"),
				".":        to.StringPtr("null"),
			},
		},
	}
	for _, check := range checks {
		log.Println("Checking", check.Name)
		translated := LabelsToTags(check.In)
		for k, v := range translated {
			value := v
			if check.Out[k] == nil {
				t.Errorf("Expected 'null'\nGot nil\n%q", k)
			}
			if check.Out[k] != nil && *value != *check.Out[k] {
				t.Errorf("Expected %s\nGot %s", *check.Out[k], *value)
			}
		}
	}
}
