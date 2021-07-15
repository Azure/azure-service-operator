// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"github.com/Azure/azure-service-operator/api/v1beta1"

	"context"

	ctrl "sigs.k8s.io/controller-runtime"
)

// AzureSqlDatabaseReconciler reconciles a AzureSqlDatabase object
type AzureSqlDatabaseReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqldatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={azuresqldatabases/status,azuresqldatabases/finalizers},verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureSqlDatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &v1beta1.AzureSqlDatabase{})
}

// SetupWithManager function sets up the functions with the controller
func (r *AzureSqlDatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.AzureSqlDatabase{}).
		Complete(r)
}
