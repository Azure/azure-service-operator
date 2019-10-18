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
	"strings"
	"time"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
)

const azureSQLFailoverGroupFinalizerName = "azuresqlfailovergroup.finalizers.azure.com"

// AzureSQLFailoverGroupReconciler reconciles a AzureSQLFailoverGroup object
type AzureSQLFailoverGroupReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfailovergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlfailovergroups/status,verbs=get;update;patch

func (r *AzureSQLFailoverGroupReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("azuresqlfailovergroup", req.NamespacedName)
	var instance azurev1alpha1.AzureSQLFailoverGroup

	defer func() {
		if err := r.Status().Update(ctx, &instance); err != nil {
			r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
	}()

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql failover group resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, azureSQLFailoverGroupFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				catch := []string{
					errhelp.AsyncOpIncompleteError,
				}
				if azerr, ok := err.(*errhelp.AzureError); ok {
					if helpers.ContainsString(catch, azerr.Type) {
						log.Info("Got ignorable error", "type", azerr.Type)
						return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
					}
				}
				msg := fmt.Sprintf("Delete AzureFailoverGroup failed with %s", err.Error())
				log.Info(msg)
				instance.Status.Message = msg
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, azureSQLFailoverGroupFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, azureSQLFailoverGroupFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			msg := fmt.Sprintf("Adding AzureFailoverGroup finalizer failed with error %s", err.Error())
			log.Info(msg)
			instance.Status.Message = msg
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, v1.EventTypeNormal, "Submitting", "starting resource reconciliation")
		// TODO: Add error handling for cases where username or password are invalid:
		// https://docs.microsoft.com/en-us/rest/api/sql/servers/createorupdate#response
		if err := r.reconcileExternal(&instance); err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
				errhelp.InvalidServerName,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					if azerr.Type == errhelp.InvalidServerName {
						msg := "Invalid Server Name"
						r.Recorder.Event(&instance, v1.EventTypeWarning, "Failed", msg)
						instance.Status.Message = msg
						return ctrl.Result{Requeue: false}, nil
					}
					msg := fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
					log.Info(msg)
					instance.Status.Message = msg
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling sql server in azure: %v", err)
		}
		// give azure some time to catch up
		log.Info("waiting for provision to take effect")
		return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	if err := r.verifyExternal(&instance); err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
			errhelp.AsyncOpIncompleteError,
		}
		if azerr, ok := err.(*errhelp.AzureError); ok {
			if helpers.ContainsString(catch, azerr.Type) {
				msg := fmt.Sprintf("Got ignorable error type: %s", azerr.Type)
				log.Info(msg)
				instance.Status.Message = msg
				return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
			}
		}
		return ctrl.Result{}, fmt.Errorf("error verifying sql failover group in azure: %v", err)
	}

	r.Recorder.Event(&instance, v1.EventTypeNormal, "Provisioned", "azurefailovergroup "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *AzureSQLFailoverGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSQLFailoverGroup{}).
		Complete(r)
}

func (r *AzureSQLFailoverGroupReconciler) reconcileExternal(instance *azurev1alpha1.AzureSQLFailoverGroup) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	servername := instance.Spec.Server
	failoverPolicy := instance.Spec.FailoverPolicy
	failoverGracePeriod := instance.Spec.FailoverGracePeriod
	secServerName := instance.Spec.SecondaryServerName
	secServerResourceGroup := instance.Spec.SecondaryServerResourceGroup
	databaseList := instance.Spec.DatabaseList

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        name,
		Location:          location,
	}

	// create the sql failover group
	instance.Status.Provisioning = true

	if _, err := sdkClient.CreateOrUpdateSQLServer(); err != nil {
		if !strings.Contains(err.Error(), "not complete") {
			msg := fmt.Sprintf("CreateOrUpdateSQLServer not complete: %v", err)
			instance.Status.Message = msg
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to provision or update instance")
			return errhelp.NewAzureError(err)
		}
	} else {
		msg := "Resource request successfully submitted to Azure"
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeNormal, "Provisioned", msg)
	}

	_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
		r.Log.Info("Creating or updating secret with SQL Server credentials")
		innerErr := controllerutil.SetControllerReference(instance, secret, r.Scheme)
		if innerErr != nil {
			return innerErr
		}
		return nil
	})
	if createOrUpdateSecretErr != nil {
		return createOrUpdateSecretErr
	}

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	return nil
}

func (r *AzureSQLFailoverGroupReconciler) verifyExternal(instance *azurev1alpha1.AzureSQLFailoverGroup) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        name,
		Location:          location,
	}

	serv, err := sdkClient.GetServer()
	if err != nil {
		azerr := errhelp.NewAzureError(err).(*errhelp.AzureError)
		if azerr.Type != errhelp.ResourceNotFound {
			return azerr
		}

		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = *serv.State
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Checking", fmt.Sprintf("instance in %s state", instance.Status.State))

	if instance.Status.State == "Ready" {
		msg := "AzureSQLFailoverGroup successfully provisioned"
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = msg
	}

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		return updateerr
	}

	return errhelp.NewAzureError(err)
}

func (r *AzureSQLFailoverGroupReconciler) deleteExternal(instance *azurev1alpha1.AzureSQLFailoverGroup) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	location := instance.Spec.Location

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        name,
		Location:          location,
	}

	_, err := sdkClient.DeleteSQLServer()
	if err != nil {
		msg := fmt.Sprintf("Couldn't delete resource in Azure: %v", err)
		instance.Status.Message = msg
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
		return errhelp.NewAzureError(err)
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", name+" deleted")
	return nil
}
