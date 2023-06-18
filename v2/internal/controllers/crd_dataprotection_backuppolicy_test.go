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
)

func Test_Dataprotection_Backuppolicy_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// tc := globalTestContext.ForTest(t) initializes the test context for this test.
	// The globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)

	// rg := tc.CreateTestResourceGroupAndWait() creates a test resource group and waits until the operation is completed.
	rg := tc.CreateTestResourceGroupAndWait()

	// Consts for BackupVault
	identityType := "SystemAssigned"
	alertsForAllJobFailures_Status := dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled
	StorageSetting_DatastoreType_Value := dataprotection.StorageSetting_DatastoreType_VaultStore
	StorageSetting_Type_Value := dataprotection.StorageSetting_Type_LocallyRedundant

	// Create a backupvault
	backupvault := &dataprotection.BackupVault{
		ObjectMeta: tc.MakeObjectMetaWithName("asotestbackupvault"),
		Spec: dataprotection.BackupVault_Spec{
			Location: tc.AzureRegion,
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

	// Note:
	// It is mandatory to create a backupvault before creating a backuppolicy

	// Consts for BackupPolicy
	backupPolicy_ObjectType := dataprotection.BackupPolicy_ObjectType_BackupPolicy

	// Consts for BackupPolicy:AzureBackupRule
	AzureBackRule_Name := "BackupHourly"
	AzureBackupRule_ObjectType := dataprotection.AzureBackupRule_ObjectType_AzureBackupRule

	AzureBackupParams_BackupType_Value := "Incremental"
	AzureBackupParams_ObjectType_Value := dataprotection.AzureBackupParams_ObjectType_AzureBackupParams

	DataStore_DataStoreType_Value := dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore
	DataStore_ObjectType_Value := "DataStoreInfoBase"

	Schedule_ObjectType_Value := dataprotection.ScheduleBasedTriggerContext_ObjectType_ScheduleBasedTriggerContext
	Schedule_Timezone_Value := "UTC"

	TaggingCriteria_isDefault_Value := true
	TaggingCriteria_TaggingPriority_Value := 99
	TaggingCriteria_TagInfo_TagName_Value := "Default"

	// Consts for BackupPolicy:AzureRetentionRule
	AzureRetentionRule_Name := "Default"
	AzureRetentionRule_ObjectType := dataprotection.AzureRetentionRule_ObjectType_AzureRetentionRule
	AzureRetentionRule_IsDefault := true

	AzureRetentionRule_Lifecycles_DeleteAfter_Duration := "P9D"
	AzureRetentionRule_Lifecycles_DeleteAfter_ObjectType := dataprotection.AbsoluteDeleteOption_ObjectType_AbsoluteDeleteOption
	AzureRetentionRule_Lifecycles_SourceDataStore_DataStoreType := dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore
	AzureRetentionRule_Lifecycles_SourceDataStore_ObjectType := "DataStoreInfoBase"

	// Create a BackupPolicy
	backuppolicy := &dataprotection.BackupVaultsBackupPolicy{
		ObjectMeta: tc.MakeObjectMeta("asotestbackuppolicy"),
		Spec: dataprotection.BackupVaults_BackupPolicy_Spec{
			Owner: testcommon.AsOwner(backupvault),
			Properties: &dataprotection.BaseBackupPolicy{
				BackupPolicy: &dataprotection.BackupPolicy{
					DatasourceTypes: []string{"Microsoft.ContainerService/managedClusters"},
					ObjectType:      &backupPolicy_ObjectType,
					PolicyRules: []dataprotection.BasePolicyRule{
						{
							AzureBackup: &dataprotection.AzureBackupRule{
								Name:       &AzureBackRule_Name,
								ObjectType: &AzureBackupRule_ObjectType,
								BackupParameters: &dataprotection.BackupParameters{
									AzureBackupParams: &dataprotection.AzureBackupParams{
										BackupType: &AzureBackupParams_BackupType_Value,
										ObjectType: &AzureBackupParams_ObjectType_Value,
									},
								},
								DataStore: &dataprotection.DataStoreInfoBase{
									DataStoreType: &DataStore_DataStoreType_Value,
									ObjectType:    &DataStore_ObjectType_Value,
								},
								Trigger: &dataprotection.TriggerContext{
									Schedule: &dataprotection.ScheduleBasedTriggerContext{
										ObjectType: &Schedule_ObjectType_Value,
										Schedule: &dataprotection.BackupSchedule{
											RepeatingTimeIntervals: []string{"R/2023-06-07T10:26:32+00:00/PT4H"},
											TimeZone:               &Schedule_Timezone_Value,
										},
										TaggingCriteria: []dataprotection.TaggingCriteria{
											{
												IsDefault:       &TaggingCriteria_isDefault_Value,
												TaggingPriority: &TaggingCriteria_TaggingPriority_Value,
												TagInfo: &dataprotection.RetentionTag{
													TagName: &TaggingCriteria_TagInfo_TagName_Value,
												},
											},
										},
									},
								},
							},
						},
						{
							AzureRetention: &dataprotection.AzureRetentionRule{
								Name:       &AzureRetentionRule_Name,
								ObjectType: &AzureRetentionRule_ObjectType,
								IsDefault:  &AzureRetentionRule_IsDefault,
								Lifecycles: []dataprotection.SourceLifeCycle{
									{
										DeleteAfter: &dataprotection.DeleteOption{
											AbsoluteDeleteOption: &dataprotection.AbsoluteDeleteOption{
												Duration:   &AzureRetentionRule_Lifecycles_DeleteAfter_Duration,
												ObjectType: &AzureRetentionRule_Lifecycles_DeleteAfter_ObjectType,
											},
										},
										SourceDataStore: &dataprotection.DataStoreInfoBase{
											DataStoreType: &AzureRetentionRule_Lifecycles_SourceDataStore_DataStoreType,
											ObjectType:    &AzureRetentionRule_Lifecycles_SourceDataStore_ObjectType,
										},
										TargetDataStoreCopySettings: []dataprotection.TargetCopySetting{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(backuppolicy)

	// Assertions and Expectations
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.DatasourceTypes).To(BeEquivalentTo([]string{"Microsoft.ContainerService/managedClusters"}))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.ObjectType).To(BeEquivalentTo(&backupPolicy_ObjectType))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup.Name).To(BeEquivalentTo(&AzureBackRule_Name))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup.ObjectType).To(BeEquivalentTo(&AzureBackupRule_ObjectType))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule:BackupParameters
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup.BackupParameters.AzureBackupParams.BackupType).To(BeEquivalentTo(&AzureBackupParams_BackupType_Value))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup.BackupParameters.AzureBackupParams.ObjectType).To(BeEquivalentTo(&AzureBackupParams_ObjectType_Value))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule:DataStore
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup.DataStore.DataStoreType).To(BeEquivalentTo(&DataStore_DataStoreType_Value))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup.DataStore.ObjectType).To(BeEquivalentTo(&DataStore_ObjectType_Value))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule:Trigger
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup.Trigger.Schedule.ObjectType).To(BeEquivalentTo(&Schedule_ObjectType_Value))

	// Assertions and Expectations for BackupPolicy:AzureRetentionRule
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention.Name).To(BeEquivalentTo(&AzureRetentionRule_Name))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention.ObjectType).To(BeEquivalentTo(&AzureRetentionRule_ObjectType))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention.IsDefault).To(BeEquivalentTo(&AzureRetentionRule_IsDefault))

	// Assertions and Expectations for BackupPolicy:AzureRetentionRule:Lifecycles
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention.Lifecycles[0].DeleteAfter.AbsoluteDeleteOption.Duration).To(BeEquivalentTo(&AzureRetentionRule_Lifecycles_DeleteAfter_Duration))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention.Lifecycles[0].DeleteAfter.AbsoluteDeleteOption.ObjectType).To(BeEquivalentTo(&AzureRetentionRule_Lifecycles_DeleteAfter_ObjectType))

	// Assertions and Expectations for BackupPolicy:AzureRetentionRule:Lifecycles:SourceDataStore
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention.Lifecycles[0].SourceDataStore.DataStoreType).To(BeEquivalentTo(&AzureRetentionRule_Lifecycles_SourceDataStore_DataStoreType))
	tc.Expect(backuppolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention.Lifecycles[0].SourceDataStore.ObjectType).To(BeEquivalentTo(&AzureRetentionRule_Lifecycles_SourceDataStore_ObjectType))

	tc.Expect(backuppolicy.Status.Id).ToNot(BeNil())

	armId := *backuppolicy.Status.Id

	tc.DeleteResourceAndWait(backuppolicy)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(dataprotection.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
