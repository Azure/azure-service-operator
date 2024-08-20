/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/
package test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

const (
	// TODO: This assumes the test is running at a particular path... I _think_ this is safe?
	KindOIDCIssuerPath   = "../out/kind-identity/azure/saissuer.txt"
	CLusterOIDCIssuerVar = "CLUSTER_OIDC_ISSUER"
	FICSubject           = "system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default"
)

// These tests are focused on multitenancy in a real deployment, where different namespaces use different credentials

func Test_Multitenant_SingleOperator_NamespacedCredential(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Get the issuer we're going to use for Workload Identity
	issuer, err := getOIDCIssuer()
	tc.Expect(err).ToNot(HaveOccurred())

	namespaceName := "multitenant-namespaced-cred"
	tc.Expect(tc.CreateTestNamespace(namespaceName)).To(Succeed())
	defer tc.DeleteResource(&v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	})

	rg := tc.CreateTestResourceGroupAndWait()

	mi := newManagedIdentity(tc, rg)
	fic := newFederatedIdentityCredential(tc, mi, issuer)
	tc.CreateResourcesAndWait(mi, fic)

	// TODO: See https://github.com/Azure/azure-service-operator/issues/3439 for the issue tracking us enabling this more
	// TODO: easily in the operator itself (rather than reading from .Status in code)
	secret := creds.NewScopedManagedIdentitySecret(tc.AzureSubscription, tc.AzureTenant, to.Value(mi.Status.ClientId), identity.NamespacedSecretName, namespaceName)
	tc.CreateResource(secret)

	testResourceGroup := tc.NewTestResourceGroup()
	testResourceGroup.Namespace = namespaceName

	// Creating new rg with restricted permissions namespaced secret should fail.
	tc.CreateResourceAndWaitForState(testResourceGroup, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We should eventually see an authorization error. There may be some initial errors related to the FederatedIdentityCredential
	// not yet existing (AAD propagation can take a few seconds)
	tc.Eventually(func() string {
		tc.GetResource(types.NamespacedName{Namespace: testResourceGroup.Namespace, Name: testResourceGroup.Name}, testResourceGroup)
		return testResourceGroup.Status.Conditions[0].Message
	}).WithTimeout(1 * time.Minute).Should(ContainSubstring("does not have authorization to perform action"))

	// Deleting the credential would default to applying the global credential with all permissions
	tc.DeleteResource(secret)

	tc.Eventually(testResourceGroup).Should(tc.Match.BeProvisioned(0))

	objKey := client.ObjectKeyFromObject(testResourceGroup)
	tc.GetResource(objKey, testResourceGroup)

	resID := genruntime.GetResourceIDOrDefault(testResourceGroup)

	// Make sure the ResourceGroup is created successfully in Azure.
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, resID, string(resources.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	tc.DeleteResourcesAndWait(testResourceGroup)
}

func Test_Multitenant_SingleOperator_PerResourceCredential(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Get the issuer we're going to use for Workload Identity
	issuer, err := getOIDCIssuer()
	tc.Expect(err).ToNot(HaveOccurred())

	rg := tc.CreateTestResourceGroupAndWait()

	mi := newManagedIdentity(tc, rg)
	fic := newFederatedIdentityCredential(tc, mi, issuer)
	tc.CreateResourcesAndWait(mi, fic)

	// TODO: See https://github.com/Azure/azure-service-operator/issues/3439 for the issue tracking us enabling this more
	// TODO: easily in the operator itself (rather than reading from .Status in code)
	secret := creds.NewScopedManagedIdentitySecret(tc.AzureSubscription, tc.AzureTenant, to.Value(mi.Status.ClientId), "credential", tc.Namespace)
	tc.CreateResource(secret)

	acct := newStorageAccount(tc, rg)
	acct.Annotations = map[string]string{annotations.PerResourceSecret: secret.Name}

	// Creating new storage account in with restricted permissions per resource secret should fail.
	tc.CreateResourceAndWaitForState(acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We should eventually see an authorization error. There may be some initial errors related to the FederatedIdentityCredential
	// not yet existing (AAD propagation can take a few seconds)
	tc.Eventually(func() string {
		tc.GetResource(types.NamespacedName{Namespace: acct.Namespace, Name: acct.Name}, acct)
		return acct.Status.Conditions[0].Message
	}).WithTimeout(1 * time.Minute).Should(ContainSubstring("does not have authorization to perform action"))

	// Deleting the per-resource credential annotation would default to applying the global credential with all permissions
	old := acct.DeepCopy()
	delete(acct.Annotations, annotations.PerResourceSecret)
	tc.Patch(old, acct)

	tc.Eventually(acct).Should(tc.Match.BeProvisioned(0))

	objKey := client.ObjectKeyFromObject(acct)
	tc.GetResource(objKey, acct)
	tc.Expect(acct.Annotations).ToNot(HaveKey(annotations.PerResourceSecret))

	resID := genruntime.GetResourceIDOrDefault(acct)

	// Make sure the StorageAccount is created successfully in Azure.
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, resID, string(storage.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	tc.DeleteResourcesAndWait(acct, rg)
}

func Test_Multitenant_SingleOperator_PerResourceCredential_MatchSubscriptionWithOwner(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Get the issuer we're going to use for Workload Identity
	issuer, err := getOIDCIssuer()
	tc.Expect(err).ToNot(HaveOccurred())

	rg := tc.CreateTestResourceGroupAndWait()

	mi := newManagedIdentity(tc, rg)
	fic := newFederatedIdentityCredential(tc, mi, issuer)
	tc.CreateResourcesAndWait(mi, fic)

	// TODO: See https://github.com/Azure/azure-service-operator/issues/3439 for the issue tracking us enabling this more
	// TODO: easily in the operator itself (rather than reading from .Status in code)
	secret := creds.NewScopedManagedIdentitySecret(uuid.New().String(), tc.AzureTenant, to.Value(mi.Status.ClientId), "credential", tc.Namespace)
	tc.CreateResource(secret)

	acct := newStorageAccount(tc, rg)
	acct.Annotations = map[string]string{annotations.PerResourceSecret: secret.Name}

	// Create a new storage account with a distinct subscription; per resource secret should fail because it does not match
	tc.CreateResourceAndWaitForFailure(acct)
	tc.Expect(acct.Status.Conditions[0].Message).To(ContainSubstring("does not match parent subscription"))
	tc.Expect(acct.Status.Conditions[0].Severity).To(Equal(conditions.ConditionSeverityError))
}

// getOIDCIssuer gets the OIDC issuer for the cluster.
// There are two supported ways of informing the tests about the OIDC issuer.
// - Set the CLUSTER_OIDC_ISSUER environment variable
// - Run in a kind cluster created by ASO's automation, which writes the OIDC issuer to {{.KIND_WORKLOAD_IDENTITY_PATH}}/azure/saissuer.txt
func getOIDCIssuer() (string, error) {
	// If we wanted this to work more generically for any OIDC issuer, we could do something like so:
	// https://azure.github.io/azure-workload-identity/docs/installation/managed-clusters.html#steps-to-get-the-oidc-issuer-url-from-a-generic-managed-cluster
	// but for now we don't really need that flexibility.

	// if the KIND OIDC issuer file exists, it takes precedence:
	content, err := os.ReadFile(KindOIDCIssuerPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
	} else {
		return strings.TrimSpace(string(content)), nil
	}

	issuer := os.Getenv(CLusterOIDCIssuerVar)
	if issuer != "" {
		return issuer, nil
	}

	return "", errors.Errorf("could not determine cluster OIDC issuer either from %s or %s", KindOIDCIssuerPath, CLusterOIDCIssuerVar)
}

func newManagedIdentity(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *managedidentity.UserAssignedIdentity {
	return &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("mi")),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}
}

func newFederatedIdentityCredential(tc *testcommon.KubePerTestContext, umi *managedidentity.UserAssignedIdentity, issuer string) *managedidentity.FederatedIdentityCredential {
	return &managedidentity.FederatedIdentityCredential{
		ObjectMeta: tc.MakeObjectMetaWithName("fic"), // Safe to always use this name as it's per-mi
		Spec: managedidentity.UserAssignedIdentities_FederatedIdentityCredential_Spec{
			Owner: testcommon.AsOwner(umi),
			// For Workload Identity, Audiences should always be "api://AzureADTokenExchange"
			Audiences: []string{
				"api://AzureADTokenExchange",
			},
			Issuer:  to.Ptr(issuer),
			Subject: to.Ptr(FICSubject),
		},
	}
}

func newStorageAccount(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *storage.StorageAccount {
	// Create a storage account
	accessTier := storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot
	kind := storage.StorageAccount_Kind_Spec_StorageV2
	sku := storage.SkuName_Standard_LRS
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Sku: &storage.Sku{
				Name: &sku,
			},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
		},
	}
	return acct
}
