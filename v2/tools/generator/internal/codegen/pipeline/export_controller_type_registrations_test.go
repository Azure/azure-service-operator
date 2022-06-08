/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func TestEnsureIndexMethodNamesUniqueResolvesConflicts(t *testing.T) {
	t.Parallel()

	type tc struct {
		chain    string
		expected string
	}

	cases := []struct {
		name   string
		chains []tc
	}{
		{
			"Top level properties only",
			[]tc{
				{"FullName", "FullName"},
				{"FamilyName", "FamilyName"},
				{"KnownAs", "KnownAs"},
			},
		},
		{
			"Nested properties that don't conflict",
			[]tc{
				{"Administrator.UserName", "UserName"},
				{"Administrator.Password", "Password"},
				{"User.Role", "Role"},
				{"User.Created", "Created"},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			var chains []*propertyChain
			for _, tc := range c.chains {
				chains = append(chains, createTestChain(tc.chain))
			}

			ensureIndexMethodNamesUnique(chains)

			for i, chain := range chains {
				g.Expect(chain.indexPropertyPath()).To(Equal(c.chains[i].expected))
			}
		})
	}
}

func createTestChain(chain string) *propertyChain {
	props := strings.Split(chain, ".")
	var result *propertyChain
	for _, prop := range props {
		result = result.add(
			astmodel.NewPropertyDefinition(astmodel.PropertyName(prop), strings.ToLower(prop), astmodel.StringType))
	}

	return result
}
