/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

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
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

const (
	fakeID = "00000000-0000-0000-0000-000000000000"
)

func NewFakeKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

func NewTestARMClientCache(client kubeclient.Client) (*ARMClientCache, error) {
	cfg, err := config.ReadFromEnvironment()
	if err != nil {
		return nil, err
	}

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	options := &genericarmclient.GenericClientOptions{Metrics: metrics.NewARMClientMetrics()}
	globalARMClient, err := genericarmclient.NewGenericClient(cfg.Cloud(), creds, options)
	if err != nil {
		return nil, err
	}

	return NewARMClientCache(globalARMClient, cfg.SubscriptionID, cfg.PodNamespace, client, cfg.Cloud(), nil, metrics.NewARMClientMetrics()), nil
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

	clientWithNoDefaultCred := NewARMClientCache(nil, "", cfg.PodNamespace, kubeClient, cfg.Cloud(), nil, metrics.NewARMClientMetrics())

	rg := newResourceGroup("")

	_, err = clientWithNoDefaultCred.GetClient(ctx, rg)
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
		Name:      NamespacedSecretName,
	}

	secret := newSecret(namespacedSecretName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup("test-namespace")
	rg.Annotations = map[string]string{reconcilers.PerResourceSecretAnnotation: perResourceCredentialName.Name}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetClient(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(details.CredentialFrom).To(BeEquivalentTo(perResourceCredentialName))
	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(details.SubscriptionID).To(BeEquivalentTo(fakeID))
	g.Expect(details.Client).To(Not(BeEquivalentTo(res.ARMClientCache.globalClient.GenericClient())))
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
	rg.Annotations = map[string]string{reconcilers.PerResourceSecretAnnotation: perResourceCredentialName.String()}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	_, err = res.ARMClientCache.GetClient(ctx, rg)
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
	rg.Annotations = map[string]string{reconcilers.PerResourceSecretAnnotation: credentialNamespacedName.Name}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetClient(ctx, rg)

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
	rg.Annotations = map[string]string{reconcilers.PerResourceSecretAnnotation: credentialNamespacedName.Name}
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetClient(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(details.CredentialFrom).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(details.SubscriptionID).To(BeEquivalentTo(fakeID))
	g.Expect(details.Client).To(Not(BeEquivalentTo(res.ARMClientCache.globalClient.GenericClient())))
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

	secret := newSecret(credentialNamespacedName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup(credentialNamespacedName.Namespace)
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	details, err := res.ARMClientCache.GetClient(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(details.CredentialFrom).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(details.SubscriptionID).To(BeEquivalentTo(fakeID))
	g.Expect(details.Client).To(Not(BeEquivalentTo(res.ARMClientCache.globalClient.GenericClient())))
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

	secret := newSecret(credentialNamespacedName)

	err = res.kubeClient.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	rg := newResourceGroup(credentialNamespacedName.Namespace)
	err = res.kubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	oldDetails, err := res.ARMClientCache.GetClient(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(oldDetails.CredentialFrom).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(oldDetails.SubscriptionID).To(BeEquivalentTo(fakeID))
	g.Expect(oldDetails.Client).To(Not(BeEquivalentTo(res.ARMClientCache.globalClient.GenericClient())))

	// change secret and check if we get a new client
	old := secret
	secret.Data[config.ClientIDVar] = []byte("11111111-1111-1111-1111-111111111111")
	err = res.kubeClient.Patch(ctx, secret, MergeFrom(old))
	g.Expect(err).ToNot(HaveOccurred())

	newDetails, err := res.ARMClientCache.GetClient(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(1))
	g.Expect(newDetails.CredentialFrom).To(BeEquivalentTo(credentialNamespacedName))
	g.Expect(newDetails.SubscriptionID).To(BeEquivalentTo(fakeID))
	g.Expect(newDetails.Client).To(Not(BeEquivalentTo(res.ARMClientCache.globalClient.GenericClient())))
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
	details, err := res.ARMClientCache.GetClient(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(res.ARMClientCache.clients)).To(BeEquivalentTo(0))
	g.Expect(details.SubscriptionID).To(BeEquivalentTo(res.ARMClientCache.globalClient.subscriptionID))
	g.Expect(details.Client).To(BeEquivalentTo(res.ARMClientCache.globalClient.GenericClient()))
}

func newSecret(namespacedName types.NamespacedName) *v1.Secret {
	secretData := make(map[string][]byte)
	secretData[config.ClientIDVar] = []byte(fakeID)
	secretData[config.ClientSecretVar] = []byte(fakeID)
	secretData[config.TenantIDVar] = []byte(fakeID)
	secretData[config.SubscriptionIDVar] = []byte(fakeID)

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
