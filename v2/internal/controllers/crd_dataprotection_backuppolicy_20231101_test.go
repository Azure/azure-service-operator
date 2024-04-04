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
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101"
	// The testcommon package includes common testing utilities.
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// The to package includes utilities for converting values to pointers.
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func newBackupPolicy20231101(tc *testcommon.KubePerTestContext, backupVault *dataprotection.BackupVault, name string) *dataprotection.BackupVaultsBackupPolicy {
	backupPolicy := &dataprotection.BackupVaultsBackupPolicy{
		ObjectMeta: tc.MakeObjectMeta(name),
		Spec: dataprotection.BackupVaults_BackupPolicy_Spec{
			Owner: testcommon.AsOwner(backupVault),
			Properties: &dataprotection.BaseBackupPolicy{
				BackupPolicy: &dataprotection.BackupPolicy{
					DatasourceTypes: []string{"Microsoft.ContainerService/managedClusters"},
					ObjectType:      to.Ptr(dataprotection.BackupPolicy_ObjectType_BackupPolicy),
					PolicyRules: []dataprotection.BasePolicyRule{
						{
							AzureBackup: createAzureBackupRule20231101(),
						},
						{
							AzureRetention: createAzureRetentionRule20231101(),
						},
					},
				},
			},
		},
	}
	return backupPolicy
}

func Test_Dataprotection_Backuppolicy_20231101_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a new backupvault resource
	backupVault := newBackupVault20231101(tc, rg, "asotestbackupvault")

	// Note:
	// It is mandatory to create a backupvault before creating a backuppolicy

	// Create a BackupPolicy
	backupPolicy := newBackupPolicy20231101(tc, backupVault, "asotestbackuppolicy")

	// Sequence of creating BackupVault and BackupPolicy is handled by ASO internally
	tc.CreateResourcesAndWait(backupVault, backupPolicy)

	// Assertions and Expectations
	tc.Expect(backupPolicy.Status.Properties.BackupPolicy.DatasourceTypes).To(BeEquivalentTo([]string{"Microsoft.ContainerService/managedClusters"}))
	tc.Expect(backupPolicy.Status.Properties.BackupPolicy.ObjectType).To(BeEquivalentTo(to.Ptr(dataprotection.BackupPolicy_ObjectType_BackupPolicy)))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule
	policyRule0 := backupPolicy.Status.Properties.BackupPolicy.PolicyRules[0].AzureBackup

	tc.Expect(policyRule0.Name).To(BeEquivalentTo(to.Ptr("BackupHourly")))
	tc.Expect(policyRule0.ObjectType).To(BeEquivalentTo(to.Ptr(dataprotection.AzureBackupRule_ObjectType_AzureBackupRule)))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule:BackupParameters
	tc.Expect(policyRule0.BackupParameters.AzureBackupParams.BackupType).To(BeEquivalentTo(to.Ptr("Incremental")))
	tc.Expect(policyRule0.BackupParameters.AzureBackupParams.ObjectType).To(BeEquivalentTo(to.Ptr(dataprotection.AzureBackupParams_ObjectType_AzureBackupParams)))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule:DataStore
	tc.Expect(policyRule0.DataStore.DataStoreType).To(BeEquivalentTo(to.Ptr(dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore)))
	tc.Expect(policyRule0.DataStore.ObjectType).To(BeEquivalentTo(to.Ptr("DataStoreInfoBase")))

	// Assertions and Expectations for BackupPolicy:AzureBackupRule:Trigger
	tc.Expect(policyRule0.Trigger.Schedule.ObjectType).To(BeEquivalentTo(to.Ptr(dataprotection.ScheduleBasedTriggerContext_ObjectType_ScheduleBasedTriggerContext)))

	// Assertions and Expectations for BackupPolicy:AzureRetentionRule
	policyRule1 := backupPolicy.Status.Properties.BackupPolicy.PolicyRules[1].AzureRetention

	tc.Expect(policyRule1.Name).To(BeEquivalentTo(to.Ptr("Default")))
	tc.Expect(policyRule1.ObjectType).To(BeEquivalentTo(to.Ptr(dataprotection.AzureRetentionRule_ObjectType_AzureRetentionRule)))
	tc.Expect(policyRule1.IsDefault).To(BeEquivalentTo(to.Ptr(true)))

	// Assertions and Expectations for BackupPolicy:AzureRetentionRule:Lifecycles
	tc.Expect(policyRule1.Lifecycles[0].DeleteAfter.AbsoluteDeleteOption.Duration).To(BeEquivalentTo(to.Ptr("P9D")))
	tc.Expect(policyRule1.Lifecycles[0].DeleteAfter.AbsoluteDeleteOption.ObjectType).To(BeEquivalentTo(to.Ptr(dataprotection.AbsoluteDeleteOption_ObjectType_AbsoluteDeleteOption)))

	// Assertions and Expectations for BackupPolicy:AzureRetentionRule:Lifecycles:SourceDataStore
	tc.Expect(policyRule1.Lifecycles[0].SourceDataStore.DataStoreType).To(BeEquivalentTo(to.Ptr(dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore)))
	tc.Expect(policyRule1.Lifecycles[0].SourceDataStore.ObjectType).To(BeEquivalentTo(to.Ptr("DataStoreInfoBase")))

	tc.Expect(backupPolicy.Status.Id).ToNot(BeNil())

	armId := *backupPolicy.Status.Id

	// Note:
	// Patch Operations are currently not allowed on BackupPolicy

	// Delete the backuppolicy
	tc.DeleteResourceAndWait(backupPolicy)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(dataprotection.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

// creating a new backup policy rule: AZURE_BACKUP_RULE
func createAzureBackupRule20231101() *dataprotection.AzureBackupRule {
	azureBackupRule := &dataprotection.AzureBackupRule{
		Name:       to.Ptr("BackupHourly"),
		ObjectType: to.Ptr(dataprotection.AzureBackupRule_ObjectType_AzureBackupRule),
		BackupParameters: &dataprotection.BackupParameters{
			AzureBackupParams: &dataprotection.AzureBackupParams{
				BackupType: to.Ptr("Incremental"),
				ObjectType: to.Ptr(dataprotection.AzureBackupParams_ObjectType_AzureBackupParams),
			},
		},
		DataStore: &dataprotection.DataStoreInfoBase{
			DataStoreType: to.Ptr(dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore),
			ObjectType:    to.Ptr("DataStoreInfoBase"),
		},
		Trigger: &dataprotection.TriggerContext{
			Schedule: &dataprotection.ScheduleBasedTriggerContext{
				ObjectType: to.Ptr(dataprotection.ScheduleBasedTriggerContext_ObjectType_ScheduleBasedTriggerContext),
				Schedule: &dataprotection.BackupSchedule{
					RepeatingTimeIntervals: []string{"R/2023-06-07T10:26:32+00:00/PT4H"},
					TimeZone:               to.Ptr("UTC"),
				},
				TaggingCriteria: []dataprotection.TaggingCriteria{
					{
						IsDefault:       to.Ptr(true),
						TaggingPriority: to.Ptr(99),
						TagInfo: &dataprotection.RetentionTag{
							TagName: to.Ptr("Default"),
						},
					},
				},
			},
		},
	}

	return azureBackupRule
}

// creating a new retention policy rule: AZURE_RETENTION_RULE
func createAzureRetentionRule20231101() *dataprotection.AzureRetentionRule {
	azureRetentionRule := &dataprotection.AzureRetentionRule{
		Name:       to.Ptr("Default"),
		ObjectType: to.Ptr(dataprotection.AzureRetentionRule_ObjectType_AzureRetentionRule),
		IsDefault:  to.Ptr(true),
		Lifecycles: []dataprotection.SourceLifeCycle{
			{
				DeleteAfter: &dataprotection.DeleteOption{
					AbsoluteDeleteOption: &dataprotection.AbsoluteDeleteOption{
						Duration:   to.Ptr("P9D"),
						ObjectType: to.Ptr(dataprotection.AbsoluteDeleteOption_ObjectType_AbsoluteDeleteOption),
					},
				},
				SourceDataStore: &dataprotection.DataStoreInfoBase{
					DataStoreType: to.Ptr(dataprotection.DataStoreInfoBase_DataStoreType_OperationalStore),
					ObjectType:    to.Ptr("DataStoreInfoBase"),
				},
				TargetDataStoreCopySettings: []dataprotection.TargetCopySetting{},
			},
		},
	}

	return azureRetentionRule
}
