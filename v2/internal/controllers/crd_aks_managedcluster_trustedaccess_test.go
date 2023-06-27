/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	// The testing package is imported for testing-related functionality.
	"testing"

	// The gomega package is used for assertions and expectations in tests.
	. "github.com/onsi/gomega"

	// The dataprotection package contains types and functions related to dataprotection resources.
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20230101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	// The testcommon package includes common testing utilities.
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	// The to package includes utilities for converting values to pointers.
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ManagedCluster_TrustedAccess(t *testing.T) {

	// indicates that this test function can run in parallel with other tests
	t.Parallel()

	// Create a test resource group and wait until the operation is completed, where the globalTestContext is a global object that provides the necessary context and utilities for testing.
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Creation of Backup Vault
	backupVault := newBackupVault(tc, rg, "asotestbackupvault")

	// Creation of Backup Policy
	backupPolicy := newBackupPolicy(tc, backupVault, "asotestbackuppolicy")
	
	// Creation of AKS Managed Cluster
	

}

// Creation of AKS Managed Cluster Trusted Access
