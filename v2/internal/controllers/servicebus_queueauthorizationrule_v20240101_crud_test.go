/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	servicebus "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20240101"
	servicebusv20240101 "github.com/Azure/azure-service-operator/v2/api/servicebus/v20240101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ServiceBus_Queue_AuthorizationRule_v20240101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sku := servicebus.SBSku_Name_Standard
	namespace := NewServiceBusNamespace_v1api20240101(tc, rg, sku)

	tc.CreateResourceAndWait(namespace)

	tc.Expect(namespace.Status.Id).ToNot(BeNil())

	queue := &servicebus.NamespacesQueue{
		ObjectMeta: tc.MakeObjectMeta("queue"),
		Spec: servicebus.NamespacesQueue_Spec{
			Owner: testcommon.AsOwner(namespace),
		},
	}

	tc.CreateResourceAndWait(queue)

	tc.Expect(queue.Status.Id).ToNot(BeNil())

	rule := &servicebusv20240101.QueueAuthorizationRule{
		ObjectMeta: tc.MakeObjectMeta("rule"),
		Spec: servicebusv20240101.QueueAuthorizationRule_Spec{
			Owner: testcommon.AsOwner(queue),
			Rights: []servicebusv20240101.QueueAuthorizationRuleRights_Spec{
				servicebusv20240101.QueueAuthorizationRuleRights_Spec_Listen,
				servicebusv20240101.QueueAuthorizationRuleRights_Spec_Send,
			},
		},
	}

	tc.CreateResourceAndWait(rule)

	tc.Expect(rule.Status.Rights).To(HaveLen(2))

	// Test secret export
	old := rule.DeepCopy()
	ruleKeysSecret := "queuerulekeyssecret"
	rule.Spec.OperatorSpec = &servicebusv20240101.QueueAuthorizationRuleOperatorSpec{
		Secrets: &servicebusv20240101.QueueAuthorizationRuleOperatorSecrets{
			PrimaryKey:                &genruntime.SecretDestination{Name: ruleKeysSecret, Key: "primaryKey"},
			SecondaryKey:              &genruntime.SecretDestination{Name: ruleKeysSecret, Key: "secondaryKey"},
			PrimaryConnectionString:   &genruntime.SecretDestination{Name: ruleKeysSecret, Key: "primaryConnectionString"},
			SecondaryConnectionString: &genruntime.SecretDestination{Name: ruleKeysSecret, Key: "secondaryConnectionString"},
		},
	}
	tc.PatchResourceAndWait(old, rule)

	tc.ExpectSecretHasKeys(
		ruleKeysSecret,
		"primaryKey",
		"secondaryKey",
		"primaryConnectionString",
		"secondaryConnectionString",
	)

	ruleArmId := *rule.Status.Id
	tc.DeleteResourceAndWait(rule)
	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(ruleArmId, string(servicebusv20240101.APIVersion_Value))

	queueArmId := *queue.Status.Id
	tc.DeleteResourceAndWait(queue)
	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(queueArmId, string(servicebus.APIVersion_Value))

	namespaceArmId := *namespace.Status.Id
	tc.DeleteResourceAndWait(namespace)
	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(namespaceArmId, string(servicebus.APIVersion_Value))
}
