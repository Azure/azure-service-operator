/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"
)

func Test_ManagedIdentity_ResourceCanBeCreated(t *testing.T) {
	t.Parallel()

	// The identity used by this test is always based on AAD Pod Identity, so all we need to do is
	// create and delete a resource group and ensure that it works.

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	tc.DeleteResourceAndWait(rg)
}
