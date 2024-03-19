/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	eventhub "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20240101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_EventHub_Namespace_20240101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	skuTier := eventhub.Sku_Tier_Standard
	skuName := eventhub.Sku_Name_Standard
	namespace := &eventhub.Namespace{
		ObjectMeta: tc.MakeObjectMeta("namespace"),
		Spec: eventhub.Namespace_Spec{
			// Fails with NoRegisteredProviderFound error in westus2
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

	// Run secret sub tests
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Namespace_SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				Namespace_20240101_SecretsWrittenToSameKubeSecret(tc, namespace)
			},
		},
	)

	// Run sub tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "EventHub CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHub_20240101_CRUD(tc, namespace)
			},
		},
		testcommon.Subtest{
			Name: "EventHub namespace auth rule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Namespace_AuthorizationRules_20240101_CRUD(tc, namespace)
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(eventhub.APIVersion_Value))
}

func Namespace_20240101_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, ns *eventhub.Namespace) {
	old := ns.DeepCopy()
	namespaceKeysSecret := "namespacekeys"
	ns.Spec.OperatorSpec = &eventhub.NamespaceOperatorSpec{
		Secrets: &eventhub.NamespaceOperatorSecrets{
			PrimaryKey:                &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primary-key"},
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primary-connection-string"},
			SecondaryKey:              &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondary-key"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondary-connection-string"},
		},
	}
	tc.PatchResourceAndWait(old, ns)

	tc.ExpectSecretHasKeys(namespaceKeysSecret, "primary-key", "primary-connection-string", "secondary-key", "secondary-connection-string")
}

func EventHub_20240101_CRUD(tc *testcommon.KubePerTestContext, namespace client.Object) {
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
				EventHub_AuthorizationRules_20240101_CRUD(tc, eh)
			},
		},
		testcommon.Subtest{
			Name: "EventHub consumer group CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHub_ConsumerGroup_20240101_CRUD(tc, eh)
			},
		},
	)

	// Perform a simple patch
	old := eh.DeepCopy()
	eh.Spec.MessageRetentionInDays = to.Ptr(3)
	tc.PatchResourceAndWait(old, eh)
	tc.Expect(eh.Status.MessageRetentionInDays).To(Equal(to.Ptr(3)))
}

func Namespace_AuthorizationRules_20240101_CRUD(tc *testcommon.KubePerTestContext, namespace client.Object) {
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

	// Run secret sub tests
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "NamespacesAuthorizationRule_SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				NamespacesAuthorizationRule_20240101_SecretsWrittenToSameKubeSecret(tc, rule)
			},
		},
	)

	// a basic assertion on a property
	tc.Expect(rule.Status.Rights).To(HaveLen(2))
}

func NamespacesAuthorizationRule_20240101_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, ns *eventhub.NamespacesAuthorizationRule) {
	old := ns.DeepCopy()
	namespaceKeysSecret := "namespaceauthrulekeys"
	ns.Spec.OperatorSpec = &eventhub.NamespacesAuthorizationRuleOperatorSpec{
		Secrets: &eventhub.NamespacesAuthorizationRuleOperatorSecrets{
			PrimaryKey:                &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primary-key"},
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primary-connection-string"},
			SecondaryKey:              &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondary-key"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondary-connection-string"},
		},
	}
	tc.PatchResourceAndWait(old, ns)

	tc.ExpectSecretHasKeys(namespaceKeysSecret, "primary-key", "primary-connection-string", "secondary-key", "secondary-connection-string")
}

func EventHub_AuthorizationRules_20240101_CRUD(tc *testcommon.KubePerTestContext, eh client.Object) {
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

	// Run secret sub tests
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "EventHubAuthorizationRule_SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHubAuthorizationRule_20240101_SecretsWrittenToSameKubeSecret(tc, rule)
			},
		},
	)

	// a basic assertion on a property
	tc.Expect(rule.Status.Rights).To(HaveLen(2))
}

func EventHubAuthorizationRule_20240101_SecretsWrittenToSameKubeSecret(tc *testcommon.KubePerTestContext, ns *eventhub.NamespacesEventhubsAuthorizationRule) {
	old := ns.DeepCopy()
	namespaceKeysSecret := "eventhubauthrulekeys"
	ns.Spec.OperatorSpec = &eventhub.NamespacesEventhubsAuthorizationRuleOperatorSpec{
		Secrets: &eventhub.NamespacesEventhubsAuthorizationRuleOperatorSecrets{
			PrimaryKey:                &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primary-key"},
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "primary-connection-string"},
			SecondaryKey:              &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondary-key"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: namespaceKeysSecret, Key: "secondary-connection-string"},
		},
	}
	tc.PatchResourceAndWait(old, ns)

	tc.ExpectSecretHasKeys(namespaceKeysSecret, "primary-key", "primary-connection-string", "secondary-key", "secondary-connection-string")
}

func EventHub_ConsumerGroup_20240101_CRUD(tc *testcommon.KubePerTestContext, eh client.Object) {
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
