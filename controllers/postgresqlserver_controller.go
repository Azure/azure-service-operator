// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	ctrl "sigs.k8s.io/controller-runtime"
)

// PostgreSQLServerReconciler reconciles a PostgreSQLServer object
type PostgreSQLServerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlservers/status,verbs=get;update;patch

func (r *PostgreSQLServerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &v1alpha2.PostgreSQLServer{})
}

func (r *PostgreSQLServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.PostgreSQLServer{}).
		Complete(r)
}
