/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	synapse "github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Workspace_BigDataPool(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	accountUrlKey := "accountUrl"

	// Create a storage account to use as workspace datalake storage
	sa := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("adlstore")),
		Spec: storage.StorageAccount_Spec{
			Location:     to.Ptr("eastus2"),
			Owner:        testcommon.AsOwner(rg),
			Kind:         to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku:          &storage.Sku{Name: to.Ptr(storage.SkuName_Standard_LRS)},
			IsHnsEnabled: to.Ptr(true),
			OperatorSpec: &storage.StorageAccountOperatorSpec{
				ConfigMaps: &storage.StorageAccountOperatorConfigMaps{
					DfsEndpoint: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  accountUrlKey,
					},
				},
			},
		},
	}

	ws := &synapse.Workspace{
		ObjectMeta: tc.MakeObjectMeta("workspace"),
		Spec: synapse.Workspace_Spec{
			Identity: &synapse.ManagedIdentity{
				Type: to.Ptr(synapse.ManagedIdentity_Type_SystemAssigned),
			},
			DefaultDataLakeStorage: &synapse.DataLakeStorageAccountDetails{
				AccountUrlFromConfig: &genruntime.ConfigMapReference{
					Name: configMapName,
					Key:  accountUrlKey,
				},
				Filesystem: to.Ptr("default"),
			},
			Location: to.Ptr("eastus2"),
			Owner:    testcommon.AsOwner(rg),
			Tags:     map[string]string{"cheese": "blue"},
		},
	}

	tc.CreateResourcesAndWait(sa, ws)

	tc.Expect(ws.Status.Id).ToNot(BeNil())
	wsArmId := *ws.Status.Id
	// Perform a simple patch
	old := ws.DeepCopy()
	ws.Spec.Tags["cheese"] = "époisses"
	tc.PatchResourceAndWait(old, ws)
	tc.Expect(ws.Status.Tags).To(Equal(map[string]string{"cheese": "époisses"}))

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Test_WorkspacesBigDataPool_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				WorkspacesBigDataPool_CRUD(tc, ws)
			},
		})

	tc.DeleteResourceAndWait(ws)

	// Ensure that the workspace was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		wsArmId,
		string(synapse.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func WorkspacesBigDataPool_CRUD(tc *testcommon.KubePerTestContext, workspaces *synapse.Workspace) {
	nodeSize := synapse.BigDataPoolResourceProperties_NodeSize_Medium
	nodeSizeFamily := synapse.BigDataPoolResourceProperties_NodeSizeFamily_MemoryOptimized

	pool := &synapse.WorkspacesBigDataPool{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("")),
		Spec: synapse.Workspaces_BigDataPool_Spec{
			Location:       to.Ptr("eastus2"),
			Owner:          testcommon.AsOwner(workspaces),
			SparkVersion:   to.Ptr("3.3"),
			NodeCount:      to.Ptr(4),
			NodeSize:       &nodeSize,
			NodeSizeFamily: &nodeSizeFamily,
			AutoScale: &synapse.AutoScaleProperties{
				Enabled: to.Ptr(false),
			},
			AutoPause: &synapse.AutoPauseProperties{
				Enabled: to.Ptr(false),
			},
		},
	}
	tc.CreateResourceAndWait(pool)

	tc.Expect(pool.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(pool)
}
