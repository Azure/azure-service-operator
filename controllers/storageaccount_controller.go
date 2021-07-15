// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// StorageAccountReconciler reconciles a Storage Account object
type StorageAccountReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=storageaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={storageaccounts/status,storageaccounts/finalizers},verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *StorageAccountReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha1.StorageAccount{})
}

// SetupWithManager sets up the controller functions
func (r *StorageAccountReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.StorageAccount{}).
		Complete(r)
}
