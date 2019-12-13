/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

package controllers

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const storageFinalizerName = "storage.finalizers.azure.com"

// StorageReconciler reconciles a Storage object
type StorageReconciler struct {
	client.Client
	Log            logr.Logger
	Recorder       record.EventRecorder
	Scheme         *runtime.Scheme
	RequeueTime    time.Duration
	StorageManager storages.StorageManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=storages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=storages/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *StorageReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("storage", req.NamespacedName)

	// Fetch the Storage instance
	var instance azurev1alpha1.Storage

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("unable to retrieve storage resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		//return ctrl.Result{}, helpers.IgnoreKubernetesResourceNotFound(err)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	//log.Info("Getting Storage Account", "Storage.Namespace", instance.Namespace, "Storage.Name", instance.Name)
	//log.V(1).Info("Describing Storage Account", "Storage", instance)

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, storageFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				catch := []string{
					errhelp.AsyncOpIncompleteError,
				}
				azerr := errhelp.NewAzureErrorAzureError(err)
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}

				instance.Status.Message = fmt.Sprintf("Delete Storage Account failed with %v", err)
				log.Info(instance.Status.Message)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, storageFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, storageFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when adding finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		err := r.reconcileExternal(&instance)
		if err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.AccountNameInvalid,
				errhelp.AlreadyExists,
				errhelp.AsyncOpIncompleteError,
			}
			azerr := errhelp.NewAzureErrorAzureError(err)
			if helpers.ContainsString(catch, azerr.Type) {
				if azerr.Type == errhelp.AlreadyExists {
					// This error could happen in two cases - when the storage account
					// exists in some other resource group or when this is a repeat
					// call to the reconcile loop for an update of this exact resource. So
					// we call a Get to check if this is the current resource and if
					// yes, we let the call go through instead of ending the reconcile loop
					_, err := r.StorageManager.GetStorage(ctx, instance.Spec.ResourceGroupName, instance.ObjectMeta.Name)
					if err != nil {
						// This means that the Server exists elsewhere and we should
						// terminate the reconcile loop
						instance.Status.Message = "Storage Account Already exists"
						instance.Status.Provisioning = false
						r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", instance.Status.Message)
						return ctrl.Result{Requeue: false}, nil
					}
				}

				if azerr.Type == errhelp.AccountNameInvalid {
					instance.Status.Message = "Invalid Storage Account Name"
					r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", instance.Status.Message)
					return ctrl.Result{Requeue: false}, nil
				}
				log.Info("Got ignorable error", "type", azerr.Type)
				return ctrl.Result{Requeue: true, RequeueAfter: time.Second * time.Duration(requeueAfter)}, nil
			}
			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		log.Info("waiting for provision to take effect")
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	if err := r.verifyExternal(ctx, &instance); err != nil {
		if err.Error() == "NotReady" {
			instance.Status.Message = fmt.Sprintf("Got ignorable error type: %s", err.Error())
			log.Info(instance.Status.Message)
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
			log.Info(instance.Status.Message)
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}

		return ctrl.Result{}, fmt.Errorf("error verifying storage account in azure: %v", err)
	}
	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "storage account "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *StorageReconciler) addFinalizer(instance *azurev1alpha1.Storage) error {
	helpers.AddFinalizer(instance, storageFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", storageFinalizerName))
	return nil
}

func (r *StorageReconciler) reconcileExternal(instance *azurev1alpha1.Storage) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	sku := instance.Spec.Sku
	kind := instance.Spec.Kind
	accessTier := instance.Spec.AccessTier
	enableHTTPSTrafficOnly := instance.Spec.EnableHTTPSTrafficOnly
	dataLakeEnabled := instance.Spec.DataLakeEnabled

	var err error

	//get owner instance of ResourceGroup
	r.Recorder.Event(instance, corev1.EventTypeNormal, "UpdatingOwner", "Updating owner ResourceGroup instance")
	var ownerInstance azurev1alpha1.ResourceGroup

	// Get resource group
	resourceGroupNamespacedName := types.NamespacedName{Name: groupName, Namespace: instance.Namespace}
	err = r.Get(ctx, resourceGroupNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get owner instance of ReourceGroup")
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Got owner instance of Resource Group and assigning controller reference now")
		innerErr := controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
		if innerErr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to set controller reference to ResourceGroup")
		}
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Owner instance assigned successfully")
	}

	// write information back to instance
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	// write information back to instance
	instance.Status.Provisioning = true

	_, err = r.StorageManager.CreateStorage(ctx, groupName, name, location, sku, kind, nil, accessTier, enableHTTPSTrafficOnly, dataLakeEnabled)
	if err != nil {
		if !strings.Contains(err.Error(), "not complete") {
			msg := fmt.Sprintf("CreateStorage not complete: %v", err)
			instance.Status.Message = msg
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to provision or update instance")
			return err
		}
		if strings.Contains(err.Error(), errhelp.AccountNameInvalid) {
			msg := fmt.Sprintf("Invalid Account Name: %v", err)
			instance.Status.Message = msg
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to provision or update instance")
			return errhelp.NewAzureError(err)
		}
	} else {
		msg := "Resource request successfully submitted to Azure"
		instance.Status.Message = msg
	}

	err = r.Status().Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}
	return nil
}

func (r *StorageReconciler) deleteExternal(instance *azurev1alpha1.Storage) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	_, err := r.StorageManager.DeleteStorage(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, v1.EventTypeWarning, "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't delete resource in azure")
		return err
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", name+" deleted")
	return nil
}

// SetupWithManager sets up the controller functions
func (r *StorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.Storage{}).
		Complete(r)
}

func (r *StorageReconciler) verifyExternal(ctx context.Context, instance *azurev1alpha1.Storage) error {
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	stor, err := r.StorageManager.GetStorage(ctx, groupName, name)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		if azerr.Type != errhelp.ResourceNotFound {
			return azerr
		}

		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = string(stor.ProvisioningState)
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Checking", fmt.Sprintf("instance in %s state", instance.Status.State))

	if instance.Status.State == "Succeeded" {
		instance.Status.Message = "StorageAccount successfully provisioned"
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		return nil
	}

	return fmt.Errorf("NotReady")

}
