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
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

const fileSystemFinalizerName = "filesystem.finalizers.azure.com"

// AzureDataLakeGen2FileSystemReconciler reconciles a AzureDataLakeGen2FileSystem object
type AzureDataLakeGen2FileSystemReconciler struct {
	client.Client
	Log               logr.Logger
	Recorder          record.EventRecorder
	RequeueTime       time.Duration
	FileSystemManager storages.FileSystemManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuredatalakegen2filesystems,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuredatalakegen2filesystems/status,verbs=get;update;patch

func (r *AzureDataLakeGen2FileSystemReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("adlsgen2FileSystem", req.NamespacedName)

	var instance azurev1alpha1.AzureDataLakeGen2FileSystem

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("unable to retrieve ADLS Gen2 resource", "err", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, fileSystemFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Error", "Delete Storage failed with ", err)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, fileSystemFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, fileSystemFinalizerName) {
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
			}
			if helpers.ContainsString(catch, err.(*errhelp.AzureError).Type) || helpers.ContainsString(catch, err.Error()) {
				log.Info("Got ignorable error", "type", err.(*errhelp.AzureError).Type)
				return ctrl.Result{Requeue: true, RequeueAfter: time.Second * time.Duration(requeueAfter)}, nil
			}
			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *AzureDataLakeGen2FileSystemReconciler) addFinalizer(instance *azurev1alpha1.AzureDataLakeGen2FileSystem) error {
	helpers.AddFinalizer(instance, fileSystemFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", fileSystemFinalizerName))
	return nil
}

func (r *AzureDataLakeGen2FileSystemReconciler) reconcileExternal(instance *azurev1alpha1.AzureDataLakeGen2FileSystem) error {
	ctx := context.Background()
	storageAccountName := instance.Spec.StorageAccountName
	groupName := instance.Spec.ResourceGroupName
	fileSystemName := instance.ObjectMeta.Name
	xMsDate := time.Now().String()

	var err error

	// write info back to instance
	instance.Status.Provisioning = true

	err = r.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "unable to update instance")
	}

	_, err = r.FileSystemManager.CreateFileSystem(ctx, groupName, fileSystemName, to.Int32Ptr(20), xMsDate, storageAccountName)
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

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, "Normal", "Updated", fileSystemName+" provisioned")

	return nil
}

func (r *AzureDataLakeGen2FileSystemReconciler) deleteExternal(instance *azurev1alpha1.AzureDataLakeGen2FileSystem) error {
	ctx := context.Background()
	fileSystemName := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	storageAccountName := instance.Spec.StorageAccountName
	xMsDate := time.Now().String()

	resp, err := r.FileSystemManager.DeleteFileSystem(ctx, groupName, fileSystemName, to.Int32Ptr(40), xMsDate, storageAccountName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, "Warning", "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resource in azure")
		return err
	}

	r.Recorder.Event(instance, "Normal", "Deleted", fileSystemName+" deleted")
	r.Recorder.Event(instance, "Normal", "Deleted", resp.Status)

	return nil
}

// SetupWithManager sets up the controller functions
func (r *AzureDataLakeGen2FileSystemReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureDataLakeGen2FileSystem{}).
		Complete(r)
}
