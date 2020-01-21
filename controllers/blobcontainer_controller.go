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
	"reflect"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/prometheus/common/log"
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
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if instance.IsBeingDeleted() {
		if helpers.HasFinalizer(&instance, blobContainerFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				azerr := errhelp.NewAzureErrorAzureError(err)

				catch := []string{
					errhelp.AsyncOpIncompleteError,
				}
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
				}
				catch = []string{
					errhelp.ParentNotFoundErrorCode,
					errhelp.ResourceGroupNotFoundErrorCode,
				}
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Error about parent resource not found which we can ignore")
					return ctrl.Result{}, nil
				}

				instance.Status.Message = fmt.Sprintf("Delete blob container failed with %v", err)
				log.Info(instance.Status.Message)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, blobContainerFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, fmt.Errorf("Error removing finalizer: %v", err)
			}
		}
		return ctrl.Result{}, nil
	}

	defer func() {
		if !helpers.IsBeingDeleted(&instance) {
			if err := r.Status().Update(ctx, &instance); err != nil {
				r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
			}
		}
	}()

	if !instance.HasFinalizer(blobContainerFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("Error adding finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "Submitting", "Starting resource reconciliation")
		if err := r.reconcileExternal(&instance); err != nil {
			instance.Status.Message = err.Error()
			// Catch most common errors
			// Blob service error codes:
			// https://docs.microsoft.com/en-us/rest/api/storageservices/blob-service-error-codes
			catchIgnorable := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catchIgnorable, azerr.Type) {
					msg := fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
					r.Log.Info(msg)
					// Requeue if ReconcileExternal errors on one of these codes
					return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
				}
			}
			catchNotIgnorable := []string{
				errhelp.ContainerOperationFailure, // Container name was invalid
				errhelp.ValidationError,           // Some validation (such as min/max length for ContainerName) failed
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catchNotIgnorable, azerr.Type) {
					r.Log.Info(fmt.Sprintf("Got error type: %s", azerr.Type))
					msg := fmt.Sprintf("Got error type: %s", azerr.Type)
					log.Info(msg)
					instance.Status.Message = msg
					// Do not requeue if ReconcileExternal errors on one of these codes
					return ctrl.Result{Requeue: false}, nil
				}
				// DEBUGGING
				r.Log.Info(fmt.Sprintf("catchNotIgnorable did not include error: %s", azerr.Type))
			}
		} else {
			r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "blobcontainer "+instance.ObjectMeta.Name+" provisioned")
		}
	}
	return ctrl.Result{}, nil
}

func (r *BlobContainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.BlobContainer{}).
		Complete(r)
}

func (r *BlobContainerReconciler) reconcileExternal(instance *azurev1alpha1.BlobContainer) error {
	ctx := context.Background()
	groupName := instance.Spec.ResourceGroup
	accountName := instance.Spec.AccountName
	containerName := instance.ObjectMeta.Name
	accessLevel := instance.Spec.AccessLevel

	//get owner instance of Storage
	r.Recorder.Event(instance, corev1.EventTypeNormal, "UpdatingOwner", "Attempting to get owner Storage Account instance")
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

	// Write owner ref information back to instance
	if ownerrefupdateerr := r.Update(ctx, instance); ownerrefupdateerr != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update owner ref information of instance")
	}

	r.Log.Info("Info", "Starting provisioning", fmt.Sprintf("Creating blob container: %s", containerName))
	instance.Status.Provisioning = true

	// Write status information back to instance
	if statusupdateerr := r.Status().Update(ctx, instance); statusupdateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance status")
	}

	// Create the blob container
	_, err := r.StorageManager.CreateBlobContainer(ctx, groupName, accountName, containerName, accessLevel)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", err.Error())

		// WIP: Validation error handling investigation
		// Log TypeOf error to identify type of MinLength/MaxLength validation error
		if strings.Contains(err.Error(), "MinLength") {
			r.Log.Info("ERROR", "VALIDATION ERROR - MinLength rule violated", reflect.TypeOf(err))
			return errhelp.NewAzureError(err)
		}
		if strings.Contains(err.Error(), "MaxLength") {
			r.Log.Info("ERROR", "VALIDATION ERROR - MaxLength rule violated", reflect.TypeOf(err))
			return errhelp.NewAzureError(err)
		}
		// END WIP: Validation error handling investigation

		instance.Status.Message = fmt.Sprintf("Unable to create blob container in Azure %v", err)

		return errhelp.NewAzureError(err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = successMsg

	// Write status information back to instance
	if statusupdateerr := r.Status().Update(ctx, instance); statusupdateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance status")
	}

	return nil
}

func (r *BlobContainerReconciler) deleteExternal(instance *azurev1alpha1.BlobContainer) error {
	ctx := context.Background()
	groupName := instance.Spec.ResourceGroup
	accountName := instance.Spec.AccountName
	containerName := instance.ObjectMeta.Name

	r.Log.Info(fmt.Sprintf("Deleting blob container: " + containerName))
	_, err := r.StorageManager.DeleteBlobContainer(ctx, groupName, accountName, containerName)
	if err != nil {

		azerr := errhelp.NewAzureErrorAzureError(err)
		if azerr.Type == errhelp.ParentNotFoundErrorCode || azerr.Type == errhelp.ResourceGroupNotFoundErrorCode {
			log.Info("Got ignorable error", "type", azerr.Type)
			msg := fmt.Sprintf("Deleted blob container: %s", containerName)
			r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", msg)
			instance.Status.Message = msg
			return nil
		}

		msg := fmt.Sprintf("Unable to delete blob container from Azure: %v", err)
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		instance.Status.Message = msg

		return errhelp.NewAzureError(err)
	}

	msg := fmt.Sprintf("Deleted blob container: %s", containerName)
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

	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("Finalizer added: %s", blobContainerFinalizerName))

	return nil
}
