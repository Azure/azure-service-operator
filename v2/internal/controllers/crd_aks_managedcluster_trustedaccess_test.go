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
	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	// The testcommon package includes common testing utilities.
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// The to package includes utilities for converting values to pointers.
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ManagedCluster_TrustedAccess(t *testing.T) {

	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait() // rg.Spec.AzureName

	// Creation of Backup Vault
	backupVault := newBackupVault(tc, rg, "asotestbackupvault")

	// Creation of AKS Managed Cluster
	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	identityKind := aks.ManagedClusterIdentity_Type_SystemAssigned
	osType := aks.OSType_Linux
	agentPoolMode := aks.AgentPoolMode_System

	cluster := &aks.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: aks.ManagedCluster_Spec{
			Location:  backupVault.Spec.Location,
			Owner:     testcommon.AsOwner(rg),
			DnsPrefix: to.Ptr("aso"),
			AgentPoolProfiles: []aks.ManagedClusterAgentPoolProfile{
				{
					Name:   to.Ptr("ap1"),
					Count:  to.Ptr(1),
					VmSize: to.Ptr("Standard_DS2_v2"),
					OsType: &osType,
					Mode:   &agentPoolMode,
				},
			},
			LinuxProfile: &aks.ContainerServiceLinuxProfile{
				AdminUsername: &adminUsername,
				Ssh: &aks.ContainerServiceSshConfiguration{
					PublicKeys: []aks.ContainerServiceSshPublicKey{
						{
							KeyData: sshPublicKey,
						},
					},
				},
			},
			Identity: &aks.ManagedClusterIdentity{
				Type: &identityKind,
			},
		},
	}

	tc.CreateResourceAndWait(cluster)

	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	armId := *cluster.Status.Id

	// Creation of AKS Managed Cluster Trusted Access
	rgName := rg.Spec.AzureName
	ResourceId := "/subscriptions/f0c630e0-2995-4853-b056-0b3c09cb673f/resourceGroups/" + rgName + "/providers/Microsoft.DataProtection/BackupVaults/" + backupVault.Name

	trustedAccess := &aks.ManagedClustersTrustedAccessRoleBinding{
		ObjectMeta: tc.MakeObjectMeta("asotest"),
		Spec: aks.ManagedClusters_TrustedAccessRoleBinding_Spec{
			Owner:     testcommon.AsOwner(cluster),
			AzureName: "asotest",
			Roles:     []string{"Microsoft.DataProtection/backupVaults/backup-operator"},
			SourceResourceReference: &genruntime.ResourceReference{
				ARMID: ResourceId,
			},
		},
	}

	tc.CreateResourceAndWait(trustedAccess)

	tc.Expect(trustedAccess.Status.ProvisioningState).To(BeEquivalentTo(to.Ptr(aks.TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Succeeded)))

	// Deletion of AKS Managed Cluster
	tc.DeleteResourceAndWait(cluster)

	// Deletion of Trusted Access
	tc.DeleteResourceAndWait(trustedAccess)

	// Ensure that the cluster was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(aks.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
