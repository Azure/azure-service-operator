package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

// Test_NewSpecConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToSpec() and
// ConvertFromSpec() functions are converting to/from spec types that aren't the hub  type
func Test_NewSpecConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
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
	// We omit the PropertyAssignment functions to reduce the amount of clutter, those are tested elsewhare
	injector := astmodel.NewFunctionInjector()
	personSpec2020, err = injector.Inject(personSpec2020, convertSpecTo, convertSpecFrom)
	g.Expect(err).To(Succeed())

	// Write to a file
	fileDef := test.CreateFileDefinition(personSpec2020)

	// When verifying the golden file, check that the implementations of ConvertSpecTo() and ConvertSpecFrom() are
	// what you expect - if you don't have expectations, check that they do the right thing.
	test.AssertFileGeneratesExpectedCode(t, fileDef, t.Name())
}

// Test_NewStatusConversionFunction_Conversion_GeneratesExpectedCode tests the code when the ConvertToStatus() and
// ConvertFromStatus() functions are converting to/from status types that aren't the hub  type
func Test_NewStatusConversionFunction_Conversion_GeneratesExpectedCode(t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personStatus2020 := test.CreateStatus(test.Pkg2020, "Person")

	// Create our downstream type
	personStatus2021 := test.CreateStatus(test.Pkg2021, "Person")

	// Create our hub type
	personStatus2022 := test.CreateStatus(test.Pkg2022, "Person")

	// Create Property Assignment functions
	types := make(astmodel.Types)
	types.AddAll(personStatus2020)
	types.AddAll(personStatus2021)

	conversionContext := conversions.NewPropertyConversionContext(types, idFactory)
	propertyAssignTo, err := NewPropertyAssignmentToFunction(personStatus2020, personStatus2021, idFactory, conversionContext)
	g.Expect(err).To(Succeed())

	propertyAssignFrom, err := NewPropertyAssignmentFromFunction(personStatus2020, personStatus2021, idFactory, conversionContext)
	g.Expect(err).To(Succeed())

	// Create Spec Conversion Functions
	convertSpecTo := NewStatusConversionFunction(personStatus2022.Name(), propertyAssignTo, idFactory)
	convertSpecFrom := NewStatusConversionFunction(personStatus2022.Name(), propertyAssignFrom, idFactory)

	// Inject these methods into personStatus2020
	// We omit the PropertyAssignment functions to reduce the amount of clutter, those are tested elsewhare
	injector := astmodel.NewFunctionInjector()
	personStatus2020, err = injector.Inject(personStatus2020, convertSpecTo, convertSpecFrom)
	g.Expect(err).To(Succeed())

	// Write to a file
	fileDef := test.CreateFileDefinition(personStatus2020)

	// When verifying the golden file, check that the implementations of ConvertStatusTo() and ConvertStatusFrom() are
	// what you expect - if you don't have expectations, check that they do the right thing.
	test.AssertFileGeneratesExpectedCode(t, fileDef, t.Name())
}
