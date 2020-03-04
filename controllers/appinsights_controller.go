// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// AppInsightsReconciler reconciles a AppInsights object
type AppInsightsReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=appinsights,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=appinsights/status,verbs=get;update;patch

// Reconcile attempts to set the desired state snapshot representation of the service in k8s
func (r *AppInsightsReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AppInsights{})
}

// SetupWithManager initializes the control loop for this operator
func (r *AppInsightsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AppInsights{}).
		Complete(r)
}
