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

	creatorv1 "Telstra.Dx.AzureOperator/api/v1"
	eventhubsresourcemanager "Telstra.Dx.AzureOperator/resourcemanager/eventhubs"
	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// EventhubReconciler reconciles a Eventhub object
type EventhubReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

func ignoreNotFound(err error) error {

	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}

// +kubebuilder:rbac:groups=creator.telstra.k8.io,resources=eventhubs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=creator.telstra.k8.io,resources=eventhubs/status,verbs=get;update;patch

//Reconcile blah
func (r *EventhubReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("eventhub", req.NamespacedName)

	// your logic here
	var instance creatorv1.Eventhub

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

	if !instance.HasFinalizer(eventhubFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		r.createEventhub(&instance)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager blah
func (r *EventhubReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&creatorv1.Eventhub{}).
		Complete(r)
}

func (r *EventhubReconciler) createEventhub(instance *creatorv1.Eventhub) {
	log := r.Log.WithValues("eventhub", instance)
	ctx := context.Background()

	var err error

	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	// write information back to instance
	instance.Status.Provisioning = true
	err = r.Update(ctx, instance)
	if err != nil {
		log.Error(err, "unable to update resourcegroup before submitting to resource manager")
	}
	_, err = eventhubsresourcemanager.CreateHub(ctx, resourcegroup, eventhubNamespace, eventhubName)
	if err != nil {
		log.Error(err, "ERROR")
	}
	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		log.Error(err, "unable to update resourcegroup after submitting to resource manager")
	}

}

func (r *EventhubReconciler) deleteEventhub(instance *creatorv1.Eventhub) error {

	log := r.Log.WithValues("eventhub", instance)
	ctx := context.Background()

	eventhubName := instance.ObjectMeta.Name
	namespaceName := instance.ObjectMeta.Name
	resourcegroup := instance.Spec.ResourceGroup

	var err error
	_, err = eventhubsresourcemanager.DeleteHub(ctx, resourcegroup, namespaceName, eventhubName)
	if err != nil {
		log.Error(err, "ERROR")
	}
	return nil
}
