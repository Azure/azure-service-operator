/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_EventGrid_Topic(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a topic
	topic := &eventgrid.Topic{
		ObjectMeta: tc.MakeObjectMeta("topic"),
		Spec: eventgrid.Topic_Spec{
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
		testcommon.Subtest{
			Name: "Topic_SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				Topic_SecretsWrittenToSameKubeSecret(tc, topic)
			},
		},
	)

	tc.DeleteResourceAndWait(topic)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(eventgrid.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Topic_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, topic *eventgrid.Topic) {
	old := topic.DeepCopy()
	topicSecret := "topickeys"
	topicConfigMap := "topicconfigmap"
	topic.Spec.OperatorSpec = &eventgrid.TopicOperatorSpec{
		Secrets: &eventgrid.TopicOperatorSecrets{
			Key1: &genruntime.SecretDestination{Name: topicSecret, Key: "key1"},
			Key2: &genruntime.SecretDestination{Name: topicSecret, Key: "key2"},
		},
		ConfigMaps: &eventgrid.TopicOperatorConfigMaps{
			Endpoint: &genruntime.ConfigMapDestination{
				Name: topicConfigMap,
				Key:  "endpoint",
			},
		},
	}
	tc.PatchResourceAndWait(old, topic)

	tc.ExpectSecretHasKeys(topicSecret, "key1", "key2")
	tc.ExpectConfigMapHasKeys(topicConfigMap, "endpoint")
}

func Topic_Subscription_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, topic *eventgrid.Topic) {
	sku := storage.SkuName_Standard_LRS
	acctName := tc.NoSpaceNamer.GenerateName("stor")
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(acctName),
		Spec: storage.StorageAccount_Spec{
			Owner:      testcommon.AsOwner(rg),
			Location:   tc.AzureRegion,
			Kind:       to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			AccessTier: to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
			Sku:        &storage.Sku{Name: &sku},
		},
	}

	tc.CreateResourceAndWait(acct)

	queueService := &storage.StorageAccountsQueueService{
		ObjectMeta: tc.MakeObjectMeta("qservice"),
		Spec: storage.StorageAccounts_QueueService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	tc.CreateResourceAndWait(queueService)

	queue := &storage.StorageAccountsQueueServicesQueue{
		ObjectMeta: tc.MakeObjectMeta("queue"),
		Spec: storage.StorageAccounts_QueueServices_Queue_Spec{
			Owner: testcommon.AsOwner(queueService),
		},
	}

	tc.CreateResourceAndWait(queue)

	acctReference := tc.MakeReferenceFromResource(acct)
	subscription := &eventgrid.EventSubscription{
		ObjectMeta: tc.MakeObjectMeta("sub"),
		Spec: eventgrid.EventSubscription_Spec{
			Owner: tc.AsExtensionOwner(topic),
			Destination: &eventgrid.EventSubscriptionDestination{
				StorageQueue: &eventgrid.StorageQueueEventSubscriptionDestination{
					EndpointType:      to.Ptr(eventgrid.StorageQueueEventSubscriptionDestination_EndpointType_StorageQueue),
					ResourceReference: acctReference,
					QueueName:         &queue.Name,
				},
			},
		},
	}

	tc.CreateResourceAndWait(subscription)
}
