// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"github.com/Azure/azure-service-operator/api/v1beta1"

	"context"

	ctrl "sigs.k8s.io/controller-runtime"
)

// AzureSqlFailoverGroupReconciler reconciles a AzureSqlFailoverGroup object
type AzureSqlFailoverGroupReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfailovergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={azuresqlfailovergroups/status,azuresqlfailovergroups/finalizers},verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureSqlFailoverGroupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &v1beta1.AzureSqlFailoverGroup{})
}

func (r *AzureSqlFailoverGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.AzureSqlFailoverGroup{}).
		Complete(r)
}
