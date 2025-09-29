/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20241001"
	network_prev "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Test_Networking_NetworkWatcher_20241001_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// NetworkWatcher
	networkWatcher := &network.NetworkWatcher{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("nwatcher")),
		Spec: network.NetworkWatcher_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(networkWatcher)

	tc.Expect(networkWatcher.Status.Id).ToNot(BeNil())
	armId := *networkWatcher.Status.Id

	// Perform a simple patch by adding tags
	old := networkWatcher.DeepCopy()
	networkWatcher.Spec.Tags = map[string]string{
		"environment": "test",
		"purpose":     "monitoring",
	}
	tc.PatchResourceAndWait(old, networkWatcher)
	tc.Expect(networkWatcher.Status.Tags).To(HaveKey("environment"))
	tc.Expect(networkWatcher.Status.Tags).To(HaveKey("purpose"))

	// Run sub tests for FlowLog
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "FlowLog CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				NetworkWatcher_FlowLog_20241001_CRUD(tc, networkWatcher)
			},
		},
	)

	tc.DeleteResourceAndWait(networkWatcher)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func NetworkWatcher_FlowLog_20241001_CRUD(tc *testcommon.KubePerTestContext, networkWatcher client.Object) {
	// Create a storage account for the flow log
	storageAccount := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(networkWatcher.GetOwnerReferences()[0]),
			Sku: &storage.Sku{
				Name: to.Ptr(storage.Sku_Name_Standard_LRS),
			},
			Kind: to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
		},
	}

	// Create an NSG as target for the flow log
	nsg := &network_prev.NetworkSecurityGroup{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("nsg")),
		Spec: network_prev.NetworkSecurityGroup_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(networkWatcher.GetOwnerReferences()[0]),
		},
	}

	tc.CreateResourcesAndWait(storageAccount, nsg)
	
	// Get storage account ARM ID for the FlowLog
	tc.Expect(storageAccount.Status.Id).ToNot(BeNil())
	storageAccountArmId := *storageAccount.Status.Id

	// NetworkWatchersFlowLog
	flowLog := &network.NetworkWatchersFlowLog{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("flowlog")),
		Spec: network.NetworkWatchersFlowLog_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(networkWatcher),
			Enabled:  to.Ptr(true),
			TargetResourceReference: &genruntime.ResourceReference{
				Reference: tc.MakeReferenceFromResource(nsg),
			},
			StorageId: &storageAccountArmId,
			RetentionPolicy: &network.RetentionPolicyParameters{
				Enabled: to.Ptr(true),
				Days:    to.Ptr(7), // 7 days retention for testing
			},
			Format: &network.FlowLogFormatParameters{
				Type:    to.Ptr(network.FlowLogFormatParameters_Type_JSON),
				Version: to.Ptr(2), // Use JSON format version 2
			},
		},
	}

	tc.CreateResourceAndWait(flowLog)

	tc.Expect(flowLog.Status.Id).ToNot(BeNil())
	flowLogArmId := *flowLog.Status.Id

	// Verify the flow log was created with correct properties
	tc.Expect(flowLog.Status.Enabled).To(Equal(to.Ptr(true)))
	tc.Expect(flowLog.Status.RetentionPolicy).ToNot(BeNil())
	tc.Expect(flowLog.Status.RetentionPolicy.Enabled).To(Equal(to.Ptr(true)))
	tc.Expect(flowLog.Status.RetentionPolicy.Days).To(Equal(to.Ptr(7)))
	tc.Expect(flowLog.Status.Format).ToNot(BeNil())
	tc.Expect(flowLog.Status.Format.Type).To(Equal(to.Ptr(network.FlowLogFormatParameters_Type_STATUS_JSON)))
	tc.Expect(flowLog.Status.Format.Version).To(Equal(to.Ptr(2)))

	// Test updating flow log - change retention to 14 days
	oldFlowLog := flowLog.DeepCopy()
	flowLog.Spec.RetentionPolicy.Days = to.Ptr(14)
	flowLog.Spec.Tags = map[string]string{
		"updated": "true",
	}
	tc.PatchResourceAndWait(oldFlowLog, flowLog)
	tc.Expect(flowLog.Status.RetentionPolicy.Days).To(Equal(to.Ptr(14)))
	tc.Expect(flowLog.Status.Tags).To(HaveKey("updated"))

	// Test disabling flow log
	oldFlowLog2 := flowLog.DeepCopy()
	flowLog.Spec.Enabled = to.Ptr(false)
	tc.PatchResourceAndWait(oldFlowLog2, flowLog)
	tc.Expect(flowLog.Status.Enabled).To(Equal(to.Ptr(false)))

	tc.DeleteResourceAndWait(flowLog)

	// Ensure that the flow log was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, flowLogArmId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

	// Clean up NSG and storage account
	tc.DeleteResourcesAndWait(nsg, storageAccount)
}