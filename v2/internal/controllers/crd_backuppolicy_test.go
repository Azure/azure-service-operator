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

func Test_Dataprotection_Backuppolicy_CRUD(t *testing.T) {
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

	// Asserts
}

// Criteria: []dataprotection.BackupCriteria{
// 	{
// 		ScheduleBasedBackupCriteria: &dataprotection.ScheduleBasedBackupCriteria{
// 			ObjectType:
// 		},
// 	},
// },

// Adhoc: &dataprotection.AdhocBasedTriggerContext{
// 	ObjectType: amon4.ObjectType,
// 	TaggingCriteria: &dataprotection.AdhocBasedTaggingCriteria{
// 		TagInfo: &dataprotection.RetentionTag{
// 			TagName: amon5.TagName,
// 		},
// 	},
// },

// Global Struct Approach

// type DeleteAfter_Mystruct struct {
// 	Duration   *string `json:"duration,omitempty"`
// 	ObjectType *string `json:"objectType,omitempty"`
// }
// type DeleteAfter_Mystruct struct {
// 	Duration   *string `json:"duration,omitempty"`
// 	ObjectType *string `json:"objectType,omitempty"`
// }

// globalval1 := "P1Y"
// globalval2 := "AzureRetentionRule"

// DeleteAfter := DeleteAfter_Mystruct{
// 	Duration:   &globalval1,
// 	ObjectType: &globalval2,
// }

// Feed the above object in the above backuppolicy object
// backuppolicy.Spec.Properties.BackupPolicy.PolicyRules[0].AzureRetention.Lifecycles[0].DeleteAfter = DeleteAfter

// fmt.Println(backuppolicy)

// res2B, error := json.Marshal(backuppolicy)
// if error != nil {
// 	fmt.Println("Error")
// }
// newvar := string(res2B)

// fmt.Println("////////////////New///Line//////////////")
// fmt.Println(newvar)

// Discarded Code:
// region := tc.AzureRegion

// // This code is meant for Monitoring Settings
// vval1 := "Enabled"
// aamons := &dataprotectionstorage.AzureMonitorAlertSettings{
// 	AlertsForAllJobFailures: &vval1,
// }
// aamon := &dataprotection.AzureMonitorAlertSettings{}
// eerr := aamon.AssignProperties_From_AzureMonitorAlertSettings(aamons)
// if eerr != nil {
// 	t.Fatalf("failed to assign properties from AzureMonitorAlertSettings: %v", eerr)
// }

// // This code is meant for Storage Settings
// vval2 := "VaultStore"
// aamons2 := &dataprotectionstorage.StorageSetting{
// 	DatastoreType: &vval2,
// }
// aamon2 := &dataprotection.StorageSetting{}
// eerr2 := aamon2.AssignProperties_From_StorageSetting(aamons2)
// if eerr2 != nil {
// 	t.Fatalf("failed to assign properties from StorageSetting: %v", eerr2)
// }

// vval3 := "LocallyRedundant"
// aamons3 := &dataprotectionstorage.StorageSetting{
// 	Type: &vval3,
// }
// aamon3 := &dataprotection.StorageSetting{}
// eerr3 := aamon3.AssignProperties_From_StorageSetting(aamons3)
// if eerr3 != nil {
// 	t.Fatalf("failed to assign properties from StorageSetting: %v", eerr3)
// }

// // This code is meant for Identity
// vval4 := "SystemAssigned"
// aamons4 := &dataprotectionstorage.DppIdentityDetails{
// 	Type: &vval4,
// }
// aamon4 := &dataprotection.DppIdentityDetails{}
// eerr4 := aamon4.AssignProperties_From_DppIdentityDetails(aamons4)
// if eerr4 != nil {
// 	t.Fatalf("failed to assign properties from DppIdentityDetails: %v", eerr4)
// }

// //Create a backupvault
// backupvault := &dataprotection.BackupVault{
// 	ObjectMeta: tc.MakeObjectMetaWithName("backupvault"),
// 	Spec: dataprotection.BackupVault_Spec{
// 		Location: region,
// 		Tags:     map[string]string{"cheese": "blue"},
// 		Owner:    testcommon.AsOwner(rg),
// 		Identity: &dataprotection.DppIdentityDetails{
// 			Type: aamon4.Type,
// 		},
// 		Properties: &dataprotection.BackupVaultSpec{
// 			MonitoringSettings: &dataprotection.MonitoringSettings{
// 				AzureMonitorAlertSettings: aamon,
// 			},
// 			StorageSettings: []dataprotection.StorageSetting{
// 				{
// 					DatastoreType: aamon2.DatastoreType,
// 					Type:          aamon3.Type,
// 				},
// 			},
// 		},
// 	},
// }

// tc.CreateResourceAndWait(backupvault)

// Below code is meant for BackupPolicy

// This code is meant for Object Type
// val1 := "BackupPolicy"
// amons := &dataprotectionstorage.BackupPolicy{
// 	ObjectType: &val1,
// }
// amon := &dataprotection.BackupPolicy{}
// err := amon.AssignProperties_From_BackupPolicy(amons)
// if err != nil {
// 	t.Fatalf("failed to assign properties from BackupPolicy: %v", err)
// }

// val2 := "BackupHourly"
// val3 := "AzureBackupRule"
// amons2 := &dataprotectionstorage.AzureBackupRule{
// 	Name:       &val2,
// 	ObjectType: &val3,
// }
// amon2 := &dataprotection.AzureBackupRule{}
// err2 := amon2.AssignProperties_From_AzureBackupRule(amons2)
// if err2 != nil {
// 	t.Fatalf("failed to assign properties from AzureBackupRule: %v", err2)
// }

// val4 := "OperationalStore"
// val5 := "DataStoreInfoBase"
// amons3 := &dataprotectionstorage.DataStoreInfoBase{
// 	DataStoreType: &val4,
// 	ObjectType:    &val5,
// }
// amon3 := &dataprotection.DataStoreInfoBase{}
// err3 := amon3.AssignProperties_From_DataStoreInfoBase(amons3)
// if err3 != nil {
// 	t.Fatalf("failed to assign properties from DataStoreInfoBase: %v", err3)
// }

// val6 := "AdhocBasedTriggerContext"
// amons4 := &dataprotectionstorage.AdhocBasedTriggerContext{
// 	ObjectType: &val6,
// }
// amon4 := &dataprotection.AdhocBasedTriggerContext{}
// err4 := amon4.AssignProperties_From_AdhocBasedTriggerContext(amons4)
// if err4 != nil {
// 	t.Fatalf("failed to assign properties from AdhocBasedTriggerContext: %v", err4)
// }

// val7 := "Default"
// amons5 := &dataprotectionstorage.RetentionTag{
// 	TagName: &val7,
// }
// amon5 := &dataprotection.RetentionTag{}
// err5 := amon5.AssignProperties_From_RetentionTag(amons5)
// if err5 != nil {
// 	t.Fatalf("failed to assign properties from RetentionTag: %v", err5)
// }

// val8 := "ScheduleBasedTriggerContext"
// amons6 := &dataprotectionstorage.ScheduleBasedTriggerContext{
// 	ObjectType: &val8,
// }
// amon6 := &dataprotection.ScheduleBasedTriggerContext{}
// err6 := amon6.AssignProperties_From_ScheduleBasedTriggerContext(amons6)
// if err6 != nil {
// 	t.Fatalf("failed to assign properties from ScheduleBasedTriggerContext: %v", err6)
// }

// val9 := "UTC"
// amons7 := &dataprotectionstorage.BackupSchedule{
// 	TimeZone: &val9,
// }
// amon7 := &dataprotection.BackupSchedule{}
// err7 := amon7.AssignProperties_From_BackupSchedule(amons7)
// if err7 != nil {
// 	t.Fatalf("failed to assign properties from BackupSchedule: %v", err7)
// }

// val10 := "Incremental"
// val11 := "AzureBackupParams"
// amons8 := &dataprotectionstorage.AzureBackupParams{
// 	BackupType: &val10,
// 	ObjectType: &val11,
// }
// amon8 := &dataprotection.AzureBackupParams{}
// err8 := amon8.AssignProperties_From_AzureBackupParams(amons8)
// if err8 != nil {
// 	t.Fatalf("failed to assign properties from AzureBackupParams: %v", err8)
// }

// val12 := "Default"
// val13 := "AzureRetentionRule"
// valtrue := true
// amons9 := &dataprotectionstorage.AzureRetentionRule{
// 	Name:       &val12,
// 	ObjectType: &val13,
// 	IsDefault:  &valtrue,
// }
// amon9 := &dataprotection.AzureRetentionRule{}
// err9 := amon9.AssignProperties_From_AzureRetentionRule(amons9)
// if err9 != nil {
// 	t.Fatalf("failed to assign properties from AzureRetentionRule: %v", err9)
// }

// val14 := "P7D"
// val15 := "AbsoluteDeleteOption"
// amons10 := &dataprotectionstorage.AbsoluteDeleteOption{
// 	Duration:   &val14,
// 	ObjectType: &val15,
// }
// amon10 := &dataprotection.AbsoluteDeleteOption{}
// err10 := amon10.AssignProperties_From_AbsoluteDeleteOption(amons10)
// if err10 != nil {
// 	t.Fatalf("failed to assign properties from AbsoluteDeleteOption: %v", err10)
// }

// val16 := "OperationalStore"
// val17 := "DataStoreInfoBase"
// amons11 := &dataprotectionstorage.DataStoreInfoBase{
// 	DataStoreType: &val16,
// 	ObjectType:    &val17,
// }
// amon11 := &dataprotection.DataStoreInfoBase{}
// err11 := amon11.AssignProperties_From_DataStoreInfoBase(amons11)
// if err11 != nil {
// 	t.Fatalf("failed to assign properties from DataStoreInfoBase: %v", err11)
// }

// val18 := true
// val19 := 99
// amons12 := &dataprotectionstorage.TaggingCriteria{
// 	IsDefault:       &val18,
// 	TaggingPriority: &val19,
// }
// amon12 := &dataprotection.TaggingCriteria{}
// err12 := amon12.AssignProperties_From_TaggingCriteria(amons12)
// if err12 != nil {
// 	t.Fatalf("failed to assign properties from TaggingCriteria: %v", err12)
// }

// val20 := "Default"
// amons13 := &dataprotectionstorage.RetentionTag{
// 	TagName: &val20,
// }
// amon13 := &dataprotection.RetentionTag{}
// err13 := amon13.AssignProperties_From_RetentionTag(amons13)
// if err13 != nil {
// 	t.Fatalf("failed to assign properties from RetentionTag: %v", err13)
// }
