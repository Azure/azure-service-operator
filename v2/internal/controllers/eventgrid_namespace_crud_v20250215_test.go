/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	eventgrid "github.com/Azure/azure-service-operator/v2/api/eventgrid/v20250215"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_EventGrid_Namespace_CRUD_20250215(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a namespace
	namespace := &eventgrid.Namespace{
		ObjectMeta: tc.MakeObjectMeta("namespace"),
		Spec: eventgrid.Namespace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Tags:     map[string]string{"cheese": "blue"},
			Sku: &eventgrid.NamespaceSku{
				Name:     to.Ptr(eventgrid.NamespaceSku_Name_Standard),
				Capacity: to.Ptr(1),
			},
		},
	}

	tc.CreateResourceAndWait(namespace)

	armId := *namespace.Status.Id
	tc.Expect(namespace.Status.MinimumTlsVersionAllowed).To(HaveValue(Equal(eventgrid.NamespaceProperties_MinimumTlsVersionAllowed_STATUS_12)))

	// Perform a simple patch.
	old := namespace.DeepCopy()
	namespace.Spec.Tags["cheese"] = "époisses"
	tc.PatchResourceAndWait(old, namespace)
	tc.Expect(namespace.Status.Tags).To(Equal(map[string]string{"cheese": "époisses"}))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Namespace_SecretsWrittenToSameKubeSecret",
			Test: func(tc *testcommon.KubePerTestContext) {
				Namespace_SecretsWrittenToSameKubeSecret_20250215(tc, namespace)
			},
		},
	)

	tc.DeleteResourceAndWait(namespace)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(eventgrid.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Namespace_SecretsWrittenToSameKubeSecret_20250215(tc *testcommon.KubePerTestContext, namespace *eventgrid.Namespace) {
	old := namespace.DeepCopy()
	namespaceSecret := "namespacekeys"
	namespace.Spec.OperatorSpec = &eventgrid.NamespaceOperatorSpec{
		Secrets: &eventgrid.NamespaceOperatorSecrets{
			Key1: &genruntime.SecretDestination{Name: namespaceSecret, Key: "key1"},
			Key2: &genruntime.SecretDestination{Name: namespaceSecret, Key: "key2"},
		},
	}
	tc.PatchResourceAndWait(old, namespace)

	tc.ExpectSecretHasKeys(namespaceSecret, "key1", "key2")
}
