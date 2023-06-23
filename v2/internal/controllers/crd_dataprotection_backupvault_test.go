/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	// The testing package is imported for testing-related functionality.
	"testing"

	// The gomega package is used for assertions and expectations in tests.
	. "github.com/onsi/gomega"

	// The dataprotection package contains types and functions related to dataprotection resources.
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	// The testcommon package includes common testing utilities.
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// The to package includes utilities for converting values to pointers.
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func newBackupVault(tc *testcommon.KubePerTestContext, rg *testcommon.TestResourceGroup, name string) *dataprotection.BackupVault {
	// rg := tc.CreateTestResourceGroupAndWait()

	backupVault := &dataprotection.BackupVault{
		ObjectMeta: tc.MakeObjectMeta(name),
		Spec: dataprotection.BackupVault_Spec{
			Location: tc.AzureRegion,
			Tags:     map[string]string{"cheese": "blue"},
			Owner:    testcommon.AsOwner(rg),
			Identity: &dataprotection.DppIdentityDetails{
				Type: to.Ptr("SystemAssigned"),
			},
			Properties: &dataprotection.BackupVaultSpec{
				MonitoringSettings: &dataprotection.MonitoringSettings{
					AzureMonitorAlertSettings: &dataprotection.AzureMonitorAlertSettings{
						AlertsForAllJobFailures: to.Ptr(dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled),
					},
				},
				StorageSettings: []dataprotection.StorageSetting{
					{
						DatastoreType: to.Ptr(dataprotection.StorageSetting_DatastoreType_VaultStore),
						Type:          to.Ptr(dataprotection.StorageSetting_Type_LocallyRedundant),
					},
				},
			},
		},
	}
	tc.CreateResourceAndWait(backupVault)

	return backupVault
}

func Test_Dataprotection_Backupvault_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	// backupvault := newBackupVault(tc, "asotestbackupvault")
	backupVault := newBackupVault(tc, rg, "asotestbackupvault")

	// Assertions and Expectations
	tc.Expect(backupVault.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(backupVault.Status.Tags).To(BeEquivalentTo(map[string]string{"cheese": "blue"}))
	tc.Expect(backupVault.Status.Identity.Type).To(BeEquivalentTo(to.Ptr("SystemAssigned")))
	tc.Expect(backupVault.Status.Properties.MonitoringSettings.AzureMonitorAlertSettings.AlertsForAllJobFailures).To(BeEquivalentTo(to.Ptr(dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled)))
	tc.Expect(backupVault.Status.Properties.StorageSettings[0].DatastoreType).To(BeEquivalentTo(to.Ptr(dataprotection.StorageSetting_DatastoreType_VaultStore)))
	tc.Expect(backupVault.Status.Properties.StorageSettings[0].Type).To(BeEquivalentTo(to.Ptr(dataprotection.StorageSetting_Type_LocallyRedundant)))
	tc.Expect(backupVault.Status.Id).ToNot(BeNil())

	armId := *backupVault.Status.Id

	// Perform a Simple Patch
	old := backupVault.DeepCopy()
	backupVault.Spec.Properties.MonitoringSettings.AzureMonitorAlertSettings.AlertsForAllJobFailures = to.Ptr(dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Disabled)

	tc.PatchResourceAndWait(old, backupVault)

	// Assertions and Expectations for Patch Operation
	tc.Expect(backupVault.Status.Properties.MonitoringSettings.AzureMonitorAlertSettings.AlertsForAllJobFailures).To(BeEquivalentTo(to.Ptr(dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Disabled)))

	// Delete the resource
	tc.DeleteResourceAndWait(backupVault)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(dataprotection.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
