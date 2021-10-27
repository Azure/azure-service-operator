/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/azure-service-operator/v2/api/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	// TODO: Do we want to use a sample object rather than a code generated one?
	batch "github.com/Azure/azure-service-operator/v2/api/microsoft.batch/v1alpha1api20210101"

	. "github.com/onsi/gomega"
)

func createResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind: "ResourceGroup",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-namespace",
			Name:      "myrg",
		},
		Spec: resources.ResourceGroupSpec{
			Location: "West US",
		},
	}
}

func createDummyResource() *batch.BatchAccount {
	return &batch.BatchAccount{
		Spec: batch.BatchAccounts_Spec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}
}

func CreateScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	err := batch.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	err = resources.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	return scheme, nil
}

type testResolverAndFriends struct {
	scheme   *runtime.Scheme
	resolver *genruntime.Resolver
	client   client.Client
}

func makeTestResolver(scheme *runtime.Scheme) (testResolverAndFriends, error) {
	groupToVersionMap, err := MakeResourceGVKLookup(scheme)
	if err != nil {
		return testResolverAndFriends{}, err
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	resolver := genruntime.NewResolver(kubeclient.NewClient(fakeClient, scheme), groupToVersionMap)

	return testResolverAndFriends{
		scheme:   scheme,
		resolver: resolver,
		client:   fakeClient,
	}, nil
}

func MakeResourceGVKLookup(scheme *runtime.Scheme) (map[schema.GroupKind]schema.GroupVersionKind, error) {
	result := make(map[schema.GroupKind]schema.GroupVersionKind)

	// Register all types used in these tests
	objs := []runtime.Object{
		new(resources.ResourceGroup),
		new(batch.BatchAccount),
	}

	for _, obj := range objs {
		gvk, err := apiutil.GVKForObject(obj, scheme)
		if err != nil {
			return nil, errors.Wrapf(err, "creating GVK for obj %T", obj)
		}
		groupKind := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
		if existing, ok := result[groupKind]; ok {
			return nil, errors.Errorf("somehow group: %q, kind: %q was already registered with version %q", gvk.Group, gvk.Kind, existing.Version)
		}
		result[groupKind] = gvk
	}

	return result, nil
}

type DummyStruct struct{}

func Test_ConvertResourceToARMResource(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	subscriptionID := "1234"
	test, err := makeTestResolver(scheme)
	g.Expect(err).ToNot(HaveOccurred())

	rg := createResourceGroup()
	g.Expect(test.client.Create(ctx, rg)).To(Succeed())
	account := createDummyResource()
	g.Expect(test.client.Create(ctx, account)).To(Succeed())

	resource, err := reflecthelpers.ConvertResourceToARMResource(ctx, test.resolver, account, subscriptionID)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect("azureName").To(Equal(resource.Spec().GetName()))
	g.Expect("2021-01-01").To(Equal(resource.Spec().GetAPIVersion()))
	g.Expect("Microsoft.Batch/batchAccounts").To(Equal(resource.Spec().GetType()))
}

func Test_FindReferences(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	scheme, err := CreateScheme()
	g.Expect(err).ToNot(HaveOccurred())

	test, err := makeTestResolver(scheme)
	g.Expect(err).ToNot(HaveOccurred())

	rg := createResourceGroup()
	g.Expect(test.client.Create(ctx, rg)).To(Succeed())
	account := createDummyResource()
	ref := genruntime.ResourceReference{ARMID: "test"}
	account.Spec.KeyVaultReference = &batch.KeyVaultReference{
		Reference: ref,
	}
	g.Expect(test.client.Create(ctx, account)).To(Succeed())

	refs, err := reflecthelpers.FindResourceReferences(account)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(refs).To(HaveLen(1))
	g.Expect(refs).To(HaveKey(ref))
}

func Test_NewStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	account := createDummyResource()

	status, err := reflecthelpers.NewEmptyStatus(account)
	g.Expect(err).To(BeNil())
	g.Expect(status).To(BeAssignableToTypeOf(&batch.BatchAccount_Status{}))
}

func Test_EmptyArmResourceStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	account := createDummyResource()

	status, err := reflecthelpers.NewEmptyArmResourceStatus(account)
	g.Expect(err).To(BeNil())
	g.Expect(status).To(BeAssignableToTypeOf(&batch.BatchAccount_StatusARM{}))
}
