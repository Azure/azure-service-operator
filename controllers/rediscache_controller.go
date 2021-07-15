// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

const redisCacheFinalizerName = "rediscache.finalizers.azure.com"

// RedisCacheReconciler reconciles a RedisCache object
type RedisCacheReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=rediscaches,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={rediscaches/status,rediscaches/finalizers},verbs=get;update;patch

func (r *RedisCacheReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha1.RedisCache{})
}

// SetupWithManager sets up the controller functions
func (r *RedisCacheReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.RedisCache{}).
		Complete(r)
}
