/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func waitForOwnerMissingError(tc *testcommon.KubePerTestContext, obj client.Object) {
	objectKey := client.ObjectKeyFromObject(obj)

	tc.Eventually(func() string {
		tc.GetResource(objectKey, obj)
		conditioner := obj.(conditions.Conditioner)
		ready, ok := conditions.GetCondition(conditioner, conditions.ConditionTypeReady)
		if !ok {
			return ""
		}

		return ready.Reason
	}).Should(Equal("WaitingForOwner"))
}

func storageAccountAndResourceGroupProvisionedOutOfOrderHelper(t *testing.T, waitHelper func(tc *testcommon.KubePerTestContext, obj client.Object)) {
	tc := globalTestContext.ForTest(t)

	// Create the resource group in-memory but don't submit it yet
	rg := tc.NewTestResourceGroup()

	acct := newStorageAccount(tc, rg)

	// Create the storage account - initially this will not succeed, but it should keep trying
	tc.CreateResource(acct)

	waitHelper(tc, acct)

	// The resource group should be created successfully
	tc.CreateResource(rg)

	// The storage account should also be created successfully
	tc.Eventually(acct).Should(tc.Match.BeProvisioned(0))
}

func subnetAndVNETCreatedProvisionedOutOfOrder(t *testing.T, waitHelper func(tc *testcommon.KubePerTestContext, obj client.Object)) {
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetwork_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}

	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworks_Subnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}

	// Create the subnet - initially this will not succeed, but it should keep trying
	tc.CreateResource(subnet)

	waitHelper(tc, subnet)

	// Now created the vnet
	tc.CreateResourceAndWait(vnet)
	// The subnet account should also be created successfully eventually
	tc.Eventually(subnet).Should(tc.Match.BeProvisioned(0))
}

func Test_StorageAccount_CreatedBeforeResourceGroup(t *testing.T) {
	t.Parallel()
	storageAccountAndResourceGroupProvisionedOutOfOrderHelper(t, waitForOwnerMissingError)
}

func Test_StorageAccount_CreatedInParallelWithResourceGroup(t *testing.T) {
	t.Parallel()
	doNotWait := func(_ *testcommon.KubePerTestContext, _ client.Object) { /* do not wait */ }
	storageAccountAndResourceGroupProvisionedOutOfOrderHelper(t, doNotWait)
}

func Test_Subnet_CreatedBeforeVNET(t *testing.T) {
	t.Parallel()
	subnetAndVNETCreatedProvisionedOutOfOrder(t, waitForOwnerMissingError)
}

func Test_CreateResourceGroupThatAlreadyExists_ReconcilesSuccessfully(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Create another resource group that points to the same Azure resource
	rgCopy := &resources.ResourceGroup{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      rg.Name,
			Namespace: rg.Namespace,
		},
		Spec: *rg.Spec.DeepCopy(),
	}
	rgCopy.Spec.AzureName = rgCopy.Name
	rgCopy.Name = rgCopy.Name + "duplicate"

	tc.CreateResourceAndWait(rgCopy)
}

func Test_CreateStorageAccountThatAlreadyExists_ReconcilesSuccessfully(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	acct := newStorageAccount(tc, rg)

	acctCopy := acct.DeepCopy()

	tc.CreateResourcesAndWait(acct)

	// Patch the account to remove the finalizer
	old := acct.DeepCopy()
	controllerutil.RemoveFinalizer(acct, "serviceoperator.azure.com/finalizer")
	tc.Patch(old, acct)

	// Delete the account
	tc.DeleteResourceAndWait(acct)

	// Create it again
	tc.CreateResourcesAndWait(acctCopy)
}

func Test_CreateStorageAccountWithoutRequiredProperties_Rejected(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	acct := newStorageAccount(tc, rg)

	acctCopy := acct.DeepCopy()

	tc.CreateResourcesAndWait(acct)

	// Patch the account to remove the finalizer
	old := acct.DeepCopy()
	controllerutil.RemoveFinalizer(acct, "serviceoperator.azure.com/finalizer")
	tc.Patch(old, acct)

	// Delete the account
	tc.DeleteResourceAndWait(acct)

	// Create it again
	tc.CreateResourcesAndWait(acctCopy)
}

func Test_AzureName_IsImmutableOnceSuccessfullyCreated(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	acct := newStorageAccount(tc, rg)
	tc.CreateResourcesAndWait(acct)

	// Patch the account to change AzureName
	newAzureName := "test123"
	old := acct.DeepCopy()
	acct.Spec.AzureName = newAzureName
	err := tc.PatchAndExpectError(old, acct)

	tc.Expect(err).ToNot(BeNil())
	tc.Expect(old.Spec.AzureName).ToNot(BeIdenticalTo(newAzureName))

	// Delete the account
	tc.DeleteResourceAndWait(acct)
}

func Test_Owner_IsImmutableOnceSuccessfullyCreated(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	acct := newStorageAccount(tc, rg)
	tc.CreateResourcesAndWait(acct)

	rg2 := tc.CreateTestResourceGroupAndWait()

	// Patch the account to change Owner
	old := acct.DeepCopy()
	acct.Spec.Owner = testcommon.AsOwner(rg2)
	err := tc.PatchAndExpectError(old, acct)

	tc.Expect(err).To(MatchError(ContainSubstring("updating 'spec.owner.name' is not allowed")))
	tc.Expect(old.Owner().Name).ToNot(BeIdenticalTo(rg2.Name))

	// Delete the account
	tc.DeleteResourceAndWait(acct)
}

func Test_OwnerARMID_IsImmutableOnceSuccessfullyCreated(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Ensure that the RG has an ARM ID set
	tc.Expect(rg.Status.Id).ToNot(BeNil())
	tc.Expect(to.Value(rg.Status.Id)).ToNot(BeEmpty())

	acct := newStorageAccount(tc, rg)
	// Manually set the ARM ID of the owner:
	acct.Spec.Owner.Name = ""
	acct.Spec.Owner.ARMID = to.Value(rg.Status.Id)
	tc.CreateResourcesAndWait(acct)

	rg2 := tc.CreateTestResourceGroupAndWait()

	// Ensure that the RG has an ARM ID set
	tc.Expect(rg2.Status.Id).ToNot(BeNil())
	tc.Expect(to.Value(rg2.Status.Id)).ToNot(BeEmpty())

	// Patch the account to change Owner
	old := acct.DeepCopy()
	acct.Spec.Owner.ARMID = to.Value(rg2.Status.Id)
	err := tc.PatchAndExpectError(old, acct)

	tc.Expect(err).To(MatchError(ContainSubstring("updating 'spec.owner.armId' is not allowed")))
	tc.Expect(old.Owner().Name).ToNot(BeIdenticalTo(rg2.Name))

	// Delete the account
	tc.DeleteResourceAndWait(acct)
}

func Test_AzureName_IsImmutable_IfAzureHasBeenCommunicatedWith(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	invalidAzureName := "==--039+"

	acct := newStorageAccount(tc, rg)
	acct.Spec.AzureName = invalidAzureName
	tc.CreateResourceAndWaitForFailure(acct)

	// Patch the account to change AzureName
	old := acct.DeepCopy()
	acct.Spec.AzureName = tc.NoSpaceNamer.GenerateName("stor")
	err := tc.PatchAndExpectError(old, acct)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("updating 'spec.azureName' is not allowed"))

	// Delete the account
	tc.DeleteResourceAndWait(acct)
}

func Test_Owner_IsMutableIfNotSuccessfullyCreated(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	invalidOwnerName := "test123"
	actualOwnerName := rg.Name
	rg.Name = invalidOwnerName

	acct := newStorageAccount(tc, rg)
	tc.CreateResourceAndWaitForState(acct, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// TODO: We have hack in here to skip re-reconcile here as to avoid race between requeue and patch,
	// TODO: which ends up in duplicate PUTs while replaying.
	// Patch the account to skip reconcile
	old := acct.DeepCopy()
	acct.Annotations = make(map[string]string)
	acct.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	tc.Patch(old, acct)

	// Patch the account to change Owner's name
	old = acct.DeepCopy()
	acct.Spec.Owner.Name = actualOwnerName
	delete(acct.Annotations, "serviceoperator.azure.com/reconcile-policy")
	tc.PatchResourceAndWait(old, acct)

	tc.Expect(acct.Owner().Name).ToNot(BeIdenticalTo(invalidOwnerName))
	// Delete the account
	tc.DeleteResourceAndWait(acct)
}

func Test_CreateCosmosAccountWithSkipReconcile_SecretsAreWritten(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	cosmosSecret1 := "keys1"

	// Custom namer because cosmosdb accounts have stricter name
	// requirements - no hyphens allowed.
	// Create a Cosmos DB account
	offerType := documentdb.DatabaseAccountOfferType_Standard
	kind := documentdb.DatabaseAccount_Kind_Spec_GlobalDocumentDB
	acct := &documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("sqlacct")),
		Spec: documentdb.DatabaseAccount_Spec{
			Location:                 tc.AzureRegion,
			Owner:                    testcommon.AsOwner(rg),
			Kind:                     &kind,
			DatabaseAccountOfferType: &offerType,
			Locations: []documentdb.Location{
				{
					LocationName: tc.AzureRegion,
				},
			},
			OperatorSpec: &documentdb.DatabaseAccountOperatorSpec{
				Secrets: &documentdb.DatabaseAccountOperatorSecrets{
					DocumentEndpoint: &genruntime.SecretDestination{
						Name: cosmosSecret1,
						Key:  "endpoint",
					},
				},
			},
		},
	}

	cosmosSecret2 := "keys2"
	skipAcct := acct.DeepCopy()
	skipAcct.Spec.AzureName = skipAcct.Name
	skipAcct.Name = skipAcct.Name + "-skip" // So we don't collide
	skipAcct.Spec.OperatorSpec.Secrets.DocumentEndpoint.Name = cosmosSecret2
	skipAcct.Annotations = map[string]string{
		annotations.ReconcilePolicy: string(annotations.ReconcilePolicySkip),
	}

	tc.CreateResourcesAndWait(acct, skipAcct)

	tc.ExpectSecretHasKeys(cosmosSecret1, "endpoint")
	tc.ExpectSecretHasKeys(cosmosSecret2, "endpoint")
}
