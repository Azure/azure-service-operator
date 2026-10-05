/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package generic

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Namespace decides whether a resource other than the one being reconciled may be touched, so unlike the
// effective policy it has to fail closed.
func Test_MergeReconcilePolicy_GivenAnnotations_ReturnsExpectedPolicies(t *testing.T) {
	t.Parallel()

	const namespace = "test-namespace"

	cases := map[string]struct {
		objectPolicy      string
		namespacePolicy   string
		namespaceMissing  bool
		expectedError     string
		expectedEffective annotations.ReconcilePolicyValue
		expectedNamespace annotations.ReconcilePolicyValue
	}{
		"Nothing annotated": {
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedNamespace: annotations.ReconcilePolicyManage,
		},
		"Namespace says skip": {
			namespacePolicy:   string(annotations.ReconcilePolicySkip),
			expectedEffective: annotations.ReconcilePolicySkip,
			expectedNamespace: annotations.ReconcilePolicySkip,
		},
		"Namespace with invalid policy uses global policy": {
			namespacePolicy:   "unknown",
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedNamespace: annotations.ReconcilePolicyManage,
		},
		"Object overrides a namespace that says skip": {
			objectPolicy:      string(annotations.ReconcilePolicyManage),
			namespacePolicy:   string(annotations.ReconcilePolicySkip),
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedNamespace: annotations.ReconcilePolicySkip,
		},
		"Object with invalid policy does not fall back to namespace": {
			objectPolicy:      "unknown",
			namespacePolicy:   string(annotations.ReconcilePolicySkip),
			expectedEffective: annotations.ReconcilePolicyManage,
			expectedNamespace: annotations.ReconcilePolicySkip,
		},
		"Unreadable namespace, returns error": {
			namespaceMissing: true,
			expectedError:    "failed to retrieve namespace object",
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

			policies, err := reconciler.mergeReconcilePolicy(context.Background(), logr.Discard(), obj)
			if c.expectedError != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedError)))
			} else {
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(policies.Effective).To(Equal(c.expectedEffective))
				g.Expect(policies.NamespacePolicy).To(Equal(c.expectedNamespace))
			}
		})
	}
}

type policyRecordingReconciler struct {
	createOrUpdatePolicies *annotations.ResolvedReconcilePolicies
	updateStatusPolicies   *annotations.ResolvedReconcilePolicies
}

func (r *policyRecordingReconciler) CreateOrUpdate(
	_ context.Context,
	_ logr.Logger,
	_ record.EventRecorder,
	_ genruntime.MetaObject,
	policies annotations.ResolvedReconcilePolicies,
) (ctrl.Result, error) {
	r.createOrUpdatePolicies = &policies
	return ctrl.Result{}, nil
}

func (*policyRecordingReconciler) Delete(
	_ context.Context,
	_ logr.Logger,
	_ record.EventRecorder,
	_ genruntime.MetaObject,
) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (*policyRecordingReconciler) Claim(
	_ context.Context,
	_ logr.Logger,
	_ record.EventRecorder,
	_ genruntime.MetaObject,
) error {
	return nil
}

func (r *policyRecordingReconciler) UpdateStatus(
	_ context.Context,
	_ logr.Logger,
	_ record.EventRecorder,
	_ genruntime.MetaObject,
	policies annotations.ResolvedReconcilePolicies,
) error {
	r.updateStatusPolicies = &policies
	return nil
}

func Test_CreateOrUpdate_PassesResolvedPoliciesToSelectedPath(t *testing.T) {
	t.Parallel()

	const namespace = "test-namespace"

	cases := map[string]struct {
		objectPolicy      string
		expectCreate      bool
		expectedEffective annotations.ReconcilePolicyValue
	}{
		"manage calls create or update": {
			objectPolicy:      string(annotations.ReconcilePolicyManage),
			expectCreate:      true,
			expectedEffective: annotations.ReconcilePolicyManage,
		},
		"skip calls status only": {
			objectPolicy:      string(annotations.ReconcilePolicySkip),
			expectedEffective: annotations.ReconcilePolicySkip,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			s := runtime.NewScheme()
			g.Expect(corev1.AddToScheme(s)).To(Succeed())

			namespaceObject := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{Name: namespace},
			}
			spy := &policyRecordingReconciler{}
			reconciler := &GenericReconciler{
				Reconciler: spy,
				KubeClient: kubeclient.NewClient(
					fake.NewClientBuilder().
						WithScheme(s).
						WithObjects(namespaceObject).
						Build(),
				),
				Config: config.Values{
					DefaultReconcilePolicy: annotations.ReconcilePolicyManage,
				},
				PositiveConditions: conditions.NewPositiveConditionBuilder(clock.New()),
			}
			obj := &resources.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "rg",
					Namespace:   namespace,
					Annotations: policyAnnotation(c.objectPolicy),
					Finalizers:  []string{genruntime.ReconcilerFinalizer},
				},
			}

			result, err := reconciler.createOrUpdate(context.Background(), logr.Discard(), obj)

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(result).To(Equal(ctrl.Result{}))

			expected := annotations.ResolvedReconcilePolicies{
				Effective:       c.expectedEffective,
				NamespacePolicy: annotations.ReconcilePolicyManage,
				NamespaceName:   namespace,
				Global:          annotations.ReconcilePolicyManage,
			}
			if c.expectCreate {
				g.Expect(spy.createOrUpdatePolicies).NotTo(BeNil())
				g.Expect(*spy.createOrUpdatePolicies).To(Equal(expected))
				g.Expect(spy.updateStatusPolicies).To(BeNil())
			} else {
				g.Expect(spy.createOrUpdatePolicies).To(BeNil())
				g.Expect(spy.updateStatusPolicies).NotTo(BeNil())
				g.Expect(*spy.updateStatusPolicies).To(Equal(expected))
			}
		})
	}
}

func policyAnnotation(policy string) map[string]string {
	if policy == "" {
		return nil
	}

	return map[string]string{annotations.ReconcilePolicy: policy}
}
