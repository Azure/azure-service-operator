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
	"strings"
	"time"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/go-autorest/autorest/to"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

const sqlActionFinalizerName = "sqlaction.finalizers.azure.com"

// SqlActionReconciler reconciles a SqlAction object
type SqlActionReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlactions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlactions/status,verbs=get;update;patch

// Reconcile function runs the actual reconcilation loop of the controller
func (r *SqlActionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqlaction", req.NamespacedName)

	var instance azurev1.SqlAction

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to fetch SqlAction", "err", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, sqlActionFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete SqlAction failed with ", "err", err)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, sqlActionFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, sqlActionFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding sqlaction finalizer failed with ", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		if err := r.reconcileExternal(&instance); err != nil {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				log.Info("Requeuing as the async operation is not complete")
				return ctrl.Result{
					Requeue:      true,
					RequeueAfter: time.Second * time.Duration(requeueAfter),
				}, nil
			} else if errhelp.IsResourceNotFound(err) {
				log.Info("Not requeueing as a specified resource was not found")
				instance.Status.Message = "Resource not found error"
				// write information back to instance
				if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
					r.Recorder.Event(&instance, "Warning", "Failed", "Unable to update instance")
				}
				return ctrl.Result{}, nil
			}
			// TODO: Add error handling for other types of errors we might encounter here
			return ctrl.Result{}, fmt.Errorf("error reconciling sqlaction in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", "SqlAction "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *SqlActionReconciler) addFinalizer(instance *azurev1.SqlAction) error {
	helpers.AddFinalizer(instance, sqlActionFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", sqlActionFinalizerName))
	return nil
}

func (r *SqlActionReconciler) reconcileExternal(instance *azurev1.SqlAction) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	serverName := instance.Spec.ServerName
	groupName := instance.Spec.ResourceGroup
	namespace := instance.Namespace

	if instance.Status.Provisioned != true {

		instance.Status.Provisioning = true
		instance.Status.Message = "SqlAction in progress"

		// write information back to instance
		if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		}

		sdkClient := sql.GoSDKClient{
			Ctx:               ctx,
			ResourceGroupName: groupName,
			ServerName:        serverName,
		}

		sqlActionNamespacedName := types.NamespacedName{Name: name, Namespace: namespace}
		if err := r.Get(ctx, sqlActionNamespacedName, instance); err != nil {
			r.Log.Info("Unable to retrieve SqlAction resource", "err", err.Error())
			// we'll ignore not-found errors, since they can't be fixed by an immediate
			// requeue (we'll need to wait for a new notification), and we can get them
			// on deleted requests.
			return (err)
		}

		// Get the Sql Server instance that corresponds to the Server name in the spec for this action
		server, err := sdkClient.GetServer()
		if err != nil {
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to get instance of SqlServer")
			r.Log.Info("Error", "Sql Server instance not found", err)
			instance.Status.Message = "Sql Server instance not found"
			return err
		}

		sdkClient.Location = *server.Location

		// rollcreds action
		if instance.Spec.ActionName == "rollcreds" {
			sqlServerProperties := sql.SQLServerProperties{
				AdministratorLogin:         server.ServerProperties.AdministratorLogin,
				AdministratorLoginPassword: server.ServerProperties.AdministratorLoginPassword,
			}

			// Generate a new password
			sqlServerProperties.AdministratorLoginPassword = to.StringPtr(RollCreds(16))

			if _, err := sdkClient.CreateOrUpdateSQLServer(sqlServerProperties); err != nil {
				if !strings.Contains(err.Error(), "not complete") {
					r.Recorder.Event(instance, "Warning", "Failed", "Unable to provision or update instance")
					return errhelp.NewAzureError(err)
				}
			} else {
				r.Recorder.Event(instance, "Normal", "Provisioned", "resource request successfully submitted to Azure")
			}

			// Update the k8s secret
			// TODO: Question: Do we need to define the entire secret as below to send the CreateOrUpdate request?
			secret := &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      serverName,
					Namespace: namespace,
				},
				Data: map[string][]byte{
					"username":           []byte(*sqlServerProperties.AdministratorLogin),
					"password":           []byte(*sqlServerProperties.AdministratorLoginPassword),
					"sqlservernamespace": []byte(instance.Namespace),
					"sqlservername":      []byte(name),
				},
				Type: "Opaque",
			}

			_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
				r.Log.Info("mutating secret bundle")
				secret.Data["password"] = []byte(*sqlServerProperties.AdministratorLoginPassword)
				return nil
			})
			if createOrUpdateSecretErr != nil {
				r.Log.Info("Error", "CreateOrUpdateSecretErr", createOrUpdateSecretErr)
				return createOrUpdateSecretErr
			}

			instance.Status.Provisioning = false
			instance.Status.Provisioned = true
			instance.Status.Message = "SqlAction completed successfully."

			// write information back to instance
			if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
				r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
			}

			// write information back to instance
			if updateerr := r.Update(ctx, instance); updateerr != nil {
				r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
			}
		}

		// Add implementations for other SqlActions here (instance.Spec.ActionName)
	}

	return nil
}

func (r *SqlActionReconciler) deleteExternal(instance *azurev1.SqlAction) error {
	r.Log.Info("deleteExternal function for SqlAction: Do nothing")
	return nil
}

func (r *SqlActionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlAction{}).
		Complete(r)
}

func RollCreds(length int) string {
	// logic to generate new password and return it
	newPassword := generateRandomString(length)
	return newPassword
}
