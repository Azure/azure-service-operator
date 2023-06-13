/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	// . "github.com/onsi/gomega"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"

	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Dataprotection_Backuppolicy_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	region := "East Asia"

	// rg := tc.CreateTestResourceGroupAndWait()
	rg := &resources.ResourceGroup{
		ObjectMeta: tc.MakeObjectMetaWithName("t-agrawals"),
		Spec: resources.ResourceGroup_Spec{
			Location: &region,
		},
	}

	tc.CreateResourceAndWaitWithoutCleanup(rg)

	// Create a backupvault
	// region := tc.AzureRegion
	region_backupvault := "East US"
	identityType := "SystemAssigned"
	alertsForAllJobFailures_Status := dataprotection.AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled
	StorageSetting_DatastoreType_Value := dataprotection.StorageSetting_DatastoreType_VaultStore
	StorageSetting_Type_Value := dataprotection.StorageSetting_Type_LocallyRedundant

	backupvault := &dataprotection.BackupVault{
		ObjectMeta: tc.MakeObjectMetaWithName("shayvault1"),
		Spec: dataprotection.BackupVault_Spec{
			Location: &region_backupvault,
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
	// tc.CreateResourceAndWait(backupvault)
	tc.CreateResourceAndWaitWithoutCleanup(backupvault)
	// Note:
	// It is mandatory to create a backupvault before creating a backuppolicy

	// Create a BackupPolicy
	backupPolicy_ObjectType := dataprotection.BackupPolicy_ObjectType_BackupPolicy

	// consts for AzureBackupRule
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

	// consts for AzureRetentionRule
	AzureRetentionRule_Name := "Default"
	AzureRetentionRule_ObjectType := dataprotection.AzureRetentionRule_ObjectType_AzureRetentionRule
	AzureRetentionRule_IsDefault := true

	AzureRetentionRule_Lifecycles_DeleteAfter_Duration := "P9D"
	AzureRetentionRule_Lifecycles_DeleteAfter_ObjectType := dataprotection.AbsoluteDeleteOption_ObjectType_AbsoluteDeleteOption
	AzureRetentionRule_Lifecycles_SourceDataStore_DataStoreType := dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore
	AzureRetentionRule_Lifecycles_SourceDataStore_ObjectType := "DataStoreInfoBase"

	// backuppolicy generation
	backuppolicy := &dataprotection.BackupVaultsBackupPolicy{
		ObjectMeta: tc.MakeObjectMetaWithName("testsbackuppolicy"),
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

	// tc.CreateResourceAndWait(backuppolicy)
	tc.CreateResourceAndWaitWithoutCleanup(backuppolicy)

	// Asserts
}
