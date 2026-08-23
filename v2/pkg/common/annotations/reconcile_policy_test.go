// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package annotations

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func TestResolvedReconcilePoliciesForResource(t *testing.T) {
	t.Parallel()

	policies := ResolvedReconcilePolicies{
		Effective:       ReconcilePolicyDetachOnDelete,
		NamespacePolicy: ReconcilePolicySkip,
		NamespaceName:   "expected-namespace",
		Global:          ReconcilePolicyManage,
	}

	cases := map[string]struct {
		namespace     string
		annotation    string
		expected      ReconcilePolicyValue
		expectedError string
	}{
		"empty uses namespace policy": {
			namespace: "expected-namespace",
			expected:  ReconcilePolicySkip,
		},
		"valid annotation overrides namespace policy": {
			namespace:  "expected-namespace",
			annotation: string(ReconcilePolicyManage),
			expected:   ReconcilePolicyManage,
		},
		"invalid annotation uses global policy": {
			namespace:  "expected-namespace",
			annotation: "unknown",
			expected:   ReconcilePolicyManage,
		},
		"different namespace returns error": {
			namespace:     "other-namespace",
			expectedError: "expected resource in namespace \"expected-namespace\", but it was in \"other-namespace\"",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			resource := &metav1.PartialObjectMetadata{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   c.namespace,
					Annotations: policyAnnotation(c.annotation),
				},
			}

			actual, err := policies.ForResource(resource)
			if c.expectedError != "" {
				g.Expect(err).To(MatchError(c.expectedError))
				return
			}

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func TestResolvedReconcilePoliciesForResource_AllowsMatchingEmptyNamespace(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	policies := ResolvedReconcilePolicies{
		NamespacePolicy: ReconcilePolicySkip,
		Global:          ReconcilePolicyManage,
	}
	resource := &metav1.PartialObjectMetadata{}

	actual, err := policies.ForResource(resource)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(actual).To(Equal(ReconcilePolicySkip))
}

func policyAnnotation(policy string) map[string]string {
	if policy == "" {
		return nil
	}

	return map[string]string{ReconcilePolicy: policy}
}
