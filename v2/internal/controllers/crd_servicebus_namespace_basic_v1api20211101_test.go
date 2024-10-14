/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ServiceBus_Namespace_Basic_v1api20211101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sku := servicebus.SBSku_Name_Basic
	namespace := NewServiceBusNamespace_v1api20211101(tc, rg, sku)

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Queue CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_Queue_v1api20211101_CRUD(tc, namespace)
			},
		},
	)

	// This must run at the end as it modifies the namespace
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Namespace secrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_Namespace_Secrets_v1api20211101(tc, namespace)
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(servicebus.APIVersion_Value))
}

func ServiceBus_Queue_v1api20211101_CRUD(tc *testcommon.KubePerTestContext, sbNamespace client.Object) {
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

func ServiceBus_Namespace_Secrets_v1api20211101(tc *testcommon.KubePerTestContext, sbNamespace client.Object) {
	namespace := sbNamespace.(*servicebus.Namespace)
	secretName := "namespace-secrets-20211101"

	old := namespace.DeepCopy()

	if namespace.Spec.OperatorSpec == nil {
		namespace.Spec.OperatorSpec = &servicebus.NamespaceOperatorSpec{}
	}

	namespace.Spec.OperatorSpec.Secrets = &servicebus.NamespaceOperatorSecrets{
		Endpoint: &genruntime.SecretDestination{
			Name: secretName,
			Key:  "Endpoint",
		},
		PrimaryKey: &genruntime.SecretDestination{
			Name: secretName,
			Key:  "PrimaryKey",
		},
		PrimaryConnectionString: &genruntime.SecretDestination{
			Name: secretName,
			Key:  "PrimaryConnectionString",
		},
		SecondaryKey: &genruntime.SecretDestination{
			Name: secretName,
			Key:  "SecondaryKey",
		},
		SecondaryConnectionString: &genruntime.SecretDestination{
			Name: secretName,
			Key:  "SecondaryConnectionString",
		},
	}

	tc.PatchResourceAndWait(old, namespace)

	tc.UpdateResource(namespace)
	tc.ExpectSecretHasKeys(secretName, "Endpoint", "PrimaryKey", "PrimaryConnectionString", "SecondaryKey", "SecondaryConnectionString")
}
