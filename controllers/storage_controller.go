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
	"time"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	"k8s.io/client-go/tools/record"
)

const storageFinalizerName = "storage.finalizers.azure.com"

// StorageReconciler reconciles a Storage object
type StorageReconciler struct {
	client.Client
	Log         logr.Logger
	Recorder    record.EventRecorder
	RequeueTime time.Duration
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=storages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=storages/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *StorageReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("storage", req.NamespacedName)

	// Fetch the Storage instance
	var instance azurev1.Storage

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
				log.Info("Error", "Delete Storage failed with ", err)
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
			}
			if helpers.ContainsString(catch, err.(*errhelp.AzureError).Type) {
				log.Info("Got ignorable error", "type", err.(*errhelp.AzureError).Type)
				return ctrl.Result{Requeue: true, RequeueAfter: time.Second * time.Duration(requeueAfter)}, nil
			}
			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *StorageReconciler) addFinalizer(instance *azurev1.Storage) error {
	helpers.AddFinalizer(instance, storageFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", storageFinalizerName))
	return nil
}

func (r *StorageReconciler) reconcileExternal(instance *azurev1.Storage) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	sku := instance.Spec.Sku
	kind := instance.Spec.Kind
	accessTier := instance.Spec.AccessTier
	enableHTTPSTrafficOnly := instance.Spec.EnableHTTPSTrafficOnly

	var err error

	// write information back to instance
	instance.Status.Provisioning = true

	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	_, err = storages.CreateStorage(ctx, groupName, name, location, sku, kind, nil, accessTier, enableHTTPSTrafficOnly)
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

	r.Recorder.Event(instance, "Normal", "Updated", name+" provisioned")

	return nil
}

func (r *StorageReconciler) deleteExternal(instance *azurev1.Storage) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	_, err := storages.DeleteStorage(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, "Warning", "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resouce in azure")
		return err
	}

	r.Recorder.Event(instance, "Normal", "Deleted", name+" deleted")
	return nil
}

// SetupWithManager sets up the controller functions
func (r *StorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.Storage{}).
		Complete(r)
}

/* Below code was from prior to refactor.
   Left here for future reference for pulling out values post deployment.

func (r *StorageReconciler) updateStatus(req ctrl.Request, resourceGroupName, deploymentName, provisioningState string, outputs interface{}) (*servicev1alpha1.Storage, error) {
	ctx := context.Background()
	log := r.Log.WithValues("storage", req.NamespacedName)

	resource := &servicev1alpha1.Storage{}
	r.Get(ctx, req.NamespacedName, resource)
	log.Info("Getting Storage Account", "Storage.Namespace", resource.Namespace, "Storage.Name", resource.Name)

	resourceCopy := resource.DeepCopy()
	resourceCopy.Status.DeploymentName = deploymentName
	resourceCopy.Status.ProvisioningState = provisioningState

	err := r.Status().Update(ctx, resourceCopy)
	if err != nil {
		log.Error(err, "unable to update Storage status")
		return nil, err
	}
	log.V(1).Info("Updated Status", "Storage.Namespace", resourceCopy.Namespace, "Storage.Name", resourceCopy.Name, "Storage.Status", resourceCopy.Status)

	if helpers.IsDeploymentComplete(provisioningState) {
		if outputs != nil {
			resourceCopy.Output.StorageAccountName = helpers.GetOutput(outputs, "storageAccountName")
			resourceCopy.Output.Key1 = helpers.GetOutput(outputs, "key1")
			resourceCopy.Output.Key2 = helpers.GetOutput(outputs, "key2")
			resourceCopy.Output.ConnectionString1 = helpers.GetOutput(outputs, "connectionString1")
			resourceCopy.Output.ConnectionString2 = helpers.GetOutput(outputs, "connectionString2")
		}

		err := r.syncAdditionalResourcesAndOutput(req, resourceCopy)
		if err != nil {
			log.Error(err, "error syncing resources")
			return nil, err
		}
		log.V(1).Info("Updated additional resources", "Storage.Namespace", resourceCopy.Namespace, "Storage.Name", resourceCopy.Name, "Storage.AdditionalResources", resourceCopy.AdditionalResources, "Storage.Output", resourceCopy.Output)
	}

	return resourceCopy, nil
}

func (r *StorageReconciler) syncAdditionalResourcesAndOutput(req ctrl.Request, s *servicev1alpha1.Storage) (err error) {
	ctx := context.Background()
	log := r.Log.WithValues("storage", req.NamespacedName)

	secrets := []string{}
	secretData := map[string]string{
		"storageAccountName": "{{.Obj.Output.StorageAccountName}}",
		"key1":               "{{.Obj.Output.Key1}}",
		"key2":               "{{.Obj.Output.Key2}}",
		"connectionString1":  "{{.Obj.Output.ConnectionString1}}",
		"connectionString2":  "{{.Obj.Output.ConnectionString2}}",
	}
	secret := helpers.CreateSecret(s, s.Name, s.Namespace, secretData)
	secrets = append(secrets, secret)

	resourceCopy := s.DeepCopy()
	resourceCopy.AdditionalResources.Secrets = secrets

	err = r.Update(ctx, resourceCopy)
	if err != nil {
		log.Error(err, "unable to update Storage status")
		return err
	}

	return nil
}
*/
