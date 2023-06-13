/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	devices "github.com/Azure/azure-service-operator/v2/api/devices/v1api20210702"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Devices_IotHub_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	iothub := &devices.IotHub{
		ObjectMeta: tc.MakeObjectMeta("iothub"),
		Spec: devices.IotHub_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &devices.IotHubSkuInfo{
				Capacity: to.Ptr(5),
				Name:     to.Ptr(devices.IotHubSkuInfo_Name_B1),
			},
			OperatorSpec: &devices.IotHubOperatorSpec{
				Secrets: &devices.IotHubOperatorSecrets{
					IotHubOwnerPrimaryKey: &genruntime.SecretDestination{Name: "newsecret", Key: "primaryConnectionString"},
				},
			},
		},
	}

	tc.CreateResourceAndWait(iothub)

	tc.Expect(iothub.Status.Id).ToNot(BeNil())
	armId := *iothub.Status.Id

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "WriteIotHubSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				IotHub_WriteSecrets(tc, iothub)
			},
		})

	tc.DeleteResourceAndWait(iothub)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(devices.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

}

func IotHub_WriteSecrets(tc *testcommon.KubePerTestContext, iotHub *devices.IotHub) {
	old := iotHub.DeepCopy()
	iotHubKeysSecret := "iothubkeyssecret"
	iotHub.Spec.OperatorSpec = &devices.IotHubOperatorSpec{
		Secrets: &devices.IotHubOperatorSecrets{
			DevicePrimaryKey:              &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "devicePrimaryKey"},
			DeviceSecondaryKey:            &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "deviceSecondaryKey"},
			IotHubOwnerPrimaryKey:         &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "iotHubOwnerPrimaryKey"},
			IotHubOwnerSecondaryKey:       &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "iotHubOwnerSecondaryKey"},
			RegistryReadPrimaryKey:        &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "registryReadPrimaryKey"},
			RegistryReadSecondaryKey:      &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "registryReadSecondaryKey"},
			RegistryReadWritePrimaryKey:   &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "registryReadWritePrimaryKey"},
			RegistryReadWriteSecondaryKey: &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "registryReadWriteSecondaryKey"},
			ServicePrimaryKey:             &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "servicePrimaryKey"},
			ServiceSecondaryKey:           &genruntime.SecretDestination{Name: iotHubKeysSecret, Key: "serviceSecondaryKey"},
		},
	}

	tc.PatchResourceAndWait(old, iotHub)
	tc.ExpectSecretHasKeys(
		iotHubKeysSecret,
		"devicePrimaryKey",
		"deviceSecondaryKey",
		"iotHubOwnerPrimaryKey",
		"iotHubOwnerSecondaryKey",
		"registryReadPrimaryKey",
		"registryReadSecondaryKey",
		"registryReadWritePrimaryKey",
		"registryReadWriteSecondaryKey",
		"servicePrimaryKey",
		"serviceSecondaryKey")

}
