package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestCreateStorageTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	testGroup := "microsoft.person"

	packageV1 := test.MakeLocalPackageReference(testGroup, "v20200101")

	// Properties for v1

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	familyNameProperty := astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
		WithDescription("Shared name of the family")

	knownAsProperty := astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
		WithDescription("How the person is generally known")

	fullAddressProperty := astmodel.NewPropertyDefinition("FullAddress", "fullAddress", astmodel.StringType).
		WithDescription("Full written address for map or postal use")

	cityProperty := astmodel.NewPropertyDefinition("City", "city", astmodel.StringType).
		WithDescription("City or town (or nearest)")

	// Properties for v2

	packageV2 := test.MakeLocalPackageReference(testGroup, "v20211231")

	addressV2 := test.CreateObjectDefinition(packageV2, "Address", fullAddressProperty, cityProperty)

	residentialAddress := astmodel.NewPropertyDefinition("ResidentialAddress", "residentialAddress", addressV2.Name())

	postalAddress := astmodel.NewPropertyDefinition("PostalAddress", "postalAddress", addressV2.Name())

	// Test Resource V1

	specV1 := test.CreateSpec(packageV1, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	statusV1 := test.CreateStatus(packageV1, "Person")
	resourceV1 := test.CreateResource(packageV1, "Person", specV1, statusV1)

	// Test Resource V2

	specV2 := test.CreateSpec(packageV2, "Person",
		fullNameProperty, familyNameProperty, knownAsProperty, residentialAddress, postalAddress)
	statusV2 := test.CreateStatus(packageV2, "Person")
	resourceV2 := test.CreateResource(packageV2, "Person", specV2, statusV2)

	types := make(astmodel.Types)
	types.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2, addressV2)

	// Run the stage

	graph := storage.NewConversionGraph()
	createStorageTypes := CreateStorageTypes(graph, idFactory)

	// Don't need a context when testing
	finalTypes, err := createStorageTypes.Run(context.TODO(), types)

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalTypes, t.Name())
}
