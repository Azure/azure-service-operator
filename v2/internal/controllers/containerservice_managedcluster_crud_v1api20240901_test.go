/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240901"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AKS_ManagedCluster_20240901_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.AzureRegion = to.Ptr("westus3") // TODO: the default test region of westus2 doesn't allow ds2_v2 at the moment

	rg := tc.CreateTestResourceGroupAndWait()

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
			Name: "AKS KubeConfig secret CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_Kubeconfig_20240901_Secrets(tc, cluster)
			},
		})
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "AKS AgentPool CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_AgentPool_20240901_CRUD(tc, cluster)
			},
		},
		testcommon.Subtest{
			Name: "AKS MaintenanceWindow CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_MaintenanceConfiguration_20240901_CRUD(tc, cluster)
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

func AKS_ManagedCluster_AgentPool_20240901_CRUD(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	osType := aks.OSType_Linux

	agentPool := &aks.ManagedClustersAgentPool{
		ObjectMeta: tc.MakeObjectMetaWithName("ap2"),
		Spec: aks.ManagedClustersAgentPool_Spec{
			Owner:  testcommon.AsOwner(cluster),
			Count:  to.Ptr(1),
			VmSize: to.Ptr("Standard_DS2_v2"),
			OsType: &osType,
			Mode:   to.Ptr(aks.AgentPoolMode_System),
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

func AKS_ManagedCluster_Kubeconfig_20240901_Secrets(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	old := cluster.DeepCopy()
	secret := "kubeconfig"
	cluster.Spec.OperatorSpec = &aks.ManagedClusterOperatorSpec{
		Secrets: &aks.ManagedClusterOperatorSecrets{
			AdminCredentials: &genruntime.SecretDestination{Name: secret, Key: "admin"},
			UserCredentials:  &genruntime.SecretDestination{Name: secret, Key: "user"},
		},
	}

	tc.PatchResourceAndWait(old, cluster)
	tc.ExpectSecretHasKeys(secret, "admin", "user")
}

func AKS_ManagedCluster_MaintenanceConfiguration_20240901_CRUD(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	mtcConfiguration := &aks.MaintenanceConfiguration{
		ObjectMeta: tc.MakeObjectMetaWithName("aksmanagedautoupgradeschedule"), // MaintenanceWindows only support a few fixed names
		Spec: aks.MaintenanceConfiguration_Spec{
			Owner: testcommon.AsOwner(cluster),
			MaintenanceWindow: &aks.MaintenanceWindow{
				Schedule: &aks.Schedule{
					Weekly: &aks.WeeklySchedule{
						DayOfWeek:     to.Ptr(aks.WeekDay_Saturday),
						IntervalWeeks: to.Ptr(1),
					},
				},
				DurationHours: to.Ptr(12),
				StartTime:     to.Ptr("00:00"),
			},
		},
	}

	tc.CreateResourceAndWait(mtcConfiguration)
	defer tc.DeleteResourceAndWait(mtcConfiguration)

	tc.Expect(mtcConfiguration.Status.Id).ToNot(BeNil())

	// a basic assertion on a few properties
	tc.Expect(mtcConfiguration.Status.MaintenanceWindow).ToNot(BeNil())
	tc.Expect(mtcConfiguration.Status.MaintenanceWindow.DurationHours).ToNot(BeNil())
	tc.Expect(*mtcConfiguration.Status.MaintenanceWindow.DurationHours).To(Equal(12))

	// Perform a simple patch
	old := mtcConfiguration.DeepCopy()
	mtcConfiguration.Spec.MaintenanceWindow.DurationHours = to.Ptr(8)
	tc.PatchResourceAndWait(old, mtcConfiguration)
	tc.Expect(mtcConfiguration.Status.MaintenanceWindow.DurationHours).ToNot(BeNil())
	tc.Expect(*mtcConfiguration.Status.MaintenanceWindow.DurationHours).To(Equal(8))
}
