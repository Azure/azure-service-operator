/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	// The testing package is imported for testing-related functionality.
	"fmt"
	"testing"

	// The gomega package is used for assertions and expectations in tests.
	. "github.com/onsi/gomega"

	// The dataprotection package contains types and functions related to dataprotection resources.
	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	akscluster "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001"
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	kubernetesconfiguration "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20230501"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	// The testcommon package includes common testing utilities.
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// The to package includes utilities for converting values to pointers.
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Dataprotection_Backupinstace_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// create storage account and blob container
	acct := newStorageAccount(tc, rg)
	tc.CreateResourceAndWait(acct)

	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccounts_BlobService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	tc.CreateResourceAndWait(blobService)

	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMeta("velero"),
		Spec: storage.StorageAccounts_BlobServices_Container_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	tc.CreateResourceAndWait(blobContainer)

	// create cluster
	osType := akscluster.OSType_Linux
	osSKU := akscluster.OSSKU_AzureLinux
	upgradeChannel := akscluster.ManagedClusterAutoUpgradeProfile_UpgradeChannel_NodeImage
	agentPoolMode := akscluster.AgentPoolMode_System
	cluster := &akscluster.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: akscluster.ManagedCluster_Spec{
			KubernetesVersion: to.Ptr("1.27.3"),
			Location:          tc.AzureRegion,
			Owner:             testcommon.AsOwner(rg),
			DnsPrefix:         to.Ptr("aso"),
			AgentPoolProfiles: []akscluster.ManagedClusterAgentPoolProfile{
				{
					Name:   to.Ptr("agentpool"),
					Count:  to.Ptr(3),
					VmSize: to.Ptr("Standard_B4ms"),
					OsType: &osType,
					OsSKU:  &osSKU,
					Mode:   &agentPoolMode,
				},
			},
			Identity: &akscluster.ManagedClusterIdentity{
				Type: to.Ptr(akscluster.ManagedClusterIdentity_Type_SystemAssigned),
			},
			AutoUpgradeProfile: &akscluster.ManagedClusterAutoUpgradeProfile{
				UpgradeChannel: &upgradeChannel,
			},
			EnableRBAC: to.Ptr(true),
		},
	}
	tc.CreateResourceAndWait(cluster)

	// create extension
	extension := &kubernetesconfiguration.Extension{
		ObjectMeta: tc.MakeObjectMeta("extension"),
		Spec: kubernetesconfiguration.Extension_Spec{
			ReleaseTrain:  to.Ptr("stable"),
			ExtensionType: to.Ptr("microsoft.dataprotection.kubernetes"),
			Owner:         tc.AsExtensionOwner(cluster),
			Scope: &kubernetesconfiguration.Scope{
				Cluster: &kubernetesconfiguration.ScopeCluster{
					ReleaseNamespace: to.Ptr("dataprotection-microsoft"),
				},
			},
			ConfigurationSettings: map[string]string{
				"configuration.backupStorageLocation.bucket":                blobContainer.AzureName(),
				"configuration.backupStorageLocation.config.resourceGroup":  rg.AzureName(),
				"configuration.backupStorageLocation.config.storageAccount": acct.AzureName(),
				"configuration.backupStorageLocation.config.subscriptionId": tc.AzureSubscription,
				"credentials.tenantId":                                      tc.AzureTenant,
			},
		},
	}

	tc.CreateResourceAndWait(extension)

	tc.Expect(extension.Status.Id).ToNot(BeNil())
	tc.Expect(extension.Status.Identity.PrincipalId).ToNot(BeNil())

	// give permission to extension msi over SA
	extenstionRoleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMeta("extenstionRoleAssignment"),
		Spec: authorization.RoleAssignment_Spec{
			Owner:       tc.AsExtensionOwner(acct),
			PrincipalId: extension.Status.Identity.PrincipalId,
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/17d1049b-9a84-46fb-8f53-869881c3d3ab", tc.AzureSubscription), // This is Storage Account Contributor Role
			},
		},
	}

	tc.CreateResourceAndWait(extenstionRoleAssignment)

	// create vault and policy
	backupVault := newBackupVault(tc, rg, "asotestbackupvault")
	backupPolicy := newBackupPolicy(tc, backupVault, "asotestbackuppolicy")
	tc.CreateResourcesAndWait(backupVault, backupPolicy)

	// give read permission to vault msi over SRG
	snapshotRGRoleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMeta("snapshotRGRoleAssignment"),
		Spec: authorization.RoleAssignment_Spec{
			Owner:       tc.AsExtensionOwner(rg),
			PrincipalId: backupVault.Status.Identity.PrincipalId,
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7", tc.AzureSubscription), // This is Reader Role
			},
		},
	}

	tc.CreateResourceAndWait(snapshotRGRoleAssignment)

	// give read permission to vault msi over cluster
	clusterRoleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMeta("clusterRoleAssignment"),
		Spec: authorization.RoleAssignment_Spec{
			Owner:       tc.AsExtensionOwner(cluster),
			PrincipalId: backupVault.Status.Identity.PrincipalId,
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7", tc.AzureSubscription), // This is Reader Role
			},
		},
	}

	tc.CreateResourceAndWait(clusterRoleAssignment)

	// give cluster msi access over snapshot rg for pv creation
	clusterMSIRoleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMeta("clusterMSIRoleAssignment"),
		Spec: authorization.RoleAssignment_Spec{
			Owner:       tc.AsExtensionOwner(rg),
			PrincipalId: cluster.Status.Identity.PrincipalId,
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is Contributor Role
			},
		},
	}

	tc.CreateResourceAndWait(clusterMSIRoleAssignment)

	// create TA role binding
	trustedAccessRoleBinding := &aks.TrustedAccessRoleBinding{
		ObjectMeta: tc.MakeObjectMetaWithName("tarb"),
		Spec: aks.ManagedClusters_TrustedAccessRoleBinding_Spec{
			Owner: testcommon.AsOwner(cluster),
			Roles: []string{
				// Microsoft.DataProtection/backupVaults/backup-operator
				backupVault.GetType() + "/backup-operator",
			},
			SourceResourceReference: tc.MakeReferenceFromResource(backupVault),
		},
	}

	tc.CreateResourceAndWait(trustedAccessRoleBinding)

	//create backup instance
	biName := "asotestbackupinstance"
	backupInstance := &dataprotection.BackupVaultsBackupInstance{
		ObjectMeta: tc.MakeObjectMeta(biName),
		Spec: dataprotection.BackupVaults_BackupInstance_Spec{
			Owner: testcommon.AsOwner(backupVault),
			Properties: &dataprotection.BackupInstance{
				FriendlyName: to.Ptr(biName),
				DataSourceInfo: &dataprotection.Datasource{
					DatasourceType:   to.Ptr(cluster.GetType()),
					ResourceUri:      cluster.Status.Id,
					ResourceName:     to.Ptr(cluster.AzureName()),
					ResourceLocation: cluster.Spec.Location,
				},
				DataSourceSetInfo: &dataprotection.DatasourceSet{
					DatasourceType:   to.Ptr(cluster.GetType()),
					ResourceUri:      cluster.Status.Id,
					ResourceName:     to.Ptr(cluster.AzureName()),
					ResourceLocation: cluster.Spec.Location,
				},
				PolicyInfo: &dataprotection.PolicyInfo{
					PolicyId: backupPolicy.Status.Id,
					PolicyParameters: &dataprotection.PolicyParameters{
						DataStoreParametersList: []dataprotection.DataStoreParameters{
							{
								AzureOperationalStoreParameters: &dataprotection.AzureOperationalStoreParameters{
									DataStoreType:   to.Ptr(dataprotection.AzureOperationalStoreParameters_DataStoreType_OperationalStore),
									ResourceGroupId: rg.Status.Id,
									ObjectType:      to.Ptr(dataprotection.AzureOperationalStoreParameters_ObjectType_AzureOperationalStoreParameters),
								},
							},
						},
						BackupDatasourceParametersList: []dataprotection.BackupDatasourceParameters{
							{
								KubernetesCluster: &dataprotection.KubernetesClusterBackupDatasourceParameters{
									SnapshotVolumes:              to.Ptr(true),
									IncludeClusterScopeResources: to.Ptr(true),
								},
							},
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(backupInstance)

	// Assertions and Expectations
	armId := *backupInstance.Status.Id
	tc.Expect(backupInstance.Status.Id).ToNot(BeNil())
	tc.Expect(backupInstance.Status.Properties.FriendlyName).To(Equal((to.Ptr(biName))))
	tc.Expect(backupInstance.Status.Properties.DataSourceInfo.ResourceID).To(Equal(cluster.Status.Id))
	tc.Expect(backupInstance.Status.Properties.DataSourceSetInfo.ResourceID).To(Equal(cluster.Status.Id))
	tc.Expect(backupInstance.Status.Properties.ProvisioningState).To(Equal(to.Ptr("Succeeded")))

	// Note:
	// Patch Operations are currently not allowed on BackupInstance currently

	// Delete the backupinstance
	tc.DeleteResourceAndWait(backupInstance)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(dataprotection.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
