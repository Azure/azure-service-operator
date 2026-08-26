/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

const ourOperator = "azureserviceoperator-system"

// managed is the usual case: nothing anywhere says to leave these resources alone
var managed = annotations.ResolvedReconcilePolicies{
	Effective:       annotations.ReconcilePolicyManage,
	NamespacePolicy: annotations.ReconcilePolicyManage,
	Global:          annotations.ReconcilePolicyManage,
}

func Test_StartAllowed_GivenPolicies_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		policies      annotations.ResolvedReconcilePolicies
		watcherPolicy string
		expected      bool
	}{
		"Managed watcher": {
			policies: managed,
			expected: true,
		},
		"Watcher annotated skip": {
			policies:      managed,
			watcherPolicy: string(annotations.ReconcilePolicySkip),
			expected:      false,
		},
		"Watcher annotated detach-on-delete is still modified": {
			policies:      managed,
			watcherPolicy: string(annotations.ReconcilePolicyDetachOnDelete),
			expected:      true,
		},
		"Skip inherited from the namespace or the operator": {
			policies: annotations.ResolvedReconcilePolicies{
				Effective:       annotations.ReconcilePolicyManage,
				NamespacePolicy: annotations.ReconcilePolicySkip,
				Global:          annotations.ReconcilePolicyManage,
			},
			expected: false,
		},
		"Watcher annotated manage overrides an inherited skip": {
			policies: annotations.ResolvedReconcilePolicies{
				Effective:       annotations.ReconcilePolicyManage,
				NamespacePolicy: annotations.ReconcilePolicySkip,
				Global:          annotations.ReconcilePolicyManage,
			},
			watcherPolicy: string(annotations.ReconcilePolicyManage),
			expected:      true,
		},
		// An unusable annotation falls back to the operator's policy, which is what the watcher's own
		// reconcile does with it
		"Unreadable policy falls back to the operator's manage": {
			policies:      managed,
			watcherPolicy: "nonsense",
			expected:      true,
		},
		// A namespace saying skip does not reach an annotated resource, so it must not decide this either
		"Unreadable policy is not overridden by a skipping namespace": {
			policies: annotations.ResolvedReconcilePolicies{
				Effective:       annotations.ReconcilePolicyManage,
				NamespacePolicy: annotations.ReconcilePolicySkip,
				Global:          annotations.ReconcilePolicyManage,
			},
			watcherPolicy: "nonsense",
			expected:      true,
		},
		"Unreadable policy under an operator that skips": {
			policies: annotations.ResolvedReconcilePolicies{
				Effective:       annotations.ReconcilePolicyManage,
				NamespacePolicy: annotations.ReconcilePolicyManage,
				Global:          annotations.ReconcilePolicySkip,
			},
			watcherPolicy: "nonsense",
			expected:      false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			watcher := &databasewatcher.Watcher{}
			if c.watcherPolicy != "" {
				watcher.SetAnnotations(map[string]string{annotations.ReconcilePolicy: c.watcherPolicy})
			}

			g.Expect(startAllowed(c.policies, watcher)).To(Equal(c.expected))
		})
	}
}

// A watcher always shares its target's namespace, so a mismatch means the policies in hand were resolved
// somewhere else and can't answer for this watcher
func Test_StartAllowed_GivenWatcherInAnotherNamespace_ReportsTheMismatch(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	watcher := &databasewatcher.Watcher{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "watcher",
			Namespace: "elsewhere",
		},
	}

	_, err := startAllowed(managed, watcher)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("elsewhere"))
}

func Test_WatcherConfigured_WhenDatastoreMissing_ReturnsExpectedReason(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	target := ourTarget()
	watcher := &databasewatcher.Watcher{
		ObjectMeta: metav1.ObjectMeta{
			Name: "watcher",
			Annotations: map[string]string{
				reconcilers.OperatorNamespaceAnnotation: ourOperator,
			},
		},
	}

	reason, ok := watcherConfigured(target, watcher)
	g.Expect(ok).To(BeTrue())
	g.Expect(reason).To(ContainSubstring("no datastore"))
}

func Test_ForeignWatcher_WhenTargetAndWatcherAreClaimedByDifferentOperators_ReturnsExpectedReason(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	target := ourTarget()

	watcher := &databasewatcher.Watcher{
		ObjectMeta: metav1.ObjectMeta{
			Name: "watcher",
			Annotations: map[string]string{
				reconcilers.OperatorNamespaceAnnotation: "other-operator",
			},
		},
		Spec: databasewatcher.Watcher_Spec{
			Datastore: &databasewatcher.Datastore{
				KustoOfferingType: to.Ptr("adx"),
			},
		},
	}

	reason, ok := foreignWatcher(target, watcher)
	g.Expect(ok).To(BeTrue())
	g.Expect(reason).To(ContainSubstring("managed by the operator"))
}

// Sharing a credential, whether by naming the same one or by neither naming any, is the ordinary case and
// must not be mistaken for a mismatch
func Test_DifferingCredential_GivenAnnotations_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		targetCredential  *string
		watcherCredential *string
		expected          bool
	}{
		"Neither annotated, so both take whatever the namespace or operator provides": {},
		"Both naming the same secret": {
			targetCredential:  to.Ptr("shared"),
			watcherCredential: to.Ptr("shared"),
		},
		// Carrying the annotation names a secret even when it names an empty one, which the provider then
		// tries and fails to load; that is not what an unannotated resource gets
		"Target's empty annotation against the watcher's none": {
			targetCredential: to.Ptr(""),
			expected:         true,
		},
		"Watcher's empty annotation against the target's none": {
			watcherCredential: to.Ptr(""),
			expected:          true,
		},
		// These two agree only if the namespace has an aso-credential secret, which can't be seen here
		"Watcher names the namespace secret, target names nothing": {
			watcherCredential: to.Ptr(identity.NamespacedSecretName),
			expected:          true,
		},
		"Target names the namespace secret, watcher names nothing": {
			targetCredential: to.Ptr(identity.NamespacedSecretName),
			expected:         true,
		},
		"Only the watcher names a secret": {
			watcherCredential: to.Ptr("watcher-credential"),
			expected:          true,
		},
		"Only the target names a secret": {
			targetCredential: to.Ptr("target-credential"),
			expected:         true,
		},
		"Each naming a different secret": {
			targetCredential:  to.Ptr("target-credential"),
			watcherCredential: to.Ptr("watcher-credential"),
			expected:          true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			target := &databasewatcher.Target{}
			if c.targetCredential != nil {
				target.SetAnnotations(map[string]string{annotations.PerResourceSecret: *c.targetCredential})
			}

			watcher := &databasewatcher.Watcher{}
			if c.watcherCredential != nil {
				watcher.SetAnnotations(map[string]string{annotations.PerResourceSecret: *c.watcherCredential})
			}

			g.Expect(differingCredential(target, watcher)).To(Equal(c.expected))
		})
	}
}

func Test_TargetPostReconcileCheck_GivenWrongType_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := runPostReconcileCheck(&TargetExtension{}, &databasewatcher.Watcher{}, nil)

	g.Expect(err).To(HaveOccurred())
}

// A watcher another operator has claimed is managed under a policy and a credential this one cannot see,
// so it is refused rather than acted on under this operator's configuration
func Test_ForeignWatcher_GivenOperators_ReturnsExpectedReason(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		targetOperator  string
		watcherOperator string
		expectBlocked   bool
	}{
		// Unreachable in practice, since a resource is claimed before any extension runs; accepting it
		// would weaken the invariant for no gain
		"Neither claimed yet": {expectBlocked: true},
		"Both claimed by the same operator": {
			targetOperator:  ourOperator,
			watcherOperator: ourOperator,
		},
		"Claimed by different operators": {
			targetOperator:  ourOperator,
			watcherOperator: "other-operator",
			expectBlocked:   true,
		},
		"Watcher claimed, target not": {
			watcherOperator: "other-operator",
			expectBlocked:   true,
		},
		"Target claimed, watcher not": {
			targetOperator: ourOperator,
			expectBlocked:  true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			target := &databasewatcher.Target{}
			if c.targetOperator != "" {
				target.SetAnnotations(map[string]string{reconcilers.OperatorNamespaceAnnotation: c.targetOperator})
			}

			watcher := &databasewatcher.Watcher{
				ObjectMeta: metav1.ObjectMeta{
					Name: "watcher",
				},
			}
			if c.watcherOperator != "" {
				watcher.SetAnnotations(map[string]string{reconcilers.OperatorNamespaceAnnotation: c.watcherOperator})
			}

			reason, ok := foreignWatcher(target, watcher)

			if !c.expectBlocked {
				g.Expect(ok).To(BeFalse())
				g.Expect(reason).To(BeEmpty())
				return
			}

			g.Expect(ok).To(BeTrue())
			g.Expect(reason).To(ContainSubstring("managed by the operator"))
		})
	}
}

// runPostReconcileCheck invokes checker with policies that permit modification, a next that reports
// success, and no ARM client, so that only the paths that don't call Azure are exercised.
func runPostReconcileCheck(
	checker extensions.PostReconciliationChecker,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
) (extensions.PostReconcileCheckResult, error) {
	next := func(
		_ context.Context,
		_ genruntime.MetaObject,
		_ genruntime.MetaObject,
		_ *resolver.Resolver,
		_ *genericarmclient.GenericClient,
		_ logr.Logger,
		_ annotations.ResolvedReconcilePolicies,
	) (extensions.PostReconcileCheckResult, error) {
		return extensions.PostReconcileCheckResultSuccess(), nil
	}

	return checker.PostReconcileCheck(
		context.Background(), obj, owner, nil, nil, logr.Discard(), managed, next,
	)
}

// ourTarget is a target this operator has claimed, which is what every target is by the time an
// extension sees it
func ourTarget() *databasewatcher.Target {
	return &databasewatcher.Target{
		ObjectMeta: metav1.ObjectMeta{
			Name: "target",
			Annotations: map[string]string{
				reconcilers.OperatorNamespaceAnnotation: ourOperator,
			},
		},
	}
}
