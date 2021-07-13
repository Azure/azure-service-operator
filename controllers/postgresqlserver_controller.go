// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

// PostgreSQLServerReconciler reconciles a PostgreSQLServer object
type PostgreSQLServerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={postgresqlservers/status,postgresqlservers/finalizers},verbs=get;update;patch

func (r *PostgreSQLServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &v1alpha2.PostgreSQLServer{})
}

func (r *PostgreSQLServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.PostgreSQLServer{}).
		Complete(r)
}
