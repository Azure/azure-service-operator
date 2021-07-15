// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// KeyVaultReconciler reconciles a KeyVault object
type KeyVaultReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=keyvaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={keyvaults/status,keyvaults/finalizers},verbs=get;update;patch

// Reconcile function runs the actual reconcilation loop of the controller
func (r *KeyVaultReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha1.KeyVault{})
}

// SetupWithManager sets up the controller functions
func (r *KeyVaultReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.KeyVault{}).
		Complete(r)
}
