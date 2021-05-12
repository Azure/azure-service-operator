/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
)

func Test_ResourceHierarchy_ResourceGroupOnly(t *testing.T) {
	g := NewWithT(t)

	resourceGroupName := "myrg"

	a := createResourceGroup(resourceGroupName)
	hierarchy := genruntime.ResourceHierarchy{a}

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
	hierarchy := genruntime.ResourceHierarchy{a, b}

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
