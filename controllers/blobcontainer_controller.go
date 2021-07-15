// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

const blobContainerFinalizerName = "blobcontainer.finalizers.com"

// BlobContainerReconciler reconciles a BlobContainer object
type BlobContainerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=blobcontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={blobcontainers/status,blobcontainers/finalizers},verbs=get;update;patch

func (r *BlobContainerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &v1alpha2.BlobContainer{})
}

func (r *BlobContainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.BlobContainer{}).
		Complete(r)
}
