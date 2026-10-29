// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	machinelearningservices "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20240401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_MachineLearning_Registry_20240401_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	registry := &machinelearningservices.Registry{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("registry")),
		Spec: machinelearningservices.Registry_Spec{
			Identity: &machinelearningservices.ManagedServiceIdentity{
				Type: to.Ptr(machinelearningservices.ManagedServiceIdentityType_SystemAssigned),
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &machinelearningservices.Sku{
				Name: new("P3"),
				Tier: to.Ptr(machinelearningservices.SkuTier_Free),
			},
			PublicNetworkAccess: new("Disabled"),
			RegionDetails: []machinelearningservices.RegistryRegionArmDetails{
				{
					AcrDetails: []machinelearningservices.AcrDetails{
						{
							SystemCreatedAcrAccount: &machinelearningservices.SystemCreatedAcrAccount{
								AcrAccountName: new("myasoamltestacr"),
								AcrAccountSku:  new("Premium"),
							},
						},
					},
					Location: tc.AzureRegion,
					StorageAccountDetails: []machinelearningservices.StorageAccountDetails{
						{
							SystemCreatedStorageAccount: &machinelearningservices.SystemCreatedStorageAccount{
								AllowBlobPublicAccess: new(false),
								StorageAccountName:    new("myasoamltestsa"),
								StorageAccountType:    new("Standard_LRS"),
							},
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(registry)
	// Ensure that the status is what we expect
	tc.Expect(registry.Status.Id).ToNot(BeNil())
	armId := *registry.Status.Id

	tc.DeleteResourceAndWait(registry)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(machinelearningservices.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
