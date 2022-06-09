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

func TestEnsureIndexIndexPropertyPathsUnique_ResolvesConflicts(t *testing.T) {
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
		{
			"Nested properties that conflict get prefixes",
			[]tc{
				{"Administrator.UserName", "AdministratorUserName"},
				{"Administrator.Password", "AdministratorPassword"},
				{"Service.UserName", "ServiceUserName"},
				{"Service.Password", "ServicePassword"},
			},
		},
		{
			"Parent properties that don't disambiguate are skipped",
			[]tc{
				{"Administrator.Properties.UserName", "AdministratorUserName"},
				{"Administrator.Properties.Password", "AdministratorPassword"},
				{"Service.Properties.UserName", "ServiceUserName"},
				{"Service.Properties.Password", "ServicePassword"},
			},
		},
		{
			"Other properties that share parents are prefixed for consistency",
			[]tc{
				{"Administrator.UserName", "AdministratorUserName"},
				{"Administrator.Password", "AdministratorPassword"},
				{"Administrator.Expiry", "AdministratorExpiry"},
				{"Service.UserName", "ServiceUserName"},
				{"Service.Password", "ServicePassword"},
			},
		},
		{
			"Original conflict needs to be resolved",
			[]tc{
				{"obj.Spec.Properties.HDInsight.Properties.AdministratorAccount.Password", "HDInsightPassword"},
				{"obj.Spec.Properties.VirtualMachine.Properties.AdministratorAccount.Password", "VirtualMachinePassword"},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			factory := make(testChainFactory)
			var chains []*propertyChain
			for _, tc := range c.chains {
				chains = append(chains, factory.createTestChain(tc.chain))
			}

			ensureIndexPropertyPathsUnique(chains)

			for i, chain := range chains {
				g.Expect(chain.indexPropertyPath()).To(Equal(c.chains[i].expected))
			}
		})
	}
}

// testChainFactory is a map of property chains to test chains, used to ensure the chains with a common prefix are
// using references to the same root property chain
type testChainFactory map[string]*propertyChain

func (factory testChainFactory) createTestChain(chain string) *propertyChain {
	// Return cached chain if present
	if c, ok := factory[chain]; ok {
		return c
	}

	var result *propertyChain
	if splitPoint := strings.LastIndex(chain, "."); splitPoint != -1 {
		// We have a multipart chain, so create the parent chain first
		parent := factory.createTestChain(chain[:splitPoint])
		name := chain[splitPoint+1:]
		prop := astmodel.NewPropertyDefinition(astmodel.PropertyName(name), strings.ToLower(name), astmodel.StringType)
		result = parent.add(prop)
	} else {
		// Top level property
		prop := astmodel.NewPropertyDefinition(astmodel.PropertyName(chain), strings.ToLower(chain), astmodel.StringType)
		result = newPropertyChain().add(prop)
	}

	factory[chain] = result
	return result
}
