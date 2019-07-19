/*
.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	servicev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// StorageReconciler reconciles a Storage object
type StorageReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=service.azure,resources=storages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=service.azure,resources=storages/status,verbs=get;update;patch

func (r *StorageReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("storage", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *StorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&servicev1alpha1.Storage{}).
		Complete(r)
}
