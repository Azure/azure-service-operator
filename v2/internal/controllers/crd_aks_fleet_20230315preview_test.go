/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	fleet "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230315preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	. "github.com/onsi/gomega"
)

func Test_AKS_Fleet_20230315_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	region := to.Ptr("westus3")
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
	// armId := *flt.Status.Id

	// patching a fleet
	old := flt.DeepCopy()
	flt.Spec.Tags = map[string]string{
		"name": "test-tag2",
	}
	tc.PatchResourceAndWait(old, flt)
	tc.Expect(flt.Spec.Tags["name"]).To(Equal("test-tag2"))

	// TODO: Move create cluster code to fleet member test
	// if *isLive {
	// 	t.Skip("can't run in live mode, as this test is creates a KeyVault which reserves the name unless manually purged")
	// }

	// adminUsername := "adminUser"
	// sshPublicKey, err := tc.GenerateSSHKey(2048)
	// tc.Expect(err).ToNot(HaveOccurred())

	// identityKind := aks.ManagedClusterIdentity_Type_SystemAssigned
	// osType := aks.OSType_Linux
	// agentPoolMode := aks.AgentPoolMode_System

	// cluster := &aks.ManagedCluster{
	// 	ObjectMeta: tc.MakeObjectMeta("mc"),
	// 	Spec: aks.ManagedCluster_Spec{
	// 		Location:  region,
	// 		Owner:     testcommon.AsOwner(rg),
	// 		DnsPrefix: to.Ptr("aso"),
	// 		AgentPoolProfiles: []aks.ManagedClusterAgentPoolProfile{
	// 			{
	// 				Name:   to.Ptr("ap1"),
	// 				Count:  to.Ptr(1),
	// 				VmSize: to.Ptr("Standard_DS2_v2"),
	// 				OsType: &osType,
	// 				Mode:   &agentPoolMode,
	// 			},
	// 		},
	// 		LinuxProfile: &aks.ContainerServiceLinuxProfile{
	// 			AdminUsername: &adminUsername,
	// 			Ssh: &aks.ContainerServiceSshConfiguration{
	// 				PublicKeys: []aks.ContainerServiceSshPublicKey{
	// 					{
	// 						KeyData: sshPublicKey,
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Identity: &aks.ManagedClusterIdentity{
	// 			Type: &identityKind,
	// 		},
	// 		OidcIssuerProfile: &aks.ManagedClusterOIDCIssuerProfile{
	// 			Enabled: to.Ptr(true),
	// 		},
	// 	},
	// }

	// tc.CreateResourceAndWait(cluster)

	// tc.Expect(cluster.Status.Id).ToNot(BeNil())
	// armId := *cluster.Status.Id

	// Run sub tests
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Fleet FleetMember CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_Fleet_FleetMember_20230315Preview_CRUD(tc, flt)
			},
		},
		testcommon.Subtest{
			Name: "Fleet UpdateRun CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_Fleet_UpdateRun_20230315Preview_CRUD(tc, flt)
			},
		})

	// delete a fleet
	// tc.DeleteResourceAndWait(flt)

	// // Ensure that fleet was really deleted in Azure
	// exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(fleet.APIVersion_Value))
	// tc.Expect(err).ToNot(HaveOccurred())
	// tc.Expect(retryAfter).To(BeZero())
	// tc.Expect(exists).To(BeFalse())
}

func AKS_Fleet_FleetMember_20230315Preview_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet) {

	// if *isLive {
	// 	t.Skip("can't run in live mode, as this test is creates a KeyVault which reserves the name unless manually purged")
	// }

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	identityKind := aks.ManagedClusterIdentity_Type_SystemAssigned
	osType := aks.OSType_Linux
	agentPoolMode := aks.AgentPoolMode_System
	rg := tc.CreateTestResourceGroupAndWait()
	region := to.Ptr("westus3")
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
	flt_member := &fleet.FleetsMember{
		ObjectMeta: tc.MakeObjectMeta("fleetmember"),
		Spec: fleet.Fleets_Member_Spec{
			Owner: testcommon.AsOwner(flt),
			ClusterResourceReference: &genruntime.ResourceReference{
				ARMID: armId,
			},
		},
	}

	// create fleet member
	tc.CreateResourceAndWait(flt_member)
	tc.Expect(flt_member.Status.Id).ToNot(BeNil())
	tc.Expect(flt_member.Status.ClusterResourceId).ToNot(BeNil())

	// delete
	// defer tc.DeleteResourceAndWait(flt_member)

}

func AKS_Fleet_UpdateRun_20230315Preview_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet) {
	updateRun := &fleet.FleetsUpdateRun{
		ObjectMeta: tc.MakeObjectMeta("updaterun"),
		Spec: fleet.Fleets_UpdateRun_Spec{
			ManagedClusterUpdate: &fleet.ManagedClusterUpdate{
				Upgrade: &fleet.ManagedClusterUpgradeSpec{
					Type:              to.Ptr(fleet.ManagedClusterUpgradeType_Full),
					KubernetesVersion: to.Ptr("1.26.8"),
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

	tc.PatchResourceAndWait(old, updateRun)

	tc.Expect(*updateRun.Status.ManagedClusterUpdate.Upgrade.Type).To(Equal(fleet.ManagedClusterUpgradeType_STATUS_NodeImageOnly))
}
