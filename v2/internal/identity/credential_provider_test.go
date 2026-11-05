/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/benbjohnson/clock"
	"github.com/google/uuid"
	authenticationv1 "k8s.io/api/authentication/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	internalconfig "github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	asocloud "github.com/Azure/azure-service-operator/v2/pkg/common/cloud"
	"github.com/Azure/azure-service-operator/v2/pkg/common/config"
)

const (
	testPodNamespace   = "azureserviceoperator-system-test"
	testSubscriptionID = "00000011-1111-0011-1100-110000000000" // Arbitrary GUID that isn't all 0s
	fakeID             = "00000000-0000-0000-0000-000000000000"
)

type testCredentialProviderResources struct {
	kubeClient                  kubeclient.Client
	Provider                    CredentialProvider
	fakeTokenCredentialProvider *mockTokenCredentialProvider
}

func testCredentialProviderSetup(cloud *cloud.Configuration) (*testCredentialProviderResources, error) {
	return testCredentialProviderSetupWithOptions(cloud, &CredentialProviderOptions{
		WorkloadIdentityAuthMode: internalconfig.WorkloadIdentityAuthModeRelaxed,
	})
}

// testCredentialProviderSetupWithOptions is like testCredentialProviderSetup but lets the caller
// override any CredentialProviderOptions field (TokenProvider/Cloud are always overridden with the
// test fake/default below, regardless of what's passed in).
func testCredentialProviderSetupWithOptions(cloud *cloud.Configuration, opts *CredentialProviderOptions) (*testCredentialProviderResources, error) {
	s := createTestScheme()

	if cloud == nil {
		cloud = to.Ptr(asocloud.Configuration{}.Cloud())
	}

	// Global creds
	tokenCreds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	creds := NewDefaultCredential(
		tokenCreds,
		testPodNamespace,
		testSubscriptionID,
		nil,
	)

	client := NewFakeKubeClient(s)

	fakeTokenCredentialProvider := &mockTokenCredentialProvider{}
	if opts == nil {
		opts = &CredentialProviderOptions{}
	}
	opts.TokenProvider = fakeTokenCredentialProvider
	opts.Cloud = cloud

	provider := NewCredentialProvider(creds, client, opts)

	return &testCredentialProviderResources{
		kubeClient:                  client,
		fakeTokenCredentialProvider: fakeTokenCredentialProvider,
		Provider:                    provider,
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

	res, err := testCredentialProviderSetup(nil)
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

	res, err := testCredentialProviderSetup(nil)
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

	res, err := testCredentialProviderSetup(nil)
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

	res, err := testCredentialProviderSetup(nil)
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

	res, err := testCredentialProviderSetup(nil)
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

	cred, err := res.Provider.GetCredential(ctx, rg)
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

	res, err := testCredentialProviderSetup(nil)
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

	cred, err := res.Provider.GetCredential(ctx, rg)
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

	res, err := testCredentialProviderSetup(nil)
	g.Expect(err).ToNot(HaveOccurred())

	clientID := uuid.New().String()
	tenantID := uuid.New().String()

	secret := newWorkloadIdentitySecret(clientID, tenantID)
	secret.Data[config.WorkloadIdentityServiceAccount] = []byte("ignored-in-relaxed-mode")

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.TokenFilePath).To(Equal(FederatedTokenFilePath))
	g.Expect(res.fakeTokenCredentialProvider.GetAssertion).To(BeNil())
}

func TestCredentialProvider_StrictWorkloadIdentityCredential_UsesDefaultServiceAccount(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()
	testClock := clock.NewMock()

	assertionProvider := newTestServiceAccountTokenProvider(
		t,
		testClock,
		func(
			_ context.Context,
			serviceAccount *v1.ServiceAccount,
			tokenRequest *authenticationv1.TokenRequest,
		) error {
			g.Expect(serviceAccount.Namespace).To(Equal("test-namespace"))
			g.Expect(serviceAccount.Name).To(Equal(DefaultWorkloadIdentityServiceAccount))
			g.Expect(tokenRequest.Spec.Audiences).To(Equal([]string{workloadIdentityAudience}))
			g.Expect(tokenRequest.Spec.ExpirationSeconds).NotTo(BeNil())
			g.Expect(*tokenRequest.Spec.ExpirationSeconds).To(Equal(int64(workloadIdentityTokenLifetime / time.Second)))
			setTokenRequestStatus(tokenRequest, "strict-assertion", testClock.Now().Add(time.Hour))
			return nil
		},
	)
	res, err := testCredentialProviderSetupWithOptions(nil, &CredentialProviderOptions{
		ServiceAccountTokenProvider: assertionProvider,
		WorkloadIdentityAuthMode:    internalconfig.WorkloadIdentityAuthModeStrict,
	})
	g.Expect(err).ToNot(HaveOccurred())

	clientID := uuid.New().String()
	tenantID := uuid.New().String()
	secret := newWorkloadIdentitySecret(clientID, tenantID)
	g.Expect(res.kubeClient.Create(ctx, secret)).To(Succeed())

	_, err = res.Provider.GetCredential(ctx, newResourceGroup("test-namespace"))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.TokenFilePath).To(BeEmpty())
	g.Expect(res.fakeTokenCredentialProvider.GetAssertion).NotTo(BeNil())

	assertion, err := res.fakeTokenCredentialProvider.GetAssertion(ctx)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(assertion).To(Equal("strict-assertion"))
}

func TestCredentialProvider_StrictWorkloadIdentityCredential_UsesConfiguredServiceAccountAndOptions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()
	testClock := clock.NewMock()
	customCloud := asocloud.Configuration{
		AzureAuthorityHost:      "https://login.example.com/",
		ResourceManagerEndpoint: "https://management.example.com/",
		ResourceManagerAudience: "https://management.example.com/",
	}.Cloud()

	assertionProvider := newTestServiceAccountTokenProvider(
		t,
		testClock,
		func(
			_ context.Context,
			serviceAccount *v1.ServiceAccount,
			tokenRequest *authenticationv1.TokenRequest,
		) error {
			g.Expect(serviceAccount.Name).To(Equal("custom-workload"))
			setTokenRequestStatus(tokenRequest, "custom-assertion", testClock.Now().Add(time.Hour))
			return nil
		},
	)
	res, err := testCredentialProviderSetupWithOptions(&customCloud, &CredentialProviderOptions{
		ServiceAccountTokenProvider: assertionProvider,
		WorkloadIdentityAuthMode:    internalconfig.WorkloadIdentityAuthModeStrict,
	})
	g.Expect(err).ToNot(HaveOccurred())

	additionalTenants := []string{uuid.New().String(), uuid.New().String()}
	secret := newWorkloadIdentitySecret(uuid.New().String(), uuid.New().String())
	secret.Data[config.WorkloadIdentityServiceAccount] = []byte("custom-workload")
	secret.Data[config.AzureAdditionalTenants] = []byte(strings.Join(additionalTenants, ","))
	g.Expect(res.kubeClient.Create(ctx, secret)).To(Succeed())

	_, err = res.Provider.GetCredential(ctx, newResourceGroup("test-namespace"))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(res.fakeTokenCredentialProvider.AdditionalTenants).To(Equal(additionalTenants))
	g.Expect(res.fakeTokenCredentialProvider.Cloud).To(Equal(customCloud))

	assertion, err := res.fakeTokenCredentialProvider.GetAssertion(ctx)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(assertion).To(Equal("custom-assertion"))
}

func TestCredentialProvider_StrictWorkloadIdentityCredential_RejectsInvalidServiceAccount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
	}{
		{name: "empty", value: ""},
		{name: "invalid DNS name", value: "Not_A_ServiceAccount"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			res, err := testCredentialProviderSetupWithOptions(nil, &CredentialProviderOptions{
				WorkloadIdentityAuthMode: internalconfig.WorkloadIdentityAuthModeStrict,
				ServiceAccountTokenProvider: newTestServiceAccountTokenProvider(
					t,
					clock.NewMock(),
					func(
						_ context.Context,
						_ *v1.ServiceAccount,
						_ *authenticationv1.TokenRequest,
					) error {
						return nil
					},
				),
			})
			g.Expect(err).ToNot(HaveOccurred())

			secret := newWorkloadIdentitySecret(uuid.New().String(), uuid.New().String())
			secret.Data[config.WorkloadIdentityServiceAccount] = []byte(test.value)
			g.Expect(res.kubeClient.Create(context.Background(), secret)).To(Succeed())

			_, err = res.Provider.GetCredential(context.Background(), newResourceGroup("test-namespace"))
			expectedError := fmt.Sprintf(
				`credential secret "test-namespace/aso-credential" contains invalid %s %q`,
				config.WorkloadIdentityServiceAccount,
				test.value,
			)
			g.Expect(err).To(MatchError(ContainSubstring(expectedError)))
			g.Expect(res.fakeTokenCredentialProvider.GetAssertion).To(BeNil())
		})
	}
}

func TestCredentialProvider_StrictMode_DoesNotAffectClientSecretCredential(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()
	res, err := testCredentialProviderSetupWithOptions(nil, &CredentialProviderOptions{
		WorkloadIdentityAuthMode: internalconfig.WorkloadIdentityAuthModeStrict,
	})
	g.Expect(err).ToNot(HaveOccurred())

	secret := newWorkloadIdentitySecret(uuid.New().String(), uuid.New().String())
	secret.Data[config.AzureClientSecret] = []byte("client-secret")
	secret.Data[config.WorkloadIdentityServiceAccount] = []byte("")
	g.Expect(res.kubeClient.Create(ctx, secret)).To(Succeed())

	_, err = res.Provider.GetCredential(ctx, newResourceGroup("test-namespace"))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(res.fakeTokenCredentialProvider.ClientSecret).To(Equal("client-secret"))
	g.Expect(res.fakeTokenCredentialProvider.GetAssertion).To(BeNil())
}

func TestCredentialProvider_WorkloadIdentityCredential_HonoursFederatedTokenFilePathOption(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := t.Context()

	const customTokenPath = "/var/run/secrets/azure/tokens/azure-identity-token" // #nosec G101 -- file path, not a credential

	res, err := testCredentialProviderSetupWithOptions(nil, &CredentialProviderOptions{
		FederatedTokenFilePath: customTokenPath,
	})
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

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.TokenFilePath).To(Equal(customTokenPath))
}

func TestResolveFederatedTokenFilePath_HonoursOverrideThenFallsBackToDefault(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Empty override -> default hardcoded path.
	g.Expect(ResolveFederatedTokenFilePath("")).To(Equal(FederatedTokenFilePath))

	// Whitespace-only override -> still treated as unset, default hardcoded path.
	g.Expect(ResolveFederatedTokenFilePath("   ")).To(Equal(FederatedTokenFilePath))

	// Set -> the override takes precedence (whitespace-trimmed).
	const customTokenPath = "/var/run/secrets/azure/tokens/azure-identity-token" // #nosec G101 -- file path, not a credential
	g.Expect(ResolveFederatedTokenFilePath("  " + customTokenPath + "  ")).To(Equal(customTokenPath))
}

func TestCredentialProvider_AdditionalTenants_AreConfiguredCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup(nil)
	g.Expect(err).ToNot(HaveOccurred())

	clientID := uuid.New().String()
	tenantID := uuid.New().String()
	clientSecret := uuid.New().String()
	additionalTenants := []string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      NamespacedSecretName,
		},
		Data: map[string][]byte{
			config.AzureSubscriptionID:    []byte(testSubscriptionID),
			config.AzureClientID:          []byte(clientID),
			config.AzureTenantID:          []byte(tenantID),
			config.AzureClientSecret:      []byte(clientSecret),
			config.AzureAdditionalTenants: []byte(strings.Join(additionalTenants, ",")),
		},
	}
	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.ClientSecret).To(Equal(clientSecret))
	g.Expect(res.fakeTokenCredentialProvider.AdditionalTenants).To(Equal(additionalTenants))
}

func TestCredentialProvider_NonstandardClouds_AreConfiguredCorrectly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	customCloud := asocloud.Configuration{
		AzureAuthorityHost:      "specialhost",
		ResourceManagerEndpoint: "specialendpoint",
		ResourceManagerAudience: "specialaudience",
	}.Cloud()
	res, err := testCredentialProviderSetup(&customCloud)
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

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(cred.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
	g.Expect(res.fakeTokenCredentialProvider.ClientID).To(Equal(clientID))
	g.Expect(res.fakeTokenCredentialProvider.TenantID).To(Equal(tenantID))
	g.Expect(res.fakeTokenCredentialProvider.ClientSecret).To(Equal(clientSecret))
	g.Expect(res.fakeTokenCredentialProvider.Cloud).To(Equal(customCloud))
}

func TestCredentialProvider_NamespaceCredentialMissingRequiredFields_Errors(t *testing.T) {
	t.Parallel()
	ctx := context.TODO()

	clientID := uuid.New().String()
	tenantID := uuid.New().String()

	tests := []struct {
		name        string
		data        map[string][]byte
		expectedErr string
	}{
		{
			name: "missing sub id",
			data: map[string][]byte{
				config.AzureClientID: []byte(clientID),
				config.AzureTenantID: []byte(tenantID),
			},
			expectedErr: "does not contain key \"AZURE_SUBSCRIPTION_ID\"",
		},
		{
			name: "missing client id",
			data: map[string][]byte{
				config.AzureSubscriptionID: []byte(testSubscriptionID),
				config.AzureTenantID:       []byte(tenantID),
			},
			expectedErr: "does not contain key \"AZURE_CLIENT_ID\"",
		},
		{
			name: "missing tenant id",
			data: map[string][]byte{
				config.AzureSubscriptionID: []byte(testSubscriptionID),
				config.AzureClientID:       []byte(clientID),
			},
			expectedErr: "does not contain key \"AZURE_TENANT_ID\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			res, err := testCredentialProviderSetup(nil)
			g.Expect(err).ToNot(HaveOccurred())

			secret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-namespace",
					Name:      NamespacedSecretName,
				},
				Data: tt.data,
			}
			err = res.kubeClient.Create(ctx, secret)
			g.Expect(err).ToNot(HaveOccurred())

			rg := newResourceGroup("test-namespace")
			err = res.kubeClient.Create(ctx, rg)
			g.Expect(err).ToNot(HaveOccurred())

			_, err = res.Provider.GetCredential(ctx, rg)
			g.Expect(err).To(MatchError(ContainSubstring(tt.expectedErr)))
		})
	}
}

func TestCredentialProvider_CrossNamespaceCredentials_Blocked(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetup(nil)
	g.Expect(err).ToNot(HaveOccurred())

	perResourceCredentialName := types.NamespacedName{
		Namespace: "test-namespace2",
		Name:      "test-secret",
	}
	secret := newSecret(perResourceCredentialName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	rg.Annotations = map[string]string{annotations.PerResourceSecret: perResourceCredentialName.String()}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	_, err = res.Provider.GetCredential(ctx, rg)
	g.Expect(err).To(MatchError(ContainSubstring("cannot contain '/'. Secret must be in same namespace as resource.")))
}

func TestCredentialProvider_AllowMultiEnvManagement_Disabled_RejectsCloudConfigInSecret(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetupWithOptions(nil, &CredentialProviderOptions{
		WorkloadIdentityAuthMode: internalconfig.WorkloadIdentityAuthModeRelaxed,
	})
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
			config.AzureSubscriptionID:     []byte(testSubscriptionID),
			config.AzureClientID:           []byte(clientID),
			config.AzureTenantID:           []byte(tenantID),
			config.AzureClientSecret:       []byte(clientSecret),
			config.ResourceManagerEndpoint: []byte("https://management.usgovcloudapi.net"),
			config.ResourceManagerAudience: []byte("https://management.core.usgovcloudapi.net/"),
			config.AzureAuthorityHost:      []byte("https://login.microsoftonline.us/"),
		},
	}
	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	_, err = res.Provider.GetCredential(ctx, rg)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring(config.AllowMultiEnvManagement))
}

func TestCredentialProvider_AllowMultiEnvManagement_Enabled_PartialCloudConfigRejected(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetupWithOptions(nil, &CredentialProviderOptions{
		AllowMultiEnvManagement:  true,
		WorkloadIdentityAuthMode: internalconfig.WorkloadIdentityAuthModeRelaxed,
	})
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
			// Only 1 of 3 cloud config fields
			config.ResourceManagerEndpoint: []byte("https://management.usgovcloudapi.net"),
		},
	}
	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	_, err = res.Provider.GetCredential(ctx, rg)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("must specify ALL or NONE"))
}

func TestCredentialProvider_AllowMultiEnvManagement_Enabled_UsesCloudConfigFromSecret(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testCredentialProviderSetupWithOptions(nil, &CredentialProviderOptions{
		AllowMultiEnvManagement:  true,
		WorkloadIdentityAuthMode: internalconfig.WorkloadIdentityAuthModeRelaxed,
	})
	g.Expect(err).ToNot(HaveOccurred())

	clientID := uuid.New().String()
	tenantID := uuid.New().String()
	clientSecret := uuid.New().String()

	customEndpoint := "https://management.usgovcloudapi.net"
	customAudience := "https://management.core.usgovcloudapi.net/"
	customAuthorityHost := "https://login.microsoftonline.us/"

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      NamespacedSecretName,
		},
		Data: map[string][]byte{
			config.AzureSubscriptionID:     []byte(testSubscriptionID),
			config.AzureClientID:           []byte(clientID),
			config.AzureTenantID:           []byte(tenantID),
			config.AzureClientSecret:       []byte(clientSecret),
			config.ResourceManagerEndpoint: []byte(customEndpoint),
			config.ResourceManagerAudience: []byte(customAudience),
			config.AzureAuthorityHost:      []byte(customAuthorityHost),
		},
	}
	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	cred, err := res.Provider.GetCredential(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	// Credential should have cloud config matching what was supplied
	cloudCfg := cred.CloudConfig()
	g.Expect(cloudCfg).ToNot(BeNil())
	g.Expect(cloudCfg.ActiveDirectoryAuthorityHost).To(Equal(customAuthorityHost))
	g.Expect(cloudCfg.Services[cloud.ResourceManager].Endpoint).To(Equal(customEndpoint))
	g.Expect(cloudCfg.Services[cloud.ResourceManager].Audience).To(Equal(customAudience))

	// And the token credential provider should have been invoked with those settings
	g.Expect(res.fakeTokenCredentialProvider.Cloud.ActiveDirectoryAuthorityHost).To(Equal(customAuthorityHost))
	g.Expect(res.fakeTokenCredentialProvider.Cloud.Services[cloud.ResourceManager].Endpoint).To(Equal(customEndpoint))
	g.Expect(res.fakeTokenCredentialProvider.Cloud.Services[cloud.ResourceManager].Audience).To(Equal(customAudience))
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

func newWorkloadIdentitySecret(clientID string, tenantID string) *v1.Secret {
	return &v1.Secret{
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
