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

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// "github.com/azure-service-operator/pkg/resourcemanager/adlsgen2s"
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"k8s.io/client-go/tools/record"
)

// AdlsGen2Reconciler reconciles a AdlsGen2 object
type AdlsGen2Reconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	// AdlsGen2Manager adlsgen2s.AdlsGen2Manager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=adlsgen2s,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=adlsgen2s/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AdlsGen2Reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("adlsgen2", req.NamespacedName)

	var instance azurev1.AdlsGen2

	// TODO: requeue after a certain amount of time. example in storage_controller

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("unable to retrieve ADLS Gen2 resource", "err", err.Error())

		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		log.Info("Success", "deletion requested successfully - stubbed", nil)
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		_ = r.reconcileExternal(&instance)
		log.Info("Success", "Here we would submit for creation - stubbed", nil)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller functions
func (r *AdlsGen2Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.AdlsGen2{}).
		Complete(r)
}

func (r *AdlsGen2Reconciler) reconcileExternal(instance *azurev1.AdlsGen2) error {
	ctx := context.Background()
	// storageAccountName := instance.Spec.StorageAccountName
	name := instance.ObjectMeta.Name

	var err error

	// write info back to instance
	instance.Status.Provisioning = true

	err = r.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "unable to update instance")
	}

	// TODO: add logic to actually create a data lake

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, "Normal", "Updated", name+" provisioned - stubbed")
	return nil
}
