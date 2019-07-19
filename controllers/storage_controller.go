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

	"github.com/go-logr/logr"
	uuid "github.com/satori/go.uuid"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	servicev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/client/group"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	storagetemplate "github.com/Azure/azure-service-operator/pkg/storage"
)

// StorageReconciler reconciles a Storage object
type StorageReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=service.azure,resources=storages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=service.azure,resources=storages/status,verbs=get;update;patch

func (r *StorageReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("storage", req.NamespacedName)

	// Fetch the Storage instance
	instance := &servicev1alpha1.Storage{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		log.Error(err, "unable to fetch Storage")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	storageFinalizerName := "storage.finalizers.azure"

	// examine DeletionTimestamp to determine if object is under deletion
	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
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
	}

	var resourceGroupName string
	if instance.Status.ResourceGroupName != "" {
		resourceGroupName = instance.Status.ResourceGroupName
	} else {
		resourceGroupName = uuid.NewV4().String()
		log.Info("Creating a new resource group", "ResourceGroupName", resourceGroupName)
		group.CreateGroup(ctx, resourceGroupName)
		_, err = r.updateStatus(req, resourceGroupName)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	log.Info("Reconciling Storage", "Storage.Namespace", instance.Namespace, "Storage.Name", instance.Name)
	template := storagetemplate.New(instance)
	_, err = template.CreateDeployment(ctx, resourceGroupName)
	if err != nil {
		log.Error(err, "Failed to reconcile Storage")
		return ctrl.Result{}, err
	}

	// Storage created successfully - don't requeue
	return ctrl.Result{}, nil
}

func (r *StorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&servicev1alpha1.Storage{}).
		Complete(r)
}

func (r *StorageReconciler) updateStatus(req ctrl.Request, resourceGroupName string) (*servicev1alpha1.Storage, error) {
	ctx := context.Background()
	log := r.Log.WithValues("storage", req.NamespacedName)

	resource := &servicev1alpha1.Storage{}
	r.Get(ctx, req.NamespacedName, resource)

	resourceCopy := resource.DeepCopy()
	resourceCopy.Status.ResourceGroupName = resourceGroupName
	log.Info("Getting Storage Account", "Storage.Namespace", resourceCopy.Namespace, "Storage.Name", resourceCopy.Name, "Storage.Status.ResourceGroupName", resourceCopy.Status.ResourceGroupName)

	if err := r.Status().Update(ctx, resourceCopy); err != nil {
		log.Error(err, "unable to update Storage status")
		return nil, err
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

	// Request object not found, could have been deleted after reconcile request.
	// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
	// Return and don't requeue
	log.Info("Deleting Storage Account", "Storage.Status.ResourceGroupName", instance.Status.ResourceGroupName)
	_, err := group.DeleteGroup(ctx, instance.Status.ResourceGroupName)
	return err
}
