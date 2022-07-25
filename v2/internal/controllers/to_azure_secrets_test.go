/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Note that this test uses a VM purely as an example resource which has a secret. The behavior will
// be the same for any resource that uses the azure_generic_arm_reconciler
func Test_MissingUserSecret_ReturnsError(t *testing.T) {
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
	secret := genruntime.SecretReference{
		Name: "thisdoesntexist",
		Key:  "key",
	}
	vm := newVM(tc, rg, networkInterface, secret)

	tc.CreateResourceAndWaitForState(vm, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal(conditions.ReasonSecretNotFound))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		ContainSubstring("failed resolving secret references: %s/%s does not exist", tc.Namespace, secret.Name))

	// Delete VM and resources.
	tc.DeleteResourcesAndWait(vm, networkInterface, subnet, vnet, rg)
}

func Test_MissingSecretKey_ReturnsError(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)

	tc.RunParallelSubtests()
	// Inefficient but avoids triggering the vnet/subnets problem.
	// https://github.com/Azure/azure-service-operator/issues/1944
	tc.CreateResourceAndWait(vnet)
	tc.CreateResourcesAndWait(subnet, networkInterface)
	secret := createVMPasswordSecretAndRef(tc)
	secret.Key = "doesnotexist" // Change the key to a key that doesn't actually exist
	vm := newVM(tc, rg, networkInterface, secret)

	tc.CreateResourceAndWaitForState(vm, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal(conditions.ReasonSecretNotFound))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		ContainSubstring("Secret \"%s/%s\" does not contain key \"%s\"", tc.Namespace, secret.Name, secret.Key))

	// Delete VM and resources.
	tc.DeleteResourcesAndWait(vm, networkInterface, subnet, vnet, rg)
}

func Test_UserSecretInDifferentNamespace_SecretNotFound(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	secretInDiffNamespace := createNamespaceAndSecret(tc, "secret-namespace", tc.Namer.GenerateName("vmsecret"))

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)

	tc.CreateResourceAndWait(vnet)
	tc.CreateResourceAndWait(subnet)
	tc.CreateResourceAndWait(networkInterface)

	vm := newVM(tc, rg, networkInterface, secretInDiffNamespace)

	tc.CreateResourceAndWaitForState(vm, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal(conditions.ReasonSecretNotFound))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		ContainSubstring("failed resolving secret references: %s/%s does not exist", tc.Namespace, secretInDiffNamespace.Name))

	// Delete VM and resources.
	tc.DeleteResourcesAndWait(vm, networkInterface, subnet, vnet, rg)
}

func createNamespaceAndSecret(tc *testcommon.KubePerTestContext, namespaceName string, secretName string) genruntime.SecretReference {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}
	// Creating a namespace here using CreateResourceUntracked as we don't want to register clean up for namespace.
	tc.CreateResourceUntracked(ns)

	passwordKey := "password"
	secret := &v1.Secret{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      secretName,
			Namespace: ns.Name,
		},
		StringData: map[string]string{
			passwordKey: tc.Namer.GeneratePasswordOfLength(40),
		},
	}

	tc.CreateResource(secret)

	secretRef := genruntime.SecretReference{
		Name: secret.Name,
		Key:  passwordKey,
	}
	return secretRef
}
