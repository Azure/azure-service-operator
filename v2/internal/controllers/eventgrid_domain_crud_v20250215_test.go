/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v20250215"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_EventGrid_Domain_CRUD_20250215(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a domain
	domain := &eventgrid.Domain{
		ObjectMeta: tc.MakeObjectMeta("domain"),
		Spec: eventgrid.Domain_Spec{
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PublicNetworkAccess: to.Ptr(eventgrid.DomainProperties_PublicNetworkAccess_Enabled),
			DisableLocalAuth:    to.Ptr(true),
		},
	}

	// Create a storage account to use as destination
	acctName := tc.NoSpaceNamer.GenerateName("dest")
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(acctName),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     to.Ptr(storage.StorageAccount_Kind_Spec_StorageV2),
			Sku:      &storage.Sku{Name: to.Ptr(storage.SkuName_Standard_LRS)},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: to.Ptr(storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot),
		},
	}

	tc.CreateResourcesAndWait(domain, acct)

	tc.Expect(domain.Status.DisableLocalAuth).To(HaveValue(BeTrue()))

	queueServices := &storage.StorageAccountsQueueService{
		ObjectMeta: tc.MakeObjectMeta("dest-queues"),
		Spec: storage.StorageAccountsQueueService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&queueServices.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	queue := &storage.StorageAccountsQueueServicesQueue{
		ObjectMeta: tc.MakeObjectMeta("dest-queue"),
		Spec: storage.StorageAccountsQueueServicesQueue_Spec{
			Owner: testcommon.AsOwner(queueServices),
		},
	}

	tc.CreateResourcesAndWait(queueServices, queue)
	armId := *domain.Status.Id
	acctReference := tc.MakeReferenceFromResource(acct)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CreateDomainTopicAndSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				DomainTopicAndSubscription_CRUD_20250215(tc, queue, domain, acctReference)
			},
		},
		testcommon.Subtest{
			Name: "CreateDomainSubscription",
			Test: func(tc *testcommon.KubePerTestContext) {
				DomainSubscription_CRUD_20250215(tc, queue, domain, acctReference)
			},
		},
	)

	tc.DeleteResourceAndWait(domain)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(eventgrid.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func DomainTopicAndSubscription_CRUD_20250215(tc *testcommon.KubePerTestContext, queue *storage.StorageAccountsQueueServicesQueue, domain *eventgrid.Domain, acctReference *genruntime.ResourceReference) {
	topic := &eventgrid.DomainsTopic{
		ObjectMeta: tc.MakeObjectMeta("topic"),
		Spec: eventgrid.DomainsTopic_Spec{
			Owner: testcommon.AsOwner(domain),
		},
	}

	tc.CreateResourceAndWait(topic)
	// don’t bother deleting; deleting domain will clean up

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
	// don’t bother deleting
}

func DomainSubscription_CRUD_20250215(tc *testcommon.KubePerTestContext, queue *storage.StorageAccountsQueueServicesQueue, domain *eventgrid.Domain, acctReference *genruntime.ResourceReference) {
	subscription := &eventgrid.EventSubscription{
		ObjectMeta: tc.MakeObjectMeta("sub"),
		Spec: eventgrid.EventSubscription_Spec{
			Owner: tc.AsExtensionOwner(domain),
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
	// don’t bother deleting
}
