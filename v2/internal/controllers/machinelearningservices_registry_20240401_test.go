// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"testing"

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
			DiscoveryUrl: nil,
			Identity: &machinelearningservices.ManagedServiceIdentity{
				Type: to.Ptr(machinelearningservices.ManagedServiceIdentityType_None),
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &machinelearningservices.Sku{
				Name: to.Ptr("P3"),
				Tier: to.Ptr(machinelearningservices.SkuTier_Free),
			},
		},
	}

	tc.CreateResourceAndWait(registry)
}
