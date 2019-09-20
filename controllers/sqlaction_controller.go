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
			}
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

	r.Log.Info("Info", "ctx", ctx)
	r.Log.Info("Info", "SqlAction name", name)
	r.Log.Info("Info", "ServerName", serverName)
	r.Log.Info("Info", "ResourceGroup", groupName)
	r.Log.Info("Info", "Namespace", namespace)

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        serverName,
	}

	// write information back to instance
	//instance.Status.Provisioning = true

	sqlActionNamespacedName := types.NamespacedName{Name: name, Namespace: namespace}
	if err := r.Get(ctx, sqlActionNamespacedName, instance); err != nil {
		r.Log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return (err)
	}

	// Get the Sql Server instance that corresponds to the Server name in the spec for this action
	server, err := sdkClient.GetServer(groupName, serverName)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to get instance of SqlServer")
	}

	sdkClient.Location = *server.Location
	sqlServerProperties := sql.SQLServerProperties{
		AdministratorLogin:         server.ServerProperties.AdministratorLogin,
		AdministratorLoginPassword: server.ServerProperties.AdministratorLoginPassword,
	}

	//debugging
	r.Log.Info("Info", "Old Username: ", *sqlServerProperties.AdministratorLogin)
	//r.Log.Info("Info", "Old Password: ", *sqlServerProperties.AdministratorLoginPassword)

	sqlServerProperties.AdministratorLoginPassword = to.StringPtr(RollCreds(16))

	//debugging
	r.Log.Info("Info", "New Username: ", *sqlServerProperties.AdministratorLogin)
	r.Log.Info("Info", "New Password: ", *sqlServerProperties.AdministratorLoginPassword)

	if _, err := sdkClient.CreateOrUpdateSQLServer(sqlServerProperties); err != nil {
		if !strings.Contains(err.Error(), "not complete") {
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to provision or update instance")
			return errhelp.NewAzureError(err)
		}
	} else {
		r.Recorder.Event(instance, "Normal", "Provisioned", "resource request successfully dubmitted to Azure")
	}

	// Update the k8s secret

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
		// innerErr := controllerutil.SetControllerReference(instance, secret, r.Scheme)
		// if innerErr != nil {
		// 	r.Log.Info("Error", "InnerErr", innerErr)
		// 	return innerErr
		// }
		return nil
	})
	if createOrUpdateSecretErr != nil {
		r.Log.Info("Error", "CreateOrUpdateSecretErr", createOrUpdateSecretErr)
		return createOrUpdateSecretErr
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	if err = r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	return nil
}

func (r *SqlActionReconciler) deleteExternal(instance *azurev1.SqlAction) error {
	r.Log.Info("deleteExternal function for SqlAction: do nothing")
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
