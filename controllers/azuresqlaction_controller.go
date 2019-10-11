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
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/go-autorest/autorest/to"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

const AzureSqlActionFinalizerName = "azuresqlaction.finalizers.azure.com"

// AzureSqlActionReconciler reconciles a AzureSqlAction object
type AzureSqlActionReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
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

	if !instance.IsSubmitted() {
		if err := r.reconcileExternal(&instance); err != nil {
			catchIgnorable := []string{
				errhelp.AsyncOpIncompleteError,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catchIgnorable, azerr.Type) {
					log.Info("Requeuing as the async operation is not complete")
					return ctrl.Result{
						Requeue:      true,
						RequeueAfter: time.Second * time.Duration(requeueAfter),
					}, nil
				}
			}
			catchNotIgnorable := []string{
				errhelp.ResourceNotFound,
				errhelp.ResourceGroupNotFoundErrorCode,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catchNotIgnorable, azerr.Type) {
					log.Info("Not requeueing as a specified resource was not found")
					return ctrl.Result{}, nil
				}
			}
			// TODO: Add error handling for other types of errors we might encounter here
			return ctrl.Result{}, fmt.Errorf("error reconciling azuresqlaction in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, corev1.EventTypeNormal, "Provisioned", "AzureSqlAction "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *AzureSqlActionReconciler) reconcileExternal(instance *azurev1alpha1.AzureSqlAction) error {
	ctx := context.Background()
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

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        serverName,
	}

	// Get the Sql Server instance that corresponds to the Server name in the spec for this action
	server, err := sdkClient.GetServer()
	if err != nil {
		if strings.Contains(err.Error(), "ResourceGroupNotFound") {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get instance of AzureSqlServer: Resource group not found")
			r.Log.Info("Error", "Unable to get instance of AzureSqlServer: Resource group not found", err)
			instance.Status.Message = "Resource group not found"
			// write information back to instance
			if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
				r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
			}
			return err
		} else {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to get instance of AzureSqlServer")
			r.Log.Info("Error", "Sql Server instance not found", err)
			instance.Status.Message = "Sql server instance not found"
			// write information back to instance
			if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
				r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
			}
			return err
		}
	}

	sdkClient.Location = *server.Location

	// rollcreds action
	if strings.ToLower(instance.Spec.ActionName) == "rollcreds" {
		azureSqlServerProperties := sql.SQLServerProperties{
			AdministratorLogin:         server.ServerProperties.AdministratorLogin,
			AdministratorLoginPassword: server.ServerProperties.AdministratorLoginPassword,
		}

		// Generate a new password
		newPassword, _ := generateRandomPassword(passwordLength)
		azureSqlServerProperties.AdministratorLoginPassword = to.StringPtr(newPassword)

		if _, err := sdkClient.CreateOrUpdateSQLServer(azureSqlServerProperties); err != nil {
			if !strings.Contains(err.Error(), "not complete") {
				r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to provision or update instance")
				return errhelp.NewAzureError(err)
			}
		} else {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Provisioned", "resource request successfully submitted to Azure")
		}

		// Update the k8s secret
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      serverName,
				Namespace: namespace,
			},
			Type: "Opaque",
		}

		_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
			r.Log.Info("Creating or updating secret with SQL Server credentials")
			secret.Data["password"] = []byte(*azureSqlServerProperties.AdministratorLoginPassword)
			return nil
		})
		if createOrUpdateSecretErr != nil {
			r.Log.Info("Error", "CreateOrUpdateSecretErr", createOrUpdateSecretErr)
			return createOrUpdateSecretErr
		}

		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.Message = "AzureSqlAction completed successfully."

		// write information back to instance
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}

		// write information back to instance
		if updateerr := r.Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}
	} else {
		r.Log.Info("Error", "reconcileExternal", "Unknown action name")

		instance.Status.Message = "Unknown action name error"

		// write information back to instance
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		return errors.New("Unknown action name")
	}

	// Add implementations for other AzureSqlActions here (instance.Spec.ActionName)

	return nil
}

func (r *AzureSqlActionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.AzureSqlAction{}).
		Complete(r)
}
