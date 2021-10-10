package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
	"github.com/Azure/azure-service-operator/v2/internal/generator/test"
)

func TestGolden_GetStatusFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	getStatusFunction := NewGetStatusFunction(idFactory)

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status, getStatusFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetStatusFunction", resource)
}
