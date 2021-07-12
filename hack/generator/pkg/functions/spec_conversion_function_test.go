package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

// Test_SpecConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertTo() and
// ConvertFrom() functions are directly converting to/from the Hub type, without any intermediate step.
func Test_SpecConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personSpec2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)

	// Create our downstream type
	personSpec2021 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty, test.CityProperty)

	// Create our hub type
	personSpec2022 := test.CreateSpec(test.Pkg2022, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty, test.CityProperty)

	// Create Property Assignment functions
	types := make(astmodel.Types)
	types.AddAll(personSpec2020)
	types.AddAll(personSpec2021)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentToFunction(personSpec2020, personSpec2021, idFactory, conversionContext)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFromFunction(personSpec2020, personSpec2021, idFactory, conversionContext)
	g.Expect(err).To(Succeed())

	// Create Spec Conversion Functions
	convertSpecTo := NewSpecConversionFunction(personSpec2022.Name(), propertyAssignTo, idFactory)
	convertSpecFrom := NewSpecConversionFunction(personSpec2022.Name(), propertyAssignFrom, idFactory)

	// Inject these methods into personSpec2020
	injector := astmodel.NewFunctionInjector()
	personSpec2020, err = injector.Inject(personSpec2020, propertyAssignTo, propertyAssignFrom, convertSpecTo, convertSpecFrom)
	g.Expect(err).To(Succeed())

	// Write to a file
	fileDef := test.CreateFileDefinition(personSpec2020)
	test.AssertFileGeneratesExpectedCode(t, fileDef, t.Name())
}
