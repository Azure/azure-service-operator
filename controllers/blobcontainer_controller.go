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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"

	storages "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
)

const blobContainerFinalizerName = "blobcontainer.finalizers.com"

// BlobContainerReconciler reconciles a BlobContainer object
type BlobContainerReconciler struct {
	client.Client
	Log            logr.Logger
	Recorder       record.EventRecorder
	Scheme         *runtime.Scheme
	StorageManager storages.BlobContainerManager
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

	// Debugging
	r.Log.Info("Info", "instance", instance)

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

	// TODO: Add error handling for creating/deleting blob containers
	// https://docs.microsoft.com/en-us/rest/api/storageservices/blob-service-error-codes
	if !instance.IsSubmitted() {
		r.Log.Info("Info", "Debugging", "Entered !instance.IsSubmitted for BlobContainer")
		r.Log.Info("Info", "instance.Status.Provisioning", instance.Status.Provisioning)
		r.Log.Info("Info", "instance.Status.Provisioned", instance.Status.Provisioned)
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
	ctx := context.Background()
	groupName := instance.Spec.ResourceGroup
	accountName := instance.Spec.AccountName
	containerName := instance.ObjectMeta.Name
	accessLevel := instance.Spec.AccessLevel

	//get owner instance of Storage
	r.Recorder.Event(instance, corev1.EventTypeNormal, "UpdatingOwner", "Updating owner Storage Account instance")
	var ownerInstance azurev1alpha1.Storage

	// Get storage account
	storageAccountNamespacedName := types.NamespacedName{Name: accountName, Namespace: instance.Namespace}
	ownerreferr := r.Get(ctx, storageAccountNamespacedName, &ownerInstance)
	if ownerreferr != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get owner instance of Storage Account")
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Got owner instance of Storage Account and assigning controller reference now")
		innerErr := controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
		if innerErr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to set controller reference to Storage Account")
		}
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Owner instance assigned successfully")
	}

	// write information back to instance
	if ownerrefupdateerr := r.Update(ctx, instance); ownerrefupdateerr != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	r.Log.Info(fmt.Sprintf("Creating blob container: %s", containerName))
	instance.Status.Provisioning = true

	// write information back to instance
	if statusupdateerr := r.Status().Update(ctx, instance); statusupdateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	_, err := r.StorageManager.CreateBlobContainer(ctx, groupName, accountName, containerName, accessLevel)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", err.Error())
		msg := "Couldn't create blob container in azure"
		instance.Status.Message = msg

		return err
	}

	msg := fmt.Sprintf("Created blob container: %s", containerName)
	r.Recorder.Event(instance, v1.EventTypeNormal, "Created", msg)
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = msg

	// write information back to instance
	if statusupdateerr := r.Status().Update(ctx, instance); statusupdateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	return nil
}

func (r *BlobContainerReconciler) deleteExternal(instance *azurev1alpha1.BlobContainer) error {
	ctx := context.Background()
	groupName := instance.Spec.ResourceGroup
	accountName := instance.Spec.AccountName
	containerName := instance.ObjectMeta.Name

	r.Log.Info(fmt.Sprintf("deleting blob container: " + containerName))
	_, err := r.StorageManager.DeleteBlobContainer(ctx, groupName, accountName, containerName)
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
	// Debugging
	r.Log.Info("Info", "instance", instance)

	// Removing to see if this gets rid of panic - NOTE: It does  // TODO: Why?
	//r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", blobContainerFinalizerName))

	return nil
}
