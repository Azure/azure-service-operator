/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/creds"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

func Test_OwnerIsARMIDSetWithName_Rejected(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Get the rg's ARM ID
	tc.Expect(rg.Status.Id).ToNot(BeNil())
	armID := *rg.Status.Id

	// Now create a storage account
	acct := newStorageAccount20230101(tc, rg)
	acct.Spec.Owner.ARMID = armID

	// Create the storage account from ARM ID
	err := tc.CreateResourceExpectRequestFailure(acct)
	tc.Expect(err).To(MatchError(ContainSubstring("the 'ARMID' field is mutually exclusive with 'Group', 'Kind', 'Namespace', and 'Name'")))
}

func Test_OwnerIsARMIDOfWrongType_Rejected(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()
	badARMID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/myvm", tc.AzureSubscription, "nonexistrg")

	// Now create a storage account
	acct := newStorageAccount20230101(tc, rg)
	acct.Spec.Owner = testcommon.AsARMIDOwner(badARMID)

	// Create the storage account from ARM ID
	err := tc.CreateResourceExpectRequestFailure(acct)
	tc.Expect(err).To(MatchError(ContainSubstring("expected owner ARM ID to be for a resource group, but was \"Microsoft.Compute/virtualMachines\"")))
}

func Test_OwnerIsKubernetesResourceFromDifferentSubscription_ResourceFails(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	cfg, err := testcommon.ReadFromEnvironmentForTest()
	tc.Expect(err).ToNot(HaveOccurred())

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	tc.Expect(rg.Status.Id).ToNot(BeNil())

	scopedCredentialName := "other-subscription-secret"

	// Now create a storage account
	acct := newStorageAccount20230101(tc, rg)
	acct.Annotations = map[string]string{annotations.PerResourceSecret: scopedCredentialName}

	uuid, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	scopedCredentialSecret := creds.NewScopedServicePrincipalSecret(
		uuid.String(),
		tc.AzureTenant,
		cfg.ClientID,
		"1234", // We're not expecting to make it all the way to provision the resource, so it's OK to use a fake client secret here
		scopedCredentialName,
		tc.Namespace)
	tc.CreateResource(scopedCredentialSecret)

	tc.CreateResourceAndWaitForFailure(acct)
}
