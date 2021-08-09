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

const AdminSecretKey = ".spec.adminSecret"

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

	// Add a field indexer which we will use to find servers by AdminSecret name
	err := mgr.GetFieldIndexer().IndexField(context.Background(), &v1alpha2.MySQLServer{}, AdminSecretKey, func(rawObj client.Object) []string {
		obj := rawObj.(*v1alpha2.MySQLServer)
		return []string{obj.Spec.AdminSecret}
	})
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.MySQLServer{}).
		Watches(&source.Kind{Type: &corev1.Secret{}}, handler.EnqueueRequestsFromMapFunc(r.makeMapFunc(mgr))).
		Complete(r)
}

// TODO: It may be possible where we construct the ctrl.Manager to limit what secrets we watch with
// TODO: this feature: https://github.com/kubernetes-sigs/controller-runtime/blob/master/designs/use-selectors-at-cache.md,
// TODO: likely scoped by a label selector? That requires additional work as we would need to manage that label
// TODO: in the MySQLServer reconcile loop probably.
func (r *MySQLServerReconciler) makeMapFunc(mgr ctrl.Manager) func(o client.Object) []reconcile.Request {
	return func(o client.Object) []reconcile.Request {
		// Safety check that we're looking at a secret, if not nothing to do
		if _, ok := o.(*corev1.Secret); !ok {
			return nil
		}

		// This should be fast since the list of items it's going through should always be cached
		// locally with the shared informer.
		// Unfortunately we don't have a ctx we can use here, see https://github.com/kubernetes-sigs/controller-runtime/issues/1628
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var matchingResources v1alpha2.MySQLServerList
		err := mgr.GetClient().List(ctx, &matchingResources, client.MatchingFields{AdminSecretKey: o.GetName()})
		if err != nil {
			r.Reconciler.Telemetry.LogError("couldn't list MySQLServers in MapFunc", err)
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
}
