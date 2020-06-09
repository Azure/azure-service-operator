// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	ctrl "sigs.k8s.io/controller-runtime"
)

// MySQLServerReconciler reconciles a MySQLServer object
type MySQLServerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlservers/status,verbs=get;update;patch

func (r *MySQLServerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &v1alpha2.MySQLServer{})
}

func (r *MySQLServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.MySQLServer{}).
		Complete(r)
}
