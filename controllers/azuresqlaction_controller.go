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
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/go-autorest/autorest/to"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

const AzureSqlActionFinalizerName = "azuresqlaction.finalizers.azure.com"

// AzureSqlActionReconciler reconciles a AzureSqlAction object
type AzureSqlActionReconciler struct {
	client.Client
	Log                   logr.Logger
	Recorder              record.EventRecorder
	Scheme                *runtime.Scheme
	AzureSqlServerManager sql.SqlServerManager
	SecretClient          secrets.SecretClient
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlactions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlactions/status,verbs=get;update;patch

// Reconcile function runs the actual reconcilation loop of the controller
func (r *AzureSqlActionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("azuresqlaction", req.NamespacedName)

	var instance azurev1alpha1.AzureSqlAction

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to fetch AzureSqlAction", "err", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	defer func() {
		if !helpers.IsBeingDeleted(&instance) {
			if err := r.Status().Update(ctx, &instance); err != nil {
				r.Recorder.Event(&instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
			}
		}
	}()

	if !instance.IsSubmitted() {
		if err := r.reconcileExternal(ctx, &instance); err != nil {
			catchIgnorable := []string{
				errhelp.AsyncOpIncompleteError,
			}
			catchNotIgnorable := []string{
				errhelp.ResourceNotFound,
				errhelp.ResourceGroupNotFoundErrorCode,
			}
			azerr := errhelp.NewAzureErrorAzureError(err)

			// handle catchable errors
			if helpers.ContainsString(catchIgnorable, azerr.Type) {
				log.Info("Requeuing as the async operation is not complete")
				return ctrl.Result{
					RequeueAfter: time.Second * time.Duration(requeueAfter),
				}, nil
			}

			// handle errors we can't ignore
			if helpers.ContainsString(catchNotIgnorable, azerr.Type) {
				log.Info("Not requeueing as a specified resource was not found")
				return ctrl.Result{}, nil
			}

			// TODO: Add error handling for other types of errors we might encounter here
			return ctrl.Result{}, fmt.Errorf("error reconciling azuresqlaction in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, corev1.EventTypeNormal, "Provisioned", "AzureSqlAction "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *AzureSqlActionReconciler) reconcileExternal(ctx context.Context, instance *azurev1alpha1.AzureSqlAction) error {
	serverName := instance.Spec.ServerName
	groupName := instance.Spec.ResourceGroup
	namespace := instance.Namespace

	instance.Status.Provisioning = true
	instance.Status.Provisioned = false
	instance.Status.Message = "AzureSqlAction in progress"
	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	//get owner instance of AzureSqlServer
	r.Recorder.Event(instance, corev1.EventTypeNormal, "UpdatingOwner", "Updating owner AzureSqlServer instance")
	var ownerInstance azurev1alpha1.AzureSqlServer
	azureSqlServerNamespacedName := types.NamespacedName{Name: serverName, Namespace: instance.Namespace}
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
	if err := r.Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	// Get the Sql Server instance that corresponds to the Server name in the spec for this action
	server, err := r.AzureSqlServerManager.GetServer(ctx, groupName, serverName)
	if err != nil {
		if strings.Contains(err.Error(), errhelp.ResourceGroupNotFoundErrorCode) {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get instance of AzureSqlServer: Resource group not found")
			r.Log.Info("Error", "Unable to get instance of AzureSqlServer: Resource group not found", err)
			instance.Status.Message = "Resource group not found"

			return err
		}

		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get instance of AzureSqlServer")
		r.Log.Info("Error", "Sql Server instance not found", err)
		instance.Status.Message = "Sql server instance not found"

		return err
	}

	// rollcreds action
	if strings.ToLower(instance.Spec.ActionName) == "rollcreds" {
		azureSqlServerProperties := sql.SQLServerProperties{
			AdministratorLogin:         server.ServerProperties.AdministratorLogin,
			AdministratorLoginPassword: server.ServerProperties.AdministratorLoginPassword,
		}

		// Generate a new password
		newPassword, _ := generateRandomPassword(passwordLength)
		azureSqlServerProperties.AdministratorLoginPassword = to.StringPtr(newPassword)

		if _, err := r.AzureSqlServerManager.CreateOrUpdateSQLServer(ctx, groupName, *server.Location, serverName, azureSqlServerProperties); err != nil {
			if !strings.Contains(err.Error(), "not complete") {
				r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to provision or update instance")
				return err
			}
		} else {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Provisioned", "resource request successfully submitted to Azure")
		}

		key := types.NamespacedName{Name: serverName, Namespace: namespace}
		data, err := r.SecretClient.Get(ctx, key)
		if err != nil {
			return err
		}

		data["password"] = []byte(*azureSqlServerProperties.AdministratorLoginPassword)
		err = r.SecretClient.Upsert(
			ctx,
			key,
			data,
			secrets.WithOwner(&ownerInstance),
			secrets.WithScheme(r.Scheme),
		)
		if err != nil {
			return err
		}

		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.Message = "AzureSqlAction completed successfully."
		return nil
	}

	// Add implementations for other AzureSqlActions here (instance.Spec.ActionName)

	r.Log.Info("Error", "reconcileExternal", "Unknown action name")
	instance.Status.Message = "Unknown action name error"

	return errors.New("Unknown AzureSqlAction name: " + instance.Spec.ActionName)

}

func (r *AzureSqlActionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlAction{}).
		Complete(r)
}
