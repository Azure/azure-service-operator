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
	"time"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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

//Reconcile reconciler for eventhubnamespace
func (r *EventhubNamespaceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("eventhubnamespace", req.NamespacedName)

	var instance azurev1.EventhubNamespace
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable tto retrieve eventhub namespace resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
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
		err := r.reconcileExternal(&instance)
		if err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
			}
			if helpers.ContainsString(catch, err.(*errhelp.AzureError).Type) {
				log.Info("Got ignorable error", "type", err.(*errhelp.AzureError).Type)
				return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
			}

			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *EventhubNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.EventhubNamespace{}).
		Complete(r)
}

func (r *EventhubNamespaceReconciler) reconcileExternal(instance *azurev1.EventhubNamespace) error {
	ctx := context.Background()

	var err error

	namespaceLocation := instance.Spec.Location
	namespaceName := instance.ObjectMeta.Name
	resourcegroup := instance.Spec.ResourceGroup

	// write information back to instance
	instance.Status.Provisioning = true

	//get owner instance
	var ownerInstance azurev1.ResourceGroup
	resourceGroupNamespacedName := types.NamespacedName{Name: resourcegroup, Namespace: instance.Namespace}
	err = r.Get(ctx, resourceGroupNamespacedName, &ownerInstance)

	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to get owner instance of resourcegroup")
	} else {
		//set owner reference for eventhubnamespace if it exists
		references := []metav1.OwnerReference{
			metav1.OwnerReference{
				APIVersion: "v1",
				Kind:       "ResourceGroup",
				Name:       ownerInstance.GetName(),
				UID:        ownerInstance.GetUID(),
			},
		}
		instance.ObjectMeta.SetOwnerReferences(references)
	}

	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	// create Event Hubs namespace
	_, err = eventhubsresourcemanager.CreateNamespaceAndWait(ctx, resourcegroup, namespaceName, namespaceLocation)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't create resource in azure")
		instance.Status.Provisioning = false
		errUpdate := r.Update(ctx, instance)
		if errUpdate != nil {
			//log error and kill it
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		}
		return errhelp.NewAzureError(err)
	}

	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, "Normal", "Updated", namespaceName+" provisioned")

	return nil

}

func (r *EventhubNamespaceReconciler) deleteEventhubNamespace(instance *azurev1.EventhubNamespace) error {

	ctx := context.Background()

	namespaceName := instance.ObjectMeta.Name
	resourcegroup := instance.Spec.ResourceGroup

	var err error
	_, err = eventhubsresourcemanager.DeleteNamespace(ctx, resourcegroup, namespaceName)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resouce in azure")
		return err
	}
	return nil
}
