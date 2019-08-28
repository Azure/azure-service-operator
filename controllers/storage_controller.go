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

	//resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	//servicev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	//"github.com/Azure/azure-service-operator/pkg/client/deployment"
	//"github.com/Azure/azure-service-operator/pkg/client/group"
	//"github.com/Azure/azure-service-operator/pkg/config"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	"k8s.io/client-go/tools/record"
	// "github.com/Azure/azure-service-operator/pkg/helpers"
	//storagetemplate "github.com/Azure/azure-service-operator/pkg/storage"
	// "github.com/Azure/go-autorest/autorest/to"
)

const storageFinalizerName = "storage.finalizers.azure.com"

// StorageReconciler reconciles a Storage object
type StorageReconciler struct {
	client.Client
	Log         logr.Logger
	Recorder    record.EventRecorder
	RequeueTime time.Duration
}

// +kubebuilder:rbac:groups=service.azure,resources=storages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=service.azure,resources=storages/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *StorageReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("storage", req.NamespacedName)

	// Fetch the Storage instance
	//instance := &servicev1alpha1.Storage{}
	var instance azurev1.Storage

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		// if err != nil {
		log.Error(err, "unable to retrieve storage resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		//return ctrl.Result{}, helpers.IgnoreKubernetesResourceNotFound(err)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Getting Storage Account", "Storage.Namespace", instance.Namespace, "Storage.Name", instance.Name)
	log.V(1).Info("Describing Storage Account", "Storage", instance)

	//storageFinalizerName := "storage.finalizers.azure"
	// examine DeletionTimestamp to determine if object is under deletion
	/*if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !helpers.ContainsString(instance.ObjectMeta.Finalizers, storageFinalizerName) {
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, storageFinalizerName)
			if err := r.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if helpers.ContainsString(instance.ObjectMeta.Finalizers, storageFinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			if err := r.deleteExternalResources(instance); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return ctrl.Result{}, err
			}

			// remove our finalizer from the list and update it.
			instance.ObjectMeta.Finalizers = helpers.RemoveString(instance.ObjectMeta.Finalizers, storageFinalizerName)
			if err := r.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, err
	}*/
	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, storageFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete Storage failed with ", err.Error())
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
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding storage finalizer failed with ", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		if err := r.reconcileExternal(&instance); err != nil {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				log.Info("Requeuing as the async operation is not complete")
				return ctrl.Result{
					Requeue:      true,
					RequeueAfter: time.Second * time.Duration(requeueAfter),
				}, nil
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling storage in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", "Storage "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil

	//resourcegroupName := instance.ObjectMeta.Name
	// log.Info("SubscriptionId: ", config.Instance.SubscriptionID)
	// log.Info("ClusterName:", config.Instance.ClusterName)
	// log.Info("instance.Name: ", instance.Name)
	// log.Info("instance namespace: ", instance.Namespace)
	//resourceGroupName := helpers.AzrueResourceGroupName(config.Instance.SubscriptionID, config.Instance.ClusterName, "storage", instance.Name, instance.Namespace)
	/*esourceGroupName := helpers.AzrueResourceGroupName("b840fb3a-097c-462e-b658-5c6364683ae2", "myAKSCluster", "storage", "myinstancename", "mynamespace")
	resourceGroupLocation := instance.Spec.Location
	log.Info("storage controller", "rgn: ", resourceGroupName)
	deploymentName := instance.Status.DeploymentName
	if deploymentName != "" {
		log.Info("Checking deployment", "ResourceGroupName", resourceGroupName, "DeploymentName", deploymentName)
		de, _ := deployment.GetDeployment(ctx, resourceGroupName, deploymentName)
		if de.Properties == nil || de.Properties.ProvisioningState == nil {
			return ctrl.Result{}, nil
		}
		provisioningState := *de.Properties.ProvisioningState
		if helpers.IsDeploymentComplete(provisioningState) {
			log.Info("Deployment is complete", "ProvisioningState", provisioningState)
			_, err = r.updateStatus(req, resourceGroupName, deploymentName, provisioningState, de.Properties.Outputs)
			if err != nil {
				return ctrl.Result{}, err
			}
			if instance.Status.Generation == instance.ObjectMeta.Generation {
				return ctrl.Result{}, nil
			}
		} else {
			log.Info("Requeue the request", "ProvisioningState", provisioningState)
			return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
		}
	}
	instance.Status.Generation = instance.ObjectMeta.Generation
	if err := r.Status().Update(ctx, instance); err != nil {
		return ctrl.Result{}, err
	}

	log.Info("Creating a new resource group", "ResourceGroupName", resourceGroupName)
	// tags := map[string]*string{
	// 	"name":      to.StringPtr(instance.Name),
	// 	"namespace": to.StringPtr(instance.Namespace),
	// 	"kind":      to.StringPtr("storage"),
	// }
	//group.CreateGroup(ctx, resourceGroupName, resourcegroupLocation, tags)
	//from RGController	_, err = resoucegroupsresourcemanager.CreateGroup(ctx, resourcegroupName, resourcegroupLocation)
	_, err = resoucegroupsresourcemanager.CreateGroup(ctx, resourceGroupName, resourceGroupLocation)

	log.Info("Reconciling Storage", "Storage.Namespace", instance.Namespace, "Storage.Name", instance.Name)
	template := storagetemplate.New(instance)
	deploymentName, err = template.CreateDeployment(ctx, resourceGroupName)
	if err != nil {
		log.Error(err, "Failed to reconcile Storage")
		return ctrl.Result{}, err
	}

	de, _ := deployment.GetDeployment(ctx, resourceGroupName, deploymentName)
	_, err = r.updateStatus(req, resourceGroupName, deploymentName, *de.Properties.ProvisioningState, nil)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Storage created successfully - don't requeue
	return ctrl.Result{}, nil
	*/
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
	//sku, kind, tags, accesstier enabblehttpstraffice

	// write information back to instance
	instance.Status.Provisioning = true

	if err := r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	_, err := storages.CreateStorage(ctx, groupName, name, location, sku, kind, nil, accessTier, enableHTTPSTrafficOnly)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Recorder.Event(instance, "Normal", "Provisioning", name+" provisioning")
			return err
		}
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't create resource in azure")
		instance.Status.Provisioning = false
		errUpdate := r.Status().Update(ctx, instance)
		if errUpdate != nil {
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		}
		return err
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	if err = r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

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

/*
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

func (r *StorageReconciler) deleteExternalResources(instance *servicev1alpha1.Storage) error {
	//
	// delete any external resources associated with the storage
	//
	// Ensure that delete implementation is idempotent and safe to invoke
	// multiple types for same object.
	ctx := context.Background()
	log := r.Log.WithValues("Storage.Namespace", instance.Namespace, "Storage.Name", instance.Name)

	resourceGroupName := helpers.AzrueResourceGroupName(config.Instance.SubscriptionID, config.Instance.ClusterName, "storage", instance.Name, instance.Namespace)
	log.Info("Deleting Storage Account", "ResourceGroupName", resourceGroupName)
	_, err := group.DeleteGroup(ctx, resourceGroupName)
	if err != nil && helpers.IgnoreAzureResourceNotFound(err) != nil {
		return err
	}

	err = helpers.DeleteSecret(instance.Name, instance.Namespace)
	if err != nil && helpers.IgnoreKubernetesResourceNotFound(err) != nil {
		return err
	}

	return nil
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
