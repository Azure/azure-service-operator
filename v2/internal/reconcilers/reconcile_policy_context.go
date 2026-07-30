/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

type reconcilePoliciesContextKey struct{}

// ReconcilePolicies are the reconcile policies in effect during a reconcile, none of which a resource
// extension can work out on its own. Read them with ReconcilePolicyFromContext or
// ReconcilePolicyForAnnotation rather than combining them.
type ReconcilePolicies struct {
	// Effective is the policy for the resource being reconciled, resolved from its own annotation, its
	// namespace, and the operator's configuration.
	Effective annotations.ReconcilePolicyValue

	// Inherited is the policy a resource in this namespace carrying no annotation of its own would get.
	Inherited annotations.ReconcilePolicyValue

	// Default is the policy the operator is configured with, which an unusable annotation falls back to.
	Default annotations.ReconcilePolicyValue
}

// WithReconcilePolicies returns a context carrying the reconcile policies in effect.
func WithReconcilePolicies(ctx context.Context, policies ReconcilePolicies) context.Context {
	return context.WithValue(ctx, reconcilePoliciesContextKey{}, policies)
}

// ReconcilePolicyFromContext returns the policy for the resource being reconciled. A post-reconcile check
// runs even when that policy forbids modification, since the skip path still updates status, so an
// extension that acts on Azure has to ask before doing anything at all.
func ReconcilePolicyFromContext(ctx context.Context) annotations.ReconcilePolicyValue {
	return reconcilePoliciesFromContext(ctx).Effective
}

// ReconcilePolicyForAnnotation returns the reconcile policy for a resource other than the one being
// reconciled, given that resource's own reconcile-policy annotation. An extension that modifies another
// resource in Azure must respect the policy it is managed under, and can't resolve it alone: the resource
// carries only its annotation, while the namespace and the operator supply the rest.
func ReconcilePolicyForAnnotation(ctx context.Context, annotation string) annotations.ReconcilePolicyValue {
	policies := reconcilePoliciesFromContext(ctx)
	if annotation == "" {
		return policies.Inherited
	}

	// The error belongs to the reconcile of the resource carrying the annotation; the policy returned
	// with it is the fallback that reconcile will itself apply
	policy, _ := ParseReconcilePolicy(annotation, policies.Default)

	return policy
}

// reconcilePoliciesFromContext returns the policies recorded on the context, or manage when there are
// none, which happens outside of a reconcile.
func reconcilePoliciesFromContext(ctx context.Context) ReconcilePolicies {
	policies, ok := ctx.Value(reconcilePoliciesContextKey{}).(ReconcilePolicies)
	if !ok {
		return ReconcilePolicies{
			Effective: annotations.ReconcilePolicyManage,
			Inherited: annotations.ReconcilePolicyManage,
			Default:   annotations.ReconcilePolicyManage,
		}
	}

	return policies
}
