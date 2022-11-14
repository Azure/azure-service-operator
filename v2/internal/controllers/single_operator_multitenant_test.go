/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/
package controllers_test

import (
	"testing"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func Test_Multitenant_SingleOperator_NamespacedCredential(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	secret := newCredentialSecret(tc.AzureSubscription, arm.NamespacedSecretName, tc.Namespace)
	tc.CreateResource(secret)

	rg := tc.NewTestResourceGroup()

	tc.CreateResourcesAndWait(rg)

	acct := newStorageAccount(tc, rg)

	// Creating new storage account with restricted permissions namespaced secret should fail.
	tc.CreateResourceAndWaitForState(acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// Deleting the credential would default to applying the global credential with all permissions
	tc.DeleteResource(secret)

	tc.Eventually(acct).Should(tc.Match.BeProvisioned(0))

	objKey := client.ObjectKeyFromObject(acct)
	tc.GetResource(objKey, acct)

	resID := genruntime.GetResourceIDOrDefault(acct)

	// Make sure the StorageAccount is created successfully in Azure.
	exists, _, err := tc.AzureClient.HeadByID(tc.Ctx, resID, string(storage.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	tc.DeleteResourcesAndWait(acct, rg)

}

func Test_Multitenant_SingleOperator_PerResourceCredential(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	secret := newCredentialSecret(tc.AzureSubscription, "credential", tc.Namespace)
	tc.CreateResource(secret)

	nsName := types.NamespacedName{
		Namespace: secret.Namespace,
		Name:      secret.Name,
	}

	rg := tc.CreateTestResourceGroupAndWait()

	acct := newStorageAccount(tc, rg)
	acct.Annotations = map[string]string{arm.PerResourceSecretAnnotation: nsName.String()}

	// Creating new storage account in with restricted permissions per resource secret should fail.
	tc.CreateResourceAndWaitForState(acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// Deleting the per-resource credential annotation would default to applying the global credential with all permissions
	old := acct.DeepCopy()
	delete(acct.Annotations, arm.PerResourceSecretAnnotation)
	tc.Patch(old, acct)

	tc.Eventually(acct).Should(tc.Match.BeProvisioned(0))

	objKey := client.ObjectKeyFromObject(acct)
	tc.GetResource(objKey, acct)

	resID := genruntime.GetResourceIDOrDefault(acct)

	// Make sure the StorageAccount is created successfully in Azure.
	exists, _, err := tc.AzureClient.HeadByID(tc.Ctx, resID, string(storage.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	tc.DeleteResourcesAndWait(acct, rg)

}

func newCredentialSecret(subscriptionID string, name, namespaceName string) *v1.Secret {
	secretData := make(map[string][]byte)
	// TODO: find a way to provision creds into secret
	secretData[config.AzureClientIDVar] = []byte("")
	secretData[config.AzureClientSecretVar] = []byte("")
	secretData[config.TenantIDVar] = []byte("")
	secretData[config.SubscriptionIDVar] = []byte(subscriptionID)

	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespaceName,
		},
		Data: secretData,
	}

}
