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
	"os"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
)

// ConsumerGroupReconciler reconciles a ConsumerGroup object
type ConsumerGroupReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=consumergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=consumergroups/status,verbs=get;update;patch

//Reconcile reconciler for consumergroup
func (r *ConsumerGroupReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("consumergroup", req.NamespacedName)

	var instance azurev1.ConsumerGroup
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Error(err, "unable to fetch consumergroup")
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

	if !instance.HasFinalizer(consumerGroupFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		err := r.createConsumerGroup(&instance)
		if err != nil {

			return ctrl.Result{}, fmt.Errorf("error when creating consumer group in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	return ctrl.Result{
		RequeueAfter: time.Second * time.Duration(requeueAfter),
	}, nil
}

func (r *ConsumerGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.ConsumerGroup{}).
		Complete(r)
}

func (r *ConsumerGroupReconciler) createConsumerGroup(instance *azurev1.ConsumerGroup) error {

	ctx := context.Background()
	var err error
	consumergroupName := instance.ObjectMeta.Name
	namespaceName := instance.Spec.NamespaceName
	resourcegroup := instance.Spec.ResourceGroupName
	eventhubName := instance.Spec.EventhubName

	// write information back to instance
	instance.Status.Provisioning = true
	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	_, err = eventhubsresourcemanager.CreateConsumerGroup(ctx, resourcegroup, namespaceName, eventhubName, consumergroupName)
	if err != nil {

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't create consumer group in azure")
		instance.Status.Provisioning = false
		errUpdate := r.Update(ctx, instance)
		if errUpdate != nil {
			//log error and kill it
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		}
		return err
	}
	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, "Normal", "Updated", consumergroupName+" provisioned")

	return nil

}

func (r *ConsumerGroupReconciler) deleteConsumerGroup(instance *azurev1.ConsumerGroup) error {
	ctx := context.Background()

	consumergroupName := instance.ObjectMeta.Name
	namespaceName := instance.Spec.NamespaceName
	resourcegroup := instance.Spec.ResourceGroupName
	eventhubName := instance.Spec.EventhubName

	var err error
	_, err = eventhubsresourcemanager.DeleteConsumerGroup(ctx, resourcegroup, namespaceName, eventhubName, consumergroupName)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete consumer group in azure")
		return err
	}
	return nil
}
