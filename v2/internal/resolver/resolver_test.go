/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver_test

import (
	"context"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1beta20210101"
	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta20210501"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"

	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

const testNamespace = "testnamespace"

func NewKubeClient(s *runtime.Scheme) kubeclient.Client {
	fakeClient := fake.NewClientBuilder().WithScheme(s).Build()
	return kubeclient.NewClient(fakeClient)
}

func NewTestResolver(client kubeclient.Client) (*resolver.Resolver, error) {
	res := resolver.NewResolver(client)
	// Register all types used in these tests
	objs := []*registration.StorageType{
		registration.NewStorageType(new(resources.ResourceGroup)),
		registration.NewStorageType(new(batch.BatchAccount)),
		registration.NewStorageType(new(storage.StorageAccount)),
		registration.NewStorageType(new(storage.StorageAccountsBlobService)),
		registration.NewStorageType(new(mysql.FlexibleServer)),
	}
	err := res.IndexStorageTypes(client.Scheme(), objs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type testResources struct {
	resolver *resolver.Resolver
	client   kubeclient.Client
}

func testSetup() (*testResources, error) {
	s := createTestScheme()

	client := NewKubeClient(s)
	res, err := NewTestResolver(client)
	if err != nil {
		return nil, err
	}

	return &testResources{
		resolver: res,
		client:   client,
	}, nil
}

func createResourceGroup(name string) *resources.ResourceGroup {
	return &resources.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind:       resolver.ResourceGroupKind,
			APIVersion: resources.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: testNamespace,
		},
		Spec: resources.ResourceGroupSpec{
			Location:  to.StringPtr("West US"),
			AzureName: name, // defaulter webhook will copy Name to AzureName
		},
	}
}

func createResourceGroupRootedResource(rgName string, name string) (genruntime.ARMMetaObject, genruntime.ARMMetaObject) {
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
		Spec: batch.BatchAccount_Spec{
			Owner: &genruntime.KnownResourceReference{
				Name: rgName,
			},
			AzureName: name, // defaulter webhook will copy Name to AzureName
		},
	}

	return a, b
}

func createDeeplyNestedResource(rgName string, parentName string, name string) resolver.ResourceHierarchy {
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
		Spec: storage.StorageAccount_Spec{
			Owner: &genruntime.KnownResourceReference{
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
		Spec: storage.StorageAccountsBlobService_Spec{
			Owner: &genruntime.KnownResourceReference{
				Name: parentName,
			},
		},
	}

	return resolver.ResourceHierarchy{a, b, c}
}

func createSimpleExtensionResource(name string, ownerName string, ownerGVK schema.GroupVersionKind) genruntime.ARMMetaObject {
	return &testcommon.SimpleExtensionResource{
		TypeMeta: metav1.TypeMeta{
			Kind:       "SimpleExtensionResource",
			APIVersion: testcommon.SimpleExtensionResourceGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: testNamespace,
		},
		Spec: testcommon.SimpleExtensionResourceSpec{
			Owner: genruntime.ResourceReference{
				Group: ownerGVK.Group,
				Kind:  ownerGVK.Kind,
				Name:  ownerName,
			},
			AzureName: name, // defaulter webhook will copy Name to AzureName
		},
	}
}

func createExtensionResourceOnResourceGroup(rgName string, name string) (genruntime.ARMMetaObject, genruntime.ARMMetaObject) {
	a := createResourceGroup(rgName)
	gvk := a.GetObjectKind().GroupVersionKind()
	b := createSimpleExtensionResource(name, a.GetName(), gvk)

	return a, b
}

func createExtensionResourceOnResourceInResourceGroup(rgName string, resourceName string, name string) resolver.ResourceHierarchy {
	a, b := createResourceGroupRootedResource(rgName, resourceName)
	gvk := b.GetObjectKind().GroupVersionKind()

	c := createSimpleExtensionResource(name, b.GetName(), gvk)

	return resolver.ResourceHierarchy{a, b, c}
}

func createExtensionResourceOnDeepHierarchyInResourceGroup(rgName string, parentName, resourceName string, name string) resolver.ResourceHierarchy {
	hierarchy := createDeeplyNestedResource(rgName, parentName, resourceName)
	extensionParent := hierarchy[len(hierarchy)-1]
	gvk := extensionParent.GetObjectKind().GroupVersionKind()

	extension := createSimpleExtensionResource(name, extensionParent.GetName(), gvk)

	return append(hierarchy, extension)
}

func Test_ResolveResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "myrg"
	a := createResourceGroup(resourceGroupName)

	hierarchy, err := test.resolver.ResolveResourceHierarchy(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(len(hierarchy)).To(Equal(1))
	g.Expect(hierarchy[0].GetName()).To(Equal(a.Name))
	g.Expect(hierarchy[0].GetNamespace()).To(Equal(a.Namespace))
}

func Test_ResolveResourceHierarchy_ResourceGroup_TopLevelResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "myrg"
	resourceName := "myresource"

	a, b := createResourceGroupRootedResource(resourceGroupName, resourceName)

	err = test.client.Create(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())

	hierarchy, err := test.resolver.ResolveResourceHierarchy(ctx, b)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(len(hierarchy)).To(Equal(2))
	g.Expect(hierarchy[0].GetName()).To(Equal(a.GetName()))
	g.Expect(hierarchy[0].GetNamespace()).To(Equal(a.GetNamespace()))
	g.Expect(hierarchy[1].GetName()).To(Equal(b.GetName()))
	g.Expect(hierarchy[1].GetNamespace()).To(Equal(b.GetNamespace()))
}

func Test_ResolveResourceHierarchy_ResourceGroup_NestedResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "myrg"
	resourceName := "myresource"
	childResourceName := "mychildresource"

	originalHierarchy := createDeeplyNestedResource(resourceGroupName, resourceName, childResourceName)

	for _, item := range originalHierarchy {
		err = test.client.Create(ctx, item)
		g.Expect(err).ToNot(HaveOccurred())
	}

	hierarchy, err := test.resolver.ResolveResourceHierarchy(ctx, originalHierarchy[len(originalHierarchy)-1])
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
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "myrg"
	resourceName := "myresource"

	_, b := createResourceGroupRootedResource(resourceGroupName, resourceName)

	// Purposefully skipping creating the RG

	_, err = test.resolver.ResolveResourceHierarchy(ctx, b)
	g.Expect(err).To(HaveOccurred())

	g.Expect(errors.Unwrap(err)).To(BeAssignableToTypeOf(&resolver.ReferenceNotFound{}))
}

func Test_ResolveReference_FindsReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "myrg"

	resourceGroup := createResourceGroup(resourceGroupName)
	err = test.client.Create(ctx, resourceGroup)
	g.Expect(err).ToNot(HaveOccurred())

	ref := genruntime.ResourceReference{Group: resolver.ResourceGroupGroup, Kind: resolver.ResourceGroupKind, Name: resourceGroupName}
	resolved, err := test.resolver.ResolveReference(ctx, ref.ToNamespacedRef(testNamespace))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resolved).To(BeAssignableToTypeOf(&resources.ResourceGroup{}))
	resolvedRg := resolved.(*resources.ResourceGroup)

	g.Expect(resolvedRg.Spec.Location).To(Equal(resourceGroup.Spec.Location))
}

func Test_ResolveReference_ReturnsErrorIfReferenceIsNotAKubernetesReference(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	ref := genruntime.ResourceReference{ARMID: "abcd"}
	_, err = test.resolver.ResolveReference(ctx, ref.ToNamespacedRef(""))
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError("reference abcd is not pointing to a Kubernetes resource"))
}

func Test_ResolveReferenceToARMID_KubernetesResource_ReturnsExpectedID(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "myrg"
	armID := "/subscriptions/00000000-0000-0000-000000000000/resources/resourceGroups/myrg"

	resourceGroup := createResourceGroup(resourceGroupName)
	genruntime.SetResourceID(resourceGroup, armID)

	err = test.client.Create(ctx, resourceGroup)
	g.Expect(err).ToNot(HaveOccurred())

	ref := genruntime.ResourceReference{Group: resolver.ResourceGroupGroup, Kind: resolver.ResourceGroupKind, Name: resourceGroupName}
	id, err := test.resolver.ResolveReferenceToARMID(ctx, ref.ToNamespacedRef(testNamespace))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(id).To(Equal(armID))
}

func Test_ResolveReferenceToARMID_ARMResource_ReturnsExpectedID(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	armID := "/subscriptions/00000000-0000-0000-000000000000/resources/resourceGroups/myrg"
	ref := genruntime.ResourceReference{ARMID: armID}
	id, err := test.resolver.ResolveReferenceToARMID(ctx, ref.ToNamespacedRef(""))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(id).To(Equal(armID))
}

func Test_ResolveSecrets_ReturnsExpectedSecretValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	resourceGroupName := "myrg"
	armID := "/subscriptions/00000000-0000-0000-000000000000/resources/resourceGroups/myrg"

	resourceGroup := createResourceGroup(resourceGroupName)
	genruntime.SetResourceID(resourceGroup, armID) // TODO: Do I actually need this here?

	err = test.client.Create(ctx, resourceGroup)
	g.Expect(err).ToNot(HaveOccurred())

	secretName := "testsecret"
	secretKey := "mysecretkey"
	secretValue := "myPinIs1234"
	// Create a secret
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: testNamespace,
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{
			secretKey: []byte(secretValue),
		},
		Type: "Opaque",
	}

	err = test.client.Create(ctx, secret)
	g.Expect(err).ToNot(HaveOccurred())

	ref := genruntime.SecretReference{Name: secretName, Key: secretKey}
	namespacedRef := ref.ToNamespacedRef(testNamespace)

	resolvedSecrets, err := test.resolver.ResolveSecretReferences(ctx, set.Make(namespacedRef))
	g.Expect(err).ToNot(HaveOccurred())

	actualSecret, err := resolvedSecrets.LookupSecret(ref)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(actualSecret).To(Equal(secretValue))
}

func Test_ResolveSecrets_ReturnsReferenceNotFound(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.TODO()

	test, err := testSetup()
	g.Expect(err).ToNot(HaveOccurred())

	secretName := "testsecret"
	secretKey := "mysecretkey"
	ref := genruntime.SecretReference{Name: secretName, Key: secretKey}
	namespacedRef := ref.ToNamespacedRef(testNamespace)

	_, err = test.resolver.ResolveSecretReferences(ctx, set.Make(namespacedRef))
	g.Expect(err).To(HaveOccurred())
	g.Expect(errors.Unwrap(err)).To(BeAssignableToTypeOf(&resolver.SecretNotFound{}))
}

func createTestScheme() *runtime.Scheme {
	s := runtime.NewScheme()
	_ = resources.AddToScheme(s)
	_ = batch.AddToScheme(s)
	_ = storage.AddToScheme(s)
	_ = mysql.AddToScheme(s)
	_ = v1.AddToScheme(s)

	return s
}
