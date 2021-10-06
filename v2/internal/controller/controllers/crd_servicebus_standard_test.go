/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	servicebus "github.com/Azure/azure-service-operator/v2/api/microsoft.servicebus/v1alpha1api20210101preview"
	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"
)

func Test_ServiceBus_Standard_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	zoneRedundant := false
	namespace := &servicebus.Namespace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sbstandard")),
		Spec: servicebus.Namespaces_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &servicebus.SBSku{
				Name: servicebus.SBSkuNameStandard,
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
			Test: func(t testcommon.KubePerTestContext) { ServiceBus_Queue_CRUD(tc, namespace) },
		},
		testcommon.Subtest{
			Name: "Topic CRUD",
			Test: func(t testcommon.KubePerTestContext) { ServiceBus_Topic_CRUD(tc, namespace) },
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(servicebus.NamespacesSpecAPIVersion20210101Preview))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

// Topics can only be created in Standard or Premium SKUs
func ServiceBus_Topic_CRUD(tc testcommon.KubePerTestContext, sbNamespace client.Object) {
	topic := &servicebus.NamespacesTopic{
		ObjectMeta: tc.MakeObjectMeta("topic"),
		Spec: servicebus.NamespacesTopics_Spec{
			Location: &tc.AzureRegion,
			Owner:    testcommon.AsOwner(sbNamespace),
		},
	}

	tc.CreateResourceAndWait(topic)
	defer tc.DeleteResourceAndWait(topic)

	tc.Expect(topic.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(topic.Status.SizeInBytes).ToNot(BeNil())
	tc.Expect(*topic.Status.SizeInBytes).To(Equal(0))
}
