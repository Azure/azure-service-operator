/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	. "github.com/onsi/gomega"
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
	secretName := tc.Namer.GenerateName("admin-pw")
	nameSpaceName := "secret-namespace"

	createNamespaces(tc, nameSpaceName)
	secretInDiffNamespace := createNamespacedSecret(tc, nameSpaceName, secretName)

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

func Test_UserSecretInDifferentNamespace_ShouldNotTriggerReconcile(t *testing.T) {
	t.Parallel()
	ns1 := "ns-1"
	ns2 := "ns-2"

	cfg, err := config.ReadFromEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	cfg.TargetNamespaces = []string{ns1, ns2}
	tc := globalTestContext.ForTestWithConfig(t, cfg)

	createNamespaces(tc, ns1, ns2)

	secretName := tc.Namer.GenerateName("server-secret")

	secretInCurrentNamespace := createNamespacedSecret(tc, ns1, secretName)
	secretInDiffNamespace := createNamespacedSecret(tc, ns2, secretName)

	rg := tc.NewTestResourceGroup()
	rg.Namespace = ns1
	tc.CreateResourceGroupAndWait(rg)

	rg2 := tc.NewTestResourceGroup()
	rg2.Namespace = ns2
	tc.CreateResourceGroupAndWait(rg2)

	server1, _ := newFlexibleServer(tc, testcommon.AsOwner(rg), secretInCurrentNamespace)
	server1.ObjectMeta = metav1.ObjectMeta{
		Name:      tc.Namer.GenerateName("mysql"),
		Namespace: ns1,
	}
	tc.CreateResourceAndWait(server1)

	resourceID := genruntime.GetResourceIDOrDefault(server1)
	tc.Expect(resourceID).ToNot(BeEmpty())

	// Deleting server1 here using AzureClient so that Operator does not know about the deletion
	_, err = tc.AzureClient.DeleteByID(tc.Ctx, resourceID, server1.GetAPIVersion())
	if err != nil {
		return
	}

	server2, _ := newFlexibleServer(tc, testcommon.AsOwner(rg2), secretInDiffNamespace)
	server2.ObjectMeta = metav1.ObjectMeta{
		Name:      tc.Namer.GenerateName("mysql2"),
		Namespace: ns2,
	}

	tc.CreateResourceAndWait(server2)

	// Here we want to make sure that the server1 we deleted in Azure should not be created again with reconcile of secret
	// in other namespace
	_, err = tc.AzureClient.GetByID(tc.Ctx, resourceID, server1.GetAPIVersion(), server1)
	tc.Expect(err).ToNot(BeNil())

	tc.DeleteResourcesAndWait(server1, server2)
	tc.DeleteResourcesAndWait(rg, rg2)

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
