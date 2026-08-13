/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

func Test_ReconcilePolicyForAnnotation_givenAnnotation_resolvesItAsTheReconcilerWould(t *testing.T) {
	t.Parallel()

	policies := annotations.ReconcilePolicies{
		// Belongs to the resource being reconciled, and must not decide another resource's policy
		Effective: annotations.ReconcilePolicyDetachOnDelete,
		Inherited: annotations.ReconcilePolicySkip,
		Default:   annotations.ReconcilePolicyManage,
	}

	cases := map[string]struct {
		annotation string
		expected   annotations.ReconcilePolicyValue
	}{
		"No annotation takes what the namespace and the operator say": {
			expected: annotations.ReconcilePolicySkip,
		},
		"An annotation is taken at its word": {
			annotation: string(annotations.ReconcilePolicyManage),
			expected:   annotations.ReconcilePolicyManage,
		},
		"So is one that forbids modification": {
			annotation: string(annotations.ReconcilePolicySkip),
			expected:   annotations.ReconcilePolicySkip,
		},
		// A reconciler parses a non-empty annotation against the operator's policy, never the namespace's
		"An unusable annotation falls back to the operator's policy, not the namespace's": {
			annotation: "nonsense",
			expected:   annotations.ReconcilePolicyManage,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctx := reconcilers.WithReconcilePolicies(context.Background(), policies)

			g.Expect(reconcilers.ReconcilePolicyForAnnotation(ctx, c.annotation)).To(Equal(c.expected))
		})
	}
}

// The reverse of the case above: an operator configured to leave things alone must not have that
// overridden by a namespace that permits management
func Test_ReconcilePolicyForAnnotation_givenUnusableAnnotation_doesNotOverrideASkippingOperator(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ctx := reconcilers.WithReconcilePolicies(context.Background(), annotations.ReconcilePolicies{
		Effective: annotations.ReconcilePolicyManage,
		Inherited: annotations.ReconcilePolicyManage,
		Default:   annotations.ReconcilePolicySkip,
	})

	g.Expect(reconcilers.ReconcilePolicyForAnnotation(ctx, "nonsense")).
		To(Equal(annotations.ReconcilePolicySkip))
}

func Test_ReconcilePolicyFromContext_givenPolicies_returnsTheEffectiveOne(t *testing.T) {
	t.Parallel()

	values := []annotations.ReconcilePolicyValue{
		annotations.ReconcilePolicyManage,
		annotations.ReconcilePolicySkip,
		annotations.ReconcilePolicyDetachOnDelete,
	}

	for _, effective := range values {
		t.Run(string(effective), func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctx := reconcilers.WithReconcilePolicies(context.Background(), annotations.ReconcilePolicies{
				Effective: effective,
				Inherited: annotations.ReconcilePolicySkip,
				Default:   annotations.ReconcilePolicyManage,
			})

			g.Expect(reconcilers.ReconcilePolicyFromContext(ctx)).To(Equal(effective))
		})
	}
}

func Test_ReconcilePolicies_givenNoPolicies_areManage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Nothing outside of a reconcile records a policy, and the operator's own default is to manage
	g.Expect(reconcilers.ReconcilePolicyFromContext(context.Background())).
		To(Equal(annotations.ReconcilePolicyManage))
	g.Expect(reconcilers.ReconcilePolicyForAnnotation(context.Background(), "")).
		To(Equal(annotations.ReconcilePolicyManage))
	g.Expect(reconcilers.ReconcilePolicyForAnnotation(context.Background(), "nonsense")).
		To(Equal(annotations.ReconcilePolicyManage))
}
