/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package multitenant_test

import (
	"testing"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"
)

const (
	resourceGroupName     = "aso-sample-rg"
	preReleaseNamespace   = "pre-release"
	vnetBeforeUpgradeName = "vnet-before-upgrade"
)

func Test_Pre_Release_ResourceCanBeCreated_BeforeUpgrade(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.Namespace = preReleaseNamespace

	rg := tc.NewTestResourceGroup()
	rg.Name = resourceGroupName

	vnet := newVnet(tc, vnetBeforeUpgradeName)
	// We want to use 'CreateResourceUntracked' here as we don't want the resource cleanup for this test. We will be checking the
	// backward compatibility for these resources after we upgrade the controller.
	tc.CreateResourceAndWaitWithoutCleanup(rg)
	tc.CreateResourceAndWaitWithoutCleanup(vnet)

}

func newVnet(tc *testcommon.KubePerTestContext, name string) *network.VirtualNetwork {
	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: network.VirtualNetworks_Spec{
			Owner: &genruntime.KnownResourceReference{
				Name: resourceGroupName,
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

	// TODO: Will have to change the version here when we go from beta to stable
	vnetAfterUpgrade := newVnet(tc, tc.Namer.GenerateName("vn"))
	tc.CreateResourceAndWait(vnetAfterUpgrade)
	tc.DeleteResourcesAndWait(vnetAfterUpgrade)

	// Patch for existing resource. This resource already will exist in kind as will be created without cleanup in test 'Test_Pre_Release_ResourceCanBeCreated_BeforeUpgrade'.
	vnetBeforeUpgrade := newVnet(tc, vnetBeforeUpgradeName)
	defer tc.DeleteResourceAndWait(vnetBeforeUpgrade)
	old := vnetBeforeUpgrade.DeepCopy()
	vnetBeforeUpgrade.Spec.Tags = map[string]string{"newTag": "afterUpgradeTag"}
	tc.Patch(old, vnetBeforeUpgrade)
	tc.Eventually(vnetBeforeUpgrade).Should(tc.Match.BeProvisioned(1))
	tc.GetResource(types.NamespacedName{Name: vnetBeforeUpgrade.Name, Namespace: vnetBeforeUpgrade.Namespace}, vnetBeforeUpgrade)
	tc.Expect(vnetBeforeUpgrade.Status.Tags).To(gomega.HaveKey("newTag"))

}
