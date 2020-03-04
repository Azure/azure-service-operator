// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// PostgreSQLFirewallRuleReconciler reconciles a PostgreSQLFirewallRule object
type PostgreSQLFirewallRuleReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlfirewallrules/status,verbs=get;update;patch

func (r *PostgreSQLFirewallRuleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.PostgreSQLFirewallRule{})
}

func (r *PostgreSQLFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.PostgreSQLFirewallRule{}).
		Complete(r)
}
