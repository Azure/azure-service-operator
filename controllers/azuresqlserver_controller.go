// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/api/v1beta1"
)

// AzureSqlServerReconciler reconciles an AzureSqlServer object
type AzureSqlServerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlservers/status,verbs=get;update;patch

func (r *AzureSqlServerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &v1beta1.AzureSqlServer{})
}

func (r *AzureSqlServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.AzureSqlServer{}).
		Complete(r)
}
