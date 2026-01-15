/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func newStorageAccountWithInvalidKeyExpiration(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *storage.StorageAccount {
	// Custom namer because storage accounts have strict names

	// Create a storage account with an invalid key expiration period
	accessTier := storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot
	kind := storage.StorageAccount_Kind_Spec_BlobStorage
	sku := storage.SkuName_Standard_LRS
	return &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Sku: &storage.Sku{
				Name: &sku,
			},
			AccessTier: &accessTier,
			KeyPolicy: &storage.KeyPolicy{
				KeyExpirationPeriodInDays: to.Ptr(-260),
			},
		},
	}
}

// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long-running operation polling. This ensures that the second case is handled correctly.
func Test_OperationAccepted_LongRunningOperationFails(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	acct := newStorageAccountWithInvalidKeyExpiration(tc, rg)
	tc.CreateResourceAndWaitForFailure(acct)

	ready, ok := conditions.GetCondition(acct, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())

	tc.Expect(ready.Status).To(Equal(metav1.ConditionFalse))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityError))
	tc.Expect(ready.Reason).To(Equal("InvalidValuesForRequestParameters"))
	tc.Expect(ready.Message).To(ContainSubstring("Values for request parameters are invalid: keyPolicy.keyExpirationPeriodInDays."))
}

// There are two ways that a long-running operation can fail. It can be rejected when initially
// submitted to the Azure API, or it can be accepted and then report a failure during
// long-running operation polling. This ensures that a resource in the second case
// can be updated to resolve the cause of the failure and successfully deployed.
func Test_OperationAccepted_LongRunningOperationFails_SucceedsAfterUpdate(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	acct := newStorageAccountWithInvalidKeyExpiration(tc, rg)
	tc.CreateResourceAndWaitForFailure(acct)

	// Remove the bad property and ensure we can successfully provision
	old := acct.DeepCopy()
	acct.Spec.KeyPolicy = nil
	tc.PatchResourceAndWait(old, acct)

	// Ensure that the old failure information was cleared away
	objectKey := client.ObjectKeyFromObject(acct)
	updated := &storage.StorageAccount{}
	tc.GetResource(objectKey, updated)

	ready, ok := conditions.GetCondition(acct, conditions.ConditionTypeReady)
	tc.Expect(ok).To(BeTrue())

	tc.Expect(ready.Status).To(Equal(metav1.ConditionTrue))
	tc.Expect(ready.Severity).To(Equal(conditions.ConditionSeverityNone))
	tc.Expect(ready.Reason).To(Equal(conditions.ReasonSucceeded))
	tc.Expect(ready.Message).To(Equal(""))
}
