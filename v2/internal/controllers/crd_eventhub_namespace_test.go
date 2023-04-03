/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	eventhub "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_EventHub_Namespace_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	skuTier := eventhub.Sku_Tier_Standard
	skuName := eventhub.Sku_Name_Standard
	namespace := &eventhub.Namespace{
		ObjectMeta: tc.MakeObjectMeta("namespace"),
		Spec: eventhub.Namespace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &eventhub.Sku{
				Name: &skuName,
				Tier: &skuTier,
			},
			IsAutoInflateEnabled:   to.Ptr(true),
			MaximumThroughputUnits: to.Ptr(1),
		},
	}

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	tc.Expect(namespace.Status.MaximumThroughputUnits).To(Equal(to.Ptr(1)))
	armId := *namespace.Status.Id

	// Perform a simple patch
	old := namespace.DeepCopy()
	namespace.Spec.MaximumThroughputUnits = to.Ptr(2)
	tc.PatchResourceAndWait(old, namespace)
	tc.Expect(namespace.Status.MaximumThroughputUnits).To(Equal(to.Ptr(2)))

	// Run sub tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "EventHub CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHub_CRUD(tc, namespace)
			},
		},
		testcommon.Subtest{
			Name: "EventHub namespace auth rule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Namespace_AuthorizationRules_CRUD(tc, namespace)
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(eventhub.APIVersion_Value))
}

func EventHub_CRUD(tc *testcommon.KubePerTestContext, namespace client.Object) {
	eh := &eventhub.NamespacesEventhub{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.Namespaces_Eventhub_Spec{
			Owner:                  testcommon.AsOwner(namespace),
			MessageRetentionInDays: to.Ptr(7),
			PartitionCount:         to.Ptr(1),
		},
	}

	tc.CreateResourceAndWait(eh)
	defer tc.DeleteResourceAndWait(eh)

	tc.Expect(eh.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(eh.Status.MessageRetentionInDays).ToNot(BeNil())
	tc.Expect(*eh.Status.MessageRetentionInDays).To(Equal(7))

	// Run sub tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "EventHub auth rule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHub_AuthorizationRules_CRUD(tc, eh)
			},
		},
		testcommon.Subtest{
			Name: "EventHub consumer group CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHub_ConsumerGroup_CRUD(tc, eh)
			},
		},
	)

	// Perform a simple patch
	old := eh.DeepCopy()
	eh.Spec.MessageRetentionInDays = to.Ptr(3)
	tc.PatchResourceAndWait(old, eh)
	tc.Expect(eh.Status.MessageRetentionInDays).To(Equal(to.Ptr(3)))
}

func Namespace_AuthorizationRules_CRUD(tc *testcommon.KubePerTestContext, namespace client.Object) {
	rule := &eventhub.NamespacesAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.Namespaces_AuthorizationRule_Spec{
			Owner: testcommon.AsOwner(namespace),
			Rights: []eventhub.Namespaces_AuthorizationRule_Properties_Rights_Spec{
				eventhub.Namespaces_AuthorizationRule_Properties_Rights_Spec_Listen,
				eventhub.Namespaces_AuthorizationRule_Properties_Rights_Spec_Send,
			},
		},
	}

	tc.CreateResourceAndWait(rule)
	defer tc.DeleteResourceAndWait(rule)

	tc.Expect(rule.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(rule.Status.Rights).To(HaveLen(2))
}

func EventHub_AuthorizationRules_CRUD(tc *testcommon.KubePerTestContext, eh client.Object) {
	rule := &eventhub.NamespacesEventhubsAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.Namespaces_Eventhubs_AuthorizationRule_Spec{
			Owner: testcommon.AsOwner(eh),
			Rights: []eventhub.Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec{
				eventhub.Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_Listen,
				eventhub.Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_Send,
			},
		},
	}

	tc.CreateResourceAndWait(rule)
	defer tc.DeleteResourceAndWait(rule)

	tc.Expect(rule.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(rule.Status.Rights).To(HaveLen(2))
}

func EventHub_ConsumerGroup_CRUD(tc *testcommon.KubePerTestContext, eh client.Object) {
	userMetadata := to.Ptr("This is some fun metadata")
	consumerGroup := &eventhub.NamespacesEventhubsConsumerGroup{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.Namespaces_Eventhubs_Consumergroup_Spec{
			Owner:        testcommon.AsOwner(eh),
			UserMetadata: to.Ptr("This is some fun metadata"),
		},
	}

	tc.CreateResourceAndWait(consumerGroup)
	defer tc.DeleteResourceAndWait(consumerGroup)

	tc.Expect(consumerGroup.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(consumerGroup.Status.UserMetadata).To(Equal(userMetadata))
}
