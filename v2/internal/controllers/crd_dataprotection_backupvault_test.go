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
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Test function Test_Dataprotection_Backupvault_CRUD which uses all available attributes in the Backupvault CRD

func Test_Dataprotection_Backupvault_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a Backupvault.
	backupvault := &dataprotection.BackupVault{
		
		ObjectMeta: tc.MakeObjectMetaWithName("backupvault"),
		Spec: dataprotection.BackupVault_Spec{
			ETag: to.StringPtr(""),
			Identity: &dataprotection.Identity{
				Type: dataprotection.IdentityType_SystemAssigned,
			},
			Location: tc.AzureRegion,
			Name:     to.StringPtr(""), // This is the name of the backup vault
			Properties: &dataprotection.BackupVaultProperties{
				monitorSettings: &dataprotection.MonitorSettings{
					azureMonitorAlertSettings: &dataprotection.AzureMonitorAlertSettings{
						alertsForAllJobFailures: to.BoolPtr(true),
					},
				},
				storageSettings : &dataprotection.StorageSettings{
					dataStoreType : dataprotection.DataStorageType_Local,
					type : dataprotection.StorageType_AzureStorage,
				},
			},
			SystemData: &dataprotection.SystemData{
				CreatedAt:      to.TimePtr(""),
				CreatedBy:      to.StringPtr(""),
				LastModifiedAt: to.TimePtr(""),
				LastModifiedBy: to.StringPtr(""),
			},
			Tags: map[string]string{"cheese": "blue"},
		},
	},

		tc.CreateResourceAndWait(backupvault)

	tc.Expect(backupvault.Status.ETag).ToNot(BeNil())
	tc.Expect(backupvault.Status.Identity.Type).To(Equal(dataprotection.IdentityType_SystemAssigned))
	tc.Expect(backupvault.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(backupvault.Status.Name).ToNot(BeNil())
	tc.Expect(backupvault.Status.Properties.MonitorSettings.AzureMonitorAlertSettings.AlertsForAllJobFailures).To(BeEquivalentTo(true))
	tc.Expect(backupvault.Status.Prperties.StorageSettings.DataStoreType).To(Equal(dataprotection.DataStorageType_Local))
	tc.Expect(backupvault.Status.Properties.StorageSettings.Type).To(Equal(dataprotection.StorageType_AzureStorage))
	tc.Expect(backupvault.Status.SystemData.CreatedAt).ToNot(BeNil())
	tc.Expect(backupvault.Status.SystemData.CreatedBy).ToNot(BeNil())
	tc.Expect(backupvault.Status.SystemData.LastModifiedAt).ToNot(BeNil())
	tc.Expect(backupvault.Status.SystemData.LastModifiedBy).ToNot(BeNil())
	tc.Expect(backupvault.Status.Tags).ToNot(BeNil())

	tc.Expect(backupvault.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(backupvault.Status.Sku.Name).To(BeEquivalentTo(dataprotection.Sku_Name_Standard))
	tc.Expect(backupvault.Status.Id).ToNot(BeNil())

	armId := *backupvault.Status.Id

}
