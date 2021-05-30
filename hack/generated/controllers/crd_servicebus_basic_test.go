/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	servicebus "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.servicebus/v1alpha1api20180101preview"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_ServiceBus_Basic_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateNewTestResourceGroupAndWait()

	zoneRedundant := false
	namespace := &servicebus.Namespace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sbnamespace")),
		Spec: servicebus.Namespaces_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &servicebus.SBSku{
				Name: servicebus.SBSkuNameBasic,
			},
			Properties: servicebus.SBNamespaceProperties{
				ZoneRedundant: &zoneRedundant,
			},
		},
	}

	tc.CreateAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Queue CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				ServiceBus_Queue_CRUD(testContext, namespace.ObjectMeta)
			},
		},
	)

	tc.DeleteAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, "2018-01-01-preview")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func ServiceBus_Queue_CRUD(testContext testcommon.KubePerTestContext, sbNamespace metav1.ObjectMeta) {
	queue := &servicebus.NamespacesQueue{
		ObjectMeta: testContext.MakeObjectMeta("queue"),
		Spec: servicebus.NamespacesQueues_Spec{
			Location: &testContext.AzureRegion,
			Owner:    testcommon.AsOwner(sbNamespace),
		},
	}

	testContext.CreateAndWait(queue)
	defer testContext.DeleteAndWait(queue)

	testContext.Expect(queue.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	testContext.Expect(queue.Status.Properties.SizeInBytes).ToNot(BeNil())
	testContext.Expect(*queue.Status.Properties.SizeInBytes).To(Equal(0))
}
