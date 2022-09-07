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

func Test_ServiceBus_Namespace_Standard_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	zoneRedundant := false
	sku := servicebus.SBSku_Name_Standard
	namespace := &servicebus.Namespace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("sbstandard")),
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
			Test: func(tc *testcommon.KubePerTestContext) { ServiceBus_Queue_CRUD(tc, namespace) },
		},
		testcommon.Subtest{
			Name: "Topic CRUD",
			Test: func(tc *testcommon.KubePerTestContext) { ServiceBus_Topic_CRUD(tc, namespace) },
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(servicebus.APIVersion_Value))
}

// Topics can only be created in Standard or Premium SKUs
func ServiceBus_Topic_CRUD(tc *testcommon.KubePerTestContext, sbNamespace client.Object) {
	topic := &servicebus.NamespacesTopic{
		ObjectMeta: tc.MakeObjectMeta("topic"),
		Spec: servicebus.Namespaces_Topic_Spec{
			Location: tc.AzureRegion,
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
