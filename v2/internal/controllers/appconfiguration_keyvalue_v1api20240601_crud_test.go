/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	appconfig "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v20240601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_AppConfiguration_KeyValue_v1api20240601_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("can't run in live mode, as ConfigurationStore retains the name after deletion which results in conflicts")
	}

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create ConfigurationStore
	cs := &appconfig.ConfigurationStore{
		ObjectMeta: tc.MakeObjectMeta("confstore"),
		Spec: appconfig.ConfigurationStore_Spec{
			CreateMode: to.Ptr(appconfig.ConfigurationStoreProperties_CreateMode_Default),
			Location:   tc.AzureRegion,
			Owner:      testcommon.AsOwner(rg),
			Sku: &appconfig.Sku{
				Name: to.Ptr("standard"),
			},
			PublicNetworkAccess: to.Ptr(appconfig.ConfigurationStoreProperties_PublicNetworkAccess_Enabled),
		},
	}

	tc.CreateResourceAndWait(cs)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "AppConfiguration KeyValue CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AppConfiguration_KeyValue_v1api20240601_CRUD(tc, cs)
			},
		},
		testcommon.Subtest{
			Name: "AppConfiguration Replica CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AppConfiguration_Replica_v1api20240601_CRUD(tc, cs)
			},
		},
		testcommon.Subtest{
			Name: "AppConfiguration Snapshot CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AppConfiguration_Snapshot_v1api20240601_CRUD(tc, cs)
			},
		})

	tc.DeleteResourceAndWait(cs)
}

func AppConfiguration_KeyValue_v1api20240601_CRUD(tc *testcommon.KubePerTestContext, cs *appconfig.ConfigurationStore) {
	tc.LogSectionf("Creating KeyValue")

	keyValueName := tc.Namer.GenerateName("keyvalue")
	keyValue := &appconfig.KeyValue{
		ObjectMeta: tc.MakeObjectMetaWithName(keyValueName),
		Spec: appconfig.KeyValue_Spec{
			AzureName:   "MyApp:Settings:TestKey",
			Owner:       testcommon.AsOwner(cs),
			Value:       to.Ptr("TestValue"),
			ContentType: to.Ptr("text/plain"),
			Tags: map[string]string{
				"Environment": "Test",
				"Team":        "Backend",
			},
		},
	}

	tc.CreateResourceAndWait(keyValue)

	armId := *keyValue.Status.Id
	tc.Expect(keyValue.Status.Key).ToNot(BeNil())
	tc.Expect(*keyValue.Status.Key).To(Equal("MyApp:Settings:TestKey"))
	tc.Expect(keyValue.Status.Value).ToNot(BeNil())
	tc.Expect(*keyValue.Status.Value).To(Equal("TestValue"))

	// Update the value
	old := keyValue.DeepCopy()
	keyValue.Spec.Value = to.Ptr("UpdatedTestValue")
	tc.PatchResourceAndWait(old, keyValue)
	tc.Expect(*keyValue.Status.Value).To(Equal("UpdatedTestValue"))

	tc.DeleteResourceAndWait(keyValue)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(appconfig.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func AppConfiguration_Replica_v1api20240601_CRUD(tc *testcommon.KubePerTestContext, cs *appconfig.ConfigurationStore) {
	tc.LogSectionf("Creating Replica")

	replicaName := tc.Namer.GenerateName("replica")
	replica := &appconfig.Replica{
		ObjectMeta: tc.MakeObjectMetaWithName(replicaName),
		Spec: appconfig.Replica_Spec{
			AzureName: "eastus2-replica",
			Owner:     testcommon.AsOwner(cs),
			Location:  to.Ptr("eastus2"),
		},
	}

	tc.CreateResourceAndWait(replica)

	armId := *replica.Status.Id
	tc.Expect(replica.Status.Location).ToNot(BeNil())
	tc.Expect(*replica.Status.Location).To(Equal("eastus2"))
	tc.Expect(replica.Status.Endpoint).ToNot(BeNil())
	tc.Expect(replica.Status.ProvisioningState).ToNot(BeNil())

	tc.DeleteResourceAndWait(replica)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(appconfig.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func AppConfiguration_Snapshot_v1api20240601_CRUD(tc *testcommon.KubePerTestContext, cs *appconfig.ConfigurationStore) {
	tc.LogSectionf("Creating Snapshot")

	snapshotName := tc.Namer.GenerateName("snapshot")
	snapshot := &appconfig.Snapshot{
		ObjectMeta: tc.MakeObjectMetaWithName(snapshotName),
		Spec: appconfig.Snapshot_Spec{
			AzureName:       "test-snapshot",
			Owner:           testcommon.AsOwner(cs),
			CompositionType: to.Ptr(appconfig.SnapshotProperties_CompositionType_Key),
			Filters: []appconfig.KeyValueFilter{
				{
					Key: to.Ptr("MyApp:*"),
				},
			},
			RetentionPeriod: to.Ptr(60 * 60 * 24 * 30), // 30 days
			Tags: map[string]string{
				"Environment": "Test",
				"Purpose":     "Testing",
			},
		},
	}

	tc.CreateResourceAndWait(snapshot)

	armId := *snapshot.Status.Id
	tc.Expect(snapshot.Status.CompositionType).ToNot(BeNil())
	tc.Expect(snapshot.Status.ProvisioningState).ToNot(BeNil())
	tc.Expect(snapshot.Status.ItemsCount).ToNot(BeNil())

	tc.DeleteResourceAndWait(snapshot)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(appconfig.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
