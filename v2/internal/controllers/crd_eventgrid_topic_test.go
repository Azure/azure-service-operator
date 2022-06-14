/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_EventGrid_Topic(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a topic
	topic := &eventgrid.Topic{
		ObjectMeta: tc.MakeObjectMeta("topic"),
		Spec: eventgrid.Topics_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Tags:     map[string]string{"cheese": "blue"},
		},
	}

	tc.CreateResourceAndWait(topic)

	armId := *topic.Status.Id

	// Perform a simple patch.
	old := topic.DeepCopy()
	topic.Spec.Tags["cheese"] = "époisses"
	tc.PatchResourceAndWait(old, topic)
	tc.Expect(topic.Status.Tags).To(Equal(map[string]string{"cheese": "époisses"}))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CreateTopicSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				Topic_Subscription_CRUD(tc, rg, topic)
			},
		},
	)

	tc.DeleteResourceAndWait(topic)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(eventgrid.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Topic_Subscription_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, topic *eventgrid.Topic) {
	kind := storage.StorageAccountsSpecKindStorageV2
	sku := storage.SkuNameStandardLRS
	acctName := tc.NoSpaceNamer.GenerateName("stor")
	tier := storage.StorageAccountPropertiesCreateParametersAccessTierHot
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(acctName),
		Spec: storage.StorageAccounts_Spec{
			Owner:      testcommon.AsOwner(rg),
			Location:   tc.AzureRegion,
			Kind:       &kind,
			AccessTier: &tier,
			Sku:        &storage.Sku{Name: &sku},
		},
	}

	tc.CreateResourceAndWait(acct)

	queueService := &storage.StorageAccountsQueueService{
		ObjectMeta: tc.MakeObjectMeta("qservice"),
		Spec: storage.StorageAccountsQueueServices_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	tc.CreateResourceAndWait(queueService)

	queue := &storage.StorageAccountsQueueServicesQueue{
		ObjectMeta: tc.MakeObjectMeta("queue"),
		Spec: storage.StorageAccountsQueueServicesQueues_Spec{
			Owner: testcommon.AsOwner(queueService),
		},
	}

	tc.CreateResourceAndWait(queue)

	acctReference := tc.MakeReferenceFromResource(acct)

	endpointType := eventgrid.StorageQueueEventSubscriptionDestinationEndpointTypeStorageQueue
	subscription := &eventgrid.EventSubscription{
		ObjectMeta: tc.MakeObjectMeta("sub"),
		Spec: eventgrid.EventSubscriptions_Spec{
			Owner: tc.AsExtensionOwner(topic),
			Destination: &eventgrid.EventSubscriptionDestination{
				StorageQueue: &eventgrid.StorageQueueEventSubscriptionDestination{
					EndpointType: &endpointType,
					Properties: &eventgrid.StorageQueueEventSubscriptionDestinationProperties{
						ResourceReference: acctReference,
						QueueName:         &queue.Name,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(subscription)
}
