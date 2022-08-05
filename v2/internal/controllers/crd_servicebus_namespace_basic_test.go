/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ServiceBus_Namespace_Basic_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	zoneRedundant := false
	sku := servicebus.SBSku_Name_Basic
	namespace := &servicebus.Namespace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sbnamespace")),
		Spec: servicebus.Namespace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &servicebus.SBSku{
				Name: &sku,
			},
			ZoneRedundant: &zoneRedundant,
		},
	}

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Queue CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_Queue_CRUD(tc, namespace)
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(servicebus.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func ServiceBus_Queue_CRUD(tc *testcommon.KubePerTestContext, sbNamespace client.Object) {
	queue := &servicebus.NamespacesQueue{
		ObjectMeta: tc.MakeObjectMeta("queue"),
		Spec: servicebus.NamespacesQueue_Spec{
			Owner: testcommon.AsOwner(sbNamespace),
		},
	}

	tc.CreateResourceAndWait(queue)
	defer tc.DeleteResourceAndWait(queue)

	tc.Expect(queue.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(queue.Status.SizeInBytes).ToNot(BeNil())
	tc.Expect(*queue.Status.SizeInBytes).To(Equal(0))
}
