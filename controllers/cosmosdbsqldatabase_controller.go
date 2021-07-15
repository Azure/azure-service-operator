// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
)

// CosmosDBSQLDatabaseReconciler reconciles a CosmosDB SQL Database object
type CosmosDBSQLDatabaseReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=cosmosdbsqldatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={cosmosdbsqldatabases/status,cosmosdbsqldatabases/finalizers},verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *CosmosDBSQLDatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &v1alpha1.CosmosDBSQLDatabase{})
}

// SetupWithManager sets up the controller functions
func (r *CosmosDBSQLDatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.CosmosDBSQLDatabase{}).
		Complete(r)
}
