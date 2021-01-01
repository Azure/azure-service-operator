// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// AzureContainerRegistryReconciler reconciles a AzureContainerRegistry object
type AzureContainerRegistryReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azurecontainerregistries,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azurecontainerregistries/status,verbs=get;update;patch

func (r *AzureContainerRegistryReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureContainerRegistry{})
}

func (r *AzureContainerRegistryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureContainerRegistry{}).
		Complete(r)
}
