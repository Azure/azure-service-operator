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
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches"
	"k8s.io/client-go/tools/record"
)

const redisCacheFinalizerName = "rediscache.finalizers.azure.com"

// RedisCacheReconciler reconciles a RedisCache object
type RedisCacheReconciler struct {
	client.Client
	Log         logr.Logger
	Recorder    record.EventRecorder
	RequeueTime time.Duration
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=rediscaches,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=rediscaches/status,verbs=get;update;patch

func (r *RedisCacheReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("rediscache", req.NamespacedName)

	// Fetch the Redis Cache instance
	var instance azurev1.RedisCache

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve redis cache resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, redisCacheFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Error", "Delete Redis Cache failed with ", err)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, redisCacheFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, redisCacheFinalizerName) {
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

func (r *RedisCacheReconciler) addFinalizer(instance *azurev1.RedisCache) error {
	helpers.AddFinalizer(instance, redisCacheFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", redisCacheFinalizerName))
	return nil
}

func (r *RedisCacheReconciler) reconcileExternal(instance *azurev1.RedisCache) error {
	ctx := context.Background()

	var err error

	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	sku := instance.Spec.Properties.Sku
	enableNonSSLPort := instance.Spec.Properties.EnableNonSslPort

	// write information back to instance
	instance.Status.Provisioning = true

	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	_, err = rediscaches.CreateRedisCache(ctx, groupName, name, location, sku, enableNonSSLPort, nil)
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

func (r *RedisCacheReconciler) deleteExternal(instance *azurev1.RedisCache) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	_, err := rediscaches.DeleteRedisCache(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, "Warning", "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resource in azure")
		return err
	}
	r.Recorder.Event(instance, "Normal", "Deleted", name+" deleted")
	return nil
}

// SetupWithManager sets up the controller functions
func (r *RedisCacheReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.RedisCache{}).
		Complete(r)
}

/* Below code was from prior to refactor.
Left here for future reference for pulling out values post deployment.

func (r *RedisCacheReconciler) updateStatus(req ctrl.Request, resourceGroupName, deploymentName, provisioningState string, outputs interface{}) (*servicev1alpha1.RedisCache, error) {
	ctx := context.Background()
	log := r.Log.WithValues("Redis Cache", req.NamespacedName)

	resource := &servicev1alpha1.RedisCache{}
	r.Get(ctx, req.NamespacedName, resource)
	log.Info("Getting Redis Cache", "RedisCache.Namespace", resource.Namespace, "RedisCache.Name", resource.Name)

	resourceCopy := resource.DeepCopy()
	resourceCopy.Status.DeploymentName = deploymentName
	resourceCopy.Status.ProvisioningState = provisioningState

	err := r.Status().Update(ctx, resourceCopy)
	if err != nil {
		log.Error(err, "unable to update Redis Cache status")
		return nil, err
	}
	log.V(1).Info("Updated Status", "Redis Cache.Namespace", resourceCopy.Namespace, "RedisCache.Name", resourceCopy.Name, "RedisCache.Status", resourceCopy.Status)

	if helpers.IsDeploymentComplete(provisioningState) {
		if outputs != nil {
			resourceCopy.Output.RedisCacheName = helpers.GetOutput(outputs, "redisCacheName")
			resourceCopy.Output.PrimaryKey = helpers.GetOutput(outputs, "primaryKey")
			resourceCopy.Output.SecondaryKey = helpers.GetOutput(outputs, "secondaryKey")
		}

		err := r.syncAdditionalResourcesAndOutput(req, resourceCopy)
		if err != nil {
			log.Error(err, "error syncing resources")
			return nil, err
		}
		log.V(1).Info("Updated additional resources", "Storage.Namespace", resourceCopy.Namespace, "RedisCache.Name", resourceCopy.Name, "RedisCache.AdditionalResources", resourceCopy.AdditionalResources, "RedisCache.Output", resourceCopy.Output)
	}

	return resourceCopy, nil
}

func (r *RedisCacheReconciler) syncAdditionalResourcesAndOutput(req ctrl.Request, s *servicev1alpha1.RedisCache) (err error) {
	ctx := context.Background()
	log := r.Log.WithValues("redisCache", req.NamespacedName)

	secrets := []string{}
	secretData := map[string]string{
		"redisCacheName": "{{.Obj.Output.RedisCacheName}}",
		"primaryKey":     "{{.Obj.Output.PrimaryKey}}",
		"secondaryKey":   "{{.Obj.Output.SecondaryKey}}",
	}
	secret := helpers.CreateSecret(s, s.Name, s.Namespace, secretData)
	secrets = append(secrets, secret)

	resourceCopy := s.DeepCopy()
	resourceCopy.AdditionalResources.Secrets = secrets

	err = r.Update(ctx, resourceCopy)
	if err != nil {
		log.Error(err, "unable to update Redis Cache status")
		return err
	}

	return nil
}
*/
