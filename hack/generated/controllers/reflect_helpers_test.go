/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.resources/v20200601"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/armresourceresolver"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"

	// TODO: Do we want to use a sample object rather than a code generated one?
	batch "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.batch/v20170901"

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
			ApiVersion: "apiVersion",
			AzureName:  "azureName",
			Location:   "westus",
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

type DummyStruct struct{}

func Test_ConvertResourceToDeployableResource(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	s := runtime.NewScheme()
	g.Expect(batch.AddToScheme(s)).To(Succeed())
	g.Expect(resources.AddToScheme(s)).To(Succeed())

	fakeClient := fake.NewFakeClientWithScheme(s)

	rg := createResourceGroup()
	g.Expect(fakeClient.Create(ctx, rg)).To(Succeed())
	account := createDummyResource()
	g.Expect(fakeClient.Create(ctx, account)).To(Succeed())
	resolver := armresourceresolver.NewResolver(kubeclient.NewClient(fakeClient, s))

	resource, err := ConvertResourceToDeployableResource(ctx, resolver, account)
	g.Expect(err).ToNot(HaveOccurred())

	rgResource, ok := resource.(*genruntime.ResourceGroupResource)
	g.Expect(ok).To(BeTrue())
	g.Expect("myrg").To(Equal(rgResource.ResourceGroup()))
	g.Expect("azureName").To(Equal(rgResource.Spec().GetName()))
	g.Expect("apiVersion").To(Equal(rgResource.Spec().GetApiVersion()))
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
	g.Expect(status).To(BeAssignableToTypeOf(&batch.BatchAccount_StatusArm{}))
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
