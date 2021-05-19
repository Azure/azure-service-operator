// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// AzurePublicIPAddressReconciler reconciles a AzurePublicIPAddress object
type AzurePublicIPAddressReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azurepublicipaddresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={azurepublicipaddresses/status,azurepublicipaddresses/finalizers},verbs=get;update;patch

func (r *AzurePublicIPAddressReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzurePublicIPAddress{})
}

func (r *AzurePublicIPAddressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzurePublicIPAddress{}).
		Complete(r)
}
