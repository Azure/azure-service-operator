package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestInjectOriginalVersionFunction(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	testGroup := "microsoft.person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	familyNameProperty := astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
		WithDescription("Shared name of the family")

	knownAsProperty := astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
		WithDescription("How the person is generally known")

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	injectOriginalVersion := InjectOriginalVersionFunction(idFactory)

	// Don't need a context when testing
	finalTypes, err := injectOriginalVersion.Run(context.TODO(), types)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalTypes, t.Name())
}
