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
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	"context"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ResourceGroupReconciler reconciles a ResourceGroup object
type ResourceGroupReconciler struct {
	client.Client
	Log                  logr.Logger
	Recorder             record.EventRecorder
	Reconciler           *AsyncReconciler
	ResourceGroupManager resourcegroupsresourcemanager.ResourceGroupManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *ResourceGroupReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	// ctx := context.Background()
	// log := r.Log.WithValues("resourcegroup", req.NamespacedName)

	return r.Reconciler.Reconcile(req, &azurev1alpha1.ResourceGroup{})

	// var instance azurev1alpha1.ResourceGroup
	// if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
	// 	log.Info("Unable to retrieve resourcegroup resource", "err", err.Error())
	// 	// we'll ignore not-found errors, since they can't be fixed by an immediate
	// 	// requeue (we'll need to wait for a new notification), and we can get them
	// 	// on deleted requests.
	// 	return ctrl.Result{}, client.IgnoreNotFound(err)
	// }

	// if instance.IsBeingDeleted() {
	// 	err := r.handleFinalizer(&instance)
	// 	if err != nil {
	// 		return reconcile.Result{}, fmt.Errorf("error when handling finalizer: %v", err)
	// 	}
	// 	return ctrl.Result{}, nil
	// }

	// if !instance.HasFinalizer(resourceGroupFinalizerName) {
	// 	err := r.addFinalizer(&instance)
	// 	if err != nil {
	// 		return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
	// 	}
	// 	return ctrl.Result{}, nil
	// }

	// if !instance.IsSubmitted() {
	// 	err := r.reconcileExternal(&instance)
	// 	if err != nil {
	// 		return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
	// 	}
	// 	return ctrl.Result{}, nil
	// }

	// return ctrl.Result{}, nil

}

// SetupWithManager function sets up the functions with the controller
func (r *ResourceGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(r)
}

func (r *ResourceGroupReconciler) reconcileExternal(instance *azurev1alpha1.ResourceGroup) error {

	ctx := context.Background()
	var err error
	resourcegroupLocation := instance.Spec.Location
	resourcegroupName := instance.ObjectMeta.Name

	// write information back to instance
	instance.Status.Provisioning = true
	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	_, err = r.ResourceGroupManager.CreateGroup(ctx, resourcegroupName, resourcegroupLocation)
	if err != nil {

		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't create resource in azure")
		instance.Status.Provisioning = false
		errUpdate := r.Update(ctx, instance)
		if errUpdate != nil {
			//log error and kill it
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		return err
	}
	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", resourcegroupName+" provisioned")

	return nil

}

func (r *ResourceGroupReconciler) deleteResourceGroup(instance *azurev1alpha1.ResourceGroup) error {
	ctx := context.Background()

	resourcegroup := instance.ObjectMeta.Name

	var err error
	_, err = r.ResourceGroupManager.DeleteGroup(ctx, resourcegroup)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't delete resource in azure")
		return err
	}
	return nil
}
