// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// APIMAPIReconciler reconciles a APIM object
type APIMAPIReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=apimgmtapis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=apimgmtapis/status,verbs=get;update;patch

// Reconcile attempts to set the desired state snapshot representation of the service in k8s
func (r *APIMAPIReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.APIMgmtAPI{})
}

// SetupWithManager initializes the control loop for this operator
func (r *APIMAPIReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.APIMgmtAPI{}).
		Complete(r)
}
