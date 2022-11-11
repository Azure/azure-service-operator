/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/
package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func Test_Multitenant_SingleOperator_NamespacedCredential(t *testing.T) {
	t.Parallel()

	secretNamespace := "secret-namespace"

	cfg, err := testcommon.ReadFromEnvironmentForTest()
	if err != nil {
		t.Fatal(err)
	}

	// TargetNamespaces is only used here for logger mapping in namespaces, and not multi-tenancy.
	cfg.TargetNamespaces = []string{secretNamespace}
	tc := globalTestContext.ForTestWithConfig(t, cfg)

	err = tc.CreateTestNamespaces(secretNamespace)
	tc.Expect(err).ToNot(HaveOccurred())

	secret := newCredentialSecret(arm.NamespacedSecretName, secretNamespace)
	tc.CreateResource(secret)

	rg := tc.NewTestResourceGroup()
	// to enforce namespaced secret on rg
	rg.Namespace = secretNamespace

	tc.CreateResourcesAndWait(rg)

	acct := newStorageAccount(tc, rg)
	acct.Namespace = secretNamespace

	tc.CreateResourceAndWaitForState(acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// Patch the account to skip reconcile
	old := acct.DeepCopy()
	acct.Annotations = make(map[string]string)
	acct.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	tc.Patch(old, acct)
	tc.DeleteResourceAndWait(acct)

	// Creating new storage account in namespace other than credential namespace
	// TODO: FIXME.. Tests get stuck here
	acct = newStorageAccount(tc, rg)
	tc.CreateResourceAndWait(acct)

	tc.DeleteResourcesAndWait(rg)

}

func Test_Multitenant_SingleOperator_PerResourceCredential(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	secret := newCredentialSecret("credential", tc.Namespace)
	tc.CreateResource(secret)

	nsName := types.NamespacedName{
		Namespace: secret.Namespace,
		Name:      secret.Name,
	}

	rg := tc.NewTestResourceGroup()
	rg.Annotations = map[string]string{arm.PerResourceSecretAnnotation: nsName.String()}

	tc.CreateResourcesAndWait(rg)
	tc.DeleteResourcesAndWait(rg)

}

func newCredentialSecret(name, namespaceName string) *v1.Secret {
	secretData := make(map[string][]byte)
	// TODO: find a way to provision creds into secret
	secretData[config.AzureClientIDVar] = []byte("")
	secretData[config.AzureClientSecretVar] = []byte("")
	secretData[config.TenantIDVar] = []byte("")
	secretData[config.SubscriptionIDVar] = []byte("")

	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespaceName,
		},
		Data: secretData,
	}

}
