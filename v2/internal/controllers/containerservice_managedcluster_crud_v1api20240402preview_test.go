/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240402preview"
	"github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AKS_ManagedCluster_20240402preview_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.AzureRegion = to.Ptr("westus3") // TODO: the default test region of westus2 doesn't allow ds2_v2 at the moment

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	cluster := NewManagedCluster20240402preview(tc, rg, adminUsername, sshPublicKey)

	tc.CreateResourceAndWait(cluster)

	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	armId := *cluster.Status.Id

	// Perform a simple patch
	skuName := aks.ManagedClusterSKU_Name_Base
	skuTier := aks.ManagedClusterSKU_Tier_Standard
	old := cluster.DeepCopy()
	cluster.Spec.Sku = &aks.ManagedClusterSKU{
		Name: &skuName,
		Tier: &skuTier,
	}
	tc.PatchResourceAndWait(old, cluster)
	tc.Expect(cluster.Status.Sku).ToNot(BeNil())
	tc.Expect(cluster.Status.Sku.Name).ToNot(BeNil())
	tc.Expect(*cluster.Status.Sku.Name).To(Equal(aks.ManagedClusterSKU_Name_STATUS_Base))
	tc.Expect(cluster.Status.Sku.Tier).ToNot(BeNil())
	tc.Expect(*cluster.Status.Sku.Tier).To(Equal(aks.ManagedClusterSKU_Tier_STATUS_Standard))

	// Run sub tests
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "AKS KubeConfig secret & configmap CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_Kubeconfig_20240402preview_OperatorSpec(tc, cluster)
			},
		})
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "AKS AgentPool CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_AgentPool_20240402preview_CRUD(tc, cluster)
			},
		},
	)

	tc.DeleteResourceAndWait(cluster)

	// Ensure that the cluster was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(aks.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func NewManagedCluster20240402preview(tc *testcommon.KubePerTestContext, rg *v1api20200601.ResourceGroup, adminUsername string, sshPublicKey *string) *aks.ManagedCluster {
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
				Type: to.Ptr(aks.ManagedClusterIdentity_Type_SystemAssigned),
			},
			NetworkProfile: &aks.ContainerServiceNetworkProfile{
				NetworkPlugin: to.Ptr(aks.NetworkPlugin_Azure),
			},
			OidcIssuerProfile: &aks.ManagedClusterOIDCIssuerProfile{
				Enabled: to.Ptr(true),
			},
		},
	}

	return cluster
}

func AKS_ManagedCluster_AgentPool_20240402preview_CRUD(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	osType := aks.OSType_Linux
	agentPoolMode := aks.AgentPoolMode_System

	agentPool := &aks.ManagedClustersAgentPool{
		ObjectMeta: tc.MakeObjectMetaWithName("ap2"),
		Spec: aks.ManagedClustersAgentPool_Spec{
			Owner:  testcommon.AsOwner(cluster),
			Count:  to.Ptr(1),
			VmSize: to.Ptr("Standard_DS2_v2"),
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
	tc.Expect(agentPool.Status.Count).To(Equal(to.Ptr(1)))
	tc.Expect(agentPool.Status.OsType).ToNot(BeNil())
	tc.Expect(string(*agentPool.Status.OsType)).To(Equal(string(osType)))

	// Perform a simple patch
	old := agentPool.DeepCopy()
	agentPool.Spec.NodeLabels = map[string]string{
		"mylabel": "label",
	}
	tc.PatchResourceAndWait(old, agentPool)
	tc.Expect(agentPool.Status.NodeLabels).To(HaveKey("mylabel"))

	// Perform another patch which removes the above label. We expect it should actually be removed.
	old = agentPool.DeepCopy()
	agentPool.Spec.NodeLabels = nil
	tc.PatchResourceAndWait(old, agentPool)
	tc.Expect(agentPool.Status.NodeLabels).To(BeEmpty())
}

func AKS_ManagedCluster_Kubeconfig_20240402preview_OperatorSpec(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	old := cluster.DeepCopy()
	secret := "kubeconfig"
	cluster.Spec.OperatorSpec = &aks.ManagedClusterOperatorSpec{
		ConfigMaps: &aks.ManagedClusterOperatorConfigMaps{
			OIDCIssuerProfile: &genruntime.ConfigMapDestination{Name: "oidc", Key: "issuer"},
		},
		Secrets: &aks.ManagedClusterOperatorSecrets{
			AdminCredentials: &genruntime.SecretDestination{Name: secret, Key: "admin"},
			UserCredentials:  &genruntime.SecretDestination{Name: secret, Key: "user"},
		},
	}

	tc.PatchResourceAndWait(old, cluster)
	tc.ExpectSecretHasKeys(secret, "admin", "user")
	tc.ExpectConfigMapHasKeysAndValues("oidc", "issuer", *cluster.Status.OidcIssuerProfile.IssuerURL)
}
