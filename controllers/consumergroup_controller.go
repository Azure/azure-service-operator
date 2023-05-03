// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// ConsumerGroupReconciler reconciles a ConsumerGroup object
type ConsumerGroupReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=consumergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={consumergroups/status,consumergroups/finalizers},verbs=get;update;patch

// Reconcile reconciler for consumergroup
func (r *ConsumerGroupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha1.ConsumerGroup{})
}

func (r *ConsumerGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ConsumerGroup{}).
		Complete(r)
}
