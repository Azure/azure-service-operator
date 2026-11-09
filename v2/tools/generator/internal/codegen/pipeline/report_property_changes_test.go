/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/storage"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

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
		CreateConversionGraph(cfg),
	)
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
