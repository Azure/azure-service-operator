// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	ctrl "sigs.k8s.io/controller-runtime"
)

// AzureSqlDatabaseReconciler reconciles a AzureSqlDatabase object
type AzureSqlDatabaseReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqldatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqldatabases/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureSqlDatabaseReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureSqlDatabase{})
}

// SetupWithManager function sets up the functions with the controller
func (r *AzureSqlDatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlDatabase{}).
		Complete(r)
}
