/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// TestResourceVersionsReport_IsDeprecatedResource_HybridMigration verifies that `v1api`-prefixed
// resources in a Hybrid group whose migration hasn't shipped yet are NOT flagged as deprecated,
// since they're still the only variants users can use. See #5534.
func Test_resourceVersionsReportPartition_isDeprecatedResource_HybridMigration(t *testing.T) {
	t.Parallel()

	authorizationLegacyPkg := test.MakeLocalPackageReference("authorization", "v20220401").
		WithVersionPrefix(astmodel.GeneratorVersion) // v1api prefix
	authorizationLegacyItem := resourceVersionsReportItem{
		name: astmodel.MakeInternalTypeName(authorizationLegacyPkg, "RoleAssignment"),
	}

	// alertsmanagement is Hybrid but has no upcoming migration release registered (migration
	// happened long ago), so its v1api variants ARE deprecated.
	alertsLegacyPkg := test.MakeLocalPackageReference("alertsmanagement", "v20210401").
		WithVersionPrefix(astmodel.GeneratorVersion)
	alertsLegacyItem := resourceVersionsReportItem{
		name: astmodel.MakeInternalTypeName(alertsLegacyPkg, "SmartDetector"),
	}

	cases := map[string]struct {
		item     resourceVersionsReportItem
		expected bool
	}{
		"UpcomingHybridMigration_LegacyNotDeprecated": {
			item:     authorizationLegacyItem,
			expected: false,
		},
		"AlreadyReleasedHybridMigration_LegacyDeprecated": {
			item:     alertsLegacyItem,
			expected: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			// CurrentRelease matches the value used in azure-arm.yaml at the time this test was
			// written; the authorization migration (v2.21.0) is later than this.
			currentRelease := "v2.20.0"

			partition := &resourceVersionsReportPartition{}

			g.Expect(partition.isDeprecatedResource(currentRelease)(c.item)).To(Equal(c.expected))
		})
	}
}
