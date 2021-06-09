package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func Test_OriginalGVKFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	originalGVKFunction := NewOriginalGVKFunction(idFactory)
	demoType := astmodel.NewObjectType().		WithFunction(originalGVKFunction)

	demoRef := astmodel.MakeLocalPackageReference("github.com/Azure/azure-service-operator/hack/generated/_apis", "Person", "vDemo")
	demoName := astmodel.MakeTypeName(demoRef, "Demo")

	demoDef := astmodel.MakeTypeDefinition(demoName, demoType)

	fileDef := test.CreateFileDefinition(demoDef)
	test.AssertFileGeneratesExpectedCode(t, fileDef, t.Name())
}
