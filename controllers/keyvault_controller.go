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

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	v1 "k8s.io/api/core/v1"
)

const keyVaultFinalizerName = "keyvault.finalizers.azure.com"

// KeyVaultReconciler reconciles a KeyVault object
type KeyVaultReconciler struct {
	client.Client
	Log             logr.Logger
	Recorder        record.EventRecorder
	RequeueTime     time.Duration
	KeyVaultManager keyvaults.KeyVaultManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=keyvaults,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=keyvaults/status,verbs=get;update;patch

// Reconcile function runs the actual reconcilation loop of the controller
func (r *KeyVaultReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("keyvault", req.NamespacedName)

	var instance azurev1.KeyVault

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to fetch KeyVault", "err", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !helpers.IsBeingDeleted(&instance) {
		if !helpers.HasFinalizer(&instance, keyVaultFinalizerName) {
			if err := r.addFinalizer(&instance); err != nil {
				log.Info("Adding keyvault finalizer failed with ", "err", err.Error())
				return ctrl.Result{}, err
			}
		}
	} else {
		if helpers.HasFinalizer(&instance, keyVaultFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete KeyVault failed with ", "err", err.Error())
				return ctrl.Result{}, err
			}
			helpers.RemoveFinalizer(&instance, keyVaultFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if err := r.reconcileExternal(&instance); err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			log.Info("Requeuing as the async operation is not complete")
			return ctrl.Result{
				Requeue:      true,
				RequeueAfter: time.Second * time.Duration(requeueAfter),
			}, nil
		}
		return ctrl.Result{}, fmt.Errorf("error reconciling keyvault in azure: %v", err)
	}

	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "Keyvault "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *KeyVaultReconciler) addFinalizer(instance *azurev1.KeyVault) error {
	helpers.AddFinalizer(instance, keyVaultFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", keyVaultFinalizerName))
	return nil
}

func (r *KeyVaultReconciler) reconcileExternal(instance *azurev1.KeyVault) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	var final error
	if vault, err := r.KeyVaultManager.CreateVault(ctx, groupName, name, location); err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Recorder.Event(instance, v1.EventTypeNormal, "Provisioning", name+" provisioning")
			return err
		}
		instance.Status.ProvisioningState = to.StringPtr("Failed")
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't create resource in azure")

		if err := r.Status().Update(ctx, instance); err != nil {
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		final = errors.Wrap(err, "failed to update status")
	} else {
		instance.Status.ProvisioningState = to.StringPtr("Succeeded")
		instance.Status.ID = vault.ID
		if err := r.Status().Update(ctx, instance); err != nil {
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		final = errors.Wrap(err, "failed to update status")
	}
	return final
}

func (r *KeyVaultReconciler) deleteExternal(instance *azurev1.KeyVault) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	_, err := r.KeyVaultManager.DeleteVault(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, v1.EventTypeWarning, "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't delete resouce in azure")
		return err
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", name+" deleted")
	return nil
}

// SetupWithManager sets up the controller functions
func (r *KeyVaultReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.KeyVault{}).
		Complete(r)
}
