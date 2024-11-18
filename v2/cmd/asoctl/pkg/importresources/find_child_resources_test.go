/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources_test

import (
	"strings"
	"testing"

	. "github.com/Azure/azure-service-operator/v2/cmd/asoctl/pkg/importresources"
	. "github.com/onsi/gomega"

	cdn "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501"
	dbforpostgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20221201"
)

func Test_FindChildResources_GivenParentResource_ReturnsExpectedChildResources(t *testing.T) {
	t.Parallel()

	postgreSQLServer := &dbforpostgresql.FlexibleServer{}
	postgreSQLConfiguration := &dbforpostgresql.FlexibleServersConfiguration{}

	ruleset := &cdn.RuleSet{}
	rule := &cdn.Rule{}

	cases := map[string]struct {
		parentResourceType        string
		expectedChildResourceType string
	}{
		"PostgreSQL servers have configurations": {
			parentResourceType:        postgreSQLServer.GetType(),
			expectedChildResourceType: postgreSQLConfiguration.GetType(),
		},
		"CDN rulesets have rules": {
			parentResourceType:        ruleset.GetType(),
			expectedChildResourceType: rule.GetType(),
		},
		"Can find PostgreSQL configurations when parent is lowercase": {
			parentResourceType:        strings.ToLower(postgreSQLServer.GetType()),
			expectedChildResourceType: postgreSQLConfiguration.GetType(),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			childResources := FindChildResourcesForResourceType(c.parentResourceType)

			g.Expect(childResources).NotTo(BeEmpty())
			g.Expect(childResources).To(ContainElement(c.expectedChildResourceType))
		})
	}
}
