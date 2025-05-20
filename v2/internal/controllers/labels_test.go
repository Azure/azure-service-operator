/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/common/labels"
)

func Test_ExpectedLabels_AreAdded(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a storage account
	acct := newStorageAccount(tc, rg)

	tc.CreateResourcesAndWait(acct)
	tc.Expect(acct.Labels).To(HaveKeyWithValue(labels.OwnerNameLabel, rg.Name))
	tc.Expect(acct.Labels).To(HaveKeyWithValue(labels.OwnerGroupKindLabel, "ResourceGroup.resources.azure.com"))
	tc.Expect(acct.Labels).To(HaveKeyWithValue(labels.OwnerUIDLabel, string(rg.UID)))
}
