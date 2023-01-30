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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Note: This file is very similar to to_azure_configmaps_test.go. The same cases should be covered there as well.

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
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal(conditions.ReasonSecretNotFound.Name))
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
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal(conditions.ReasonSecretNotFound.Name))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		ContainSubstring("Secret \"%s/%s\" does not contain key \"%s\"", tc.Namespace, secret.Name, secret.Key))

	// Delete VM and resources.
	tc.DeleteResourceAndWait(rg)
}

func Test_UserSecretInDifferentNamespace_SecretNotFound(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	secretName := tc.Namer.GenerateName("admin-pw")
	namespaceName := "secret-namespace"

	err := tc.CreateTestNamespaces(namespaceName)
	tc.Expect(err).ToNot(HaveOccurred())

	secretInDiffNamespace := createNamespacedSecret(tc, namespaceName, secretName)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)

	tc.CreateResourcesAndWait(vnet, subnet, networkInterface)

	vm := newVirtualMachine20201201(tc, rg, networkInterface, secretInDiffNamespace)

	tc.CreateResourceAndWaitForState(vm, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	// We expect the ready condition to include details of the error
	tc.Expect(vm.Status.Conditions[0].Reason).To(Equal(conditions.ReasonSecretNotFound.Name))
	tc.Expect(vm.Status.Conditions[0].Message).To(
		ContainSubstring("failed resolving secret references: %s/%s does not exist", tc.Namespace, secretInDiffNamespace.Name))

	// Delete VM and resources.
	tc.DeleteResourcesAndWait(vm, networkInterface, subnet, vnet, rg)
}

func Test_UserSecretInDifferentNamespace_ShouldNotTriggerReconcile(t *testing.T) {
	t.Parallel()
	ns1 := "secret-in-different-namespace-ns-1"
	ns2 := "secret-in-different-namespace-ns-2"

	cfg, err := testcommon.ReadFromEnvironmentForTest()
	if err != nil {
		t.Fatal(err)
	}

	// TargetNamespaces is only used here for logger mapping in namespaces, and not multi-tenancy.
	cfg.TargetNamespaces = []string{ns1, ns2}
	tc := globalTestContext.ForTestWithConfig(t, cfg)

	err = tc.CreateTestNamespaces(ns1, ns2)
	tc.Expect(err).ToNot(HaveOccurred())

	secretName := tc.Namer.GenerateName("vm-secret")
	ns1Secret := createNamespacedSecret(tc, ns1, secretName)
	ns2Secret := createNamespacedSecret(tc, ns2, secretName)

	rg := tc.NewTestResourceGroup()
	rg.Namespace = ns1
	tc.CreateResourceGroupAndWait(rg)

	rg2 := tc.NewTestResourceGroup()
	rg2.Namespace = ns2
	tc.CreateResourceGroupAndWait(rg2)

	// VM1 with same secret name in ns1
	vm1 := createNamespacedVM(tc, rg, ns1, ns1Secret)

	resourceID := genruntime.GetResourceIDOrDefault(vm1)
	tc.Expect(resourceID).ToNot(BeEmpty())

	// Deleting server1 here using AzureClient so that Operator does not know about the deletion.
	// If operator reconciles that resource again it will be recreated, and we will notice
	resp, err := tc.AzureClient.BeginDeleteByID(tc.Ctx, resourceID, vm1.GetAPIVersion())
	tc.Expect(err).ToNot(HaveOccurred())
	_, err = resp.Poller.PollUntilDone(tc.Ctx, nil)
	tc.Expect(err).ToNot(HaveOccurred())

	// VM2 with same secret name in ns2
	_ = createNamespacedVM(tc, rg2, ns2, ns2Secret)

	// Here we want to make sure that the server1 we deleted from Azure was not created again by a reconcile triggered
	// when the second secret was created in other namespace
	_, err = tc.AzureClient.GetByID(tc.Ctx, resourceID, vm1.GetAPIVersion(), vm1)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(genericarmclient.IsNotFoundError(err)).To(BeTrue())

	tc.GetResource(client.ObjectKeyFromObject(vm1), vm1)
	tc.Expect(vm1.Status.Conditions).ToNot(HaveLen(0))
	tc.Expect(vm1.Status.Conditions[0].Type).To(BeEquivalentTo(conditions.ConditionTypeReady))
	tc.Expect(vm1.Status.Conditions[0].Status).To(BeEquivalentTo(metav1.ConditionTrue))
	tc.DeleteResourcesAndWait(rg, rg2)
}

func createNamespacedVM(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, ns string, ns1Secret genruntime.SecretReference) *v1beta20201201.VirtualMachine {
	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	vnet.Namespace = ns
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	subnet.Namespace = ns
	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)
	networkInterface.Namespace = ns

	tc.CreateResourcesAndWait(vnet, subnet, networkInterface)

	vm1 := newVirtualMachine20201201(tc, rg, networkInterface, ns1Secret)
	vm1.Namespace = ns
	tc.CreateResourceAndWait(vm1)
	return vm1
}

func createNamespacedSecret(tc *testcommon.KubePerTestContext, namespaceName string, secretName string) genruntime.SecretReference {
	passwordKey := "password"
	secret := &v1.Secret{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      secretName,
			Namespace: namespaceName,
		},
		StringData: map[string]string{
			passwordKey: tc.Namer.GeneratePasswordOfLength(40),
		},
	}

	tc.CreateResource(secret)

	return genruntime.SecretReference{
		Name: secret.Name,
		Key:  passwordKey,
	}
}
