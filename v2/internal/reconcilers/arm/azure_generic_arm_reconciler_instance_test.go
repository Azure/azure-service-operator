/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"net/http"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

func TestAzureDeploymentReconcilerInstance_CheckSubscription_GivenResourceID_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	// Test cases can use the following substitutions for clarity:
	// {expectedSub} - a subscription ID that matches that on the credential used for the check
	// {unexpectedSub} - a subscription ID that does not match that on the credential used for the check
	//
	cases := map[string]struct {
		resourceID             string
		expectedErrorSubstring string
	}{
		"resource group in expected subscription": {
			resourceID: "/subscriptions/{expectedSub}/resourceGroups/asotest-rg",
		},
		"resource group in expected subscription (lowercase)": {
			resourceID: "/subscriptions/{expectedSubLower}/resourceGroups/asotest-rg",
		},
		"resource group in different subscription": {
			resourceID:             "/subscriptions/{unexpectedSub}/resourceGroups/asotest-rg",
			expectedErrorSubstring: "resource does not match with Client Credential",
		},
		"Subscription aliases have no subscription ID to check, so should always succeed": {
			resourceID: "/providers/Microsoft.Subscription/aliases/asotest-sub-zzqblm",
		},
		"Role assignment to resource group with matching subscription": {
			resourceID: "/subscriptions/{expectedSub}/resourceGroups/asotest-rg-mnlufu/providers/Microsoft.Authorization/roleAssignments/ed3c9f84-ae20-5aea-8602-58ed2eeffef1",
		},
		"Role assignment to resource group with non-matching subscription": {
			resourceID:             "/subscriptions/{unexpectedSub}/resourceGroups/asotest-rg/providers/Microsoft.Authorization/roleAssignments/asotest-role-assignment",
			expectedErrorSubstring: "resource does not match with Client Credential",
		},
		"Role assignment to subscription with matching subscription": {
			resourceID: "/subscriptions/{expectedSub}/providers/Microsoft.Authorization/roleAssignments/ed3c9f84-ae20-5aea-8602-58ed2eeffef1",
		},
		"Role assignment to subscription with non-matching subscription": {
			resourceID:             "/subscriptions/{unexpectedSub}/providers/Microsoft.Authorization/roleAssignments/asotest-role-assignment",
			expectedErrorSubstring: "resource does not match with Client Credential",
		},
	}

	expectedSub := "00000000-0000-DEAD-C0DE-000000000000"
	expectedSubLower := "00000000-0000-dead-c0de-000000000000"
	unexpectedSub := "11111111-1111-1111-1111-111111111111"

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			reconciler := &azureDeploymentReconcilerInstance{
				ARMConnection: &fakeConnection{
					subscriptionID: expectedSub,
				},
			}

			// Replace the placeholders in the resource ID with the actual subscription IDs
			resourceID := c.resourceID
			resourceID = strings.ReplaceAll(resourceID, "{expectedSub}", expectedSub)
			resourceID = strings.ReplaceAll(resourceID, "{unexpectedSub}", unexpectedSub)
			resourceID = strings.ReplaceAll(resourceID, "{expectedSubLower}", expectedSubLower)

			err := reconciler.checkSubscription(resourceID)
			if c.expectedErrorSubstring == "" {
				g.Expect(err).ToNot(HaveOccurred())
			} else {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedErrorSubstring)))
			}
		})
	}
}

func TestAzureDeploymentReconcilerInstance_HandleDeleteFailed_NormalizesVerySlowRetry(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	scheme := runtime.NewScheme()
	g.Expect(resources.AddToScheme(scheme)).To(Succeed())

	obj := &resources.ResourceGroup{}
	obj.SetGroupVersionKind(resources.GroupVersion.WithKind("ResourceGroup"))
	reconciler := &azureDeploymentReconcilerInstance{
		ARMOwnedResourceReconcilerCommon: reconcilers.ARMOwnedResourceReconcilerCommon{
			ResourceResolver: resolver.NewResolver(NewFakeKubeClient(scheme)),
		},
		Obj: obj,
		Log: logr.Discard(),
	}

	http400Error := genericarmclient.NewCloudError(&azcore.ResponseError{
		StatusCode:  http.StatusBadRequest,
		RawResponse: &http.Response{Header: make(http.Header)},
	})

	deleteErr := reconciler.handleDeleteFailed(http400Error)
	readyDeleteErr, ok := conditions.AsReadyConditionImpactingError(deleteErr)
	g.Expect(ok).To(BeTrue())
	g.Expect(readyDeleteErr.Severity).To(Equal(conditions.ConditionSeverityWarning))
	g.Expect(readyDeleteErr.RetryClassification).To(Equal(retry.Slow))
}

type fakeConnection struct {
	subscriptionID string
	credentialFrom types.NamespacedName
}

var _ Connection = &fakeConnection{}

// Client implements [Connection].
func (f *fakeConnection) Client() *genericarmclient.GenericClient {
	panic("unimplemented")
}

// CredentialFrom implements [Connection].
func (f *fakeConnection) CredentialFrom() types.NamespacedName {
	return f.credentialFrom
}

// SubscriptionID implements [Connection].
func (f *fakeConnection) SubscriptionID() string {
	return f.subscriptionID
}
