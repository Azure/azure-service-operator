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

// CosmosDBReconciler reconciles a CosmosDB object
type CosmosDBReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=service.azure,resources=cosmosdbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=service.azure,resources=cosmosdbs/status,verbs=get;update;patch

func (r *CosmosDBReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("cosmosdb", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *CosmosDBReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&servicev1alpha1.CosmosDB{}).
		Complete(r)
}
