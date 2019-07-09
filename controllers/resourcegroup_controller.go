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
	"fmt"

	azurev1 "Telstra.Dx.AzureOperator/api/v1"
	resoucegroupsresourcemanager "Telstra.Dx.AzureOperator/resourcemanager/resourcegroups"

	"context"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ResourceGroupReconciler reconciles a ResourceGroup object
type ResourceGroupReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=azure.telstra.k8.io,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.telstra.k8.io,resources=resourcegroups/status,verbs=get;update;patch

func (r *ResourceGroupReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("resourcegroup", req.NamespacedName)

	var instance azurev1.ResourceGroup
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Error(err, "unable to fetch resourcegroup")
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

	if !instance.HasFinalizer(resouceGroupFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		r.createResourceGroup(&instance)
	}

	return ctrl.Result{}, nil

}

func (r *ResourceGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.ResourceGroup{}).
		Complete(r)
}

func (r *ResourceGroupReconciler) createResourceGroup(instance *azurev1.ResourceGroup) {

	log := r.Log.WithValues("resourcegroup", instance)
	ctx := context.Background()
	var err error
	resourcegroupLocation := instance.Spec.Location
	resourcegroupName := instance.ObjectMeta.Name

	// write information back to instance
	instance.Status.Provisioning = true
	err = r.Update(ctx, instance)
	if err != nil {
		log.Error(err, "unable to update resourcegroup before submitting to resource manager")
	}

	_, err = resoucegroupsresourcemanager.CreateGroup(ctx, resourcegroupName, resourcegroupLocation)
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

	r.Recorder.Event(instance, "Normal", "Updated", resourcegroupName+" provisioned")

}
func (r *ResourceGroupReconciler) deleteResourceGroup(instance *azurev1.ResourceGroup) error {
	log := r.Log.WithValues("eventhub", instance)
	ctx := context.Background()

	resourcegroup := instance.ObjectMeta.Name

	var err error
	_, err = resoucegroupsresourcemanager.DeleteGroup(ctx, resourcegroup)
	if err != nil {
		log.Error(err, "ERROR")
	}
	return nil
}
