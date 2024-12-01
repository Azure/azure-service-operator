/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20240402preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func Test_Latest_Reconciled_Generation_Reconciles_AllEvents(t *testing.T) {
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
			KubernetesVersion: to.Ptr("1.30.0"),
		},
	}

	tc.CreateResourceAndWait(cluster)

	agentPool := &aks.ManagedClustersAgentPool{
		ObjectMeta: tc.MakeObjectMetaWithName("ap2"),
		Spec: aks.ManagedClustersAgentPool_Spec{
			Owner:               testcommon.AsOwner(cluster),
			Count:               to.Ptr(1),
			VmSize:              to.Ptr("Standard_DS2_v2"),
			OsType:              to.Ptr(aks.OSType_Linux),
			Mode:                to.Ptr(aks.AgentPoolMode_User),
			OrchestratorVersion: to.Ptr("1.29.5"),
		},
	}

	tc.CreateResourceAndWait(agentPool)

	old := agentPool.DeepCopy()
	agentPool.Spec.OrchestratorVersion = to.Ptr("1.30.0") // Start an upgrade that we know will take a bit

	tc.PatchResourceAndWaitForState(old, agentPool, metav1.ConditionFalse, conditions.ConditionSeverityInfo)

	old = agentPool.DeepCopy()
	agentPool.Spec.Count = to.Ptr(2) // Scale count immediately after upgrade starts

	tc.Patch(old, agentPool)

	objectKey := client.ObjectKeyFromObject(agentPool)

	tc.Eventually(func() bool {
		var updated aks.ManagedClustersAgentPool
		tc.GetResource(objectKey, &updated)
		return *updated.Status.OrchestratorVersion == "1.30.0" && *updated.Status.Count == 2
	}).Should(BeTrue())

	tc.DeleteResourcesAndWait(agentPool, cluster)
}
