// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"os"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/config"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
	resources "github.com/Azure/azure-service-operator/v2/api/microsoft.resources/v1alpha1api20200601"

	. "github.com/onsi/gomega"
)

func TestOperatorModeWebhooks(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	envVal := os.Getenv("AZURE_OPERATOR_MODE")
	if envVal == "" {
		envVal = config.OperatorModeBoth.String()
	}
	operatorMode, err := config.ParseOperatorMode(envVal)
	tc.Expect(err).NotTo(HaveOccurred())

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

	// Check assumptions.
	tc.Expect(rg.Spec.AzureName).To(Equal(""))

	_, err = tc.CreateResourceGroup(&rg)
	if operatorMode.IncludesWebhooks() {
		// AzureName should have been defaulted on the group on the
		// way in (it doesn't require waiting for a reconcile).
		tc.Expect(err).ToNot(HaveOccurred())
		// TODO: this works here, but it's possibly because envtest
		// has the apiserver running in-process - I'm not sure that
		// the client would reflect changes made by webhooks in a
		// remote cluster. Does this test need to work if it's run
		// against a non-envtest cluster as well?
		tc.Expect(rg.Spec.AzureName).To(Equal(rg.ObjectMeta.Name))
	} else {
		// Otherwise we should fail because the webhook isn't
		// registered (in a real multi-operator deployment it would be
		// routed to a different operator running in webhook-only
		// mode).
		tc.Expect(err.Error()).To(MatchRegexp(`failed calling webhook .* connection refused`))
	}
}

func TestOperatorModeWatchers(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	envVal := os.Getenv("AZURE_OPERATOR_MODE")
	if envVal == "" {
		envVal = config.OperatorModeBoth.String()
	}
	operatorMode, err := config.ParseOperatorMode(envVal)
	tc.Expect(err).NotTo(HaveOccurred())

	if !operatorMode.IncludesWebhooks() {
		// If we can't create a resource then we can't confirm whether
		// it gets reconciled, and unlike in v1 there aren't any
		// resources without at least the webhook to default
		// AzureName.
		t.Skip("without webhooks we can't check watchers behaviour")
	}

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

	_, err = tc.CreateResourceGroup(&rg)
	tc.Expect(err).NotTo(HaveOccurred())

	name := types.NamespacedName{
		Name:      rg.ObjectMeta.Name,
		Namespace: rg.ObjectMeta.Namespace,
	}

	hasFinalizer := gotFinalizer(tc, name)
	if operatorMode.IncludesWatchers() {
		tc.G.Eventually(
			hasFinalizer,
			timeoutFast,
			retry,
		).Should(
			BeTrue(),
			"instance never got finalizer even though operator mode is %q",
			operatorMode,
		)
	} else {
		tc.G.Consistently(
			hasFinalizer,
			20*time.Second,
			time.Second,
		).Should(
			BeFalse(),
			"instance got a finalizer when operator mode is %q",
			operatorMode,
		)
	}
}
