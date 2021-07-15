// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"context"

	ctrl "sigs.k8s.io/controller-runtime"
)

// AzureSqlActionReconciler reconciles a AzureSqlAction object
type AzureSqlActionReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlactions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={azuresqlactions/status,azuresqlactions/finalizers},verbs=get;update;patch

// Reconcile function runs the actual reconcilation loop of the controller
func (r *AzureSqlActionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha1.AzureSqlAction{})
}

func (r *AzureSqlActionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlAction{}).
		Complete(r)
}
