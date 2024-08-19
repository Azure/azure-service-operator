/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20200801preview"
	akscluster "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001"
	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240402preview"
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101"
	kubernetesconfiguration "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20230501"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DataProtection_BackupInstance_20231101_CRUD(t *testing.T) {
	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	contributorRoleId := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription)
	readerRoleId := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7", tc.AzureSubscription)
	saContributorRoleId := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/17d1049b-9a84-46fb-8f53-869881c3d3ab", tc.AzureSubscription)

	// create cluster
	clusterConfigMapName := "cluster-configmap"
	clusterPrincipalIdKey := "principalId"
	cluster := newBackupInstanceManagedCluster(tc, rg)
	cluster.Spec.OperatorSpec = &akscluster.ManagedClusterOperatorSpec{
		ConfigMaps: &akscluster.ManagedClusterOperatorConfigMaps{
			PrincipalId: &genruntime.ConfigMapDestination{Name: clusterConfigMapName, Key: clusterPrincipalIdKey},
		},
	}

	// create vault and policy
	backupVaultConfigMapName := "backupvault-configmap"
	backupVaultPrincipalIdKey := "principalId"
	backupVault := newBackupVault20231101(tc, rg, "asotestbackupvault")
	backupVault.Spec.OperatorSpec = &dataprotection.BackupVaultOperatorSpec{
		ConfigMaps: &dataprotection.BackupVaultOperatorConfigMaps{
			PrincipalId: &genruntime.ConfigMapDestination{Name: backupVaultConfigMapName, Key: backupVaultPrincipalIdKey},
		},
	}

	backupPolicy := newBackupPolicy20231101(tc, backupVault, "asotestbp")

	// create storage account and blob container

	acct := newStorageAccount(tc, rg)
	// Not allowed by the policy anymore. Will need to update the storage account and respective tests
	acct.Spec.AllowBlobPublicAccess = to.Ptr(false)

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
	extension := newBackupInstanceKubernetesExtension(tc, cluster, blobContainer, rg, acct)
	extension.Spec.OperatorSpec = &kubernetesconfiguration.ExtensionOperatorSpec{
		ConfigMaps: &kubernetesconfiguration.ExtensionOperatorConfigMaps{
			PrincipalId: &genruntime.ConfigMapDestination{Name: extConfigMapName, Key: extPrincipalIdKey},
		},
	}

	trustedAccessRoleBinding := newBackupInstanceTrustedAccessRoleinding(tc, cluster, backupVault)
	// give permission to extension msi over SA
	extensionRoleAssignment := newRoleAssignment(tc, acct, "extensionroleassignment", saContributorRoleId)
	extensionRoleAssignment.Spec.PrincipalIdFromConfig = &genruntime.ConfigMapReference{
		Name: extConfigMapName,
		Key:  extPrincipalIdKey,
	}

	// give read permission to vault msi over cluster
	clusterRoleAssignment := newRoleAssignment(tc, cluster, "clusterroleassignment", readerRoleId)
	clusterRoleAssignment.Spec.PrincipalIdFromConfig = &genruntime.ConfigMapReference{
		Name: backupVaultConfigMapName,
		Key:  backupVaultPrincipalIdKey,
	}

	// give cluster msi access over snapshot rg for pv creation
	clusterMSIRoleAssignment := newRoleAssignment(tc, rg, "clustermsiroleassignment", contributorRoleId)
	clusterMSIRoleAssignment.Spec.PrincipalIdFromConfig = &genruntime.ConfigMapReference{
		Name: clusterConfigMapName,
		Key:  clusterPrincipalIdKey,
	}
	// give read permission to vault msi over SRG
	snapshotRGRoleAssignment := newRoleAssignment(tc, rg, "snapshotrgroleassignment", readerRoleId)
	snapshotRGRoleAssignment.Spec.PrincipalIdFromConfig = &genruntime.ConfigMapReference{
		Name: backupVaultConfigMapName,
		Key:  backupVaultPrincipalIdKey,
	}

	// create backup instance
	biName := "asobackupinstance"
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
									DataStoreType:          to.Ptr(dataprotection.AzureOperationalStoreParameters_DataStoreType_OperationalStore),
									ResourceGroupReference: tc.MakeReferenceFromResource(rg),
									ObjectType:             to.Ptr(dataprotection.AzureOperationalStoreParameters_ObjectType_AzureOperationalStoreParameters),
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

	tc.CreateResourcesAndWait(
		cluster,
		acct,
		blobService,
		blobContainer,
		backupVault,
		backupPolicy,
		extension,
		trustedAccessRoleBinding,
		extensionRoleAssignment,
		clusterRoleAssignment,
		clusterMSIRoleAssignment,
		snapshotRGRoleAssignment,
		backupInstance)

	objectKey := client.ObjectKeyFromObject(backupInstance)

	// Assertions and Expectations
	tc.Expect(backupInstance.Status.Id).ToNot(BeNil())
	tc.Expect(backupInstance.Status.Properties.FriendlyName).To(BeEquivalentTo(to.Ptr(biName)))
	tc.Expect(backupInstance.Status.Properties.ProvisioningState).To(BeEquivalentTo(to.Ptr("Succeeded")))

	// Ensure state got updated in Azure.
	tc.Eventually(func() bool {
		var updated dataprotection.BackupVaultsBackupInstance
		tc.GetResource(objectKey, &updated)
		return checkProtectionSuccess(updated)
	}).Should(Equal(true)) // This is the expected value

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

func checkProtectionSuccess(updated dataprotection.BackupVaultsBackupInstance) bool {
	status := *updated.Status.Properties.ProtectionStatus.Status
	return status == dataprotection.ProtectionStatusDetails_Status_STATUS_ProtectionConfigured ||
		status == dataprotection.ProtectionStatusDetails_Status_STATUS_ConfiguringProtection
}

func newRoleAssignment(tc *testcommon.KubePerTestContext, owner client.Object, name, role string) *authorization.RoleAssignment {
	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMeta(name),
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(owner),
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: role,
			},
		},
	}
	return roleAssignment
}

func newBackupInstanceTrustedAccessRoleinding(tc *testcommon.KubePerTestContext, cluster *akscluster.ManagedCluster, backupVault *dataprotection.BackupVault) *aks.TrustedAccessRoleBinding {
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
	return trustedAccessRoleBinding
}

func newBackupInstanceKubernetesExtension(tc *testcommon.KubePerTestContext, cluster *akscluster.ManagedCluster, blobContainer *storage.StorageAccountsBlobServicesContainer, rg *resources.ResourceGroup, acct *storage.StorageAccount) *kubernetesconfiguration.Extension {
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
		},
	}
	return extension
}

func newBackupInstanceManagedCluster(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *akscluster.ManagedCluster {
	cluster := &akscluster.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: akscluster.ManagedCluster_Spec{
			Location:  tc.AzureRegion,
			Owner:     testcommon.AsOwner(rg),
			DnsPrefix: to.Ptr("aso"),
			AgentPoolProfiles: []akscluster.ManagedClusterAgentPoolProfile{
				{
					Name:   to.Ptr("agentpool"),
					Count:  to.Ptr(3),
					VmSize: to.Ptr("Standard_DS2_v2"),
					OsType: to.Ptr(akscluster.OSType_Linux),
					OsSKU:  to.Ptr(akscluster.OSSKU_AzureLinux),
					Mode:   to.Ptr(akscluster.AgentPoolMode_System),
				},
			},
			Identity: &akscluster.ManagedClusterIdentity{
				Type: to.Ptr(akscluster.ManagedClusterIdentity_Type_SystemAssigned),
			},
			AutoUpgradeProfile: &akscluster.ManagedClusterAutoUpgradeProfile{
				UpgradeChannel: to.Ptr(akscluster.ManagedClusterAutoUpgradeProfile_UpgradeChannel_NodeImage),
			},
			EnableRBAC: to.Ptr(true),
			AddonProfiles: map[string]akscluster.ManagedClusterAddonProfile{
				"azurepolicy": {
					Enabled: to.Ptr(true),
				},
			},
		},
	}
	return cluster
}
