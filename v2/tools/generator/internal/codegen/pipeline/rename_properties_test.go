package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func Test_RenameProperties_RenamesExpectedProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	personSpec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	personStatus := test.CreateStatus(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person := test.CreateResource(test.Pkg2020, "Person", personSpec, personStatus)

	defs := astmodel.MakeTypeDefinitionSetFromDefinitions(person, personSpec, personStatus)

	cfg := config.NewConfiguration()
	omc := cfg.ObjectModelConfiguration
	omc.ModifyProperty(
		personSpec.Name(),
		"KnownAs",
		func(p *config.PropertyConfiguration) error {
			p.RenameTo.Set("Alias")
			return nil
		})

	initialState, err := RunTestPipeline(
		NewState(defs))

	g.Expect(err).To(Succeed())
	finalState, err := RunTestPipeline(
		initialState,
		RenameProperties(omc))

	// When verifying the golden file, ensure the property has been renamed as expected
	test.AssertPackagesGenerateExpectedCode(t, finalState.definitions, test.DiffWithTypes(initialState.definitions))
}
