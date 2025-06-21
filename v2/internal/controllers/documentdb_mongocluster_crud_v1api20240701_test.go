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

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_DocumentDB_MongoCluster_v1api20240701_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("eastus")

	// Create our resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a secret for admin password
	password := tc.Namer.GeneratePasswordOfLength(40)

	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("aso-mongodb-password"),
		StringData: map[string]string{
			"password": password,
		},
	}

	adminPasswordSecretRef := genruntime.SecretReference{
		Name: secret.Name,
		Key:  "password",
	}

	// Declare a MongoCluster
	clusterName := tc.Namer.GenerateName("aso-mongocluster")
	mongoCluster := documentdb.MongoCluster{
		ObjectMeta: tc.MakeObjectMetaWithName(clusterName),
		Spec: documentdb.MongoCluster_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &documentdb.MongoClusterProperties{
				Administrator: &documentdb.AdministratorProperties{
					UserName: to.Ptr("testadmin"),
					Password: &adminPasswordSecretRef,
				},
				Compute: &documentdb.ComputeProperties{
					Tier: to.Ptr("M30"),
				},
				Storage: &documentdb.StorageProperties{
					SizeGb: to.Ptr(32),
				},
				ServerVersion:       to.Ptr("7.0"),
				PublicNetworkAccess: to.Ptr(documentdb.PublicNetworkAccess_Enabled),
				HighAvailability: &documentdb.HighAvailabilityProperties{
					TargetMode: to.Ptr(documentdb.HighAvailabilityMode_ZoneRedundantPreferred),
				},
				PreviewFeatures: []documentdb.PreviewFeature{documentdb.PreviewFeature_GeoReplicas},
				Sharding: &documentdb.ShardingProperties{
					ShardCount: to.Ptr(1),
				},
			},
			OperatorSpec: &documentdb.MongoClusterOperatorSpec{
				SecretExpressions: []*core.DestinationExpression{
					{
						Name:  "authsecret",
						Key:   "connectionString",
						Value: "self.status.properties.connectionString",
					},
					{
						Name:  "authsecret",
						Key:   "username",
						Value: "self.spec.properties.administrator.userName",
					},
					{
						Name:  "authsecret",
						Key:   "servername",
						Value: "self.spec.azureName",
					},
				},
			},
			Tags: map[string]string{
				"environment": "test",
				"project":     "aso",
			},
		},
	}

	// Create the MongoCluster in Azure
	tc.CreateResourcesAndWait(secret, &mongoCluster)

	// Perform some assertions on the resource we just created
	tc.Expect(mongoCluster.Status.Id).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Name).ToNot(BeNil())
	tc.Expect(*mongoCluster.Status.Name).To(Equal(clusterName))
	tc.Expect(mongoCluster.Status.Properties).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Properties.ClusterStatus).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Properties.ProvisioningState).ToNot(BeNil())

	// Verify administrator properties
	tc.Expect(mongoCluster.Status.Properties.Administrator).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Properties.Administrator.UserName).ToNot(BeNil())
	tc.Expect(*mongoCluster.Status.Properties.Administrator.UserName).To(Equal("testadmin"))

	// Verify compute properties
	tc.Expect(mongoCluster.Status.Properties.Compute).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Properties.Compute.Tier).ToNot(BeNil())
	tc.Expect(*mongoCluster.Status.Properties.Compute.Tier).To(Equal("M30"))

	// Verify storage properties
	tc.Expect(mongoCluster.Status.Properties.Storage).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Properties.Storage.SizeGb).ToNot(BeNil())
	tc.Expect(*mongoCluster.Status.Properties.Storage.SizeGb).To(Equal(32))

	// Verify server version
	tc.Expect(mongoCluster.Status.Properties.ServerVersion).ToNot(BeNil())
	tc.Expect(*mongoCluster.Status.Properties.ServerVersion).To(Equal("7.0"))

	// Verify high availability
	tc.Expect(mongoCluster.Status.Properties.HighAvailability).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Properties.HighAvailability.TargetMode).ToNot(BeNil())

	// Verify sharding
	tc.Expect(mongoCluster.Status.Properties.Sharding).ToNot(BeNil())
	tc.Expect(mongoCluster.Status.Properties.Sharding.ShardCount).ToNot(BeNil())
	tc.Expect(*mongoCluster.Status.Properties.Sharding.ShardCount).To(Equal(1))

	// Verify that the connection string secret was created
	tc.LogSectionf("Verifying connection string secret was created")
	tc.ExpectSecretHasKeys("authsecret", "connectionString", "username", "servername")

	// Run sub-tests that may create child resources
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "MongoCluster FirewallRule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DocumentDB_MongoCluster_FirewallRule_v1api20240701_CRUD(tc, &mongoCluster)
			},
		})

	// Delete the cluster and make sure it goes away
	armId := *mongoCluster.Status.Id
	tc.DeleteResourceAndWait(&mongoCluster)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(documentdb.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func DocumentDB_MongoCluster_FirewallRule_v1api20240701_CRUD(tc *testcommon.KubePerTestContext, mongoCluster client.Object) {
	// Create a firewall rule for the mongo cluster
	firewallRuleName := tc.Namer.GenerateName("firewallrule")
	firewallRule := documentdb.FirewallRule{
		ObjectMeta: tc.MakeObjectMetaWithName(firewallRuleName),
		Spec: documentdb.FirewallRule_Spec{
			Owner: testcommon.AsOwner(mongoCluster),
			Properties: &documentdb.FirewallRuleProperties{
				StartIpAddress: to.Ptr("192.168.1.1"),
				EndIpAddress:   to.Ptr("192.168.1.255"),
			},
		},
	}

	tc.T.Log("creating firewall rule")
	tc.CreateResourceAndWait(&firewallRule)
	defer func() {
		tc.LogSectionf("Cleaning up firewall rule")
		tc.DeleteResourceAndWait(&firewallRule)
	}()

	// Verify the firewall rule was created correctly
	tc.Expect(firewallRule.Status.Id).ToNot(BeNil())
	tc.Expect(firewallRule.Status.Name).ToNot(BeNil())
	tc.Expect(*firewallRule.Status.Name).To(Equal(firewallRuleName))
	tc.Expect(firewallRule.Status.Properties).ToNot(BeNil())
	tc.Expect(firewallRule.Status.Properties.StartIpAddress).ToNot(BeNil())
	tc.Expect(*firewallRule.Status.Properties.StartIpAddress).To(Equal("192.168.1.1"))
	tc.Expect(firewallRule.Status.Properties.EndIpAddress).ToNot(BeNil())
	tc.Expect(*firewallRule.Status.Properties.EndIpAddress).To(Equal("192.168.1.255"))

	// Update the firewall rule
	tc.LogSectionf("Updating firewall rule IP range")
	old := firewallRule.DeepCopy()
	firewallRule.Spec.Properties.StartIpAddress = to.Ptr("10.0.0.1")
	firewallRule.Spec.Properties.EndIpAddress = to.Ptr("10.0.0.255")
	tc.PatchResourceAndWait(old, &firewallRule)
	tc.Expect(firewallRule.Status.Properties.StartIpAddress).ToNot(BeNil())
	tc.Expect(*firewallRule.Status.Properties.StartIpAddress).To(Equal("10.0.0.1"))
	tc.Expect(firewallRule.Status.Properties.EndIpAddress).ToNot(BeNil())
	tc.Expect(*firewallRule.Status.Properties.EndIpAddress).To(Equal("10.0.0.255"))
}
