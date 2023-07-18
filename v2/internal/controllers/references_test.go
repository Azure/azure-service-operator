/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Note that this test uses a VM purely as an example resource which has a secret. The behavior will
// be the same for any resource that uses the azure_generic_arm_reconciler
func Test_MissingCrossResourceReference_ReturnsError(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)
	// Don't actually create any of the networking resources.
	// We're testing to make sure we get the right error when they're missing
	secret := createVMPasswordSecretAndRef(tc)
	vm := newVirtualMachine20201201(tc, rg, networkInterface, secret)

	tc.CreateResourceAndWaitForState(vm, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal("ReferenceNotFound"))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		Equal(fmt.Sprintf("failed resolving ARM IDs for references: %s/%s does not exist (NetworkInterface.network.azure.com \"%s\" not found)",
			tc.Namespace,
			networkInterface.Name,
			networkInterface.Name)))

	// Delete VM and resources.
	tc.DeleteResourcesAndWait(vm)
}
