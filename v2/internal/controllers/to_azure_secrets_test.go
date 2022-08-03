/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Note that this test uses a VM purely as an example resource which has a secret. The behavior will
// be the same for any resource that uses the azure_generic_arm_reconciler
func Test_MissingSecret_ReturnsError_ReconcilesSuccessfullyWhenSecretAdded(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)
	// Inefficient but avoids triggering the vnet/subnets problem.
	// https://github.com/Azure/azure-service-operator/issues/1944
	tc.CreateResourceAndWait(vnet)
	tc.CreateResourcesAndWait(subnet, networkInterface)
	secretRef := genruntime.SecretReference{
		Name: "thisdoesntexist",
		Key:  "key",
	}
	vm := newVirtualMachine20201201(tc, rg, networkInterface, secretRef)

	tc.CreateResourceAndWaitForState(vm, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal("SecretNotFound"))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		ContainSubstring("failed resolving secret references: %s/%s does not exist", tc.Namespace, secretRef.Name))

	// Now create the secret
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMetaWithName(secretRef.Name),
		StringData: map[string]string{
			secretRef.Key: tc.Namer.GeneratePasswordOfLength(40),
		},
	}
	tc.CreateResource(secret)

	// Wait for the VM to make it to a steady state
	tc.Eventually(vm).Should(tc.Match.BeProvisioned(0))

	// Delete VM.
	tc.DeleteResourcesAndWait(vm)
}

func Test_MissingSecretKey_ReturnsError(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)
	// Inefficient but avoids triggering the vnet/subnets problem.
	// https://github.com/Azure/azure-service-operator/issues/1944
	tc.CreateResourceAndWait(vnet)
	tc.CreateResourcesAndWait(subnet, networkInterface)
	secret := createVMPasswordSecretAndRef(tc)
	secret.Key = "doesnotexist" // Change the key to a key that doesn't actually exist
	vm := newVirtualMachine20201201(tc, rg, networkInterface, secret)

	tc.CreateResourceAndWaitForState(vm, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal("SecretNotFound"))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		ContainSubstring("Secret \"%s/%s\" does not contain key \"%s\"", tc.Namespace, secret.Name, secret.Key))

	// Delete VM and resources.
	tc.DeleteResourceAndWait(rg)
}
