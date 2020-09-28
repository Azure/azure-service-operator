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

func createTopLevelResource(rgName string, name string) *storage.StorageAccount {
	return &storage.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "foo",
			Name:      name,
		},
		Spec: storage.StorageAccounts_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: rgName,
			},
		},
	}
}

func createResourceAndParent(rgName string, parentName string, name string) (genruntime.MetaObject, genruntime.MetaObject) {
	a := &batch.BatchAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: parentName,
		},
		Spec: batch.BatchAccounts_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: rgName,
			},
		},
	}

	b := &batch.BatchAccountsPool{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: batch.BatchAccountsPools_Spec{
			Owner: genruntime.KnownResourceReference{
				Name: parentName,
			},
		},
	}

	return a, b
}

func Test_ResolveResourceHierarchy_TopLevelResource(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := runtime.NewScheme()
	_ = storage.AddToScheme(s)
	resolver := NewTestResolver(s)

	resourceGroupName := "myrg"
	name := "myresource"
	a := createTopLevelResource(resourceGroupName, name)

	hierarchy, err := resolver.ResolveResourceHierarchy(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(1).To(Equal(len(hierarchy)))
	g.Expect(a.Name).To(Equal(hierarchy[0].GetName()))
	g.Expect(a.Namespace).To(Equal(hierarchy[0].GetNamespace()))
}

func Test_ResolveResourceHierarchy_ChildResource(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := runtime.NewScheme()
	_ = batch.AddToScheme(s)

	resolver := NewTestResolver(s)

	resourceGroupName := "myrg"
	parentName := "myresource"
	childName := "myresource2"

	a, b := createResourceAndParent(resourceGroupName, parentName, childName)

	err := resolver.client.Client.Create(ctx, a)
	g.Expect(err).ToNot(HaveOccurred())

	hierarchy, err := resolver.ResolveResourceHierarchy(ctx, b)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(2).To(Equal(len(hierarchy)))
	g.Expect(a.GetName()).To(Equal(hierarchy[0].GetName()))
	g.Expect(a.GetNamespace()).To(Equal(hierarchy[0].GetNamespace()))
	g.Expect(b.GetName()).To(Equal(hierarchy[1].GetName()))
	g.Expect(b.GetNamespace()).To(Equal(hierarchy[1].GetNamespace()))
}

func Test_ResourceHierarchy_TopLevelResource(t *testing.T) {
	g := NewWithT(t)

	resourceGroupName := "myrg"
	name := "myresource"

	a := createTopLevelResource(resourceGroupName, name)
	hierarchy := ResourceHierarchy{a}

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullAzureName()).To(Equal(name))
}

func Test_ResourceHierarchy_ChildResource(t *testing.T) {
	g := NewWithT(t)

	resourceGroupName := "myrg"
	parentName := "myresource"
	childName := "myresource2"

	a, b := createResourceAndParent(resourceGroupName, parentName, childName)
	hierarchy := ResourceHierarchy{a, b}

	g.Expect(hierarchy.ResourceGroup()).To(Equal(resourceGroupName))
	g.Expect(hierarchy.FullAzureName()).To(Equal(fmt.Sprintf("%s/%s", parentName, childName)))
}
