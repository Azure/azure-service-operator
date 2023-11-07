/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	datafactory "github.com/Azure/azure-service-operator/v2/api/datafactory/v1api20180601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Data_Factory_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	globalParameterSpecification := datafactory.GlobalParameterSpecification{
		Type: to.Ptr(datafactory.GlobalParameterSpecification_Type_String),
		Value: map[string]v1.JSON{
			"foo": {
				Raw: []byte(`"value"`),
			},
		},
	}

	// Create a data factory instance
	factory := &datafactory.Factory{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("datafactory")),
		Spec: datafactory.Factory_Spec{
			Identity: &datafactory.FactoryIdentity{
				Type: to.Ptr(datafactory.FactoryIdentity_Type_SystemAssigned),
			},
			Location:         tc.AzureRegion,
			Owner:            testcommon.AsOwner(rg),
			GlobalParameters: map[string]datafactory.GlobalParameterSpecification{"testKey": globalParameterSpecification},
			Tags:             map[string]string{"cheese": "blue"},
		},
	}

	tc.CreateResourcesAndWait(factory)

	tc.Expect(factory.Status.Id).ToNot(BeNil())
	tc.Expect(factory.Status.GlobalParameters).To(HaveKey("testKey"))
	factoryArmId := *factory.Status.Id
	// Perform a simple patch
	old := factory.DeepCopy()
	factory.Spec.Tags["cheese"] = "époisses"
	tc.PatchResourceAndWait(old, factory)
	tc.Expect(factory.Status.Tags).To(Equal(map[string]string{"cheese": "époisses"}))

	tc.DeleteResourceAndWait(factory)
	// Ensure that the data factory was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		factoryArmId,
		string(datafactory.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
