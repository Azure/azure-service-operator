/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	// . "github.com/onsi/gomega"

	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Dataprotection_Backupinstance_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a backupvault
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

	tc.CreateResourceAndWait(backuppolicy)

	// Create a BackupInstance
	
	// consts for BackupInstance
	BackupInstance_FriendlyName_Value := "test_backup_instance"
	BackupInstance_ObjectType_Value := "BackupInstance"

	// consts for BackupInstance:DataSourceInfo
	DataSourceInfo_ObjectType_Value := "Datasource"
	DataSourceInfo_ResourceType_Value := "Microsoft.ContainerService/managedClusters"
	DataSourceInfo_DatasourceType_Value := "Microsoft.ContainerService/managedClusters"

	// consts for BackupInstance:DataSourceSetInfo
	DataSourceSetInfo_ObjectType_Value := "DatasourceSet"
	DataSourceSetInfo_ResourceType_Value := "Microsoft.ContainerService/managedClusters"
	DataSourceSetInfo_DatasourceType_Value := "Microsoft.ContainerService/managedClusters"

	// consts for BackupInstance:PolicyInfo:DataStoreParameters
	DataStoreParameters_DataStoreType_Value:= dataprotection.AzureOperationalStoreParameters_DataStoreType_OperationalStore
	DataStoreParameters_ObjectType_Value := dataprotection.AzureOperationalStoreParameters_ObjectType_AzureOperationalStoreParameters

	// consts for BackupInstance:PolicyInfo:BackupDatasourceParameters
	BackupDatasourceParameters_ObjectType_Value := dataprotection.KubernetesClusterBackupDatasourceParameters_ObjectType_KubernetesClusterBackupDatasourceParameters
	BackupDatasourceParameters_SnapshotVolumes_Value := true
	BackupDatasourceParameters_IncludeClusterScopeResources_Value := true

	// backupinstance generation
	backupinstance := &dataprotection.BackupVaultsBackupInstance{
		ObjectMeta: tc.MakeObjectMetaWithName("testsbackupinstance"),
		Spec: dataprotection.BackupVaults_BackupInstance_Spec{
			Owner: testcommon.AsOwner(backupvault),
			Tags: map[string]string{"cheese": "blue"},
			Properties: &dataprotection.BackupInstance{
				FriendlyName: &BackupInstance_FriendlyName_Value,
				ObjectType: &BackupInstance_ObjectType_Value,
				DataSourceInfo: &dataprotection.Datasource{
					ObjectType: &DataSourceInfo_ObjectType_Value,
					// ResourceReference: "resource ID here  ........??",
					ResourceName: "testAKSCluster",
					ResourceType: &DataSourceInfo_ResourceType_Value,
					ResourceLocation: region, // Resource Location is same as of BackupVault
					ResourceUri: "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.ContainerService/managedClusters/testAKSCluster ......... reource URI here  ........??",
					DatasourceType: &DataSourceInfo_DatasourceType_Value,
				},
				DataSourceSetInfo: &dataprotection.DatasourceSet{
					ObjectType: &DataSourceSetInfo_ObjectType_Value,
					ResourceReference: "resource ID here  ........??",
					ResourceName: "testAKSCluster",
					ResourceType: &DataSourceSetInfo_ResourceType_Value,
					ResourceLocation: region, // Resource Location is same as of BackupVault
					ResourceUri: "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.ContainerService/managedClusters/testAKSCluster ......... reource URI here  ........??",
					DatasourceType: &DataSourceSetInfo_DatasourceType_Value,
				},
				PolicyInfo: &dataprotection.PolicyInfo{
					PolicyId: "policy ID here  ........??",
					PolicyParameters: &dataprotection.PolicyParameters{
						DataStoreParametersList: []dataprotection.DataStoreParameters{
							{
								AzureOperationalStoreParameters : &dataprotection.AzureOperationalStoreParameters{
									DataStoreType: &DataStoreParameters_DataStoreType_Value,
									ObjectType: &DataStoreParameters_ObjectType_Value,
									ResourceGroupId: "resourceGroupID",
								},
							},	
						},
						BackupDatasourceParametersList: []dataprotection.BackupDatasourceParameters{
							{
								KubernetesCluster: &dataprotection.KubernetesClusterBackupDatasourceParameters{
									ObjectType: &BackupDatasourceParameters_ObjectType_Value,
									IncludedNamespaces: []string{"default"},
									ExcludedNamespaces: []string{"kube-system"},
									IncludedResourceTypes: []string{"pods"},
									ExcludedResourceTypes: []string{"v1/Secret"},
									LabelSelectors: []string{"app=nginx"},
									SnapshotVolumes: &BackupDatasourceParameters_SnapshotVolumes_Value,
									IncludeClusterScopeResources: &BackupDatasourceParameters_IncludeClusterScopeResources_Value,
								},
							},
						},
					},
				},
				// DatasourceAuthCredentials: &dataprotection.AuthCredentials{},
				// ValidationType: &dataprotection.BackupInstance_ValidationType{},
			},
		},
	},

}