/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// Test_PropertyChangesReporter_SaveReports_GivenTwoVersions_WritesReportsForPublicAndStorageVersions
// verifies that reports are generated in every package containing a non-hub resource.
func Test_PropertyChangesReporter_SaveReports_GivenTwoVersions_WritesReportsForPublicAndStorageVersions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person2020Spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	person2020Status := test.CreateStatus(test.Pkg2020, "Person")
	person2020 := test.CreateResource(test.Pkg2020, "Person", person2020Spec, person2020Status)

	person2021Spec := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.KnownAsProperty)
	person2021Status := test.CreateStatus(test.Pkg2021, "Person")
	person2021 := test.CreateResource(test.Pkg2021, "Person", person2021Spec, person2021Status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020Spec, person2020Status, person2020, person2021Spec, person2021Status, person2021)

	cfg := config.NewConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateStorageTypes(),
		CreateConversionGraph(cfg))
	g.Expect(err).NotTo(HaveOccurred())

	graph, err := GetStateData[*storage.ConversionGraph](state, ConversionGraphInfo)
	g.Expect(err).NotTo(HaveOccurred())

	reporter := NewPropertyChangesReporter(state.Definitions(), graph, cfg.ObjectModelConfiguration)

	outputFolder := t.TempDir()
	err = reporter.SaveReports(outputFolder)
	g.Expect(err).NotTo(HaveOccurred())

	// Public resources transition to their storage counterpart, so both public versions get reports.
	public2020Report := filepath.Join(outputFolder, test.Pkg2020.FolderPath(), "person-changes.md")
	g.Expect(public2020Report).To(BeAnExistingFile())

	content, err := os.ReadFile(public2020Report)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(content)).To(ContainSubstring(astmodel.CodeGenerationComments[0]))
	g.Expect(string(content)).To(ContainSubstring(test.Pkg2020.PackageName()))
	g.Expect(string(content)).To(ContainSubstring(test.Pkg2020s.PackageName()))
	g.Expect(string(content)).To(ContainSubstring("Person"))

	public2021Report := filepath.Join(outputFolder, test.Pkg2021.FolderPath(), "person-changes.md")
	g.Expect(public2021Report).To(BeAnExistingFile())

	// The earlier storage resource transitions to the later storage resource.
	storage2020Report := filepath.Join(outputFolder, test.Pkg2020s.FolderPath(), "person-changes.md")
	g.Expect(storage2020Report).To(BeAnExistingFile())

	// The final storage resource is the hub, so it has no report.
	storage2021Report := filepath.Join(outputFolder, test.Pkg2021s.FolderPath(), "person-changes.md")
	g.Expect(storage2021Report).NotTo(BeAnExistingFile())
}

// Test_PropertyChangesReporter_FindResources_ReturnsPublicAndStorageResources confirms that
// findResources() includes resources from both package kinds, while excluding non-resource types
// and resources from ARM, webhook, and compatibility packages.
func Test_PropertyChangesReporter_FindResources_ReturnsPublicAndStorageResources(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	person2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	personStatus2020 := test.CreateStatus(test.Pkg2020, "Person")
	personResource2020 := test.CreateResource(test.Pkg2020, "Person", person2020, personStatus2020)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(person2020, personStatus2020, personResource2020)

	cfg := config.NewConfiguration()

	state, err := RunTestPipeline(
		NewState(defs),
		CreateStorageTypes(),
		CreateConversionGraph(cfg))
	g.Expect(err).NotTo(HaveOccurred())

	graph, err := GetStateData[*storage.ConversionGraph](state, ConversionGraphInfo)
	g.Expect(err).NotTo(HaveOccurred())

	reportDefinitions := state.Definitions().Copy()
	excludedPackages := []astmodel.InternalPackageReference{
		astmodel.MakeSubPackageReference(astmodel.ARMPackageName, test.Pkg2020),
		astmodel.MakeSubPackageReference(astmodel.WebhookPackageName, test.Pkg2020),
		astmodel.MakeCompatPackageReference(test.Pkg2020),
	}
	for _, pkg := range excludedPackages {
		reportDefinitions.Add(personResource2020.WithName(personResource2020.Name().WithPackageReference(pkg)))
	}

	reporter := NewPropertyChangesReporter(reportDefinitions, graph, cfg.ObjectModelConfiguration)
	resources := reporter.findResources()

	g.Expect(resources).To(HaveLen(2))
	g.Expect(resources).To(ContainElement(personResource2020.Name()))
	g.Expect(resources).To(ContainElement(personResource2020.Name().WithPackageReference(test.Pkg2020s)))
}
