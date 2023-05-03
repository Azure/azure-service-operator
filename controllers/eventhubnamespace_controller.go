// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	ctrl "sigs.k8s.io/controller-runtime"
)

// EventhubNamespaceReconciler reconciles a EventhubNamespace object
type EventhubNamespaceReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubnamespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={eventhubnamespaces/status,eventhubnamespaces/finalizers},verbs=get;update;patch

// Reconcile reconciler for eventhubnamespace
func (r *EventhubNamespaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha1.EventhubNamespace{})
}

// SetupWithManager sets up the functions for the controller
func (r *EventhubNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.EventhubNamespace{}).
		Complete(r)
}
