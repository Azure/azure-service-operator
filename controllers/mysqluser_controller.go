// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// MySQLUserReconciler reconciles a MySQLUser object
type MySQLUserReconciler struct {
	Reconciler *AsyncReconciler
}

// Reconcile for mysqluser
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlusers/status,verbs=get;update;patch
func (r *MySQLUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.MySQLUser{})
}

// SetupWithManager runs reconcile loop with manager
func (r *MySQLUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.MySQLUser{}).
		Complete(r)
}
