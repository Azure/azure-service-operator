// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// AzureManagedDiskReconciler reconciles a AzureManagedDisk object
type AzureManagedDiskReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuremanageddisks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuremanageddisks/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureManagedDiskReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureManagedDisk{})
}

// SetupWithManager function sets up the functions with the controller
func (r *AzureManagedDiskReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureManagedDisk{}).
		Complete(r)
}
