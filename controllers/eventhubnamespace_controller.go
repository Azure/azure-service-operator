/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	creatorv1 "Telstra.Dx.AzureOperator/api/v1"
	eventhubsresourcemanager "Telstra.Dx.AzureOperator/resourcemanager/eventhubs"
)

// EventhubNamespaceReconciler reconciles a EventhubNamespace object
type EventhubNamespaceReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=creator.telstra.k8.io,resources=eventhubnamespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=creator.telstra.k8.io,resources=eventhubnamespaces/status,verbs=get;update;patch

func (r *EventhubNamespaceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("eventhubnamespace", req.NamespacedName)

	// your logic here

	var instance creatorv1.EventhubNamespace
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Error(err, "unable to fetch Eventhub")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}
	if instance.Status.Provisioning || instance.Status.Provisioned {
		return ctrl.Result{}, nil
	}
	r.createeventhubnamespace(&instance)
	return ctrl.Result{}, nil
}

func (r *EventhubNamespaceReconciler) createeventhubnamespace(instance *creatorv1.EventhubNamespace) {
	log := r.Log.WithValues("eventhubnamespace", instance)
	ctx := context.Background()

	var err error

	namespaceLocation := instance.Spec.Location
	namespaceName := instance.ObjectMeta.Name
	resourcegroup := instance.Spec.ResourceGroupName

	//todo: check if resource group is not provided find first avaliable resource group

	// create Event Hubs namespace
	_, err = eventhubsresourcemanager.CreateNamespace(ctx, resourcegroup, namespaceName, namespaceLocation)
	if err != nil {
		log.Error(err, "ERROR")
	}

}
func (r *EventhubNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&creatorv1.EventhubNamespace{}).
		Complete(r)
}
