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

func Test_ServiceBus_Standard_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateNewTestResourceGroupAndWait()

	zoneRedundant := false
	namespace := &servicebus.Namespace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sbstandard")),
		Spec: servicebus.Namespaces_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &servicebus.SBSku{
				Name: servicebus.SBSkuNameStandard,
			},
			Properties: servicebus.SBNamespaceProperties{
				ZoneRedundant: &zoneRedundant,
			},
		},
	}

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Queue CRUD",
			Test: func(t testcommon.KubePerTestContext) { ServiceBus_Queue_CRUD(tc, namespace.ObjectMeta) },
		},
		testcommon.Subtest{
			Name: "Topic CRUD",
			Test: func(t testcommon.KubePerTestContext) { ServiceBus_Topic_CRUD(tc, namespace.ObjectMeta) },
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, "2018-01-01-preview")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

// Topics can only be created in Standard or Premium SKUs
func ServiceBus_Topic_CRUD(tc testcommon.KubePerTestContext, sbNamespace metav1.ObjectMeta) {
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
	tc.Expect(topic.Status.Properties.SizeInBytes).ToNot(BeNil())
	tc.Expect(*topic.Status.Properties.SizeInBytes).To(Equal(0))
}
