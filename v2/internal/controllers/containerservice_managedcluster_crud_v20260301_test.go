/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v20260301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_AKS_ManagedCluster_20260301_CRUD(t *testing.T) {
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
					OsType: to.Ptr(aks.ManagedClusterAgentPoolProfile_OsType_Linux),
					OsSKU:  to.Ptr(aks.OSSKU_AzureContainerLinux),
					Mode:   to.Ptr(aks.AgentPoolMode_System),
				},
			},
			Identity: &aks.ManagedClusterIdentity{
				Type: to.Ptr(aks.ResourceIdentityType_SystemAssigned),
			},
		},
	}

	tc.CreateResourceAndWait(cluster)

	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	armId := *cluster.Status.Id
	tc.Expect(cluster.Status.AgentPoolProfiles).To(HaveLen(1))
	tc.Expect(cluster.Status.AgentPoolProfiles[0].OsSKU).ToNot(BeNil())
	tc.Expect(*cluster.Status.AgentPoolProfiles[0].OsSKU).To(Equal(aks.OSSKU_STATUS_AzureContainerLinux))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "AKS AgentPool CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_ManagedCluster_AgentPool_20260301_CRUD(tc, cluster)
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

func AKS_ManagedCluster_AgentPool_20260301_CRUD(tc *testcommon.KubePerTestContext, cluster *aks.ManagedCluster) {
	osType := aks.ManagedClusterAgentPoolProfileProperties_OsType_Linux
	osSKU := aks.OSSKU_AzureContainerLinux

	agentPool := &aks.ManagedClustersAgentPool{
		ObjectMeta: tc.MakeObjectMetaWithName("ap2"),
		Spec: aks.ManagedClustersAgentPool_Spec{
			Owner:  testcommon.AsOwner(cluster),
			Count:  to.Ptr(1),
			VmSize: to.Ptr("Standard_DS2_v2"),
			OsType: &osType,
			OsSKU:  &osSKU,
			Mode:   to.Ptr(aks.AgentPoolMode_User),
		},
	}

	tc.CreateResourceAndWait(agentPool)
	defer tc.DeleteResourceAndWait(agentPool)

	tc.Expect(agentPool.Status.Id).ToNot(BeNil())
	tc.Expect(agentPool.Status.Count).To(Equal(to.Ptr(1)))
	tc.Expect(agentPool.Status.OsType).ToNot(BeNil())
	tc.Expect(string(*agentPool.Status.OsType)).To(Equal(string(osType)))
	tc.Expect(agentPool.Status.OsSKU).ToNot(BeNil())
	tc.Expect(*agentPool.Status.OsSKU).To(Equal(aks.OSSKU_STATUS_AzureContainerLinux))
}
