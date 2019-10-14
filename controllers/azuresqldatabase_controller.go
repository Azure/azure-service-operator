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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

const azureSQLDatabaseFinalizerName = "azuresqldatabase.finalizers.azure.com"

// AzureSqlDatabaseReconciler reconciles a AzureSqlDatabase object
type AzureSqlDatabaseReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqldatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqldatabases/status,verbs=get;update;patch

func (r *AzureSqlDatabaseReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("azuresqldatabase", req.NamespacedName)

	var instance azurev1alpha1.AzureSqlDatabase

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve azure-sql-database resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, azureSQLDatabaseFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete AzureSqlDatabase failed with ", "err", err.Error())
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, azureSQLDatabaseFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, azureSQLDatabaseFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding AzureSqlDatabase finalizer failed with ", "error", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, corev1.EventTypeNormal, "Submitting", "starting resource reconciliation for AzureSqlDatabase")
		if err := r.reconcileExternal(&instance); err != nil {

			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling azure sql database in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, corev1.EventTypeNormal, "Provisioned", "azuresqldatabase "+instance.ObjectMeta.Name+" provisioned ")

	return ctrl.Result{}, nil
}

func (r *AzureSqlDatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlDatabase{}).
		Complete(r)
}

func (r *AzureSqlDatabaseReconciler) reconcileExternal(instance *azurev1alpha1.AzureSqlDatabase) error {
	ctx := context.Background()
	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.ObjectMeta.Name
	dbEdition := instance.Spec.Edition

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        server,
		Location:          location,
	}

	azureSqlDatabaseProperties := sql.SQLDatabaseProperties{
		DatabaseName: dbName,
		Edition:      dbEdition,
	}

	r.Log.Info("Calling createorupdate Azure SQL database")
	// instance.Status.Provisioning = true

	//get owner instance of AzureSqlServer
	r.Recorder.Event(instance, corev1.EventTypeNormal, "UpdatingOwner", "Updating owner AzureSqlServer instance")
	var ownerInstance azurev1alpha1.AzureSqlServer
	azureSqlServerNamespacedName := types.NamespacedName{Name: server, Namespace: instance.Namespace}
	err := r.Get(ctx, azureSqlServerNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get owner instance of AzureSqlServer")
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Got owner instance of Sql Server and assigning controller reference now")
		innerErr := controllerutil.SetControllerReference(&ownerInstance, instance, r.Scheme)
		if innerErr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to set controller reference to AzureSqlServer")
		}
		r.Recorder.Event(instance, corev1.EventTypeNormal, "OwnerAssign", "Owner instance assigned successfully")
	}

	// write information back to instance
	if updateerr := r.Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	_, err = sdkClient.CreateOrUpdateDB(azureSqlDatabaseProperties)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Log.Info("Async operation not complete or group not found")
			instance.Status.Provisioning = true
			if errup := r.Status().Update(ctx, instance); errup != nil {
				r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
			}
		}

		return errhelp.NewAzureError(err)
	}

	_, err = sdkClient.GetDB(dbName)
	if err != nil {
		return errhelp.NewAzureError(err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	if err = r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	return nil
}

func (r *AzureSqlDatabaseReconciler) deleteExternal(instance *azurev1alpha1.AzureSqlDatabase) error {
	ctx := context.Background()
	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.ObjectMeta.Name

	// create the Go SDK client with relevant info
	sdk := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        server,
		Location:          location,
	}

	r.Log.Info(fmt.Sprintf("deleting external resource: group/%s/server/%s/database/%s"+groupName, server, dbName))
	_, err := sdk.DeleteDB(dbName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Couldn't delete resouce in azure")
		return err
	}
	r.Recorder.Event(instance, corev1.EventTypeNormal, "Deleted", dbName+" deleted")
	return nil
}

func (r *AzureSqlDatabaseReconciler) addFinalizer(instance *azurev1alpha1.AzureSqlDatabase) error {
	helpers.AddFinalizer(instance, azureSQLDatabaseFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", azureSQLDatabaseFinalizerName))
	return nil
}
