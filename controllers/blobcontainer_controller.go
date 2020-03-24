// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

const blobContainerFinalizerName = "blobcontainer.finalizers.com"

// BlobContainerReconciler reconciles a BlobContainer object
type BlobContainerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=blobcontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=blobcontainers/status,verbs=get;update;patch

func (r *BlobContainerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.BlobContainer{})
}

func (r *BlobContainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.BlobContainer{}).
		Complete(r)
}
