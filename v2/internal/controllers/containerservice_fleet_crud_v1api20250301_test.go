/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	fleet "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250301"
	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AKS_Fleet_20250301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("westus3") // TODO: the default test region of westus2 doesn't allow ds2_v2 at the moment

	rg := tc.CreateTestResourceGroupAndWait()
	region := tc.AzureRegion
	flt := &fleet.Fleet{
		ObjectMeta: tc.MakeObjectMeta("fleet"),
		Spec: fleet.Fleet_Spec{
			Location: region,
			Owner:    testcommon.AsOwner(rg),
			HubProfile: &fleet.FleetHubProfile{
				DnsPrefix: to.Ptr("aso"),
			},
			Tags: map[string]string{
				"name": "test-tag",
			},
		},
	}
	// creating a fleet
	tc.CreateResourceAndWait(flt)
	tc.Expect(flt.Status.Id).ToNot(BeNil())
	tc.Expect(flt.Status.Tags).ToNot(BeNil())
	tc.Expect(flt.Spec.Tags["name"]).To(Equal("test-tag"))
	armId := *flt.Status.Id

	// patching a fleet
	old := flt.DeepCopy()
	flt.Spec.Tags = map[string]string{
		"name": "test-tag2",
	}
	tc.PatchResourceAndWait(old, flt)
	tc.Expect(flt.Spec.Tags["name"]).To(Equal("test-tag2"))

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

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
			OidcIssuerProfile: &aks.ManagedClusterOIDCIssuerProfile{
				Enabled: to.Ptr(true),
			},
			KubernetesVersion: to.Ptr("1.33.5"),
		},
	}
	tc.CreateResourceAndWait(cluster)
	tc.Expect(cluster.Status.Id).ToNot(BeNil())
	clusterArmID := *cluster.Status.Id

	// Run sub tests - updateRun subtest depends on fleetMember subtest
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Fleet FleetMember CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_Fleet_FleetMember_20250301_CRUD(tc, flt, clusterArmID)
			},
		},
		testcommon.Subtest{
			Name: "Fleet UpdateRun CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_Fleet_UpdateRun_20250301_CRUD(tc, flt)
			},
		},
		testcommon.Subtest{
			Name: "Fleet UpdateStrategy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_Fleet_UpdateStrategy_202150301_CRUD(tc, flt)
			},
		},
	)

	// Delete the resources
	tc.DeleteResourcesAndWait(cluster, flt)

	// Ensure that fleet was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(fleet.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func AKS_Fleet_FleetMember_20250301_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet, clusterArmID string) {
	fltMember := &fleet.FleetsMember{
		ObjectMeta: tc.MakeObjectMeta("fleetmember"),
		Spec: fleet.FleetsMember_Spec{
			Owner: testcommon.AsOwner(flt),
			ClusterResourceReference: &genruntime.ResourceReference{
				ARMID: clusterArmID,
			},
		},
	}

	// create fleet member
	tc.CreateResourceAndWait(fltMember)
	tc.Expect(fltMember.Status.Id).ToNot(BeNil())
	tc.Expect(fltMember.Status.ClusterResourceId).ToNot(BeNil())

	// not deleting fleet member since updateRun test depends on resource
}

func AKS_Fleet_UpdateRun_20250301_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet) {
	updateRun := &fleet.FleetsUpdateRun{
		ObjectMeta: tc.MakeObjectMeta("updaterun"),
		Spec: fleet.FleetsUpdateRun_Spec{
			ManagedClusterUpdate: &fleet.ManagedClusterUpdate{
				Upgrade: &fleet.ManagedClusterUpgradeSpec{
					Type:              to.Ptr(fleet.ManagedClusterUpgradeType_Full),
					KubernetesVersion: to.Ptr("1.32.7"),
				},
			},
			Owner: testcommon.AsOwner(flt),
		},
	}

	tc.CreateResourceAndWait(updateRun)

	defer tc.DeleteResourceAndWait(updateRun)

	tc.Expect(updateRun.Status.Id).ToNot(BeNil())

	// a basic assertion on a few properties
	tc.Expect(*updateRun.Status.ManagedClusterUpdate.Upgrade.Type).To(Equal(fleet.ManagedClusterUpgradeType_STATUS_Full))

	// Perform a simple patch
	old := updateRun.DeepCopy()
	updateRun.Spec.ManagedClusterUpdate.Upgrade.Type = to.Ptr(fleet.ManagedClusterUpgradeType_NodeImageOnly)
	updateRun.Spec.ManagedClusterUpdate.Upgrade.KubernetesVersion = nil // This must be nil for NodeImageOnly upgrade

	tc.PatchResourceAndWait(old, updateRun)

	tc.Expect(*updateRun.Status.ManagedClusterUpdate.Upgrade.Type).To(Equal(fleet.ManagedClusterUpgradeType_STATUS_NodeImageOnly))
}

func AKS_Fleet_UpdateStrategy_202150301_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet) {
	updateStrategy := &fleet.FleetsUpdateStrategy{
		ObjectMeta: tc.MakeObjectMeta("updatestrategy"),
		Spec: fleet.FleetsUpdateStrategy_Spec{
			Owner: testcommon.AsOwner(flt),
			Strategy: &fleet.UpdateRunStrategy{
				Stages: []fleet.UpdateStage{
					{
						Name: to.Ptr("stage1"),
						Groups: []fleet.UpdateGroup{
							{
								Name: to.Ptr("group1"),
							},
						},
					},
				},
			},
		},
	}

	// Create and verify UpdateStrategy
	tc.CreateResourceAndWait(updateStrategy)
	tc.Expect(updateStrategy.Status.Id).ToNot(BeNil())
	tc.Expect(updateStrategy.Status.Strategy).ToNot(BeNil())
	tc.Expect(updateStrategy.Status.Strategy.Stages).To(HaveLen(1))
	tc.Expect(updateStrategy.Status.Strategy.Stages[0].Name).ToNot(BeNil())
	tc.Expect(*updateStrategy.Status.Strategy.Stages[0].Name).To(Equal("stage1"))

	tc.DeleteResourceAndWait(updateStrategy)
}
