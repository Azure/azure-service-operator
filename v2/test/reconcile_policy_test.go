/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// This test cannot be run in record/replay mode because the state it looks for at the beginning (Ready = false with warning)
// is not "stable" (the reconciler keeps retrying). Since it keeps retrying there isn't a deterministic number of
// retry attempts it makes which means a recording test may run out of recorded retries.
func Test_ReconcilePolicy_SkipReconcile_DoesntCreateResourceInAzure(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.NewTestResourceGroup()
	rg.Annotations = map[string]string{
		reconcilers.ReconcilePolicyAnnotation: string(reconcilers.ReconcilePolicySkip),
	}
	tc.CreateResourceAndWaitForState(rg, metav1.ConditionFalse, conditions.ConditionSeverityWarning)

	// We expect status to be empty
	tc.Expect(rg.Status.Name).To(BeNil())
	tc.Expect(rg.Status.Location).To(BeNil())

	// We expect the ready condition to include details of the error
	tc.Expect(rg.Status.Conditions[0].Reason).To(Equal("AzureResourceNotFound"))
	tc.Expect(rg.Status.Conditions[0].Message).To(ContainSubstring("ResourceGroupNotFound"))

	tc.Expect(rg.Annotations[genruntime.ResourceIDAnnotation]).ToNot(BeEmpty())

	// The actual Azure resource shouldn't exist
	armID := rg.Annotations[genruntime.ResourceIDAnnotation]
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armID,
		"2020-06-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())

	// Now when we remove the annotation the resource should move to ready state
	// Update the tags
	old := rg.DeepCopy()
	delete(rg.Annotations, reconcilers.ReconcilePolicyAnnotation)
	tc.Patch(old, rg)
	tc.Eventually(rg).Should(tc.Match.BeProvisioned(0))

	// We expect status to not be empty
	tc.Expect(rg.Status.Name).ToNot(BeNil())
	tc.Expect(rg.Status.Location).ToNot(BeNil())

	// The actual Azure resource should exist
	exists, _, err = tc.AzureClient.HeadByID(
		tc.Ctx,
		armID,
		"2020-06-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())
}
