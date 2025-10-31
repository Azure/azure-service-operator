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
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_EventHub_Namespace_v20211101_CRUD(t *testing.T) {
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

	// Run secret sub tests
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Namespace_SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				Namespace_SecretsWrittenToSameKubeSecret_v20211101(tc, namespace)
			},
		},
	)

	// Run sub tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "EventHub CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHub_CRUD_v20211101(tc, namespace)
			},
		},
		testcommon.Subtest{
			Name: "EventHub namespace auth rule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Namespace_AuthorizationRules_CRUD_v20211101(tc, namespace)
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(eventhub.APIVersion_Value))
}

func Namespace_SecretsWrittenToSameKubeSecret_v20211101(tc *testcommon.KubePerTestContext, ns *eventhub.Namespace) {
	old := ns.DeepCopy()

	nsSecretName := "namespacesecret"
	ns.Spec.OperatorSpec = &eventhub.NamespaceOperatorSpec{
		Secrets: &eventhub.NamespaceOperatorSecrets{
			PrimaryKey:                &genruntime.SecretDestination{Name: nsSecretName, Key: "primary-key"},
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: nsSecretName, Key: "primary-connection-string"},
			SecondaryKey:              &genruntime.SecretDestination{Name: nsSecretName, Key: "secondary-key"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: nsSecretName, Key: "secondary-connection-string"},
		},
	}
	tc.PatchResourceAndWait(old, ns)

	tc.ExpectSecretHasKeys(nsSecretName, "primary-key", "primary-connection-string", "secondary-key", "secondary-connection-string")
}

func EventHub_CRUD_v20211101(tc *testcommon.KubePerTestContext, namespace client.Object) {
	eh := &eventhub.NamespacesEventhub{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.NamespacesEventhub_Spec{
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
				EventHub_AuthorizationRules_CRUD_v20211101(tc, eh)
			},
		},
		testcommon.Subtest{
			Name: "EventHub consumer group CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				EventHub_ConsumerGroup_CRUD_v20211101(tc, eh)
			},
		},
	)

	// Perform a simple patch
	old := eh.DeepCopy()
	eh.Spec.MessageRetentionInDays = to.Ptr(3)
	tc.PatchResourceAndWait(old, eh)
	tc.Expect(eh.Status.MessageRetentionInDays).To(Equal(to.Ptr(3)))
}

func Namespace_AuthorizationRules_CRUD_v20211101(tc *testcommon.KubePerTestContext, namespace client.Object) {
	rule := &eventhub.NamespacesAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.NamespacesAuthorizationRule_Spec{
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
				NamespacesAuthorizationRule_SecretsWrittenToSameKubeSecret_v20211101(tc, rule)
			},
		},
	)

	// a basic assertion on a property
	tc.Expect(rule.Status.Rights).To(HaveLen(2))
}

func NamespacesAuthorizationRule_SecretsWrittenToSameKubeSecret_v20211101(tc *testcommon.KubePerTestContext, ns *eventhub.NamespacesAuthorizationRule) {
	old := ns.DeepCopy()

	namespaceAuthRuleSecretName := "namespaceauthrulesecret"
	ns.Spec.OperatorSpec = &eventhub.NamespacesAuthorizationRuleOperatorSpec{
		Secrets: &eventhub.NamespacesAuthorizationRuleOperatorSecrets{
			PrimaryKey:                &genruntime.SecretDestination{Name: namespaceAuthRuleSecretName, Key: "primary-key"},
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: namespaceAuthRuleSecretName, Key: "primary-connection-string"},
			SecondaryKey:              &genruntime.SecretDestination{Name: namespaceAuthRuleSecretName, Key: "secondary-key"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: namespaceAuthRuleSecretName, Key: "secondary-connection-string"},
		},
	}
	tc.PatchResourceAndWait(old, ns)

	tc.ExpectSecretHasKeys(namespaceAuthRuleSecretName, "primary-key", "primary-connection-string", "secondary-key", "secondary-connection-string")
}

func EventHub_AuthorizationRules_CRUD_v20211101(tc *testcommon.KubePerTestContext, eh client.Object) {
	rule := &eventhub.NamespacesEventhubsAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.NamespacesEventhubsAuthorizationRule_Spec{
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
				EventHubAuthorizationRule_SecretsWrittenToSameKubeSecret_v20211101(tc, rule)
			},
		},
	)

	// a basic assertion on a property
	tc.Expect(rule.Status.Rights).To(HaveLen(2))
}

func EventHubAuthorizationRule_SecretsWrittenToSameKubeSecret_v20211101(tc *testcommon.KubePerTestContext, ns *eventhub.NamespacesEventhubsAuthorizationRule) {
	old := ns.DeepCopy()
	//nolint:gosec
	eventHubAuthRuleSecretName := "eventhubauthrulesecret"
	ns.Spec.OperatorSpec = &eventhub.NamespacesEventhubsAuthorizationRuleOperatorSpec{
		Secrets: &eventhub.NamespacesEventhubsAuthorizationRuleOperatorSecrets{
			PrimaryKey:                &genruntime.SecretDestination{Name: eventHubAuthRuleSecretName, Key: "primary-key"},
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: eventHubAuthRuleSecretName, Key: "primary-connection-string"},
			SecondaryKey:              &genruntime.SecretDestination{Name: eventHubAuthRuleSecretName, Key: "secondary-key"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: eventHubAuthRuleSecretName, Key: "secondary-connection-string"},
		},
	}
	tc.PatchResourceAndWait(old, ns)

	tc.ExpectSecretHasKeys(eventHubAuthRuleSecretName, "primary-key", "primary-connection-string", "secondary-key", "secondary-connection-string")
}

func EventHub_ConsumerGroup_CRUD_v20211101(tc *testcommon.KubePerTestContext, eh client.Object) {
	userMetadata := to.Ptr("This is some fun metadata")
	consumerGroup := &eventhub.NamespacesEventhubsConsumerGroup{
		ObjectMeta: tc.MakeObjectMeta("eventhub"),
		Spec: eventhub.NamespacesEventhubsConsumerGroup_Spec{
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
