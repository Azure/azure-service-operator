// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// PostgreSQLUserReconciler reconciles a PSQLUser object
type PostgreSQLUserReconciler struct {
	Reconciler *AsyncReconciler
}

//Reconcile for postgresqluser
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={postgresqlusers/status,postgresqlusers/finalizers},verbs=get;update;patch
func (r *PostgreSQLUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.PostgreSQLUser{})
}

// SetupWithManager runs reconcile loop with manager
func (r *PostgreSQLUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.PostgreSQLUser{}).
		Complete(r)
}
