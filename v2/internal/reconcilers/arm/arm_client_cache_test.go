/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	. "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	config2 "github.com/Azure/azure-service-operator/v2/pkg/common/config"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

const testPodNamespace = "azureserviceoperator-system-test"
const testSubscriptionID = "00000011-1111-0011-1100-110000000000" // Arbitrary GUID that isn't all 0s
const fakeID = "00000000-0000-0000-0000-000000000000"

func NewFakeKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

func NewTestCredentialProvider(client kubeclient.Client) (identity.CredentialProvider, error) {
	tokenCreds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	creds := identity.NewDefaultCredential(tokenCreds, testPodNamespace, testSubscriptionID)
	return identity.NewCredentialProvider(creds, client), nil
}

func NewTestARMClientCache(client kubeclient.Client) (*ARMClientCache, error) {
	cfg, err := config.ReadFromEnvironment()
	if err != nil {
		return nil, err
	}

	credentialProvider, err := NewTestCredentialProvider(client)
	if err != nil {
		return nil, err
	}

	return NewARMClientCache(credentialProvider, client, cfg.Cloud(), nil, metrics.NewARMClientMetrics()), nil
}

type testResources struct {
	ARMClientCache *ARMClientCache
	kubeClient     kubeclient.Client
}

func createTestScheme() *runtime.Scheme {
	s := runtime.NewScheme()

	_ = v1.AddToScheme(s)
	_ = resources.AddToScheme(s)

	return s
}

func testSetup() (*testResources, error) {
	s := createTestScheme()

	client := NewFakeKubeClient(s)
	cache, err := NewTestARMClientCache(client)
	if err != nil {
		return nil, err
	}

	return &testResources{
		ARMClientCache: cache,
		kubeClient:     client,
	}, nil
}

func Test_DefaultCredential_NotSet_ReturnsErrorWhenTryToUseGlobalCredential(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	s := createTestScheme()
	kubeClient := NewFakeKubeClient(s)

	cfg, err := config.ReadFromEnvironment()
	g.Expect(err).To(BeNil())

	providerWithNoDefaultCred := identity.NewCredentialProvider(nil, kubeClient)
	clientWithNoDefaultCred := NewARMClientCache(providerWithNoDefaultCred, kubeClient, cfg.Cloud(), nil, metrics.NewARMClientMetrics())

	rg := newResourceGroup("")

	_, err = clientWithNoDefaultCred.GetConnection(ctx, rg)
	g.Expect(err).ToNot(BeNil())
}

func Test_ARMClientCache_ReturnsPerResourceScopedClientOverNamespacedClient(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
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
		Name:      identity.NamespacedSecretName,
	}

	secret := newSecret(namespacedSecretName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	rg.Annotations = map[string]string{
		annotations.PerResourceSecret: perResourceCredentialName.Name,
	}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetConnection(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(details.CredentialFrom()).To(BeEquivalentTo(perResourceCredentialName))
	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(details.SubscriptionID()).To(BeEquivalentTo(fakeID))
}

func Test_ARMClientCache_PerResourceSecretInDifferentNamespace_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	perResourceCredentialName := types.NamespacedName{
		Namespace: "test-namespace-2",
		Name:      "test-secret",
	}

	perResourceSecret := newSecret(perResourceCredentialName)

	err = res.kubeClient.Create(ctx, perResourceSecret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	rg.Annotations = map[string]string{
		annotations.PerResourceSecret: perResourceCredentialName.String(),
	}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	_, err = res.ARMClientCache.GetConnection(ctx, rg)
	g.Expect(err).To(HaveOccurred())
	var target *core.SecretNotFound
	g.Expect(errors.As(err, &target)).To(BeTrue())
}

func Test_ARMClientCache_ReturnsError_IfSecretNotFound(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-namespace",
		Name:      "test-secret",
	}

	rg := newResourceGroup("")
	rg.Annotations = map[string]string{
		annotations.PerResourceSecret: credentialNamespacedName.Name,
	}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetConnection(ctx, rg)

	g.Expect(err).To(HaveOccurred())
	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	g.Expect(details).To(BeNil())
}

func Test_ARMClientCache_ReturnsPerResourceScopedClient(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-namespace",
		Name:      "test-secret",
	}

	secret := newSecret(credentialNamespacedName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	rg.Annotations = map[string]string{
		annotations.PerResourceSecret: credentialNamespacedName.Name,
	}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetConnection(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(details.CredentialFrom()).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(details.SubscriptionID()).To(BeEquivalentTo(fakeID))
}

func Test_ARMClientCache_ReturnsNamespaceScopedClient(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-secret",
		Name:      identity.NamespacedSecretName,
	}

	secret := newSecret(credentialNamespacedName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup(credentialNamespacedName.Namespace)
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetConnection(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(details.CredentialFrom()).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(details.SubscriptionID()).To(BeEquivalentTo(fakeID))
}

func Test_ARMClientCache_ReturnsNamespaceScopedClient_SecretChanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-secret",
		Name:      identity.NamespacedSecretName,
	}

	secret := newSecret(credentialNamespacedName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup(credentialNamespacedName.Namespace)
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	oldDetails, err := res.ARMClientCache.GetConnection(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(oldDetails.CredentialFrom()).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(oldDetails.SubscriptionID()).To(BeEquivalentTo(fakeID))

	// change secret and check if we get a new client
	old := secret
	secret.Data[config2.AzureClientID] = []byte("11111111-1111-1111-1111-111111111111")
	err = res.kubeClient.Patch(ctx, secret, MergeFrom(old))
	g.Expect(err).ToNot(HaveOccurred())

	newDetails, err := res.ARMClientCache.GetConnection(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(newDetails.CredentialFrom()).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(newDetails.SubscriptionID()).To(BeEquivalentTo(fakeID))
}

func Test_ARMClientCache_ReturnsGlobalClient(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("")
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetConnection(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(details.SubscriptionID()).To(BeEquivalentTo(testSubscriptionID))
}

func newSecret(namespacedName types.NamespacedName) *v1.Secret {
	secretData := make(map[string][]byte)
	secretData[config2.AzureClientID] = []byte(fakeID)
	secretData[config2.AzureClientSecret] = []byte(fakeID)
	secretData[config2.AzureTenantID] = []byte(fakeID)
	secretData[config2.AzureSubscriptionID] = []byte(fakeID)

	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespacedName.Name,
			Namespace: namespacedName.Namespace,
		},
		Data: secretData,
	}
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
