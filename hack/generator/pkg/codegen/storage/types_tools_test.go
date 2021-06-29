package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

var (
	testGroup = "microsoft.person"

	testPackage = test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty = astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	familyNameProperty = astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
		WithDescription("Shared name of the family")

	knownAsProperty = astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
		WithDescription("How the person is generally known")
)

func TestFindSpecTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := createTestSpec(testPackage, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := createTestStatus(testPackage, "Person")
	resource := createTestResource(testPackage, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	specs := FindSpecTypes(types)

	g.Expect(specs).To(HaveLen(1))
	g.Expect(specs.Contains(spec.Name())).To(BeTrue())
}

func TestFindStatusTypes(t *testing.T) {
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := createTestSpec(testPackage, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	status := createTestStatus(testPackage, "Person")
	resource := createTestResource(testPackage, "Person", spec, status)

	types := make(astmodel.Types)
	types.AddAll(resource, status, spec)

	statuses := FindStatusTypes(types)

	g.Expect(statuses).To(HaveLen(1))
	g.Expect(statuses.Contains(status.Name())).To(BeTrue())
}


func createTestResource(
	pkg astmodel.PackageReference,
	name string,
	spec astmodel.TypeDefinition,
	status astmodel.TypeDefinition) astmodel.TypeDefinition {

	return astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(pkg, name),
		astmodel.NewResourceType(spec.Name(), status.Name()))
}

func createTestSpec(
	pkg astmodel.PackageReference, name string, properties ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {
	specName := astmodel.MakeTypeName(pkg, name+"_Spec")
	return astmodel.MakeTypeDefinition(
		specName,
		astmodel.NewObjectType().WithProperties(properties...))
}

func createTestStatus(pkg astmodel.PackageReference, name string) astmodel.TypeDefinition {
	statusProperty := astmodel.NewPropertyDefinition("Status", "status", astmodel.StringType)
	statusName := astmodel.MakeTypeName(pkg, name+"_Status")
	return astmodel.MakeTypeDefinition(
		statusName,
		astmodel.NewObjectType().WithProperties(statusProperty))
}
