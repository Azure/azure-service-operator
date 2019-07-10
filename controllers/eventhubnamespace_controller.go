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
	"fmt"

	azurev1 "Telstra.Dx.AzureOperator/api/v1"
	eventhubsresourcemanager "Telstra.Dx.AzureOperator/resourcemanager/eventhubs"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// EventhubNamespaceReconciler reconciles a EventhubNamespace object
type EventhubNamespaceReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubnamespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubnamespaces/status,verbs=get;update;patch

func (r *EventhubNamespaceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("eventhubnamespace", req.NamespacedName)

	// your logic here

	var instance azurev1.EventhubNamespace
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Error(err, "unable to fetch Eventhub")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}
	if instance.IsBeingDeleted() {
		err := r.handleFinalizer(&instance)
		if err != nil {
			return reconcile.Result{}, fmt.Errorf("error when handling finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.HasFinalizer(eventhubNamespaceFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		r.createEventHubNamespace(&instance)
	}

	return ctrl.Result{}, nil
}

func (r *EventhubNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.EventhubNamespace{}).
		Complete(r)
}

func (r *EventhubNamespaceReconciler) createEventHubNamespace(instance *azurev1.EventhubNamespace) {
	log := r.Log.WithValues("eventhubnamespace", instance)
	ctx := context.Background()

	var err error

	namespaceLocation := instance.Spec.Location
	namespaceName := instance.ObjectMeta.Name
	resourcegroup := instance.Spec.ResourceGroup

	// write information back to instance
	instance.Status.Provisioning = true
	err = r.Update(ctx, instance)
	if err != nil {
		log.Error(err, "unable to update eventhubnamespace before submitting to resource manager")
	}

	//todo: check if resource group is not provided find first avaliable resource group

	// create Event Hubs namespace
	_, err = eventhubsresourcemanager.CreateNamespaceAndWait(ctx, resourcegroup, namespaceName, namespaceLocation)
	if err != nil {
		log.Error(err, "ERROR")
	}

	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		log.Error(err, "unable to update eventhubnamespace after submitting to resource manager")
	}

	r.Recorder.Event(instance, "Normal", "Updated", namespaceName+" provisioned")

}
func (r *EventhubNamespaceReconciler) deleteEventhubNamespace(instance *azurev1.EventhubNamespace) error {

	log := r.Log.WithValues("eventhub", instance)
	ctx := context.Background()

	namespaceName := instance.ObjectMeta.Name
	resourcegroup := instance.Spec.ResourceGroup

	var err error
	_, err = eventhubsresourcemanager.DeleteNamespace(ctx, resourcegroup, namespaceName)
	if err != nil {
		log.Error(err, "ERROR")
	}
	return nil
}
