// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// MySQLVNetRuleReconciler reconciles a MySQLVNetRule object
type MySQLVNetRuleReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlvnetrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlvnetrules/status,verbs=get;update;patch

func (r *MySQLVNetRuleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("mysqlvnetrule", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *MySQLVNetRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.MySQLVNetRule{}).
		Complete(r)
}
