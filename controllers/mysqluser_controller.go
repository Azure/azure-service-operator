// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// MySqlServerPort is the default server port for psql server
const MySqlServerPort = 3306

// MDriverName is driver name for db connection
const MDriverName = "postgres"

// MSecretUsernameKey is the username key in secret
const MSecretUsernameKey = "username"

// MSecretPasswordKey is the password key in secret
const MSecretPasswordKey = "password"

// MySQLUserFinalizerName is the name of the finalizer
const MySQLUserFinalizerName = "mysqluser.finalizers.azure.com"

// MySQLUserReconciler reconciles a MySQLUser object
type MySQLUserReconciler struct {
	Reconciler *AsyncReconciler
}

// Reconcile for mysqluser
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=MySQLUsers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=MySQLUsers/status,verbs=get;update;patch
func (r *MySQLUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.MySQLUser{})
}

// SetupWithManager runs reconcile loop with manager
func (r *MySQLUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.MySQLUser{}).
		Complete(r)
}
