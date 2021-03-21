/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	servicebus "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.servicebus/v1alpha1api20180101preview"
	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

func Test_ServiceBus_Namespace_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	rg, err := testContext.CreateNewTestResourceGroup(testcommon.WaitForCreation)
	g.Expect(err).ToNot(HaveOccurred())

	zoneRedundant := false
	namespace := &servicebus.Namespace{
		ObjectMeta: testContext.MakeObjectMetaWithName(testContext.Namer.GenerateName("sbnamespace")),
		Spec: servicebus.Namespaces_Spec{
			Location: testContext.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &servicebus.SBSku{
				Name: servicebus.SBSkuNameBasic,
			},
			Properties: servicebus.SBNamespaceProperties{
				ZoneRedundant: &zoneRedundant,
			},
		},
	}

	err = testContext.KubeClient.Create(ctx, namespace)
	g.Expect(err).ToNot(HaveOccurred())

	// It should be created in Kubernetes
	g.Eventually(namespace).Should(testContext.Match.BeProvisioned(ctx))

	// Run sub-tests
	t.Run("Queue CRUD", func(t *testing.T) {
		ServiceBus_Queue_CRUD(t, testContext, namespace.ObjectMeta)
	})

	g.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	// Delete
	err = testContext.KubeClient.Delete(ctx, namespace)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(namespace).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := testContext.AzureClient.HeadResource(ctx, armId, "2018-01-01-preview")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}

func ServiceBus_Queue_CRUD(t *testing.T, testContext testcommon.KubePerTestContext, sbNamespace metav1.ObjectMeta) {
	ctx := context.Background()

	g := NewGomegaWithT(t)

	queue := &servicebus.NamespacesQueue{
		ObjectMeta: testContext.MakeObjectMeta("queue"),
		Spec: servicebus.NamespacesQueues_Spec{
			Location: &testContext.AzureRegion,
			Owner:    testcommon.AsOwner(sbNamespace),
		},
	}

	// Create
	err := testContext.KubeClient.Create(ctx, queue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(queue).Should(testContext.Match.BeProvisioned(ctx))

	g.Expect(queue.Status.Id).ToNot(BeNil())

	// Just a basic assertion on a property
	g.Expect(queue.Status.Properties.SizeInBytes).ToNot(BeNil())
	g.Expect(*queue.Status.Properties.SizeInBytes).To(Equal(0))

	err = testContext.KubeClient.Delete(ctx, queue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(queue).Should(testContext.Match.BeDeleted(ctx))
}
