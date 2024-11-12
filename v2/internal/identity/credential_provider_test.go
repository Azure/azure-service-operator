/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/common/config"
)

const (
	testPodNamespace   = "azureserviceoperator-system-test"
	testSubscriptionID = "00000011-1111-0011-1100-110000000000" // Arbitrary GUID that isn't all 0s
	fakeID             = "00000000-0000-0000-0000-000000000000"
)

type testCredentialProviderResources struct {
	Provider                    CredentialProvider
	kubeClient                  kubeclient.Client
	fakeProvider                CredentialProvider
	fakeTokenCredentialProvider *mockTokenCredentialProvider
}

func NewTestCredentialProvider(client kubeclient.Client) (CredentialProvider, error) {
	tokenCreds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	creds := NewDefaultCredential(tokenCreds, testPodNamespace, testSubscriptionID)
	return NewCredentialProvider(creds, client, nil), nil
}

func testCredentialProviderSetup() (*testCredentialProviderResources, error) {
	s := createTestScheme()

	client := NewFakeKubeClient(s)
	provider, err := NewTestCredentialProvider(client)
	if err != nil {
		return nil, err
	}

	fakeTokenCredentialProvider := &mockTokenCredentialProvider{}
	fakeProvider := NewCredentialProvider(
		nil,
		client,
		&CredentialProviderOptions{TokenProvider: fakeTokenCredentialProvider})

	return &testCredentialProviderResources{
		Provider:                    provider,
		kubeClient:                  client,
		fakeTokenCredentialProvider: fakeTokenCredentialProvider,
		fakeProvider:                fakeProvider,
	}, nil
}

func TestCredentialProvider_DefaultCredentialNotSet_ReturnsErrorWhenTryToUseGlobalCredential(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	s := createTestScheme()
	kubeClient := NewFakeKubeClient(s)

	providerWithNoDefaultCred := NewCredentialProvider(nil, kubeClient, nil)
	rg := newResourceGroup("")

	_, err := providerWithNoDefaultCred.GetCredential(ctx, rg)
	g.Expect(err).ToNot(BeNil())
}

func TestCredentialProvider_ResourceScopeCredentialAndNamespaceCredential_PrefersResourceScopedCredential(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup()
	g.Expect(err).ToNot(HaveOccurred())

	perResourceCredentialName := types.NamespacedName{
		Namespace: "test-namespace",
		Name:      "test-secret",
	}

	perResourceSecret := newSecret(perResourceCredentialName)

	err = res.kubeClient.Create(ctx, perResourceSecret)
	g.Expect(err).ToNot(HaveOccurred())

	namespacedSecretName := types.NamespacedName{
		Namespace: "test-namespace",
		Name:      NamespacedSecretName,
	}

	secret := newSecret(namespacedSecretName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	rg.Annotations = map[string]string{annotations.PerResourceSecret: perResourceCredentialName.Name}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.CredentialFrom()).To(BeEquivalentTo(perResourceCredentialName))
	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(fakeID))
}

func TestCredentialProvider_SecretDoesNotExist_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-namespace",
		Name:      "test-secret",
	}

	rg := newResourceGroup("test-namespace")
	rg.Annotations = map[string]string{annotations.PerResourceSecret: credentialNamespacedName.Name}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).To(HaveOccurred())
	g.Expect(cred).To(BeNil())
}

func TestCredentialProvider_NamespaceCredential_IsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-secret",
		Name:      NamespacedSecretName,
	}

	secret := newSecret(credentialNamespacedName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup(credentialNamespacedName.Namespace)
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.CredentialFrom()).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(fakeID))
}

func TestCredentialProvider_GlobalCredential_IsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup()
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(cred.CredentialFrom()).To(BeEquivalentTo(types.NamespacedName{Namespace: testPodNamespace, Name: globalCredentialSecretName}))
}

func TestCredentialProvider_ServicePrincipalCredential_IsConfiguredCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup()
	g.Expect(err).ToNot(HaveOccurred())

	clientID := uuid.New().String()
	tenantID := uuid.New().String()
	clientSecret := uuid.New().String()

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      NamespacedSecretName,
		},
		Data: map[string][]byte{
			config.AzureSubscriptionID: []byte(testSubscriptionID),
			config.AzureClientID:       []byte(clientID),
			config.AzureTenantID:       []byte(tenantID),
			config.AzureClientSecret:   []byte(clientSecret),
		},
	}
	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.fakeProvider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.ClientSecret).To(Equal(clientSecret))
}

func TestCredentialProvider_CertificateCredential_IsConfiguredCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup()
	g.Expect(err).ToNot(HaveOccurred())

	clientID := uuid.New().String()
	tenantID := uuid.New().String()
	cert := uuid.New().String()
	certPassword := uuid.New().String()

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      NamespacedSecretName,
		},
		Data: map[string][]byte{
			config.AzureSubscriptionID:            []byte(testSubscriptionID),
			config.AzureClientID:                  []byte(clientID),
			config.AzureTenantID:                  []byte(tenantID),
			config.AzureClientCertificate:         []byte(cert),
			config.AzureClientCertificatePassword: []byte(certPassword),
		},
	}

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.fakeProvider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.ClientCertificate).To(Equal([]byte(cert)))
	g.Expect(res.fakeTokenCredentialProvider.Password).To(Equal([]byte(certPassword)))
}

func TestCredentialProvider_WorkloadIdentityCredential_IsConfiguredCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup()
	g.Expect(err).ToNot(HaveOccurred())

	clientID := uuid.New().String()
	tenantID := uuid.New().String()

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      NamespacedSecretName,
		},
		Data: map[string][]byte{
			config.AzureSubscriptionID: []byte(testSubscriptionID),
			config.AzureClientID:       []byte(clientID),
			config.AzureTenantID:       []byte(tenantID),
		},
	}

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.fakeProvider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.TokenFilePath).To(Equal(FederatedTokenFilePath))
}

func newResourceGroup(namespace string) *resources.ResourceGroup {
	return &resources.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind:       resolver.ResourceGroupKind,
			APIVersion: resources.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-rg",
			Namespace: namespace,
		},
		Spec: resources.ResourceGroup_Spec{
			Location:  to.Ptr("West US"),
			AzureName: "my-rg", // defaulter webhook will copy Name to AzureName
		},
	}
}

func newSecret(namespacedName types.NamespacedName) *v1.Secret {
	secretData := make(map[string][]byte)
	secretData[config.AzureClientID] = []byte(fakeID)
	secretData[config.AzureClientSecret] = []byte(fakeID)
	secretData[config.AzureTenantID] = []byte(fakeID)
	secretData[config.AzureSubscriptionID] = []byte(fakeID)

	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespacedName.Name,
			Namespace: namespacedName.Namespace,
		},
		Data: secretData,
	}
}

func createTestScheme() *runtime.Scheme {
	s := runtime.NewScheme()

	_ = v1.AddToScheme(s)
	_ = resources.AddToScheme(s)

	return s
}

func NewFakeKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}
