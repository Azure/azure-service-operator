/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_ReportAllResourceVersions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	gold := goldie.New(t)

	person2020desc := []string{
		"This is an older version",
		"of the Person resource",
	}
	person2020 := test.CreateResource(
		test.Pkg2020,
		"Person",
		test.CreateSpec(test.Pkg2020, "Person"),
		test.CreateStatus(test.Pkg2020, "Person"),
	).
		WithDescription(person2020desc...)

	address2020 := test.CreateResource(
		test.Pkg2020,
		"Address",
		test.CreateSpec(test.Pkg2020, "Address"),
		test.CreateStatus(test.Pkg2020, "Address"),
	)

	person2021desc := []string{
		"This is a newer version",
		"of the Person resource",
	}
	person2021 := test.CreateResource(
		test.Pkg2021,
		"Person",
		test.CreateSpec(test.Pkg2021, "Person"),
		test.CreateStatus(test.Pkg2021, "Person"),
	).
		WithDescription(person2021desc...)

	address2021 := test.CreateResource(
		test.Pkg2021,
		"Address",
		test.CreateSpec(test.Pkg2021, "Address"),
		test.CreateStatus(test.Pkg2021, "Address"),
	)

	batch2021 := test.CreateResource(
		test.BatchPkg2021,
		"BatchAccount",
		test.CreateSpec(test.BatchPkg2021, "BatchAccount"),
		test.CreateStatus(test.BatchPkg2021, "BatchAccount"),
	)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, address2020, person2021, address2021, batch2021)

	// utility function used to configure a which ASO version from which a resource was supported
	supportedFrom := func(from string) func(tc *config.TypeConfiguration) error {
		return func(tc *config.TypeConfiguration) error {
			tc.SupportedFrom.Set(from)
			return nil
		}
	}

	cfg := config.NewConfiguration()
	cfg.RootURL = "https://github.com/Azure/azure-service-operator/tree/main/v2"
	cfg.SamplesPath = "../../../../../samples"

	omc := cfg.ObjectModelConfiguration
	g.Expect(omc.ModifyType(person2020.Name(), supportedFrom("beta.0"))).To(Succeed())
	g.Expect(omc.ModifyType(address2020.Name(), supportedFrom("beta.0"))).To(Succeed())
	g.Expect(omc.ModifyType(person2021.Name(), supportedFrom("beta.2"))).To(Succeed())
	g.Expect(omc.ModifyType(address2021.Name(), supportedFrom("beta.2"))).To(Succeed())
	g.Expect(omc.ModifyType(batch2021.Name(), supportedFrom("beta.2"))).To(Succeed())

	report, err := NewResourceVersionsReport(defs, cfg)
	g.Expect(err).ToNot(HaveOccurred())

	var buffer strings.Builder
	g.Expect(report.WriteAllResourcesReportToBuffer(
		"", // No Frontmatter
		&buffer,
	)).
		To(Succeed())

	gold.Assert(t, t.Name(), []byte(buffer.String()))
}

func TestResourceVersionsReportGroupInfo_GivenGroup_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	storagePkg := test.MakeLocalPackageReference("storage", "v20230101")

	storageAccount := resourceVersionsReportItem{
		name:          astmodel.MakeInternalTypeName(storagePkg, "StorageAccount"),
		armType:       "Microsoft.Storage/storageAccounts",
		armVersion:    "2023-01-01",
		supportedFrom: "v2.0.0",
	}

	alertsManagementPkg := test.MakeLocalPackageReference("alertsmanagement", "v20210401")

	smartDetector := resourceVersionsReportItem{
		name:          astmodel.MakeInternalTypeName(alertsManagementPkg, "SmartDetector"),
		armType:       "microsoft.alertsManagement/smartDetectorAlertRules",
		armVersion:    "2021-04-01",
		supportedFrom: "v2.11.0",
	}

	prometheusRuleGroup := resourceVersionsReportItem{
		name:    astmodel.MakeInternalTypeName(alertsManagementPkg, "PrometheusRuleGroup"),
		armType: "Microsoft.AlertsManagement/prometheusRuleGroups",
	}

	cases := map[string]struct {
		group            string
		items            set.Set[resourceVersionsReportItem]
		expectedGroup    string
		expectedProvider string
		expectedTitle    string
	}{
		"StorageAccount": {
			group:            "storage",
			items:            set.Make(storageAccount),
			expectedGroup:    "storage",
			expectedProvider: "Microsoft.Storage",
			expectedTitle:    "Storage",
		},
		"SmartDetector": {
			group:            "alertsmanagement",
			items:            set.Make(smartDetector),
			expectedGroup:    "alertsmanagement",
			expectedProvider: "Microsoft.alertsManagement",
			expectedTitle:    "AlertsManagement",
		},
		"PrometheusRuleGroup": {
			group:            "alertsmanagement",
			items:            set.Make(prometheusRuleGroup),
			expectedGroup:    "alertsmanagement",
			expectedProvider: "Microsoft.AlertsManagement",
			expectedTitle:    "AlertsManagement",
		},
		"Prefers Correct Case": {
			group:            "alertsmanagement",
			items:            set.Make(smartDetector, prometheusRuleGroup),
			expectedGroup:    "alertsmanagement",
			expectedProvider: "Microsoft.AlertsManagement",
			expectedTitle:    "AlertsManagement",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			report := &ResourceVersionsReport{} // empty report

			info := report.groupInfo(c.group, c.items)
			g.Expect(info).ToNot(BeNil())
			g.Expect(info.Group).To(Equal(c.expectedGroup))
			g.Expect(info.Provider).To(Equal(c.expectedProvider))
			g.Expect(info.Title).To(Equal(c.expectedTitle))
		})
	}
}

// TestResourceVersionsReport_SupportedFrom_OverridesForHybridMigrationGroup verifies that we
// report the correct "Supported From" version for resources in a group whose migration to
// Hybrid versioning has not yet been released. See #5534.
func TestResourceVersionsReport_SupportedFrom_OverridesForHybridMigrationGroup(t *testing.T) {
	t.Parallel()

	// authorization is registered in versionMigrationHybridReleases with an upcoming migration
	// release, so pre-migration resources in that group should be reported as available in that
	// upcoming release (not their original release).
	authorizationHybridRelease, ok := astmodel.HybridMigrationReleaseForGroup("authorization")
	NewGomegaWithT(t).Expect(ok).To(BeTrue(), "authorization must be registered as a hybrid migration group")

	newStylePkg := test.MakeLocalPackageReference("authorization", "v20220401")
	legacyStylePkg := newStylePkg.WithVersionPrefix(astmodel.GeneratorVersion) // v1api prefix
	cdnPkg := test.MakeLocalPackageReference("cdn", "v20230501").WithVersionPrefix(astmodel.GeneratorVersion)

	roleAssignmentNewStyle := astmodel.MakeInternalTypeName(newStylePkg, "RoleAssignment")
	roleAssignmentLegacyStyle := astmodel.MakeInternalTypeName(legacyStylePkg, "RoleAssignment")
	cdnProfile := astmodel.MakeInternalTypeName(cdnPkg, "Profile")

	cases := map[string]struct {
		name           astmodel.InternalTypeName
		configuredFrom string
		expected       string
	}{
		"NewStyleAuthorizationResource_UsesMigrationRelease": {
			// The `v`-prefixed variant of a resource introduced pre-v2.17 in an unreleased hybrid group
			// should have its Supported From bumped to the migration release.
			name:           roleAssignmentNewStyle,
			configuredFrom: "v2.4.0",
			expected:       authorizationHybridRelease,
		},
		"LegacyStyleAuthorizationResource_KeepsOriginal": {
			// The `v1api`-prefixed variant retains the original supportedFrom because it has always
			// been available at that release.
			name:           roleAssignmentLegacyStyle,
			configuredFrom: "v2.4.0",
			expected:       "v2.4.0",
		},
		"NonHybridGroup_KeepsOriginal": {
			// cdn is Legacy (not Hybrid) so the original supportedFrom is preserved regardless.
			name:           cdnProfile,
			configuredFrom: "v2.0.0",
			expected:       "v2.0.0",
		},
		"PreGARelease_MappedToV200": {
			// v2.0.0-alpha.X strings continue to be normalized to v2.0.0 for older legacy resources.
			name:           roleAssignmentLegacyStyle,
			configuredFrom: "v2.0.0-alpha.2",
			expected:       "v2.0.0",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			cfg := config.NewConfiguration()
			g.Expect(cfg.ObjectModelConfiguration.ModifyType(
				c.name,
				func(tc *config.TypeConfiguration) error {
					tc.SupportedFrom.Set(c.configuredFrom)
					return nil
				},
			)).To(Succeed())

			report := &ResourceVersionsReport{
				reportConfiguration:      cfg.SupportedResourcesReport,
				objectModelConfiguration: cfg.ObjectModelConfiguration,
			}

			g.Expect(report.supportedFrom(c.name)).To(Equal(c.expected))
		})
	}
}
