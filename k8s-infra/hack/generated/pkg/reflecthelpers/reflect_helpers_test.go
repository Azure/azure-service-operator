/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"

	// TODO: Do we want to use a sample object rather than a code generated one?
	batch "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.batch/v1alpha1api20170901"

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
			Properties: batch.BatchAccountCreateProperties{},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}
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

func Test_ConvertResourceToDeployableResource(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	scheme := runtime.NewScheme()
	g.Expect(batch.AddToScheme(scheme)).To(Succeed())
	g.Expect(resources.AddToScheme(scheme)).To(Succeed())

	groupToVersionMap, err := MakeResourceGVKLookup(scheme)
	g.Expect(err).ToNot(HaveOccurred())

	fakeClient := fake.NewFakeClientWithScheme(scheme)

	rg := createResourceGroup()
	g.Expect(fakeClient.Create(ctx, rg)).To(Succeed())
	account := createDummyResource()
	g.Expect(fakeClient.Create(ctx, account)).To(Succeed())
	resolver := genruntime.NewResolver(kubeclient.NewClient(fakeClient, scheme), groupToVersionMap)

	resource, err := ConvertResourceToDeployableResource(ctx, resolver, account)
	g.Expect(err).ToNot(HaveOccurred())

	rgResource, ok := resource.(*genruntime.ResourceGroupResource)
	g.Expect(ok).To(BeTrue())
	g.Expect("myrg").To(Equal(rgResource.ResourceGroup()))
	g.Expect("azureName").To(Equal(rgResource.Spec().GetName()))
	g.Expect("2017-09-01").To(Equal(rgResource.Spec().GetApiVersion()))
	g.Expect(string(batch.BatchAccountsSpecTypeMicrosoftBatchBatchAccounts)).To(Equal(rgResource.Spec().GetType()))

}

func Test_NewStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	account := createDummyResource()

	status, err := NewEmptyStatus(account)
	g.Expect(err).To(BeNil())
	g.Expect(status).To(BeAssignableToTypeOf(&batch.BatchAccount_Status{}))
}

func Test_EmptyArmResourceStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	account := createDummyResource()

	status, err := NewEmptyArmResourceStatus(account)
	g.Expect(err).To(BeNil())
	g.Expect(status).To(BeAssignableToTypeOf(&batch.BatchAccount_StatusARM{}))
}

func Test_HasStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	account := createDummyResource()
	result, err := HasStatus(account)
	g.Expect(err).To(BeNil())
	g.Expect(result).To(BeFalse())
}

func Test_NewPtrFromStruct_ReturnsPtr(t *testing.T) {
	g := NewGomegaWithT(t)

	v := DummyStruct{}
	ptr := NewPtrFromValue(v)
	g.Expect(ptr).To(BeAssignableToTypeOf(&DummyStruct{}))
}

func Test_NewPtrFromPrimitive_ReturnsPtr(t *testing.T) {
	g := NewGomegaWithT(t)

	v := 5
	ptr := NewPtrFromValue(v)

	expectedValue := &v

	g.Expect(ptr).To(Equal(expectedValue))
}

func Test_NewPtrFromPtr_ReturnsPtrPtr(t *testing.T) {
	g := NewGomegaWithT(t)

	ptr := &DummyStruct{}
	ptrPtr := NewPtrFromValue(ptr)
	g.Expect(ptrPtr).To(BeAssignableToTypeOf(&ptr))
}
