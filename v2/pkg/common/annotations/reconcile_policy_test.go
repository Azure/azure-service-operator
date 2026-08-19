// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package annotations

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestParseReconcilePolicy(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		policy        string
		defaultPolicy ReconcilePolicyValue
		expected      ReconcilePolicyValue
		expectError   bool
	}{
		"manage": {
			policy:        string(ReconcilePolicyManage),
			defaultPolicy: ReconcilePolicySkip,
			expected:      ReconcilePolicyManage,
		},
		"skip": {
			policy:        string(ReconcilePolicySkip),
			defaultPolicy: ReconcilePolicyManage,
			expected:      ReconcilePolicySkip,
		},
		"detach-on-delete": {
			policy:        string(ReconcilePolicyDetachOnDelete),
			defaultPolicy: ReconcilePolicySkip,
			expected:      ReconcilePolicyDetachOnDelete,
		},
		"empty uses default": {
			defaultPolicy: ReconcilePolicySkip,
			expected:      ReconcilePolicySkip,
		},
		"unknown returns default and error": {
			policy:        "unknown",
			defaultPolicy: ReconcilePolicySkip,
			expected:      ReconcilePolicySkip,
			expectError:   true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual, err := ParseReconcilePolicy(c.policy, c.defaultPolicy)

			if c.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).NotTo(HaveOccurred())
			}
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func TestResolvedReconcilePoliciesForAnnotation(t *testing.T) {
	t.Parallel()

	policies := ResolvedReconcilePolicies{
		Effective: ReconcilePolicyDetachOnDelete,
		Namespace: ReconcilePolicySkip,
		Global:    ReconcilePolicyManage,
	}

	cases := map[string]struct {
		annotation string
		expected   ReconcilePolicyValue
	}{
		"empty uses namespace policy": {
			expected: ReconcilePolicySkip,
		},
		"valid annotation overrides namespace policy": {
			annotation: string(ReconcilePolicyManage),
			expected:   ReconcilePolicyManage,
		},
		"invalid annotation uses global policy": {
			annotation: "unknown",
			expected:   ReconcilePolicyManage,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(policies.ForAnnotation(c.annotation)).To(Equal(c.expected))
		})
	}
}
