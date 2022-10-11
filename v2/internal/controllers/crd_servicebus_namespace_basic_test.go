/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ServiceBus_Namespace_Basic_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sku := servicebus.SBSku_Name_Basic
	namespace := NewServiceBusNamespace(tc, rg, sku)

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
		testcommon.Subtest{
			Name: "Namespace secrets",
			Test: func(tc *testcommon.KubePerTestContext) { ServiceBus_Namespace_Secrets(tc, namespace) },
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(servicebus.APIVersion_Value))
}

func ServiceBus_Queue_CRUD(tc *testcommon.KubePerTestContext, sbNamespace client.Object) {
	queue := &servicebus.NamespacesQueue{
		ObjectMeta: tc.MakeObjectMeta("queue"),
		Spec: servicebus.Namespaces_Queue_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(sbNamespace),
		},
	}

	tc.CreateResourceAndWait(queue)
	defer tc.DeleteResourceAndWait(queue)

	tc.Expect(queue.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(queue.Status.SizeInBytes).ToNot(BeNil())
	tc.Expect(*queue.Status.SizeInBytes).To(Equal(0))
}

func ServiceBus_Namespace_Secrets(tc *testcommon.KubePerTestContext, namespace *servicebus.Namespace) {
	old := namespace.DeepCopy()
	secret := "s1"
	namespace.Spec.OperatorSpec = &servicebus.NamespaceOperatorSpec{
		Secrets: &servicebus.NamespaceOperatorSecrets{
			Endpoint: &genruntime.SecretDestination{
				Name: secret,
				Key:  "endpoint",
			},
		},
	}
	tc.PatchResourceAndWait(old, namespace)

	tc.ExpectSecretHasKeys(secret, "endpoint")
}
