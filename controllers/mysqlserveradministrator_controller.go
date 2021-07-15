// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// MySQLServerAdministratorReconciler reconciles a MySQLServerAdministrator object
type MySQLServerAdministratorReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlserveradministrators,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={mysqlserveradministrators/status,mysqlserveradministrators/finalizers},verbs=get;update;patch

func (r *MySQLServerAdministratorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha1.MySQLServerAdministrator{})
}

func (r *MySQLServerAdministratorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.MySQLServerAdministrator{}).
		Complete(r)
}
