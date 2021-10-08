package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
	"github.com/Azure/azure-service-operator/v2/internal/generator/test"
)

func TestMarkLatestAPIVersionAsStorageVersion(t *testing.T) {
	g := NewGomegaWithT(t)

	// Test Resource V1
	specV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	statusV1 := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", specV1, statusV1)

	// Test Resource V2
	specV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty)
	statusV2 := test.CreateStatus(test.Pkg2021, "Person")
	resourceV2 := test.CreateResource(test.Pkg2021, "Person", specV2, statusV2)

	types := make(astmodel.Types)
	types.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2)

	initialState, err := RunTestPipeline(
		NewState(types),
	)
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		MarkLatestAPIVersionAsStorageVersion())
	g.Expect(err).To(Succeed())

	// Check that the expected types are flagged as storage types
	test.AssertPackagesGenerateExpectedCode(t, finalState.types, test.DiffWithTypes(initialState.Types()))
}
