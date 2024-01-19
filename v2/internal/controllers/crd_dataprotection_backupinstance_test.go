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
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku: &storage.Sku{
				Name: to.Ptr(storage.SkuName_Standard_LRS),
			},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
		},
	}
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

	tc.CreateResourceAndWait(acct, blobContainer)

	// create cluster
	cluster := &aks.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: aks.ManagedCluster_Spec{
			Location:  tc.AzureRegion,
			Owner:     testcommon.AsOwner(rg),
			DnsPrefix: to.Ptr("aso"),
			AgentPoolProfiles: []aks.ManagedClusterAgentPoolProfile{
				{
					Name:   to.Ptr("ap1"),
					Count:  to.Ptr(1),
					VmSize: to.Ptr("Standard_DS2_v2"),
					OsType: to.Ptr(aks.OSType_Linux),
					Mode:   to.Ptr(aks.AgentPoolMode_System),
				},
			},
			Identity: &aks.ManagedClusterIdentity{
				Type: to.Ptr(aks.ManagedClusterIdentity_Type_SystemAssigned),
			},
			KubernetesVersion: to.Ptr("1.27.1"),
		},
	}

	tc.CreateResourceAndWait(cluster)

	// create extension
	extension := &kubernetesconfiguration.Extension{
		ObjectMeta: tc.MakeObjectMeta("extension"),
		Spec: kubernetesconfiguration.Extension_Spec{
			ReleaseTrain:  to.Ptr("stable"),
			ExtensionType: to.Ptr("microsoft.dataprotection"),
			Owner:         tc.AsExtensionOwner(cluster),
			Scope: &kubernetesconfiguration.Scope{
				Cluster: &kubernetesconfiguration.ScopeCluster{
					ReleaseNamespace: to.Ptr("dataprotection-microsoft"),
				},
			},
			ConfigurationSettings: map[string]string{
				"configuration.backupStorageLocation.bucket":                blobContainer.AzureName(),
				"configuration.backupStorageLocation.config.resourceGroup":  tc.ResourceGroupName,
				"configuration.backupStorageLocation.config.storageAccount": acct.AzureName(),
				"configuration.backupStorageLocation.config.subscriptionId": tc.AzureSubscription,
				"credentials.tenantId":                                      tc.AzureTenant,
			},
		},
	}

	tc.CreateResourceAndWait(extension)

	// give permission to extension msi over SA
	extenstionRoleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateUUID().String()),
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
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateUUID().String()),
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
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateUUID().String()),
		Spec: authorization.RoleAssignment_Spec{
			Owner:       tc.AsExtensionOwner(cluster),
			PrincipalId: backupVault.Status.Identity.PrincipalId,
			RoleDefinitionReference: &genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7", tc.AzureSubscription), // This is Reader Role
			},
		},
	}

	tc.CreateResourceAndWait(clusterRoleAssignment)

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

	// Assertions and Expectations

	// Note:
	// Patch Operations are currently not allowed on BackupInstance

	// Delete the backupinstance

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(dataprotection.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
