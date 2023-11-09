/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	"github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AKS_ManagedCluster_20230202Preview_CRUD(t *testing.T) {
	t.Parallel()

	// TODO: We can include this once we support AutoPurge or CreateOrRecover mode for KeyVault.
	// TODO: See https://github.com/Azure/azure-service-operator/issues/1415
	if *isLive {
		t.Skip("can't run in live mode, as this test is creates a KeyVault which reserves the name unless manually purged")
	}

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	region := to.Ptr("westus3") // TODO: the default test region of westus2 doesn't allow ds2_v2 at the moment
	//region := tc.AzureRegion

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	identityKind := aks.ManagedClusterIdentity_Type_SystemAssigned
	osType := aks.OSType_Linux
	agentPoolMode := aks.AgentPoolMode_System

	cluster := &aks.ManagedCluster{
		ObjectMeta: tc.MakeObjectMeta("mc"),
		Spec: aks.ManagedCluster_Spec{
			Location:  region,
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
			OidcIssuerProfile: &aks.ManagedClusterOIDCIssuerProfile{
				Enabled: to.Ptr(true),
			},
		},
	}

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
				AKS_ManagedCluster_Kubeconfig_20230102Preview_OperatorSpec(tc, cluster)
			},
		},
		testcommon.Subtest{
			Name: "AKS RoleBinding CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_TrustedAccessRoleBinding_20230102Preview_CRUD(tc, rg, cluster)
			},
		})
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "AKS AgentPool CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_AgentPool_20230102Preview_CRUD(tc, cluster)
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

func AKS_ManagedCluster_AgentPool_20230102Preview_CRUD(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	osType := aks.OSType_Linux
	agentPoolMode := aks.AgentPoolMode_System

	agentPool := &aks.ManagedClustersAgentPool{
		ObjectMeta: tc.MakeObjectMetaWithName("ap2"),
		Spec: aks.ManagedClusters_AgentPool_Spec{
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
}

func AKS_ManagedCluster_Kubeconfig_20230102Preview_OperatorSpec(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
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

func AKS_ManagedCluster_TrustedAccessRoleBinding_20230102Preview_CRUD(
	tc *testcommon.KubePerTestContext,
	resourceGroup *v1api20200601.ResourceGroup,
	cluster *aks.ManagedCluster,
) {

	// Create a storage account and key vault to use for the workspace
	sa := newStorageAccount(tc, resourceGroup)
	tc.CreateResourceAndWait(sa)

	kv := newVault("kv", tc, resourceGroup)
	tc.CreateResourceAndWait(kv)

	// Create workspace
	workspace := newWorkspace(
		tc,
		testcommon.AsOwner(resourceGroup),
		sa,
		kv,
		resourceGroup.Spec.Location,
	)
	tc.CreateResourceAndWait(workspace)

	roleBinding := &aks.TrustedAccessRoleBinding{
		ObjectMeta: tc.MakeObjectMetaWithName("tarb"),
		Spec: aks.ManagedClusters_TrustedAccessRoleBinding_Spec{
			Owner: testcommon.AsOwner(cluster),
			Roles: []string{
				// Microsoft.MachineLearningServices/workspaces/mlworkload
				workspace.GetType() + "/mlworkload",
			},
			SourceResourceReference: tc.MakeReferenceFromResource(workspace),
		},
	}

	// Create the role binding
	tc.CreateResourceAndWait(roleBinding)

	// Clean up when we're done
	defer tc.DeleteResourcesAndWait(roleBinding, workspace, kv, sa)

	// Perform a simple patch
	old := roleBinding.DeepCopy()
	roleBinding.Spec.Roles = []string{
		// Microsoft.MachineLearningServices/workspaces/inference-v1
		workspace.GetType() + "/inference-v1",
	}

	tc.PatchResourceAndWait(old, roleBinding)
	tc.Expect(roleBinding.Status.Roles).To(HaveLen(1))
	tc.Expect(roleBinding.Status.Roles[0]).To(Equal(roleBinding.Spec.Roles[0]))
}
