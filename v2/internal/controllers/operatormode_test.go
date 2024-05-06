// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func TestOperatorMode_Webhooks(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTestWithConfig(t, config.Values{
		OperatorMode: config.OperatorModeWebhooks,
	})

	rg := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: tc.Namespace,
		},
		Spec: resources.ResourceGroup_Spec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	tc.Expect(rg.Spec.AzureName).To(Equal(""))

	tc.CreateResource(&rg)
	// AzureName should have been defaulted on the group on the
	// way in (it doesn't require waiting for a reconcile).
	tc.Expect(rg.Spec.AzureName).To(Equal(rg.ObjectMeta.Name))

	checkNeverGetsFinalizer(tc, &rg,
		"instance got a finalizer when operator mode is webhooks")
}

func TestOperatorMode_Watchers(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTestWithConfig(t, config.Values{
		OperatorMode: config.OperatorModeWatchers,
	})

	rg := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: tc.Namespace,
		},
		Spec: resources.ResourceGroup_Spec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	tc.Expect(rg.Spec.AzureName).To(Equal(""))

	err := tc.CreateResourceExpectRequestFailure(&rg)
	// We should fail because the webhook isn't registered (in a real
	// multi-operator deployment it would be routed to a different
	// operator running in webhook-only mode).
	tc.Expect(err.Error()).To(MatchRegexp(`failed calling webhook .* connection refused`))

	// Nothing else to check since we can't create the resource
	// without a webhook.
}

func TestOperatorMode_Both(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTestWithConfig(t, config.Values{
		OperatorMode: config.OperatorModeBoth,
	})

	rg := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: tc.Namespace,
		},
		Spec: resources.ResourceGroup_Spec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	tc.Expect(rg.Spec.AzureName).To(Equal(""))

	tc.CreateResource(&rg)

	// AzureName should have been defaulted on the group on the
	// way in (it doesn't require waiting for a reconcile).
	tc.Expect(rg.Spec.AzureName).To(Equal(rg.ObjectMeta.Name))
	tc.Eventually(&rg).Should(tc.Match.BeProvisioned(0))
}
