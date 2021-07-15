// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
)

// CosmosDBReconciler reconciles a CosmosDB object
type CosmosDBReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=cosmosdbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={cosmosdbs/status,cosmosdbs/finalizers},verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *CosmosDBReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &v1alpha1.CosmosDB{})
}

// SetupWithManager sets up the controller functions
func (r *CosmosDBReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.CosmosDB{}).
		Complete(r)
}
