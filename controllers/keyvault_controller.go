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
	"os"
	"strconv"
	"time"

	"github.com/Azure/azure-service-operator/resourcemanager/keyvaults"
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/helpers"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const keyVaultFinalizerName = "keyvault.finalizers.azure.com"

// KeyVaultReconciler reconciles a KeyVault object
type KeyVaultReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=keyvaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=keyvaults/status,verbs=get;update;patch

func (r *KeyVaultReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("keyvault", req.NamespacedName)

	var instance azurev1.KeyVault
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Error(err, "unable to fetch KeyVault")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, keyVaultFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, keyVaultFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, keyVaultFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		if err := r.reconcileExternal(&instance); err != nil {
			return ctrl.Result{}, fmt.Errorf("error reconciling keyvault in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	return ctrl.Result{
		RequeueAfter: time.Second * time.Duration(requeueAfter),
	}, nil
}

func (r *KeyVaultReconciler) addFinalizer(instance *azurev1.KeyVault) error {
	helpers.AddFinalizer(instance, keyVaultFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", keyVaultFinalizerName))
	return nil
}

func (r *KeyVaultReconciler) reconcileExternal(instance *azurev1.KeyVault) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	// write information back to instance
	instance.Status.Provisioning = true
	
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	_, err := keyvaults.CreateVaultAndWait(ctx, groupName, name, location)
	if err != nil { 
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't create resource in azure")
		instance.Status.Provisioning = false
		errUpdate := r.Update(ctx, instance)
		if errUpdate != nil {
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		}
		return err
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	if 	err = r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, "Normal", "Updated", name+" provisioned")
	return nil
}

func (r *KeyVaultReconciler) deleteExternal(instance *azurev1.KeyVault) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	_, err := keyvaults.DeleteVault(ctx, groupName, name)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resouce in azure")
		return err
	}
	return nil
}

func (r *KeyVaultReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.KeyVault{}).
		Complete(r)
}
