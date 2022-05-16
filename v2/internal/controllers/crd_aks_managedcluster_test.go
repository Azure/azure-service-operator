/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_AKS_ManagedCluster_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	identityKind := aks.ManagedClusterIdentityTypeSystemAssigned
	osType := aks.ManagedClusterAgentPoolProfileOsTypeLinux
	agentPoolMode := aks.ManagedClusterAgentPoolProfileModeSystem

	cluster := &aks.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: aks.ManagedClusters_Spec{
			Location:  tc.AzureRegion,
			Owner:     testcommon.AsOwner(rg),
			DnsPrefix: to.StringPtr("aso"),
			AgentPoolProfiles: []aks.ManagedClusterAgentPoolProfile{
				{
					Name:   to.StringPtr("ap1"),
					Count:  to.IntPtr(1),
					VmSize: to.StringPtr("Standard_DS2_v2"),
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

	// Perform a simple patch
	skuName := aks.ManagedClusterSKUNameBasic
	skuTier := aks.ManagedClusterSKUTierPaid
	old := cluster.DeepCopy()
	cluster.Spec.Sku = &aks.ManagedClusterSKU{
		Name: &skuName,
		Tier: &skuTier,
	}
	tc.PatchResourceAndWait(old, cluster)
	tc.Expect(cluster.Status.Sku).ToNot(BeNil())
	tc.Expect(cluster.Status.Sku.Name).ToNot(BeNil())
	tc.Expect(*cluster.Status.Sku.Name).To(Equal(aks.ManagedClusterSKUStatusNameBasic))
	tc.Expect(cluster.Status.Sku.Tier).ToNot(BeNil())
	tc.Expect(*cluster.Status.Sku.Tier).To(Equal(aks.ManagedClusterSKUStatusTierPaid))

	// Run sub tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "AKS AgentPool CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_AgentPool_CRUD(testContext, cluster)
			},
		},
	)

	tc.DeleteResourceAndWait(cluster)

	// Ensure that the cluster was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(aks.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func AKS_ManagedCluster_AgentPool_CRUD(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	osType := aks.ManagedClusterAgentPoolProfilePropertiesOsTypeLinux
	agentPoolMode := aks.ManagedClusterAgentPoolProfilePropertiesModeSystem

	agentPool := &aks.ManagedClustersAgentPool{
		ObjectMeta: tc.MakeObjectMetaWithName("ap2"),
		Spec: aks.ManagedClustersAgentPools_Spec{
			Owner:  testcommon.AsOwner(cluster),
			Count:  to.IntPtr(1),
			VmSize: to.StringPtr("Standard_DS2_v2"),
			OsType: &osType,
			Mode:   &agentPoolMode,
		},
	}

	tc.CreateResourceAndWait(agentPool)
	// Note: If this test is just too slow to record (it's pretty slow), we can remove this
	// delete entirely which will speed it up some. Leaving it here for now though as we
	// need it to test the DELETE path of agent pool. If this is removed, the agent pool
	// will still be cleaned up by the cluster delete that happens later.
	defer tc.DeleteResourceAndWait(agentPool)

	tc.Expect(agentPool.Status.Id).ToNot(BeNil())

	// a basic assertion on a few properties
	tc.Expect(agentPool.Status.Count).To(Equal(to.IntPtr(1)))
	tc.Expect(agentPool.Status.OsType).ToNot(BeNil())
	tc.Expect(string(*agentPool.Status.OsType)).To(Equal(string(osType)))

	// Perform a simple patch
	old := agentPool.DeepCopy()
	agentPool.Spec.NodeLabels = map[string]string{
		"mylabel": "label",
	}
	tc.PatchResourceAndWait(old, agentPool)
	tc.Expect(agentPool.Status.NodeLabels).To(HaveKey("mylabel"))
}
