/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	dataprotectionstorage "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101storage"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// "github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Dataprotection_Backupvault_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	region := tc.AzureRegion

	// This code is meant for Monitoring Settings
	val1 := "Enabled"
	amons := &dataprotectionstorage.AzureMonitorAlertSettings{
		AlertsForAllJobFailures: &val1,
	}
	amon := &dataprotection.AzureMonitorAlertSettings{}
	err := amon.AssignProperties_From_AzureMonitorAlertSettings(amons)
	if err != nil {
		t.Fatalf("failed to assign properties from AzureMonitorAlertSettings: %v", err)
	}

	// This code is meant for Storage Settings
	val2 := "VaultStore"
	amons2 := &dataprotectionstorage.StorageSetting{
		DatastoreType: &val2,
	}
	amon2 := &dataprotection.StorageSetting{}
	err2 := amon2.AssignProperties_From_StorageSetting(amons2)
	if err2 != nil {
		t.Fatalf("failed to assign properties from StorageSetting: %v", err2)
	}

	val3 := "LocallyRedundant"
	amons3 := &dataprotectionstorage.StorageSetting{
		Type: &val3,
	}
	amon3 := &dataprotection.StorageSetting{}
	err3 := amon3.AssignProperties_From_StorageSetting(amons3)
	if err3 != nil {
		t.Fatalf("failed to assign properties from StorageSetting: %v", err3)
	}

	// This code is meant for Identity
	val4 := "SystemAssigned"
	amons4 := &dataprotectionstorage.DppIdentityDetails{
		Type: &val4,
	}
	amon4 := &dataprotection.DppIdentityDetails{}
	err4 := amon4.AssignProperties_From_DppIdentityDetails(amons4)
	if err4 != nil {
		t.Fatalf("failed to assign properties from DppIdentityDetails: %v", err4)
	}

	//Create a backupvault
	backupvault := &dataprotection.BackupVault{
		ObjectMeta: tc.MakeObjectMetaWithName("backupvault"),
		Spec: dataprotection.BackupVault_Spec{
			Location: region,
			Tags:     map[string]string{"cheese": "blue"},
			Owner:    testcommon.AsOwner(rg),
			Identity: &dataprotection.DppIdentityDetails{
				Type: amon4.Type,
			},
			Properties: &dataprotection.BackupVaultSpec{
				MonitoringSettings: &dataprotection.MonitoringSettings{
					AzureMonitorAlertSettings: amon,
				},
				StorageSettings: []dataprotection.StorageSetting{
					{
						DatastoreType: amon2.DatastoreType,
						Type:          amon3.Type,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(backupvault)

	// Assert that the backupvault exists in Azure
	tc.Expect(backupvault.Status.Location).To(Equal(region))
	tc.Expect(backupvault.Status.Tags).To(Equal(map[string]string{"cheese": "blue"}))
	tc.Expect(backupvault.Status.Identity.Type).To(Equal(amon4.Type))
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
