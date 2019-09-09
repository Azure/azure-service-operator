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
	"math/rand"
	"strings"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// SqlServerReconciler reconciles a SqlServer object
type SqlServerReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlservers/status,verbs=get;update;patch

func (r *SqlServerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqlserver", req.NamespacedName)

	// your logic here
	var instance azurev1.SqlServer

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve sql-server resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, SQLServerFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete SqlServer failed with ", "error", err.Error())
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, SQLServerFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, SQLServerFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding SqlServer finalizer failed with ", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		r.Recorder.Event(&instance, "Normal", "Submitting", "starting resource reconciliation")
		if err := r.reconcileExternal(&instance); err != nil {
			if strings.Contains(err.Error(), "asynchronous operation has not completed") {
				r.Recorder.Event(&instance, "Normal", "Provisioning", "async op still running")
				return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling sql server in azure: %v", err)
		}
		// if the request was just sent to azure, the resource probably isn't ready yet
		return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
	}

	if err := r.verifyExternal(&instance); err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		if azerr, ok := err.(*errhelp.AzureError); ok {
			if helpers.ContainsString(catch, azerr.Type) {
				log.Info("Got ignorable error", "type", azerr.Type)
				return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
			}
		}
		return ctrl.Result{}, fmt.Errorf("error verifying sql server in azure: %v", err)
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", "sqlserver "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *SqlServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.SqlServer{}).
		Complete(r)
}

func (r *SqlServerReconciler) reconcileExternal(instance *azurev1.SqlServer) error {
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

	sqlServerProperties := sql.SQLServerProperties{
		AdministratorLogin:         to.StringPtr(generateRandomString(8)),
		AdministratorLoginPassword: to.StringPtr(generateRandomString(16)),
	}

	// Check to see if secret already exists for admin username/password
	var checkForSecretsErr error
	secret := &v1.Secret{}

	checkForSecretsErr = r.Get(context.Background(), types.NamespacedName{Name: name, Namespace: instance.Namespace}, secret)

	// If secret doesn't exist, generate creds
	// Note: sql server enforces password policy.  Details can be found here:
	// https://docs.microsoft.com/en-us/sql/relational-databases/security/password-policy?view=sql-server-2017
	if checkForSecretsErr != nil {
		r.Log.Info("secret does not exist, using randomly generated creds")
		secret = &v1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "apps/v1beta1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: instance.Namespace,
			},
			Data: map[string][]byte{
				"username":           []byte(*sqlServerProperties.AdministratorLogin),
				"password":           []byte(*sqlServerProperties.AdministratorLoginPassword),
				"sqlservernamespace": []byte(instance.Namespace),
				"sqlservername":      []byte(name),
			},
			Type: "Opaque",
		}
	} else {
		r.Log.Info("secret already exists, pulling creds now")
		sqlServerProperties.AdministratorLogin = to.StringPtr(string(secret.Data["username"]))
		sqlServerProperties.AdministratorLoginPassword = to.StringPtr(string(secret.Data["password"]))
	}

	_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
		r.Log.Info("mutating secret bundle")
		innerErr := controllerutil.SetControllerReference(instance, secret, r.Scheme)
		if innerErr != nil {
			return innerErr
		}
		return nil
	})
	if createOrUpdateSecretErr != nil {
		return createOrUpdateSecretErr
	}

	// create the sql server
	instance.Status.Provisioning = true
	_, err := sdkClient.CreateOrUpdateSQLServer(sqlServerProperties)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to provision or update instance")
		instance.Status.Provisioning = false
		err = errhelp.NewAzureError(err)
	} else {
		r.Recorder.Event(instance, "Normal", "Provisioned", "resource request successfully dubmitted to Azure")
	}

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	return err
}

func (r *SqlServerReconciler) verifyExternal(instance *azurev1.SqlServer) error {
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

	serv, err := sdkClient.GetServer(groupName, name)
	if err != nil {
		azerr := errhelp.NewAzureError(err).(*errhelp.AzureError)
		if azerr.Type != errhelp.ResourceNotFound {
			return azerr
		}

		instance.Status.State = "NotReady"
	} else {
		instance.Status.State = *serv.State
	}

	r.Recorder.Event(instance, "Normal", "Checking", fmt.Sprintf("instance in %s state", instance.Status.State))

	if instance.Status.State == "Ready" {

		if instance.Spec.AllowAzureServiceAccess == true {
			// Add firewall rule to allow azure service access
			_, err := sdkClient.CreateOrUpdateSQLFirewallRule("AllowAzureAccess", "0.0.0.0", "0.0.0.0")
			if err != nil {
				r.Recorder.Event(instance, "Warning", "Failed", "Unable to add firewall rule to SQL server")
				return errhelp.NewAzureError(err)
			}
		}
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
	}

	// write information back to instance
	if updateerr := r.Status().Update(ctx, instance); updateerr != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		return updateerr
	}

	return errhelp.NewAzureError(err)
}

func (r *SqlServerReconciler) deleteExternal(instance *azurev1.SqlServer) error {
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
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resouce in azure")
		return err
	}

	r.Recorder.Event(instance, "Normal", "Deleted", name+" deleted")
	return nil
}

// helper function to generate username/password for secrets
func generateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	const characterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+-=<>"

	// TODO: add logic to enforce password policy rules for sql server
	b := make([]byte, n)
	for i := range b {
		b[i] = characterBytes[rand.Intn(len(characterBytes))]
	}

	return string(b)
}
