/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package generic

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

// Inherited decides whether a resource other than the one being reconciled may be touched, so unlike the
// effective policy it has to fail closed
func Test_MergeReconcilePolicy_givenAnnotations_returnsExpectedPolicies(t *testing.T) {
	t.Parallel()

	const namespace = "test-namespace"

	cases := map[string]struct {
		objectPolicy      string
		namespacePolicy   string
		namespaceMissing  bool
		expectedEffective annotations.ReconcilePolicyValue
		expectedInherited annotations.ReconcilePolicyValue
	}{
		"Nothing annotated": {
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedInherited: annotations.ReconcilePolicyManage,
		},
		"Namespace says skip": {
			namespacePolicy:   string(annotations.ReconcilePolicySkip),
			expectedEffective: annotations.ReconcilePolicySkip,
			expectedInherited: annotations.ReconcilePolicySkip,
		},
		"Object overrides a namespace that says skip": {
			objectPolicy:      string(annotations.ReconcilePolicyManage),
			namespacePolicy:   string(annotations.ReconcilePolicySkip),
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedInherited: annotations.ReconcilePolicySkip,
		},
		"Unreadable namespace, object annotated": {
			objectPolicy:      string(annotations.ReconcilePolicyManage),
			namespaceMissing:  true,
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedInherited: annotations.ReconcilePolicySkip,
		},
		// The namespace we cannot read may be the one that says to leave everything alone
		"Unreadable namespace, nothing annotated": {
			namespaceMissing:  true,
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedInherited: annotations.ReconcilePolicySkip,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			s := runtime.NewScheme()
			g.Expect(corev1.AddToScheme(s)).To(Succeed())

			builder := fake.NewClientBuilder().WithScheme(s)
			if !c.namespaceMissing {
				builder = builder.WithObjects(&corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name:        namespace,
						Annotations: policyAnnotation(c.namespacePolicy),
					},
				})
			}

			reconciler := &GenericReconciler{
				KubeClient: kubeclient.NewClient(builder.Build()),
				Config: config.Values{
					DefaultReconcilePolicy: annotations.ReconcilePolicyManage,
				},
			}

			obj := &resources.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "rg",
					Namespace:   namespace,
					Annotations: policyAnnotation(c.objectPolicy),
				},
			}

			policies := reconciler.mergeReconcilePolicy(context.Background(), logr.Discard(), obj)

			g.Expect(policies.Effective).To(Equal(c.expectedEffective))
			g.Expect(policies.Inherited).To(Equal(c.expectedInherited))
		})
	}
}

func policyAnnotation(policy string) map[string]string {
	if policy == "" {
		return nil
	}

	return map[string]string{annotations.ReconcilePolicy: policy}
}
