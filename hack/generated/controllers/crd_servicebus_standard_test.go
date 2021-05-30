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

	servicebus "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.servicebus/v1alpha1api20180101preview"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_ServiceBus_Standard_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	rg, err := testContext.CreateNewTestResourceGroup(testcommon.WaitForCreation)
	g.Expect(err).ToNot(HaveOccurred())

	zoneRedundant := false
	namespace := &servicebus.Namespace{
		ObjectMeta: testContext.MakeObjectMetaWithName(testContext.Namer.GenerateName("sbstandard")),
		Spec: servicebus.Namespaces_Spec{
			Location: testContext.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &servicebus.SBSku{
				Name: servicebus.SBSkuNameStandard,
			},
			Properties: servicebus.SBNamespaceProperties{
				ZoneRedundant: &zoneRedundant,
			},
		},
	}

	// Create
	g.Expect(testContext.KubeClient.Create(ctx, namespace)).To(Succeed())
	g.Eventually(namespace, remainingTime(t)).Should(testContext.Match.BeProvisioned(ctx))

	g.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	RunParallelSubtests(t,
		subtest{
			name: "Queue CRUD",
			test: func(t *testing.T) { ServiceBus_Queue_CRUD(t, testContext, namespace.ObjectMeta) },
		},
		subtest{
			name: "Topic CRUD",
			test: func(t *testing.T) { ServiceBus_Topic_CRUD(t, testContext, namespace.ObjectMeta) },
		},
	)

	// Delete
	g.Expect(testContext.KubeClient.Delete(ctx, namespace)).To(Succeed())
	g.Eventually(namespace, remainingTime(t)).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := testContext.AzureClient.HeadResource(ctx, armId, "2018-01-01-preview")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}

// Topics can only be created in Standard or Premium SKUs
func ServiceBus_Topic_CRUD(t *testing.T, testContext testcommon.KubePerTestContext, sbNamespace metav1.ObjectMeta) {
	ctx := context.Background()
	g := NewGomegaWithT(t)

	topic := &servicebus.NamespacesTopic{
		ObjectMeta: testContext.MakeObjectMeta("topic"),
		Spec: servicebus.NamespacesTopics_Spec{
			Location: &testContext.AzureRegion,
			Owner:    testcommon.AsOwner(sbNamespace),
		},
	}

	// Create
	g.Expect(testContext.KubeClient.Create(ctx, topic)).To(Succeed())
	g.Eventually(topic, remainingTime(t)).Should(testContext.Match.BeProvisioned(ctx))

	g.Expect(topic.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	g.Expect(topic.Status.Properties.SizeInBytes).ToNot(BeNil())
	g.Expect(*topic.Status.Properties.SizeInBytes).To(Equal(0))

	g.Expect(testContext.KubeClient.Delete(ctx, topic)).To(Succeed())
	g.Eventually(topic, remainingTime(t)).Should(testContext.Match.BeDeleted(ctx))
}
