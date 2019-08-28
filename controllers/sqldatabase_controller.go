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
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"

	//"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
)

const SQLDatabaseFinalizerName = "sqldatabase.finalizers.azure.com"

// SqlDatabaseReconciler reconciles a SqlDatabase object
type SqlDatabaseReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqldatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqldatabases/status,verbs=get;update;patch

func (r *SqlDatabaseReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqldatabase", req.NamespacedName)

	// your logic here
	var instance azurev1.SqlDatabase

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql-database resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, SQLDatabaseFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete SqlDatabase failed with ", err.Error())
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, SQLDatabaseFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, SQLDatabaseFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding SqlDatabase finalizer failed with ", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, "Normal", "Submitting", "starting resource reconciliation for SqlDatabase")
		if err := r.reconcileExternal(&instance); err != nil {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				r.Recorder.Event(&instance, "Normal", "Provisioning", "async op still running")
				return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling sql database in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", "sqldatabase "+instance.ObjectMeta.Name+" provisioned ")

	return ctrl.Result{}, nil
}

func (r *SqlDatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlDatabase{}).
		Complete(r)
}

func (r *SqlDatabaseReconciler) reconcileExternal(instance *azurev1.SqlDatabase) error {
	ctx := context.Background()
	location := instance.Spec.Location
	server := instance.Spec.Server
	groupName := instance.Spec.ResourceGroup
	dbEdition := instance.Spec.Edition

	dbName := instance.ObjectMeta.Name

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        server,
		Location:          location,
	}

	sqlDatabaseProperties := sql.SQLDatabaseProperties{
		DatabaseName: dbName,
		Edition:      dbEdition,
	}

	r.Log.Info("Calling createorupdate SQL database")
	instance.Status.Provisioning = true

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	_, err := sdkClient.CreateOrUpdateDB(sqlDatabaseProperties)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Log.Info("Async operation not complete or group not found")
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

func (r *SqlDatabaseReconciler) deleteExternal(instance *azurev1.SqlDatabase) error {
	ctx := context.Background()
	dbname := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	location := instance.Spec.Location

	// create the Go SDK client with relevant info
	sdk := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        server,
		Location:          location,
	}

	_, err := sdk.DeleteDB(dbname)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, "Warning", "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resouce in azure")
		return err
	}
	r.Recorder.Event(instance, "Normal", "Deleted", dbname+" deleted")
	return nil
}

func (r *SqlDatabaseReconciler) addFinalizer(instance *azurev1.SqlDatabase) error {
	helpers.AddFinalizer(instance, SQLDatabaseFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", SQLDatabaseFinalizerName))
	return nil
}
