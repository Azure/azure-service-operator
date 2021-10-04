// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	resources "github.com/Azure/azure-service-operator/v2/api/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/controller/config"
	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"

	. "github.com/onsi/gomega"
)

func TestOperatorMode_Webhooks(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTestWithConfig(t, config.Values{
		OperatorMode: config.OperatorModeWebhooks,
	})

	rg := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: tc.Namespace(),
		},
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	tc.Expect(rg.Spec.AzureName).To(Equal(""))

	_, err := tc.CreateResourceGroup(&rg)
	tc.Expect(err).ToNot(HaveOccurred())
	// AzureName should have been defaulted on the group on the
	// way in (it doesn't require waiting for a reconcile).
	// TODO: this works here, but it's possibly because envtest
	// has the apiserver running in-process - I'm not sure that
	// the client would reflect changes made by webhooks in a
	// remote cluster. Does this test need to work if it's run
	// against a non-envtest cluster as well?
	tc.Expect(rg.Spec.AzureName).To(Equal(rg.ObjectMeta.Name))

	hasFinalizer := gotFinalizer(tc, types.NamespacedName{
		Name:      rg.ObjectMeta.Name,
		Namespace: rg.ObjectMeta.Namespace,
	})

	tc.G.Consistently(
		hasFinalizer,
		20*time.Second,
		time.Second,
	).Should(
		BeFalse(),
		"instance got a finalizer when operator mode is webhooks",
	)
}

func TestOperatorMode_Watchers(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTestWithConfig(t, config.Values{
		OperatorMode: config.OperatorModeWatchers,
	})

	rg := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: tc.Namespace(),
		},
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	tc.Expect(rg.Spec.AzureName).To(Equal(""))

	_, err := tc.CreateResourceGroup(&rg)
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
			Namespace: tc.Namespace(),
		},
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	tc.Expect(rg.Spec.AzureName).To(Equal(""))

	_, err := tc.CreateResourceGroup(&rg)
	tc.Expect(err).NotTo(HaveOccurred())

	// AzureName should have been defaulted on the group on the
	// way in (it doesn't require waiting for a reconcile).
	// TODO: this works here, but it's possibly because envtest
	// has the apiserver running in-process - I'm not sure that
	// the client would reflect changes made by webhooks in a
	// remote cluster. Does this test need to work if it's run
	// against a non-envtest cluster as well?
	tc.Expect(rg.Spec.AzureName).To(Equal(rg.ObjectMeta.Name))

	hasFinalizer := gotFinalizer(tc, types.NamespacedName{
		Name:      rg.ObjectMeta.Name,
		Namespace: rg.ObjectMeta.Namespace,
	})
	tc.G.Eventually(
		hasFinalizer,
		timeoutFast,
		retry,
	).Should(
		BeTrue(),
		"instance never got finalizer even though operator mode is watchers-and-webhooks",
	)
	// Wait for resource group to be provisioned, otherwise we
	// might get no requests to ARM, and then trying to run the
	// tests without credentials fails because there's no
	// recording.
	tc.Eventually(&rg).Should(tc.Match.BeProvisioned())
}
