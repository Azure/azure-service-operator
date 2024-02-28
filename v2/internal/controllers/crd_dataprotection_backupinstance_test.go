/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"

	akscluster "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001"
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101"
	/*
		"fmt"
		authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
		aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
		kubernetesconfiguration "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20230501"
		storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	*/)

func Test_Dataprotection_Backupinstace_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// create cluster
	clusterConfigMapName := "cluster-configmap"
	clusterPrincipalIdKey := "principalId"
	osType := akscluster.OSType_Linux
	osSKU := akscluster.OSSKU_AzureLinux
	upgradeChannel := akscluster.ManagedClusterAutoUpgradeProfile_UpgradeChannel_NodeImage
	agentPoolMode := akscluster.AgentPoolMode_System
	cluster := &akscluster.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: akscluster.ManagedCluster_Spec{
			Location:  tc.AzureRegion,
			Owner:     testcommon.AsOwner(rg),
			DnsPrefix: to.Ptr("aso"),
			AgentPoolProfiles: []akscluster.ManagedClusterAgentPoolProfile{
				{
					Name:   to.Ptr("agentpool"),
					Count:  to.Ptr(2),
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
			AddonProfiles: map[string]akscluster.ManagedClusterAddonProfile{
				"azurepolicy": {
					Enabled: to.Ptr(true),
				},
			},
			OperatorSpec: &akscluster.ManagedClusterOperatorSpec{
				ConfigMaps: &akscluster.ManagedClusterOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{Name: clusterConfigMapName, Key: clusterPrincipalIdKey},
				},
			},
		},
	}

	// create vault and policy
	backupVaultConfigMapName := "backupvault-configmap"
	backupVaultPrincipalIdKey := "principalId"
	backupVault := newBackupVault(tc, rg, "asotestbackupvault")
	backupVault.Spec.OperatorSpec = &dataprotection.BackupVaultOperatorSpec{
		ConfigMaps: &dataprotection.BackupVaultOperatorConfigMaps{
			PrincipalId: &genruntime.ConfigMapDestination{Name: backupVaultConfigMapName, Key: backupVaultPrincipalIdKey},
		},
	}

	backupPolicy := newBackupPolicy(tc, backupVault, "asotestbackuppolicy")

	// create storage account and blob container
	/*
		acct := newStorageAccount(tc, rg)

		blobService := &storage.StorageAccountsBlobService{
			ObjectMeta: tc.MakeObjectMeta("blobservice"),
			Spec: storage.StorageAccounts_BlobService_Spec{
				Owner: testcommon.AsOwner(acct),
			},
		}

		blobContainer := &storage.StorageAccountsBlobServicesContainer{
			ObjectMeta: tc.MakeObjectMeta("velero"),
			Spec: storage.StorageAccounts_BlobServices_Container_Spec{
				Owner: testcommon.AsOwner(blobService),
			},
		}

		// create extension
		extConfigMapName := "ext-configmap"
		extPrincipalIdKey := "principalId"
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
					"configuration.backupStorageLocation.bucket":                blobContainer.Name,
					"configuration.backupStorageLocation.config.resourceGroup":  rg.Name,
					"configuration.backupStorageLocation.config.storageAccount": acct.Name,
					"configuration.backupStorageLocation.config.subscriptionId": tc.AzureSubscription,
					"credentials.tenantId":                                      tc.AzureTenant,
				},
				OperatorSpec: &kubernetesconfiguration.ExtensionOperatorSpec{
					ConfigMaps: &kubernetesconfiguration.ExtensionOperatorConfigMaps{
						PrincipalId: &genruntime.ConfigMapDestination{Name: extConfigMapName, Key: extPrincipalIdKey},
					},
				},
			},
		}

		// create TA role binding
		trustedAccessRoleBinding := &aks.TrustedAccessRoleBinding{
			ObjectMeta: tc.MakeObjectMetaWithName("tarb"),
			Spec: aks.ManagedClusters_TrustedAccessRoleBinding_Spec{
				Owner: testcommon.AsOwner(cluster),
				Roles: []string{
					"Microsoft.DataProtection/backupVaults/backup-operator",
				},
				SourceResourceReference: tc.MakeReferenceFromResource(backupVault),
			},
		}

		// give permission to extension msi over SA
		extenstionRoleAssignmentGUID, err := tc.Namer.GenerateUUID()
		tc.Expect(err).ToNot(HaveOccurred())
		extenstionRoleAssignment := &authorization.RoleAssignment{
			ObjectMeta: tc.MakeObjectMeta(extenstionRoleAssignmentGUID.String()),
			Spec: authorization.RoleAssignment_Spec{
				Owner: tc.AsExtensionOwner(acct),
				PrincipalIdFromConfig: &genruntime.ConfigMapReference{
					Name: extConfigMapName,
					Key:  extPrincipalIdKey,
				},
				RoleDefinitionReference: &genruntime.ResourceReference{
					ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/17d1049b-9a84-46fb-8f53-869881c3d3ab", tc.AzureSubscription), // This is Storage Account Contributor Role
				},
			},
		}

		// give read permission to vault msi over cluster
		clusterRoleAssignmentGUID, err := tc.Namer.GenerateUUID()
		tc.Expect(err).ToNot(HaveOccurred())
		clusterRoleAssignment := &authorization.RoleAssignment{
			ObjectMeta: tc.MakeObjectMeta(clusterRoleAssignmentGUID.String()),
			Spec: authorization.RoleAssignment_Spec{
				Owner: tc.AsExtensionOwner(cluster),
				PrincipalIdFromConfig: &genruntime.ConfigMapReference{
					Name: backupVaultConfigMapName,
					Key:  backupVaultPrincipalIdKey,
				},
				RoleDefinitionReference: &genruntime.ResourceReference{
					ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7", tc.AzureSubscription), // This is Reader Role
				},
			},
		}

		// give cluster msi access over snapshot rg for pv creation
		clusterMSIRoleAssignmentAssignmentGUID, err := tc.Namer.GenerateUUID()
		tc.Expect(err).ToNot(HaveOccurred())
		clusterMSIRoleAssignment := &authorization.RoleAssignment{
			ObjectMeta: tc.MakeObjectMeta(clusterMSIRoleAssignmentAssignmentGUID.String()),
			Spec: authorization.RoleAssignment_Spec{
				Owner: tc.AsExtensionOwner(rg),
				PrincipalIdFromConfig: &genruntime.ConfigMapReference{
					Name: clusterConfigMapName,
					Key:  clusterPrincipalIdKey,
				},
				RoleDefinitionReference: &genruntime.ResourceReference{
					ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is Contributor Role
				},
			},
		}


		// give read permission to vault msi over SRG
		snapshotRGRoleAssignmentGUID, err := tc.Namer.GenerateUUID()
		tc.Expect(err).ToNot(HaveOccurred())
		snapshotRGRoleAssignment := &authorization.RoleAssignment{
			ObjectMeta: tc.MakeObjectMeta(snapshotRGRoleAssignmentGUID.String()),
			Spec: authorization.RoleAssignment_Spec{
				Owner: tc.AsExtensionOwner(rg),
				PrincipalIdFromConfig: &genruntime.ConfigMapReference{
					Name: backupVaultConfigMapName,
					Key:  backupVaultPrincipalIdKey,
				},
				RoleDefinitionReference: &genruntime.ResourceReference{
					ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7", tc.AzureSubscription), // This is Reader Role
				},
			},
		}
	*/

	//create backup instance
	biName := "asotestbackupinstance"
	backupInstance := &dataprotection.BackupVaultsBackupInstance{
		ObjectMeta: tc.MakeObjectMeta(biName),
		Spec: dataprotection.BackupVaults_BackupInstance_Spec{
			Owner: testcommon.AsOwner(backupVault),
			Properties: &dataprotection.BackupInstance{
				ObjectType:   to.Ptr("BackupInstance"),
				FriendlyName: to.Ptr(biName),
				DataSourceInfo: &dataprotection.Datasource{
					ObjectType:        to.Ptr("Datasource"),
					DatasourceType:    to.Ptr(cluster.GetType()),
					ResourceUri:       cluster.Status.Id,
					ResourceName:      to.Ptr(cluster.AzureName()),
					ResourceLocation:  cluster.Spec.Location,
					ResourceType:      to.Ptr(cluster.GetType()),
					ResourceReference: tc.MakeReferenceFromResource(cluster),
				},
				DataSourceSetInfo: &dataprotection.DatasourceSet{
					ObjectType:        to.Ptr("DatasourceSet"),
					DatasourceType:    to.Ptr(cluster.GetType()),
					ResourceUri:       cluster.Status.Id,
					ResourceName:      to.Ptr(cluster.AzureName()),
					ResourceType:      to.Ptr(cluster.GetType()),
					ResourceLocation:  cluster.Spec.Location,
					ResourceReference: tc.MakeReferenceFromResource(cluster),
				},
				PolicyInfo: &dataprotection.PolicyInfo{
					PolicyReference: tc.MakeReferenceFromResource(backupPolicy),
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
									ObjectType:                   to.Ptr(dataprotection.KubernetesClusterBackupDatasourceParameters_ObjectType_KubernetesClusterBackupDatasourceParameters),
								},
							},
						},
					},
				},
			},
		},
	}

	tc.CreateResourcesAndWait(cluster, backupVault, backupPolicy, backupInstance)

	// tc.CreateResourcesAndWait(cluster, acct, blobService, blobContainer, backupVault, backupPolicy, extension,
	// 	trustedAccessRoleBinding, extenstionRoleAssignment, clusterRoleAssignment, clusterMSIRoleAssignment, snapshotRGRoleAssignment, backupInstance)

	objectKey := client.ObjectKeyFromObject(backupInstance)

	// Assertions and Expectations
	tc.Expect(backupInstance.Status.Id).ToNot(BeNil())
	tc.Expect(backupInstance.Status.Properties.FriendlyName).To(BeEquivalentTo(to.Ptr(biName)))
	tc.Expect(backupInstance.Status.Properties.ProvisioningState).To(BeEquivalentTo(to.Ptr("Succeeded")))

	// Ensure state got updated in Azure.
	tc.Eventually(func() *dataprotection.ProtectionStatusDetails_Status_STATUS {
		var updated dataprotection.BackupVaultsBackupInstance
		tc.GetResource(objectKey, &updated)
		return updated.Status.Properties.ProtectionStatus.Status
	}).Should(Equal(to.Ptr(dataprotection.ProtectionStatusDetails_Status_STATUS_ProtectionConfigured))) // This is the expected value

	// Note:
	// Patch Operations are currently not allowed on BackupInstance currently

	// Delete the backupinstance
	tc.DeleteResourceAndWait(backupInstance)

	// Ensure that the resource was really deleted in Azure
	armId := *backupInstance.Status.Id
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(dataprotection.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
