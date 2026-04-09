/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"
	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v20251015"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DocumentDB_CassandraCluster_v1api20251015_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("eastus")

	// Create our resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a password secret for the Cassandra admin
	password := tc.Namer.GeneratePasswordOfLength(40)

	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("cass-pwd"),
		StringData: map[string]string{
			"password": password,
		},
	}

	adminPasswordSecretRef := genruntime.SecretReference{
		Name: secret.Name,
		Key:  "password",
	}

	// Create the VNet for Cassandra networking
	vnet := newCassandraVirtualNetwork(tc, testcommon.AsOwner(rg))

	// Create the management subnet (for DelegatedManagementSubnetId)
	mgmtSubnet := newCassandraManagementSubnet(tc, testcommon.AsOwner(vnet))
	mgmtSubnetRoleAssignment := newCassandraRoleAssignment(tc, "mgmtsubnetroleassignment", mgmtSubnet)

	// Create the data center subnet (for DelegatedSubnetId) with delegation
	dcSubnet := newCassandraDataCenterSubnet(tc, testcommon.AsOwner(vnet))
	dcSubnetRoleAssignment := newCassandraRoleAssignment(tc, "dcsubnetroleassignment", dcSubnet)

	// Declare the CassandraCluster
	clusterName := tc.Namer.GenerateName("cassandracluster")
	cassandraCluster := &documentdb.CassandraCluster{
		ObjectMeta: tc.MakeObjectMetaWithName(clusterName),
		Spec: documentdb.CassandraCluster_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &documentdb.CassandraCluster_Properties_Spec{
				CassandraVersion: to.Ptr("4.0"),
				DelegatedManagementSubnetReference: &genruntime.ResourceReference{
					Group: network.GroupVersion.Group,
					Kind:  "VirtualNetworksSubnet",
					Name:  mgmtSubnet.Name,
				},
				InitialCassandraAdminPassword: &adminPasswordSecretRef,
				RepairEnabled:                 to.Ptr(true),
			},
		},
	}

	// Create a data center for the Cassandra cluster
	dataCenter := newCassandraDataCenter(tc, cassandraCluster, dcSubnet)

	// Create all resources together, mirroring what a user experiences when applying a YAML file
	tc.CreateResourcesAndWait(
		secret,
		vnet,
		mgmtSubnet,
		dcSubnet,
		mgmtSubnetRoleAssignment,
		dcSubnetRoleAssignment,
		cassandraCluster,
		dataCenter)

	// Perform some assertions on the cluster we just created
	tc.Expect(cassandraCluster.Status.Id).ToNot(BeNil())
	tc.Expect(cassandraCluster.Status.Name).ToNot(BeNil())
	tc.Expect(*cassandraCluster.Status.Name).To(Equal(clusterName))
	tc.Expect(cassandraCluster.Status.Properties).ToNot(BeNil())
	tc.Expect(cassandraCluster.Status.Properties.ProvisioningState).ToNot(BeNil())

	// Verify the data center was created correctly
	tc.Expect(dataCenter.Status.Id).ToNot(BeNil())
	tc.Expect(dataCenter.Status.Name).ToNot(BeNil())
	tc.Expect(*dataCenter.Status.Name).To(Equal(dataCenter.Name))
	tc.Expect(dataCenter.Status.Properties).ToNot(BeNil())
	tc.Expect(dataCenter.Status.Properties.ProvisioningState).ToNot(BeNil())
	tc.Expect(dataCenter.Status.Properties.NodeCount).ToNot(BeNil())
	tc.Expect(*dataCenter.Status.Properties.NodeCount).To(Equal(3))

	// No test for Cluster Updates as it appears every property is read/only once after creation as the RP silently rejects changes

	// Update the data center to ensure that works
	oldDataCenter := dataCenter.DeepCopy()
	dataCenter.Spec.Properties.NodeCount = to.Ptr(5)
	tc.PatchResourceAndWait(oldDataCenter, dataCenter)

	// Delete the cluster and make sure it goes away
	armId := *cassandraCluster.Status.Id
	tc.DeleteResourceAndWait(cassandraCluster)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(documentdb.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func DocumentDB_CassandraCluster_DataCenter_v1api20251015_CRUD(
	tc *testcommon.KubePerTestContext,
	cassandraCluster client.Object,
	dcSubnet *network.VirtualNetworksSubnet,
) {
}

func newCassandraDataCenter(tc *testcommon.KubePerTestContext, cassandraCluster client.Object, dcSubnet *network.VirtualNetworksSubnet) *documentdb.CassandraDataCenter {
	dcName := tc.Namer.GenerateName("dc")
	dataCenter := &documentdb.CassandraDataCenter{
		ObjectMeta: tc.MakeObjectMetaWithName(dcName),
		Spec: documentdb.CassandraDataCenter_Spec{
			Owner: testcommon.AsOwner(cassandraCluster),
			Properties: &documentdb.CassandraClusters_DataCenter_Properties_Spec{
				DataCenterLocation: to.Ptr("eastus"),
				DelegatedSubnetReference: &genruntime.ResourceReference{
					Group: network.GroupVersion.Group,
					Kind:  "VirtualNetworksSubnet",
					Name:  dcSubnet.Name,
				},
				NodeCount: to.Ptr(3),
			},
		},
	}

	return dataCenter
}

func newCassandraVirtualNetwork(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetwork {
	return &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetwork_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
}

func newCassandraManagementSubnet(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetworksSubnet {
	return &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         owner,
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}
}

func newCassandraDataCenterSubnet(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetworksSubnet {
	return &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         owner,
			AddressPrefix: to.Ptr("10.0.1.0/24"),
			Delegations: []network.Delegation{
				{
					Name:        to.Ptr("cassandraClustersDelegation"),
					ServiceName: to.Ptr("Microsoft.DocumentDB/cassandraClusters"),
				},
			},
		},
	}
}

func newCassandraRoleAssignment(
	tc *testcommon.KubePerTestContext,
	name string,
	owner client.Object,
) *authorization.RoleAssignment {
	result := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: authorization.RoleAssignment_Spec{
			Owner: tc.AsExtensionOwner(owner),
			// Needs to be the principle within the current subscription
			// Hard coded for testing, because ASO doesn't currently allow looking this up automatically
			// TODO: Remove this once ASO can do the lookup.
			PrincipalId:   to.Ptr("e5007d2c-4b13-4a74-9b6a-605d99f03501"),
			PrincipalType: to.Ptr(authorization.RoleAssignmentProperties_PrincipalType_ServicePrincipal),
			RoleDefinitionReference: &genruntime.WellKnownResourceReference{
				WellKnownName: "Network Contributor",
			},
		},
	}

	// Don't try to delete directly, we get authorization denied errors doing this.
	// Instead, use this annotation, knowing the assignment will be cleaned up when the subnet is deleted.
	tc.AddAnnotation(&result.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	return result
}
