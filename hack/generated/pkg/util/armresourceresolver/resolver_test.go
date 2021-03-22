/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armresourceresolver

import (
	"context"
	"fmt"
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

func NewTestResolver(s *runtime.Scheme, reconciledResourceLookup map[schema.GroupKind]schema.GroupVersionKind) *Resolver {
	fakeClient := fake.NewFakeClientWithScheme(s)

	return NewResolver(kubeclient.NewClient(fakeClient, s), reconciledResourceLookup)
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
			Kind: ResourceGroupKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
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
			Kind: "BatchAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
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

func createDeeplyNestedResource(rgName string, parentName string, name string) ResourceHierarchy {
	a := createResourceGroup(rgName)

	b := &storage.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: parentName,
		},
		Spec: storage.StorageAccounts_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: rgName,
			},
			AzureName: parentName, // defaulter webhook will copy Name to AzureName
		},
	}

	c := &storage.StorageAccountsBlobService{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: storage.StorageAccountsBlobServices_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: parentName,
			},
		},
	}

	return ResourceHierarchy{a, b, c}
}

// ResolveResourceHierarchy tests

func Test_ResolveResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	resolver := NewTestResolver(s, reconciledResourceLookup)

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
	resolver := NewTestResolver(s, reconciledResourceLookup)

	resourceGroupName := "myrg"
	resourceName := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, resourceName)

	err = resolver.client.Client.Create(ctx, a)
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
	resolver := NewTestResolver(s, reconciledResourceLookup)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	originalHierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	for _, item := range originalHierarchy {
		err := resolver.client.Client.Create(ctx, item)
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

func Test_ResolveResourceHierarchy_ReturnsOwnerNotFoundError(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := createTestScheme()

	reconciledResourceLookup, err := MakeResourceGVKLookup(s)
	g.Expect(err).ToNot(HaveOccurred())
	resolver := NewTestResolver(s, reconciledResourceLookup)

	resourceGroupName := "myrg"
	resourceName := "myresource"

	_, b := createResourceGroupRootedResource(resourceGroupName, resourceName)

	// Purposefully skipping creating the RG

	_, err = resolver.ResolveResourceHierarchy(ctx, b)
	g.Expect(err).To(HaveOccurred())

	g.Expect(errors.Unwrap(err)).To(BeAssignableToTypeOf(&OwnerNotFound{}))
}

// ResourceHierarchy tests

func Test_ResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	g := NewWithT(t)

	resourceGroupName := "myrg"

	a := createResourceGroup(resourceGroupName)
	hierarchy := ResourceHierarchy{a}

	// This is expected to fail
	_, err := hierarchy.ResourceGroup()
	g.Expect(err).To(HaveOccurred())

	location, err := hierarchy.Location()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(a.Spec.Location).To(Equal(location))
	g.Expect(hierarchy.FullAzureName()).To(Equal(resourceGroupName))
}

func Test_ResourceHierarchy_ResourceGroup_TopLevelResource(t *testing.T) {
	g := NewWithT(t)

	resourceGroupName := "myrg"
	name := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, name)
	hierarchy := ResourceHierarchy{a, b}

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullAzureName()).To(Equal(name))
}

func Test_ResourceHierarchy_ResourceGroup_NestedResource(t *testing.T) {
	g := NewWithT(t)

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	hierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	rg, err := hierarchy.ResourceGroup()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullAzureName()).To(Equal(fmt.Sprintf("%s/%s", hierarchy[1].AzureName(), hierarchy[2].AzureName())))
}

func createTestScheme() *runtime.Scheme {
	s := runtime.NewScheme()
	_ = resources.AddToScheme(s)
	_ = batch.AddToScheme(s)
	_ = storage.AddToScheme(s)

	return s
}
