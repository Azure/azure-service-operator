/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ServiceBus_Namespace_Standard_v1api20221001preview_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sku := servicebus.SBSku_Name_Standard
	namespace := NewServiceBusNamespace_v1api20221001preview(tc, rg, sku)

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())
	armId := *namespace.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Queue CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_Queue_v1api20221001preview_CRUD(tc, namespace)
			},
		},
		testcommon.Subtest{
			Name: "Topic CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_Topic_v1api20221001preview_CRUD(tc, namespace)
			},
		},
		testcommon.Subtest{
			Name: "AuthorizationRule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_AuthorizationRule_v1api20221001preview_CRUD(tc, namespace)
			},
		},
		testcommon.Subtest{
			Name: "NamespaceSecrets CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_Namespace_Secrets_v1api20221001preview(tc, namespace)
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(servicebus.APIVersion_Value))
}

func NewServiceBusNamespace_v1api20221001preview(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, sku servicebus.SBSku_Name) *servicebus.Namespace {
	zoneRedundant := false
	namespace := &servicebus.Namespace{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("namespace")),
		Spec: servicebus.Namespace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &servicebus.SBSku{
				Name: &sku,
			},
			ZoneRedundant: &zoneRedundant,
		},
	}
	return namespace
}

func ServiceBus_AuthorizationRule_v1api20221001preview_CRUD(tc *testcommon.KubePerTestContext, sbNamespace client.Object) {
	secretName := "rule-secrets"

	rule := &servicebus.NamespacesAuthorizationRule{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("rule")),
		Spec: servicebus.Namespaces_AuthorizationRule_Spec{
			Owner: testcommon.AsOwner(sbNamespace),
			Rights: []servicebus.Namespaces_AuthorizationRule_Properties_Rights_Spec{
				servicebus.Namespaces_AuthorizationRule_Properties_Rights_Spec_Listen,
				servicebus.Namespaces_AuthorizationRule_Properties_Rights_Spec_Send,
			},
			OperatorSpec: &servicebus.NamespacesAuthorizationRuleOperatorSpec{
				Secrets: &servicebus.NamespacesAuthorizationRuleOperatorSecrets{
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
				},
			},
		},
	}

	tc.CreateResourceAndWait(rule)
	defer tc.DeleteResourceAndWait(rule)

	tc.Expect(rule.Status.Rights).To(HaveLen(2))

	tc.ExpectSecretHasKeys(secretName, "PrimaryKey", "PrimaryConnectionString", "SecondaryKey", "SecondaryConnectionString")
}

// Topics can only be created in Standard or Premium SKUs
func ServiceBus_Topic_v1api20221001preview_CRUD(tc *testcommon.KubePerTestContext, sbNamespace client.Object) {
	topic := &servicebus.NamespacesTopic{
		ObjectMeta: tc.MakeObjectMeta("topic"),
		Spec: servicebus.Namespaces_Topic_Spec{
			Owner: testcommon.AsOwner(sbNamespace),
		},
	}

	tc.CreateResourceAndWait(topic)
	defer tc.DeleteResourceAndWait(topic)

	tc.Expect(topic.Status.Id).ToNot(BeNil())

	// a basic assertion on a property
	tc.Expect(topic.Status.SizeInBytes).ToNot(BeNil())
	tc.Expect(*topic.Status.SizeInBytes).To(Equal(0))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Subscription CRUD",
			Test: func(tc *testcommon.KubePerTestContext) { ServiceBus_Subscription_v1api20221001preview_CRUD(tc, topic) },
		},
	)
}

func ServiceBus_Subscription_v1api20221001preview_CRUD(tc *testcommon.KubePerTestContext, sbTopic client.Object) {
	subscription := &servicebus.NamespacesTopicsSubscription{
		ObjectMeta: tc.MakeObjectMeta("subscription"),
		Spec: servicebus.Namespaces_Topics_Subscription_Spec{
			Owner: testcommon.AsOwner(sbTopic),
		},
	}

	tc.CreateResourceAndWait(subscription)

	tc.Expect(subscription.Status.Id).ToNot(BeNil())
	armId := *subscription.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SubscriptionsRule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ServiceBus_Subscriptions_Rule_v1api20221001preview_CRUD(tc, subscription)
			},
		},
	)

	tc.DeleteResourceAndWait(subscription)
	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(servicebus.APIVersion_Value))
}

func ServiceBus_Subscriptions_Rule_v1api20221001preview_CRUD(tc *testcommon.KubePerTestContext, sbSubscription client.Object) {
	rule := &servicebus.NamespacesTopicsSubscriptionsRule{
		ObjectMeta: tc.MakeObjectMeta("subrule"),
		Spec: servicebus.Namespaces_Topics_Subscriptions_Rule_Spec{
			Owner: testcommon.AsOwner(sbSubscription),
		},
	}

	tc.CreateResourceAndWait(rule)

	tc.Expect(rule.Status.Id).ToNot(BeNil())
	armId := *rule.Status.Id

	tc.DeleteResourceAndWait(rule)
	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(servicebus.APIVersion_Value))
}
