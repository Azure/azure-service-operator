// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// AzureVirtualMachineExtensionReconciler reconciles a AzureVirtualMachineExtension object
type AzureVirtualMachineExtensionReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azurevirtualmachineextensions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azurevirtualmachineextensions/status,verbs=get;update;patch

func (r *AzureVirtualMachineExtensionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureVirtualMachineExtension{})
}

func (r *AzureVirtualMachineExtensionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureVirtualMachineExtension{}).
		Complete(r)
}
