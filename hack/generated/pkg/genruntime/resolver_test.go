/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	batch "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.batch/v1alpha1api20170901"
	resources "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	storage "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.storage/v1alpha1api20190401"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"
)

const testNamespace = "testnamespace"

func NewKubeClient(s *runtime.Scheme) *kubeclient.Client {
	fakeClient := fake.NewFakeClientWithScheme(s)
	return kubeclient.NewClient(fakeClient, s)
}

func NewTestResolver(client *kubeclient.Client, reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind) *genruntime.Resolver {
	return genruntime.NewResolver(client, reconciledResourceLookup)
}

func MakeResourceGVKLookup(scheme *runtime.Scheme) (map[schema.GroupKind]schema.GroupVersionKind, error) {
	result := make(map[schema.GroupKind]schema.GroupVersionKind)

	// Register all types used in these tests
	objs := []runtime.Object{
		new(resources.ResourceGroup),
		new(batch.BatchAccount),
		new(storage.StorageAccount),
		new(storage.StorageAccountsBlobService),
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

func createResourceGroup(name string) *resources.ResourceGroup {
	return &resources.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind:       genruntime.ResourceGroupKind,
			APIVersion: resources.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: testNamespace,
		},
		Spec: resources.ResourceGroupSpec{
			Location:  "West US",
			AzureName: name, // defaulter webhook will copy Name to AzureName
		},
	}
}

func createResourceGroupRootedResource(rgName string, name string) (genruntime.MetaObject, genruntime.MetaObject) {
	a := createResourceGroup(rgName)

	b := &batch.BatchAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "BatchAccount",
			APIVersion: batch.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: testNamespace,
		},
		Spec: batch.BatchAccounts_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: rgName,
			},
			AzureName: name, // defaulter webhook will copy Name to AzureName
		},
	}

	return a, b
}

func createDeeplyNestedResource(rgName string, parentName string, name string) genruntime.ResourceHierarchy {
	a := createResourceGroup(rgName)

	b := &storage.StorageAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StorageAccount",
			APIVersion: storage.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      parentName,
			Namespace: testNamespace,
		},
		Spec: storage.StorageAccounts_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: rgName,
			},
			AzureName: parentName, // defaulter webhook will copy Name to AzureName
		},
	}

	c := &storage.StorageAccountsBlobService{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StorageAccountsBlobService",
			APIVersion: storage.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: testNamespace,
		},
		Spec: storage.StorageAccountsBlobServices_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: parentName,
			},
		},
	}

	return genruntime.ResourceHierarchy{a, b, c}
}

func Test_ResolveResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	resolver := NewTestResolver(NewKubeClient(s), reconciledResourceLookup)

	resourceGroupName := "myrg"
	a := createResourceGroup(resourceGroupName)

	hierarchy, err := resolver.ResolveResourceHierarchy(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(len(hierarchy)).To(Equal(1))
	g.Expect(hierarchy[0].GetName()).To(Equal(a.Name))
	g.Expect(hierarchy[0].GetNamespace()).To(Equal(a.Namespace))
}

func Test_ResolveResourceHierarchy_ResourceGroup_TopLevelResource(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	client := NewKubeClient(s)
	resolver := NewTestResolver(client, reconciledResourceLookup)

	resourceGroupName := "myrg"
	resourceName := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, resourceName)

	err = client.Client.Create(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())

	hierarchy, err := resolver.ResolveResourceHierarchy(ctx, b)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(len(hierarchy)).To(Equal(2))
	g.Expect(hierarchy[0].GetName()).To(Equal(a.GetName()))
	g.Expect(hierarchy[0].GetNamespace()).To(Equal(a.GetNamespace()))
	g.Expect(hierarchy[1].GetName()).To(Equal(b.GetName()))
	g.Expect(hierarchy[1].GetNamespace()).To(Equal(b.GetNamespace()))
}

func Test_ResolveResourceHierarchy_ResourceGroup_NestedResource(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	client := NewKubeClient(s)
	resolver := NewTestResolver(client, reconciledResourceLookup)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	originalHierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	for _, item := range originalHierarchy {
		err := client.Client.Create(ctx, item)
		g.Expect(err).ToNot(HaveOccurred())
	}

	hierarchy, err := resolver.ResolveResourceHierarchy(ctx, originalHierarchy[len(originalHierarchy)-1])
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(len(originalHierarchy)).To(Equal(len(hierarchy)))

	for i := 0; i < len(hierarchy); i++ {
		desired := originalHierarchy[i]
		actual := hierarchy[i]

		g.Expect(actual.GetName()).To(Equal(desired.GetName()))
		g.Expect(actual.GetNamespace()).To(Equal(desired.GetNamespace()))
	}
}

func Test_ResolveResourceHierarchy_ReturnsReferenceNotFound(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	resolver := NewTestResolver(NewKubeClient(s), reconciledResourceLookup)

	resourceGroupName := "myrg"
	resourceName := "myresource"

	_, b := createResourceGroupRootedResource(resourceGroupName, resourceName)

	// Purposefully skipping creating the RG

	_, err = resolver.ResolveResourceHierarchy(ctx, b)
	g.Expect(err).To(HaveOccurred())

	g.Expect(errors.Unwrap(err)).To(BeAssignableToTypeOf(&genruntime.ReferenceNotFound{}))
}

func Test_ResolveReference_FindsReference(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	client := NewKubeClient(s)
	resolver := NewTestResolver(client, reconciledResourceLookup)

	resourceGroupName := "myrg"

	resourceGroup := createResourceGroup(resourceGroupName)
	err = client.Client.Create(ctx, resourceGroup)
	g.Expect(err).ToNot(HaveOccurred())

	ref := genruntime.ResourceReference{Group: genruntime.ResourceGroupGroup, Kind: genruntime.ResourceGroupKind, Namespace: testNamespace, Name: resourceGroupName}
	resolved, err := resolver.ResolveReference(ctx, ref)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resolved).To(BeAssignableToTypeOf(&resources.ResourceGroup{}))
	resolvedRg := resolved.(*resources.ResourceGroup)

	g.Expect(resolvedRg.Spec.Location).To(Equal(resourceGroup.Spec.Location))
}

func Test_ResolveReference_ReturnsErrorIfReferenceIsNotAKubernetesReference(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()
	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	resolver := NewTestResolver(NewKubeClient(s), reconciledResourceLookup)

	ref := genruntime.ResourceReference{ARMID: "abcd"}
	_, err = resolver.ResolveReference(ctx, ref)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError("reference abcd is not pointing to a Kubernetes resource"))
}

func Test_ResolveReferenceToARMID_KubernetesResource_ReturnsExpectedID(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	client := NewKubeClient(s)
	resolver := NewTestResolver(client, reconciledResourceLookup)

	resourceGroupName := "myrg"
	armID := "/subscriptions/00000000-0000-0000-000000000000/resources/resourceGroups/myrg"

	resourceGroup := createResourceGroup(resourceGroupName)
	genruntime.SetResourceID(resourceGroup, armID)

	err = client.Client.Create(ctx, resourceGroup)
	g.Expect(err).ToNot(HaveOccurred())

	ref := genruntime.ResourceReference{Group: genruntime.ResourceGroupGroup, Kind: genruntime.ResourceGroupKind, Namespace: testNamespace, Name: resourceGroupName}
	id, err := resolver.ResolveReferenceToARMID(ctx, ref)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(id).To(Equal(armID))
}

func Test_ResolveReferenceToARMID_ARMResource_ReturnsExpectedID(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	client := NewKubeClient(s)
	resolver := NewTestResolver(client, reconciledResourceLookup)

	armID := "/subscriptions/00000000-0000-0000-000000000000/resources/resourceGroups/myrg"
	ref := genruntime.ResourceReference{ARMID: armID}
	id, err := resolver.ResolveReferenceToARMID(ctx, ref)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(id).To(Equal(armID))
}

func createTestScheme() *runtime.Scheme {
	s := runtime.NewScheme()
	_ = resources.AddToScheme(s)
	_ = batch.AddToScheme(s)
	_ = storage.AddToScheme(s)

	return s
}
