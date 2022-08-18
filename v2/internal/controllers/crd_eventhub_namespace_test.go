/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	eventhub "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_EventHub_Namespace_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	skuTier := eventhub.SkuTier_Standard
	skuName := eventhub.SkuName_Standard
	namespace := &eventhub.Namespace{
		ObjectMeta: tc.MakeObjectMeta("namespace"),
		Spec: eventhub.Namespaces_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &eventhub.Sku{
				Name: &skuName,
				Tier: &skuTier,
			},
			IsAutoInflateEnabled:   to.BoolPtr(true),
			MaximumThroughputUnits: to.IntPtr(1),
		},
	}

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	tc.Expect(namespace.Status.MaximumThroughputUnits).To(Equal(to.IntPtr(1)))
	armId := *namespace.Status.Id

	// Perform a simple patch
	old := namespace.DeepCopy()
	namespace.Spec.MaximumThroughputUnits = to.IntPtr(2)
	tc.PatchResourceAndWait(old, namespace)
	tc.Expect(namespace.Status.MaximumThroughputUnits).To(Equal(to.IntPtr(2)))

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
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(eventhub.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func EventHub_CRUD(tc *testcommon.KubePerTestContext, namespace client.Object) {
	eh := &eventhub.NamespacesEventhub{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.NamespacesEventhubs_Spec{
			Owner:                  testcommon.AsOwner(namespace),
			MessageRetentionInDays: to.IntPtr(7),
			PartitionCount:         to.IntPtr(1),
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
	eh.Spec.MessageRetentionInDays = to.IntPtr(3)
	tc.PatchResourceAndWait(old, eh)
	tc.Expect(eh.Status.MessageRetentionInDays).To(Equal(to.IntPtr(3)))
}

func Namespace_AuthorizationRules_CRUD(tc *testcommon.KubePerTestContext, namespace client.Object) {
	rule := &eventhub.NamespacesAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.NamespacesAuthorizationRules_Spec{
			Owner: testcommon.AsOwner(namespace),
			Rights: []eventhub.AuthorizationRulePropertiesRights{
				eventhub.AuthorizationRulePropertiesRights_Listen,
				eventhub.AuthorizationRulePropertiesRights_Send,
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
		Spec: eventhub.NamespacesEventhubsAuthorizationRules_Spec{
			Owner: testcommon.AsOwner(eh),
			Rights: []eventhub.AuthorizationRulePropertiesRights{
				eventhub.AuthorizationRulePropertiesRights_Listen,
				eventhub.AuthorizationRulePropertiesRights_Send,
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
	userMetadata := to.StringPtr("This is some fun metadata")
	consumerGroup := &eventhub.NamespacesEventhubsConsumerGroup{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.NamespacesEventhubsConsumergroups_Spec{
			Owner:        testcommon.AsOwner(eh),
			UserMetadata: to.StringPtr("This is some fun metadata"),
		},
	}

	tc.CreateResourceAndWait(consumerGroup)
	defer tc.DeleteResourceAndWait(consumerGroup)

	tc.Expect(consumerGroup.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(consumerGroup.Status.UserMetadata).To(Equal(userMetadata))
}
