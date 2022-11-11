/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	. "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

const (
	fakeID = "00000000-0000-0000-0000-000000000000"
)

func NewFakeKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

func NewTestARMClientCache() (*ARMClientCache, error) {
	cfg, err := config.ReadFromEnvironment()
	if err != nil {
		return nil, err
	}

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	globalARMClient, err := genericarmclient.NewGenericClient(cfg.Cloud(), creds, cfg.SubscriptionID, metrics.NewARMClientMetrics())
	if err != nil {
		return nil, err
	}

	return NewARMClientCache(globalARMClient, cfg.PodNamespace, cfg.Cloud(), nil), nil
}

type testResources struct {
	armClientCache *ARMClientCache
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
	cache, err := NewTestARMClientCache()
	if err != nil {
		return nil, err
	}

	return &testResources{
		armClientCache: cache,
		kubeClient:     client,
	}, nil
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

	perResourceSecret := newSecret(perResourceCredentialName.Name, perResourceCredentialName.Namespace)

	err = res.kubeClient.Create(ctx, perResourceSecret)
	g.Expect(err).ToNot(HaveOccurred())

	namespacedSecretName := types.NamespacedName{
		Namespace: "test-secret",
		Name:      NamespacedSecretName,
	}

	secret := newSecret(namespacedSecretName.Name, namespacedSecretName.Namespace)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("")
	rg.Annotations = map[string]string{PerResourceSecretAnnotation: perResourceCredentialName.String()}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	client, credentialFrom, err := res.armClientCache.GetClient(ctx, rg, res.kubeClient)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(credentialFrom).To(BeEquivalentTo(perResourceCredentialName.String()))
	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(client.SubscriptionID()).To(BeEquivalentTo(fakeID))
	g.Expect(client).To(Not(BeEquivalentTo(res.armClientCache.globalClient.GenericClient())))
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
	rg.Annotations = map[string]string{PerResourceSecretAnnotation: credentialNamespacedName.String()}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	client, _, err := res.armClientCache.GetClient(ctx, rg, res.kubeClient)

	g.Expect(err).To(HaveOccurred())
	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	g.Expect(client).To(Not(BeEquivalentTo(res.armClientCache.globalClient.GenericClient())))
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

	secret := newSecret(credentialNamespacedName.Name, credentialNamespacedName.Namespace)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("")
	rg.Annotations = map[string]string{PerResourceSecretAnnotation: credentialNamespacedName.String()}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	client, credentialFrom, err := res.armClientCache.GetClient(ctx, rg, res.kubeClient)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(credentialFrom).To(BeEquivalentTo(credentialNamespacedName.String()))
	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(client.SubscriptionID()).To(BeEquivalentTo(fakeID))
	g.Expect(client).To(Not(BeEquivalentTo(res.armClientCache.globalClient.GenericClient())))
}

func Test_ARMClientCache_ReturnsNamespaceScopedClient(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-secret",
		Name:      NamespacedSecretName,
	}

	secret := newSecret(credentialNamespacedName.Name, credentialNamespacedName.Namespace)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup(credentialNamespacedName.Namespace)
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	client, credentialFrom, err := res.armClientCache.GetClient(ctx, rg, res.kubeClient)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(credentialFrom).To(BeEquivalentTo(credentialNamespacedName.String()))
	g.Expect(client.SubscriptionID()).To(BeEquivalentTo(fakeID))
	g.Expect(client).To(Not(BeEquivalentTo(res.armClientCache.globalClient.GenericClient())))
}

func Test_ARMClientCache_ReturnsNamespaceScopedClient_SecretChanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	res, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	credentialNamespacedName := types.NamespacedName{
		Namespace: "test-secret",
		Name:      NamespacedSecretName,
	}

	secret := newSecret(credentialNamespacedName.Name, credentialNamespacedName.Namespace)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup(credentialNamespacedName.Namespace)
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	oldClient, credentialFrom, err := res.armClientCache.GetClient(ctx, rg, res.kubeClient)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(credentialFrom).To(BeEquivalentTo(credentialNamespacedName.String()))
	g.Expect(oldClient.SubscriptionID()).To(BeEquivalentTo(fakeID))
	g.Expect(oldClient).To(Not(BeEquivalentTo(res.armClientCache.globalClient.GenericClient())))

	// change secret and check if we get a new client
	old := secret
	secret.Data[config.AzureClientIDVar] = []byte("11111111-1111-1111-1111-111111111111")
	err = res.kubeClient.Patch(ctx, secret, MergeFrom(old))
	g.Expect(err).ToNot(HaveOccurred())

	newClient, credentialFrom, err := res.armClientCache.GetClient(ctx, rg, res.kubeClient)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(credentialFrom).To(BeEquivalentTo(credentialNamespacedName.String()))
	g.Expect(newClient.SubscriptionID()).To(BeEquivalentTo(fakeID))
	g.Expect(newClient).To(Not(BeEquivalentTo(res.armClientCache.globalClient.GenericClient())))
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

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	client, _, err := res.armClientCache.GetClient(ctx, rg, res.kubeClient)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.armClientCache.clients)).To(BeEquivalentTo(0))
	g.Expect(client.SubscriptionID()).To(BeEquivalentTo(res.armClientCache.globalClient.GenericClient().SubscriptionID()))
	g.Expect(client).To(BeEquivalentTo(res.armClientCache.globalClient.GenericClient()))

}

func newSecret(name string, namespace string) *v1.Secret {
	secretData := make(map[string][]byte)
	secretData[config.AzureClientIDVar] = []byte(fakeID)
	secretData[config.AzureClientSecretVar] = []byte(fakeID)
	secretData[config.TenantIDVar] = []byte(fakeID)
	secretData[config.SubscriptionIDVar] = []byte(fakeID)

	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
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
		Spec: resources.ResourceGroupSpec{
			Location:  to.StringPtr("West US"),
			AzureName: "my-rg", // defaulter webhook will copy Name to AzureName
		},
	}
}
