/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v20250215"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_EventGrid_NamespaceTopic_CRUD_20250215(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a namespace
	namespace := &eventgrid.Namespace{
		ObjectMeta: tc.MakeObjectMeta("namespace"),
		Spec: eventgrid.Namespace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &eventgrid.NamespaceSku{
				Name:     to.Ptr(eventgrid.NamespaceSku_Name_Standard),
				Capacity: to.Ptr(1),
			},
		},
	}

	// Create a namespace topic
	namespaceTopic := &eventgrid.NamespaceTopic{
		ObjectMeta: tc.MakeObjectMeta("namespacetopic"),
		Spec: eventgrid.NamespaceTopic_Spec{
			Owner:                testcommon.AsOwner(namespace),
			EventRetentionInDays: to.Ptr(1),
			InputSchema:          to.Ptr(eventgrid.NamespaceTopicProperties_InputSchema_CloudEventSchemaV1_0),
		},
	}

	tc.CreateResourcesAndWait(namespace, namespaceTopic)

	armId := *namespaceTopic.Status.Id
	tc.Expect(namespaceTopic.Status.EventRetentionInDays).To(HaveValue(Equal(1)))
	tc.Expect(namespaceTopic.Status.InputSchema).To(HaveValue(Equal(eventgrid.NamespaceTopicProperties_InputSchema_STATUS_CloudEventSchemaV1_0)))

	// Perform a simple patch.
	old := namespaceTopic.DeepCopy()
	namespaceTopic.Spec.EventRetentionInDays = to.Ptr(2)
	tc.PatchResourceAndWait(old, namespaceTopic)
	tc.Expect(namespaceTopic.Status.EventRetentionInDays).To(HaveValue(Equal(2)))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "NamespaceTopic_SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				NamespaceTopic_SecretsWrittenToSameKubeSecret_20250215(tc, namespaceTopic)
			},
		},
	)

	tc.DeleteResourceAndWait(namespaceTopic)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(eventgrid.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func NamespaceTopic_SecretsWrittenToSameKubeSecret_20250215(tc *testcommon.KubePerTestContext, namespaceTopic *eventgrid.NamespaceTopic) {
	old := namespaceTopic.DeepCopy()
	namespaceTopicSecret := "namespacetopickeys"
	namespaceTopic.Spec.OperatorSpec = &eventgrid.NamespaceTopicOperatorSpec{
		Secrets: &eventgrid.NamespaceTopicOperatorSecrets{
			Key1: &genruntime.SecretDestination{Name: namespaceTopicSecret, Key: "key1"},
			Key2: &genruntime.SecretDestination{Name: namespaceTopicSecret, Key: "key2"},
		},
	}
	tc.PatchResourceAndWait(old, namespaceTopic)

	tc.ExpectSecretHasKeys(namespaceTopicSecret, "key1", "key2")
}
