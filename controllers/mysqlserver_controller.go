// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

// MySQLServerReconciler reconciles a MySQLServer object
type MySQLServerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={mysqlservers/status,mysqlservers/finalizers},verbs=get;update;patch

func (r *MySQLServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &v1alpha2.MySQLServer{})
}

func (r *MySQLServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.MySQLServer{}).
		Complete(r)
}
