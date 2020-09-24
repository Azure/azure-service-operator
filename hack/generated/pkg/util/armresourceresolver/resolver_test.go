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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/util/kubeclient"

	batch "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.batch/v20170901"
	storage "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.storage/v20190401"
)

func NewTestResolver(s *runtime.Scheme) *Resolver {
	fakeClient := fake.NewFakeClientWithScheme(s)
	return NewResolver(kubeclient.NewClient(fakeClient, s))
}

func Test_GetFullAzureNameAndResourceGroup_TopLevelResource(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := runtime.NewScheme()
	_ = storage.AddToScheme(s)

	resolver := NewTestResolver(s)

	resourceGroupName := "myrg"
	name := "myresource"

	a := &storage.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: storage.StorageAccounts_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: resourceGroupName,
			},
		},
	}

	rg, fullName, err := resolver.GetResourceGroupAndFullAzureName(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(fullName).To(Equal(name))
}

func Test_GetFullAzureNameAndResourceGroup_ChildResource(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := runtime.NewScheme()
	_ = batch.AddToScheme(s)

	resolver := NewTestResolver(s)

	resourceGroupName := "myrg"
	parentName := "myresource"
	childName := "myresource2"

	a := &batch.BatchAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: parentName,
		},
		Spec: batch.BatchAccounts_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: resourceGroupName,
			},
		},
	}
	err := resolver.client.Client.Create(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())

	b := &batch.BatchAccountsPool{
		ObjectMeta: metav1.ObjectMeta{
			Name: childName,
		},
		Spec: batch.BatchAccountsPools_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: parentName,
			},
		},
	}

	rg, fullName, err := resolver.GetResourceGroupAndFullAzureName(ctx, b)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rg).To(Equal(resourceGroupName))
	g.Expect(fullName).To(Equal(fmt.Sprintf("%s/%s", parentName, childName)))
}
