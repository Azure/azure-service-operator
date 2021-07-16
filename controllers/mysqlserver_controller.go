// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

// MySQLServerReconciler reconciles a MySQLServer object
type MySQLServerReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={mysqlservers/status,mysqlservers/finalizers},verbs=get;update;patch

func (r *MySQLServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &v1alpha2.MySQLServer{})
}

func (r *MySQLServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// In order to watch a secret associated with this server:
	// * We cannot use "For" - this automatically reconciles the type in question, and since the type
	//   is a secret we don't want that.
	// * We cannot use "Owns" - this reconciles the resources owner.
	// * We must use Watches

	// In order to actually make a change to the MySQLServer password, you must issue an update to the server which
	// includes "administratorLogin": "<user>", "administratorLoginPassword": "<password>",

	// TODO: There are some issues with the below because we need the controller to add an annotation to the
	// TODO: secret but then if that resource is deleted we need to remove it. So we have to check in the finalizer
	// TODO: before we delete and remove it from the adminSecret. BUT if we get updated to clear adminsecret
	// TODO: we don't know that we previously had a secret and so will miss removing it...

	adminSecretKey := ".spec.adminSecret"

	// Add a field indexer
	err := mgr.GetFieldIndexer().IndexField(context.Background(), &v1alpha2.MySQLServer{}, adminSecretKey, func(rawObj client.Object) []string {
		obj := rawObj.(*v1alpha2.MySQLServer)
		return []string{obj.Spec.AdminSecret}
	})
	if err != nil {
		return err
	}

	// Where we construct the ctrl.Manager we may want to limit what secrets we watch with
	// this feature: https://github.com/kubernetes-sigs/controller-runtime/blob/master/designs/use-selectors-at-cache.md,
	// likely scoped by a label selector
	mapFunc := func(o client.Object) []reconcile.Request {
		// Safety check that we're looking at a secret, if not nothing to do
		if _, ok := o.(*corev1.Secret); !ok {
			return nil
		}

		// TODO: This should be fast? We don't have a ctx we can use here... maybe need to file a bug?
		ctx, _ := context.WithTimeout(context.Background(), 1 * time.Minute)
		var matchingResources v1alpha2.MySQLServerList
		err := mgr.GetClient().List(ctx, &matchingResources, client.MatchingFields{adminSecretKey: o.GetName()})
		if err != nil {
			// TODO: log the error at least
			return nil
		}


		var result []reconcile.Request

		for _, matchingResource := range matchingResources.Items {
			result = append(result, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: matchingResource.GetNamespace(),
					Name:      matchingResource.GetName(),
				},
			})
		}

		return result
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.MySQLServer{}).
		Watches(&source.Kind{Type: &corev1.Secret{}}, handler.EnqueueRequestsFromMapFunc(mapFunc)).
		Complete(r)
}
