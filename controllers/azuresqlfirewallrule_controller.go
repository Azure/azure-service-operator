// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/api/v1beta1"
)

// AzureSqlFirewallRuleReconciler reconciles a AzureSqlFirewallRule object
type AzureSqlFirewallRuleReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={azuresqlfirewallrules/status,azuresqlfirewallrules/finalizers},verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureSqlFirewallRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, err error) {
	return r.Reconciler.Reconcile(ctx, req, &v1beta1.AzureSqlFirewallRule{})
}

// SetupWithManager function sets up the functions with the controller
func (r *AzureSqlFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.AzureSqlFirewallRule{}).
		Complete(r)
}
