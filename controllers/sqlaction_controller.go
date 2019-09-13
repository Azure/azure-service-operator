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

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	"github.com/Azure/go-autorest/autorest/to"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

// SqlActionReconciler reconciles a SqlAction object
type SqlActionReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlactions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=sqlactions/status,verbs=get;update;patch

func (r *SqlActionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sqlaction", req.NamespacedName)

	// your logic here
	// Debugging
	r.Log.Info("starting sqlaction reconciler")
	var instance azurev1.SqlServer
	var action azurev1.SqlAction

	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	sdkClient := sql.GoSDKClient{
		Ctx:               ctx,
		ResourceGroupName: groupName,
		ServerName:        name,
		Location:          location,
	}

	if err := r.Get(ctx, req.NamespacedName, &action); err != nil {
		log.Info("Unable to retrieve sql-action resource", "err", err.Error())
		return ctrl.Result{}, err
	}

	sqlServerProperties := sql.SQLServerProperties{
		AdministratorLogin:         to.StringPtr(""),
		AdministratorLoginPassword: to.StringPtr(""),
	}

	// TODO: Find a central location for this const or make it configurable by user
	const passwordLength = 16

	if action.Spec.ActionName == "rollcreds" {
		// Check for and get secret containing creds
		var checkForSecretsErr error
		secret := &v1.Secret{}

		checkForSecretsErr = r.Get(context.Background(), types.NamespacedName{Name: instance.ObjectMeta.Name, Namespace: instance.Namespace}, secret)

		if checkForSecretsErr == nil {
			// Secret exists
			// Update sqlserverproperties object with secret admin login + new password
			// Testing
			r.Log.Info("Old username:")
			r.Log.Info(string(secret.Data["username"]))
			sqlServerProperties.AdministratorLogin = to.StringPtr(string(RollCreds(8))) // hard-coding for testing
			sqlServerProperties.AdministratorLoginPassword = to.StringPtr(RollCreds(passwordLength))
			r.Log.Info("New username:")
			r.Log.Info(*sqlServerProperties.AdministratorLogin)

			// Update sql server with new sqlserverproperties
			instance.Status.Provisioning = true
			_, err := sdkClient.CreateOrUpdateSQLServer(sqlServerProperties)
			if err != nil {
				r.Recorder.Event(&instance, "Warning", "Failed", "Unable to provision or update instance")
				instance.Status.Provisioning = false
				err = errhelp.NewAzureError(err)
			} else {
				r.Recorder.Event(&instance, "Normal", "Provisioned", "resource request successfully submitted to Azure")
			}

			// write information back to instance
			if updateerr := r.Status().Update(ctx, &instance); updateerr != nil {
				r.Recorder.Event(&instance, "Warning", "Failed", "Unable to update instance")
			}

			// Update secret with new password
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

			_, createOrUpdateSecretErr := controllerutil.CreateOrUpdate(context.Background(), r.Client, secret, func() error {
				r.Log.Info("mutating secret bundle")
				innerErr := controllerutil.SetControllerReference(&instance, secret, r.Scheme)
				if innerErr != nil {
					return innerErr
				}
				return nil
			})
			if createOrUpdateSecretErr != nil {
				r.Log.Info("Could not update secret")
			}
		} else {
			// Error: Secret did not exist
			// TODO: How do we want to handle this?
		}
	}

	return ctrl.Result{}, nil
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
