// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// PSqlServerPort is the default server port for psql server
const PSqlServerPort = 5432

// PDriverName is driver name for db connection
const PDriverName = "postgres"

// PSecretUsernameKey is the username key in secret
const PSecretUsernameKey = "username"

// PSecretPasswordKey is the password key in secret
const PSecretPasswordKey = "password"

// PSQLUserFinalizerName is the name of the finalizer
const PSQLUserFinalizerName = "postgresqluser.finalizers.azure.com"

// PostgreSQLUserReconciler reconciles a PSQLUser object
type PostgreSQLUserReconciler struct {
	Reconciler *AsyncReconciler
}

//Reconcile for postgresqluser
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=postgresqlusers/status,verbs=get;update;patch
func (r *PostgreSQLUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.PostgreSQLUser{})
}

// SetupWithManager runs reconcile loop with manager
func (r *PostgreSQLUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.PostgreSQLUser{}).
		Complete(r)
}
