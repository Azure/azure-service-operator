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

// WithReconcilePolicies returns a context carrying the reconcile policies in effect.
func WithReconcilePolicies(ctx context.Context, policies annotations.ReconcilePolicies) context.Context {
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
	return reconcilePoliciesFromContext(ctx).ForAnnotation(annotation)
}

// reconcilePoliciesFromContext returns the policies recorded on the context, or manage when there are
// none, which happens outside of a reconcile.
func reconcilePoliciesFromContext(ctx context.Context) annotations.ReconcilePolicies {
	policies, ok := ctx.Value(reconcilePoliciesContextKey{}).(annotations.ReconcilePolicies)
	if !ok {
		return annotations.ReconcilePolicies{
			Effective: annotations.ReconcilePolicyManage,
			Inherited: annotations.ReconcilePolicyManage,
			Default:   annotations.ReconcilePolicyManage,
		}
	}

	return policies
}
