/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Dataprotection_Backupvault_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	//Create a backupvault
	region := tc.AzureRegion
	identityType := "SystemAssigned"
	alertsForAllJobFailures_Status := dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled
	StorageSetting_DatastoreType_Value := dataprotection.StorageSetting_DatastoreType_VaultStore
	StorageSetting_Type_Value := dataprotection.StorageSetting_Type_LocallyRedundant

	backupvault := &dataprotection.BackupVault{
		ObjectMeta: tc.MakeObjectMetaWithName("backupvault"),
		Spec: dataprotection.BackupVault_Spec{
			Location: region,
			Tags:     map[string]string{"cheese": "blue"},
			Owner:    testcommon.AsOwner(rg),
			Identity: &dataprotection.DppIdentityDetails{
				Type: &identityType,
			},
			Properties: &dataprotection.BackupVaultSpec{
				MonitoringSettings: &dataprotection.MonitoringSettings{
					AzureMonitorAlertSettings: &dataprotection.AzureMonitorAlertSettings{
						AlertsForAllJobFailures: &alertsForAllJobFailures_Status,
					},
				},
				StorageSettings: []dataprotection.StorageSetting{
					{
						DatastoreType: &StorageSetting_DatastoreType_Value,
						Type:          &StorageSetting_Type_Value,
					},
				},
			},
		},
	}
	tc.CreateResourceAndWait(backupvault)

	// Assert that the backupvault exists in Azure
	tc.Expect(backupvault.Status.Location).To(Equal(region))
	tc.Expect(backupvault.Status.Tags).To(Equal(map[string]string{"cheese": "blue"}))
	// tc.Expect(backupvault.Status.Identity.Type).To(Equal(amon4.Type))
	// tc.Expect(backupvault.Status.Properties.MonitoringSettings.AzureMonitorAlertSettings.AlertsForAllJobFailures).To(Equal(amon))
	// tc.Expect(backupvault.Status.Properties.StorageSettings[0].DatastoreType).To(Equal(amon2.DatastoreType))
	// tc.Expect(backupvault.Status.Properties.StorageSettings[0].Type).To(Equal(amon3.Type))

	// armId := *backupvault.Status.Id

	// tc.DeleteResourceAndWait(backupvault)

	// // Ensure that the resource group was really deleted in Azure
	// exists, _, err := tc.AzureClient.HeadByID(
	// 	tc.Ctx,
	// 	armId,
	// 	string(dataprotection.APIVersion_Value))
	// tc.Expect(err).ToNot(HaveOccurred())
	// tc.Expect(exists).To(BeFalse())

}
