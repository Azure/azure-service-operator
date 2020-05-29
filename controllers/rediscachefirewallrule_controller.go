// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// RedisCacheFirewallRuleReconciler reconciles a RedisCacheFirewallRule object
type RedisCacheFirewallRuleReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=rediscachefirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=rediscachefirewallrules/status,verbs=get;update;patch

func (r *RedisCacheFirewallRuleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.RedisCacheFirewallRule{})
}

func (r *RedisCacheFirewallRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.RedisCacheFirewallRule{}).
		Complete(r)
}
