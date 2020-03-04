// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// SqlServerPort is the default server port for sql server
const SqlServerPort = 1433

// DriverName is driver name for db connection
const DriverName = "sqlserver"

// SecretUsernameKey is the username key in secret
const SecretUsernameKey = "username"

// SecretPasswordKey is the password key in secret
const SecretPasswordKey = "password"

// AzureSQLUserFinalizerName is the name of the finalizer
const AzureSQLUserFinalizerName = "azuresqluser.finalizers.azure.com"

// AzureSQLUserReconciler reconciles a AzureSQLUser object
type AzureSQLUserReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=AzureSQLUsers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=AzureSQLUsers/status,verbs=get;update;patch

func (r *AzureSQLUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureSQLUser{})
}

// SetupWithManager runs reconcile loop with manager
func (r *AzureSQLUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSQLUser{}).
		Complete(r)
}
