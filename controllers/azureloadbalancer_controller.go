// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// AzureLoadBalancerReconciler reconciles a AzureLoadBalancer object
type AzureLoadBalancerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azureloadbalancers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azureloadbalancers/status,verbs=get;update;patch

func (r *AzureLoadBalancerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureLoadBalancer{})
}

func (r *AzureLoadBalancerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureLoadBalancer{}).
		Complete(r)
}
