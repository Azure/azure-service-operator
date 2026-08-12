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

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v20260315"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DocumentDB_CassandraCluster_v20260315_CRUD(t *testing.T) {
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

	// Declare the CassandraCluster, exercising the properties 2026-03-15 adds over 2025-10-15:
	// extensions, autoReplicate and scheduledEventStrategy.
	clusterName := tc.Namer.GenerateName("cassandracluster")
	cassandraCluster := &documentdb.CassandraCluster{
		ObjectMeta: tc.MakeObjectMetaWithName(clusterName),
		Spec: documentdb.CassandraCluster_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &documentdb.ClusterResourceProperties{
				AutoReplicate:    to.Ptr(documentdb.AutoReplicate_SystemKeyspaces),
				CassandraVersion: to.Ptr("4.0"),
				DelegatedManagementSubnetReference: &genruntime.ResourceReference{
					Group: network.GroupVersion.Group,
					Kind:  "VirtualNetworksSubnet",
					Name:  mgmtSubnet.Name,
				},
				Extensions:                    []string{"cassandra-lucene-index"},
				InitialCassandraAdminPassword: &adminPasswordSecretRef,
				RepairEnabled:                 to.Ptr(true),
				ScheduledEventStrategy:        to.Ptr(documentdb.ScheduledEventStrategy_StopByRack),
			},
		},
	}

	// Create a data center for the Cassandra cluster
	dataCenter := newCassandraDataCenterV20260315(tc, cassandraCluster, dcSubnet)

	// Create all resources together, mirroring what a user experiences when applying a YAML file
	tc.CreateResourcesAndWait(
		secret,
		vnet,
		mgmtSubnet,
		dcSubnet,
		mgmtSubnetRoleAssignment,
		dcSubnetRoleAssignment,
		cassandraCluster,
		dataCenter,
	)

	// Perform some assertions on the cluster we just created
	tc.Expect(cassandraCluster.Status.Id).ToNot(BeNil())
	tc.Expect(cassandraCluster.Status.Name).ToNot(BeNil())
	tc.Expect(*cassandraCluster.Status.Name).To(Equal(clusterName))
	tc.Expect(cassandraCluster.Status.Properties).ToNot(BeNil())
	tc.Expect(cassandraCluster.Status.Properties.ProvisioningState).ToNot(BeNil())
	tc.Expect(cassandraCluster.Status.Properties.Extensions).To(ContainElement("cassandra-lucene-index"))
	tc.Expect(cassandraCluster.Status.Properties.AutoReplicate).ToNot(BeNil())
	tc.Expect(cassandraCluster.Status.Properties.ScheduledEventStrategy).ToNot(BeNil())

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
		string(documentdb.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func newCassandraDataCenterV20260315(tc *testcommon.KubePerTestContext, cassandraCluster client.Object, dcSubnet *network.VirtualNetworksSubnet) *documentdb.CassandraDataCenter {
	dcName := tc.Namer.GenerateName("dc")
	dataCenter := &documentdb.CassandraDataCenter{
		ObjectMeta: tc.MakeObjectMetaWithName(dcName),
		Spec: documentdb.CassandraDataCenter_Spec{
			Owner: testcommon.AsOwner(cassandraCluster),
			Properties: &documentdb.DataCenterResourceProperties{
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
