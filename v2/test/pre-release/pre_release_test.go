/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package multitenant_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	resourceGroupName     = "upgrade-test-rg"
	preReleaseNamespace   = "pre-release"
	vnetBeforeUpgradeName = "vnet-before-upgrade"
)

func Test_Pre_Release_ResourceCanBeCreated_BeforeUpgrade(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	newNamer := tc.Namer.WithNumRandomChars(0)
	rgName := newNamer.GenerateName(resourceGroupName)

	tc.Namespace = preReleaseNamespace

	rg := tc.NewTestResourceGroup()
	rg.Name = rgName

	vnet := newVnet(tc, newNamer.GenerateName(vnetBeforeUpgradeName), rgName)
	// We want to use 'CreateResourceAndWaitWithoutCleanup' here as we don't want the resource cleanup for this test. We will be checking the
	// backward compatibility for these resources after we upgrade the controller.
	tc.CreateResourceAndWaitWithoutCleanup(rg)
	tc.CreateResourceAndWaitWithoutCleanup(vnet)
}

func newVnet(tc *testcommon.KubePerTestContext, name string, rgName string) *network.VirtualNetwork {
	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: network.VirtualNetwork_Spec{
			Owner: &genruntime.KnownResourceReference{
				Name: rgName,
			},
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/8"},
			},
		},
	}
	return vnet
}

func Test_Pre_Release_ResourceCanBeCreated_AfterUpgrade(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.Namespace = preReleaseNamespace
	newNamer := tc.Namer.WithNumRandomChars(0)
	rgName := newNamer.GenerateName(resourceGroupName)

	// Ensure expected RG still exists
	rg := &resources.ResourceGroup{}
	tc.GetResource(types.NamespacedName{Namespace: tc.Namespace, Name: rgName}, rg)

	tc.Expect(rg.Status.Id).ToNot(BeNil())
	defer tc.DeleteResourcesAndWait(rg)

	// This resource already will exist in kind as will be created without cleanup in test 'Test_Pre_Release_ResourceCanBeCreated_BeforeUpgrade'.
	vnetBeforeUpgrade := newVnet(tc, newNamer.GenerateName(vnetBeforeUpgradeName), rgName)
	defer tc.DeleteResourceAndWait(vnetBeforeUpgrade)

	vnetAfterUpgrade := newVnet(tc, tc.Namer.GenerateName("vn"), rgName)
	tc.CreateResourceAndWait(vnetAfterUpgrade)
	tc.DeleteResourcesAndWait(vnetAfterUpgrade)

	// Patch for existing resource.
	old := vnetBeforeUpgrade.DeepCopy()
	vnetBeforeUpgrade.Spec.Tags = map[string]string{"newTag": "afterUpgradeTag"}
	tc.Patch(old, vnetBeforeUpgrade)
	tc.Eventually(vnetBeforeUpgrade).Should(tc.Match.BeProvisioned(1))
	tc.GetResource(types.NamespacedName{Name: vnetBeforeUpgrade.Name, Namespace: vnetBeforeUpgrade.Namespace}, vnetBeforeUpgrade)
	tc.Expect(vnetBeforeUpgrade.Status.Tags).To(gomega.HaveKey("newTag"))
}
