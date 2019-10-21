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

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"

	storages "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
)

const blobContainerFinalizerName = "blobcontainer.finalizers.com"

// BlobContainerReconciler reconciles a BlobContainer object
type BlobContainerReconciler struct {
	client.Client
	Log                  logr.Logger
	Recorder             record.EventRecorder
	BlobContainerManager storages.BlobContainerManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=blobcontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=blobcontainers/status,verbs=get;update;patch

func (r *BlobContainerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("blobcontainer", req.NamespacedName)

	var instance azurev1alpha1.BlobContainer
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve blobcontainer resource", "err", err.Error())
		// TODO: What is the requeue logic here?  What exactly is client.IgnoreNotFound() doing
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if instance.IsBeingDeleted() {
		if helpers.HasFinalizer(&instance, blobContainerFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Error", "Delete blob container failed with ", err)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, blobContainerFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !instance.HasFinalizer(blobContainerFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when adding finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		err := r.reconcileExternal(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *BlobContainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.BlobContainer{}).
		Complete(r)
}

func (r *BlobContainerReconciler) reconcileExternal(instance *azurev1alpha1.BlobContainer) error {
	r.Log.Info("Info", "Info", "Entered reconcileExternal for BlobContainer")
	return nil
}

func (r *BlobContainerReconciler) deleteExternal(instance *azurev1alpha1.BlobContainer) error {
	ctx := context.Background()
	groupName := instance.Spec.ResourceGroup
	accountName := instance.Spec.AccountName
	containerName := instance.Spec.ContainerName

	r.Log.Info(fmt.Sprintf("deleting blob container: " + containerName))
	_, err := storages.BlobContainerManager.DeleteBlobContainer(ctx, groupName, accountName, containerName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, v1.EventTypeWarning, "DoesNotExist", "Resource to delete does not exist")
			return nil
		}
		msg := "Couldn't delete resouce in azure"
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		instance.Status.Message = msg

		return err
	}

	msg := fmt.Sprintf("Deleted %s", containerName)
	r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", msg)
	instance.Status.Message = msg

	return nil
}

func (r *BlobContainerReconciler) addFinalizer(instance *azurev1alpha1.BlobContainer) error {
	helpers.AddFinalizer(instance, blobContainerFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		msg := fmt.Sprintf("Failed to update finalizer: %v", err)
		instance.Status.Message = msg

		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", blobContainerFinalizerName))
	return nil
}
