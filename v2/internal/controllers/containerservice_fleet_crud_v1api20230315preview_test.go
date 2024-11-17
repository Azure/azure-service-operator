/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201"
	fleet "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230315preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AKS_Fleet_20230315_CRUD(t *testing.T) {
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
				Type: to.Ptr(aks.ManagedClusterIdentity_Type_SystemAssigned),
			},
			OidcIssuerProfile: &aks.ManagedClusterOIDCIssuerProfile{
				Enabled: to.Ptr(true),
			},
			KubernetesVersion: to.Ptr("1.29.5"),
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
				AKS_Fleet_FleetMember_20230315Preview_CRUD(tc, flt, clusterArmID)
			},
		},
		testcommon.Subtest{
			Name: "Fleet UpdateRun CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_Fleet_UpdateRun_20230315Preview_CRUD(tc, flt)
			},
		})

	// delete managed cluster
	tc.DeleteResourceAndWait(cluster)

	// delete a fleet
	tc.DeleteResourceAndWait(flt)

	// Ensure that fleet was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(fleet.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func AKS_Fleet_FleetMember_20230315Preview_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet, clusterArmID string) {
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

func AKS_Fleet_UpdateRun_20230315Preview_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet) {
	updateRun := &fleet.FleetsUpdateRun{
		ObjectMeta: tc.MakeObjectMeta("updaterun"),
		Spec: fleet.FleetsUpdateRun_Spec{
			ManagedClusterUpdate: &fleet.ManagedClusterUpdate{
				Upgrade: &fleet.ManagedClusterUpgradeSpec{
					Type:              to.Ptr(fleet.ManagedClusterUpgradeType_Full),
					KubernetesVersion: to.Ptr("1.30.0"),
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
